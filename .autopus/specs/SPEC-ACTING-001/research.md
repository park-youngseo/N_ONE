# SPEC-ACTING-001 Research: N-ONE 성인 연기학원 수강생 대시보드

## Codebase Analysis
본 기능은 완전히 새로운 요구사항으로, 기존 코드베이스 내에 "연기학원", "수강생", "출결", "대본"과 관련된 비즈니스 로직이 존재하지 않음을 확인하였습니다. 따라서 Greenfield 프로젝트로 진행합니다.

### Target Files
- 신규 패키지 및 디렉터리 구조 생성이 필요합니다 (`pkg/acting`, `web/dashboard-app` 등).

### Dependencies
- 기존 `autopus-adk` CLI 엔진과는 독립적으로 동작하도록 설계하되, 공통 유틸리티(로깅, 설정 관리 등)가 필요할 경우 `pkg/config` 등을 참조할 수 있습니다.

## Lore Decisions
- 현재까지 연기학원 도메인과 관련된 과거 의사결정(Lore)은 발견되지 않았습니다.

## Architecture Compliance
- `pkg/` 하위에는 도메인 중심의 비즈니스 로직을 배치하고, `internal/` 하위에는 API 핸들러 등 인프라 세부 사항을 배치하는 기존의 계층 구조를 따를 것을 권장합니다.

## Key Findings
- **Keyword Search**: `grep_search` 결과 "연기학원", "수강생" 등에 대한 매칭이 없음을 확인했습니다.
- **Project Scope**: 단순한 CLI 도구가 아닌 웹 대시보드 형태의 확장이 필요하므로, 별도의 웹 서버(API)와 프론트엔드 빌드 파이프라인이 필요합니다.

## Recommendations
- **UI Framework**: React + TypeScript 사용을 강력히 권장하며, 스타일링은 프로젝트의 깔끔한 디자인 기조를 유지하기 위해 Vanilla CSS를 사용합니다.
- **API Specification**: 프론트엔드와 백엔드 간의 명확한 계약을 위해 OpenAPI(Swagger) 또는 간단한 JSON 스키마를 선제적으로 정의하는 것이 좋습니다.
- **Testing**: 웹 UI 검증을 위해 `auto verify` (Playwright 기반)를 적극적으로 활용할 수 있는 구조로 설계하십시오.
