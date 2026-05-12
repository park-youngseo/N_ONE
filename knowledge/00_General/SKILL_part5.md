    Issues: <list>
    Recommended Agent: executor | tester | planner
  """
)
```

- **PASS** → proceed to Phase 3
- **FAIL** → spawn the Recommended Agent to fix issues, then re-validate (up to 3 retries)

WHEN `LOOP_MODE = true`: extend retry limit to 5. Display RALF iteration status after each retry. Apply circuit breaker if 3 consecutive no-progress iterations.

**Validation failure recovery (Gate 2):**
```
✗ Validation 실패: {N}개 이슈 발견
  복구 옵션:
  1. /auto go {SPEC-ID} --continue  (중단점에서 재개)
  2. /auto fix "{specific issue}"   (개별 이슈 수정)
```

After validation, display:
```
🐙 validator ────────────────────────
  Verdict: {PASS|FAIL} | 검사 항목: {N}개 | 이슈: {N}개
  다음: {pass → 테스트 단계 | fail → 수정 후 재검증}
```

> **⏭ POST-AGENT**: validator returned. NEXT REQUIRED STEP: Phase 2.5: Annotation. Do NOT skip to Completion.

**[REQUIRED] Phase 2.5 — Annotation (MUST call Agent tool)**

```
Agent(
  subagent_type = "annotator",
  mode = "bypassPermissions",
  prompt = """
    Apply @AX tags to all files modified during Phase 2.
    Follow ax-annotation skill rules.
    Validate per-file limits: ANCHOR ≤ 3, WARN ≤ 5.
  """
)
```

> **⏭ POST-AGENT**: annotator returned. NEXT REQUIRED STEP: Phase 3: Testing. Do NOT skip to Completion.

**[REQUIRED] Phase 3 — Testing (MUST call Agent tool)**

```
Agent(
  subagent_type = "tester",
  mode = "bypassPermissions",
  prompt = """
    Raise test coverage to 85%+.
    Add missing edge case tests.
    Run: go test -race -cover ./...
    Ensure all tests pass.
  """
)
```

After tester, display:
```
🐙 tester ───────────────────────────
  테스트: {N}개 추가 | 커버리지: {before}% → {after}% | 통과: {N}/{N}
  다음: 리뷰 단계로 진행
```

> **⏭ POST-AGENT**: tester returned. NEXT REQUIRED STEP: Gate 3: Coverage Check. Do NOT skip to Completion.

**[GATE] Gate 3: Coverage** — verify 85%+ coverage from tester output. If below, re-spawn tester. For harness-only tasks → [INTENDED SKIP: Gate 3]

WHEN `LOOP_MODE = true`: extend coverage retry limit to 3 (default: 1). Display RALF iteration status after each retry.

**[REQUIRED] Phase 4 — Review (MUST call Agent tool)**

Run reviewer and security-auditor in **parallel**:

```
# Call BOTH in a SINGLE message for parallel execution
Agent(
  subagent_type = "reviewer",
  mode = PERMISSION_MODE == "bypass" ? "bypassPermissions" : "plan",
  prompt = """
    Review all changes using TRUST 5 criteria.
    Return format:
    Verdict: APPROVE | REQUEST_CHANGES
    Issues: <list with file:line references>
  """
)
Agent(
  subagent_type = "security-auditor",
  mode = PERMISSION_MODE == "bypass" ? "bypassPermissions" : "plan",
  prompt = """
    Audit all changed files for security vulnerabilities.
    Return format:
    Verdict: PASS | FAIL
    Issues: <list with file:line references and severity>
  """
)
```

Both must return APPROVE/PASS for the gate to pass. If results conflict, Lead consolidates issue lists with priority: security issues > code quality issues. Consolidated issue list sent to executor for remediation.

- **Both APPROVE/PASS** → pipeline complete
- **Either REQUEST_CHANGES/FAIL** → spawn executor to fix consolidated issues, then re-review (up to 2 retries)

WHEN `LOOP_MODE = true`: extend review retry limit to 3. Display RALF iteration status after each retry. Apply circuit breaker if 3 consecutive identical issues.

**Review failure recovery (retries exceeded):**
```
✗ Review 실패: {N}개 변경 요청 (재시도 초과)
  복구 옵션:
  1. /auto go {SPEC-ID} --continue  (수정 후 재개)
  2. /auto review                   (변경사항 직접 검토)
```

After reviewer, display:
```
🐙 reviewer ─────────────────────────
  Verdict: {APPROVE|REQUEST_CHANGES} | Issues: {N}개 | 심각도: {low|medium|high}
  다음: {verdict-based guidance}
```

> **⏭ POST-AGENT**: reviewer returned. NEXT REQUIRED STEP: Gate 4: Review Result Handling. Do NOT skip to Completion.

**[GATE] Gate 4: Review Result Handling**

- **APPROVE** → pipeline complete, proceed to Pre-Completion Verification
- **REQUEST_CHANGES** → spawn executor to fix issues, then re-review (max 2 retries, see **Retry Limits Reference**)

**Review failure recovery (retries exceeded):**
```
✗ Review 실패: {N}개 변경 요청 (재시도 초과)
  복구 옵션:
  1. /auto go {SPEC-ID} --continue  (수정 후 재개)
  2. /auto review                   (변경사항 직접 검토)
```

**Parallel execution rules:**
- Only parallelize tasks whose Mode is "parallel" in the planner's assignment table
- File ownership conflict (same path pattern) → force sequential
- Pass SPEC requirements, task description, and relevant file context to each subagent

**Retry limits:** See **Retry Limits Reference** table above.

**Fallback:** If a subagent fails 2 consecutive times, handle in the main session directly.

**Subagent failure recovery:**
```
✗ {agent-name} 실패: {error description}
  복구 옵션:
  1. /auto go {SPEC-ID} --continue  (파이프라인 재개)
  2. {manual fallback instruction}  (수동 처리)
```

### Pipeline (--multi mode: multi-provider)

When `--multi` flag is set, activate multi-provider orchestration:

- Use the `auto orchestra review` engine to review with multiple AI providers (claude, gemini, codex)
- Validate with multi-provider consensus at the review gate
- Can be combined with `--strategy`: consensus (default), debate, pipeline, fastest

`--team --multi` combination: Agent Teams + multi-provider review activated in the Review Phase.

Usage: `/auto go SPEC-ID` (subagent pipeline), `/auto go SPEC-ID --team` (Agent Teams), `/auto go SPEC-ID --solo` (single session), `/auto go SPEC-ID --quality ultra`, `/auto go SPEC-ID --auto --loop`

### Flags

| Flag | Description |
|------|-------------|
| `--auto` | Autonomous mode: run the full pipeline without user confirmation |
| `--loop` | RALF loop mode: auto-retry failed gates with extended iteration limits until PASS or circuit break |
| `--continue` | Resume an interrupted implementation from a previous session |
| `--team` | Agent Teams mode: use Gemini CLI Agent Teams (experimental) for multi-agent collaboration |
| `--solo` | Single session mode: no subagents, main session implements directly |
| `--multi` | Multi-provider mode: activate orchestra engine-based review |
| `--strategy` | Multi-provider strategy: consensus, debate, pipeline, fastest (requires --multi) |
| `--quality` | Quality mode: ultra (all agents Opus), balanced (mixed) |
| `--skip-scaffold` | Skip Phase 1.5 Test Scaffold |

### Harness-Only Tasks

If all tasks modify only `.md` files:
- Skip Go build/test validation
- Validator performs format-only checks (frontmatter validity, 300-line limit)

### [REQUIRED] Pre-Completion Verification

Before displaying completion output, verify ALL phases were evaluated:

- [ ] Phase 1: Planning — completed
- [ ] Phase 1.5: Test Scaffold — completed OR skipped (--skip-scaffold)
- [ ] Gate 1: Approval — completed OR skipped (AUTO_MODE)
- [ ] Phase 1.8: Doc Fetch — completed OR skipped (no external libs)
- [ ] Phase 2: Implementation — all tasks completed
- [ ] Phase 2.1: Worktree Merge — completed (if parallel tasks existed) OR N/A
- [ ] Gate 2: Validation — PASS
- [ ] Phase 2.5: Annotation — completed
- [ ] Phase 3: Testing — completed (or skipped for harness-only)
- [ ] Gate 3: Coverage — 85%+ verified (or N/A for harness-only)
- [ ] Phase 4: Review — APPROVE

IF any item is unchecked → return to that phase/gate. Do NOT display Completion Guidance.

### Completion Guidance

After go completes, display the workflow lifecycle bar:

```
🐙 Workflow: {SPEC-ID}
  ✓ plan  →  ● go  →  ○ sync
```

Then show state-aware next step guidance:

```
다음 단계: /auto sync {SPEC-ID}
```

If the SPEC status is now `implemented`, also check project state:
- SPEC status `implemented` → suggest `/auto sync {SPEC-ID}`
- No project docs found → suggest `/auto setup` before sync


---

## verify — Frontend UX Verification

@.gemini/skills/autopus/frontend-verify.md

Verify frontend UX changes using VLM-powered visual analysis and Playwright E2E tests.

### Pipeline

#### [REQUIRED] Phase 0: Change Analysis

Run `git diff` to determine impact scope. Identify which UI components and pages were modified.

> **⏭ POST-PHASE**: Change scope identified. NEXT REQUIRED STEP: Phase 1: Test Gen/Heal. Do NOT skip to Completion.

#### [REQUIRED] Phase 1: Test Gen/Heal

Generate new Playwright tests or repair existing ones based on the change scope from Phase 0.

> **⏭ POST-PHASE**: Tests generated/repaired. NEXT REQUIRED STEP: Phase 2: Test Runner. Do NOT skip to Completion.

#### [REQUIRED] Phase 2: Test Runner (MUST call Bash tool)

IMPORTANT: You MUST execute Playwright tests via the Bash tool. Do NOT simulate test results.

```bash
npx playwright test
```

**Fallback rule**: Only if the Bash call returns an error, display the exact error to the user and ask for fallback approval.

> **⏭ POST-PHASE**: Tests executed, screenshots captured. NEXT REQUIRED STEP: Phase 3: VLM Verifier. Do NOT skip to Completion.

#### [REQUIRED] Phase 3: VLM Verifier

Perform semantic visual analysis on captured screenshots. Return PASS/WARN/FAIL per component.

> **⏭ POST-PHASE**: Visual verification complete. NEXT REQUIRED STEP: Phase 4: Reporter. Do NOT skip to Completion.

#### [REQUIRED] Phase 4: Reporter

Apply fixes (unless `--report-only`) and produce the verification report.

### [REQUIRED] Pre-Completion Verification

Before displaying the final report, verify ALL phases were executed:

- [ ] Phase 0: Change Analysis — git diff run, scope identified
- [ ] Phase 1: Test Gen/Heal — Playwright tests generated/repaired
- [ ] Phase 2: Test Runner — Bash tool called with `npx playwright test` (not simulated)
- [ ] Phase 3: VLM Verifier — visual analysis completed
- [ ] Phase 4: Reporter — report produced

IF any item is unchecked → return to that phase. Do NOT display completion output.

### Flags

| Flag | Description |
|------|-------------|
| `--fix` | Enable auto-fix mode (default: enabled) |
| `--report-only` | Skip auto-fix, output findings only |
| `--viewport <size>` | Viewport size: desktop, mobile, or WxH (default: desktop) |

Usage: `/auto verify`, `/auto verify --report-only`, `/auto verify --viewport mobile`

---

## browse — Browser Automation

@.gemini/skills/autopus/browser-automation.md

Open a browser and interact with web pages. Automatically selects cmux browser (in cmux terminal) or agent-browser (fallback) based on terminal detection.

### Prerequisites
