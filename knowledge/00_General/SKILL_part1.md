---
name: auto
description: Autopus-ADK 메인 명령어 — 개발 워크플로우, 품질 검증, 탐색, 관리 서브커맨드 라우터
---

# /auto — Autopus-ADK

> 🐙 Autopus v0.4 | autopus-adk | full

At the start of every response, output the following banner:

```
🐙 Autopus ─────────────────────────
```

Then continue with the task output. End each subcommand response with `🐙`.

ARGUMENTS: $ARGUMENTS

## Context Load (on every session start)

Before processing any subcommand, load project context by reading these files if they exist:

1. `ARCHITECTURE.md` — architecture map (domains, layers, dependencies)
2. `.autopus/project/product.md` — project overview, core features
3. `.autopus/project/structure.md` — directory structure, package roles
4. `.autopus/project/tech.md` — tech stack, build, testing
5. `.autopus/project/scenarios.md` — E2E test scenarios (if exists)
6. `.autopus/project/canary.md` — health check configuration (if exists)
7. `.autopus/context/signatures.md` — exported API signatures (if exists)
7. `.autopus/learnings/pipeline.jsonl` — pipeline learning patterns (if exists)

WHEN `.autopus/learnings/pipeline.jsonl` exists AND contains 5+ entries, display a non-intrusive notification after loading context:

```
💡 학습 패턴 {N}개 — 다음 파이프라인에 자동 반영됩니다
```

If none of these files exist, display:

```
No project context documents found. Run `/auto setup` to initialize.
```

## SPEC Path Resolution

WHEN any subcommand receives a SPEC-ID, resolve the actual file path using this standard procedure:

1. Check `.autopus/specs/{SPEC-ID}/spec.md` (top-level — legacy and cross-module SPECs)
2. Glob `*/.autopus/specs/{SPEC-ID}/spec.md` (submodule depth 1)

From the resolved path, extract:
- `SPEC_PATH`: full relative path to spec.md
- `SPEC_DIR`: SPEC directory path
- `TARGET_MODULE`: submodule path (or "." if top-level)
- `WORKING_DIR`: same as TARGET_MODULE — the directory where build/test commands run

Error handling:
- 0 results → error: "SPEC-{ID} not found. Available SPECs:" then list all found via `*/.autopus/specs/*/spec.md`
- 2+ results → error: "Duplicate SPEC-{ID} found:" then list each path

All subcommands that reference a SPEC-ID MUST use this resolution procedure instead of hardcoded paths.

## Subcommand Routing

**Strip global flags first** from $ARGUMENTS, then determine the subcommand from the first remaining word.

### Global Flags

| Flag | Description |
|------|-------------|
| `--think` | Activate Sequential Thinking MCP for deep analysis mode. Load the `mcp__sequential-thinking__sequentialthinking` tool via ToolSearch, then perform step-by-step reasoning. |
| `--ultrathink` | Same as `--think` but with higher reasoning effort. |
| `--auto` | Autonomous mode: skip user confirmations and approval gates. Quality defaults to `autopus.yaml → quality.default` (not hardcoded). |
| `--loop` | RALF loop mode: auto-retry failed quality gates with extended iteration limits until PASS or circuit break. |
| `--multi` | Multi-provider mode: activate orchestra engine for reviews and decisions. |
| `--quality <value>` | Agent quality mode: `ultra` (all agents Opus) or `balanced` (mixed). Overrides `autopus.yaml → quality.default`. |

### Routing Rules

1. **Strip global flags**: Separate `--think`, `--ultrathink`, `--auto`, `--loop`, `--multi`, `--quality <value>`, and other global flags from $ARGUMENTS
2. **Match subcommand**: Use the first remaining word to determine the subcommand
3. **Triage-based Smart Routing (natural language → difficulty-aware flow)**:
   - If the first word does not match a known subcommand and the remaining text has 2+ words:
   - Execute the **Triage Process** (see below) to assess difficulty and recommend a flow
   - Route to the recommended flow after user confirmation (or auto-proceed with `--auto`)
4. **Empty arguments**: Show the subcommand list

### Triage Process

WHEN a natural language request is detected (rule 3 above), THE SYSTEM SHALL assess task difficulty before routing.

#### Step 1: Signal Extraction

Analyze the request text for the following signals:

