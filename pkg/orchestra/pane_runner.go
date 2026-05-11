package orchestra

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/insajin/autopus-adk/pkg/terminal"
)

/*
[Global Connectivity Table]
| Node | Responsibility | Connectivity |
| :--- | :--- | :--- |
| pane_runner.go | Main execution flow for pane-based orchestration | Uses pane_runner_io.go for I/O and command building |
| pane_runner_io.go | Low-level I/O, sentinel detection, command string building | Helper for pane_runner.go |
| pane_shell.go | Shell escaping and argument sanitization | Utility for pane_runner.go |
*/

// paneInfo tracks a provider's pane and output file.
type paneInfo struct {
	paneID     terminal.PaneID
	outputFile string
	provider   ProviderConfig
	skipWait   bool // true when SendCommand failed — skip sentinel wait
}

// RunPaneOrchestra runs orchestration using terminal panes when available.
// Falls back to RunOrchestra for plain terminals or when pane creation fails.
func RunPaneOrchestra(ctx context.Context, cfg OrchestraConfig) (*OrchestraResult, error) {
	if cfg.Terminal == nil || cfg.Terminal.Name() == "plain" {
		return RunOrchestra(ctx, cfg)
	}
	if cfg.Interactive {
		return RunInteractivePaneOrchestra(ctx, cfg)
	}
	if cfg.Strategy == StrategyRelay {
		return runRelayPaneOrchestra(ctx, cfg)
	}

	start := time.Now()
	timeout := cfg.TimeoutSeconds
	if timeout <= 0 {
		timeout = 120
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	panes, failed, err := splitProviderPanes(timeoutCtx, cfg)
	if err != nil {
		return runFallback(ctx, cfg)
	}
	defer cleanupPanes(cfg.Terminal, panes)

	sendFailed := sendPaneCommands(timeoutCtx, cfg, panes)
	failed = append(failed, sendFailed...)

	responses, waitFailed := collectPaneResults(timeoutCtx, panes, start)
	failed = append(failed, waitFailed...)

	total := time.Since(start)
	merged, summary := mergeByStrategy(cfg.Strategy, responses, cfg)
	if merged == "" {
		merged = fmt.Sprintf("[pane mode] %d providers executed", len(responses))
	}

	return &OrchestraResult{
		Strategy:        cfg.Strategy,
		Responses:       responses,
		Merged:          merged,
		Duration:        total,
		Summary:         summary,
		FailedProviders: failed,
	}, nil
}

// splitProviderPanes creates a pane and temp file for each provider.
func splitProviderPanes(ctx context.Context, cfg OrchestraConfig) ([]paneInfo, []FailedProvider, error) {
	surfCreator, hasSurface := cfg.Terminal.(terminal.SurfaceCreator)
	panes := make([]paneInfo, 0, len(cfg.Providers))
	for _, p := range cfg.Providers {
		var paneID terminal.PaneID
		var err error
		if hasSurface {
			paneID, err = surfCreator.CreateSurface(ctx)
		} else {
			paneID, err = cfg.Terminal.SplitPane(ctx, terminal.Horizontal)
		}
		if err != nil {
			cleanupPanes(cfg.Terminal, panes)
			return nil, nil, err
		}
		safeName := sanitizeProviderName(p.Name)
		tmpFile, err := os.CreateTemp("", "autopus-orch-"+safeName+"-")
		if err != nil {
			cleanupPanes(cfg.Terminal, panes)
			return nil, nil, err
		}
		tmpFile.Close()
		panes = append(panes, paneInfo{paneID: paneID, outputFile: tmpFile.Name(), provider: p})
	}
	return panes, nil, nil
}

// sendPaneCommands sends the interactive command to each pane.
func sendPaneCommands(ctx context.Context, cfg OrchestraConfig, panes []paneInfo) []FailedProvider {
	var failed []FailedProvider
	for i, pi := range panes {
		cmd := buildPaneCommand(pi.provider, cfg.Prompt, pi.outputFile) + "\n"
		if err := cfg.Terminal.SendCommand(ctx, pi.paneID, cmd); err != nil {
			failed = append(failed, FailedProvider{
				Name:  pi.provider.Name,
				Error: fmt.Sprintf("SendCommand failed: %v", err),
			})
			panes[i].skipWait = true
		}
	}
	return failed
}

// collectPaneResults waits for each pane's sentinel and collects output.
func collectPaneResults(ctx context.Context, panes []paneInfo, start time.Time) ([]ProviderResponse, []FailedProvider) {
	var (
		responses []ProviderResponse
		failed    []FailedProvider
		mu        sync.Mutex
		wg        sync.WaitGroup
	)

	for _, pi := range panes {
		if pi.skipWait {
			responses = append(responses, ProviderResponse{
				Provider: pi.provider.Name,
				Duration: time.Since(start),
				TimedOut: true,
			})
			continue
		}
		wg.Add(1)
		go func(pi paneInfo) {
			defer wg.Done()
			err := waitForSentinel(ctx, pi.outputFile)
			output := readOutputFile(pi.outputFile)

			mu.Lock()
			defer mu.Unlock()

			responses = append(responses, ProviderResponse{
				Provider: pi.provider.Name,
				Output:   output,
				Duration: time.Since(start),
				TimedOut: err != nil,
			})
			if err != nil {
				failed = append(failed, FailedProvider{
					Name:  pi.provider.Name,
					Error: err.Error(),
				})
			}
		}(pi)
	}
	wg.Wait()
	return responses, failed
}

// paneArgs returns the args to use in pane mode for the given provider.
func paneArgs(p ProviderConfig) []string {
	if len(p.PaneArgs) > 0 {
		return p.PaneArgs
	}
	return p.Args
}

// mergeByStrategy applies the configured merge strategy to responses.
func mergeByStrategy(s Strategy, responses []ProviderResponse, cfg OrchestraConfig) (string, string) {
	switch s {
	case StrategyPipeline:
		return FormatPipeline(responses), fmt.Sprintf("파이프라인: %d단계 완료", len(responses))
	case StrategyDebate:
		return buildDebateMerged(responses, cfg)
	case StrategyFastest:
		if len(responses) > 0 {
			return responses[0].Output, fmt.Sprintf("최속 응답: %s", responses[0].Provider)
		}
		return "", "응답 없음"
	default:
		return MergeConsensus(responses, 0.66)
	}
}

// runFallback runs the standard non-pane orchestration as fallback.
func runFallback(ctx context.Context, cfg OrchestraConfig) (*OrchestraResult, error) {
	fallbackCfg := cfg
	fallbackCfg.Terminal = nil
	return RunOrchestra(ctx, fallbackCfg)
}

// cleanupPanes closes panes and removes temporary output files.
func cleanupPanes(term terminal.Terminal, panes []paneInfo) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, pi := range panes {
		_ = term.Close(ctx, string(pi.paneID))
		_ = os.Remove(pi.outputFile)
	}
}
