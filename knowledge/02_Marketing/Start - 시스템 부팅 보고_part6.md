  sample_n   INTDEFAULT0,
PRIMARY KEY(keyword_id, source_id, yyyymm)
);
```

```
-- 스코어(표준화, 급상승, 계절성, 신뢰)
CREATETABLE scores (
  keyword_id INTREFERENCES keywords(id),
  yyyymm     CHAR(7) NOTNULL,
  norm       NUMERIC(6,2),           -- 0~100
  surge      NUMERIC(6,3),           -- 3m/9m 비율가중
  seasonal   NUMERIC(6,3),
  trust_grade CHAR(1) CHECK (trust_grade IN ('A','B','C')),
PRIMARY KEY(keyword_id, yyyymm)
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

28/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
);
```

```
-- 라벨(감정/의도)
CREATETABLE labels (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUENOTNULL,         -- 기쁨, 기대, 긴장, 좌절, 분노, 중립, 준비, 도
  type TEXT CHECK (type IN ('emotion','intent')) NOTNULL
);
```

```
-- 월별라벨집계(텍스트미저장)
CREATETABLE label_snapshots (
  keyword_id INTREFERENCES keywords(id),
  source_id  INTREFERENCES sources(id),
  yyyymm     CHAR(7) NOTNULL,
  label_id   INTREFERENCES labels(id),
  count      INTNOTNULL,
  ratio      NUMERIC(5,4) NOTNULL,  -- 0~1
  surge      NUMERIC(6,3),
PRIMARY KEY(keyword_id, source_id, yyyymm, label_id)
);
-- 인덱스
CREATE INDEX idx_snapshots_month ON snapshots(yyyymm);
CREATE INDEX idx_scores_norm ON scores(yyyymm, norm DESC);
CREATE INDEX idx_label_snapshots_month ON label_snapshots(yyyymm);
```

## **3) 코드 골격 (디렉터리 & 필수 파일)** 

arduino 

코드 복사 

```
actortrend/
├─ api/
│  ├─ main.py
│  ├─ deps.py
```

```
│  ├─ routers/
│  │  ├─ trending.py
│  │  ├─ compare.py
│  │  ├─ explore.py
│  │  ├─ sentiment.py
│  │  └─ alerts.py
│  └─ models.py
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

29/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
├─ collector/
```

```
│  ├─ run_daily.py
```

```
│  ├─ instagram_hashtag.py
```

- `│  ├─ youtube_search.py` 

```
│  ├─ naver_blog_search.py
```

```
│  ├─ nlp_normalize.py
```

```
│  ├─ labeler.py
│  └─ config.py
├─ seed/
│  └─ keywords_non_admission.csv
├─ scripts/
│  ├─ init_db.sql
│  └─ cron_example.sh
├─ docker-compose.yml
└─ requirements.txt
```

## **requirements.txt** 

makefile 

코드 복사 

```
fastapi==0.115.0
uvicorn[standard]==0.30.6
SQLAlchemy==2.0.36
psycopg2-binary==2.9.9
pydantic==2.9.2
httpx==0.27.2
python-dateutil==2.9.0
pandas==2.2.2
kiwipiepy==0.17.1
scikit-learn==1.5.2
sentence-transformers==3.0.1
statsmodels==0.14.2
```

## **docker-compose.yml (요약)** 

yaml 

코드 복사 

```
version:"3.9"
services:
db:
image:postgres:16
environment:
POSTGRES_DB:actortrend
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

30/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
POSTGRES_USER:actor
POSTGRES_PASSWORD:secret
ports: ["5432:5432"]
volumes: ["pgdata:/var/lib/postgresql/data"]
api:
build:.
command:uvicornapi.main:app--host0.0.0.0--port8000
environment:
DATABASE_URL:postgresql+psycopg2://actor:secret@db:5432/actortrend
TZ:Asia/Seoul
depends_on: [db]
ports: ["8000:8000"]
volumes: { pgdata: {} }
```

## **api/main.py (핵심)** 

python 

코드 복사 

```
from fastapi import FastAPI
from .routers import trending, compare, explore, sentiment, alerts
```

```
app = FastAPI(title="ActorTrend API", version="0.4", root_path="")
app.include_router(trending.router)
app.include_router(compare.router)
app.include_router(explore.router)
app.include_router(sentiment.router)
app.include_router(alerts.router)
```

```
@app.get("/health")
defhealth():
return {"ok": True, "tz": "Asia/Seoul"}
```

## **api/models.py (SQLAlchemy 요약)** 

python 

코드 복사 

```
import os
```

```
from sqlalchemy import create_engine, text
from sqlalchemy.orm import sessionmaker
```

```
DATABASE_URL = os.getenv("DATABASE_URL")
engine = create_engine(DATABASE_URL, pool_pre_ping=True)
SessionLocal = sessionmaker(bind=engine, autocommit=False, autoflush=False)
```

https://chatgpt.com/g/g-p-68ef210f425481918a2bd6c6d79d6c3b/c/68f1a9cf-7b20-8322-9dd9-b2abb830ae8a 

31/46 

25. 11. 30. 오후 1:50 

Start - 시스템 부팅 보고 

 

```
deffetchall(q:str, **params):
with engine.connect() as conn:
return conn.execute(text(q), params).mappings().all()
```

## **api/routers/trending.py (예시)** 

python 코드 복사 

```
from fastapi import APIRouter, Query
from ..models import fetchall
```

```
router = APIRouter(prefix="/trending", tags=["trending"])
```

```
@router.get("")
```

```
defget_trending(group: str = Query(...), months: int = 12, limit: int = 20):
    q = """
    SELECT k.lemma AS keyword, s.norm, s.surge, s.seasonal, s.yyyymm,
           array_agg(DISTINCT so.name) AS sources,
           s.trust_grade AS trust
    FROM scores s
    JOIN keywords k ON k.id=s.keyword_id
    JOIN snapshots sn ON sn.keyword_id=k.id AND sn.yyyymm=s.yyyymm
    JOIN sources so ON so.id=sn.source_id
    JOIN groups g ON g.id=k.group_id
    WHERE g.name=:group
      AND s.yyyymm >= to_char((date_trunc('month', now()) - INTERVAL ':months mont
    GROUP BY k.lemma, s.norm, s.surge, s.seasonal, s.trust_grade, s.yyyymm
    ORDER BY s.norm DESC
    LIMIT :limit;
    """
    rows = fetchall(q, group=group, months=months, limit=limit)
# 최신월기준상위만묶어반환(간단화)
    latest = max([r["yyyymm"] for r in rows]) if rows elseNone
    items = [r for r in rows if r["yyyymm"]==latest]
return {"meta":{"group":group,"months":months}, "items":items}
```

  

## **collector/config.py** 

python 코드 복사 

```
import os
""