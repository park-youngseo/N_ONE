# SPEC-ACTING-001 Plan: N-ONE 성인 연기학원 수강생 대시보드

## Implementation Strategy
- **Frontend**: React (TypeScript) + Vanilla CSS를 사용하여 빠르고 현대적인 웹 대시보드를 구축합니다. WebSocket 기반의 `useWebSocket` 커스텀 훅을 구현하여 서버 측 상태 변화를 실시간으로 Redux 또는 Context API 상태 트리에 반영합니다.
- **Backend (API)**: Go (Gin or Echo)를 사용하여 RESTful API를 제공하며, WebSocket 핸들러를 추가하여 커넥션 풀을 관리합니다.
- **Data Persistence & Cache**: PostgreSQL을 주 저장소로 사용하고, Redis를 읽기 캐시 및 Pub/Sub 채널 브로커로 구성하여 분산 서버 환경에서도 이벤트를 안정적으로 브로드캐스팅할 수 있도록 합니다.
- **Pattern**: Clean Architecture를 적용하여 도메인 로직(Entity, UseCase)을 외부 인프라(DB, Redis, WebSocket)와 철저히 분리합니다.

## File Impact Analysis

| 파일 | 작업 | 설명 |
|------|------|------|
| `pkg/acting/domain/` | 기존 수정 | Student, Script, Attendance 도메인 및 실시간 이벤트 구조체 추가 |
| `pkg/acting/service/` | 기존 수정 | 출결 및 진도 계산 비즈니스 로직에 Redis Pub/Sub 발행 로직 결합 |
| `pkg/acting/infrastructure/redis/` | 신규 생성 | Redis 클라이언트 및 Pub/Sub 구독/발행 어댑터 구현 |
| `internal/api/handler/acting/` | 기존 수정 | HTTP 핸들러 및 `ws_handler.go` (WebSocket 연결 관리) 추가 |
| `web/dashboard-app/` | 기존 수정 | React 기반 대시보드에 WebSocket 구독 훅(`useActingEvent.ts`) 연결 |

## Architecture Considerations
- `pkg/acting/domain`은 외부 라이브러리(Redis, Gorilla WebSocket 등)를 참조하지 않아야 합니다. 오직 Go 표준 라이브러리만으로 이벤트 페이로드를 정의합니다.
- 프론트엔드의 오프라인 대응: 브라우저 IndexedDB를 활용해 오프라인 상태의 변경 이벤트를 적재하고, WebSocket `onopen` 콜백 발생 시 미전송 큐를 플러시(Flush)하는 아키텍처를 도입합니다.
- 모든 API 응답 및 WebSocket 메시지 포맷은 JSON 스키마로 강제합니다.

## Tasks
- [x] M1: 데이터베이스 스키마 설계 (InMemory 모킹으로 대체 완료)
- [x] M2: Go 백엔드 기본 API 구현 및 유닛 테스트 완료
- [x] M3: React 기반 대시보드 레이아웃 및 컴포넌트 개발 (스캐폴딩 완료)
- [x] M4: WebSocket 핸들러 및 Redis Pub/Sub 메시지 브로커 통합 구현
- [x] M5: 프론트엔드 실시간 동기화 훅 개발 및 백엔드 연동
- [x] M6: 오프라인/온라인 단절 복구 알고리즘(Exponential Backoff) 구현
- [x] M7: 백엔드-프론트엔드 통합 E2E 테스트(다중 클라이언트 접속 환경 모사) 작성

## Risks & Mitigations

| 리스크 | 영향도 | 대응 |
|--------|--------|------|
| WebSocket 연결 폭증으로 인한 서버 OOM | 높음 | Connection 개수 제한 및 Ping-Pong 기반의 Dead Connection Pruning 로직 구현 |
| 오프라인 시 데이터 유실 | 중간 | 클라이언트 측 IndexedDB 기반 Event Queue 적용 |
| 대본 데이터의 복잡성 | 중간 | 대본을 Scene 단위로 정규화하여 관리 |

## Dependencies
- Go: `github.com/gorilla/websocket` (웹소켓), `github.com/redis/go-redis/v9` (레디스), `github.com/gin-gonic/gin`
- Frontend: `react`, `typescript`, `recharts`

## Exit Criteria
- [ ] M4~M7의 모든 고도화 Requirements 구현 완료
- [ ] 다중 탭 접속 시 브로드캐스팅 지연율 100ms 이내 달성
- [ ] 테스트 커버리지 85%+ 유지 (WebSocket 레이어 포함)
- [ ] E2E 환경에서 연결 끊김 및 복구 테스트 완벽 통과
