                w.writerow(header)
            w.writerow([now_kst(), event, detail])
except Exception as e:
print("로그기록실패:", e)
```

```
defsafe_session():
```

```
    s = requests.Session()
```

```
    s.headers.update({"User-Agent": USER_AGENT})
    retries = Retry(total=3, backoff_factor=0.8,
```

```
500,502,503,504))
```

```
    s.mount("https://", HTTPAdapter(max_retries=retries))
    s.mount("http://", HTTPAdapter(max_retries=retries))
return s
```

```
defdomain_of(url):
```

```
try:
return urlparse(url).netloc.lower().replace("www.", "")
except:
return""
```

```
defclean_text(s):
return" ".join(s.split()).strip()
```

```
# ----- 검색엔진함수들 -----
```

```
defsearch_duckduckgo_html(query, max_results=50, session=None):
"""DuckDuckGo HTML 파싱(라이브러리없을때안전한대체)."""
    s = session or safe_session()
    results = []
    base = "https://html.duckduckgo.com/html/"
    start = 0
whilelen(results) < max_results:
        data = {"q": query, "s": start}
try:
            r = s.post(base, data=data, timeout=REQUEST_TIMEOUT)
except Exception as e:
            write_log("ddg_html_error", str(e))
break
if r.status_code != 200:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 17/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
            write_log("ddg_html_status", f"{r.status_code}")
break
        soup = BeautifulSoup(r.text, "html.parser")
        anchors = soup.select("a.result__a")
ifnot anchors:
break
for a in anchors:
            href = a.get("href")
            title = clean_text(a.get_text(" ", strip=True) or"")
if href and title:
                results.append({"title": title, "url": href})
iflen(results) >= max_results:
break
        start += 50
        time.sleep(0.6)
return results
```

```
defsearch_duckduckgo_lib(query, max_results=50):
"""duckduckgo_search 라이브러리이용 (설치권장)."""
ifnot HAS_DDG_LIB:
return []
try:
        out = ddg(query, max_results=max_results)
        results = []
for item in out:
# ddg returns dicts with 'title','href','body' 등
            results.append({"title": item.get("title",""), "url": item.get("href"
return results
except Exception as e:
        write_log("ddg_lib_error", str(e))
return []
defsearch_google_api(query, max_results=10):
"""Google Custom Search API (환경변수필요)."""
    api_key = os.getenv("GOOGLE_API_KEY")
    cx = os.getenv("GOOGLE_CSE_CX")
ifnot api_key ornot cx:
raise RuntimeError("GOOGLE API 키또는 CX가설정되어있지않습니다.")
try:
from googleapiclient.discovery import build
        service = build("customsearch", "v1", developerKey=api_key)
        results = []
        start_index = 1
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 18/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
whilelen(results) < max_results:
```

```
            resp = service.cse().list(q=query, cx=cx, start=start_index, num=min(1
            items = resp.get("items", [])
for it in items:
                results.append({"title": it.get("title",""), "url": it.get("link"
if"nextPage"notin resp.get("queries", {}):
break
            start_index += len(items)
return results
except Exception as e:
        write_log("google_api_error", str(e))
raise
```

```
defsearch_naver_api(query, max_results=10):
    client_id = os.getenv("NAVER_CLIENT_ID")
    client_secret = os.getenv("NAVER_CLIENT_SECRET")
ifnot client_id ornot client_secret:
raise RuntimeError("NAVER API 키가설정되어있지않습니다.")
    url = "https://openapi.naver.com/v1/search/webkr.json"
    s = safe_session()
    headers = {"X-Naver-Client-Id": client_id, "X-Naver-Client-Secret": client_se
    params = {"query": query, "display": max_results}
    r = s.get(url, headers=headers, params=params, timeout=REQUEST_TIMEOUT)
if r.status_code != 200:
        write_log("naver_status", f"{r.status_code}")
raise RuntimeError("Naver API 오류")
    js = r.json()
    results = []
for it in js.get("items", []):
        results.append({"title": it.get("title",""), "url": it.get("link","")})
return results
```

```
defsearch_kakao_api(query, max_results=10):
    kakao_key = os.getenv("KAKAO_REST_KEY")
ifnot kakao_key:
raise RuntimeError("KAKAO REST KEY 없음")
    url = "https://dapi.kakao.com/v2/search/web"
    s = safe_session()
    headers = {"Authorization": f"KakaoAK {kakao_key}"}
    params = {"query": query, "size": max_results}
    r = s.get(url, headers=headers, params=params, timeout=REQUEST_TIMEOUT)
if r.status_code != 200:
        write_log("kakao_status", f"{r.status_code}")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 19/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
raise RuntimeError("Kakao API 오류")
    js = r.json()
    results = []
for it in js.get("documents", []):
        results.append({"title": it.get("title",""), "url": it.get("url","")})
return results
# ----- 설명(메타) 수집 -----
deffetch_description(url, session=None):
"""
각링크의 meta description 또는첫문단을시도해서가져옴.
실패시빈문자열반환(안정성우선)."""
    s = session or safe_session()
try:
        r = s.get(url, timeout=REQUEST_TIMEOUT)
if r.status_code != 200:
return""
        soup = BeautifulSoup(r.content, "lxml")
# meta description
        d = soup.find("meta", attrs={"name":"description"}) or soup.find("meta", a
if d and d.get("content"):
return clean_text(d.get("content"))
# 첫문단
        p = soup.find("p")
if p:
return clean_text(p.get_text()[:300])
except Exception:
return""
return""
# ----- 저장함수 -----
defsave_results(rows):
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
try:
withopen(OUT_CSV, "w", newline="", encoding="utf-8-sig") as f:
            w = csv.DictWriter(f, fieldnames=headers)
            w.writeheader()
for r in rows:
                w.writerow(r)
except Exception as e:
        write_log("save_error", str(e))
raise
```

```
defsave_top(rows):
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 20/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# 도메인집계
```

```
    agg = {}
for r in rows:
        d = r.get("domain","")
ifnot d: continue
if d notin agg:
            agg[d] = {"count":0, "best_rank": math.inf, "examples":[]}
        agg[d]["count"] += 1
        agg[d]["best_rank"] = min(agg[d]["best_rank"], r.get("rank",9999))
iflen(agg[d]["examples"]) < 3:
            agg[d]["examples"].append({"rank": r.get("rank"), "title": r.get("tit
    rows_out = []
for d,v in agg.items():
        rows_out.append({
"domain": d,
"count": v["count"],
"best_rank": v["best_rank"],
"examples": " | ".join([f"{e['rank']}:{e['title']}" for e in v["examp
        })
    rows_out.sort(key=lambda x: (-x["count"], x["best_rank"]))
# 저장
try:
withopen(TOP_CSV, "w", newline="", encoding="utf-8-sig") as f:
            w = csv.DictWriter(f, fieldnames=["domain","count","best_rank","examp
            w.writeheader()
for r in rows_out:
                w.writerow(r)
except Exception as e:
        write_log("save_top_error", str(e))
```

```
# ----- 메인실행흐름 -----
```

```
defcollect_links(query, max_n=50, provider=DEFAULT_PROVIDER, fetch_desc=False):
    write_log("collect_start", f"q={query},n={max_n},prov={provider},desc={fetch_d
    rows = []
    session = safe_session()
try:
if provider == "duckduckgo":
if HAS_DDG_LIB:
                hits = search_duckduckgo_lib(query, max_n)
else:
                hits = search_duckduckgo_html(query, max_n, session=session)
elif provider == "google":
            hits = search_google_api(query, max_n)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 21/169 