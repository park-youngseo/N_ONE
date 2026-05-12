# Changelog — autopus-adk

All notable changes to this project will be documented in this file.

## [Unreleased]

## [v0.40.37] — 2026-04-19

### Changed

- **Residual golangci-lint cleanup sweep across ADK** (2026-04-19): 남아 있던 `staticcheck`/`ineffassign`/test-style 경고를 일괄 정리해 현재 `golangci-lint run --max-issues-per-linter=0 --max-same-issues=0` 기준 0 issue 상태로 수렴
  - `.golangci.yml`, `internal/cli/**`, `pkg/orchestra/**`, `pkg/setup/**`, `pkg/worker/**` — 빈 에러 브랜치, 비효율 할당, 루프/append 패턴, 테스트 fixture/헬퍼 표현을 정리
  - `pkg/adapter/opencode/opencode_router_contract.go`, `pkg/content/agent_transformer_condense.go`, `internal/cli/issue_auto.go` — 더 이상 쓰이지 않는 보조 경로와 dead code를 제거
  - 광범위한 테스트/헬퍼 파일에서 lint 친화적 표현으로 정렬해 release gate를 통과하도록 회귀 범위를 동기화

## [v0.40.36] — 2026-04-19

### Fixed

- **Install bootstrap now separates install from init** (2026-04-19): installer가 `auto init`/`auto update`를 자동 실행하지 않고, 필수 도구만 점검한 뒤 `auto init`, `auto update --self`, `auto update`의 역할을 명시적으로 안내하도록 정리
  - `install.sh`, `install.ps1` — post-install 단계에서 required dependency만 자동 설치하고, 자동 project init/update 분기 제거
  - `internal/cli/doctor.go`, `internal/cli/doctor_fix.go` — `--required-only` 플래그와 required dependency filter 추가
  - `pkg/detect/detect.go` — `gh`를 필수 도구로 승격하고 Gemini CLI npm 패키지를 `@google/gemini-cli`로 정정
  - `README.md`, `docs/README.ko.md`, `internal/cli/doctor_fix_runtime_test.go`, `internal/cli/doctor_fix_test.go`, `pkg/detect/fullmode_deps_test.go` — 설치 가이드/회귀 테스트 동기화 및 테스트 파일 분할로 300-line limit 유지

- **E2E scenario runner backend submodule path correction** (2026-04-19): Backend build 시나리오가 `Autopus/`를 cwd로 잡아 존재하지 않는 `cmd/server` 경로를 참조하던 문제를 실제 backend 소스 경로인 `Autopus/backend/`로 정렬
  - `pkg/e2e/build.go`, `pkg/e2e/build_test.go` — default submodule map을 canary H2/H3 build cwd와 일치시키고 회귀 테스트 추가

- **Permission detection tests now use injected process-tree stubs** (2026-04-19): `--dangerously-skip-permissions`가 걸린 세션에서 `pkg/detect` 테스트가 실제 부모 프로세스 트리에 오염되던 문제를 제거
  - `pkg/detect/permission.go`, `pkg/detect/permission_test.go` — `checkProcessTreeFn` 주입 지점과 결정적 stub helper 추가

- **CC21 monitor runtime flake removed via Claude version injection hook** (2026-04-19): `claude --version` subprocess timeout으로 인해 `TestResolveCC21MonitorRuntime_Enabled`가 간헐적으로 실패하던 문제를 테스트 전용 version injector로 제거
  - `pkg/platform/claude.go`, `internal/cli/orchestra_cc21_test.go` — `claudeVersionFn`/`SetClaudeVersionForTest` 추가 및 monitor runtime 회귀 테스트 보강

## [v0.40.35] — 2026-04-19

### Fixed

- **Release workflow bootstrap ordering** (2026-04-19): `goreleaser-action@v7`가 `cosign`이 PATH에 있을 때 GoReleaser 다운로드 자체의 sigstore bundle을 추가 검증하는데, upstream bundle 검증 실패로 `v0.40.34` release workflow가 즉시 중단되던 문제를 우회
  - `.github/workflows/release.yaml` — action을 `install-only`로 먼저 실행해 checksum 검증만 수행하고, 이후 `cosign` 설치와 `goreleaser release --clean` 직접 실행으로 실제 checksum signing 단계만 유지하도록 순서 조정

## [v0.40.34] — 2026-04-19

### Added

- **Test Profile 기반 시나리오 요구조건 스킵** (2026-04-19): `auto test run`에 `--profile` capability 집합을 도입해 시나리오의 `Requires` 조건이 충족되지 않으면 FAIL 대신 SKIP으로 처리
  - `internal/cli/test.go`, `internal/cli/test_profile_test.go` — `--profile` 플래그, SKIP 집계, JSON 출력 회귀 테스트 추가
  - `pkg/config/test_profiles.go`, `pkg/config/test_profiles_test.go`, `pkg/config/schema.go` — profile별 capability 기본값 및 `autopus.yaml` 확장
  - `pkg/e2e/requires.go`, `pkg/e2e/scenario.go`, `pkg/e2e/scenario_requires_test.go` — `Requires` 파싱 및 capability mismatch 계산 로직 추가
  - `templates/shared/scenarios-*.md.tmpl` — 시나리오 템플릿에 `Requires` 필드 추가

### Fixed

