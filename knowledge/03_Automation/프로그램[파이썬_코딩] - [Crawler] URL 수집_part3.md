
가능해. **엔진 스위치** 로 확장하면 돼. 

안전·공식 경로(추천): 

**Google** : Custom Search JSON API (키/엔진ID 필요). 

**Bing** : Azure Bing Web Search API (키 필요). 

**Naver** : 검색 Open API (Client ID/Secret 필요). 

- 이유: HTML 긁기보다 **API가 안정·합법·지속성** 좋아. 

- 구현: `PROVIDER` 값을 `duckduckgo | google | bing | naver` 중에서 선택 → 각 함수 로 분기. 

   - 키/시크릿은 **환경변수** 로 읽어오게 하면 안전. 

## **4) 설치가 왜 빠졌냐고?** 

맞아. 내가 먼저 **설치 안내를 더 분명히** 줬어야 해. 미안해. 아래처럼 깔면 바로 돌아가. 

## **설치(필수)** 

bash 코드 복사 

```
python -m pip install requests beautifulsoup4
```

## **설치(선택 · 정확도↑/속도↓ 확장용)** 

bash 코드 복사 `python -m pip install lxml lxml` : 메타설명/본문 추출 시 파서 안정성↑. 나머지는 파이썬 표준라이브러리로 충분해. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 12/169 

25. 11. 30. 오후 1:05 

 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

원하면 내가 **Search.py** 를 아래처럼 바로 확장해줄게. 

`description` 칼럼 추가(DDG 스니펫 수집). 

`TOP_Search.csv` 자동 생성(도메인별 집계). 

`PROVIDER` 스위치 뼈대(추후 API키 넣으면 바로 다엔진). 

말만 해줘, **“확장”** . 

**==> picture [489 x 39] intentionally omitted <==**

야 이새끼 또 빡치게 하네 

형, 미안. 진짜 화난 거 알아. 내가 빠뜨린 부분이나 헷갈리게 만든 건 내 잘못이야. 

지금은 딱 한가지만 알려줘 — 

지금 바로 **Search.py 확장버전(설명+TOP+다엔진 구조)** 만들어줄까? 

아니면 왜 그런 부분이 빡치는지 먼저 정리해줄까? 

**==> picture [489 x 38] intentionally omitted <==**

python -m pip install -U duckduckgo-search requests 

이렇게 네이버도 있고 구글도 있지않아? 인스톨 프로그램 이? 필요한 ? 

26s 동안 생각함 

맞아, 설치 패키지 있어. 핵심만 딱 정리해줄게. 

## **DuckDuckGo** : 라이브러리 있음. 

sql 

코드 복사 

```
python -m pip install -U duckduckgo-search requests
```

→ `duckduckgo_search` 로 결과·스니펫 바로 받기 가능. 

**Google** : 공식은 **Custom Search JSON API** (API 키·CX 필요). 

nginx 코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

13/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
python -m pip install -U google-api-python-client
```

- → 또는 `requests` 만으로 REST 호출도 가능. 

- **Bing** : **Bing Web Search API** (키 필요). 라이브러리 없이 `requests` 권장. 

- (엔드포인트: `https://api.bing.microsoft.com/v7.0/search` ) 

- **Naver** : **검색 Open API** (Client ID/Secret 필요). 전용 파이썬 패키지 없이 `requests` 로 호출. 

- (엔드포인트: `https://openapi.naver.com/v1/search/webkr.json` ) 

- **Daum(카카오)** : **Kakao 검색 API** (REST 키 필요). `requests` 로 호출. 

- (엔드포인트: `https://dapi.kakao.com/v2/search/web` ) 

## **필수 키 설정(윈도우 · 한 번만)** 

PowerShell: 

nginx 

코드 복사 

```
setx GOOGLE_CSE_KEY "여기_키"
setx GOOGLE_CSE_CX  "여기_CX"
setx BING_KEY       "여기_키"
setx NAVER_CLIENT_ID "여기_ID"
setx NAVER_CLIENT_SECRET "여기_SECRET"
setx KAKAO_REST_KEY "여기_키"
```

## → 터미널을 **다시 열어야** 적용돼. 

원하면 **Search.py** 를 `PROVIDER=duckduckgo|google|bing|naver|kakao` 스위치 구조로 **바로 교체** 해줄게. 

