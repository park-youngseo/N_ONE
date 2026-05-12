> **⏭ POST-STEP**: Root cause identified. NEXT REQUIRED STEP: Step 3: Minimal Fix. Do NOT skip to Completion.

#### [REQUIRED] Step 3: Minimal Fix

Apply the smallest change that fixes the root cause. Avoid scope creep.

> **⏭ POST-STEP**: Fix applied. NEXT REQUIRED STEP: Step 4: Verify. Do NOT skip to Completion.

#### [REQUIRED] Step 4: Verify (MUST call Bash tool)

IMPORTANT: You MUST run the test suite via the Bash tool. Do NOT assume the fix works without running tests.

```bash
go test -race ./...
```

**Fallback rule**: Only if the Bash call returns an error, display the exact error to the user and ask for fallback approval.

### [REQUIRED] Pre-Completion Verification

Before reporting the fix as complete, verify ALL steps were executed:

- [ ] Step 1: Failing test written (or reproduction steps documented)
- [ ] Step 2: Root cause identified and documented
- [ ] Step 3: Minimal fix applied
- [ ] Step 4: Bash tool called — `go test -race ./...` passed

IF any item is unchecked → return to that step. Do NOT report completion.

#### [OPTIONAL] Step 5: Learning Capture (R11)

WHEN the fix is successfully verified (Step 4 passed), THE SYSTEM SHALL record the fix pattern:

```bash
auto learn record --type fix_pattern --files {changed_files} --pattern "{error → root cause → resolution}" --spec {SPEC_ID_if_available}
```

If the command fails, append directly to `.autopus/learnings/pipeline.jsonl`:
```json
{"id":"L-{next}","timestamp":"{now}","type":"fix_pattern","phase":"fix","files":[...],"packages":[...],"pattern":"{error description}","resolution":"{how it was fixed}","severity":"low","reuse_count":0}
```

Display:
```
💡 수정 패턴 학습됨: "{one-line pattern summary}"
```

Usage: `/auto fix "bug description"`, `/auto fix --file path/to/file.go`

---

## map — Analyze Codebase Structure

Analyze and visualize the codebase structure, dependencies, and key patterns.

Analysis includes:
- Directory structure and package hierarchy
- Core entry points (main, handlers)
- Dependency flow (inter-module imports)
- @AX ANCHOR points (high fan-in functions)
- Complexity hotspots

Usage: `/auto map`, `/auto map path/to/package`

---

## review — Code Review

@.gemini/skills/autopus/review.md

Review changed code against TRUST 5 criteria (Tested, Readable, Unified, Secured, Trackable).

### Pipeline

#### [REQUIRED] Step 1: Run Automated Gates (MUST call Bash tool)

IMPORTANT: You MUST run automated quality checks via the Bash tool before the TRUST 5 analysis.

```bash
go test -race ./... && golangci-lint run
```

**Fallback rule**: Only if the Bash call returns an error, display the exact error to the user and ask for fallback approval.

> **⏭ POST-STEP**: Automated gates complete. NEXT REQUIRED STEP: Step 2: TRUST 5 Analysis. Do NOT skip to Completion.

#### [REQUIRED] Step 2: TRUST 5 Analysis

Analyze all changed files against each TRUST dimension:
- **T**ested: test coverage, edge cases
- **R**eadable: naming, complexity, comments
- **U**nified: style consistency, patterns
- **S**ecured: input validation, secrets, OWASP
- **T**rackable: @AX tags, Lore entries, SPEC refs

### [REQUIRED] Pre-Completion Verification

Before displaying the review verdict, verify ALL steps were executed:

- [ ] Step 1: Bash tool called — `go test -race ./...` and `golangci-lint run` ran (not skipped)
- [ ] Step 2: TRUST 5 analysis completed for all changed files

IF any item is unchecked → return to that step. Do NOT display the verdict.

Usage: `/auto review`, `/auto review HEAD~3..HEAD`, `/auto review --pr 123`


---

## spec review — SPEC Multi-Provider Review

@.gemini/skills/autopus/spec-review.md

Review a SPEC document using multiple providers (claude, gemini, etc.) to validate quality. Delivers a PASS/REVISE/REJECT verdict via the Orchestra engine.

### Pipeline

#### [REQUIRED] Step 1: Load SPEC

Run SPEC Path Resolution for {SPEC-ID}. Load the resolved SPEC_PATH and extract requirements. Verify the file exists before proceeding.

> **⏭ POST-STEP**: SPEC loaded. NEXT REQUIRED STEP: Step 2: Collect Code Context. Do NOT skip to Completion.

#### [REQUIRED] Step 2: Collect Code Context

If `auto_collect_context: true` in `autopus.yaml`, gather relevant source code files referenced by the SPEC.

> **⏭ POST-STEP**: Context collected. NEXT REQUIRED STEP: Step 3: Run Multi-Provider Review. Do NOT skip to Completion.

#### [REQUIRED] Step 3: Run Multi-Provider Review (MUST call Bash tool)

IMPORTANT: You MUST execute the multi-provider review via the Bash tool. Do NOT simulate review results.

```bash
auto spec review {SPEC-ID} --strategy {STRATEGY}
```

**Fallback rule**: Only if the Bash call returns an error (binary not found, provider config missing, etc.), display the exact error to the user and ask for fallback approval.

> **⏭ POST-STEP**: Multi-provider review returned. NEXT REQUIRED STEP: Step 4: Merge Verdicts. Do NOT skip to Completion.

#### [REQUIRED] Step 4: Merge Verdicts

Determine final verdict according to the configured strategy (debate, consensus):
- **PASS** → proceed to Step 5
- **REVISE** → return finding list to user, apply fixes, re-run (up to 2 iterations)
- **REJECT** → return finding list, guide user to redesign

