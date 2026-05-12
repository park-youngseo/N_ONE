- **Intent Stream** : 의도 스택면적(12개월) 

- **Rising Context** : 급등 키워드 옆에 “주요 라벨 3종 뱃지” 

- **Label Map** : (감정 × 의도) 그리드 히트맵 

## **6) 품질·가드레일** 

텍스트는 **실시간 처리 후 즉시 폐기** (메모리 내), 집계만 저장 라벨 신뢰점수: RuleScore vs EmbedScore 일치도 

수동 검수 큐: 급등 라벨 샘플 20개만 임시 뷰(가려진 형태) → 확인 후 폐기 

## **7) 라이트 코드 골격(요약)** 

python 코드 복사 _`# pseudo:`_ `라벨러 from collections import Counter import re` 

`EMOJI_MAP = {"` ✨ `":"기대","` 🙌 `":"기대","` 😬 `":"긴장","` 😰 `":"긴장","` 😞 `":"좌절","` 🏆 `":"합 LEXICON = {` 

```
"기대":["기대","설렘","곧공개","곧","기다"],
```

```
"긴장":["긴장","떨림","초조","ㄷㄷ"],
"좌절":["어렵","막힘","탈락","포기","망했"],
```

```
"준비":["연습","루틴","독백연습","발성훈련","워밍업","#연습중"],
```

```
"도전":["도전","첫도전","트라이","#challenge"],
```

`"합격/성과":["합격","캐스팅","콜백","섭외","축하","` 🏆 `","` ✅ `"],` 

```
"정보공유":["꿀팁","방법","가이드","체크리스트","Q&A"],
```

```
"모집/공고":["모집","캐스팅공고","서류접수","오디션공지","마감임박"],
```

```
"후기":["후기","리뷰","소감","리캡","회고"]
}
```

```
defrule_labels(text:str):
```

```
  t = text.lower()
  hits = Counter()
for k, words in LEXICON.items():
for w in words:
if w.lower() in t:
        hits[k]+=1
for ch in text:
if ch in EMOJI_MAP:
      hits[EMOJI_MAP[ch]]+=1
return hits
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

24/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
deffuse(rule_hits, embed_probs, zero_shot=None):
```

```
# 가중결합(예: rule 0.6, embed 0.4, 제로샷타이브레이크)
  scores = Counter(rule_hits)
for k,v in (embed_probs or {}).items():
    scores[k] += 0.4*v
# 상위감정 1 + 의도 1 선택
  emotion = max([("기쁨",scores["기대"]),("긴장",scores["긴장"]),
                 ("좌절",scores["좌절"]),("분노",scores["분노"]),("중립",scores["중
  intent  = max([("준비",scores["준비"]),("도전",scores["도전"]),("합격/성과",scores
                 ("정보공유",scores["정보공유"]),("모집/공고",scores["모집/공고"]),(
return {"emotion":emotion, "intent":intent, "score":dict(scores)}
```

## **8) 알림 룰(예)** 

`surge(label='모집/공고') > 1.8` AND `source ∈ {instagram, youtube}` → **신규 오** 

## **디션 이슈 감지** 

> `ratio(label='준비') ↑ AND ratio(label='긴장') ↑` → **준비 수요 증가(초급자 컨텐츠** 

**기회)** 

`ratio(label='합격/성과') ↑` → **합격 후기 확산(브랜딩·리포스트 타이밍)** 

원하면 **v0.4(API+DDL+수집기 코드 프레임)**까지 한 번에 뽑아줄게. 

명령: **“승인”** → 실행 스펙+DDL+파이썬 베이스 코드 세트 생성. 

**==> picture [489 x 39] intentionally omitted <==**

승인 

형, 약속한 **ActorTrend v0.4 — API 스펙 + DB DDL + 수집기(파이썬 베이스)** 세트 바로 낸다. 

(입시 그룹은 이미 제거된 **Non-입시 Edition** 기준) 

## **A단계: 한 줄 개요** 

지난 12개월 동안 인스타/유튜브/네이버 블로그 **메타 데이터만** 수집→정규화→지표화(표 준화·급상승·계절성 + 감정/의도 라벨)→대시보드 API 제공. 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

25/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

## **B1단계: 핵심 요약** 

- **API** : `/trending` , `/compare` , `/explore` , `/sentiment` , `/tones` , `/alerts` 

- **DB** : Postgres (키워드/소스/스냅샷/점수/라벨/라벨집계) 

**Collector** : Python(Async) + 공식 API 우선(YouTube, Naver), 인스타는 해시태그 메타 (본문·개인정보 미수집) 

**NLP** : 한글 표기 정규화 + 동의어군 병합, 감정/의도 라벨러(룰+임베딩) 

- **스케줄** : 매일 증분(UTC 00:30 ≈ KST 09:30 이전 완료), 주 1회 재정규화 

