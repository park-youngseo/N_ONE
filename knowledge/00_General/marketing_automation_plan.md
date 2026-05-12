# Implementation Plan: Marketing Automation Solution (Project NONE)

## 1. 개요
`knowledge` 폴더의 178개 마케팅 자산을 활용하여 인스타그램과 블로그로의 유입을 자동화하는 시스템을 구축합니다. 2026 마케팅 트렌드인 '페르소나 중심의 전문성'을 반영하여 AI 기반의 콘텐츠를 생성하고 배포합니다.

## 2. 핵심 모듈 설계

### [Module 1] Knowledge-RAG Processor (T1)
- **기능**: `knowledge/` 폴더의 대규모 마케팅 데이터를 벡터화(추론용)하고 핵심 인사이트를 추출합니다.
- **파일**: [NEW] `knowledge_processor.py`

### [Module 2] Persona Content Engine (T2)
- **기능**: 추출된 지식을 바탕으로 인스타그램 릴스 대본 및 블로그 포스팅을 자동 생성합니다.
- **특징**: 70/30 하이브리드 모델(AI 기초 + 인간 인사이트) 적용.
- **파일**: [NEW] `content_generator.py`

### [Module 3] Multi-Channel Auto-Distributor (T3)
- **기능**: 인스타그램 Graph API 및 블로그 API를 연동하여 4-1-1 법칙 기반의 스케줄링 게시를 수행합니다.
- **파일**: [NEW] `distributor_engine.py`

### [Module 4] Engagement & Tracking Bot (T4)
- **기능**: 타겟 해시태그 기반 자동 소통 및 유입 URL 성과 분석을 담당합니다.
- **파일**: [NEW] `engagement_bot.py`

## 3. 사용자 리뷰 필요 사항 (User Review Required)
> [!IMPORTANT]
> - 인스타그램/블로그 API 연동을 위해 **Access Token** 발급이 필요합니다.
> - 자동 포스팅 시 **최종 승인 단계**를 넣을지, 100% 무인으로 가동할지 결정이 필요합니다.

## 4. 검증 계획 (Verification Plan)
- **단위 테스트**: 각 모듈의 API 연동 및 콘텐츠 생성 퀄리티 검증.
- **E2E 테스트**: 지식 추출부터 실제 포스팅 예약까지의 전체 프로세스 확인.
- **지표 확인**: 자동 포스팅 후 유입 클릭 수(CTR) 측정.