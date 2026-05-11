---
name: auto
description: Autopus-ADK 메인 명령어 — 개발 워크플로우, 품질 검증, 탐색, 관리 서브커맨드 라우터
---

# /auto — Autopus-ADK [Registry Node]

> [!IMPORTANT]
> **AX-300 Protocol Compliance**: This file acts as a high-density Registry. All detailed logic is delegated to atomic nodes.

## Global Connectivity Table

| Domain | Node Path | Status |
|--------|-----------|--------|
| **Setup** | [nodes/setup.md](nodes/setup.md) | 원자화 완료 |
| **Idea** | [nodes/idea.md](nodes/idea.md) | 원자화 완료 |
| **Plan** | [nodes/plan.md](nodes/plan.md) | 원자화 완료 |
| **Go** | [nodes/go.md](nodes/go.md) | 원자화 완료 |
| **Verify** | [nodes/verify.md](nodes/verify.md) | 원자화 완료 |
| **Manage** | [nodes/manage.md](nodes/manage.md) | 원자화 완료 |

## Core Banner

```
🐙 Autopus ─────────────────────────
```

## Global Flags

| Flag | Description |
|------|-------------|
| `--think` | Activate Sequential Thinking MCP |
| `--auto` | Autonomous mode: skip confirmations |
| `--multi` | Multi-provider orchestra mode |
| `--quality` | Agent quality: ultra | balanced |

## Subcommand Routing

1. **Setup**: `/auto setup` → See [setup.md](nodes/setup.md)
2. **Idea**: `/auto idea` → See [idea.md](nodes/idea.md)
3. **Plan**: `/auto plan` → See [plan.md](nodes/plan.md)
4. **Go**: `/auto go` → See [go.md](nodes/go.md)
5. **Verify**: `/auto review`, `/auto verify`, `/auto sync` → See [verify.md](nodes/verify.md)
6. **Manage**: `/auto doctor`, `/auto update`, `/auto platform` → See [manage.md](nodes/manage.md)

---
🐙
