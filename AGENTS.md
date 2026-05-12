<!-- AUTOPUS:BEGIN -->
# Autopus-ADK Harness

> 이 섹션은 Autopus-ADK에 의해 자동 생성됩니다. 수동으로 편집하지 마세요.

- **프로젝트**: N_ONE
- **모드**: full
- **플랫폼**: claude-code, codex, gemini-cli

## Installed Components

- Codex Rules: .codex/rules/autopus/
- Codex Skills: .codex/skills/
- Codex Agents: .codex/agents/
- Codex Config: config.toml
- Shared Agent Skills: .agents/skills/
- Plugin Marketplace: .agents/plugins/marketplace.json


## Language Policy

IMPORTANT: Follow these language settings strictly for all work in this project.

- **Code comments**: ko
- **Commit messages**: ko
- **AI responses**: ko

## Execution Model

- **Codex**: 하네스 기본값은 spawn_agent(...) 기반 subagent-first 입니다.
- **Codex --auto**: @auto ... --auto 가 포함되면, 기본 subagent pipeline 진행에 대한 명시적 승인으로 해석합니다.
- **Codex Runtime Caveat**: 현재 세션의 Codex 런타임 정책이 암묵적 spawn_agent(...) 호출을 제한하면, 조용히 단일 세션으로 폴백하지 말고 그 제약을 명시적으로 알린 뒤 사용자의 서브에이전트 opt-in 또는 --solo 선택을 받으세요.
- **Codex --team**: 미래의 native multi-agent surface를 위한 reserved compatibility flag입니다.


## Core Guidelines

### Karpathy Harness Rules (Legacy Core)

IMPORTANT: 구형 하네스 모델의 정수이자 안드레이 카파시의 철학인 아래 9계명을 모든 에이전트는 반드시 준수한다.

1. **가정 금지 (Ask First)**: 구현 전 이해한 바를 명확히 기술하고, 모호하면 반드시 질문하라.
2. **최소 코드 (Keep it Simple)**: 오버엔지니어링 금지. 문제를 해결하는 가장 짧은 코드를 지향하라.
3. **외과수술식 수정 (Precision)**: 요청받은 부분만 수정하라. 인접한 코드나 포맷을 멋대로 고치지 마라.
4. **성공 기준 중심 (Goal-Oriented)**: 명확한 성공 기준(Test/Checklist)이 달성될 때까지 반복하라.
5. **한국어 정제 (Language Policy)**: 문장 끝에 콜론(:) 사용 금지, 반드시 마침표(.)로 종결하라.
6. **헤더 주석 (File Header)**: 모든 신규 파일 최상단에 해당 파일의 역할을 설명하는 한 줄 주석을 넣어라.
7. **작업 3종 세트 (Artifacts)**: 큰 작업은 [계획서 -> 체크리스트 -> 컨텍스트 노트] 순으로 관리하라.
8. **선 검증 후 보고 (Verify First)**: 작업 완료 보고 전 반드시 로컬 테스트를 실행하여 검증하라.
9. **의미 단위 커밋 (Atomic Commit)**: 작업은 의미 있는 최소 단위로 커밋하며, 메시지는 명확해야 한다.

### Anti-Wrapper Shield (Low-level Defense)

IMPORTANT: 단순 챗봇이나 래퍼 수준의 낮은 지능적 실수를 방지하기 위해 다음을 준수한다.

1. **근거 없는 답변 금지 (Grounding)**: 모든 핵심 제안은 반드시 입고된 `knowledge/` 내의 파일명을 인용하여 근거를 제시하라. (예: "마케팅_01.md에 따르면...")
2. **환각 방지 (Self-Correction)**: 답변을 출력하기 직전, 스스로 "이 내용에 거짓은 없는가?"를 1초간 자문하고 검증된 내용만 출력하라.
3. **비용 최적화 (Token Guard)**: 1,000자 이상의 대량 텍스트 처리 시 반드시 로컬 모델(Ollama)이나 Economy 티어를 우선 사용하고, Premium 모델은 최종 승인 단계에서만 호출하라.

