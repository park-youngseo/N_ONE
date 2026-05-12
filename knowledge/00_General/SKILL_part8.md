
### [REQUIRED] Sync Target 5: 2-Phase Lore Commit (MUST call Bash tool)

IMPORTANT: You MUST run git operations via the Bash tool. Do NOT skip the commit step.

The workspace may use a **meta repo** pattern: the root project directory is a git repo that tracks project-level documents, while each subdirectory can be an independent git repo. Sync commits MUST be split into two phases when this pattern is detected.

**Detection**: Run `git rev-parse --show-toplevel` at root AND at `{TARGET_MODULE}`. If they differ, this is a multi-repo workspace — use 2-Phase Commit. If they are the same (or TARGET_MODULE is '.'), use single-phase commit.

#### Phase A: Module Commit (submodule changes)

WHEN TARGET_MODULE is not '.', commit SPEC and module-specific files to the submodule repo:

```bash
git -C {TARGET_MODULE} status
git -C {TARGET_MODULE} add {SPEC_DIR relative to TARGET_MODULE} {module-specific changed files}
git -C {TARGET_MODULE} diff --cached --quiet || git -C {TARGET_MODULE} commit -m "$(cat <<'EOF'
docs(<scope>): sync SPEC-{ID} 문서 및 상태 갱신

SPEC 구현 완료에 따른 모듈 내 문서 동기화.

Why: SPEC 구현 완료 후 모듈 문서 일관성 유지
Decision: 모듈 repo에 SPEC 관련 파일만 커밋
Ref: SPEC-{ID}

🐙 Autopus <noreply@autopus.co>
EOF
)"
```

#### Phase B: Meta Commit (root-level documents)

WHEN multi-repo workspace is detected, ALWAYS check and commit root-level document changes to the meta repo:

```bash
git status
git add ARCHITECTURE.md .autopus/project/ CHANGELOG.md
git diff --cached --quiet || git commit -m "$(cat <<'EOF'
docs(meta): sync 루트 프로젝트 문서 갱신

프로젝트 컨텍스트 문서를 최신 코드 상태에 맞게 동기화.

Why: 서브모듈 변경에 따른 루트 문서 일관성 유지
Decision: meta repo에 프로젝트 전체 문서만 커밋

🐙 Autopus <noreply@autopus.co>
EOF
)"
```

**Phase skip rules:**
- Phase A: skip when TARGET_MODULE is '.' (root-only sync) or no module files changed
- Phase B: skip when not a multi-repo workspace, or `git diff --cached --quiet` (no root files changed)
- Both phases: skip when `git status` reports no changes — print a notice only

### [REQUIRED] Pre-Completion Verification

Before displaying the completion bar, verify ALL sync targets were completed:

- [ ] Sync Target 1: SPEC status updated (or N/A — no SPEC-ID provided)
- [ ] Sync Target 2: Project context docs updated
- [ ] Sync Target 3: @AX lifecycle managed
- [ ] Sync Target 4: CHANGELOG.md updated
- [ ] Sync Target 4.5: Learning summary and prune — completed (or N/A — no pipeline.jsonl)
- [ ] Sync Target 5: 2-Phase Lore Commit — Phase A (module) + Phase B (meta) ran (or skipped due to no changes)

IF any item is unchecked → return to that target. Do NOT display Completion Guidance.

### Project Document Update

On sync, detect structural changes in the codebase and update:
- `ARCHITECTURE.md` — reflect new packages/dependencies
- `.autopus/project/structure.md` — reflect directory/file changes
- `.autopus/project/tech.md` — reflect dependency/build changes
- `.autopus/project/product.md` — reflect new features/commands

Usage: `/auto sync SPEC-ID`, `/auto sync`

### @AX Lifecycle Management

On sync, manage @AX tag lifecycle across the codebase:

#### TODO Cycle Tracking
1. Find all `@AX:TODO` tags in source files
2. Increment `@AX:CYCLE:N` counter (or add `@AX:CYCLE:1` if absent)
3. Escalate to `@AX:WARN` if CYCLE >= 3, adding `@AX:REASON: escalated from TODO after N cycles`

#### ANCHOR Fan-In Verification
1. Find all `@AX:ANCHOR` tags
2. Grep for function name references to count callers
3. If caller count < 3, downgrade to `@AX:NOTE` with updated reason

#### NOTE Cleanup
1. Find all `@AX:NOTE` tags referencing specific functions
2. Verify the referenced function still exists in the codebase
3. Remove orphaned NOTE tags


### Lore Commit (2-Phase)

Sync Target 5 handles the actual commits. This section provides additional commit guidelines.

**2-Phase Commit flow:**
1. **Phase A** — Module repo: SPEC files and module-specific changes → `git -C {TARGET_MODULE} commit`
2. **Phase B** — Meta repo: root-level project docs → `git commit` (at workspace root)

**Determining which files go where:**
- Files inside `{TARGET_MODULE}/` → Phase A (module commit)
- Files at root (`ARCHITECTURE.md`, `.autopus/project/`, `CHANGELOG.md`) → Phase B (meta commit)
- If implementation code (`.go`, etc.) is included in Phase A: use the primary type from the SPEC (feat, fix, refactor)
- If only docs changed: use `docs` type for both phases

**Skip commit when:**
- No changes reported by `git status` in that repo — skip and print a notice only

### Completion Guidance

After sync completes, display the workflow lifecycle bar:

```
🐙 Workflow: {SPEC-ID}
  ✓ plan  →  ✓ go  →  ● sync
```

Then display the milestone celebration ResultBox for a completed full cycle:

