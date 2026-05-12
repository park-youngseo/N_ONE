| Gate / Phase | Max Retries | With --loop | On Exceeded |
|-------------|-------------|-------------|-------------|
| Gate 1: Approval | skip if AUTO_MODE | skip if AUTO_MODE | N/A |
| Gate 2: Validation | 3 | 5 | Abort, guide: `--continue` |
| Gate 3: Coverage | re-spawn tester | 3 retries | Until 85%+ / circuit break |
| Phase 4: Review | 2 | 3 | Abort, request user intervention |

### RALF Loop Protocol

WHEN `LOOP_MODE = true`, the pipeline applies the RALF (RED-GREEN-REFACTOR-LOOP) cycle at each quality gate:

```
RALF Loop (per Phase):
  ┌─→ RED    : Phase 실행 (구현/테스트/리뷰)
  │   GREEN  : Gate 검증 (PASS/FAIL 판정)
  │   REFACTOR: FAIL 시 이슈 수정 에이전트 스폰
  └── LOOP   : Gate PASS까지 반복 (circuit breaker 적용)
```

#### Iteration Limits

| Gate / Phase | Without --loop | With --loop |
|-------------|---------------|-------------|
| Gate 2 (Validation) | 3 retries | 5 retries |
| Gate 3 (Coverage) | 1 retry | 3 retries |
| Phase 4 (Review) | 2 retries | 3 retries |

WHERE `LOOP_MODE = false`, existing retry limits apply with no behavioral change (R8 backward compatibility).

#### Circuit Breaker

WHILE `LOOP_MODE = true`, THE SYSTEM SHALL track consecutive no-progress iterations per Phase:
- Progress is measured by: issue count, coverage %, review issue count
- WHERE 3 consecutive iterations produce no new commits or file changes, THE SYSTEM SHALL break the loop
- On circuit break: log stuck state, display remaining issues, suggest recovery options

#### Progress Detection

WHEN each RALF iteration completes, THE SYSTEM SHALL compare current state against previous iteration:
- Issue count: decreased = progress
- Coverage %: increased = progress
- Review issues: decreased = progress
- WHERE no measurable improvement for 3 consecutive iterations → circuit breaker triggers

#### Escalation on Circuit Break

WHEN circuit breaker triggers, THE SYSTEM SHALL display:

```
🐙 RALF [CIRCUIT BREAK] ──────────────
  Phase: {phase} | Iteration: {N}/{max}
  Stuck: 3 consecutive iterations with no progress
  Remaining Issues:
  - {issue list}
  복구 옵션:
  1. /auto go {SPEC-ID} --continue  (수정 후 재개)
  2. /auto fix "{specific issue}"   (개별 이슈 수정)
```

#### Flag Combination Matrix

| Flags | Gate 1 (Approval) | Gates 2-4 |
|-------|-------------------|-----------|
| (none) | User approval | Retry with default limits, stop |
| --auto | Skip | Retry with default limits, stop |
| --loop | User approval | RALF loop with extended limits |
| --auto --loop | Skip | RALF loop with extended limits (완전 자율) |

#### RALF Iteration Display

WHEN `LOOP_MODE = true` AND a retry iteration occurs, THE SYSTEM SHALL display:

```
🐙 RALF [{phase}] ──────────────────
  Iteration: {N}/{max} | Issues: {before} → {after}
  Progress: {description}
  Status: {RETRY | PASS | CIRCUIT_BREAK}
```

**Step 2.1: Determine Quality Mode**

Priority order:
1. If `QUALITY` is set via `--quality` flag → use it. If invalid value, print error and stop.
2. Read `autopus.yaml` → `quality.default` → use as default (applies to both `--auto` and interactive modes)
3. If `AUTO_MODE = true` and no default found → fallback to "balanced"
4. Otherwise → show interactive selection UI with `quality.default` pre-selected

**Reading quality.default**: Use Bash tool to extract: `grep '^quality:' -A1 autopus.yaml | grep 'default:' | awk '{print $2}'` or Read autopus.yaml directly and parse the `quality.default` field.

Quality mode determines the execution profile in Agent() calls:
- **Ultra**: force the premium path for every agent
- **Balanced**: omit explicit overrides and rely on each agent's default profile

Display after determination:
```
품질 모드: {Ultra | Balanced} — {description}
```

**[REQUIRED] Step 0.5: Detect Permission Mode**

```bash
PERMISSION_MODE=$(auto permission detect)
```

If the command fails or is unavailable, default to `PERMISSION_MODE="safe"`.

Display: `권한 모드: {bypass | safe}`

**[REQUIRED] Phase 0.8 — Learning Injection (R6)**

