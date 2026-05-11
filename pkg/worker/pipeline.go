package worker

import (
	"context"
	"fmt"
	"strings"

	"github.com/insajin/autopus-adk/pkg/worker/a2a"
	"github.com/insajin/autopus-adk/pkg/worker/adapter"
	"github.com/insajin/autopus-adk/pkg/worker/budget"
	"github.com/insajin/autopus-adk/pkg/worker/compress"
	"github.com/insajin/autopus-adk/pkg/worker/routing"
)

type Phase string

const (
	PhasePlanner  Phase = "planner"
	PhaseExecutor Phase = "executor"
	PhaseTester   Phase = "tester"
	PhaseReviewer Phase = "reviewer"
)

type PhaseResult struct {
	Phase      Phase
	Output     string
	CostUSD    float64
	DurationMS int64
	SessionID  string
	ToolCalls  int
}

var defaultPipelinePhases = []Phase{PhasePlanner, PhaseExecutor, PhaseTester, PhaseReviewer}

type PipelineExecutor struct {
	provider             adapter.ProviderAdapter
	mcpConfig            string
	workDir              string
	envVars              map[string]string
	phaseInstructions    map[Phase]string
	phasePromptTemplates map[Phase]string
	allocator            *budget.PhaseAllocator
	iterationBudget      *budget.IterationBudget
	compressor           compress.ContextCompressor
	router               *routing.Router
}

func NewPipelineExecutor(provider adapter.ProviderAdapter, mcpConfig, workDir string) *PipelineExecutor {
	return &PipelineExecutor{provider: provider, mcpConfig: mcpConfig, workDir: workDir}
}

func (pe *PipelineExecutor) SetBudget(total int, alloc budget.PhaseAllocation) {
	pe.allocator = budget.NewPhaseAllocator(total, alloc)
}

func (pe *PipelineExecutor) SetIterationBudget(ib budget.IterationBudget) {
	pe.iterationBudget = &ib
	pe.allocator = budget.NewPhaseAllocator(ib.Limit, budget.DefaultAllocation())
}

func (pe *PipelineExecutor) SetCompressor(c compress.ContextCompressor) { pe.compressor = c }
func (pe *PipelineExecutor) SetEnvVars(ev map[string]string) { pe.envVars = ev }
func (pe *PipelineExecutor) SetPhaseInstructions(i map[Phase]string) { pe.phaseInstructions = i }
func (pe *PipelineExecutor) SetPhasePromptTemplates(t map[Phase]string) { pe.phasePromptTemplates = t }
func (pe *PipelineExecutor) SetRouter(r *routing.Router) { pe.router = r }

func (pe *PipelineExecutor) Execute(ctx context.Context, taskID, prompt string) (adapter.TaskResult, error) {
	return pe.ExecuteWithPlan(ctx, taskID, prompt, "", nil)
}

func (pe *PipelineExecutor) ExecuteWithPlan(ctx context.Context, taskID, prompt, model string, phases []Phase) (adapter.TaskResult, error) {
	routedModel := model
	if routedModel == "" && pe.router != nil && !a2a.SignedControlPlaneEnforced() {
		routedModel = pe.router.Route(pe.provider.Name(), prompt)
	}

	var results []PhaseResult
	var totalCost float64
	var totalDuration int64
	prevOutput := prompt
	normalizedPhases := normalizePhasePlan(phases)
	if len(normalizedPhases) == 0 { return adapter.TaskResult{}, fmt.Errorf("missing plan") }

	for _, phase := range normalizedPhases {
		select {
		case <-ctx.Done(): return adapter.TaskResult{}, ctx.Err()
		default:
		}

		phasePrompt, _ := pe.phasePrompt(phase, prevOutput)
		pr, err := pe.runPhase(ctx, taskID, phase, phasePrompt, routedModel)
		if err != nil { return adapter.TaskResult{}, err }

		if pe.allocator != nil { pe.allocator.CompletePhase(string(phase), pr.ToolCalls) }
		results = append(results, pr)
		totalCost += pr.CostUSD
		totalDuration += pr.DurationMS

		if pe.compressor != nil {
			prevOutput = pe.compressor.Compress(string(phase), pr.Output, pe.provider.Name())
		} else {
			prevOutput = pr.Output
		}
	}
	return pe.aggregateResults(results, totalCost, totalDuration), nil
}

func (pe *PipelineExecutor) phasePrompt(phase Phase, input string) (string, error) {
	if template, ok := pe.phasePromptTemplates[phase]; ok && strings.TrimSpace(template) != "" {
		return strings.ReplaceAll(template, "{{input}}", input), nil
	}
	if instruction, ok := pe.phaseInstructions[phase]; ok && strings.TrimSpace(instruction) != "" {
		return fmt.Sprintf("%s\n\n%s", instruction, input), nil
	}
	switch phase {
	case PhasePlanner: return pe.plannerPrompt(input), nil
	case PhaseExecutor: return pe.executorPrompt(input), nil
	case PhaseTester: return pe.testerPrompt(input), nil
	case PhaseReviewer: return pe.reviewerPrompt(input), nil
	default: return "", fmt.Errorf("unsupported phase")
	}
}
