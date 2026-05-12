```
# 작성자: 박성민형시스템기준
```

```
import os, csv, time, requests
```

```
from datetime import datetime
```

```
from urllib.parse import urlparse
```

```
from duckduckgo_search import DDGS
```

```
# ===== 폴더 =====
```

```
base_dir = os.path.dirname(os.path.abspath(__file__))
out_dir = os.path.join(base_dir, "out")
```

```
os.makedirs(out_dir, exist_ok=True)
```

```
# ===== API =====
```

```
NAVER_ID = os.getenv("NAVER_CLIENT_ID")
NAVER_SECRET = os.getenv("NAVER_CLIENT_SECRET")
SERPER_KEY = os.getenv("SERPER_API_KEY")
```

```
# ===== 네이버 =====
```

```
defnaver_search(query, n=30):
```

```
    headers = {"X-Naver-Client-Id": NAVER_ID, "X-Naver-Client-Secret": NAVER_SECR
    url = f"https://openapi.naver.com/v1/search/webkr.json?query={query}&display=
    res = requests.get(url, headers=headers)
```

```
    data = res.json()
```

```
return [(i["title"], i["link"], i["description"]) for i in data.get("items",
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 134/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# ===== Serper =====
```

```
defserper_search(query, n=30):
```

```
    headers = {"X-API-KEY": SERPER_KEY, "Content-Type": "application/json"}
    payload = {"q": query, "num": n}
    r = requests.post("https://google.serper.dev/search", headers=headers, json=pa
    data = r.json()
return [(i["title"], i["link"], i.get("snippet", "")) for i in data.get("orga
```

```
# ===== Google (DuckDuckGo 방식) =====
defgoogle_search(query, n=30):
    results = []
with DDGS() as ddgs:
for r in ddgs.text(query, max_results=n):
            results.append((r["title"], r["href"], r["body"]))
return results
```

```
# ===== Domain =====
defget_domain(url):
try:
return urlparse(url).netloc.replace("www.", "")
except:
return"unknown"
```

```
# ===== CSV Save =====
defsave_csv(path, rows):
withopen(path, "w", encoding="utf-8-sig", newline="") as f:
        w = csv.writer(f)
        w.writerow(["title", "url", "desc"])
        w.writerows(rows)
```

```
# ===== Log =====
defwrite_log(keyword, msg):
    path = os.path.join(out_dir, f"LOG_{keyword}_all.csv")
withopen(path, "a", encoding="utf-8-sig", newline="") as f:
        w = csv.writer(f)
        w.writerow([datetime.now().strftime("%Y-%m-%d %H:%M"), msg])
```

```
# ===== 병합 =====
defmerged_search(keyword, n=30):
    all_results = []
```

```
    engines = {
"serper": serper_search,
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 135/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
"google": google_search,
"naver": naver_search,
    }
for name, func in engines.items():
try:
            res = func(keyword, n)
            save_csv(os.path.join(out_dir, f"{keyword}_{name}.csv"), res)
            all_results.extend(res)
            write_log(keyword, f"{name}성공 ({len(res)}개)")
except Exception as e:
            write_log(keyword, f"{name}실패: {e}")
# 중복제거
    unique = {}
for t, u, d in all_results:
if u notin unique:
            unique[u] = (t, d)
    merged = [(v[0], k, v[1]) for k, v in unique.items()]
# 본결과
    save_csv(os.path.join(out_dir, f"{keyword}_all.csv"), merged)
```

```
# 도메인별 Top10
    dom_count = {}
for _, url, _ in merged:
        dom = get_domain(url)
        dom_count[dom] = dom_count.get(dom, 0) + 1
    top10 = sorted(dom_count.items(), key=lambda x: x[1], reverse=True)[:10]
```

```
withopen(os.path.join(out_dir, f"TOP_{keyword}_all.csv"), "w", encoding="utf
        w = csv.writer(f)
        w.writerow(["domain", "count"])
        w.writerows(top10)
```

`write_log(keyword, f"병합 완료 ({len(merged)}개)") print(f"\n` ✅ `병합 완료: {len(merged)}개 → {keyword}_all.csv / TOP_{keyword}_al` _`# =====`_ `실행` _`=====`_ `if __name__ == "__main__": kw = input("검색어를 입력하세요: ").strip() n = input("수집 개수 (기본 30): ").strip() n = int(n) if n else 30` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 136/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

`print(f"\n` 🔍 `검색 시작: {kw}")` 

 

```
    merged_search(kw, n)
```

## 💾 **저장 이름** 

📁 `C:\Users\psm23\OneDrive\Desktop\Search\Search_v2.7_fixed.py` 

## 이제 실행만 하면, 

📂 결과는 다음처럼 깔끔하게 유지돼 👇 

python 

코드 복사 

```
out/
```

```
 ├─ 드라마오디션_serper.csv
 ├─ 드라마오디션_google.csv
 ├─ 드라마오디션_naver.csv
 ├─ 드라마오디션_all.csv
```

```
 ├─ TOP_드라마오디션_all.csv
```

```
 └─ LOG_드라마오디션_all.csv
```

## 이 버전으로 바로 쓸까? 

(원하면 `meta설명 수집` 기능도 같이 넣어서 다음 버전 `v2.8` 으로 업그레이드 가능해.) 

**==> picture [489 x 39] intentionally omitted <==**

설명수집도 같이 넣어서 만들고 Search_v2.7_fixed.py 하지 말고 

기존처럼 Search로만들어 

9s 동안 생각함 

형, 요청대로 **설명 수집 포함 + 파일명 통일 + 단일 파일** `Search.py` 완성본이야. 

- (엔진: `serper / google / naver / duckduckgo / mix` 선택 가능, 결과는 `out/검색어_엔 진.csv` 와 `out/검색어_all.csv` 형식) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 137/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

코드 복사 

python 

 

```
# N·ONE Link Collector v2.9 / 2025-11-04 19:40 KST
```

```
# FILE: Search.py  (결과폴더: out/)
```

```
# 엔진선택형: serper / google / naver / duckduckgo / mix
```

```
# 파일명규칙: 개별엔진 → out/[검색어]_[엔진].csv
```

```
#             병합본   → out/[검색어]_all.csv, out/TOP_[검색어]_all.csv, out/LOG_[검
```

```
import os, csv, re, time, traceback, requests
```

```
from urllib.parse import urlparse
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
```

```
from bs4 import BeautifulSoup
```

```
# ───────────────────── 공통설정 ─────────────────────
```

```
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
OUT_DIR  = os.path.join(BASE_DIR, "out")
```