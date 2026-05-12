- [추론] GPT 기반으로 동일한 논리 구조를 재현할 수 있으며, 

- [가정] 형의 N·ONE 시스템에 맞춰 **“맞춤형 트렌드 분석기”**로 통합 가능하다. 

형, 이걸 **“실제 작동 가능한 형태 (노션 + 엑셀 + GPT 통합 버전)”**으로 만들어줄까? 예를 들어 형이 키워드만 입력하면 자동으로 트렌드 리포트가 완성되는 구조로. 

**==> picture [489 x 38] intentionally omitted <==**

해줘 

10s 동안 생각함 

몇 초 동안 생각함 

형, 완성했어. 

**N·ONE Insight Engine Starter (Sometrend-style) v1.0** 엑셀 킷을 만들어서 바로 쓸 수 있게 했다. 

Download the Excel starter 

## **A단계: 전체 개요** 

- [사실] 이 파일은 **수동 수집 + 자동 계산** 구조야. Google Trends / Naver Data Lab / 

- (선택) 썸트렌드에서 뽑은 수치를 붙여 넣으면, 지표·문장 요약이 자동으로 형성돼. 

- [추론] “연기학원”처럼 저빈도 키워드는 **변형 키워드 + 지역 키워드** 를 함께 넣을 때 분석이 안정적이야. 

## **B1: 사용 순서(핵심)** 

## **1. 01_INPUT** 

`keyword_main` 에 핵심 키워드 입력(예: 연기학원). 

`keyword_variants` 에 변형(배우학원; 연기 레슨; 오디션 학원; 서울 연기학원). 

## **2. 02_RAW_PASTE** 

- 외부 도구 결과를 붙여넣기(열 고정). 포맷: 

Google Trends → `date, interest` 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f04947-1b4c-8324-acfd-3e892e04a4d2 

12/13 

Start - 00_홍보[로드 맵] 

25. 11. 30. 오후 1:49 

 

Naver Data Lab → `date, index` 

Sometrend/기타 → `date, mentions, pos, neu, neg, related_terms` 

채널/출처는 `source / channel` 열에 메모. 

## **3. 03_KPI** 

자동 계산: 총 언급량, 일평균, MoM, 긍·부정률, 평균 관심도/지수. 

## **4. 04_RELATED_TERMS** 

연관어 Top N을 정리(예: 입시, 오디션, 수업료…). 보고서 문장에 자동 반영됨. 

## **5. 05_REPORT_Auto** 

날짜(KST) 기준 보고서 제목/핵심요약/인사이트 문장이 **자동 생성** 됨 → 복사해서 카톡 공지·노션에 사용. 

- 확장(옵션) — 

   - “배우학원 / 연기 레슨” 등 **2~3개 키워드를 병렬 운영** 하려면 `01_INPUT` 을 복수 행으 로 확장하고, `03_KPI` / `05_REPORT_Auto` 블록을 복제해 각각 연결 가능. 

원하면 **시각화(PPT/PNG) 자동 생성 모듈** 도 붙일 수 있어(그래프 2~3종). 

원하는 채널(예: 네이버 블로그만, 인스타만 등)과 기간(KST 절대 날짜) 알려주면, 샘플 데 이터까지 채워 넣은 **완성 리포트 버전** 으로 바로 뽑아줄게. 

**==> picture [489 x 66] intentionally omitted <==**

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f04947-1b4c-8324-acfd-3e892e04a4d2 

13/13 
