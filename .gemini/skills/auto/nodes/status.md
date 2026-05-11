# Node: auto status
## Global Connectivity Table
| Node | Role | Connectivity |
| :--- | :--- | :--- |
| SKILL.md | Root Router | Calls this node for `status` |
| setup.md | Context Generation | Provides architecture context for status |
| status.md | SPEC Dashboard | Scans `.autopus/specs/` for lifecycle state |

## Description
Display the SPEC dashboard for the current project and submodules.

## Functionality
1. **SPEC Scan**: Identify all `spec.md` files in the workspace.
2. **State Analysis**: Parse `status` field (active, completed, draft).
3. **Dashboard Display**:
    - Project Summary Table
    - Module-wise SPEC Breakdown
    - Completion Percentage (Burndown)
4. **Submodule Support**: Recursively scan submodules and aggregate statuses.

## State Definitions
| State | Description |
| :--- | :--- |
| `draft` | Under initial development |
| `active` | Currently being implemented/refined |
| `completed` | Fully implemented and verified |
| `stale` | Requires review due to architectural changes |
