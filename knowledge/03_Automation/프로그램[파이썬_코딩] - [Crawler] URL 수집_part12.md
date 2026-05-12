# ── 고정설정 ─────────────────────────────────────────────────────────────
BASE_NAME = "Search"
BASE_DIR  = os.path.dirname(os.path.abspath(__file__))
OUT_DIR   = os.path.join(BASE_DIR, "out")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

58/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
os.makedirs(OUT_DIR, exist_ok=True)
```

```
OUT_CSV = os.path.join(OUT_DIR, f"{BASE_NAME}.csv")
LOG_CSV = os.path.join(OUT_DIR, f"LOG_{BASE_NAME}.csv")
TOP_CSV = os.path.join(OUT_DIR, f"TOP_{BASE_NAME}.csv")
```

```
USER_AGENT = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0 Safari/537.36"
DEFAULT_PROVIDER = "duckduckgo"
MAX_SAFE_RESULTS = 500
REQUEST_TIMEOUT  = 15
# ──────────────────────────────────────────────────────────────────────────
```

```
defnow_kst():
return time.strftime("%Y-%m-%d %H:%M", time.localtime())
```

```
defwrite_log(event, detail=""):
    header = ["time_kst","event","detail"]
    new_file = not os.path.exists(LOG_CSV)
try:
withopen(LOG_CSV, "a", newline="", encoding="utf-8") as f:
            w = csv.writer(f)
if new_file: w.writerow(header)
            w.writerow([now_kst(), event, detail])