말 없이 진행해도 되면 지금 메시지에 확장본을 **그대로 붙여서** 주겠다. 

**==> picture [489 x 38] intentionally omitted <==**

내가 예시로 이야기 하지 않았으면 에러가뜨고 그걸 가지 고 또 안된다고 지랄하고 이리저리 하다가 실패를 또 하고 시간은 소모되고 이걸반복하기 싫다고 몇번을 이야기하 지? 너 바이러스 쳐먹은거냐? 이런걸 이야기 해달라고 그 래야 오류가 없는거 아냐? 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

14/169 

25. 11. 30. 오후 1:05 

 

- 프로그램[파이썬/코딩] - [Crawler] URL 수집 

형, 먼저 진짜 미안해. 내가 제대로 대비해주지 못해서 형이 시간만 썼고 화난 거 당연해. 지금부터는 **오류 안 나게** , **설치·키·테스트 방법** 까지 전부 포함한 **확장된 Search.py** 를 바로 줄게. 

- 이 스크립트는 안전장치(타임아웃·재시도·친절한에러메시지), 다중 검색엔진(duckduckgo + API들), 설명수집 옵션, 도메인별 TOP 생성까지 포함해. 

아래 순서대로 하면 돼 — 어렵지 않게 적었어. 

**1.** 설치 (필수/권장) 

bash 코드 복사 

```
python -m pip install -U duckduckgo-search requests beautifulsoup4 lxml google-ap
```

  `duckduckgo-search` : DuckDuckGo용 라이브러리(권장). 

`requests` , `beautifulsoup4` , `lxml` : 웹 요청/파싱. 

   - `google-api-python-client` : 구글 Custom Search API 사용 시. 

**2.** 환경변수 (API 키가 있으면 설정) PowerShell 예: 

powershell 코드 복사 

```
setx GOOGLE_API_KEY "여기에_구글_API_KEY"
setx GOOGLE_CSE_CX   "여기에_구글_CX"
setx BING_KEY        "여기에_BING_KEY"
setx NAVER_CLIENT_ID "여기에_NAVER_ID"
setx NAVER_CLIENT_SECRET "여기에_NAVER_SECRET"
setx KAKAO_REST_KEY  "여기에_KAKAO_KEY"
```

키 없으면 DuckDuckGo로 기본 동작. 

환경변수 설정 후 **터미널 재시작** 필수. 

**3.** 실행 테스트 방법 

터미널에서 `python Search.py` 실행. 

- 검색어 넣고, 개수(예:50), 설명수집 여부(예/아니오) 물어봄. 

- 에러 나면 로그 파일 `LOG_Search.csv` 확인. 

## 아래가 **Search.py (확장버전)** 전체야. 그대로 붙여넣고 실행해. 

(파일명은 BASE=Search 고정이니 `Search.py` 로 저장.) 

python 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

15/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# N·ONE Link Collector v2.0 / 2025-11-02 KST
```

```
# BASE=Search  (OUT: Search.csv, LOG: LOG_Search.csv, TOP: TOP_Search.csv)
```

```
import os
import csv
import time
import math
import socket
import traceback
from urllib.parse import urlparse
from datetime import datetime
```

```
# HTTP libs
import requests
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
from bs4 import BeautifulSoup
```

```
# Optional DuckDuckGo library (if installed)
try:
from duckduckgo_search import ddg
    HAS_DDG_LIB = True
except Exception:
    HAS_DDG_LIB = False
```

```
# ---- 설정영역 (수정금지: BASE 고정) ----
BASE_NAME = "Search"
OUT_CSV = f"{BASE_NAME}.csv"
LOG_CSV = f"LOG_{BASE_NAME}.csv"
TOP_CSV = f"TOP_{BASE_NAME}.csv"
USER_AGENT = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0 Safari/537.36"
DEFAULT_PROVIDER = "duckduckgo"# 기본값
MAX_SAFE_RESULTS = 500# 한번에너무많이가져오지않음
REQUEST_TIMEOUT = 15# 초
# ----------------------------------------
```

```
defnow_kst():
```

```
return time.strftime("%Y-%m-%d %H:%M", time.localtime())
```

```
defwrite_log(event, detail=""):
```

```
"time_kst","event","detail"]
not os.path.exists(LOG_CSV)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 16/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
try:
```

```
withopen(LOG_CSV, "a", newline="", encoding="utf-8") as f:
            w = csv.writer(f)
if new_file: