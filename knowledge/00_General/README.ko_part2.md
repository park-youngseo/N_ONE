
내장 **서킷 브레이커**로 무한 반복을 방지합니다. **단순성 점수**가 과도하게 복잡한 솔루션에 패널티를 부여합니다. 각 반복은 git 커밋으로 기록되어 리뷰 및 롤백이 용이합니다.

> **상태: 실험적** — CLI 명령어(`auto experiment`)는 사용 가능하지만 스킬 레벨 통합은 진행 중입니다. 핵심 반복 루프는 작동하며, 전체 파이프라인 통합은 준비 중입니다.

### 🧠 실패에서 배우는 파이프라인

Autopus 파이프라인은 단순히 실패하지 않습니다 — **왜 실패했는지 기억하고** 다음에 같은 실수를 방지합니다.

```
Gate 2 FAIL: golangci-lint — pkg/auth/에서 미사용 변수
→ .autopus/learnings/pipeline.jsonl에 자동 기록
→ 다음 /auto go: 학습 내용이 executor 프롬프트에 주입
→ 같은 실수 반복 없음
```

모든 파이프라인 실패가 구조화된 학습 항목으로 기록됩니다. 다음 실행 시 관련 학습이 자동으로 에이전트 프롬프트에 주입되어 — 세션을 넘어선 **조직 기억**을 파이프라인에 부여합니다.

### 🏥 배포 후 헬스 체크

먼저 배포하고, 즉시 검증합니다. `canary`는 빌드 검증, E2E 테스트, 브라우저 헬스 체크를 라이브 배포에 대해 실행합니다.

```bash
/auto canary                          # 빌드 + E2E + 브라우저 자동 검증
/auto canary --url https://myapp.com  # 특정 배포 URL 대상
/auto canary --watch 5m               # 5분마다 반복
/auto canary --compare                # 이전 카나리 리포트와 비교
```

빌드 상태, 테스트 결과, 접근성 점수, 스크린샷 diff가 포함된 `canary.md` 진단 보고서를 생성합니다.

### 🔀 스마트 모델 라우팅

모든 태스크에 Opus가 필요하지 않습니다. Autopus가 메시지 복잡도를 분석하고 적절한 모델로 자동 라우팅합니다.

```
단순 질의     → Haiku  (빠르고 저렴)
코드 리뷰     → Sonnet (균형)
아키텍처      → Opus   (깊은 추론)
```

설정 불필요 — 라우터가 토큰 수, 코드 복잡도, 도메인 신호를 평가하여 최적 모델을 선택합니다. `--quality ultra`로 언제든 오버라이드 가능.

### 🔌 프로바이더 연결 마법사

AI 프로바이더 설정에 문서를 읽을 필요 없습니다. `auto connect`가 3단계 가이드 설정을 안내합니다.

```bash
auto connect    # 대화형 마법사: 감지 → 설정 → 검증
```

설치된 CLI 도구를 감지하고, API 키를 검증하고, 연결을 테스트하고, 프로바이더 설정을 작성합니다 — 모두 한 명령으로.

### 🤖 ADK Worker — 로컬 에이전트 실행

Autopus CLI와 Autopus 플랫폼 간의 브릿지입니다. ADK Worker는 A2A + MCP 하이브리드 태스크를 로컬에서 실행하여, 클라우드 의존 없이 플랫폼급 오케스트레이션을 가능하게 합니다.

어떤 기능인가요?
- 로컬 워크스페이스를 Autopus 플랫폼의 worker loop에 연결합니다
- 플랫폼에서 내려온 태스크를 로컬 도구로 실행합니다
- 메인 하네스와 같은 보안, 예산, 감사 레일을 재사용합니다

지금은 어떻게 보면 되나요?
- `auto init`, Codex `@auto ...`, OpenCode `/auto ...`를 쓰려는 용도라면 Worker는 일단 무시해도 됩니다
- `auto worker ...`는 선택형 고급 기능이며, 현재는 점진적으로 정리 및 문서화 중입니다

### 💰 반복 예산 관리

Worker가 영원히 실행되지 않습니다. 각 executor에 도구 호출 예산이 할당되어 — 복잡한 태스크를 완료할 충분한 여유를 보장하면서 폭주 에이전트를 방지합니다.

### 📦 컨텍스트 압축

파이프라인이 단계를 진행하면서 이전 컨텍스트가 자동으로 압축됩니다 — 중요한 정보를 잃지 않으면서 에이전트 프롬프트를 집중적이고 토큰 한도 내로 유지합니다.

### 🔄 죽지 않는 파이프라인

파이프라인 중간에 크래시? 마지막 체크포인트에서 정확히 재개합니다.

```bash
/auto go SPEC-AUTH-001 --continue    # 마지막 체크포인트에서 재개
```

