25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

## 🎭 **B1단계: 핵심 요약** 

유지 그룹: 

① 오디션 

② 학원 

- ③ 기술 

- ④ 작품 

⑤ 장비·도구 삭제 그룹: 

- ❌ 입시 (연기 입시, 연극영화과, 실기시험 등 전체 제거) 키워드 총 92개 

- 분석 단위는 lemma(대표어) 기준 

## 🎯 **B2단계: 세부 구조** 

## **① Group 1 — 오디션** 

|**① Group 1 — 오디션**||
|---|---|
|**대표 키워드**|**주요 변형(variants)**|
|오디션|오디숀, audition|
|영화 오디션|영화오디션, movie audition|
|드라마 오디션|드라마오디션, drama audition|
|뮤지컬 오디션|뮤지컬오디션, musical audition|
|연극 오디션|연극오디션, theater audition|
|오디션 준비|오디션준비, audition ready|
|오디션 팁|오디션팁, audition tip|
|카메라 테스트|camera test, 테스트 촬영|



## **② Group 2 — 학원** 

**대표 키워드 주요 변형(variants)** 연기학원 연기 학원, 배우학원 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

18/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

**대표 키워드 주요 변형(variants)** 연기레슨 연기 레슨, acting lesson 기초반 기초 연기반, 초급반, beginner class 심화반 심화 연기반, advanced class 원데이 클래스 1일 클래스, 하루 클래스 카메라 연기 카메라연기, camera acting 발성 발성훈련, vocal training 호흡 호흡연습, breathing control 

## **③ Group 3 — 기술** 

**대표 키워드 주요 변형(variants)** 감정이입 감정 이입, 감정 몰입, empathy acting 표정연기 표정 연기, facial acting 대사 암기 대사암기, script memorize 즉흥연기 즉흥 연기, improv acting 리딩 reading, 대본리딩 모노로그 monologue, 독백연기 워크샵 워크샵, 워크숍, workshop 씬스터디 씬 스터디, scene study 

## **④ Group 4 — 작품** 

**대표 키워드 주요 변형(variants)** 단편영화 단편 영화, short film 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

19/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

**대표 키워드 주요 변형(variants)** 웹드라마 web drama, 웹 드라마 연극 theater, stage play 뮤지컬 musical, 뮤지컬공연 단편 출연 short casting, 단편 출연 프로필영상 profile video, 프로필 영상 리딩영상 reading video, 리딩 영상 오디션영상 audition video, 오디션 영상 

## **⑤ Group 5 — 장비·도구** 

**대표 키워드 주요 변형(variants)** 프로필사진 프로필 사진, profile photo, 프로필촬영 조명 lighting, 조명세팅 배경천 백드롭, background cloth 대본 script, 시나리오, 대본인쇄 마이크 microphone, 무선마이크 카메라 camera, 미러리스카메라 삼각대 tripod, 스탠드 녹음 녹음기, sound record, 보이스레코더 

## ⚙ **표기변형 규칙팩 (정규식 요약)** 

python 

코드 복사 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

20/46 

25. 11. 30. 오후 1:50 

 

Start - 시스템 부팅 보고 

```
rules = [
```

```
  (r'([가-힣])\s+([가-힣])', r'\1\2'),
  (r'오디션|오디숀|audition', '오디션'),
  (r'프로필\s?(사진|촬영)|profile\s?(photo|shoot)', '프로필사진'),
  (r'연기\s?학원', '연기학원'),
  (r'카메라\s?연기|camera\s?acting', '카메라연기'),
  (r'워크\s?샵|워크\s?숍', '워크샵'),
  (r'[^\w가-힣 ]', '')
]
```

## 📁 **CSV 예시 (Non-입시 Edition)** 

csv 

코드 복사 

```
group,lemma,variants
오디션,오디션,"오디숀,audition"
오디션,영화오디션,"영화오디션,movie audition"
```

```
학원,연기학원,"연기학원,배우학원"
```

```
기술,감정이입,"감정이입,감정몰입,empathy acting"
```

```
작품,웹드라마,"web drama,웹드라마"
```

```
장비,프로필사진,"프로필사진,profile photo,프로필촬영"
```

## 📡 **다음 단계 선택** 

## 1️⃣ **“승인”** → 

→ **ActorTrend v0.4 (API 스펙 + DB DDL + 수집기 기본 코드)** 생성 

(즉, 이걸 실제 작동 가능한 형태로 만들기 위한 코어 코드 구조) 

## 2️⃣ **“확장”** → 

