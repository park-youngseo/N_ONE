# Node: auto idea
## Global Connectivity Table
| Node | Role | Connectivity |
| :--- | :--- | :--- |
| SKILL.md | Root Router | Calls this node for `idea` |
| idea.md | Brainstorming | Executes multi-provider debate/consensus |
| plan.md | SPEC Planning | Receives BS-{ID} from `idea` via `--from-idea` |

## Description
Brainstorm ideas using multi-provider orchestra, evaluate with ICE scoring, and save as BS file.

## Flags
| Flag | Description | Default |
| :--- | :--- | :--- |
| `--strategy` | Orchestra strategy (debate, consensus, pipeline, fastest) | `debate` |
| `--providers` | Provider list override | `all` |
| `--auto` | Auto-chain to `/auto plan` after completion | `false` |

## Pipeline Steps
1. **Parse Input**: Extract `IDEA_DESC`, `STRATEGY`, and `AUTO_MODE`.
2. **Structure Context**: Analyze as What/Why/Who/When.
3. **Call Orchestra**: Execute `auto orchestra brainstorm` via Bash tool.
4. **ICE Scoring**: Calculate Score = (Impact × Confidence × Ease) / 100.
5. **Persistence**: Save to `.autopus/brainstorms/BS-{ID}.md`.

## ICE Scoring Matrix
| Metric | Range | Weight |
| :--- | :--- | :--- |
| Impact | 1-10 | High value delivery |
| Confidence | 1-10 | Certainty of success |
| Ease | 1-10 | Simplicity of implementation |
