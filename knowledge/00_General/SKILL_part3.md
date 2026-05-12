| `--skip-prd` | Skip PRD generation and go straight to SPEC. Recommended for MEDIUM difficulty tasks. |
| `--prd-mode <mode>` | PRD mode: `standard` (10 sections, default for HIGH) or `minimal` (5 sections). |
| `--strategy` | Multi-provider strategy: consensus (default), debate, pipeline, fastest (requires `--multi`) |
| `--target <module>` | Target submodule for SPEC storage (e.g., `autopus-adk`). Auto-detected if omitted. |

Global flags `--auto` and `--multi` also apply:
- `--auto` → skip confirmations
- `--multi` → enable multi-provider review after SPEC generation

Usage: `/auto plan "feature description"`, `/auto plan --from-idea BS-001`, `/auto plan "feature" --skip-prd`, `/auto plan "feature" --prd-mode minimal`, `/auto plan "feature" --multi --strategy debate`

### Pipeline (execute in this exact order)

#### [REQUIRED] Step 1: Parse Flags

Global flags (`--auto`, `--loop`, `--multi`, `--quality`) are already stripped. Extract plan-specific flags:
- Extract `--from-idea <BS-ID>` → `FROM_IDEA = BS-ID` (load brainstorm context)
- Extract `--skip-prd` → `SKIP_PRD = true/false` (default: false)
- Extract `--prd-mode <value>` → `PRD_MODE = value` (default: standard)
- Extract `--strategy <value>` → `STRATEGY = value` (default: consensus)
- Extract `--target <module>` → `TARGET_MODULE` (auto-detect if omitted)
- `MULTI_FLAG` ← inherited from global `--multi`
- Remaining text → `FEATURE_DESC`

#### [CONDITIONAL] Step 1.5: Generate PRD

@.gemini/skills/autopus/prd.md

```
┌─────────────────────────────────────────────┐
│ SKIP_PRD == true?                           │
├──────────┬──────────────────────────────────┤
│ TRUE     │ FALSE                            │
│ ↓        │ ↓                                │
│ Step 2   │ Generate PRD, then Step 2        │
│ [SKIP]   │                                  │
└──────────┴──────────────────────────────────┘
```

WHEN `SKIP_PRD = false`, THE SYSTEM SHALL generate a PRD before spawning the spec-writer:

1. **Determine PRD mode**: Use `PRD_MODE` (default: `standard`). If `--prd-mode` is not set, auto-select based on scope:
   - Single package / small change → `minimal` (5 sections)
   - Multi-package / new feature / API design → `standard` (10 sections)

2. **Spawn planner agent** to generate the PRD:

```
Agent(
  subagent_type = "planner",
  prompt = """
    Phase: PRD Generation (Step 1.5)
    Feature description: {FEATURE_DESC}
    Target module: {TARGET_MODULE or auto-detect}
    PRD mode: {PRD_MODE}
    Template: templates/shared/prd-{PRD_MODE}.md.tmpl
    Brainstorm context: {BS file content if FROM_IDEA is set, otherwise omit}

    Generate a PRD following the prd skill (prd.md):
    1. Analyze the request (What/Why/Who/When)
    2. Collect codebase context (related files, existing patterns, prior SPECs)
    3. Author PRD sections using the {PRD_MODE} template
    4. Run quality validation checklist
    5. Save to {TARGET_MODULE}/.autopus/specs/SPEC-{ID}/prd.md

    Return: SPEC-ID, PRD file path, and quality checklist result.
  """
)
```

3. **Extract SPEC-ID** from the planner's output. This SPEC-ID will be reused by the spec-writer in Step 2.

WHEN `SKIP_PRD = true` → [INTENDED SKIP: Step 1.5]. Proceed directly to Step 2.

> **⏭ POST-STEP**: PRD generated (or skipped). NEXT REQUIRED STEP: Step 2: Spawn spec-writer Agent. Do NOT skip to Completion.

#### [REQUIRED] Step 2: Spawn spec-writer Agent

If `FROM_IDEA` is set, search for the BS file: check `.autopus/brainstorms/BS-{FROM_IDEA}.md` (top-level) then `*/.autopus/brainstorms/BS-{FROM_IDEA}.md` (submodules). Include its content as brainstorm context.

If Step 1.5 generated a PRD, include the PRD file path and SPEC-ID as context for the spec-writer.

Spawn a `spec-writer` agent to generate the SPEC:

```
Agent(
  subagent_type = "spec-writer",
  prompt = """
    Feature description: {FEATURE_DESC}
    Target module: {TARGET_MODULE or auto-detect}
    Project directory: {current directory}
    Brainstorm context: {BS file content if FROM_IDEA is set, otherwise omit}
    PRD context: {PRD file path and SPEC-ID if Step 1.5 was executed, otherwise omit}
  """
)
```

