package orchestra

import (
	"bufio"
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
[Global Connectivity Table]
| Node | Responsibility | Parent |
| :--- | :--- | :--- |
| pane_runner_io.go | Command build, sentinel detection, file I/O | pane_runner.go |
*/

// @AX:NOTE [AUTO] sentinel marker written to output files to signal provider completion
const sentinel = "__AUTOPUS_DONE__"

// buildPaneCommand constructs the shell command to execute in a pane.
func buildPaneCommand(provider ProviderConfig, prompt, outputFile string) string {
	args := shellEscapeArgs(paneArgs(provider))
	binary := shellEscapeArg(provider.Binary)
	safeOutput := shellEscapeArg(outputFile)

	if provider.PromptViaArgs {
		return fmt.Sprintf("%s %s %s | tee %s; echo %s >> %s",
			binary, args, shellEscapeArg(prompt), safeOutput, sentinel, safeOutput)
	}
	delim := uniqueHeredocDelimiter("PROMPT_EOF", prompt, randomHex())
	return fmt.Sprintf("%s %s <<'%s'\n%s\n%s\n | tee %s; echo %s >> %s",
		binary, args, delim, prompt, delim, safeOutput, sentinel, safeOutput)
}

// waitForSentinel polls the output file until the sentinel marker is found.
func waitForSentinel(ctx context.Context, outputFile string) error {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if hasSentinel(outputFile) {
				return nil
			}
		}
	}
}

// hasSentinel checks if the output file contains the sentinel marker.
func hasSentinel(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), sentinel) {
			return true
		}
	}
	return false
}

// readOutputFile reads the output file and strips the sentinel marker.
func readOutputFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	output := strings.ReplaceAll(string(data), sentinel, "")
	return strings.TrimSpace(output)
}

// randomHex returns an 8-character random hex string.
func randomHex() string {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%08x", time.Now().UnixNano()&0xFFFFFFFF)
	}
	return fmt.Sprintf("%x", b)
}