- **SPEC review finding status breakdown summary** (2026-04-19): `auto spec review` 최종 요약이 단순 unique count 대신 `open/resolved/out_of_scope` 상태별 집계를 함께 출력하도록 개선해 운영자가 `review-findings.json`을 별도로 집계하지 않아도 열린 finding 수를 바로 확인 가능
  - `pkg/spec/findings_summary.go`, `pkg/spec/findings_test.go` — `ReviewFinding` slice를 상태별로 집계하는 `SummarizeFindings` / `FindingsSummary.Format` 로직과 회귀 테스트 추가
  - `internal/cli/spec_review.go` — 최종 CLI 요약을 status breakdown 표면으로 교체

- **Pipeline worktree remove canonical path fallback** (2026-04-19): macOS의 `/tmp` → `/private/tmp`, `/var` → `/private/var` symlink 환경에서 `git worktree remove`가 symlink path를 실제 worktree로 인식하지 못해 release gate의 `pkg/pipeline` 테스트가 실패하던 문제를 수정
  - `pkg/pipeline/worktree.go` — remove 시 원본 path와 canonical path를 순차 재시도하고, 실제 git worktree가 아닌 fallback 디렉터리는 안전하게 `os.RemoveAll`로 정리하도록 보강
  - `pkg/pipeline/worktree_internal_test.go` — symlink alias로 생성한 실제 worktree를 remove 하는 회귀 테스트 추가

- **SPEC 리뷰 체크리스트 런타임 주입 및 self-verify 기록 경로 복구 (SPEC-SPECWR-002)** (2026-04-19): `auto spec review`가 `content/rules/spec-quality.md`를 실제 런타임 프롬프트에 주입하고, `CHECKLIST:` 응답을 구조화 파싱하며, `auto spec self-verify`로 결정적 JSONL 기록을 남길 수 있도록 동기화.
  - `pkg/spec/checklist.go`, `pkg/spec/prompt.go` — embed 우선 + 디스크 fallback 체크리스트 로더, `## Quality Checklist` 주입, checklist response examples 추가
  - `pkg/spec/types.go`, `pkg/spec/reviewer.go`, `internal/cli/spec_review_loop.go`, `internal/cli/spec_review.go` — `ChecklistOutcome` 타입, `CHECKLIST:` 파싱, provider outcome 집계, 최종 요약 출력 연결
  - `pkg/spec/selfverify.go`, `internal/cli/spec.go`, `internal/cli/spec_self_verify.go`, `.gitignore` — `auto spec self-verify` 서브커맨드, 100라인 retention, `.self-verify.log` ignore 규칙 추가
  - `pkg/spec/checklist_test.go`, `pkg/spec/reviewer_checklist_test.go`, `pkg/spec/selfverify_test.go`, `internal/cli/spec_review_checklist_test.go`, `internal/cli/spec_self_verify_test.go` — checklist injection/parser/CLI/self-verify 회귀 테스트 추가

- **SPEC 리뷰 수렴성 재구축 (SPEC-REVFIX-001)** (2026-04-19): `auto spec review --multi`가 대부분의 SPEC에서 PASS에 도달하지 못하고 REVISE 루프를 소진한 뒤 circuit breaker로 종료되던 7개 복합 결함 제거.
  - **REQ-01 Supermajority verdict**: `MergeVerdicts`가 `spec.review_gate.verdict_threshold`(기본 0.67) 기준 supermajority를 적용. 1 REJECT 단독 override는 유지(security gate). `pkg/spec/reviewer.go`
  - **REQ-02 Revision 루프 내 재로드**: `runSpecReview`가 iteration마다 `spec.Load(specDir)` 재호출. 외부 수정이 다음 round에 반영됨. `internal/cli/spec_review_loop.go`
  - **REQ-03 다중 문서 주입**: `BuildReviewPrompt`가 plan.md / research.md / acceptance.md 본문을 별도 섹션으로 주입. `doc_context_max_lines`(기본 200)로 trim. `pkg/spec/prompt.go`
  - **REQ-04 Verdict 판정 기준 명문화**: 프롬프트에 `critical==0 && security==0 && major<=2 → PASS` 규칙 포함. `pass_criteria` override 지원.
  - **REQ-05 FINDING 포맷 강제 + empty RawContent guard**: structured FINDING few-shot(positive 2 + negative 1), `doc.RawContent == ""` 시 early error.
  - **REQ-06 DeduplicateFindings / MergeSupermajority 프로덕션 통합**: REVCONV-001이 구현했으나 호출되지 않던 dead code를 `runSpecReview` 경로에 연결. critical/security는 supermajority 우회.
  - **REQ-07 Finding ID 전역 유니크**: `parseDiscoverFindings`가 ID 비어있게 두고 `DeduplicateFindings`가 global `F-001..` 재발급. `ApplyScopeLock` 오동작 해결.
  - 신규: `pkg/spec/merge.go`, `pkg/config/schema_spec.go`, `internal/cli/spec_review_loop.go`, `pkg/spec/prompt_test.go`, `pkg/spec/reviewer_supermajority_test.go`, `internal/cli/spec_review_scaffold_test.go`
  - `autopus.yaml` 샘플에 `verdict_threshold`, `pass_criteria`, `doc_context_max_lines` 주석 예시 추가

### Changed

