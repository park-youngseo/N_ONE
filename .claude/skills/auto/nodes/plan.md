# auto plan — Write a SPEC

> [!NOTE]
> Global Connectivity: [Registry](../SKILL.md) | [Idea](idea.md) | [Go](go.md)

Write a SPEC document. Delegate to the `spec-writer` agent.

## Flags

| Flag | Description |
|------|-------------|
| `--from-idea <BS-ID>` | Create SPEC from a brainstorm file. |
| `--skip-prd` | Skip PRD generation (for MEDIUM difficulty). |
| `--prd-mode <mode>` | `standard` or `minimal`. |
| `--strategy` | consensus (default), debate, etc. |

## Pipeline

### [REQUIRED] Step 1: Parse Flags

Extract `--from-idea`, `--skip-prd`, `--prd-mode`, `--strategy`, `--target`.

### [CONDITIONAL] Step 1.5: Generate PRD

WHEN `SKIP_PRD = false`, spawn planner agent to generate PRD.
Save to `{TARGET_MODULE}/.autopus/specs/SPEC-{ID}/prd.md`.

### [REQUIRED] Step 2: Spawn spec-writer Agent

Spawn a `spec-writer` agent to generate the SPEC.
Include BS/PRD context if available.

### [CONDITIONAL] Step 3: Review Gate Decision

Check `autopus.yaml` → `review_gate.enabled`.

### [CONDITIONAL] Step 4: Run Multi-Provider Review

```bash
auto spec review {SPEC-ID} --strategy {STRATEGY}
```

### [REQUIRED] Step 5: Completion Guidance

Display workflow lifecycle bar: `● plan → ○ go → ○ sync`
Guide to `/auto go {SPEC-ID}`.
