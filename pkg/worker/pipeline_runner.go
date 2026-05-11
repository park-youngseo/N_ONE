package worker

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/insajin/autopus-adk/pkg/worker/adapter"
	"github.com/insajin/autopus-adk/pkg/worker/budget"
	"github.com/insajin/autopus-adk/pkg/worker/security"
	"github.com/insajin/autopus-adk/pkg/worker/stream"
)

// runPhase spawns a single subprocess for the given phase.
func (pe *PipelineExecutor) runPhase(ctx context.Context, taskID string, phase Phase, prompt, model string) (PhaseResult, error) {
	sessionID := fmt.Sprintf("pipeline-%s-%s", taskID, phase)
	taskCfg := adapter.TaskConfig{
		TaskID:    fmt.Sprintf("%s-%s", taskID, phase),
		SessionID: sessionID,
		Prompt:    prompt,
		MCPConfig: pe.mcpConfig,
		WorkDir:   pe.workDir,
		Model:     model,
	}
	if len(pe.envVars) > 0 {
		taskCfg.EnvVars = pe.envVars
	}

	cmd := pe.provider.BuildCommand(ctx, taskCfg)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	if err := cmd.Start(); err != nil {
		return PhaseResult{}, fmt.Errorf("start subprocess: %w", err)
	}

	var emergencyStop *security.EmergencyStop
	phaseBudget := pe.phaseIterationBudget(phase)
	if phaseBudget != nil {
		emergencyStop = security.NewEmergencyStop()
		emergencyStop.SetProcess(cmd)
		defer emergencyStop.ClearProcess()
	}

	go func() {
		defer stdin.Close()
		_, _ = io.Copy(stdin, strings.NewReader(prompt))
	}()

	result, parseErr := pe.parsePhaseStream(stdout, phase, phaseBudget, emergencyStop)
	waitErr := cmd.Wait()

	if parseErr != nil {
		return PhaseResult{}, fmt.Errorf("parse stream: %w", parseErr)
	}
	if waitErr != nil && result.Output == "" {
		return PhaseResult{}, fmt.Errorf("subprocess exit: %w", waitErr)
	}

	return result, nil
}

func (pe *PipelineExecutor) parsePhaseStream(r io.Reader, phase Phase, pb *budget.IterationBudget, es *security.EmergencyStop) (PhaseResult, error) {
	scanner := bufio.NewScanner(r)
	var result PhaseResult
	result.Phase = phase
	hasResult := false
	var counter *budget.Counter
	if pb != nil { counter = budget.NewCounter(*pb) }

	for scanner.Scan() {
		evt, err := pe.provider.ParseEvent(scanner.Bytes())
		if err != nil { continue }

		if evt.Type == stream.EventToolCall || evt.Type == "tool_use" {
			result.ToolCalls++
			if counter != nil {
				if state := counter.Increment(); state.Level == budget.LevelExhausted && es != nil {
					_ = es.Stop("budget_exceeded")
					return result, fmt.Errorf("budget exceeded")
				}
			}
		}

		if evt.Type == "result" {
			tr := pe.provider.ExtractResult(evt)
			result.Output, result.CostUSD, result.DurationMS, result.SessionID = tr.Output, tr.CostUSD, tr.DurationMS, tr.SessionID
			hasResult = true
		}
	}
	if !hasResult { return PhaseResult{}, fmt.Errorf("no result") }
	return result, nil
}

func (pe *PipelineExecutor) phaseIterationBudget(phase Phase) *budget.IterationBudget {
	if pe.iterationBudget == nil || pe.allocator == nil { return nil }
	pb := *pe.iterationBudget
	pb.Limit = pe.allocator.PhaseLimit(string(phase))
	if pb.Limit <= 0 { return nil }
	return &pb
}

func (pe *PipelineExecutor) aggregateResults(results []PhaseResult, cost float64, duration int64) adapter.TaskResult {
	var sb strings.Builder
	for _, r := range results {
		fmt.Fprintf(&sb, "## Phase: %s\n\n%s\n\n", r.Phase, r.Output)
	}
	return adapter.TaskResult{CostUSD: cost, DurationMS: duration, Output: sb.String()}
}

func (pe *PipelineExecutor) plannerPrompt(i string) string { return fmt.Sprintf("PLANNER: %s", i) }
func (pe *PipelineExecutor) executorPrompt(i string) string { return fmt.Sprintf("EXECUTOR: %s", i) }
func (pe *PipelineExecutor) testerPrompt(i string) string { return fmt.Sprintf("TESTER: %s", i) }
func (pe *PipelineExecutor) reviewerPrompt(i string) string { return fmt.Sprintf("REVIEWER: %s", i) }
