# auto setup — Generate/Update Project Context Documents

> [!NOTE]
> Global Connectivity: [Registry](../SKILL.md) | [Idea](idea.md) | [Plan](plan.md)

## Pipeline

### [REQUIRED] Step 1: Analyze Codebase (MUST call Agent tool)

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
    - Root repo role (product repo / monorepo root / meta workspace) and nested git repo boundaries
    - Generated surface / runtime artifact paths and source-of-truth repo locations
  """
)
```

> **⏭ POST-STEP**: Explorer returned. NEXT REQUIRED STEP: Step 2: Generate ARCHITECTURE.md. Do NOT skip to Completion.

### [REQUIRED] Step 2: Generate/Update ARCHITECTURE.md

Write or overwrite `ARCHITECTURE.md` based on the explorer's analysis. Include domains, layers, dependency map, and rule violations.

> **⏭ POST-STEP**: ARCHITECTURE.md written. NEXT REQUIRED STEP: Step 3: Generate project docs. Do NOT skip to Completion.

### [REQUIRED] Step 3: Generate/Update Project Docs

Create or update the 4 files under `.autopus/project/` using the explorer's output:
- `product.md` — project name, description, core features, use cases, mode
- `structure.md` — directory structure, package roles, entry points, file stats
- `tech.md` — tech stack, build, testing, architecture patterns
- `workspace.md` — root repo role, nested repo boundaries, generated/runtime paths, tracking/commit policy, source-of-truth locations

> **⏭ POST-STEP**: Project docs written. NEXT REQUIRED STEP: Step 4: Generate scenarios.md. Do NOT skip to Completion.

### [REQUIRED] Step 4: Generate/Update scenarios.md

Analyze codebase to extract user-facing E2E test scenarios (SPEC-E2E-001 R1).

### [REQUIRED] Step 5: Generate/Update canary.md

Analyze the project to auto-generate `.autopus/project/canary.md`.

### [REQUIRED] Pre-Completion Verification

Before displaying the completion message, verify ALL steps were executed:

- [ ] Step 1: Explore agent called — codebase analyzed
- [ ] Step 2: ARCHITECTURE.md written/updated
- [ ] Step 3: `.autopus/project/` docs written
- [ ] Step 4: scenarios.md written
- [ ] Step 5: canary.md written
