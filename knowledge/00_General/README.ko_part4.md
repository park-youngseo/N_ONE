---

## 📊 멀티 모델 오케스트레이션

| 전략 | 작동 방식 | 적합한 용도 |
|------|----------|------------|
| **🤝 Consensus** | 독립 응답을 키 합의로 병합 | 기획, 코드 리뷰 |
| **⚔️ Debate** | 2단계 토론 + 판사 판정 | 중요 결정, 보안 |
| **🔗 Pipeline** | N번째 출력 → N+1번째 입력 | 반복 정제 |
| **⚡ Fastest** | 가장 먼저 완료된 응답 사용 | 빠른 질문 |

프로바이더: **Claude** · **Codex** · **Gemini** · **OpenCode** — 그레이스풀 디그레이드 지원.

**인터랙티브 토론**을 실시간 창 시각화(cmux/tmux)로. **훅 기반 결과 수집**으로 구조화된 JSON 출력. Context7 문서 불가 시 **WebSearch 폴백**.

---

## 📖 전체 명령어

<details>
<summary><strong>CLI 명령어</strong> (루트 28개, 서브커맨드 포함 110개 이상)</summary>

| 명령어 | 설명 |
|--------|------|
| `auto init` | 하네스 초기화 — 플랫폼 감지, 파일 생성 |
| `auto update` | 하네스 업데이트 (마커 기반, 사용자 편집 보존) |
| `auto doctor` | 상태 진단 |
| `auto platform` | 플랫폼 관리 (list / add / remove) |
| `auto arch` | 아키텍처 분석 (generate / enforce) |
| `auto spec` | SPEC 관리 (new / validate / review) |
| `auto lore` | 의사결정 추적 (context / commit / validate / stale) |
| `auto orchestra` | 멀티 모델 오케스트레이션 (review / plan / secure / brainstorm / job-status / job-wait / job-result) |
| `auto setup` | 프로젝트 컨텍스트 문서 (generate / update / validate / status) |
| `auto status` | SPEC 대시보드 (done / in-progress / draft) |
| `auto telemetry` | 파이프라인 텔레메트리 (record / summary / cost / compare) |
| `auto skill` | 스킬 관리 (list / info / create) |
| `auto search` | 지식 검색 (Exa) |
| `auto docs` | 라이브러리 문서 조회 (Context7) |
| `auto lsp` | LSP 연동 (diagnostics / refs / rename / symbols / definition) |
| `auto verify` | 프론트엔드 UX 검증 (Playwright + VLM) |
| `auto check` | 하네스 규칙 검사 (안티패턴 스캔) |
| `auto hash` | 파일 해싱 (xxhash) |
| `auto issue` | 자동 이슈 리포터 (report / list / search) |
| `auto experiment` | 자율 실험 루프 (init / metric / record / commit / reset / summary / status) |
| `auto test` | E2E 시나리오 실행기 (run) |
| `auto react` | Reaction Engine (check / apply) |
| `auto agent` | 에이전트 관리 (create / run) |
| `auto terminal` | 터미널 멀티플렉서 관리 (detect / workspace / split / send / notify) |
| `auto pipeline` | 파이프라인 상태 관리 및 모니터링 |
| `auto permission` | 권한 모드 감지 (bypass / safe) |
| `auto browse` | 브라우저 자동화 (cmux browser / agent-browser) |
| `auto canary` | 배포 후 헬스 체크 (빌드 + E2E + 브라우저) |
| `auto connect` | 프로바이더 연결 마법사 (감지 → 설정 → 검증) |
| `auto update --self` | CLI 바이너리 자동 업데이트 (GitHub Releases + SHA256) |

</details>

<details>
<summary><strong>슬래시 명령어</strong> (AI 코딩 CLI 내부)</summary>

