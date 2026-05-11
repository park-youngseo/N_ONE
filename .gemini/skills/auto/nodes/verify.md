# Node: auto verify
## Global Connectivity Table
| Node | Role | Connectivity |
| :--- | :--- | :--- |
| SKILL.md | Root Router | Calls this node for `verify` |
| verify.md | UX Verification | VLM analysis & Playwright tests |
| browse.md | Browser | Provides browser interaction primitives |

## Description
Verify frontend UX changes using VLM-powered visual analysis and Playwright E2E tests.

## Pipeline Phases
1. **Phase 0: Change Analysis**: Run `git diff` to identify UI impact.
2. **Phase 1: Test Gen/Heal**: Generate/repair Playwright tests.
3. **Phase 2: Test Runner**: Execute `npx playwright test` (Bash tool).
4. **Phase 3: VLM Verifier**: Semantic visual analysis on screenshots.
5. **Phase 4: Reporter**: Apply fixes and produce report.

## Verification Primitives
| Type | Primitives |
| :--- | :--- |
| CLI | `exit_code(N)`, `stdout_contains(str)` |
| API | `status_code(N)`, `json_path(path, value)` |
| Frontend | `element_visible(selector)`, `page_title(str)` |

## Flags
| Flag | Description | Default |
| :--- | :--- | :--- |
| `--fix` | Enable auto-fix mode | `true` |
| `--report-only` | Skip auto-fix, output findings | `false` |
| `--viewport` | desktop, mobile, or WxH | `desktop` |
