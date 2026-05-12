### Phase 2.1: Worktree Merge

WHEN all parallel executors complete, THE SYSTEM SHALL merge their worktree branches into the working branch before proceeding to Gate 2.

**Sequential tasks**: Already merged immediately after each task completion during Phase 2.

**Parallel tasks (batch merge)**:
1. Collect all worktree branches with changes
2. Merge in task-ID order (T1 → T2 → T3 ...)
3. For each branch: `git -c gc.auto=0 merge <branch>` → on success: `git worktree remove <path>`
4. On merge conflict: `git merge --abort` → abort pipeline → report error

See @.claude/skills/autopus/worktree-isolation.md for full merge strategy and safety rules.

### Gate 2: Validation

```
Agent(
  subagent_type = "validator",
  prompt = """
    Validate the implementation result.

    Run ALL 6 verification checks:
    1. Build — compile/transpile passes
    2. Test — all tests pass
    3. Lint — no lint warnings
    4. Coverage — measure test coverage
    5. Structure — no file exceeds 300 lines
    6. Seam Verification:
       a. Stub Detection — grep changed files for TODO/stub/placeholder/NotImplemented patterns
       b. Smoke Test — run CLI/API entry point (--help or /health) if applicable
       c. Contract Parity — if both client and server code changed, verify endpoint paths match

    Return format:
    Verdict: PASS | FAIL
    Issues: <list of issues>
    Recommended Agent: executor | tester | planner
  """
)
```

### Phase 2.5: Annotation (Post-Validation)

WHEN Gate 2 returns PASS, THE SYSTEM SHALL execute an annotation step before proceeding to Phase 3.

A dedicated annotator agent is spawned to apply @AX tags:

```
Agent(
  subagent_type = "annotator",
  prompt = """
    Apply @AX tags to modified files based on the ax-annotation skill.
    Reference: pkg/content/ax.go:GenerateAXInstruction() for canonical rules.

    Executor work log: {modified files list, change intent from Phase 2}

    For each modified file:
    1. Scan for NOTE triggers (magic constants, undocumented exports >100 lines)
    2. Scan for WARN triggers (goroutines without context, complexity >= 15, global state mutation)
    3. Scan for ANCHOR triggers (grep for fan_in >= 3 callers)
    4. Scan for TODO triggers (public functions without tests)
    5. Validate per-file limits (ANCHOR max 3, WARN max 5)
    6. Apply overflow strategy if limits exceeded

    All tags MUST include the [AUTO] prefix.
  """,
  mode = "bypassPermissions"
)
```

Annotation is skipped for harness-only tasks (all `.md` files).

### Phase 3.5: UX Verification (Optional)

WHEN the target project contains frontend components (.tsx/.jsx files) AND the pipeline is running in subagent or Agent Teams mode (not `--solo`), THE SYSTEM SHALL execute UX verification between Testing and Review.

```
Agent(
  subagent_type = "frontend-specialist",
  prompt = """
    Run frontend UX verification on all modified frontend components.
    Reference: .claude/skills/autopus/frontend-verify.md for the full pipeline.

    1. Analyze git diff to identify changed .tsx/.jsx files
    2. Generate or heal Playwright E2E tests for affected components
    3. Execute tests and capture screenshots
    4. Analyze screenshots for visual issues (layout, readability, responsiveness)
    5. Attempt auto-fix for WARN/FAIL items (max 2 attempts)

    Return format:
    Verdict: PASS | WARN | FAIL
    Screenshots: N analyzed
    Issues: <list of issues with file references>
    Fixes: <list of auto-applied fixes>
  """,
  mode = "bypassPermissions"
)
```

Activation conditions:
- Frontend files (.tsx/.jsx) exist in the changed file set
- Skip if all changes are backend-only (.go, .md)

Phase 3.5 does NOT renumber existing phases. Testing remains Phase 3, Review remains Phase 4.

### Phase 3: Testing

```
Agent(
  subagent_type = "tester",
  prompt = """
    ## Reference Documentation
    {ctx7_docs}

    Raise coverage to 85%+.
    Add missing edge case tests.
  """,
  mode = "bypassPermissions"
)
```

### Phase 4: Review (Parallel)

reviewer and security-auditor run in parallel:

```
Agent(subagent_type = "reviewer", prompt = """
    Perform a code review using TRUST 5 criteria. Return format:
    Verdict: APPROVE | REQUEST_CHANGES
    Issues: <list of issues>
""")
Agent(subagent_type = "security-auditor", prompt = """
    Perform a security audit. Return format:
    Verdict: PASS | FAIL
    Issues: <list of security issues>
""")
```

Both must return PASS/APPROVE. On conflict, Lead (planner) consolidates issue lists.
Priority: security issues > code quality issues.