| Signal | IDEA indicators | LOW indicators | MEDIUM indicators | HIGH indicators |
|--------|----------------|---------------|-------------------|-----------------|
| **Change type** | 고민, 생각, idea, brainstorm, 가능할까, 어떻게, 테스트해보고싶어, 탐색, explore | fix, 수정, 고쳐, typo, rename, 오타, 삭제, 로그추가, config | 리팩토링, refactor, 개선, improve, 옵션추가, 확장, extend | 새로운, new, feature, 기능, 모듈, 시스템, 아키텍처, API설계 |
| **Scope** | unclear or exploratory; no specific file/package | specific file/function mentioned | package/module mentioned | "전체", "모든", multi-domain, cross-cutting |
| **Verb intensity** | think, consider, wonder, explore, 고민, 검토 | change, update, move, remove | add, enhance, restructure | design, build, create, implement (large scope) |

#### Step 2: Impact Scan (fast, ≤ 5 seconds)

Run a quick codebase scan to validate the heuristic:

1. **Grep** for keywords from the request to estimate affected files
2. **Glob** for patterns matching the described scope
3. Count: affected files, affected packages

This step is OPTIONAL if the text-based signals are unambiguous. Skip if `--auto` flag is set.

#### Step 3: Difficulty Classification

| Difficulty | Criteria (ANY of) | Recommended Flow |
|-----------|-------------------|------------------|
| **IDEA** | Exploratory question; "고민", "어떻게 하면", "가능할까"; feasibility check; brainstorm; no clear implementation target | `/auto idea "{desc}"` — brainstorm and evaluate |
| **LOW** | Single file change; bug fix with known location; config/doc edit; ≤ 1 package affected | `/auto fix "{desc}"` — direct fix, no SPEC |
| **MEDIUM** | 2-3 files; existing feature extension; single-package refactor; test additions | `/auto plan "{desc}" --skip-prd` — SPEC only, no PRD |
| **HIGH** | 3+ files across packages; new feature/module; API design; architecture change; multi-domain impact | `/auto idea "{desc}" --multi` — multi-provider brainstorm → ICE 검증 → plan 체이닝 |
| **CRITICAL** | Security vulnerability; urgent production fix | `/auto fix "{desc}"` + recommend `/auto secure` after |

#### Step 4: Triage Display and Confirmation

WHEN `--auto` is NOT set:

1. Display the triage analysis as text:

```
🐙 Triage ───────────────────────────
  요청: "{natural language request}"

  분석:
  - 변경 유형: {bug fix | enhancement | new feature | refactor | config}
  - 예상 영향: ~{N}개 파일 / {N}개 패키지
  - 난이도: {LOW | MEDIUM | HIGH | CRITICAL}

  추천: {recommended flow description}
```

2. Use the `AskUserQuestion` tool to present the selection (do NOT use text-based numbered options):

```
AskUserQuestion(
  questions = [{
    question: "어떤 플로우로 진행할까요?",
    header: "Triage",
    multiSelect: false,
    options: [
      { label: "아이디어 탐색", description: "/auto idea --multi — 멀티 프로바이더 브레인스토밍 + ICE 검증" },
      { label: "바로 수정", description: "/auto fix — SPEC 없이 직접 수정" },
      { label: "SPEC만 생성", description: "/auto plan --skip-prd — PRD 생략, SPEC만 작성" },
      { label: "PRD + SPEC 전체", description: "/auto plan — PRD 작성 후 SPEC 생성" }
    ]
  }]
)
```

Adjust the recommended option (add "(Recommended)" suffix) based on the difficulty classification. Place the recommended option FIRST in the options list.

**HIGH difficulty routing**: WHEN difficulty is HIGH, THE SYSTEM SHALL place "아이디어 탐색" as the FIRST option with "(Recommended)" suffix. The rationale: HIGH-impact features should pass through multi-provider brainstorm and ICE validation before committing to a SPEC. This serves as an upstream decision gate — "이게 정말 만들 가치가 있는가?" — preventing premature SPEC creation for ideas that haven't been validated.

WHEN `--auto` IS set, display triage result and auto-proceed:

```
🐙 Triage ───────────────────────────
  요청: "{request}"
  난이도: {level} → {flow} (자동 진행)
```