Extract the **SPEC-ID** (e.g., SPEC-UX-001) from the agent's output. If Step 1.5 already created a SPEC-ID, the spec-writer MUST reuse the same SPEC-ID.

> **⏭ POST-AGENT**: spec-writer returned. NEXT REQUIRED STEP: [CONDITIONAL] Step 3: Review Gate Decision. Do NOT skip to Completion.

#### [CONDITIONAL] Step 3: Review Gate Decision

```
┌─────────────────────────────────────────────┐
│ MULTI_FLAG == true                          │
│   OR autopus.yaml → review_gate.enabled?    │
├──────────┬──────────────────────────────────┤
│ TRUE     │ FALSE                            │
│ ↓        │ ↓                                │
│ Step 4   │ Step 5 [INTENDED SKIP: Step 4]   │
└──────────┴──────────────────────────────────┘
```

**How to check:**
- Run `cat autopus.yaml | grep -A2 "review_gate"` via Bash to check the enabled value
- Or Read `autopus.yaml` directly

#### [CONDITIONAL] Step 4: Run Multi-Provider Review

Run the `auto spec review` CLI command via the Bash tool:

```bash
auto spec review {SPEC-ID} --strategy {STRATEGY}
```

**Handling results:**
- **PASS** → Update Status in the resolved SPEC_PATH to `approved` (use Edit tool)
- **REVISE** → Print finding list, apply fixes, and re-run review (up to 2 iterations)
- **REJECT** → Print finding list and guide the user to redesign

**Fallback:**
- If `auto spec review` fails (provider not installed, etc.):
  - Print warning: "Multi-provider review could not be executed. Please check your provider configuration."
  - Keep Status as `draft`
  - Proceed to Step 5

#### [REQUIRED] Pre-Completion Verification

Before displaying completion output, verify ALL steps were evaluated:

- [ ] Step 1: Parse Flags — completed
- [ ] Step 1.5: Generate PRD — completed OR [INTENDED SKIP: --skip-prd]
- [ ] Step 2: Spawn spec-writer Agent — completed, SPEC-ID extracted
- [ ] Step 3: Review Gate Decision — evaluated (result: Step 4 or INTENDED SKIP)
- [ ] Step 4: Multi-Provider Review — completed OR [INTENDED SKIP]

IF any item is unchecked → return to that step. Do NOT display Completion Guidance.

#### [REQUIRED] Step 5: Completion Guidance

Display the workflow lifecycle bar:

```
🐙 Workflow: {SPEC-ID}
  ● plan  →  ○ go  →  ○ sync
```

Then show state-aware next step guidance:

- If Status is `approved`:
  ```
  ✓ SPEC {SPEC-ID} approved
  다음 단계: /auto go {SPEC-ID}
  ```
- If Status is `draft` (review skipped or not run):
  ```
  SPEC {SPEC-ID} 생성됨 (status: draft)
  다음 단계: /auto go {SPEC-ID}
  리뷰 실행: /auto spec review {SPEC-ID}
  ```

After displaying, also show agent result summary:

```
🐙 spec-writer ──────────────────────
  SPEC: {SPEC-ID} | 파일: 4개 | 요구사항: {N}개
  다음: /auto go {SPEC-ID}
```

---

## go — Implement a SPEC

@.gemini/skills/autopus/tdd.md
@.gemini/skills/autopus/agent-pipeline.md
@.gemini/skills/autopus/worktree-isolation.md

Implement code based on a SPEC document. Follows TDD methodology.

### Step 0: Parse Flags (REQUIRED — do this FIRST)

IMPORTANT: You MUST parse these flags from $ARGUMENTS before any other action.

Global flags (`--auto`, `--loop`, `--multi`, `--quality`) are already stripped. Extract the remaining go-specific flags:
- `--team` → `TEAMS_MODE = true/false` (Agent Teams)
- `--solo` → `SOLO_MODE = true/false` (single session, no subagents)
- `--strategy <value>` → `STRATEGY = value`
- `--continue` → `CONTINUE_MODE = true/false`
- `--skip-scaffold` → `SKIP_SCAFFOLD = true/false`
- Remaining text (non-flag word) → `SPEC_ID`

Inherit from global flags:
- `AUTO_MODE` ← `--auto`
- `LOOP_MODE` ← `--loop`
- `MULTI_MODE` ← `--multi`
- `QUALITY` ← `--quality`

