
 

```
# ── 공통 ─────────────────────────────────────────────────
```

```
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
```

```
OUT_DIR  = os.path.join(BASE_DIR, "out")
os.makedirs(OUT_DIR, exist_ok=True)
```

```
USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML
REQUEST_TIMEOUT = 12
```

```
MAX_RESULTS_CAP = 200
```

```
defnow_kst(): return time.strftime("%Y-%m-%d %H:%M", time.localtime())
```

```
defclean_text(s): return" ".join((s or"").split()).strip()
```

```
defsafe_name(t): return re.sub(r"[^0-9A-Za-z가-힣]+", "", t)[:60]
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
defsession():
```

```
    s = requests.Session()
```

```
    s.headers.update({"User-Agent": USER_AGENT})
```

```
    s.mount("https://", HTTPAdapter(max_retries=Retry(total=3, backoff_factor=0.6
return s
```

```
defwrite_log(keyword, msg):
```

```
    log_path = os.path.join(OUT_DIR, f"LOG_{keyword}_all.csv")
```

```
    new = not os.path.exists(log_path)
```

```
withopen(log_path, "a", encoding="utf-8-sig", newline="") as f:
```

```
if new: w.writerow(["time_kst","event"])
```

```
# ── 설명수집 ────────────────────────────────────────────
```

```
deffetch_description(url):
```

```
try:
```

```
8)
```

```
if r.status_code != 200: return""
        soup = BeautifulSoup(r.content, "lxml")
        meta = soup.find("meta", attrs={"name":"description"}) or soup.find("meta
if meta and meta.get("content"): return clean_text(meta.get("content"))
# 본문요약(간단)
```

```
        body_html = doc.summary(html_partial=True)
        body = BeautifulSoup(body_html, "lxml").get_text(" ")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 152/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
return clean_text(body[:300])
except Exception:
return""
```

```
# ── 엔진별검색 ──────────────────────────────────────────
defsearch_serper(query, n):
    key = os.getenv("SERPER_API_KEY")
ifnot key: write_log(safe_name(query), "Serper 키없음"); return []
    url = "https://google.serper.dev/search"
    headers = {"X-API-KEY": key, "Content-Type": "application/json"}
    data = {"q": query, "num": n}
try:
        r = requests.post(url, headers=headers, json=data, timeout=REQUEST_TIMEOUT
if r.status_code != 200: write_log(safe_name(query), f"Serper 오류{r.stat
        js = r.json()
return [{"title": clean_text(i.get("title","")), "url": i.get("link",""),
except Exception as e:
```

```
        write_log(safe_name(query), f"Serper 예외: {e}"); return []
```

```
defsearch_google(query, n):
```

```
    key, cx = os.getenv("GOOGLE_API_KEY"), os.getenv("GOOGLE_CSE_CX")
ifnot key ornot cx: write_log(safe_name(query), "Google 키/CX 없음"); return
from googleapiclient.discovery import build
try:
        svc = build("customsearch", "v1", developerKey=key)
        out, start = [], 1
whilelen(out) < n:
            res = svc.cse().list(q=query, cx=cx, start=start, num=min(10, n-len(o
for it in res.get("items", []):
                out.append({"title": clean_text(it.get("title","")), "url": it.get
if"nextPage"notin res.get("queries", {}): break
            start += len(res.get("items", []))
return out
except Exception as e:
```

```
        write_log(safe_name(query), f"Google 예외: {e}"); return []
```

```
defsearch_naver(query, n):
```

```
    cid, sec = os.getenv("NAVER_CLIENT_ID"), os.getenv("NAVER_CLIENT_SECRET")
```

```
ifnot cid ornot sec: write_log(safe_name(query), "Naver 키없음"); return []
    url = "https://openapi.naver.com/v1/search/webkr.json"
    headers = {"X-Naver-Client-Id": cid, "X-Naver-Client-Secret": sec, "User-Agent
    params = {"query": query, "display": min(100, n), "start": 1}
try:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 153/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
        r = requests.get(url, headers=headers, params=params, timeout=REQUEST_TIM
```

```
if r.status_code != 200: write_log(safe_name(query), f"Naver 오류{r.statu
        js = r.json()
return [{"title": clean_text(i.get("title","")), "url": i.get("link",""),
except Exception as e:
```

```
        write_log(safe_name(query), f"Naver 예외: {e}"); return []
```

```
defsearch_ddg(query, n):
try:
from duckduckgo_search import DDGS
        out=[]
with DDGS() as ddgs:
for r in ddgs.text(query, region="kr-kr", safesearch="off", max_result
                out.append({"title": clean_text(r.get("title","")), "url": r.get(
return out
except Exception as e:
        write_log(safe_name(query), f"DDG 예외: {e}"); return []
```

```
# ── 저장/병합 ────────────────────────────────────────────
```

```
defsave_rows(path, rows):
```

```
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
withopen(path, "w", encoding="utf-8-sig", newline="") as f:
```

```
        w = csv.DictWriter(f, fieldnames=headers); w.writeheader()
for r in rows: w.writerow(r)
```

```
defengine_to_csv(keyword, engine_name, rows, fetch_desc, query):
    fn = os.path.join(OUT_DIR, f"{keyword}_{engine_name}.csv")
    out=[]; seen=set()
for i, it inenumerate(rows, start=1):
        u = (it.get("url") or"").strip()
ifnot u or u in seen: continue
        seen.add(u)
        out.append({
"rank": i, "title": it.get("title",""), "url": u, "source": engine_nam
"query": query, "fetched_at_kst": now_kst(),
"description": fetch_description(u) if fetch_desc else"",
"domain": domain_of(u)
        })
    save_rows(fn, out)
return out
```

```
defmerge_and_save(keyword, query, parts):
    merged=[]; seen=set()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 154/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
for block in parts:
for r in block:
            u=r["url"]
ifnot u or u in seen: continue
            seen.add(u); merged.append(r)
for i,r inenumerate(merged, start=1): r["rank"]=i
    all_fn = os.path.join(OUT_DIR, f"{keyword}_all.csv")
    save_rows(all_fn, merged)
# TOP
    dom={}
for r in merged:
        d=r.get("domain","");
ifnot d: continue
if d notin dom: dom[d]={"count":0,"best":10**9}
        dom[d]["count"]+=1; dom[d]["best"]=min(dom[d]["best"], r["rank"])
    tops = [{"domain":k,"count":v["count"],"best_rank":v["best"]} for k,v in dom.
    tops.sort(key=lambda x:(-x["count"], x["best_rank"]))
    top_fn = os.path.join(OUT_DIR, f"TOP_{keyword}_all.csv")
withopen(top_fn, "w", encoding="utf-8-sig", newline="") as f:
        w=csv.DictWriter(f, fieldnames=["domain","count","best_rank"]); w.writehea
for t in tops: w.writerow(t)
    write_log(keyword, f"병합저장{len(merged)}개")
return merged, tops
```