| 명령어 | 설명 |
|--------|------|
| `/auto plan "설명"` | 새 기능의 SPEC 작성 |
| `/auto go SPEC-ID` | 전체 파이프라인으로 구현 |
| `/auto go SPEC-ID --auto --loop` | 완전 자율 + 자가 치유 |
| `/auto go SPEC-ID --team` | Agent Teams (Lead/Builder/Guardian) |
| `/auto go SPEC-ID --multi` | 멀티 프로바이더 오케스트레이션 |
| `/auto fix "버그"` | 재현 우선 버그 수정 |
| `/auto review` | TRUST 5 코드 리뷰 |
| `/auto secure` | OWASP Top 10 보안 감사 |
| `/auto map` | 코드베이스 구조 분석 |
| `/auto sync SPEC-ID` | 구현 후 문서 동기화 |
| `/auto dev "설명"` | 풀 파워: plan(--multi --ultrathink) → go(--team --loop) → sync |
| `/auto setup` | 프로젝트 컨텍스트 문서 생성/업데이트 |
| `/auto stale` | 오래된 결정 및 패턴 감지 |
| `/auto why "질문"` | 의사결정 근거 조회 |
| `/auto experiment` | 자율 실험 루프 (메트릭 기반 반복) |
| `/auto test` | 프로젝트 E2E 시나리오 실행 |
| `/auto go SPEC-ID --continue` | 중단된 파이프라인 체크포인트에서 재개 |
| `/auto browse` | 브라우저 자동화 — 열기, 스냅샷, 클릭, 검증 |
| `/auto idea "설명"` | 멀티 프로바이더 브레인스토밍 + ICE 채점 |
| `/auto canary` | 배포 후 헬스 체크 (빌드 + E2E + 브라우저) |

</details>

---

## ⚙️ 설정

<details>
<summary><strong><code>autopus.yaml</code></strong> — 모든 것을 위한 단일 설정</summary>

```yaml
mode: full                    # full 또는 lite
project_name: my-project
platforms:
  - claude-code

architecture:
  auto_generate: true
  enforce: true

lore:
  enabled: true
  required_trailers: [Why, Decision]
  stale_threshold_days: 90

spec:
  review_gate:
    enabled: true
    strategy: debate
    providers: [claude, gemini]
    judge: claude

methodology:
  mode: tdd
  enforce: true

orchestra:
  enabled: true
  default_strategy: consensus
  providers:
    claude:
      binary: claude
    codex:
      binary: codex
    gemini:
      binary: gemini
    opencode:
      binary: opencode
```

</details>

---

## 🏗️ 아키텍처

```
autopus-adk/
├── cmd/auto/           # 진입점
├── internal/cli/       # Cobra 명령어 28개 (서브커맨드 포함 110개 이상)
├── pkg/
│   ├── adapter/        # 플랫폼 어댑터 4개 (Claude, Codex, Gemini, OpenCode)
│   ├── arch/           # 아키텍처 분석 + 규칙 강제
│   ├── browse/         # 브라우저 자동화 백엔드 (cmux/agent-browser 라우팅)
│   ├── config/         # 설정 스키마 + YAML 로딩
│   ├── constraint/     # 안티패턴 스캔
│   ├── content/        # 에이전트/스킬/훅/프로필 생성 + 스킬 활성화
│   ├── cost/           # 토큰 기반 비용 추정
│   ├── detect/         # 플랫폼/프레임워크/권한 감지
│   ├── e2e/            # E2E 시나리오 생성, 실행, 검증
│   ├── experiment/     # 자율 실험 루프 (메트릭 실행, 서킷 브레이커)
│   ├── issue/          # 자동 이슈 리포터 (컨텍스트 수집, 정제)
│   ├── lore/           # 의사결정 추적 (9-trailer 프로토콜)
│   ├── lsp/            # LSP 연동
│   ├── orchestra/      # 멀티 모델 오케스트레이션 (전략 4개 + brainstorm + 인터랙티브 토론 + 훅)
│   ├── pipeline/       # 파이프라인 상태 지속성 + 체크포인트 + 팀 모니터
│   ├── search/         # 지식 검색 (Context7/Exa) + 해시 기반 검색
│   ├── selfupdate/     # CLI 바이너리 자동 업데이트 (SHA256, 원자적 교체)
│   ├── setup/          # 프로젝트 문서 생성 + 검증
│   ├── sigmap/         # AST 기반 API 시그니처 추출 (Go + TypeScript)
│   ├── spec/           # EARS 요구사항 파싱/검증
│   ├── telemetry/      # 파이프라인 텔레메트리 (JSONL 이벤트 기록)
│   ├── template/       # Go 템플릿 렌더링
│   ├── terminal/       # 터미널 멀티플렉서 어댑터 (cmux, tmux, plain)
│   └── version/        # 빌드 메타데이터
├── templates/          # 플랫폼별 템플릿
├── content/            # 임베디드 콘텐츠 (16개 에이전트, 40개 스킬)
└── configs/            # 기본 설정
```

---

## 🔒 보안

### 🛡️ 패키지 공급망 공격 방어

