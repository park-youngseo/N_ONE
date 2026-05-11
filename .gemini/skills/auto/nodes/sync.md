# Node: auto sync
## Global Connectivity Table
| Node | Role | Connectivity |
| :--- | :--- | :--- |
| SKILL.md | Root Router | Calls this node for `sync` |
| go.md | Implementation | Completes implementation before `sync` |
| sync.md | Documentation | Synchronizes SPEC, Docs, and Learnings |

## Description
Synchronize code and documentation after implementation. Updates SPEC status and project context.

## Sync Targets
1. **SPEC Status**: Update to `completed`.
2. **Project Docs**: Update `ARCHITECTURE.md`, `product.md`, `tech.md`, etc.
3. **@AX Lifecycle**: Manage CYCLE tracking, TODO escalation, and ANCHOR verification.
4. **Learning Summary**: Prune old learnings (90d) and display summary.
5. **Lore Commit**: Execute 2-Phase commit (Module + Meta repos).

## @AX Lifecycle Rules
| Tag | Action | Condition |
| :--- | :--- | :--- |
| `@AX:TODO` | Increment CYCLE | Every sync iteration |
| `@AX:WARN` | Escalate | If CYCLE >= 3 |
| `@AX:ANCHOR` | Verify Fan-in | If callers < 3, downgrade to NOTE |

## 2-Phase Commit
- **Phase A (Module)**: Commit changes in `{TARGET_MODULE}/`.
- **Phase B (Meta)**: Commit project-level docs at root.
