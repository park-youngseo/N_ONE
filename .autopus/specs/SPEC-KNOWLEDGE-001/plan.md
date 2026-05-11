# SPEC-KNOWLEDGE-001 Plan: 지식 습득 시스템의 품질 검증 로직 설계

## Implementation Strategy
- **Layer**: 로직은 CLI 엔진의 백엔드 레이어인 `pkg/worker/knowledge` 패키지에 집중 배치한다.
- **Component 1: Ingestion Gateway**: 외부 데이터가 들어오는 즉시 `LosslessValidator` 객체를 통과하게 하여 원본 크기, 라인 수, 앞뒤 3줄을 검증하는 미들웨어를 구축한다.
- **Component 2: Semantic Comparator**: 데이터 병합 전 코사인 유사도를 측정하기 위해 내부 임베딩 비교기(Threshold=1.0)를 강제하는 구조를 적용한다.
- **Component 3: Domain Lens Analyzer**: 5대 렌즈(Bitmask 적용)를 순차적으로 통과해야만 최종 데이터로 승격(Promote)되는 상태 머신(State Machine)을 도입한다.

## File Impact Analysis

| 파일 | 작업 | 설명 |
|------|------|------|
| `pkg/worker/knowledge/lossless_validator.go` | 신규 | 무손실 수집 검증 및 자가 증명 로직 |
| `pkg/worker/knowledge/semantic_comparator.go` | 신규 | 100% 논리적 일치 판단 알고리즘 |
| `pkg/worker/knowledge/domain_lens.go` | 신규 | 5대 렌즈 추론 및 등급(A~F) 스코어링 함수 |
| `internal/cli/learn.go` | 수정 | 지식 수집 커맨드 호출 시 Ingestion Gateway를 통하도록 파이프라인 수정 |

## Architecture Considerations
- 외부 의존성 최소화: 코사인 유사도 연산 등은 외부 무거운 라이브러리(e.g., Python AI 모델)에 의존하지 않고 Go 기반의 경량화된 토큰 벡터 비교 알고리즘 또는 기존 프로젝트 내장 모듈만 활용하여 속도를 보장한다.
- Fail-Fast: `LosslessValidator`에서 실패 시 지연 없이 `panic` 기반 에러 핸들링으로 즉시 셧다운(Fail)을 트리거한다.

## Tasks
- [ ] Task 1: `LosslessValidator` 구현 및 무손실 자가 증명(터미널 출력) 로직 작성
- [ ] Task 2: 코사인 유사도 1.0 검증 전용 `SemanticComparator` 구현
- [ ] Task 3: `domain_lens.go`에 5대 렌즈 상태 머신 및 등급(A~F) 판별 수식 추가
- [ ] Task 4: 기존 CLI 커맨드와의 통합 및 `lessons_learned.md` 파일 쓰기 훅 연결

## Risks & Mitigations

| 리스크 | 영향도 | 대응 |
|--------|--------|------|
| 대용량 파일(1GB 이상) 파싱 시 라인 단위 비교로 인한 메모리 부족(OOM) | 매우 높음 | 메모리에 전체를 올리지 않고 Chunk/Stream 단위의 스트리밍 해싱(Streaming Hashing) 로직 적용 |
| 코사인 유사도 1.0 달성 불가 | 높음 | 부동소수점 오차(Floating Point Error)를 감안하여 `1.0 - epsilon(1e-9)`을 임계값으로 설정 |

## Exit Criteria
- [ ] 모든 Requirement를 만족하는 3개의 핵심 컴포넌트 유닛 테스트 통과 (커버리지 100% 필수)
- [ ] E2E 레벨에서 터미널 출력 형식(라인 수, 첫/끝 3줄)이 정규식에 부합함을 증명