YAML 기반 체크포인트가 매 단계 후 파이프라인 상태를 저장합니다. Stale 감지가 오래된 세션 재개를 방지합니다. `--auto --loop`과 결합하면 **완전 복원력 있는 자율 파이프라인**을 얻습니다.

### 🧪 코드에서 E2E 시나리오 자동 생성

E2E 테스트 시나리오를 자동 생성하고 실행합니다 — 수동 테스트 작성 불필요.

```bash
auto test run                    # 모든 시나리오 실행
auto test run -s init --verbose  # 특정 시나리오 실행
```

Autopus가 코드베이스(Cobra 커맨드, API 라우트, 프론트엔드 페이지)를 분석하여 **검증 프리미티브** (`exit_code`, `stdout_contains`, `status_code`, `json_path` 등)가 포함된 타입 시나리오를 생성합니다. 증분 동기화로 코드 변경에 따라 시나리오가 자동 갱신됩니다.

### 🔧 더 많은 도구

| 기능 | 명령어 | 설명 |
|------|--------|------|
| **Reaction Engine** | `auto react check/apply` | CI 실패 감지, 로그 분석, 수정 보고서 자동 생성 |
| **Meta-Agent Builder** | `auto agent create` / `auto skill create` | 패턴 기반 커스텀 에이전트/스킬 스캐폴드 생성 |
| **Hard Gate** | `auto check --gate` | 파이프라인 게이트 강제 (mandatory/advisory 모드) |
| **Self-Update** | `auto update --self` | 원자적 바이너리 업데이트 — GitHub Releases 확인 + SHA256 검증 |
| **비용 추적** | `auto telemetry cost` | 모델별 토큰 기반 파이프라인 비용 추정 |
| **이슈 리포터** | `auto issue report` | 에러 컨텍스트 자동 수집, 시크릿 정제, GitHub 이슈 생성 |
| **시그니처 맵** | `auto setup` | AST 분석으로 exported API 시그니처 추출 (Go + TypeScript) |
| **테스트 러너 감지** | `auto init` | jest, vitest, pytest, cargo 테스트 프레임워크 자동 감지 |

### 🌐 하나의 설정, 네 개 플랫폼

```bash
auto init   # 지원되는 설치된 AI 코딩 CLI 자동 감지
```

하나의 `autopus.yaml`이 감지된 지원 플랫폼에 **네이티브 설정**을 생성합니다.

| 플랫폼 | 생성되는 파일 |
|--------|-------------|
| **Claude Code** | `.claude/rules/`, `.claude/skills/`, `.claude/agents/`, `CLAUDE.md` |
| **Codex** | `.codex/`, `.agents/skills/`, `.agents/plugins/marketplace.json`, `.autopus/plugins/auto/`, `AGENTS.md` |
| **Gemini CLI** | `.gemini/`, `GEMINI.md` |
| **OpenCode** | `.opencode/rules/`, `.opencode/agents/`, `.opencode/commands/`, `.opencode/plugins/`, `.agents/skills/`, `AGENTS.md`, `opencode.json` |
동일한 16개 에이전트. 동일한 40개 스킬. 동일한 규칙. **모든 플랫폼.**

Codex 참고:
- `auto init` 또는 `auto update` 직후에는 `$auto plan ...`, `$auto go ...`, `$auto idea ...`를 바로 사용할 수 있습니다
- `.agents/plugins/marketplace.json`에 등록된 로컬 플러그인(`.autopus/plugins/auto`)을 설치하면 더 자연스러운 `@auto ...` 문법을 사용할 수 있습니다

OpenCode 참고:
- `/auto ...`와 `/auto-plan ...` 같은 직접 alias가 `.opencode/commands/`에 생성됩니다
- 네이티브 규칙/에이전트/플러그인은 `.opencode/` 아래에, 재사용 스킬은 `.agents/skills/` 아래에 생성됩니다
- `/auto status`, `/auto map`, `/auto why`, `/auto verify`, `/auto secure`, `/auto test`, `/auto dev`, `/auto doctor` 같은 helper workflow도 OpenCode 명령 래퍼로 함께 생성됩니다
- `opencode.json`이 관리형 hook plugin을 자동 등록하므로 `auto init` 또는 `auto update` 직후 `.opencode/plugins/autopus-hooks.js`가 바로 활성화됩니다

### Codex vs OpenCode 비교