WHEN `.autopus/learnings/pipeline.jsonl` exists, THE SYSTEM SHALL query for entries matching the current SPEC's file paths, packages, or domain keywords:

```bash
auto learn query --spec {SPEC_ID} --limit 5 --format prompt
```

If the command fails or is unavailable, read `.autopus/learnings/pipeline.jsonl` directly, parse JSON entries, and match by:
1. File paths referenced in plan.md tasks
2. Package names from SPEC requirements
3. Domain keywords from SPEC title

Inject the top 5 matching entries (max 2000 tokens) into the planner prompt as a `## Previous Learnings` section. Display:

```
💡 관련 학습 패턴 {N}개 주입됨
```

Skip silently if no matching entries or file doesn't exist.

**[REQUIRED] Phase 1 — Planning (MUST call Agent tool)**

```
Agent(
  subagent_type = "planner",
  model = "opus",          ← Ultra only; omit for Balanced
  mode = PERMISSION_MODE == "bypass" ? "bypassPermissions" : "plan",
  prompt = """
    SPEC file: {SPEC_PATH}
    Plan file: {SPEC_DIR}/plan.md
    Working directory: {WORKING_DIR}

    Decompose the SPEC into implementation tasks.
    Return an agent assignment table:
    | Task ID | Description | Agent    | Mode       | File Ownership  |
    |---------|-------------|----------|------------|-----------------|
    | T1      | ...         | executor | parallel   | pkg/foo/*.go    |
    | T2      | ...         | executor | sequential | pkg/foo/*.go    |
  """
)
```

Display progress after completion:
```
🐙 Pipeline ─────────────────────────
  ✓ Phase 1: Planning
  → Phase 2: Implementation [0/N tasks]
  ○ Phase 3: Testing
  ○ Phase 4: Review
```

> **⏭ POST-AGENT**: planner returned. NEXT REQUIRED STEP: Phase 1.5: Test Scaffold. Do NOT skip to Completion.

**[REQUIRED] Phase 1.5 — Test Scaffold (MUST call Agent tool)**

WHEN `SKIP_SCAFFOLD = false`:
```
Agent(
  subagent_type = "tester",
  mode = "bypassPermissions",
  prompt = """
    Phase: Test Scaffold (Phase 1.5)
    SPEC file: {SPEC_PATH}
    Create failing test skeletons for P0/P1 requirements.
    All tests MUST FAIL. Report any passing tests.
  """
)
```

WHEN `SKIP_SCAFFOLD = true` → [INTENDED SKIP: Phase 1.5]

> **⏭ POST-AGENT**: tester (scaffold) returned. NEXT REQUIRED STEP: Gate 1: Approval. Do NOT skip to Completion.

**[GATE] Gate 1: Approval** — If `AUTO_MODE = false`, show the planner's assignment table and ask user to approve before proceeding. If `AUTO_MODE = true` → [INTENDED SKIP: Gate 1]

**[REQUIRED] Phase 1.8 — Doc Fetch (Context7 MCP)**

WHEN Gate 1 completes (or is skipped), THE SYSTEM SHALL fetch latest documentation for external libraries referenced in the SPEC using Context7 MCP. This phase runs in the **main session** — subagents cannot access MCP tools.

1. **Detect technologies**: Scan SPEC requirements, plan.md tasks, and file imports for external library names (skip standard library modules)
2. **Fetch docs** (up to 5 libraries): Call `mcp__context7__resolve-library-id` → `mcp__context7__query-docs` for each detected library
3. **Cache and prepare**: Adaptive token budget (1 lib→~5000, 2→~3000, 3→~2500, 4-5→~2000 tokens/lib), format as `## Reference Documentation` section
4. **Skip condition**: If no external libraries detected, skip Phase 1.8 entirely

Error handling: Log and skip on any Context7 failure — documentation is supplementary, never blocks the pipeline.

Ref: `.gemini/rules/autopus/context7-docs.md` for full detection heuristics, token limits, and anti-patterns.

> **⏭ POST-PHASE**: Doc Fetch complete (or skipped). NEXT REQUIRED STEP: Phase 2: Implementation. Do NOT skip to Completion.

**[REQUIRED] Phase 2 — Implementation (MUST call Agent tool)**

For each task in the planner's assignment table, spawn executor agents:

**GC Suppression**: During parallel worktree execution, all git commands MUST use `git -c gc.auto=0` to prevent garbage collection interference. See `.gemini/rules/autopus/worktree-safety.md`.

