from bs4 import BeautifulSoup
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
```

```
# ────────────────────────────────────────────────────────────
```

```
BASE_DIR  = os.path.dirname(os.path.abspath(__file__))
OUT_DIR   = os.path.join(BASE_DIR, "out")
os.makedirs(OUT_DIR, exist_ok=True)
```

```
USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML
REQUEST_TIMEOUT = 15
```

```
MAX_SAFE_RESULTS = 200
```

```
# ────────────────────────────────────────────────────────────
```

```
defnow_kst(): return time.strftime("%Y-%m-%d %H:%M", time.localtime())
```

```
defsafe_name(t): return re.sub(r"[^0-9A-Za-z가-힣]+", "", t)[:50]
```

```
defclean_text(s): return" ".join((s or"").split()).strip()
```

```
defdomain_of(url):
```

```
try: return urlparse(url).netloc.lower().replace("www.", "")
```

```
except: return""
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
    retries = Retry(total=3, backoff_factor=0.5, status_forcelist=[500,502,503,504
```

```
    s.mount("https://", HTTPAdapter(max_retries=retries))
return s
```

```
defwrite_log(path, event, detail=""):
```

```
    new_file = not os.path.exists(path)
```

```
withopen(path, "a", newline="", encoding="utf-8") as f:
```

```
        w = csv.writer(f)
```

```
if new_file: w.writerow(["time_kst","event","detail"])
```

```
        w.writerow([now_kst(), event, detail])
```

```
# ── Serper.dev 검색 ───────────────────────────────────────
```

## `def search_serper(query, n):` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

95/169 

25. 11. 30. 오후 1:05 

 

후 1:05 1:05 프로그램[파이썬/코딩] - [Crawler] URL 수집 `api_key = os.getenv("SERPER_API_KEY") if not api_key:` 

```
raise RuntimeError("Serper.dev API 키가없습니다. setx SERPER_API_KEY 명령으
    url = "https://google.serper.dev/search"
    headers = {"X-API-KEY": api_key, "Content-Type": "application/json"}
    data = {"q": query, "num": n}
try:
        r = requests.post(url, headers=headers, json=data, timeout=REQUEST_TIMEOUT
if r.status_code != 200:
raise RuntimeError(f"Serper API 오류 ({r.status_code})")
        js = r.json()
        out = []
for i in js.get("organic", []):
            out.append({"title": clean_text(i.get("title","")), "url": i.get("lin
return out
except Exception as e:
print("Serper 오류:", e)
return []
```

```
# ── Naver 검색 ─────────────────────────────────────────────
defsearch_naver(query, n):
    cid, csec = os.getenv("NAVER_CLIENT_ID"), os.getenv("NAVER_CLIENT_SECRET")
ifnot cid ornot csec:
raise RuntimeError("네이버 API 키가없습니다.")
    url = "https://openapi.naver.com/v1/search/webkr.json"
    headers = {
"X-Naver-Client-Id": cid,
"X-Naver-Client-Secret": csec,
"User-Agent": USER_AGENT
    }
    params = {"query": query, "display": min(100, n), "start": 1}
try:
```

```
        r = requests.get(url, headers=headers, params=params, timeout=REQUEST_TIM
if r.status_code != 200: raise RuntimeError(f"Naver API 오류 ({r.status_co
        js = r.json()
        out = []
for i in js.get("items", []):
            out.append({"title": clean_text(i.get("title","")), "url": i.get("lin
return out
except Exception as e:
print("Naver 오류:", e)
return []
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 96/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# ── 결과정리 ───────────────────────────────────────────────
```

```
deffetch_description(url):
try:
        s = safe_session()
        r = s.get(url, timeout=8)
if r.status_code != 200: return""
        soup = BeautifulSoup(r.text, "lxml")
        meta = soup.find("meta", attrs={"name":"description"}) or soup.find("meta
if meta and meta.get("content"): return clean_text(meta["content"])
        p = soup.find("p")
return clean_text(p.get_text()[:200]) if p else""
except: return""
```

```
defmerge_and_clean(*sources):
    seen, merged = set(), []
for src in sources:
for it in src:
            u = it.get("url","").strip()
ifnot u or u in seen: continue
            seen.add(u)
            merged.append(it)
for i,r inenumerate(merged, start=1):
        r["rank"] = i
        r["domain"] = domain_of(r["url"])
        r["fetched_at_kst"] = now_kst()
return merged
defsave_csv(rows, path):
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=headers); w.writeheader()
for r in rows: w.writerow(r)
defsave_top(rows, path):
    agg = {}
for r in rows:
        d = r.get("domain","")
ifnot d: continue
if d notin agg: agg[d] = {"count":0, "best_rank":9999}
        agg[d]["count"] += 1
        agg[d]["best_rank"] = min(agg[d]["best_rank"], r.get("rank",9999))
    out = [{"domain":k, "count":v["count"], "best_rank":v["best_rank"]} for k,v i
    out.sort(key=lambda x:(-x["count"], x["best_rank"]))
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 97/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=["domain","count","best_rank"]); w.write
for r in out: w.writerow(r)
```

_`# ──`_ `메인` _`───────────────────────────────────────────────────`_ `def main(): print("검색어를 입력하세요:") q = input("> ").strip() if not q: print("검색어 없음. 종료."); return q_clean = safe_name(q) try: print("최대 개수(기본 50):") n = int(input("> ").strip() or "50") n = max(1, min(MAX_SAFE_RESULTS, n)) except: n = 50 print("설명(meta)도 모을까요? (y/N)") fetch_desc = (input("> ").strip().lower() == "y") OUT_CSV = os.path.join(OUT_DIR, f"{q_clean}_mix.csv") LOG_CSV = os.path.join(OUT_DIR, f"LOG_{q_clean}_mix.csv") TOP_CSV = os.path.join(OUT_DIR, f"TOP_{q_clean}_mix.csv") write_log(LOG_CSV, "run_start", f"query={q}, n={n}") try: s1 = search_serper(q, n) s2 = search_naver(q, n) rows = merge_and_clean(s1, s2) if fetch_desc: for r in rows: r["description"] = fetch_description(r["url"]) else: for r in rows: r["description"] = "" if not rows: print("결과 없음. API키 및 검색어 확인") write_log(LOG_CSV, "no_results", q) return save_csv(rows, OUT_CSV) save_top(rows, TOP_CSV) print(f"` ✅ `완료: {len(rows)}개 저장 → {OUT_CSV}") write_log(LOG_CSV, "done", f"saved={len(rows)}") except Exception as e:` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

98/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
print("오류:", str(e))
        write_log(LOG_CSV, "error", str(e))
```

```
if __name__ == "__main__":
    main()
```

## ⚙ **사전 준비 (1회만)** 

## 1️⃣ **Serper.dev 가입** 

https://serper.dev 접속 → 로그인 후 API Key 복사 

2️⃣ PowerShell에서 등록: 

powershell 코드 복사 `setx SERPER_API_KEY "복사한_API키"` (선택) 네이버도 같이 쓰려면: powershell 코드 복사 `setx NAVER_CLIENT_ID "아이디" setx NAVER_CLIENT_SECRET "시크릿"` 새 PowerShell 열고 확인: powershell 코드 복사 `echo %SERPER_API_KEY% echo %NAVER_CLIENT_ID%` 

## 3️⃣ (선택) 네이버도 같이 쓰려면: 

## 4️⃣ 새 PowerShell 열고 확인: 

## 🚀 **실행** 

powershell 코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
python .\Search.py
```

## 입력 예시 👇 

makefile 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

99/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
검색어: 연기오디션
개수: 50
설명: n
```

## → 결과 

csharp `out/연기오디션_mix.csv out/LOG_연기오디션_mix.csv out/TOP_연기오디션_mix.csv` 