| 항목 | Codex | OpenCode |
|------|-------|----------|
| 기본 명령 문법 | `@auto <subcommand> ...` | `/auto <subcommand> ...` |
| `auto init` 직후 바로 되는 것 | `$auto ...` repo-skill fallback | `/auto ...`, `/auto-<subcommand> ...` 래퍼 |
| 추가 설치 단계 | 있음. `.agents/plugins/marketplace.json`의 로컬 플러그인을 설치해야 `@auto ...` 사용 가능 | 라우터 기준 추가 설치 불필요. `opencode.json`이 관리형 plugin을 자동 연결 |
| 생성되는 표면 | `.codex/`, `.agents/skills/`, `.agents/plugins/marketplace.json`, `.autopus/plugins/auto/`, `AGENTS.md` | `.opencode/commands/`, `.opencode/agents/`, `.opencode/rules/`, `.opencode/plugins/`, `.agents/skills/`, `AGENTS.md`, `opencode.json` |
| 지금 잘 되는 것 | 핵심 `auto` 워크플로우, repo skill, 로컬 플러그인 기반 `@auto` 라우팅 | 핵심 `auto` 워크플로우, 네이티브 명령 래퍼, 관리형 hook plugin wiring |
| 현재 경계 | `@auto ...`는 로컬 플러그인 설치가 전제이며, 설치 전에는 `$auto ...`를 사용 | 현재 패리티 목표는 핵심 워크플로우 표면입니다. Claude식 native settings/statusline 폭까지 주장하지 않습니다 |
| Worker 표면 | 지금은 선택 사항입니다. 플랫폼 연결 worker 실행이 필요하지 않다면 무시해도 됩니다 | 지금은 선택 사항입니다. 플랫폼 연결 worker 실행이 필요하지 않다면 무시해도 됩니다 |

---

## 🚀 빠른 시작 가이드

5분 안에 첫 AI 기반 기능을 만들어 보세요.

### 1단계 · 설치 (한 줄)

> **AI 코딩 도구(Claude Code, Codex, OpenCode 등)의 채팅창에 아래 명령을 붙여넣으세요** — 에이전트가 대신 실행해서 설치까지 처리합니다. 터미널에서 직접 실행해도 됩니다.

```bash
# macOS / Linux — 바이너리 설치 + 필수 도구 점검
cd your-project    # 프로젝트 폴더로 이동 (예: cd ~/my-app)
curl -sSfL https://raw.githubusercontent.com/Insajin/autopus-adk/main/install.sh | sh

# Windows (CMD or PowerShell)
cd your-project
powershell -c "irm https://raw.githubusercontent.com/Insajin/autopus-adk/main/install.ps1 | iex"
```

설치 스크립트는 `auto` CLI와 `autopus` alias를 함께 설치하고, 필수 도구를 점검한 뒤 이미 있는 것은 건너뛰고 누락된 `git`, GitHub CLI 같은 핵심 도구를 자동 설치합니다. 대신 `auto init`은 자동으로 실행하지 않고 사용자에게 맡깁니다.

<details>
<summary>기타 설치 방법</summary>

```bash
# go install (Go 1.26+ 필요)
go install github.com/Insajin/autopus-adk/cmd/auto@latest

# 소스에서 빌드
git clone https://github.com/Insajin/autopus-adk.git
cd autopus-adk && make build && make install
```

</details>

설치가 끝나면 스크립트가 아래 세 명령을 설명해줍니다.

- `auto init`: 현재 프로젝트를 초기화하고 `autopus.yaml`과 플랫폼 파일을 생성
- `auto update --self`: `auto` CLI 바이너리 자체를 최신 버전으로 업데이트
- `auto update`: 현재 프로젝트의 규칙, 스킬, 에이전트, 생성 파일을 최신 템플릿으로 갱신

### 2단계 · 프로젝트 초기화

```bash
cd your-project
auto init
```

Windows에서 `powershell -c ...`를 Git Bash 안에서 실행한 경우에는, 부모 Bash 프로세스의 `PATH`가 즉시 갱신되지 않으므로 설치 후 Git Bash를 다시 열어야 합니다. 이 경우 설치 스크립트가 실제 설치 경로와 `export PATH=...` 안내를 함께 출력합니다.

`auto init`은 설치된 AI 코딩 CLI(Claude Code, Codex, Gemini CLI, OpenCode)를 자동 감지하고, 각 플랫폼에 맞는 **네이티브 설정** — 규칙, 스킬, 에이전트 — 을 하나의 `autopus.yaml`에서 생성합니다.

플랫폼별 명령 문법:
- Codex: 생성된 로컬 플러그인을 설치한 뒤 `@auto ...`를 사용하고, 그 전에는 `$auto ...`를 사용합니다
- OpenCode: `/auto ...` 또는 `/auto-<subcommand> ...`를 사용합니다
- Claude Code / Gemini CLI: `/auto ...`를 사용합니다

