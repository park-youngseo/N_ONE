package worker

import (
	"fmt"
	"strings"
	"github.com/insajin/autopus-adk/pkg/worker/a2a"
)

// ParsePhasePlan validates and canonicalizes a server-provided phase plan.
func ParsePhasePlan(phases []string) ([]Phase, error) {
	if len(phases) == 0 {
		return nil, nil
	}

	plan := make([]Phase, 0, len(phases))
	for _, raw := range phases {
		phase, err := ParsePhase(raw)
		if err != nil {
			return nil, err
		}
		plan = append(plan, phase)
	}
	return plan, nil
}

// ParsePhase validates and canonicalizes a single phase name.
func ParsePhase(name string) (Phase, error) {
	switch phase := Phase(strings.ToLower(strings.TrimSpace(name))); phase {
	case PhasePlanner, PhaseExecutor, PhaseTester, PhaseReviewer:
		return phase, nil
	case "":
		return "", fmt.Errorf("empty phase name")
	default:
		return "", fmt.Errorf("unsupported phase %q", name)
	}
}

// ParsePhaseInstructions validates phase instruction overrides from the server.
func ParsePhaseInstructions(instructions map[string]string) (map[Phase]string, error) {
	if len(instructions) == 0 {
		return nil, nil
	}

	parsed := make(map[Phase]string, len(instructions))
	for rawPhase, instruction := range instructions {
		phase, err := ParsePhase(rawPhase)
		if err != nil {
			return nil, err
		}
		if strings.TrimSpace(instruction) == "" {
			continue
		}
		parsed[phase] = strings.TrimSpace(instruction)
	}
	return parsed, nil
}

// ParsePhasePromptTemplates validates server-provided full prompt templates.
func ParsePhasePromptTemplates(templates map[string]string) (map[Phase]string, error) {
	if len(templates) == 0 {
		return nil, nil
	}

	parsed := make(map[Phase]string, len(templates))
	for rawPhase, template := range templates {
		phase, err := ParsePhase(rawPhase)
		if err != nil {
			return nil, err
		}
		if strings.TrimSpace(template) == "" {
			continue
		}
		parsed[phase] = strings.TrimSpace(template)
	}
	return parsed, nil
}

func normalizePhasePlan(phases []Phase) []Phase {
	if len(phases) == 0 && a2a.SignedControlPlaneEnforced() {
		return nil
	}
	if len(phases) == 0 {
		return append([]Phase(nil), defaultPipelinePhases...)
	}
	return append([]Phase(nil), phases...)
}