Parallel tasks (Mode = "parallel" AND no file ownership conflict):
```
# Call MULTIPLE Agent() in a SINGLE message for parallel execution
# Each parallel executor runs in an isolated git worktree (R1)
# Use git -c gc.auto=0 for all git commands during parallel execution (R5)
# Include {ctx7_docs} from Phase 1.8 in each executor prompt
Agent(
  subagent_type = "executor",
  model = "opus",          ← Ultra only; omit for Balanced
  mode = "bypassPermissions",
  isolation = "worktree",  ← R1: worktree isolation for parallel tasks
  prompt = "## Reference Documentation\n{ctx7_docs}\n\nImplement T1: {task description}. Files: {file ownership}. SPEC requirements: {relevant requirements}. Working directory: {WORKING_DIR}. Run all build/test commands inside {WORKING_DIR}: cd {WORKING_DIR} && go build ./..."
)
Agent(
  subagent_type = "executor",
  model = "opus",          ← Ultra only; omit for Balanced
  mode = "bypassPermissions",
  isolation = "worktree",  ← R1: worktree isolation for parallel tasks
  prompt = "## Reference Documentation\n{ctx7_docs}\n\nImplement T2: {task description}. Files: {file ownership}. SPEC requirements: {relevant requirements}. Working directory: {WORKING_DIR}. Run all build/test commands inside {WORKING_DIR}: cd {WORKING_DIR} && go build ./..."
)
```

Max concurrent worktrees: **5**. Queue overflow tasks and spawn as slots free.

Sequential tasks (Mode = "sequential" OR file ownership conflict) do NOT use worktree isolation:
```
# Call Agent() one at a time, passing previous result to next
# Sequential: merge worktree branch immediately after each task before spawning next
#   git -c gc.auto=0 merge <branch_T1> && git worktree remove <path_T1>
result_t1 = Agent(subagent_type = "executor", prompt = "Implement T1: ... Working directory: {WORKING_DIR}. Run all build/test commands inside {WORKING_DIR}: cd {WORKING_DIR} && go build ./...")
Agent(subagent_type = "executor", prompt = "Implement T2 (depends on T1). T1 result: {result_t1}. ... Working directory: {WORKING_DIR}. Run all build/test commands inside {WORKING_DIR}: cd {WORKING_DIR} && go build ./...")
```

> **⏭ POST-AGENT**: executor(s) returned. NEXT REQUIRED STEP: [REQUIRED] Phase 2.1: Worktree Merge. Do NOT skip to Completion.

**[REQUIRED] Phase 2.1 — Worktree Merge (parallel tasks only)**

WHEN all parallel executors complete, merge their worktree branches into the working branch before Gate 2:

```bash
# Merge in task-ID order (R3)
git -c gc.auto=0 merge <branch_T1>   # R5: gc.auto=0 suppresses auto GC
git worktree remove <path_T1>
git -c gc.auto=0 merge <branch_T2>
git worktree remove <path_T2>
```

On lock error (`.git/refs.lock`, etc.), retry with exponential backoff: 3s → 6s → 12s (max 3 retries). See `.gemini/rules/autopus/worktree-safety.md`.

After merge, display:
```
🐙 merge ────────────────────────────
  브랜치: {N}개 머지 | 충돌: 0 | 정리: {N}개 워크트리 제거
  다음: 검증 단계로 진행
```

**[CHECKPOINT] Phase 2.1 Complete** — All worktree branches merged. Proceed to Gate 2.

**Merge conflict handling** (R4):
- `git merge --abort` to restore clean state
- Log: `[MERGE ERROR] T{id} merge failed: ownership validation gap detected. Files: {list}`
- Abort pipeline even in `--auto` mode (never auto-resolve to prevent data loss)

**Worktree merge failure recovery:**
```
✗ Worktree 머지 실패: {file list}
  원인: 파일 소유권 검증 갭 (ownership validation gap)
  복구 옵션:
  1. /auto go {SPEC-ID} --continue  (수정 후 재개)
  2. git worktree list              (워크트리 상태 확인)
  3. git worktree remove --force    (수동 정리)
```

After each executor, display:
```
🐙 executor ─────────────────────────
  파일: {N}개 수정 | 테스트: {N}개 추가 | 커버리지: {N}%
  다음: Phase 2.1 Worktree Merge (병합 필수)
```

**[GATE] Gate 2 — Validation (MUST call Agent tool)**

```
Agent(
  subagent_type = "validator",
  mode = PERMISSION_MODE == "bypass" ? "bypassPermissions" : "plan",
  prompt = """
    Validate the implementation. Check:
    - go build ./... passes
    - go vet ./... passes
    - golangci-lint run passes
    - No file exceeds 300 lines

    Return format:
    Verdict: PASS | FAIL