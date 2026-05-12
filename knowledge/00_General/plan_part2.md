| **Requirements** | — |

**File Ownership:**
- `.autopus/project/scenarios.md` — MODIFY

---

## 태스크 요약 테이블

| Task ID | Description | Agent | Mode | File Ownership | Complexity | Phase |
|---------|-------------|-------|------|----------------|------------|-------|
| T0-1 | 크로스 플랫폼 템플릿 헬퍼 | executor | sequential | `pkg/template/helpers.go` | MEDIUM | 0 |
| T0-2 | 공유 규칙/에이전트 소스 | executor | sequential | `content/rules/*.md`, `content/agents/*.md` | MEDIUM | 0 |
| T0-3 | 테스트 유틸리티 | executor | sequential | `pkg/adapter/testutil_test.go` | LOW | 0 |
| T1-1 | Codex 어댑터 파일 분리 | executor | parallel | `pkg/adapter/codex/*.go` | HIGH | 1 |
| T1-2 | Gemini 어댑터 파일 분리 | executor | parallel | `pkg/adapter/gemini/*.go` | HIGH | 1 |
| T2-1 | Codex 커스텀 커맨드 + 템플릿 | executor | parallel | `pkg/adapter/codex/codex_prompts.go`, `templates/codex/prompts/auto-*.md.tmpl` | MEDIUM | 2 |
| T2-2 | Codex 에이전트 + 훅 + 설정 | executor | parallel | `pkg/adapter/codex/codex_agents.go` 등 | HIGH | 2 |
| T2-3 | Gemini 커스텀 커맨드 + 규칙 | executor | parallel | `pkg/adapter/gemini/gemini_commands.go` 등 | MEDIUM | 2 |
| T2-4 | Gemini 에이전트 + 설정 | executor | parallel | `pkg/adapter/gemini/gemini_agents.go` 등 | HIGH | 2 |
| T2-5 | 서브에이전트 패턴 매핑 | executor | sequential | `templates/codex/skills/auto-go.md.tmpl` 등 | MEDIUM | 2 |
| T3-1 | 어댑터 단위 테스트 | executor | parallel | `pkg/adapter/codex/*_test.go`, `pkg/adapter/gemini/*_test.go` | MEDIUM | 3 |
| T3-2 | E2E 통합 테스트 | executor | sequential | `pkg/adapter/integration_test.go` | MEDIUM | 3 |
| T4-1 | 300줄 제한 + 매니페스트 | executor | sequential | `.autopus/claude-code-manifest.json` | LOW | 4 |
| T4-2 | E2E 시나리오 문서 | executor | sequential | `.autopus/project/scenarios.md` | LOW | 4 |

---

## 의존성 그래프

```
T0-1, T0-2, T0-3  (Phase 0 — sequential)
       │
       ▼
T1-1 ──┬── T1-2  (Phase 1 — parallel)
       │
       ▼
T2-1 ──┬── T2-2 ──┬── T2-3 ──┬── T2-4  (Phase 2 — parallel)
       │           │           │
       └───────────┴───────────┘
                   │
                   ▼
                 T2-5  (Phase 2 — sequential tail)
                   │
                   ▼
         T3-1 (parallel: codex ∥ gemini tests)
                   │
                   ▼
                 T3-2  (Phase 3 — E2E)
                   │
                   ▼
         T4-1 ──── T4-2  (Phase 4 — sequential)
```

---

## 예상 워크로드

| Phase | 태스크 수 | 예상 시간 | 병렬화 |
|-------|----------|----------|--------|
| Phase 0 | 3 | ~30분 | 순차 |
| Phase 1 | 2 | ~40분 | 병렬 (x2) |
| Phase 2 | 5 | ~60분 | 병렬 (x4) + 순차 (x1) |
| Phase 3 | 2 | ~40분 | 병렬 + 순차 |
| Phase 4 | 2 | ~15분 | 순차 |
| **Total** | **14** | **~3시간** | — |