#### [REQUIRED] Step 5: Save Results

Write `review.md` to {SPEC_DIR} with the full verdict and findings. Update SPEC status if PASS.

### Verdict Criteria

- **PASS**: SPEC approved, update status to "approved"
- **REVISE**: Revision needed, return with finding list
- **REJECT**: Fundamental redesign required

### [REQUIRED] Pre-Completion Verification

Before displaying final output, verify ALL steps were executed:

- [ ] Step 1: SPEC loaded
- [ ] Step 2: Code context collected (or skipped if `auto_collect_context: false`)
- [ ] Step 3: Bash tool called — `auto spec review` ran (not simulated)
- [ ] Step 4: Verdicts merged, final verdict determined
- [ ] Step 5: `review.md` written, SPEC status updated

IF any item is unchecked → return to that step. Do NOT display completion output.

Usage: `/auto spec review SPEC-ID`, `/auto spec review SPEC-ID --strategy debate --timeout 180`


---

## secure — Security Audit

@.gemini/skills/autopus/security-audit.md

Detect security vulnerabilities in the codebase and suggest fixes. Includes OWASP Top 10 coverage.

### Pipeline

#### [REQUIRED] Step 1: Run govulncheck (MUST call Bash tool)

IMPORTANT: You MUST run `govulncheck` via the Bash tool. Do NOT skip static analysis.

```bash
govulncheck ./...
```

**Fallback rule**: Only if the Bash call returns an error (tool not installed, etc.), display the exact error to the user and ask for fallback approval. If user approves fallback, proceed with manual analysis only and note the gap.

> **⏭ POST-STEP**: govulncheck complete. NEXT REQUIRED STEP: Step 2: OWASP Analysis. Do NOT skip to Completion.

#### [REQUIRED] Step 2: OWASP Analysis

Analyze the codebase for OWASP Top 10 vulnerabilities relevant to the detected tech stack:
- A01: Broken Access Control
- A02: Cryptographic Failures
- A03: Injection
- A04: Insecure Design
- A05: Security Misconfiguration
- Others as applicable

### [REQUIRED] Pre-Completion Verification

Before displaying the security report, verify ALL steps were executed:

- [ ] Step 1: Bash tool called — `govulncheck ./...` ran (not skipped)
- [ ] Step 2: OWASP analysis completed for the detected tech stack

IF any item is unchecked → return to that step. Do NOT display the report.

Usage: `/auto secure`, `/auto secure --owasp`, `/auto secure --secrets`


---

## stale — Detect Stale Decisions

Detect outdated decisions, expired TODOs, and obsolete architecture patterns.

Detects:
- Lore entries older than 90 days
- Stale @AX:TODO tags
- Unused interfaces and redundant abstractions

Usage: `/auto stale`, `/auto stale --lore`, `/auto stale --ax`

---

## sync — Synchronize Documentation

Synchronize code and documentation after implementation is complete. Updates SPEC status and related documents.

### [REQUIRED] Sync Target 1: Update SPEC Status

If a SPEC-ID is provided, run SPEC Path Resolution for {SPEC-ID}. Update the resolved SPEC_PATH status to `completed`. Use TARGET_MODULE for git operations.

> **⏭ POST-TARGET**: SPEC status updated. NEXT REQUIRED TARGET: Sync Target 2: Project Docs. Do NOT skip to Lore Commit.

### [REQUIRED] Sync Target 2: Update Project Context Documents

Detect structural changes in the codebase and update:
- `ARCHITECTURE.md` — reflect new packages/dependencies
- `.autopus/project/structure.md` — reflect directory/file changes
- `.autopus/project/tech.md` — reflect dependency/build changes
- `.autopus/project/product.md` — reflect new features/commands
- `.autopus/project/canary.md` — reflect new health endpoints, pages, deploy config changes

> **⏭ POST-TARGET**: Project docs updated. NEXT REQUIRED TARGET: Sync Target 3: @AX Lifecycle. Do NOT skip to Lore Commit.

### [REQUIRED] Sync Target 3: @AX Tag Lifecycle Management

Manage @AX tag lifecycle (CYCLE tracking, TODO escalation, ANCHOR verification) — see rules below.

> **⏭ POST-TARGET**: @AX lifecycle managed. NEXT REQUIRED TARGET: Sync Target 4: CHANGELOG. Do NOT skip to Lore Commit.

### [REQUIRED] Sync Target 4: Update CHANGELOG.md

Append the change summary to `CHANGELOG.md` (create if absent).

### [REQUIRED] Sync Target 4.5: Learning Summary and Prune (R8, R9)

WHEN `.autopus/learnings/pipeline.jsonl` exists, THE SYSTEM SHALL:

**Step 1: Auto-Prune (R9)**

```bash
auto learn prune --max-age 90d
```

If the command fails, read the file directly and remove entries older than 90 days. Display only if entries were pruned:
```
정리: {N}개 학습 항목 만료 삭제 (90일 초과)
```

**Step 2: Learning Summary (R8)**

Display a learning summary BEFORE the completion bar:

```bash
auto learn summary --since-last-sync
```

If the command fails, generate from the JSONL file directly:
```
🐙 학습 요약 ────────────────────────
  신규 기록: {N}개 (gate_fail: {n}, review_issue: {n}, ...)
  반복 패턴 Top 3:
    1. "{pattern}" — {reuse_count}회 주입됨
    2. "{pattern}" — {reuse_count}회
    3. "{pattern}" — {reuse_count}회
  개선 영역: {이전 sync 대비 줄어든 실패 유형}
```

WHERE no learning entries exist, display: `학습 요약: 신규 항목 없음 ✓`

> **⏭ POST-TARGET**: Learning summary displayed. NEXT REQUIRED TARGET: Sync Target 5. Do NOT skip to Completion.