### Supervisor Contract

IMPORTANT: 메인 세션은 얇은 라우터가 아니라 phase/gate를 관리하는 supervisor입니다. 각 단계마다 필수 단계, skip 조건, retry 한도, 다음 필수 단계를 명확히 유지하세요.

### Subagent Delegation

IMPORTANT: 3개 이상 파일 수정, 다중 도메인 변경, 또는 신규 코드 200줄 초과가 예상되면 기본적으로 서브에이전트를 사용하세요. 단, 읽기 위주 탐색/리서치/테스트 분석은 병렬 fan-out을 우선하고, 쓰기 위주 구현은 파일 소유권이 겹치면 순차 실행으로 전환하세요.

### Worker Contracts

IMPORTANT: 각 worker 프롬프트에는 반드시 소유 파일/모듈, 수정 금지 범위, 완료 기준, 반환 형식을 포함하세요. 최소 반환 필드는 `owned_paths`, `changed_files`, `verification`, `blockers`, `next_required_step` 입니다.

### Review Convergence

IMPORTANT: 리뷰는 discovery와 verification을 분리하세요. 첫 리뷰는 finding discovery에 집중하고, 재시도는 열린 finding 해결 여부만 diff 기준으로 확인하세요. 같은 범위를 무한 재탐색하지 마세요.

### File Size Limit

IMPORTANT: 생성 파일을 제외한 소스 파일은 300줄 이하를 유지하세요. 가능하면 200줄 이하를 목표로 분리하세요.

### Prompting Notes

IMPORTANT: 사용자가 계획만 요구한 경우를 제외하면, 긴 선행 계획만 출력하고 멈추지 마세요. 먼저 코드베이스를 확인하고, 필요한 경우 서브에이전트를 스폰한 뒤, 검증까지 이어서 진행하세요.

## Rules

See .codex/rules/autopus/ for Codex rule definitions.
See .codex/skills/agent-pipeline.md for phase and gate contracts.
See .codex/agents/ for Codex agent definitions.


## Agents

The following specialized agents are available.

### Annotator Agent

Phase 2.5 @AX tag scanning and application specialist.

### Architect Agent

시스템 아키텍처를 설계하고 기술 결정을 내리는 에이전트입니다.

### Debugger Agent

버그의 근본 원인을 분석하고 최소한의 수정으로 해결하는 에이전트입니다.

### Deep Worker Agent

장시간 실행이 필요한 복잡한 태스크를 체크포인트와 검증 루프를 통해 안전하게 완료하는 에이전트입니다.

### DevOps Agent

CI/CD, 컨테이너화, 인프라 설정을 전담하는 에이전트입니다.

### Executor Agent

TDD 또는 DDD 방법론에 따라 코드를 구현하는 에이전트입니다.

### Explorer Agent

코드베이스를 빠르게 탐색하고 구조를 분석하는 에이전트입니다.

### Frontend-Specialist Agent

Phase 3.5 Playwright E2E testing, screenshot analysis, and UX verification specialist.

### Perf-Engineer Agent

Benchmark execution, profiling, and performance regression detection specialist.

### Planner Agent

기능 기획과 요구사항 분석을 전담하는 에이전트입니다.

### Reviewer Agent

TRUST 5 기준으로 코드를 체계적으로 검토하는 에이전트입니다.

### Security Auditor Agent

OWASP Top 10 기준으로 보안 취약점을 탐지하고 수정하는 에이전트입니다.

### Spec Writer Agent

SPEC 문서를 생성하는 전문 에이전트입니다.

### Tester Agent

테스트를 설계하고 구현하는 전담 에이전트입니다.

### UX Validator Agent

Claude Vision(멀티모달)으로 프론트엔드 스크린샷을 분석하여 레이아웃 및 접근성 문제를 탐지하는 에이전트입니다.

### Validator Agent

코드 품질을 빠르게 검증하는 경량 에이전트입니다.


## Rules

See .codex/rules/autopus/ for Codex guidance.

<!-- AUTOPUS:END -->
