
https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

52/169 

25. 11. 30. 오후 1:05 

 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

```
-OUT_CSV = f"{BASE_NAME}.csv"
-LOG_CSV = f"LOG_{BASE_NAME}.csv"
-TOP_CSV = f"TOP_{BASE_NAME}.csv"
+BASE_DIR = os.path.dirname(os.path.abspath(__file__))
+OUT_CSV = os.path.join(BASE_DIR, f"{BASE_NAME}.csv")
+LOG_CSV = os.path.join(BASE_DIR, f"LOG_{BASE_NAME}.csv")
+TOP_CSV = os.path.join(BASE_DIR, f"TOP_{BASE_NAME}.csv")
```

diff 코드 복사 

```
-def save_top(rows):
+def save_top(rows):
     agg = {}
     for r in rows:
         d = r.get("domain","");
         if not d: continue
         ...
-    out.sort(key=lambda x: (-x["count"], x["best_rank"]))
-    with open(TOP_CSV, "w", newline="", encoding="utf-8-sig") as f:
+    out.sort(key=lambda x: (-x["count"], x["best_rank"]))
+    # 결과 0이어도헤더만생성
+    with open(TOP_CSV, "w", newline="", encoding="utf-8-sig") as f:
         w = csv.DictWriter(f, fieldnames=["domain","count","best_rank","examples
-        w.writeheader(); [w.writerow(r) for r in out]
+        w.writeheader()
+        for r in out: w.writerow(r)
```

  

## **(B) DuckDuckGo: 라이브러리 우선, HTML은 완전 차선** 

(이미 그 구조인데, **봇방어** 빈도 낮추려 **대기 시간↑/파라미터 튜닝** ) 

## [REPAIR-16] 

diff 코드 복사 

```
-        start += 50; time.sleep(0.6)
+        start += 30
+        time.sleep(1.2)
```

## **(C) Naver: 요청 헤더/파라미터 보강** 

[REPAIR-16] 

diff 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