**`--auto` HIGH routing**: WHEN `--auto` AND difficulty is HIGH, THE SYSTEM SHALL auto-proceed with `/auto idea "{desc}" --multi --auto` (which auto-chains to `/auto plan --from-idea BS-{ID}` via idea's `--auto` flag). This ensures HIGH-impact features always pass through the upstream decision gate even in autonomous mode.

#### Triage Override

The user can always force a specific flow by using the subcommand directly:
- `/auto idea "..."` — skip triage, go directly to idea brainstorm
- `/auto fix "..."` — skip triage, go directly to fix
- `/auto plan "..."` — skip triage, go directly to plan
- `/auto plan "..." --skip-prd` — skip triage, plan without PRD

Triage ONLY activates for bare natural language input without a recognized subcommand.

### Subcommand List

When displaying the subcommand list (empty `/auto` invocation), use the categorized format below.

**개발 워크플로우**

| Subcommand | Description |
|-----------|-------------|
| idea ⚡ | Brainstorm and evaluate ideas |
| plan | Write a SPEC |
| go | Implement a SPEC |
| sync | Synchronize documentation |
| fix ⚡ | Fix a bug (no SPEC needed) |
| dev | Full cycle: plan → go → sync |

**품질 & 리뷰** ⚡ = SPEC 없이 독립 실행 가능

| Subcommand | Description |
|-----------|-------------|
| review ⚡ | Code review (TRUST 5 criteria) |
| spec review | SPEC multi-provider review |
| secure ⚡ | Security audit (OWASP Top 10) |
| stale ⚡ | Detect stale decisions |
| verify ⚡ | Frontend UX verification |
| browse ⚡ | Browser automation (cmux/agent-browser auto-detect) |
| test | Run E2E scenarios |
| canary ⚡ | Post-deploy health check (build + E2E + browser) |

**탐색 & 분석** ⚡ = 독립 실행 가능

| Subcommand | Description |
|-----------|-------------|
| map ⚡ | Analyze codebase structure |
| why ⚡ | Query decision rationale |
| status ⚡ | SPEC dashboard |

**관리**

| Subcommand | Description |
|-----------|-------------|
| setup ⚡ | Generate project context — **첫 사용 시 추천** |
| init | Initialize harness |
| update | Update harness |
| doctor ⚡ | Health diagnostics |
| platform | Platform management |

> 자연어로 작업을 설명하면 난이도를 자동 분석하여 적절한 플로우를 추천합니다.
> ⚡ 표시된 커맨드는 SPEC이나 파이프라인 없이 바로 사용할 수 있습니다.

---

## setup — Generate/Update Project Context Documents

Analyze the project's architecture, structure, and tech stack to produce context documents.
Agents are re-initialized on every session, so these documents provide continuity across sessions.

### Pipeline

#### [REQUIRED] Step 1: Analyze Codebase (MUST call Agent tool)

IMPORTANT: You MUST spawn an Explore agent to analyze the project. Do NOT skip to Step 2 without calling Agent tool.

```
Agent(
  subagent_type = "explorer",
  mode = PERMISSION_MODE == "bypass" ? "bypassPermissions" : "plan",
  prompt = """
    Explore the full project structure. Return:
    - Directory layout and package roles
    - Tech stack (languages, frameworks, build tools)
    - Entry points and exported APIs
    - Architecture patterns and domain boundaries
  """
)
```

> **⏭ POST-STEP**: Explorer returned. NEXT REQUIRED STEP: Step 2: Generate ARCHITECTURE.md. Do NOT skip to Completion.

#### [REQUIRED] Step 2: Generate/Update ARCHITECTURE.md

Write or overwrite `ARCHITECTURE.md` based on the explorer's analysis. Include domains, layers, dependency map, and rule violations.

> **⏭ POST-STEP**: ARCHITECTURE.md written. NEXT REQUIRED STEP: Step 3: Generate project docs. Do NOT skip to Completion.

#### [REQUIRED] Step 3: Generate/Update Project Docs

Create or update the 4 files under `.autopus/project/` using the explorer's output:
- `product.md` — project name, description, core features, use cases, mode
- `structure.md` — directory structure, package roles, entry points, file stats
- `tech.md` — tech stack, build, testing, architecture patterns

> **⏭ POST-STEP**: Project docs written. NEXT REQUIRED STEP: Step 4: Generate scenarios.md. Do NOT skip to Completion.

#### [REQUIRED] Step 4: Generate/Update scenarios.md

Analyze codebase to extract user-facing E2E test scenarios (SPEC-E2E-001 R1).

### Output Files

| File | Content |
|------|---------|
| `ARCHITECTURE.md` | Domains, layers, dependency map, rule violations |
| `.autopus/project/product.md` | Project name, description, core features, use cases, mode |
| `.autopus/project/structure.md` | Directory structure, package roles, entry points, file stats |
| `.autopus/project/tech.md` | Tech stack, build, testing, architecture patterns |
| `.autopus/project/scenarios.md` | User-facing E2E test scenarios extracted from the codebase |
| `.autopus/project/canary.md` | Health check configuration: build, endpoints, browser targets, deploy platform |

### Scenario Generation (Step 4)

Analyze the project codebase to generate `.autopus/project/scenarios.md`.

#### Step 4.1: Detect Project Type

Scan the project root to determine the primary stack. Multiple types can coexist (e.g., Go CLI + API).

| Project Type | Detection Signals | Extraction Target |
|---|---|---|