```
✓ 감지됨: claude-code, codex, gemini-cli, opencode
✓ 생성됨: .claude/rules/, .claude/skills/, .claude/agents/, CLAUDE.md
✓ 생성됨: .codex/, AGENTS.md
✓ 생성됨: .gemini/, GEMINI.md
✓ 생성됨: .opencode/, .agents/skills/, AGENTS.md, opencode.json
✓ 생성됨: autopus.yaml
```

### 3단계 · 프로젝트 컨텍스트 생성 (`/auto setup`)

**가장 중요한 단계입니다.** AI 에이전트는 세션 간 모든 기억을 잃습니다 — 매번 프로젝트를 처음 보는 것과 같습니다. `/auto setup`은 에이전트가 프로젝트를 즉시 이해할 수 있게 해주는 "온보딩 문서"를 생성합니다.

```bash
/auto setup     # Claude Code, Gemini CLI, OpenCode
@auto setup     # Codex 로컬 플러그인 설치 후
$auto setup     # Codex 플러그인 설치 전 fallback
```

코드베이스를 분석하여 5개의 컨텍스트 문서를 생성합니다:

```
ARCHITECTURE.md                    # 도메인, 레이어, 의존성 맵
.autopus/project/product.md       # 프로젝트 설명, 핵심 기능
.autopus/project/structure.md     # 디렉토리 구조, 패키지 역할, 엔트리포인트
.autopus/project/tech.md          # 기술 스택, 빌드, 테스트 전략
.autopus/project/scenarios.md     # 코드에서 추출된 E2E 테스트 시나리오
```

> 💡 **왜 중요한가요?** 이 문서 없이 AI가 프로젝트를 보는 것은, 온보딩 없이 첫 출근한 신입사원과 같습니다 — 아키텍처를 추측하고, 컨벤션을 놓치고, 이미 존재하는 패턴을 다시 만들게 됩니다. `/auto setup`으로 모든 에이전트 세션이 정보를 갖고 시작합니다.

### 4단계 · 첫 기능 만들기

준비 완료. 원하는 것을 자연어로 설명하세요:

```bash
# 1. 기획 — AI가 전체 SPEC 생성 (요구사항, 태스크, 수락 기준)
/auto plan "GET /healthz 헬스 체크 엔드포인트 추가"

# 2. 구현 — 16개 에이전트가 구현, 테스트, 리뷰 처리
/auto go SPEC-HEALTH-001 --auto

# 3. 배포 — 문서 동기화, SPEC 상태 업데이트, 의사결정 이력과 함께 커밋
/auto sync SPEC-HEALTH-001
```

```
╭────────────────────────────────────╮
│ 🐙 파이프라인 완료!                 │
│ SPEC-HEALTH-001: 헬스 체크          │
│ 태스크: 3/3 │ 커버리지: 92%         │
│ 리뷰: APPROVE                      │
╰────────────────────────────────────╯
```

이게 전부입니다 — 테스트, 보안 감사, 완전한 문서화가 포함된 프로덕션 수준 코드가 세 개의 명령으로 완성됩니다.

### 빠른 참조

| 하고 싶은 것 | 명령어 |
|-------------|--------|
| **아이디어 브레인스토밍** | `/auto idea "설명" --multi --ultrathink` |
| **풀 사이클 (추천)** | `/auto dev "설명"` |
| 새 기능 기획 | `/auto plan "설명"` |
| SPEC 구현 | `/auto go SPEC-ID --auto --loop --team` |
| 버그 수정 (SPEC 불필요) | `/auto fix "설명"` |
| 자연어로 설명만 | `/auto 로그인 페이지에 2FA 추가` |
| 배포 후 헬스 체크 | `/auto canary` |
| 코드 리뷰 | `/auto review` |
| 보안 감사 | `/auto secure` |
| 중단된 파이프라인 재개 | `/auto go SPEC-ID --continue` |
| 변경 후 문서 동기화 | `/auto sync SPEC-ID` |

### 업데이트

Autopus-ADK는 두 가지 업데이트가 있습니다:

**1. 바이너리 업데이트** — `auto` CLI 자체를 최신 버전으로:

```bash
auto update --self
```

GitHub Releases에서 최신 버전을 확인하고, SHA256 체크섬 검증 후 바이너리를 교체합니다. 현재 버전은 `auto version`으로 확인하세요.

**2. 하네스 업데이트** — 프로젝트의 규칙/스킬/에이전트를 최신 템플릿으로:

```bash
auto update
```

`.claude/*`, `.codex/*`, `.gemini/*`, `.opencode/*`, `.agents/skills/*` 등의 하네스 파일을 갱신합니다. `AUTOPUS:BEGIN`~`AUTOPUS:END` 마커 바깥의 사용자 편집은 보존됩니다. 새로 설치된 플랫폼이 있으면 자동 감지하여 해당 파일도 생성합니다.

**한 줄로 둘 다:**

```bash