53/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
-    r = s.get(url, headers={"X-Naver-Client-Id": cid, "X-Naver-Client-Secret": c
-              params={"query": query, "display": max_results}, timeout=REQUEST_T
+    headers = {
+        "X-Naver-Client-Id": cid,
+        "X-Naver-Client-Secret": csec,
+        "User-Agent": USER_AGENT,
+        "Accept": "application/json"
+    }
+    params = {
+        "query": query,
+        "display": min(100, max(1, int(max_results))),  # 네이버제한 1~100
+        "start": 1                                     # 1부터시작
+    }
+    r = s.get(url, headers=headers, params=params, timeout=REQUEST_TIMEOUT)
```

  

## **(D) 예외가 나도 부분 결과 저장하고 TOP도 만들기** 

[REPAIR-16] 

diff 코드 복사 `-    try: -        rows = collect_links(q, n, prov, fetch_desc) -        if not rows: -            print("결과 없음. 검색어를 바꿔보세요."); write_log("no_results", q); r -        save_results(rows); save_top(rows) +    try: +        rows = collect_links(q, n, prov, fetch_desc) +        # 부분 결과라도 저장 +        save_results(rows) +        save_top(rows) +        if not rows: +            print("결과 없음. 검색어/엔진을 바꿔보세요.") +            write_log("no_results", q) +            return` 

  

## **3) 실행 체크 루틴** 

## **1. DuckDuckGo(빠름)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

54/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

powershell 

 

```
python .\Search.py
# 검색어: 연기커뮤니티
# 개수: 50
# 엔진: duckduckgo
# 설명: n
```

→ **Search.csv** 에 최소 수십 개 나와야 정상. 

## **2. Google(키 연결 확인)** 

powershell _`#`_ `엔진` _`: google /`_ `설명` _`: n`_ 

코드 복사 

## **3. Naver(401 해결 확인)** 

powershell 코드 복사 _`#`_ `엔진` _`: naver /`_ `설명` _`: n`_ 

## **4) 예상되는 현상별 빠른 대응** 

- **DuckDuckGo가 또 202** → 잠시 후 재시도, 검색어 바꾸기, 개수 30으로 줄이기. **Naver 401** → 거의 100% **SECRET 재발급/재등록** 로 해결됨. 

- **TOP이 또 비어 있음** → 결과가 진짜 0개면 헤더만 있는 파일이 정상(이제 생성됨). 

## 형, 위 **차분 그대로** 붙여서 저장하고 

우선 **duckduckgo / 설명 n / 개수 30** 으로 한 번만 돌려봐. 

`Search.csv` 첫 3행하고 **LOG_Search.csv** 의 마지막 5줄만 보내줘. 

그걸로 ddg 상태부터 확실히 잡고, 바로 naver 재시도까지 이어갈게. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

55/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

**==> picture [289 x 126] intentionally omitted <==**

네이버는 다시 발급 받아야 할듯하고 여기서 필요없는 파 일이 있으면 없애주고 out폴더를 생성해서 csv파일만 모아줘 

## 19s 동안 생각함 

## 형, 깔끔하게 정리해줄게. 

**삭제 없이 안전하게 정리 → out 폴더에 CSV만 모으기 → 코드도 out으로 저장되게 패치** 순서야. 

## **1) 지금 있는 파일 정리(안전 이동)** 

PowerShell에서 현재 폴더가 `...\Desktop\Search` 인지 확인한 뒤 실행해. 

powershell 

코드 복사 

```
# ① out 폴더만들기
mkdir ".\out"-Force
```

```
# ② CSV 전부 out으로이동 (로그포함)
```

```
Move-Item-LiteralPath".\*.csv"-Destination".\out"-Force-ErrorAction Silently
```

```
# ③ 옛파일은버리지말고보관
```

```
mkdir ".\_old"-Force
Move-Item".\platform_search_to_urls.py"," .\search_to_urls.py"," .\search_urls"
```

  

## 이후 폴더 상태: 

`Search.py` (실행용) 

`out\Search.csv` , `out\TOP_Search.csv` , `out\LOG_Search.csv` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

56/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

`_old\...` (예전 파일 보관) 

## **2) Search.py를 “항상 out에 저장”하도록 패치** 

아래 **차분** 만 복사해서 Search.py에 반영해줘. 

diff `@@ -import os, csv, time, math, traceback +import os, csv, time, math, traceback @@ -BASE_NAME = "Search" -OUT_CSV = f"{BASE_NAME}.csv" -LOG_CSV = f"LOG_{BASE_NAME}.csv" -TOP_CSV = f"TOP_{BASE_NAME}.csv" +BASE_NAME = "Search" +BASE_DIR  = os.path.dirname(os.path.abspath(__file__)) +# 저장 폴더: out (없으면 자동 생성) +OUT_DIR   = os.path.join(BASE_DIR, "out") +os.makedirs(OUT_DIR, exist_ok=True) +OUT_CSV = os.path.join(OUT_DIR, f"{BASE_NAME}.csv") +LOG_CSV = os.path.join(OUT_DIR, f"LOG_{BASE_NAME}.csv") +TOP_CSV = os.path.join(OUT_DIR, f"TOP_{BASE_NAME}.csv")` 

코드 복사 

이렇게 하면 어디서 실행해도 결과/로그 CSV는 전부 `out` 에 떨어져. 

## **3) 빠른 테스트** 

powershell 코드 복사 `python .\Search.py` _`#`_ `검색어` _`:`_ `연기 커뮤니티` _`#`_ `개수` _`: 30 #`_ `엔진` _`: duckduckgo  (`_ `또는` _`google) #`_ `설명` _`: n`_ 

끝나면 `out` 폴더 안에 **Search.csv / TOP_Search.csv / LOG_Search.csv** 있어야 정상. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

57/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## **4) 네이버 401 해결 팁(재발급 예정이라면)** 

콘솔에서 **SECRET 재발급 → setx로 재입력 → PowerShell 재시작** 

환경: **웹서버 /** http://localhost 확인 

사용 API: **검색** 체크 

필요하면 **삭제(완전 제거)**도 명령어 바로 만들어줄게. 

일단은 안전 보관 방식으로 정리해뒀어. 

**==> picture [489 x 38] intentionally omitted <==**

## 다시 py에 니가 만들어줬던 내용 적용해서 만들어줘 

## 확인. 아래 코드를 **Search.py** 로 붙여넣어 저장해줘. 

(결과 파일은 모두 `out` 폴더에 자동 저장) 

python 

코드 복사 

```
# N·ONE Link Collector v2.2 / 2025-11-04 14:10 KST
```

```
# BASE=Search  (OUT: out/Search.csv, LOG: out/LOG_Search.csv, TOP: out/TOP_Search
```

```
import os, csv, time, math, traceback
from urllib.parse import urlparse
import requests
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
from bs4 import BeautifulSoup
```

```
# DuckDuckGo 라이브러리(있으면사용)
try:
```

```
from duckduckgo_search import ddg
    HAS_DDG_LIB = True
except Exception:
    HAS_DDG_LIB = False