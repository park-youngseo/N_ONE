YOUTUBE_API_KEY = os.getenv("YOUTUBE_API_KEY",)
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

32/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
""
NAVER_CLIENT_ID = os.getenv("NAVER_CLIENT_ID",)
""
NAVER_CLIENT_SECRET = os.getenv("NAVER_CLIENT_SECRET",)
DATABASE_URL = os.getenv("DATABASE_URL")
```

## **collector/nlp_normalize.py (요약)** 

python 코드 복사 `import re, json NORMAL_RULES = [` 

```
  (re.compile(r'([가-힣])\s+([가-힣])'), r'\1\2'),
  (re.compile(r'오디션|오디숀|audition', re.I), '오디션'),
  (re.compile(r'프로필\s?(사진|촬영)|profile\s?(photo|shoot)', re.I), '프로필사진'),
  (re.compile(r'연기\s?학원'), '연기학원'),
```

```
  (re.compile(r'카메라\s?연기|camera\s?acting', re.I), '카메라연기'),
```

```
  (re.compile(r'워크\s?샵|워크\s?숍'), '워크샵'),
  (re.compile(r'[^\w가-힣 ]'), '')
]
defnormalize(text:str)->str:
  t = text.strip()
for p, repl in NORMAL_RULES:
    t = p.sub(repl, t)
