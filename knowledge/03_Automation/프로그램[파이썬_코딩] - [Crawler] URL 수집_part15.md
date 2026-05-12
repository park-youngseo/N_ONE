프로그램[파이썬/코딩] - [Crawler] URL 수집 

```
print("검색어를입력하세요:")
    q = input("> ").strip()
ifnot q: print("검색어없음. 종료."); return
try:
print("최대개수(기본 50):")
        n = int(input("> ").strip() or"50")
        n = max(1, min(MAX_SAFE_RESULTS, n))
except: n = 50
```

```
print("엔진선택: duckduckgo / google / naver (기본 duckduckgo)")
    prov = (input("> ").strip().lower() or"duckduckgo")
if prov notin ("duckduckgo","google","naver"):
print("지원하지않는엔진. duckduckgo 사용"); prov = "duckduckgo"
print("설명(meta)도모을까요? (y/N)")
    fetch_desc = (input("> ").strip().lower() == "y")
    prov_suffix = f"_{prov}"
    OUT_CSV = os.path.join(OUT_DIR, f"{BASE_NAME}{prov_suffix}.csv")
    LOG_CSV = os.path.join(OUT_DIR, f"LOG_{BASE_NAME}{prov_suffix}.csv")
    TOP_CSV = os.path.join(OUT_DIR, f"TOP_{BASE_NAME}{prov_suffix}.csv")
    write_log(LOG_CSV, "run_request", f"q={q},n={n},prov={prov},desc={fetch_desc}
try:
        rows = collect_links(q, n, prov, fetch_desc)
        save_results(rows, OUT_CSV)
        save_top(rows, TOP_CSV)
ifnot rows:
"
print(결과없음. 검색어/엔진을바꿔보세요.")
            write_log(LOG_CSV, "no_results", q)
return
print(f"완료: {len(rows)}개저장 → {OUT_CSV}, {TOP_CSV}")
        write_log(LOG_CSV, "run_done", f"saved={len(rows)}")
except Exception as e:
print("오류:", str(e))
print(f"자세한내용은{LOG_CSV}확인")
        write_log(LOG_CSV, "run_exception", traceback.format_exc())
```

```
if __name__ == "__main__":
    main()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

72/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

코드 복사 

## ✅ **실행 예시** 

powershell 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
python .\Search.py
```

## 입력 

- 검색어: `연기 커뮤니티` 

- 개수: `30` 

- 엔진: `google` 

- 설명: `n` 

## → 결과: 

csharp 

코드 복사 

```
out/Search_google.csv
out/TOP_Search_google.csv
out/LOG_Search_google.csv
```

## 이 버전(v2.3)은 

- ✔ 자동 out폴더 생성 

- ✔ 포털명 자동 추가 

- ✔ 부분 결과도 저장 

- ✔ TOP, LOG 항상 생성 

## 전부 반영된 완성판이야. 

실행 후 `out` 폴더 캡처만 보내주면 로그까지 정리해줄게. 

**==> picture [489 x 39] intentionally omitted <==**

**==> picture [289 x 117] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

73/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

이렇게 생성이 되어 있으니 파일이 뭐가 있는지 모르겠네 out 파일이름이 검색어 이름으로 만들어줄수 없어? 가령 TOP_Search_google 이라면 검색어 TOP_연기오디션 _google 이런식으로 

좋아 형 👌 완전 좋은 개선 포인트야. 이제 v2.4에서는 파일 이름이 **검색어 중심으로 자동 저장** 되게 바꿔줄게. 

## 🔧 **새로운 저장 규칙 (v2.4)** 

예시: 

검색어 → `연기 오디션` 

엔진 → `google` 

그럼 아래처럼 저장돼 👇 

csharp 코드 복사 

```
out/연기오디션_google.csv
out/TOP_연기오디션_google.csv
out/LOG_연기오디션_google.csv
```

(검색어에 띄어쓰기나 특수문자 있으면 자동으로 안전하게 정리돼) 

## 🔁 **수정된 완전 버전 (Search.py v2.4)** 

아래 전체를 그대로 `Search.py` 에 붙여 넣어줘. 

python 코드 복사 

```
# N·ONE Link Collector v2.4 / 2025-11-04 15:45 KST
```

```
# BASE=Search  (OUT: out/[검색어]_[엔진].csv)
```

```
import os, csv, time, math, traceback, re
from urllib.parse import urlparse
import requests
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

74/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
from bs4 import BeautifulSoup
```

```
try:
```

```
from duckduckgo_search import ddg
    HAS_DDG_LIB = True
except Exception:
    HAS_DDG_LIB = False
```

```
# ────────────────────────────────────────────────────────
BASE_DIR  = os.path.dirname(os.path.abspath(__file__))
OUT_DIR   = os.path.join(BASE_DIR, "out")
os.makedirs(OUT_DIR, exist_ok=True)
```

```
USER_AGENT = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
```

```
"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0 Safari/537.36"
MAX_SAFE_RESULTS = 500
```

```
REQUEST_TIMEOUT  = 15
```

```
# ────────────────────────────────────────────────────────
```

```
defnow_kst(): return time.strftime("%Y-%m-%d %H:%M", time.localtime())
```

```
defsafe_name(text): return re.sub(r"[^0-9A-Za-z가-힣]+", "", text)[:50]
```

```
defsafe_session():
```

```
    s = requests.Session()
```

```
    s.headers.update({"User-Agent": USER_AGENT})
```

```
    retries = Retry(total=3, backoff_factor=0.8, status_forcelist=(500,502,503,
"https://", HTTPAdapter(max_retries=retries))
```

```
"http://",  HTTPAdapter(max_retries=retries))
return s
```

```
defclean_text(s): return" ".join((s or"").split()).strip()
```

```
defdomain_of(url):
try: return urlparse(url).netloc.lower().replace("www.", "")
except: return""
```

```
defwrite_log(path, event, detail=""):
```

```
not os.path.exists(path)
```

```
try:
```

```
withopen(path, "a", newline="", encoding="utf-8") as f:
            w = csv.writer(f)
if new_file: w.writerow(["time_kst","event","detail"])
            w.writerow([now_kst(), event, detail])
except Exception: pass
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 75/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# ── 검색엔진별함수 ───────────────────────────────────────
```

```
defsearch_duckduckgo(query, n, s):
```

```
try:
```

```
if HAS_DDG_LIB:
            out = ddg(query, max_results=n) or []
return [{"title": clean_text(i.get("title","")), "url": i.get("href",
        results, base, start = [], "https://html.duckduckgo.com/html/", 0
whilelen(results) < n:
            r = s.post(base, data={"q": query, "s": start}, timeout=REQUEST_TIMEOU