Determine execution mode:
- `TEAMS_MODE = true` → Agent Teams mode
- `SOLO_MODE = true` → Single session mode
- Otherwise → **Subagent pipeline** (default)

After parsing, display the detected configuration:

```
Mode: {subagent | teams | teams+multi | solo}
SPEC: {SPEC_ID}
Target: {TARGET_MODULE}
Quality: {ultra | balanced | pending selection}
```

### Step 1: Load SPEC

Run SPEC Path Resolution for {SPEC_ID}. Load the resolved SPEC_PATH and check status. Extract TARGET_MODULE and WORKING_DIR for subsequent phases.

**Draft Guard**: If status is "draft" and review_gate is enabled, refuse implementation and prompt:
```
SPEC {SPEC_ID} is in draft status. Run review first: /auto spec review {SPEC_ID}
```

### Step 2: Route to Pipeline

Route based on the execution mode determined in Step 0.

#### Route A: Subagent Pipeline (default, no --team, no --solo)

IMPORTANT: This is the DEFAULT route. It REQUIRES spawning subagents using the Agent tool. Each Phase below MUST use an Agent() call — do NOT implement these phases yourself in the main session.

For parallel tasks, include `isolation: "worktree"` in Agent() calls to enable worktree isolation.

#### Route B: Agent Teams (`--team`)

IMPORTANT: This route uses Gemini CLI Agent Teams (experimental). Requires `CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS=1`.

If `CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS=1` is not set, abort with:
```
Agent Teams requires CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS=1. Use default mode (subagent) instead.
```

Use TeamCreate to create a team, then spawn specialized teammates using `Agent(subagent_type=..., team_name=..., name=...)`. Each teammate loads its agent definition from `.gemini/agents/autopus/`, inheriting tools, skills, model, and domain expertise. Teammates communicate directly via SendMessage.

Team composition — each member maps 1:1 to a specialized agent definition:

**Core members (always spawned):**
| Name | subagent_type | Agent Definition | Phase |
|------|---------------|------------------|-------|
| `lead` | `planner` | planner.md | Phase 1, coordination |
| `builder-1` | `executor` | executor.md | Phase 2 |
| `tester` | `tester` | tester.md | Phase 1.5, Phase 3 |
| `guardian` | `validator` | validator.md | Gate 2 |

**Phase-dependent members (spawned when needed):**
| Name | subagent_type | Agent Definition | Condition |
|------|---------------|------------------|-----------|
| `builder-2` | `executor` | executor.md | 2+ parallel tasks |
| `annotator` | `annotator` | annotator.md | Phase 2.5 |
| `auditor` | `security-auditor` | security-auditor.md | Phase 4 |
| `reviewer` | `reviewer` | reviewer.md | Phase 4 |

Quality Mode applies to execution profile selection: Ultra forces the premium path for all builders. Balanced keeps per-role defaults, and Adaptive Quality only escalates HIGH-complexity tasks. LOW-complexity work stays on the standard path in this workspace.

Reference: `.gemini/skills/autopus/agent-teams.md`

The pipeline phases are the same as Route A, but orchestrated through Agent Teams instead of individual Agent() calls. Worktree safety rules R1-R5 apply equally — each builder teammate works in an independent worktree. See @.gemini/skills/autopus/worktree-isolation.md and `.gemini/rules/autopus/worktree-safety.md`.

#### Route C: Single Session (`--solo`)

1. **TDD Implementation**: Execute RED → GREEN → REFACTOR based on tasks in plan.md
2. **Post-Implementation**: Validate against acceptance.md criteria and update status to done

#### Subagent Pipeline Phases (Route A)

```
Phase 0.5: Permission    → detect      (auto permission detect)
Phase 1: Planning        → planner     (mode: plan)
Phase 1.5: Test Scaffold → tester      (mode: bypassPermissions) — skip if --skip-scaffold
Gate 1:  Approval        → skip if AUTO_MODE
Phase 1.8: Doc Fetch     → main session (Context7 MCP) — skip if no external libs detected
Phase 2: Implementation  → executor×N  (mode: bypassPermissions, parallel with worktree isolation)
Phase 2.1: Worktree Merge → main session (merge worktree branches into working branch)
Gate 2:  Validation      → validator   (mode: plan)  — retry up to 3× on FAIL
Phase 2.5: Annotation    → annotator   (mode: bypassPermissions) — @AX tags on modified files
Phase 3: Testing         → tester      (mode: bypassPermissions)
Gate 3:  Coverage        → verify 85%+ coverage
Phase 4: Review          → reviewer + security-auditor (parallel, mode: plan)  — retry up to 2× on REQUEST_CHANGES
```

### Retry Limits Reference