> *"월 수천만 다운로드의 인기 Python 패키지에 악성 코드가 삽입됐습니다. 단순 `pip install`만으로 SSH 키, AWS 자격증명, DB 비밀번호를 탈취할 수 있었죠 — 직접 설치한 패키지가 아니라, 의존성 트리 깊숙한 곳에서."* — [Andrej Karpathy](https://x.com/karpathy)

AI 코딩 환경은 이 문제를 더 심각하게 만듭니다: 에이전트가 패키지를 자동 설치하고, 의존성 트리를 확장하고, 코드를 실행합니다 — 전부 사람의 검토 없이. **Autopus는 파이프라인 자체에 방어를 내장합니다.**

#### Autopus가 개발 워크플로우를 보호하는 방법

| 레이어 | 보호 내용 | 방법 |
|--------|----------|------|
| **파이프라인 게이트** | 모든 `/auto go`에서 의존성 취약점 스캔 | Security Auditor 에이전트가 Phase 4에서 `govulncheck ./...` 실행 |
| **시크릿 탐지** | 하드코딩된 자격증명을 커밋 전 탐지 | `gitleaks detect`로 변경 파일 전체 스캔 |
| **의존성 감사** | 의존성 트리의 알려진 CVE 탐지 | Go 프로젝트: `go list -m -json all \| nancy sleuth` |
| **락 파일 무결성** | 체크섬 검증된 의존성 | Go의 `go.sum`으로 재현 가능하고 변조 불가능한 빌드 보장 |
| **OWASP Top 10** | 인젝션, 인증 우회, SSRF 등 체계적 검사 | Security Auditor가 A01–A10 전체 커버 |
| **AI 에이전트 가드레일** | 에이전트가 패키지를 맹목적으로 설치 불가 | 하네스 규칙이 에이전트 행동 제약; 보안 게이트 FAIL 시 배포 차단 |

#### Go 이외 프로젝트

Autopus가 Python, Node.js 등 다른 생태계를 관리할 때도 동일한 원칙 적용:

```yaml
# autopus.yaml — 생태계별 보안 스캔 설정
security:
  scanners:
    go: "govulncheck ./..."
    python: "pip-audit && safety check"
    node: "npm audit --audit-level=high"
```

**하네스가 강제하는 모범 사례:**
- **버전 고정** — 모든 의존성을 정확한 버전으로 잠금 (`go.sum`, `package-lock.json`, `requirements.txt`)
- **의존성 최소화** — 300줄 파일 제한과 단일 책임 원칙이 자연스럽게 불필요한 임포트를 줄임
- **격리** — 병렬 executor가 격리된 git 워크트리에서 실행; 태스크 간 교차 오염 없음
- **맹목적 설치 금지** — Security Auditor 에이전트가 미확인/미검증 패키지를 코드베이스 진입 전 차단

### 바이너리 배포 안전성

모든 바이너리 릴리즈에는 **SHA256 체크섬** (`checksums.txt`)이 포함되며, 설치 시 자동으로 검증됩니다.

**권장: 설치 전 스크립트 확인**

```bash
# 1. 스크립트를 먼저 다운로드 — 실행 전에 내용 확인
curl -sSfL https://raw.githubusercontent.com/Insajin/autopus-adk/main/install.sh -o install.sh
less install.sh          # 무엇을 하는지 확인
sh install.sh            # 확인 후에만 실행
```

**또는 수동 검증:**

```bash
# 바이너리 + 체크섬 별도 다운로드
VERSION=$(curl -s https://api.github.com/repos/Insajin/autopus-adk/releases/latest | grep tag_name | sed 's/.*"v\(.*\)".*/\1/')
curl -LO "https://github.com/Insajin/autopus-adk/releases/download/v${VERSION}/autopus-adk_${VERSION}_$(uname -s | tr A-Z a-z)_$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/').tar.gz"
curl -LO "https://github.com/Insajin/autopus-adk/releases/download/v${VERSION}/checksums.txt"

# SHA256 검증
shasum -a 256 -c checksums.txt --ignore-missing
```

`auto update --self`도 바이너리 교체 전에 SHA256 체크섬을 검증합니다.

### 하지 않는 것들

- 텔레메트리나 분석 데이터 수집 없음
- 명시적 명령(`orchestra`, `search`, `update --self`) 외 네트워크 호출 없음
- AI 프로바이더 API 키 접근 없음 — Autopus는 API가 아닌 CLI 도구를 오케스트레이션

---

## 🤝 기여하기

Autopus-ADK는 MIT 라이선스의 오픈 소스입니다. PR 환영합니다!

```bash
make test       # 레이스 디텍션 테스트 실행
make lint       # go vet 실행
make coverage   # 커버리지 리포트 생성
```

---

<div align="center">

**🐙 Autopus** — 에이전트로. 에이전트에 의해. 에이전트를 위해.

</div>