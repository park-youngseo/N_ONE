# Node: auto management
## Global Connectivity Table
| Node | Role | Connectivity |
| :--- | :--- | :--- |
| SKILL.md | Root Router | Calls this node for management commands |
| manage.md | ADK Management | `init`, `update`, `doctor`, `platform` |

## Description
Manage the Autopus-ADK harness and platform configurations.

## Commands
| Command | Action | Description |
| :--- | :--- | :--- |
| `init` | Initialize | Install fresh harness files |
| `update` | Sync | Update harness to latest version |
| `doctor` | Diagnose | Health check for providers and wiring |
| `platform` | Config | Manage supported platforms (claude, codex, etc.) |

## Diagnostics (Doctor)
- **Check 1**: Harness file integrity (checksums).
- **Check 2**: Provider connectivity (API keys, binaries).
- **Check 3**: Git worktree status.
- **Check 4**: Environment dependencies (Node, Go, Python).