return re.sub(r'\s+',' ', t).strip()
```

  

## **collector/labeler.py (요약)** 

python 코드 복사 

```
from collections import Counter
```

`EMOJI_MAP = {"` ✨ `":"기대","` 🙌 `":"기대","` 😬 `":"긴장","` 😰 `":"긴장","` 😞 `":"좌절","` 🏆 `":"합 LEXICON = {` 

`"기쁨":["기쁨","행복","신남","뿌듯"], "기대":["기대","설렘","곧","곧공개"], "긴장":["긴장","떨림","초조","ㄷㄷ"], "좌절":["어렵","막힘","탈락","포기","망했"], "중립":["공지","업로드","업데이트"], "준비":["연습","루틴","독백연습","발성훈련","워밍업","#연습중"], "도전":["도전","첫도전","트라이","#challenge"], "합격/성과":["합격","캐스팅","콜백","섭외","축하","` 🏆 `","` ✅ `"], "정보공유":["꿀팁","방법","가이드","체크리스트","Q&A"], "모집/공고":["모집","캐스팅공고","서류접수","오디션공지","마감임박"],` 

```
"후기":["후기","리뷰","소감","리캡","회고"]
```

```
}
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

33/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
deflabel_text(text:str):
  t=text.lower()
  c=Counter()
for k,ws in LEXICON.items():
for w in ws:
if w.lower() in t: c[k]+=1
for ch in text:
if ch in EMOJI_MAP: c[EMOJI_MAP[ch]]+=1
```

```
# 감정(기쁨/기대/긴장/좌절/중립) 중최댓값, 의도(준비/도전/합격/정보공유/모집/후기) 중
  emo = max([("기쁨",c["기쁨"]),("기대",c["기대"]),("긴장",c["긴장"]),("좌절",c["좌절
  intent = max([("준비",c["준비"]),("도전",c["도전"]),("합격/성과",c["합격/성과"]),(
return emo, intent
```

## **collector/youtube_search.py (공식 API 사용)** 

python 코드 복사 

```
import httpx, pandas as pd, datetime as dt
from .config import YOUTUBE_API_KEY
from .nlp_normalize import normalize
```

```
BASE = "https://www.googleapis.com/youtube/v3/search"
```

```
deffetch_youtube_counts(keyword:str, published_after:str):
```

```
  params = {"part":"snippet","q":keyword,"maxResults":50,"type":"video","key":YOUT
  items=[]
```

`with httpx.Client(timeout=20) as client: r=client.get(BASE, params=params) r.raise_for_status()`   `data=r.json() for it in data.get("items",[]): title = normalize(it["snippet"]["title"]) date = it["snippet"]["publishedAt"][:7]` _`# YYYY-MM`_ `items.append({"yyyymm":date, "title":title})` _`#`_ `월별 건수 집계 df=pd.DataFrame(items) if df.empty: return {} return df.groupby("yyyymm").size().to_dict()`   **collector/naver_blog_search.py (Open API)** python 코드 복사 

## **collector/naver_blog_search.py (Open API)** 

python 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

34/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
import httpx, hashlib, datetime as dt
from .config import NAVER_CLIENT_ID, NAVER_CLIENT_SECRET
from .nlp_normalize import normalize
```

```
BASE = "https://openapi.naver.com/v1/search/blog.json"
deffetch_blog_counts(keyword:str, display:int=50):
  headers={"X-Naver-Client-Id":NAVER_CLIENT_ID,"X-Naver-Client-Secret":NAVER_CLIEN
  params={"query":keyword,"display":display,"sort":"date"}
with httpx.Client(timeout=20, headers=headers) as client:
    r=client.get(BASE, params=params)
    r.raise_for_status()
    data=r.json()
    items=data.get("items",[])
    mm={}
for it in items:
      date=it.get("postdate","")  # YYYYMMDD
      yyyymm=f"{date[:4]}-{date[4:6]}" iflen(date)==8elseNone
if yyyymm: mm[i:=yyyymm]=mm.get(i,0)+1
return mm
```

  

## **collector/instagram_hashtag.py (메타만)** 

python 코드 복사 

- _`#`_ `인스타그램` _`:`_ `공개 해시태그 메타` _`(`_ `게시수` _`)`_ `만 수집` _`.`_ `본문` _`/`_ `이미지` _`/`_ `댓글 미수집` _`.`_ 

- _`#`_ `구현은 실제 운영환경에서 해시태그 검색 결과의` _`'`_ `게시물 수` _`'`_ `메타를 확인 가능한 범위 내에` 

- _`#`_ `차단 방지를 위해 레이트리밋` _`/`_ `백오프 필수` _`.`_ `약관 준수` _`.`_ 

```
deffetch_instagram_hashtag_counts(hashtag:str)->dict:
```

- _`#`_ `실제 구현은 환경에 맞게 구성` _`.`_ `여기서는 인터페이스만 정의` _`.`_ 

- _`# return {"2025-07": 120, "2025-08": 141, ...}`_ 

```
return {}
```

  

## **collector/run_daily.py (증분 수집→DB 적재)** 

python 코드 복사 

```
import os, datetime as dt
from dateutil.relativedelta import relativedelta
from sqlalchemy import create_engine, text
from .youtube_search import fetch_youtube_counts
from .naver_blog_search import fetch_blog_counts
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

35/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
from .instagram_hashtag.py import fetch_instagram_hashtag_counts
```

```
ENGINE = create_engine(os.getenv("DATABASE_URL"))
```

```
defupsert_snapshot(keyword_id:int, source_id:int, mm_counts:dict):
with ENGINE.begin() as con:
for yyyymm, cnt in mm_counts.items():
      con.execute(text("""
        INSERT INTO snapshots(keyword_id, source_id, yyyymm, count)
        VALUES(:kid,:sid,:mm,:cnt)
        ON CONFLICT (keyword_id, source_id, yyyymm)
        DO UPDATE SET count=EXCLUDED.count;
      """), {"kid":keyword_id,"sid":source_id,"mm":yyyymm,"cnt":int(cnt)})
```

```
defrun(keyword:str, keyword_id:int):
```

```
  since = (dt.date.today().replace(day=1) - relativedelta(months=12)).strftime("%Y
# youtube
  y=fetch_youtube_counts(keyword, since)
  upsert_snapshot(keyword_id, source_id=2, mm_counts=y)  # 예: 2=youtube
# naver blog
```

```
  b=fetch_blog_counts(keyword)
  upsert_snapshot(keyword_id, source_id=3, mm_counts=b)  # 3=naver_blog
# instagram (옵션)
```

```
  i=fetch_instagram_hashtag_counts(keyword.replace(" ",""))
  upsert_snapshot(keyword_id, source_id=1, mm_counts=i)  # 1=instagram
```

```
if __name__=="__main__":
# TODO: keywords 테이블에서대상키워드읽어순회
pass
```

## **scripts/cron_example.sh** 

bash 코드 복사 `#!/usr/bin/env bash`  `export TZ=Asia/Seoul`  `python -m collector.run_daily` 

## **지표 산출 로직(요약)** 

**표준화(norm 0~100)** : 같은 월 내 `keyword 합산 count` 를 min-max 정규화 

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

36/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

- **급상승(surge)** : `(최근3개월 평균) / (이전9개월 평균)` × 출처 신뢰 가중 

- **계절성(seasonal)** : monthly series → STL 분해 → seasonal 성분 평균 

- **신뢰도(trust)** : 출처 다양성(≥2) + 표본수(분위수 기준) + 최근성(최근 60일 이내 관 측) → A/B/C 

## **운영 가드레일** 

- **본문/이미지/개인정보 저장 금지** , 메타 집계만 저장 

- robots.txt·약관 준수, 공식 API 우선 

- 레이트리밋/백오프/실패 큐(3회 재시도) 

- 로그: 수집시각·출처·요청량·에러코드 기록 

## **다음 액션** 

   - 원하면 내가 **seed CSV(Non-입시 92키워드) 파일 본문** 을 바로 드롭해줄게. 또는 **Docker용 .env 예시** 와 함께 “로컬 부팅 가이드(3단계)”까지 내줄게. 

- 명령: **“다음”** (seed CSV 드롭) / **“부팅”** (로컬 실행 가이드) / **“압축”** (전체 코드 번들 요약) 

**==> picture [489 x 40] intentionally omitted <==**

우선 만들기전에 이 트랜드를 마케팅에 어떻게 활용하는 거지? 