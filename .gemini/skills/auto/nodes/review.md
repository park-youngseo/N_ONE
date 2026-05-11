# Node: auto review
## Global Connectivity Table
| Node | Role | Connectivity |
| :--- | :--- | :--- |
| SKILL.md | Root Router | Calls this node for `review` |
| review.md | Quality Control | TRUST 5 & Security Audit |
| go.md | Implementation | Calls `review` in Phase 4 |

## Description
Perform code review and security audit using multi-provider consensus (Orchestra).

## Trust 5 Criteria
1. **Traceability**: Is the logic traceable to requirements?
2. **Robustness**: Are edge cases and errors handled?
3. **Understandability**: Is the code readable and clean?
4. **Security**: Does it follow security best practices?
5. **Testability**: Is the code easily testable?

## Review Process
1. **Parallel Audit**: Run `reviewer` and `security-auditor` agents concurrently.
2. **Consolidation**: Lead consolidates findings (Security > Quality).
3. **Verdict**:
    - **APPROVE**: Proceed to completion.
    - **REQUEST_CHANGES**: Spawn executor to fix issues.
4. **Multi-Provider**: [Conditional] Use `auto orchestra review` for consensus.

## Flags
| Flag | Description |
| :--- | :--- |
| `--multi` | Enable multi-provider review |
| `--strategy` | Consensus, Debate, Pipeline, Fastest |