Freeze the review output into a checklist of open findings.

- If the checklist still contains actionable findings and the retry budget remains, immediately delegate a focused fixer/executor task inside the same invocation.
- Keep the checklist stable across retries unless the patch meaningfully changes scope.
- Do not ask the user to manually fix, rerun, or confirm while the next repair step is still actionable within the current `/auto go` invocation.

## Parallel vs Sequential Decision Criteria

| Condition                                     | Execution         | Worktree Isolation |
|-----------------------------------------------|-------------------|--------------------|
| planner specifies Mode = "parallel"           | Parallel          | Yes (`isolation: "worktree"`) |
| planner specifies Mode = "sequential"         | Sequential        | No (main worktree) |
| File ownership conflict detected (R2)         | Switch to sequential | No (main worktree) |
| Task uses previous task result as input       | Sequential        | No (main worktree) |

File ownership conflict always forces sequential execution, even when worktree isolation is available (R2). The planner SHOULD design non-overlapping file ownership to maximize parallel execution with worktree isolation.

## Quality Gate Handling

```
PASS  → Proceed to next Phase
FAIL  → Delegate fix to the Recommended Agent from Gate Verdict → re-validate
```

Retry limits:

- Gate 2 (Validation): maximum 3 retries
- Phase 4 (Review): maximum 2 retries

While the review retry budget remains, keep the repair -> validate -> verify cycle inside the same invocation.

Only when the retry limit is exhausted or the pipeline hits a real blocker/circuit break should it abort and notify the user:

```
Pipeline aborted: failed to resolve [Gate name] after [N] retries.
Manual intervention required. Last issue: [Issues content]
```

## Agent Failure Handling

| Failure Type              | Handling                                           |
|---------------------------|----------------------------------------------------|
| Exits due to maxTurns     | Detect remaining work → spawn new Agent()          |
| Subagent returns error    | Analyze error content → retry with revised prompt  |
| Retry limit exceeded      | Main session implements directly (fallback)        |

Fallback condition: if a subagent fails 2 consecutive times, the main session handles the task directly.

## Pipeline Monitoring Integration

### Log Path Injection (R5)

WHEN spawning agents in any Phase, THE SYSTEM SHALL inject the pipeline log file path into each agent's prompt.

**Injection format:**

```
## Pipeline Monitor
Log file: /tmp/autopus-pipeline-{spec-id}.log
Write structured log entries: [timestamp] [your-role] [phase] message
```

**Usage in Agent() calls:**

```python
logger = PipelineLogger(log_dir)
Agent(
  subagent_type = "executor",
  prompt = f"""
    {logger.prompt_injection()}

    ## Task
    {task_description}
  """
)
```

### Dashboard Refresh (R4/R8)

WHEN a Phase transition occurs (e.g., Phase 1 → Phase 2), THE SYSTEM SHALL refresh the dashboard pane:

```python
# After phase transition, refresh dashboard pane
term.SendCommand(ctx, dashboard_pane_id, f"auto pipeline dashboard {spec_id}")
```

### Monitor Session Lifecycle

```
Pipeline Start   → MonitorSession.Start(ctx)  → creates 2 panes (cmux only)
Phase Transition → logger.LogEvent(event)      → writes to JSONL + text log
                 → term.SendCommand(dashboard) → refreshes dashboard
Pipeline End     → MonitorSession.Close(ctx)   → closes panes, removes temp files
```

### Event Types

| Event | When Emitted |
|-------|-------------|
| `phase_start` | Phase begins |
| `phase_end` | Phase completes |
| `agent_spawn` | Agent is spawned |
| `agent_done` | Agent finishes |
| `checkpoint` | Checkpoint saved |
| `error` | Error occurs |
| `blocker` | Blocker detected |

## Harness-Only Task Handling

When all tasks modify only `.md` files:

- Skip Go build/test validation
- Validator checks only file format (frontmatter YAML, section structure)
- Coverage gate (85%) is not applied

Determination: if all "file ownership" entries in the planner's assignment table are `*.md`, treat as harness-only.

## Result Integration and Completion

Once all Phases are complete:

1. Collect results from each agent and output a final summary
2. Update the SPEC file status to `"done"`
3. Guide next steps: `/auto sync <SPEC-ID>`

### Final Summary Format

```
## Pipeline Completion Summary

SPEC: <SPEC-ID>
Tasks: <completed> / <total>
Coverage: <measured>%
Review: APPROVE

Completed Files:
- <file path 1>
- <file path 2>
```

## Completion Criteria

- [ ] All Phases executed in order
- [ ] PASS verdict received at each Gate
- [ ] Coverage 85%+ confirmed
- [ ] SPEC status = "done" updated
- [ ] Final summary output complete