except Exception:
pass
```

```
defsafe_session():
    s = requests.Session()
    s.headers.update({"User-Agent": USER_AGENT})
    retries = Retry(total=3, backoff_factor=0.8, status_forcelist=(500,502,503,504
    s.mount("https://", HTTPAdapter(max_retries=retries))
    s.mount("http://",  HTTPAdapter(max_retries=retries))
return s
```

```
defdomain_of(url):
try:
return urlparse(url).netloc.lower().replace("www.", "")
except:
return""
defclean_text(s):
return" ".join((s or"").split()).strip()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

59/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# ── 검색엔진들 ───────────────────────────────────────────────────────────
defsearch_duckduckgo_lib(query, max_results=50):
ifnot HAS_DDG_LIB: return []
try:
        out = ddg(query, max_results=max_results) or []
return [{"title": clean_text(i.get("title","")), "url": i.get("href","")}
except Exception as e:
        write_log("ddg_lib_error", str(e)); return []
defsearch_duckduckgo_html(query, max_results=50, session=None):
    s = session or safe_session()
    results, base, start = [], "https://html.duckduckgo.com/html/", 0
whilelen(results) < max_results:
try:
            r = s.post(base, data={"q": query, "s": start}, timeout=REQUEST_TIMEOU
except Exception as e:
            write_log("ddg_html_error", str(e)); break
if r.status_code != 200:
            write_log("ddg_html_status", f"{r.status_code}"); break
        soup = BeautifulSoup(r.text, "html.parser")
        anchors = soup.select("a.result__a")
ifnot anchors: break
for a in anchors:
            href = a.get("href"); title = clean_text(a.get_text(" ", strip=True))
if href and title:
                results.append({"title": title, "url": href})
iflen(results) >= max_results: break
        start += 30
        time.sleep(1.2)  # 봇방어완화
return results
```

```
defsearch_google_api(query, max_results=10):
    api_key, cx = os.getenv("GOOGLE_API_KEY"), os.getenv("GOOGLE_CSE_CX")
ifnot api_key ornot cx: raise RuntimeError("구글 API 키또는 CX가없습니다.")
try:
from googleapiclient.discovery import build
        service = build("customsearch", "v1", developerKey=api_key)
        results, start_index = [], 1
whilelen(results) < max_results:
            resp = service.cse().list(q=query, cx=cx, start=start_index,
                                      num=min(10, max_results - len(results))).exe
            items = resp.get("items", [])
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 60/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
for it in items:
```

```
                results.append({"title": clean_text(it.get("title","")), "url": it
if"nextPage"notin resp.get("queries", {}): break
            start_index += len(items)
return results
```

```
except Exception as e:
```

```
        write_log("google_api_error", str(e)); raise
```

```
defsearch_naver_api(query, max_results=10):
```

```
    cid, csec = os.getenv("NAVER_CLIENT_ID"), os.getenv("NAVER_CLIENT_SECRET")
```

```
ifnot cid ornot csec: raise RuntimeError("네이버 API 키가없습니다.")
```

```
    url = "https://openapi.naver.com/v1/search/webkr.json"
    s = safe_session()
    headers = {
"X-Naver-Client-Id": cid,
"X-Naver-Client-Secret": csec,
"User-Agent": USER_AGENT,
"Accept": "application/json",
    }
```

```
    params = {"query": query, "display": min(100, max(1, int(max_results))), "sta
    r = s.get(url, headers=headers, params=params, timeout=REQUEST_TIMEOUT)
if r.status_code != 200:
        write_log("naver_status", f"{r.status_code}")
raise RuntimeError("네이버 API 오류")
    js = r.json()
return [{"title": clean_text(i.get("title","")), "url": i.get("link","")} for
```

```
# ── 설명(메타) 수집 ───────────────────────────────────────────────────────
deffetch_description(url, session=None):
    s = session or safe_session()
try:
        r = s.get(url, timeout=REQUEST_TIMEOUT)
if r.status_code != 200: return""
        soup = BeautifulSoup(r.content, "lxml")
        meta = soup.find("meta", attrs={"name":"description"}) or soup.find("meta
if meta and meta.get("content"): return clean_text(meta.get("content"))
        p = soup.find("p")
return clean_text(p.get_text()[:300]) if p else""
except Exception:
return""
```

```
# ── 저장 ──────────────────────────────────────────────────────────────────
defsave_results(rows):
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 61/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
withopen(OUT_CSV, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=headers)
        w.writeheader()
for r in rows: w.writerow(r)
defsave_top(rows):
    agg = {}
for r in rows:
        d = r.get("domain","")
ifnot d: continue
if d notin agg: agg[d] = {"count":0, "best_rank": 10**9, "ex":[]}
        agg[d]["count"] += 1
        agg[d]["best_rank"] = min(agg[d]["best_rank"], r.get("rank", 10**9))
iflen(agg[d]["ex"]) < 3: agg[d]["ex"].append(f"{r.get('rank')}:{r.get('t
    out = [{"domain": k, "count": v["count"], "best_rank": v["best_rank"], "examp
for k,v in agg.items()]
    out.sort(key=lambda x: (-x["count"], x["best_rank"]))
withopen(TOP_CSV, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=["domain","count","best_rank","examples"
        w.writeheader()
for r in out: w.writerow(r)
```

```
# ── 수집흐름 ────────────────────────────────────────────────────────────
```

```
defcollect_links(query, max_n=50, provider=DEFAULT_PROVIDER, fetch_desc=False):
```

```
    write_log("collect_start", f"q={query},n={max_n},prov={provider},desc={fetch_d
    session = safe_session()
if provider == "duckduckgo":
        hits = search_duckduckgo_lib(query, max_n) or search_duckduckgo_html(query
elif provider == "google":
        hits = search_google_api(query, max_n)
elif provider == "naver":
        hits = search_naver_api(query, max_n)
else:
raise RuntimeError("지원하지않는 provider")
```

```
    seen, rows, rank = set(), [], 0
for it in hits:
        u = (it.get("url") or"").strip()
        t = clean_text(it.get("title",""))
ifnot u or u in seen: continue
        seen.add(u); rank += 1
        desc = fetch_description(u, session=session) if fetch_desc else""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

62/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
        rows.append({
"rank": rank, "title": t, "url": u, "source": provider,
"query": query, "fetched_at_kst": now_kst(),
"description": desc, "domain": domain_of(u)
        })
if rank >= max_n: break
return rows
defmain():
print("검색어를입력하세요:")
    q = input("> ").strip()
ifnot q: print("검색어없음. 종료."); return
try:
print("최대개수(기본 50):")
        n = int(input("> ").strip() or"50")
        n = max(1, min(MAX_SAFE_RESULTS, n))
except: n = 50