→ **SNS 감정·톤 분석 모듈** 추가 (예: #오디션준비 감정어 자동 분류·시각화) 

형, 어느 쪽으로 갈까? 

**==> picture [489 x 38] intentionally omitted <==**

**==> picture [49 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
확장<br>**----- End of picture text -----**<br>


https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

21/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

5s 동안 생각함 

## 형, **ActorTrend v0.3a → SNS 감정·톤 분석 모듈(Emotion/Tone v1)** 붙인다. 

핵심은 “원문 저장 없이, 공개 텍스트의 **요약·집계 결과만** 보관”하는 하이브리드 레이블링 이야. 

## **A단계: 전체 개요** 

목표: 인스타/블로그/유튜브의 공개 메타텍스트(제목·요약·해시태그)를 **감정·의도·톤** 으로 분류해 월별 추세, 급상승 맥락을 보여줌. 저장은 **집계치만** (개인정보·본문 미저장). 

## **B1단계: 핵심** 

## **라벨 체계(12클래스)** 

감정: 기쁨, 기대, 긴장, 좌절, 분노, 중립 

의도/상황: 준비, 도전, 합격/성과, 정보공유, 모집/공고, 후기 

## **모델 전략(하이브리드)** 

**1. 룰·사전(Seed Lexicon)** : 한글 감정/의도 키워드·이모지·해시태그 기반 약한 신호 생성 

**2. 임베딩 기반 약지도** : 한국어 문장임베딩 → 로지스틱/라이트GBM으로 멀티라벨 

**3. 제로샷 백업** : NLI 계열(ko 다국어)로 저빈도 클래스 보정 

- **저장 원칙** : 텍스트 폐기, **월×키워드×출처×라벨 빈도** 만 저장 

## **B2단계: 세부** 

## **1) 파이프라인** 

**1. 수집(기존)** : 해시태그/제목/요약, 업로드일, 출처 

**2. 정제** : 이모지→토큰, 해시태그 분리, 표기 정규화(기존 rules 재사용) 

## **3. 라벨러(순서 적용)** 

RuleScore(사전·이모지) → EmbedScore(모델) → ZeroShotTieBreak 멀티라벨 허용(감정 1 + 의도 1 이상) 

## **4. 집계/스코어** 

`label_count` , `label_ratio` , `month_over_month Δ` , `surge` 키워드군/출처별 가중 합산 

## **5. 노출** 

Dashboard: “Top Rising + 감정 분포” 링차트, “톤 흐름(3개월)” Compare: 키워드별 감정 구성비/의도비 Explore: 감정×의도 버블(월축) 

## **2) 라벨 사전(요약 샘플)** 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

22/46 

25. 11. 30. 오후 1:50 Start - 시스템 부팅 보고 

 

- 기대: 기대, 설렘, 기다, 곧, 곧공개, 🔜, ✨, 🙌 

- 긴장: 떨림, 긴장, 초조, ㄷㄷ, 😬, 😰 

- 좌절: 어렵, 막힘, 탈락, 망했, 포기, 😞, 💢 

- 준비: 연습, 루틴, 독백연습, 발성훈련, 워밍업, #연습중 

- 도전: 도전, 도전기, 첫도전, 트라이, #challenge 

- 합격/성과: 합격, 캐스팅, 콜백, 섭외, 축하, 🏆, ✅ 정보공유: 꿀팁, 팁, 방법, 가이드, 체크리스트, Q&A 

- 모집/공고: 모집, 캐스팅공고, 서류접수, 오디션공지, 마감임박 

- 후기: 후기, 리뷰, 소감, 리캡, 회고 

## **3) DB 스키마(추가)** 

sql 코드 복사 

```
-- 라벨정의
CREATETABLE labels (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE,          -- joy, expectancy, tension, setback, anger, neutral,
  type TEXT CHECK (type IN ('emotion','intent'))
);
```

```
-- 월별집계
CREATETABLE label_snapshots (
  keyword_id INTREFERENCES keywords(id),
  source_id  INTREFERENCES sources(id),
  yyyymm     CHAR(7),        -- YYYY-MM
  label_id   INTREFERENCES labels(id),
  count      INT,
  ratio      NUMERIC(5,4),   -- 0~1
  surge      NUMERIC(6,3),   -- MoM or 3m/9m
PRIMARY KEY(keyword_id, source_id, yyyymm, label_id)
);
```

  

## **4) API 스펙(추가)** 

- `GET /sentiment?kw=오디션영상&months=12` → 월별 감정·의도 분포 

- `GET /tones?group=오디션&months=12` → 그룹 상위 키워드의 톤 비교 

- `GET /cooccurrence?kw=카메라연기` → 함께 등장한 감정·의도 Top N 

- `GET /alerts/label?label=성공&threshold=surge>1.5` → 라벨 급등 알림 

## **5) 대시보드 위젯** 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

23/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

- **Emotion Ring** : 감정 비중 링차트(최근월) + 3개월 Δ 