- **Claude Code 2.1 CC21 경로 연결 및 precedence 정렬 (SPEC-CC21-001)** (2026-04-19): effort frontmatter, TaskCreated hook, initial prompt 검사, monitor 기반 완료 감지를 source-of-truth와 CLI/runtime 경로에 연결
  - `internal/cli/effort*.go`, `internal/cli/check_initial_prompt*.go`, `internal/cli/orchestra_cc21.go`, `internal/cli/check_cc21.go`, `internal/cli/cc21_runtime.go` — CC21 전역 플래그, runtime precedence, check 명령, orchestra wiring 추가
  - `pkg/orchestra/cc21_monitor.go`, `pkg/platform/claude.go`, `pkg/platform/claude_test*.go` — Claude Code 2.1 capability 감지와 monitor contract 연결
  - `content/hooks/task-created-validate.sh`, `content/hooks/README.md`, `pkg/content/hooks.go`, `pkg/adapter/claude/claude_task_created_test.go` — TaskCreated generated default와 runtime override precedence 정렬
  - `content/skills/monitor-patterns.md`, `content/embed.go`, `content/skills/adaptive-quality.md`, `content/skills/idea.md`, `content/skills/agent-pipeline.md` — CC21 monitor/effort 규칙과 문서 표면 동기화
  - `pkg/adapter/claude/claude_generate.go`, `pkg/adapter/claude/claude_prepare_files.go`, `pkg/adapter/claude/claude_update.go` — Claude adapter 파일 생성/업데이트 경로를 300줄 제한에 맞게 분리 정리

- **Claude deferred-tools 선로딩 규칙 추가** (2026-04-18): Claude Code의 지연 로드 도구(`AskUserQuestion`, `TaskCreate`, `TeamCreate` 등)가 스키마 미로드 상태로 호출될 때 생기던 평문 downgrade / validation error를 줄이기 위해 전역 규칙을 추가
  - `content/rules/deferred-tools.md` — `/auto triage`, Gate 1 승인, `--team` 진입 시 `ToolSearch`로 스키마를 먼저 로드하도록 trigger point 규칙 추가

