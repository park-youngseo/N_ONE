
Skip missing files with a warning log.

### debate.go Change

Add a conditional in the debate round execution:

```go
var isolation string
if cfg.ContextMode {
    isolation = contextAwareInstruction + projectContext
} else {
    isolation = topicIsolationInstruction
}
```

Add new constant:

```go
const contextAwareInstruction = "Use the project context below to ground your ideas in the actual codebase. Focus on the given topic.\n\n"
```

---

## T10: Structured JSON output for --no-judge

**Priority**: P1
**File**: `pkg/orchestra/yield.go`
**Requirement**: R9

### Implementation

Extend the `YieldOutput` struct (created in T5) to handle the full --no-judge output format. This is a refinement of T5's JSON output to match the exact R9 schema including `round_history` with per-round responses.

---

## T11: idea.md skill update

**Priority**: P1
**File**: `.claude/skills/autopus/idea.md`
**Requirement**: R10

### Changes

Update the skill's Step 3-4 flow:

- Step 3: Replace direct orchestra call with `auto orchestra brainstorm --no-judge --yield-rounds`
- Add Steps 3.5-3.8 for the main-session judge workflow
- Step 4: Keep ICE scoring but ensure it runs in the main session

---

## T12: auto orchestra inject command

**Priority**: P2
**Files**: `internal/cli/orchestra_inject.go` (new), `internal/cli/orchestra.go`
**Requirement**: R11

### Command Structure

```
auto orchestra inject --session-id <ID> --provider <name> "prompt text"
```

### Implementation

1. Load session from persistence file
2. Find pane ID for the specified provider
3. Use `SendLongText` (which internally uses cmux set-buffer/paste-buffer) to inject prompt
4. Output confirmation to stderr

---

## T13: Unit tests for stability fixes (T1-T4)

**Priority**: P0
**Files**: `pkg/orchestra/surface_manager_test.go`, `pkg/orchestra/interactive_debate_test.go`

### Test Cases

- T1: Test sendPromptWithRetry retries on same pane before recreation
- T1: Test sendPromptWithRetry falls back to recreation after 2 same-pane failures
- T1: Test successful send on first try (no retry)
- T2: Verify pollUntilPrompt timeout constant is 30s
- T3: Verify default scrollback lines is 500
- T4: Test judge skip when NoJudge is true
- T4: Test judge runs when NoJudge is false (default)

---

## T14: Unit tests for yield/collect/cleanup (T5-T8)

**Priority**: P0
**Files**: `pkg/orchestra/yield_test.go`, `pkg/orchestra/session_test.go`, `internal/cli/orchestra_collect_test.go`, `internal/cli/orchestra_cleanup_test.go`

### Test Cases

- T5: Test YieldOutput JSON serialization matches schema
- T6: Test collect reads session and returns JSON
- T7: Test cleanup removes panes and session file
- T8: Test session save/load/remove lifecycle
- T8: Test session file permissions are 0600
- T8: Test NewSessionID generates unique IDs