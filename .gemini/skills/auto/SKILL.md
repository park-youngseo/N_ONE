---
name: auto
description: Autopus-ADK 메인 명령어 — 개발 워크플로우, 품질 검증, 탐색, 관리 서브커맨드 라우터
---

# /auto — Autopus-ADK (Registry Node)

/*
[Global Connectivity Table]
| Node | Responsibility | Subcommands |
| :--- | :--- | :--- |
| SKILL.md | Root Router & Triage | idea, fix, plan, go, review, etc. |
| nodes/setup.md | Context Generation | setup |
| nodes/status.md | SPEC Dashboard | status |
| nodes/plan.md | Planning & SPEC | plan |
| nodes/go.md | Implementation | go, dev |
| nodes/review.md | Quality & Security | review, secure |
| nodes/sync.md | Documentation Sync | sync |
| nodes/verify.md | UI Verification | verify, browse |
| nodes/manage.md | ADK Management | init, update, doctor |
*/

## Global Flags
| Flag | Description |
| :--- | :--- |
| `--think` | Activate Sequential Thinking analysis |
| `--auto` | Autonomous mode (skip confirmations) |
| `--loop` | RALF loop mode (self-healing) |
| `--multi` | Multi-provider mode (Orchestra) |
| `--quality` | Agent quality: `ultra` or `balanced` |

## Subcommand Routing (Node Mapping)

### 개발 워크플로우
- **idea**: @nodes/idea.md - Brainstorm and evaluate ideas
- **plan**: @nodes/plan.md - Write a SPEC/PRD
- **go**: @nodes/go.md - Implement a SPEC (TDD Pipeline)
- **dev**: @nodes/go.md - Full cycle: plan → go → sync
- **fix**: @nodes/go.md (or standalone) - Reproduction-First bug fix

### 품질 & 리뷰
- **review**: @nodes/review.md - Code review (TRUST 5)
- **secure**: @nodes/review.md - Security audit (OWASP)
- **verify**: @nodes/verify.md - Frontend UX verification
- **test**: @nodes/verify.md - Run E2E scenarios
- **canary**: @nodes/verify.md - Post-deploy health check

### 탐색 & 관리
- **setup**: @nodes/setup.md - Generate project context
- **status**: @nodes/status.md - SPEC dashboard
- **doctor**: @nodes/manage.md - Health diagnostics
- **sync**: @nodes/sync.md - Document synchronization

## Triage Process (Smart Routing)
WHEN natural language input is detected without a subcommand, assess difficulty:
1. **IDEA**: Exploratory question → Route to `idea`
2. **LOW**: Single file fix → Route to `fix`
3. **MEDIUM**: Feature extension → Route to `plan --skip-prd`
4. **HIGH**: Architecture change → Route to `idea --multi` (Decision Gate)

> 💡 모든 서브커맨드 상세 로직은 해당 `nodes/` 파일에 원자화되어 보관되어 있습니다.
