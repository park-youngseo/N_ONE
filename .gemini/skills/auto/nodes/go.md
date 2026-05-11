# Node: auto go
## Global Connectivity Table
| Node | Role | Connectivity |
| :--- | :--- | :--- |
| SKILL.md | Root Router | Calls this node for `go` |
| plan.md | SPEC Planning | Provides SPEC-{ID} for `go` |
| go.md | Implementation | Executes TDD pipeline (Phases 1-4) |
| sync.md | Documentation | Updates docs after `go` completion |

## Description
Implement code based on a SPEC document. Follows TDD methodology and subagent pipeline.

## Flags
| Flag | Description |
| :--- | :--- |
| `--team` | Agent Teams mode (Claude Code only) |
| `--solo` | Single session mode (no subagents) |
| `--loop` | Activate RALF loop (self-healing) |
| `--continue` | Resume from last checkpoint |

## Execution Phases (Subagent Pipeline)
| Phase | Agent | Mode | Task |
| :--- | :--- | :--- | :--- |
| 1: Planning | planner | plan | Break SPEC into atomic tasks |
| 1.5: Scaffold | tester | bypass | Generate test files & boilerplate |
| 2: Implement | executor | bypass | Concurrent implementation (Worktree) |
| 2.1: Merge | main | N/A | Merge worktree branches |
| Gate 2: Validate | validator | plan | Code quality & build verification |
| 3: Testing | tester | bypass | Run full test suite |
| 4: Review | reviewer | plan | TRUST 5 & Security audit |

## RALF Loop Protocol
| Step | Action | Description |
| :--- | :--- | :--- |
| **RED** | Execution | Run phase logic |
| **GREEN** | Validation | Judge PASS/FAIL |
| **REFACTOR** | Repair | Spawn fix agent on FAIL |
| **LOOP** | Iteration | Repeat until PASS or Circuit Break |

## Circuit Breaker Rules
- Triggered after 3 consecutive iterations with **no progress**.
- Progress measured by: Issue count decrease, Coverage increase.
- On break: Abort loop and report remaining issues to Architect.