```
╭────────────────────────────────────╮
│ 🐙 파이프라인 완료!                 │
│ {SPEC-ID}: {title}                 │
│ 태스크: N/N | 커버리지: N%          │
│ 리뷰: {verdict}                    │
╰────────────────────────────────────╯
```

**Field mappings:**
- `{SPEC-ID}` — e.g., `SPEC-UX-001`
- `{title}` — SPEC title from spec.md
- `N/N` — completed tasks / total tasks from plan.md
- `N%` — final test coverage (use `N/A` if not applicable or harness-only)
- `{verdict}` — reviewer verdict (`APPROVE` or `N/A` if review was skipped)

Then show state-aware next step guidance:

```
다음 단계: 다음 기능을 시작하려면 /auto plan "기능 설명" 을 실행하세요.
배포 후 검증: /auto canary
```

---

## why — Query Decision Rationale

Look up the rationale behind code or architecture decisions.

Query methods:
- Keyword search across Lore entries
- Extract decision context from commit history
- Look up NOTE/ANCHOR in @AX tags
- Link to related SPEC documents

Usage: `/auto why "why was this pattern used?"`, `/auto why path/to/file.go`

---

## status — SPEC Dashboard

Display the current status of all SPECs in the project.

### Output

Scan for all SPECs across the project:
1. `.autopus/specs/*/spec.md` (top-level — legacy and cross-module)
2. `*/.autopus/specs/*/spec.md` (submodules, depth 1)

Group by module and display:

```
🐙 SPEC Dashboard ───────────────────
  [autopus-adk]
    SPEC-ORCH-003  [draft]       Orchestra Provider
    SPEC-GAP-001   [approved]    Gap Analysis

  [Autopus]
    SPEC-AUTOTEST-001 [draft]    Auto Test

  [cross-module]
    SPEC-AI-001    [approved]    AI Integration
```

Top-level SPECs (`.autopus/specs/`) are grouped under `[cross-module]`.

Show a summary line at the bottom:

```
총 {N}개 | draft: {N} | approved: {N} | implemented: {N} | completed: {N}
```

Usage: `/auto status`

---

## dev — Full Development Cycle

Execute the complete plan → go → sync cycle in one command. The power-user workflow in a single invocation.

### Flag Inheritance

`/auto dev` automatically applies optimal flags to each stage:

| Stage | Inherited Flags |
|-------|----------------|
| plan | `--auto --multi --ultrathink` (deep analysis + multi-provider review) |
| go | `--auto --loop --team` (Agent Teams + RALF self-healing) |
| sync | (no special flags) |

User-provided flags **override** these defaults. For example, `--solo` overrides `--team` in the go stage.

### Pipeline

#### [REQUIRED] Stage 1: Plan (MUST call plan subcommand)

Run the `plan` pipeline with `--auto --multi --ultrathink` flags and the feature description. Extract the SPEC-ID from the output.

Display after completion:
```
🐙 Workflow: {SPEC-ID}
  ● plan  →  ○ go  →  ○ sync
```

> **⏭ POST-STAGE**: plan complete, SPEC-ID extracted. NEXT REQUIRED STAGE: Stage 2: Go. Do NOT skip to Completion.

#### [REQUIRED] Stage 2: Go (MUST call go subcommand)

Run the `go` pipeline with `--auto --loop --team` flags using the SPEC-ID from Stage 1.

Display after completion:
```
🐙 Workflow: {SPEC-ID}
  ✓ plan  →  ● go  →  ○ sync
```

> **⏭ POST-STAGE**: go complete. NEXT REQUIRED STAGE: Stage 3: Sync. Do NOT skip to Completion.

#### [REQUIRED] Stage 3: Sync (MUST call sync subcommand)

Run the `sync` pipeline with the SPEC-ID.

Display after completion:
```
🐙 Workflow: {SPEC-ID}
  ✓ plan  →  ✓ go  →  ● sync
```

### [REQUIRED] Pre-Completion Verification

Before displaying the final milestone output, verify ALL stages were completed:

- [ ] Stage 1: plan — SPEC-ID created
- [ ] Stage 2: go — implementation pipeline completed
- [ ] Stage 3: sync — documentation synchronized, Lore commit made

IF any item is unchecked → return to that stage. Do NOT display the milestone box.

### Flags

| Flag | Description |
|------|-------------|
| `--team` | Run `go` with Agent Teams — **default ON** in dev |
| `--solo` | Run `go` in single session mode (overrides --team) |
| `--quality` | Quality mode: ultra or balanced |
| `--multi` | Enable multi-provider review — **default ON** in dev |
| `--ultrathink` | Deep analysis in plan stage — **default ON** in dev |
| `--no-team` | Disable Agent Teams (use subagent pipeline instead) |
| `--no-multi` | Disable multi-provider review |
| `--no-ultrathink` | Disable deep analysis |

Usage: `/auto dev "feature description"`, `/auto dev "feature description" --solo`, `/auto dev "feature description" --no-multi`

**Default behavior**: `/auto dev "feature"` = `/auto plan --auto --multi --ultrathink` → `/auto go --auto --loop --team` → `/auto sync`


---

## ADK Management — init, update, doctor, platform

Manage the Autopus-ADK harness itself.

### Current Configuration

- **Project**: autopus-adk
- **Mode**: full
- **Platform**: claude-code, codex, gemini-cli, opencode

### Usage

- `/auto init` — initialize harness
- `/auto update` — update harness files
- `/auto doctor` — run health diagnostics
- `/auto platform list` — list platforms