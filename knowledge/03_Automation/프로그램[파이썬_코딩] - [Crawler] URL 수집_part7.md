ifnot anchors: break
for a in anchors:
            href = a.get("href"); title = clean_text(a.get_text(" ", strip=True))
if href and title:
                results.append({"title": title, "url": href})
iflen(results) >= max_results: break
        start += 50; time.sleep(0.6)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 31/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
return results
```

```
defsearch_duckduckgo_lib(query, max_results=50):
ifnot HAS_DDG_LIB: return []
try:
        out = ddg(query, max_results=max_results) or []
return [{"title": clean_text(i.get("title","")), "url": i.get("href",""
except Exception as e:
        write_log("ddg_lib_error", str(e)); return []
```

```
defsearch_google_api(query, max_results=10):
    api_key, cx = os.getenv("GOOGLE_API_KEY"), os.getenv("GOOGLE_CSE_CX")
ifnot api_key ornot cx: raise RuntimeError("구글키또는 CX 없음")
try:
from googleapiclient.discovery import build
        service = build("customsearch", "v1", developerKey=api_key)
        results, start_index = [], 1
whilelen(results) < max_results:
            resp = service.cse().list(q=query, cx=cx, start=start_index, num=min(1
            items = resp.get("items", [])
for it in items:
                results.append({"title": clean_text(it.get("title","")), "url": it
if"nextPage"notin resp.get("queries", {}): break
            start_index += len(items)
return results
except Exception as e:
        write_log("google_api_error", str(e)); raise
```

```
defsearch_naver_api(query, max_results=10):
```

```
    cid, csec = os.getenv("NAVER_CLIENT_ID"), os.getenv("NAVER_CLIENT_SECRET")
ifnot cid ornot csec: raise RuntimeError("네이버키없음")
```

```
    url = "https://openapi.naver.com/v1/search/webkr.json"
```

```
    s = safe_session()
    r = s.get(url, headers={"X-Naver-Client-Id": cid, "X-Naver-Client-Secret": cse
              params={"query": query, "display": max_results}, timeout=REQUEST_TIM
if r.status_code != 200:
        write_log("naver_status", f"{r.status_code}"); raise RuntimeError("네이버
    js = r.json()
return [{"title": clean_text(i.get("title","")), "url": i.get("link","")} for
```

```
# ---- 설명수집 ----
deffetch_description(url, session=None):
    s = session or safe_session()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 32/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
try:
```

```
        r = s.get(url, timeout=REQUEST_TIMEOUT)
if r.status_code != 200: return""
        soup = BeautifulSoup(r.content, "lxml")
        meta = soup.find("meta", attrs={"name":"description"}) or soup.find("meta
if meta and meta.get("content"): return clean_text(meta.get("content"))
        p = soup.find("p")
return clean_text(p.get_text()[:300]) if p else""
except Exception:
return""
# ---- 저장 ----
defsave_results(rows):
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
withopen(OUT_CSV, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=headers); w.writeheader()
for r in rows: w.writerow(r)
```

```
defsave_top(rows):
    agg = {}
for r in rows:
        d = r.get("domain","");
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
        w.writeheader(); [w.writerow(r) for r in out]
```

```
# ---- 수집 ----
```

```
defcollect_links(query, max_n=50, provider=DEFAULT_PROVIDER, fetch_desc=False):
    write_log("collect_start", f"q={query},n={max_n},prov={provider},desc={fetch_d
    session = safe_session()
```

```
if provider == "duckduckgo":
        hits = search_duckduckgo_lib(query, max_n) or search_duckduckgo_html(query
elif provider == "google":
        hits = search_google_api(query, max_n)
elif provider == "naver":
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 33/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
        hits = search_naver_api(query, max_n)
else:
raise RuntimeError("지원하지않는 provider")
    seen, rows, rank = set(), [], 0
for it in hits:
        u = (it.get("url") or"").strip(); t = clean_text(it.get("title",""))
ifnot u or u in seen: continue
        seen.add(u); rank += 1
        desc = fetch_description(u, session=session) if fetch_desc else""
        rows.append({"rank": rank, "title": t, "url": u, "source": provider,
"query": query, "fetched_at_kst": now_kst(),
"description": desc, "domain": domain_of(u)})
if rank >= max_n: break
return rows
defmain():
print("검색어를입력하세요:")
    q = input("> ").strip()
ifnot q: print("검색어없음. 종료."); return
try:
print("최대개수(기본 50):"); n = int(input("> ").strip() or"50")
        n = max(1, min(MAX_SAFE_RESULTS, n))
except: n = 50
print("엔진선택: duckduckgo / google / naver (기본 duckduckgo)")
    prov = (input("> ").strip().lower() or"duckduckgo")
if prov notin ("duckduckgo","google","naver"):
print("지원하지않는엔진. duckduckgo 사용"); prov = "duckduckgo"
print("설명(meta)도모을까요? (y/N)")
    fetch_desc = (input("> ").strip().lower() == "y")
    write_log("run_request", f"q={q},n={n},prov={prov},desc={fetch_desc}")
try:
        rows = collect_links(q, n, prov, fetch_desc)
ifnot rows:
"
print(결과없음. 검색어를바꿔보세요."); write_log("no_results", q); re
        save_results(rows); save_top(rows)
print(f"완료: {len(rows)}개저장 → {OUT_CSV}, {TOP_CSV}")
        write_log("run_done", f"saved={len(rows)}")
except Exception as e:
print("오류:", str(e)); print("LOG_Search.csv를확인하세요.")
        write_log("run_exception", traceback.format_exc())
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

34/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
if __name__ == "__main__":
    main()
```

## **4) 실행** 

PowerShell에서 실행해. 

nginx 

코드 복사 

```
python Search.py
```

## 예시 입력: 

- 검색어: `연기 커뮤니티` 

- 개수: `50` 

- 엔진: `duckduckgo` (키 없이 확인) 

- 설명 수집: `n` (속도 빠름) 

## **5) 확인** 

폴더에 **Search.csv** , **TOP_Search.csv** , **LOG_Search.csv** 가 보여야 해. 

## **6) 문제 시 빠른 점검** 

- “구글 키 없음”이면 키 넣고 창 재시작. 

- “네이버 키 없음”이면 ID/SECRET 확인. 

- 결과 0개면 검색어를 넓혀봐. 

다음 단계 지시해 줘. 

**==> picture [489 x 39] intentionally omitted <==**

**==> picture [289 x 140] intentionally omitted <==**

**==> picture [188 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
구글 클라우드 어떻게 쓴다고?<br>**----- End of picture text -----**<br>


## 12s 동안 생각함 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

35/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

형, **구글 클라우드로 구글 검색 API 쓰는 법** 을 초간단 순서로 줄게. 

(파일명은 계속 **Search.py / Search.csv** 그대로) 

## **1) 가입·프로젝트** 

**1.** 화면에서 **무료로 시작하기** 누르기 → 로그인. 

**2.** 상단 프로젝트 선택 → **새 프로젝트 만들기** . 

## **2) API 켜기 (Custom Search)** 

**1.** 좌측 **API 및 서비스 → 라이브러리** . 

**2. Custom Search API** 찾기 → **사용 설정** (Enable). 

## **3) API 키 만들기** 

**1. API 및 서비스 → 사용자 인증 정보** . 

**2. + 사용자 인증 정보 → API 키** → 복사. 

**3.** (권장) **키 제한** 클릭 → **HTTP referrer 제한** or **API 제한** 설정. 

## **4) Programmable Search Engine(CSE) 만들기** 

**1.** 새 탭에서 **Programmable Search Engine** 열기. 

**2. 새 검색 엔진 만들기** → “검색할 사이트”는 **전체 웹** 선택(또는 기본 허용). 