- **Claude Code Agent Teams + mode 파라미터 동기화** (2026-04-18): Agent Teams 공식 스펙(https://code.claude.com/docs/en/agent-teams)을 반영하고, Agent() 호출 파라미터 이름을 `permissionMode` → `mode` 로 통일. 플랫폼별 `--team` 플래그 동작 명시.
  - `content/skills/agent-pipeline.md`, `content/skills/worktree-isolation.md` — 본문 `Agent(... permissionMode=)` 10건 → `mode=`
  - `templates/codex/skills/agent-pipeline.md.tmpl`, `templates/codex/skills/worktree-isolation.md.tmpl`, `templates/gemini/skills/agent-pipeline/SKILL.md.tmpl`, `templates/gemini/skills/worktree-isolation/SKILL.md.tmpl` — 동일 변경 (각 4-6건)
  - `content/skills/agent-teams.md` — Prerequisites 섹션(v2.1.32+ 버전 요구) + Team Constraints 섹션(nested 금지, leader-only cleanup, 3-5명 권장, 영속 경로) 신설. Team Creation Pattern의 `Teammate()` → `Agent(team_name=..., name=...)` 공식 문법으로 교정
  - `templates/claude/commands/auto-router.md.tmpl` — Route B preflight 2단계(버전 + 환경변수) 추가, 에러 메시지 개선
  - `templates/codex/skills/agent-teams.md.tmpl` — 상단 ⚠️ Platform Note: Claude Code 전용 명시, Codex는 `spawn_agent` fallback
  - `templates/gemini/commands/auto-router.md.tmpl`, `templates/gemini/skills/agent-teams/SKILL.md.tmpl` — Platform Note 배너 + Route B 비활성화 + `--team` 경고 후 Route A fallback, 스테일 "Gemini CLI Agent Teams" 참조 제거
  - **Subagent frontmatter `permissionMode:` 필드는 공식 스펙이므로 그대로 유지** (Agent() 호출 파라미터와 별개 레이어)

### Docs

- **spec-writer 자체 품질 체크리스트 도입 문서 동기화 (SPEC-SPECWR-001)** (2026-04-19): `content/rules/spec-quality.md` 신규 체크리스트, `content/skills/spec-review.md`의 pre-review self-check, `content/agents/spec-writer.md`의 자체 검증 루프를 실제 산출물 기준으로 정렬하고 SPEC 문서를 completed 상태로 동기화
  - `content/rules/spec-quality.md`, `content/skills/spec-review.md`, `content/agents/spec-writer.md` — 체크리스트, pre-review self-check, 자체 검증 루프 source-of-truth 반영
  - `.autopus/specs/SPEC-SPECWR-001/{spec,plan,acceptance,research}.md` — completed 상태 동기화, validator/review 기준 정렬
  - 후속 보강: `research.md`의 `Self-Verify Summary` 관측 지점과 구조화된 `Open Issues` 스키마를 문서 규약으로 추가해 reviewer가 retry 경로를 문서 안에서 추적 가능하도록 보강

- **`/auto go --team` Route B 실행 절차 공백 수정** (2026-04-18): `--team` 플래그로 실행해도 core 4명 중 lead 1명만 spawn되어 멀티에이전트 협업이 작동하지 않던 문제를 수정. 실측 증거: `~/.claude/teams/spec-waitux-001/config.json` 의 members 배열에 team-lead 1명만 등록. 근본 원인: Route B 문서가 TeamCreate 호출 주체·시점, ToolSearch 선행 의존성, 4명 병렬 spawn 규칙, members 검증 게이트, phase별 SendMessage 디스패치를 명시하지 않음
  - `templates/claude/commands/auto-router.md.tmpl` — Route B에 **Team Orchestration Procedure (B1~B5)** 신설: ToolSearch → TeamCreate → 4명 병렬 Agent() spawn → `.members | length == 4` HARD GATE → SendMessage 오케스트레이션
  - `content/skills/agent-teams.md` — Lead 책임에서 "Creates the team" 문구 제거(teammates MUST NOT call TeamCreate), Team Creation Pattern을 top-level session 주체 + ToolSearch 선행 + verification gate 구조로 재작성
  - `templates/codex/skills/agent-teams.md.tmpl`, `templates/gemini/skills/agent-teams/SKILL.md.tmpl` — 플랫폼 비지원 명시를 유지한 채 Lead 문구와 코드 주석 정정

- **Route B 실측 smoke-test 기반 절차 정정** (2026-04-18): 1차 패치의 Route B 절차를 실제 `TeamCreate` + 3명 `Agent()` 호출로 smoke-test 한 결과, 공식 Claude Code Agent Teams API와 어긋난 4가지 세부 사항을 확인하고 정정. 실측 증거: `~/.claude/teams/team-probe-001/config.json` members=4 (team-lead + builder-1 + tester + guardian) 정상 생성 후 `SendMessage({type:"shutdown_request"})` ×3 + `TeamDelete()` 사이클 E2E 통과
  - **TeamCreate 파라미터명 정정**: `TeamCreate(name=...)` → `TeamCreate(team_name=..., agent_type="planner")` — 공식 스키마 파라미터는 `team_name` (기존 `name`은 오타)
  - **Lead 자동 등록 명시**: `TeamCreate`는 호출 시점에 메인 세션을 자동으로 `name: "team-lead"`, `agentType: <agent_type>`로 등록한다. Step B3은 **lead 제외 3명만 spawn**(builder-1 / tester / guardian)으로 축소 — lead Agent() 중복 spawn 방지
  - **SendMessage 주소 교정**: phase 오케스트레이션 매핑 표의 `to="lead"` → `to="team-lead"`. Phase 1 Planning은 메인 세션이 직접 담당하므로 SendMessage 불필요
  - **Step B6: Teardown 신설**: 구조화된 `{type:"shutdown_request"}`는 **per-teammate** 발송 필수 (broadcast `to:"*"`는 plain text 전용, structured payload rejected). `TeamDelete()`는 active members 남아 있으면 실패하므로 shutdown_request 후 `sleep 8` 대기 필수
  - 수정 파일: `templates/claude/commands/auto-router.md.tmpl`, `content/skills/agent-teams.md`, `templates/codex/skills/agent-teams.md.tmpl`, `templates/gemini/skills/agent-teams/SKILL.md.tmpl`

### Chore

- **SPEC review 산출물 ignore 정리** (2026-04-19): review 실행이 생성하는 `review.md`, `review-findings.json`을 runtime artifact로 간주하고 git 추적 대상에서 제외
  - `.gitignore` — `**/.autopus/specs/**/review.md`, `**/.autopus/specs/**/review-findings.json` 패턴 추가

## [v0.40.32] — 2026-04-17

### Changed

- **Claude Opus 4.7 Alignment**: 2026-04-16 Anthropic Opus 4.7 공식 출시에 맞춰 하네스 모델 ID/가격을 전면 동기화. 기존 cost estimator가 Opus 가격을 $15/$75로 과대 산정하던 오류도 함께 보정
  - `pkg/cost/pricing.go` — 모델 ID를 `claude-opus-4-7` / `claude-sonnet-4-6` / `claude-haiku-4-5`로 버전 명시, Opus 입력/출력 가격을 공식가 $5/$25로, Haiku를 $1/$5로 정정 (이전 $15/$75, $0.80/$4)
  - `pkg/cost/pricing_test.go`, `pkg/cost/estimator_test.go`, `pkg/cost/estimator_extra_test.go` — 모델명 assertion과 실제 달러 기대값(ultra/executor 4k 토큰 시 $0.04 등) 재계산
  - `pkg/worker/routing/config.go`, `pkg/worker/routing/{config,router}_test.go`, `pkg/worker/routing_integration_test.go` — Complex tier를 `claude-opus-4-7`로 승격
  - `pkg/config/defaults.go`, `autopus.yaml`, `configs/autopus.yaml` — Full 모드 기본 router tier `premium` / `ultra` 를 Opus 4.7로 갱신
  - `demo/simulate-claude.sh` — welcome banner 모델 표기를 `claude-opus-4-7`로 교체

### Docs

- **using-autopus Router Tier 예시 동기화**: `auto init` 이 생성하는 `configs/autopus.yaml` 기본값이 이미 `claude-opus-4-7` / `claude-sonnet-4-6` 버전 명시형인데, 가이드 문서의 예시 블록은 unversioned alias 로 남아 있어 사용자 혼란을 유발하던 불일치 제거
  - `content/skills/using-autopus.md`, `templates/codex/skills/using-autopus.md.tmpl`, `templates/gemini/skills/using-autopus/SKILL.md.tmpl` — router.tiers 예시 블록 통일

## [v0.40.29] — 2026-04-16

### Fixed

- **Codex Auto-Go Completion Handoff Gate Recovery**: Codex `@auto go ... --auto --loop` 가 구현/검증 요약만 남기고 종료하지 않도록 completion handoff contract를 source-of-truth와 회귀 테스트에 고정
  - `templates/codex/skills/auto-go.md.tmpl`, `templates/codex/prompts/auto-go.md.tmpl` — `Completion Handoff Gates` 와 `Final Output Contract` 를 추가해 `current_gate`, `phase_4_review_verdict`, `next_required_step`, `next_command`, `auto_progression_state` 가 비면 success-style completion summary로 닫지 못하게 보강
  - `pkg/adapter/codex/codex_surface_test.go`, `pkg/adapter/codex/codex_prompts_test.go` — generated Codex skill/prompt surface가 workflow lifecycle 뒤에 next-step handoff contract를 유지하는지 회귀 테스트 추가

## [v0.40.28] — 2026-04-16

### Fixed

- **Legacy SPEC Status Sync Recovery**: `auto spec review` 가 PASS 후 `approved` 상태를 새 scaffold SPEC뿐 아니라 기존 legacy SPEC 형식에도 안전하게 반영하도록 메타데이터 파서와 상태 갱신 경로를 복구
  - `pkg/spec/metadata.go` — `# SPEC: ...` + `**SPEC-ID**:` / `**Status**:` legacy metadata를 읽도록 보강하고, frontmatter 탐지를 문서 상단으로 제한해 본문 `---` 구분선을 잘못된 frontmatter로 오인하지 않도록 수정
  - `pkg/spec/metadata_test.go` — legacy ID/status 파싱, legacy status rewrite, 본문 separator 보호 회귀 테스트 추가

- **Status Dashboard Legacy Title Recovery**: `status` 대시보드가 legacy `# SPEC: ...` 헤더를 쓰는 SPEC에서도 ID, 상태, 제목을 다시 함께 표시하도록 회귀를 보강
  - `internal/cli/status_legacy_test.go` — `# SPEC: ...` + `**SPEC-ID**:` 형식의 legacy SPEC가 대시보드에서 제목과 상태를 유지하는지 검증

## [v0.40.27] — 2026-04-16

### Fixed

- **Auto Sync Completion Gate Recovery**: Codex `auto sync` 가 더 이상 컨텍스트/주석/커밋 게이트를 빠뜨린 채 완료를 선언하지 않도록 completion discipline을 source-of-truth와 테스트에 고정
  - `templates/codex/skills/auto-sync.md.tmpl`, `templates/codex/prompts/auto-sync.md.tmpl` — `Context Load`, `SPEC Path Resolution`, `@AX Lifecycle Management`, `Lore commit hash 또는 blocked reason`, `2-Phase Commit decision` 을 `Completion Gates` 로 승격하고, 암묵적 subagent 제한 시 사용자 opt-in 또는 `--solo` 확인을 먼저 요구하도록 보강
  - `pkg/adapter/codex/codex_prompts_test.go`, `pkg/adapter/codex/codex_surface_test.go` — generated Codex prompt/skill surface가 `@AX: no-op`, `commit hash`, completion gate 문구를 유지하는지 회귀 테스트 추가

- **OpenCode Runtime Wording Parity**: OpenCode generated `auto sync` skill에 Codex 전용 런타임 문구가 새지 않도록 변환기와 회귀 테스트를 보강
  - `pkg/adapter/opencode/opencode_util.go` — `task(...)` 문맥에서 `Codex 런타임 정책` 잔여 문구를 `OpenCode 런타임 정책` 으로 정규화
  - `pkg/adapter/opencode/opencode_test.go`, `pkg/adapter/opencode/opencode_sync_gate_test.go` — shared `.agents/skills/auto-sync/SKILL.md` 에 completion gate와 OpenCode wording parity가 유지되는지 검증

## [v0.40.26] — 2026-04-16

### Fixed

- **Workspace Policy Context Propagation**: `auto setup` 이 루트 저장소 역할과 nested repo 경계, generated/runtime 추적 정책을 별도 `workspace.md` 문서로 기록하고 이후 라우터가 공통 컨텍스트로 다시 읽도록 정렬
  - `templates/codex/skills/auto-setup.md.tmpl`, `templates/codex/prompts/auto-setup.md.tmpl` — `workspace.md` 를 `.autopus/project/` 핵심 산출물로 승격하고 meta workspace / source-of-truth / generated-runtime 경로 기록 규약 추가
  - `templates/codex/skills/auto-go.md.tmpl`, `templates/codex/prompts/auto-go.md.tmpl`, `templates/codex/skills/auto-sync.md.tmpl`, `templates/codex/prompts/auto-sync.md.tmpl` — 구현/동기화 단계가 `.autopus/project/workspace.md` 를 공통 프로젝트 컨텍스트로 로드하도록 보강
  - `pkg/adapter/codex/codex_context_docs.go`, `pkg/adapter/codex/codex_skill_render.go`, `pkg/adapter/opencode/opencode_router_contract.go`, `pkg/adapter/opencode/opencode_util.go`, `templates/claude/commands/auto-router.md.tmpl` — Codex prompt/plugin router, OpenCode shared router/alias command, Claude router가 모두 동일한 workspace policy context load 및 canonical router hand-off 계약을 따르도록 정렬
  - `pkg/adapter/codex/codex_workspace_context_test.go`, `pkg/adapter/opencode/opencode_workspace_context_test.go`, `pkg/adapter/claude/claude_workspace_context_test.go` — `workspace.md` 전파 회귀 테스트를 추가해 플랫폼별 contract drift를 다시 통과하지 못하게 보강

## [v0.40.25] — 2026-04-16

### Fixed

- **Codex Router Prompt Contract Recovery**: Codex `@auto` 메인 prompt surface가 workflow skill 쪽에만 있던 브랜딩/실행 계약을 prompt에도 동일하게 주입하고, 대형 프로젝트 문서가 잘리지 않도록 기본 project doc budget을 상향
  - `pkg/adapter/codex/codex_prompts.go`, `pkg/adapter/codex/codex_skill_render.go` — generated `.codex/prompts/auto*.md` 에 canonical branding block과 `Router Execution Contract` 를 주입
  - `templates/codex/config.toml.tmpl`, `pkg/adapter/codex/codex_lifecycle.go` — `project_doc_max_bytes` 기본값을 `262144` 로 상향하고, router prompt / config drift를 `validate` 에서 탐지하도록 보강
  - `pkg/adapter/codex/codex_*_test.go` — branding, router contract, Context7 rule, doc budget 회귀 테스트 추가

- **Context7 Web Fallback Contract Recovery**: 외부 라이브러리 문서 조회 규칙이 이제 `Context7 MCP 우선 → 실패 시 web search fallback` 계약을 공통 rule, pipeline skill, Codex/OpenCode generated surface 전반에서 일관되게 유지
  - `content/rules/context7-docs.md`, `content/skills/agent-pipeline.md`, `pkg/adapter/codex/codex_extended_skill_rewrites_agents.go` — Context7 실패 시 official docs / release notes / API reference 중심 web fallback 절차를 문서화
  - `pkg/content/skill_transformer_replace.go` — non-Claude platform surface에서 `mcp__context7__*` references를 단순 `WebSearch` 치환이 아니라 Context7-first / web-fallback 의미가 보존되는 안내로 변환
  - `pkg/adapter/opencode/opencode_lifecycle.go`, `pkg/adapter/opencode/opencode_test.go`, `pkg/content/*test.go` — OpenCode/Codex validate와 content transformer 회귀 테스트로 fallback 계약 누락을 다시 통과하지 못하게 보강

## [v0.40.24] — 2026-04-16

### Fixed

- **Acceptance Gate Lifecycle Recovery**: `spec validate` 와 pipeline validate/review 경로가 더 이상 `acceptance.md` 를 무시하지 않고, scaffold 기본 시나리오 형식도 실제 Gherkin 파서와 일치하도록 복구
  - `pkg/spec/template.go`, `pkg/spec/gherkin_parser.go` — `spec.Load()` 가 `acceptance.md` 를 함께 로드해 `AcceptanceCriteria` 를 채우고, `### Scenario 1:` / `### Edge Case 1:` scaffold 헤더를 파싱하도록 정렬
  - `pkg/pipeline/phase_prompt.go`, `pkg/spec/template_test.go`, `pkg/pipeline/phase_prompt_test.go`, `internal/cli/cli_extra_test.go` — `test_scaffold` / `implement` / `validate` / `review` 프롬프트에 acceptance context를 주입하고, scaffolded SPEC validate 회귀를 추가

- **Codex Shared Skill Branding Recovery**: Codex 에서 `@auto` 브랜드 배너가 간헐적으로 사라지던 문제를, 실제 우선 선택되던 shared `.agents/skills/` 경로에도 canonical branding block을 주입하도록 보강
  - `pkg/adapter/opencode/opencode_util.go`, `pkg/adapter/opencode/opencode_skills.go`, `pkg/adapter/opencode/opencode_workflow_custom.go` — OpenCode가 소유하는 shared skill surface에도 `## Autopus Branding` 과 canonical banner injection을 적용
  - `pkg/adapter/opencode/opencode_test.go` — generated `.agents/skills/auto*.md` 가 branding header를 유지하는지 회귀 테스트 추가

## [v0.40.20] — 2026-04-15

### Fixed

- **OpenCode Router SPEC Path Resolution Contract Recovery**: OpenCode `auto` command/skill 생성물이 shared router contract의 `SPEC Path Resolution` 섹션을 다시 포함하고, OpenCode 표면에 Codex 전용 wording이 새지 않도록 정렬
  - `pkg/adapter/opencode/opencode_router_contract.go`, `pkg/adapter/opencode/opencode_commands.go`, `pkg/adapter/opencode/opencode_skills.go` — Claude canonical router에서 SPEC path resolution block을 추출해 OpenCode `auto` surfaces에 재주입하고, `TARGET_MODULE` / `WORKING_DIR` / `Available SPECs` 계약을 복원
  - `pkg/adapter/opencode/opencode_test.go` — 생성된 `.opencode/commands/auto.md` 와 `.agents/skills/auto/SKILL.md` 가 `SPEC Path Resolution` 을 유지하고 Codex wording leak이 없는지 회귀 테스트 추가

- **Workspace-Root Submodule SPEC Resolution Regression Coverage**: workspace root에서 실행되는 OpenCode SPEC 워크플로우가 `Autopus/.autopus/specs/...` 같은 실제 서브모듈 SPEC를 놓치지 않도록 회귀 케이스를 보강
  - `pkg/spec/resolve_test.go` — `SPEC-OPCOCK-001` 이 workspace root 기준으로 `Autopus` 서브모듈에서 정확히 resolve 되는지 검증

## [v0.40.18] — 2026-04-14

### Fixed

- **Codex `@auto` Branding Injection**: Codex local plugin skill surface가 router/prompt에는 있던 문어 배너 지시를 실제 `@auto` plugin workflow skill에도 동일하게 주입하도록 정렬
  - `pkg/adapter/codex/codex_skill_render.go`, `pkg/adapter/codex/codex_workflow_custom.go` — router skill과 workflow/custom workflow skill 생성 경로 모두에 canonical Autopus branding block을 삽입
  - `pkg/adapter/codex/codex_surface_test.go` — `.agents` / `.autopus/plugins` Codex skill surfaces가 branding header를 유지하는지 회귀 테스트 추가

## [v0.40.17] — 2026-04-14

### Added

- **OpenCode Strategic Skill Canonical Sources**: OpenCode가 더 이상 Claude 전용 산출물에 의존하지 않도록 `product-discovery`, `competitive-analysis`, `metrics`를 canonical `content/skills/`에 추가
  - `content/skills/product-discovery.md`, `content/skills/competitive-analysis.md`, `content/skills/metrics.md` — platform-agnostic source로 승격하여 OpenCode `.agents/skills/`에도 동일하게 배포되도록 정렬

### Fixed

- **Codex Workflow and Rule Parity Recovery**: Codex 하네스가 Claude Code 기준 workflow surface와 규칙 패키징을 다시 충족하도록 정렬
  - `pkg/adapter/codex/codex_workflow_specs.go`, `pkg/adapter/codex/codex_workflow_custom.go`, `pkg/adapter/codex/codex_prompts.go`, `templates/codex/prompts/auto.md.tmpl` — `@auto` router와 workflow generation이 `status`, `map`, `why`, `verify`, `secure`, `test`, `dev`, `doctor`를 포함한 전체 helper flow surface를 생성하도록 복구
  - `pkg/adapter/codex/codex_rules.go`, `pkg/adapter/codex/codex_skill_render.go`, `pkg/adapter/codex/codex_skill_template_mappings.go`, `pkg/adapter/codex/codex_standard_skills.go` — Codex rule/skill rendering이 stub `@import` 대신 canonical content와 Codex-native semantics를 사용하고 `branding`, `project-identity` rule parity를 회복
  - `pkg/adapter/codex/codex_*_test.go`, `pkg/adapter/parity_test.go`, `pkg/adapter/integration_test.go` — prompt/rule count와 cross-platform parity 회귀 테스트를 추가해 workflow 누락과 규칙 드리프트를 다시 통과하지 못하게 보강

- **OpenCode Helper Flow Surface Recovery**: OpenCode router와 command surface가 `setup` 외 helper flow도 노출하고, Codex prompt 단일 의존 없이 OpenCode 전용 contract를 사용하도록 정리
  - `pkg/adapter/opencode/opencode_specs.go`, `pkg/adapter/opencode/opencode_router_contract.go`, `pkg/adapter/opencode/opencode_workflow_custom.go` — `status`, `map`, `why`, `verify`, `secure`, `test`, `dev`, `doctor` helper flow inventory와 custom skill/command body 추가
  - `pkg/adapter/opencode/opencode_commands.go`, `pkg/adapter/opencode/opencode_skills.go` — router/command generation이 OpenCode-native helper semantics와 상세 스킬 목록을 사용하도록 갱신

- **OpenCode Plugin Wiring Diagnostics**: hook plugin이 파일만 생성되고 `opencode.json`에는 연결되지 않던 결손을 수정하고, registration 누락을 validation에서 탐지하도록 보강
  - `pkg/adapter/opencode/opencode_config.go`, `pkg/adapter/opencode/opencode.go`, `pkg/adapter/opencode/opencode_lifecycle.go`, `pkg/adapter/opencode/opencode_util.go` — managed plugin 경로를 기본 등록하고 plugin array parsing/validation을 보강
  - `pkg/adapter/opencode/opencode_runtime_test.go`, `pkg/adapter/opencode/opencode_test.go` — helper flow surface, plugin registration, strategic skill generation 회귀 테스트 추가

- **Queued Task Deadline Guard**: 이미 만료된 worker task가 semaphore 슬롯을 선점하거나 subprocess를 시작하지 않도록 acquire 단계의 cancellation 우선순위를 보강
  - `pkg/worker/parallel/semaphore.go`, `pkg/worker/loop_runtime_fix_test.go` — 만료된 context는 즉시 거절하고 queued-task expiry 회귀 테스트 기대를 다시 만족하도록 정렬
  - `pkg/adapter/integration_test.go` — Codex prompt surface 확장에 맞춰 E2E prompt count 기대치를 갱신

- **Worker MCP Startup Compatibility**: Codex가 worker MCP 서버를 startup 단계에서 타입 오류 없이 수용하도록 초기 lifecycle, tool schema, resource 응답 형식을 최신 MCP 계약에 가깝게 정렬
  - `pkg/worker/mcpserver/server.go`, `pkg/worker/mcpserver/server_test.go` — `initialize` protocol negotiation, `tools/list` schema metadata, `tools/call` structured result envelope, `resources/templates/list`, `resources/read` contents wrapper 추가
  - `pkg/worker/mcpserver/resources.go`, `pkg/worker/mcpserver/resources_test.go` — resource title/template metadata를 추가해 execution URI template discovery를 노출
  - `templates/codex/config.toml.tmpl` — Codex generated config가 `autopus` MCP를 다시 기본 등록해도 startup validation을 통과하도록 정렬

## [v0.40.13] — 2026-04-14

### Fixed

- **OpenCode Workflow Surface Alignment**: OpenCode가 `auto` workflow를 얇은 prompt entrypoint가 아니라 실제 skill 템플릿과 맞는 표면으로 생성하도록 정렬
  - `pkg/adapter/opencode/opencode_specs.go`, `pkg/adapter/opencode/opencode_skills.go` — workflow별 prompt와 skill source를 분리하고, `auto`는 thin router / 하위 workflow는 실제 skill 템플릿으로 생성되도록 조정
  - `pkg/adapter/opencode/opencode_util.go` — OpenCode `task(...)` / command entrypoint semantics에 맞는 body normalization과 예제 치환 보강
  - `pkg/adapter/opencode/opencode_test.go` — workflow skill / command surface 회귀 테스트 추가

- **Codex Router Thin-Skill Stabilization**: Codex router skill이 더 이상 Claude router rewrite에 의존하지 않고 Codex thin router semantics로 생성되도록 정리
  - `pkg/adapter/codex/codex_standard_skills.go`, `pkg/adapter/codex/codex_skill_render.go`, `pkg/adapter/codex/codex_plugin_manifest.go` — router rendering과 plugin metadata를 분리하고 300-line limit를 만족하도록 파일 분할
  - `pkg/adapter/codex/codex_test.go` — `.agents/.autopus/.codex` 전 surface 회귀 테스트 추가

- **Gemini Canary Workflow Parity**: Gemini `canary` command가 참조하던 `auto-canary` skill 누락을 보완해 command-skill 정합성을 복구
  - `templates/gemini/skills/auto-canary/SKILL.md.tmpl` — Gemini 전용 `auto-canary` skill 추가
  - `pkg/adapter/gemini/gemini_test.go` — workflow command와 대응 skill 생성 정합성 회귀 테스트 추가

## [v0.40.12] — 2026-04-14

### Fixed

- **`auto update` New Platform Detection**: 바이너리 업데이트 후 새로 설치한 OpenCode 같은 supported CLI가 기존 프로젝트의 `auto update` 경로에서 자동 반영되지 않던 문제 수정
  - `internal/cli/update.go`, `internal/cli/init_helpers.go` — `update`가 현재 설치된 supported platform을 다시 감지해 `autopus.yaml`에 누락된 플랫폼을 추가하고, 같은 실행에서 해당 하네스를 생성하도록 정렬
  - `internal/cli/update_test.go` — 기존 `claude-code` 프로젝트에서 `opencode` 설치 후 `auto update`가 `opencode.json`과 `.opencode/` 하네스를 생성하는 회귀 테스트 추가

## [v0.40.11] — 2026-04-14

### Fixed

- **Worker Queue Timeout Separation**: worker 실행 대기와 provider 세마포어 대기를 분리해, 혼잡 상황에서도 queue starvation과 잘못된 타임아웃 해석이 줄어들도록 정리
  - `pkg/worker/loop.go`, `pkg/worker/loop_exec.go`, `pkg/worker/loop_test.go` — worker loop가 queue wait / execution timeout을 구분해 처리하고 직렬화 경로를 더 명확히 검증하도록 보강
  - `internal/cli/worker_start.go`, `internal/cli/worker_start_test.go` — worker start 경로가 새 timeout semantics와 직렬화 보강을 반영하도록 조정

- **Codex Worker Concurrency Stabilization**: Codex worker 동시 실행 시 output artifact와 setup 경로가 더 안정적으로 유지되도록 보강
  - `internal/cli/worker_setup_wizard.go`, `internal/cli/worker_setup_wizard_test.go` — setup wizard가 최신 worker concurrency 흐름과 일치하도록 조정

## [v0.40.10] — 2026-04-14

### Added

- **OpenCode Native Harness Generation**: `auto init/update`가 이제 OpenCode를 정식 하네스 설치 플랫폼으로 지원하여 `.opencode/` 네이티브 산출물과 `.agents/skills/` 표준 스킬을 함께 생성
  - `pkg/adapter/opencode/*` — OpenCode 어댑터를 stub에서 실제 generate/update/validate/clean 구현으로 확장하고 `AGENTS.md`, `opencode.json`, `.opencode/rules/`, `.opencode/agents/`, `.opencode/commands/`, `.opencode/plugins/`를 생성
  - `internal/cli/init_helpers.go`, `internal/cli/update.go`, `internal/cli/doctor.go`, `internal/cli/platform.go`, `internal/cli/init.go` — OpenCode를 init/update/doctor/platform add-remove 및 gitignore 경로에 연결
  - `pkg/adapter/opencode/opencode_test.go`, `pkg/content/opencode_transform_test.go` — OpenCode 산출물 생성, 설정 병합, CLI 연결, 변환 규칙 회귀 테스트 추가