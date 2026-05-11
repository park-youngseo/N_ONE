# SPEC-KNOWLEDGE-001: 지식 습득 시스템의 품질 검증 로직 설계

---
id: SPEC-KNOWLEDGE-001
title: 지식 습득 시스템의 품질 검증 로직 설계
version: 0.1.0-draft
status: draft
priority: HIGH
---

## Purpose
외부 지식(URL, 문서 등)을 시스템에 반입할 때 발생하는 정보 손실과 왜곡을 원천 차단하고, 무손실 수집 증명 및 시맨틱 중복 제거의 100% 논리적 일치를 알고리즘으로 판별하여 최고 품질의 데이터만 섭취하도록 보장합니다.

## Background
기존 시스템은 외부 데이터를 파싱할 때 블랙박스 형태로 병합(Merge)을 수행하여 미세한 컨텍스트 유실(1%의 뉘앙스 차이 등)을 탐지하지 못할 위험이 존재했습니다. 통합 무결성 프로토콜(`multi_domain_zero_loss_ingestion_protocol`)이 발효됨에 따라, 지식 섭취의 모든 중간 과정을 증명하고 5대 멀티 도메인 렌즈를 거치는 강력한 검증 게이트웨이가 필요해졌습니다.

## Requirements

### Ubiquitous
- 시스템은 외부 지식 추출 직후, 반드시 터미널에 '총 추출 라인 수(Line Count)'와 '원본의 처음 3줄 및 마지막 3줄'을 출력하여 압축이 없었음을 자가 증명해야 한다.
- 시스템은 외부 지식 습득 시 내부 지식(Truth)과의 정합성, 출처 신뢰도, 정보 밀도를 기반으로 A등급(무조건 수용)부터 F등급(즉시 폐기)까지 등급을 매기는 판별 알고리즘을 수행해야 한다.

### Event-Driven
- WHEN 데이터 병합(Semantic Deduplication) 이벤트가 발생할 때, THEN 시스템은 두 데이터 간의 유사도가 정확히 1.0(100% 일치)인지 검사하고, 단 1%의 뉘앙스 차이라도 발생하면 병합을 거부하고 원본을 각각 보존해야 한다.
- WHEN 지식 수집이 완료되었을 때, THEN 시스템은 결론 도출을 유보하고 반드시 5가지 멀티 산업 렌즈(IT, 비즈니스, 교육, 예술, 실생활)를 통과하는 추론 과정을 거쳐야 한다.

### Unwanted
- IF 무손실 검증(라인 수 불일치, 텍스트 매칭 실패)이 실패할 경우, THEN 시스템은 즉각적인 패닉(Panic) 상태에 진입하여 프로세스를 중단하고 `.autopus/history/lessons_learned.md` 파일에 실패 사실과 훼손 지점을 물리적으로 로깅해야 한다.

### Optional
- WHERE 사용자가 5대 방향성 중 특정 항목의 심층 열람을 지시한 경우, THEN 시스템은 `expert_ideation_protocol`을 즉시 호출하여 압축 없는 논리 전개와 10x 리뷰 루프를 개시해야 한다.

## Acceptance Criteria
- [ ] 수집된 원본의 크기(Byte)와 추출된 객체의 크기가 정확히 일치하며 라인 수 검증에 통과한다.
- [ ] 코사인 유사도 0.99인 데이터 병합 시도 시, 병합이 거부되고 두 개의 독립된 지식 객체로 저장됨을 증명한다.

## Out of Scope
- 지식의 시각화(Graph Node UI 등 프론트엔드 표출 영역)

## Traceability
| Requirement | Test | Status |
|-------------|------|--------|
| 무손실 증명 | pkg/worker/knowledge/verify_lossless_test.go | pending |
| 시맨틱 일치 | pkg/worker/knowledge/verify_semantic_test.go | pending |
| 등급 판별 | pkg/worker/knowledge/score_grade_test.go | pending |