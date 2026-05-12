**File Ownership:** `content/rules/*.md`, `templates/codex/rules/*.md.tmpl`, `templates/gemini/rules/*.md.tmpl`

### R23 — 에이전트 메타데이터 공유

> WHEN Codex/Gemini 에이전트 생성 시,
> THE SYSTEM SHALL `content/agents/` 내 공유 에이전트 메타데이터를 플랫폼별 포맷으로 변환한다.

**Acceptance Criteria:**
- AC-23.1: 에이전트 name, description, core instructions가 플랫폼 간 일관성 유지
- AC-23.2: 플랫폼별 포맷 차이(TOML vs YAML frontmatter)만 분기

**File Ownership:** `content/agents/*.md`, `templates/codex/agents/*.toml.tmpl`, `templates/gemini/agents/*.md.tmpl`

### R24 — 테스트 유틸리티

> WHEN 어댑터 단위 테스트 작성 시,
> THE SYSTEM SHALL 공통 테스트 헬퍼(임시 디렉터리, 파일 검증, 매니페스트 비교)를 제공한다.

**Acceptance Criteria:**
- AC-24.1: `pkg/adapter/testutil_test.go`에 `setupTestDir()`, `assertFileExists()`, `assertFileContains()` 헬퍼 존재
- AC-24.2: Codex, Gemini 테스트에서 공통 헬퍼 사용

**File Ownership:** `pkg/adapter/testutil_test.go`

### R25 — 매니페스트 호환성

> WHEN 새 파일 타입(prompts, agents, hooks, settings)이 매니페스트에 추가될 때,
> THE SYSTEM SHALL 기존 매니페스트 스키마와 하위 호환성을 유지한다.

**Acceptance Criteria:**
- AC-25.1: 새 파일이 기존 `ManifestFromFiles()` 함수로 정상 등록됨
- AC-25.2: 이전 버전 매니페스트에서 `LoadManifest()` 호출 시 오류 없이 로드됨
- AC-25.3: 매니페스트 파일 수가 Codex +20, Gemini +20 증가

**File Ownership:** `pkg/adapter/manifest.go`, `pkg/adapter/manifest_test.go`

---

## Domain 4: Integration (R26-R30)

### R26 — auto init E2E 테스트 (3 플랫폼)

> WHEN `auto init --platform {codex|gemini|claude}` 실행 시,
> THE SYSTEM SHALL 각 플랫폼에 대해 생성 파일 수, 디렉터리 구조, 파일 내용을 검증하는 E2E 테스트를 통과한다.

**Acceptance Criteria:**
- AC-26.1: Codex init 후 `.codex/prompts/auto/`, `.codex/agents/`, `.codex/skills/`, `hooks.json` 존재 검증
- AC-26.2: Gemini init 후 `.gemini/commands/auto/`, `.gemini/agents/`, `.gemini/skills/autopus/`, `.gemini/rules/autopus/` 존재 검증
- AC-26.3: Claude init 후 기존 동작 무변경 (regression test)

**File Ownership:** `pkg/adapter/integration_test.go`

### R27 — auto update E2E 테스트

> WHEN `auto update` 실행 시,
> THE SYSTEM SHALL 사용자 커스터마이징이 보존되고 autopus 섹션만 업데이트됨을 검증하는 E2E 테스트를 통과한다.

**Acceptance Criteria:**
- AC-27.1: AGENTS.md 마커 외부 사용자 내용 보존 검증
- AC-27.2: GEMINI.md 마커 외부 사용자 내용 보존 검증
- AC-27.3: settings.json 사용자 키 보존 검증
- AC-27.4: hooks.json 사용자 훅 보존 검증

**File Ownership:** `pkg/adapter/integration_test.go`

### R28 — 매니페스트 파일 수 검증

> WHEN 어댑터가 Generate()를 완료할 때,
> THE SYSTEM SHALL 매니페스트의 파일 수가 Codex >= 26, Gemini >= 26임을 검증한다.

**Acceptance Criteria:**
- AC-28.1: Codex 매니페스트: AGENTS.md(1) + skills(10+) + prompts(6) + agents(5+) + hooks(1) + config(1) >= 24
- AC-28.2: Gemini 매니페스트: GEMINI.md(1) + skills(10+) + commands(6) + agents(5+) + rules(5+) + settings(1) >= 28

**File Ownership:** `pkg/adapter/codex/codex_test.go`, `pkg/adapter/gemini/gemini_test.go`

### R29 — 300줄 제한 준수 검증

> WHEN 새 Go 파일이 생성될 때,
> THE SYSTEM SHALL 이 SPEC에서 새로 생성하거나 수정한 Go 파일이 300줄 이하임을 검증한다.

**Acceptance Criteria:**
- AC-29.1: `pkg/adapter/codex/`에서 **신규 또는 수정된** 파일 300줄 이하
- AC-29.2: `pkg/adapter/gemini/`에서 **신규 또는 수정된** 파일 300줄 이하
- AC-29.3: 기존 300줄 초과 파일은 이 SPEC 범위 외 (별도 리팩토링 SPEC으로 처리)

**File Ownership:** CI pipeline, `Makefile`

### R30 — 매니페스트 JSON 스키마 업데이트

> WHEN 새 파일 타입이 추가될 때,
> THE SYSTEM SHALL `claude-code-manifest.json`에 새 어댑터 파일 정보를 반영한다.

**Acceptance Criteria:**
- AC-30.1: `.autopus/claude-code-manifest.json`에 codex, gemini 어댑터 파일 목록 업데이트
- AC-30.2: 매니페스트 버전 범프

**File Ownership:** `.autopus/claude-code-manifest.json`