## **B2단계: 실행 청사진** 

## **1) API 스펙 (FastAPI)** 

## **공통** 

Timezone: 응답 헤더 `X-Timezone: Asia/Seoul` 

- 기간 파라미터: `months` (기본 12), `from=YYYY-MM` , `to=YYYY-MM` 

## **① 트렌드 랭킹** 

```
GET /trending?group=오디션&months=12&limit=20
```

json 코드 복사 

`{ "meta": {"group": "오디션", "months": 12, "tz": "Asia/Seoul"}, "items": [ {"keyword":"오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour"keyword":"오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour:"오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour"오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour,"norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour"norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour:87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour,"surge":1.63,"seasonal":0.12,"trust":"A","sour"surge":1.63,"seasonal":0.12,"trust":"A","sour:1.63,"seasonal":0.12,"trust":"A","sour1.63,"seasonal":0.12,"trust":"A","sour,"seasonal":0.12,"trust":"A","sour"seasonal":0.12,"trust":"A","sour:0.12,"trust":"A","sour0.12,"trust":"A","sour,"trust":"A","sour"trust":"A","sour:"A","sour"A","sour,"sour"sour {"keyword":"카메라 테스트","norm":54.1,"surge":1.21,"seasonal":0.03,"trust":"B ] }`   

```
    {"keyword":"오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour"keyword":"오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour:"오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour"오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour오디션","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour","norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour,"norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour"norm":87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour:87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour87.2,"surge":1.63,"seasonal":0.12,"trust":"A","sour,"surge":1.63,"seasonal":0.12,"trust":"A","sour"surge":1.63,"seasonal":0.12,"trust":"A","sour:1.63,"seasonal":0.12,"trust":"A","sour1.63,"seasonal":0.12,"trust":"A","sour,"seasonal":0.12,"trust":"A","sour"seasonal":0.12,"trust":"A","sour:0.12,"trust":"A","sour0.12,"trust":"A","sour,"trust":"A","sour"trust":"A","sour:"A","sour"A","sour,"sour"sour
```

## **② 비교 시계열** 

```
GET /compare?kw=연기학원,오디션,프로필사진&months=12
```

json 코드 복사 

```
{
"series":[
    {"keyword":"연기학원","points":[{"yyyymm":"2024-11","norm":31.2},...]}
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

26/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
  ]
```

```
}
```

## **③ 탐색(연관/클러스터)** 

```
GET /explore?seed=카메라연기&limit=30
```

json 코드 복사 

```
{"seed":"카메라연기","neighbors":[{"keyword":"테스트촬영","similarity":0.82},...]}
```

  

## **④ 감정/의도 분포** 

```
GET /sentiment?kw=오디션영상&months=12
```

json 코드 복사 

```
{
```

`"keyword":"오디션영상", "by_month":[ {"yyyymm":"2025-03","emotion":{"기쁨":0.22,"긴장":0.18,"좌절":0.07,"중립":0.53} "intent":{"준비":0.31,"도전":0.19,"합격/성과":0.08,"정보공유":0.24,"모집/공고": ] }`   

## **⑤ 그룹 톤 비교** 

```
GET /tones?group=학원&months=12
```

json 코드 복사 

```
{"group":"학원","top":[{"keyword":"연기학원","emotion_top":"중립","intent_top":"정보
```

  

## **⑥ 라벨 급등 알림 후보** 

```
GET /alerts?label=모집/공고&threshold=1.5&since=2025-06
```

json 코드 복사 

```
{"label":"모집/공고","hits":[{"keyword":"오디션","yyyymm":"2025-09","surge":1.84,"s
```

  

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

27/46 

25. 11. 30. 오후 1:50 

 

Start - 시스템 부팅 보고 

## **2) DB DDL (Postgres)** 

sql 

코드 복사 

```
-- 그룹(입시제외)
CREATETABLEgroups (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUENOTNULL-- 오디션, 학원, 기술, 작품, 장비
);
CREATETABLE keywords (
  id SERIAL PRIMARY KEY,
  lemma TEXT UNIQUENOTNULL,        -- 분석단위(대표어)
  display TEXT NOTNULL,             -- 노출명
  group_id INTREFERENCESgroups(id),
  variants JSONB DEFAULT'[]'::jsonb,
  created_at TIMESTAMPTZ DEFAULT now()
);
```

```
CREATETABLE sources (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUENOTNULL,         -- instagram, youtube, naver_blog, naver_cafe
  type TEXT NOTNULL-- social, video, blog, community
);
-- 월별원시집계
CREATETABLE snapshots (
  keyword_id INTREFERENCES keywords(id),
  source_id  INTREFERENCES sources(id),
  yyyymm     CHAR(7) NOTNULL,       -- YYYY-MM
  count      INTNOTNULL,