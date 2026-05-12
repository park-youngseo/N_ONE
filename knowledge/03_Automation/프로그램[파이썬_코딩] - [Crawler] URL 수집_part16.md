if r.status_code != 200: break
            soup = BeautifulSoup(r.text, "html.parser")
for a in soup.select("a.result__a"):
                href = a.get("href"); title = clean_text(a.get_text(" ", strip=Tr
if href and title: results.append({"title": title, "url": href})
iflen(results) >= n: break
            start += 30; time.sleep(1.2)
return results
except Exception: return []
```

```
defsearch_google(query, n):
```

```
    api_key, cx = os.getenv("GOOGLE_API_KEY"), os.getenv("GOOGLE_CSE_CX")
ifnot api_key ornot cx: raise RuntimeError("구글 API 키또는 CX가없습니다.")
from googleapiclient.discovery import build
    service = build("customsearch", "v1", developerKey=api_key)
    results, start_index = [], 1
whilelen(results) < n:
        resp = service.cse().list(q=query, cx=cx, start=start_index,
                                  num=min(10, n - len(results))).execute()
for it in resp.get("items", []):
            results.append({"title": clean_text(it.get("title","")), "url": it.get
if"nextPage"notin resp.get("queries", {}): break
        start_index += len(resp.get("items", []))
return results
```

```
defsearch_naver(query, n):
```

```
    cid, csec = os.getenv("NAVER_CLIENT_ID"), os.getenv("NAVER_CLIENT_SECRET")
ifnot cid ornot csec: raise RuntimeError("네이버 API 키가없습니다.")
    url = "https://openapi.naver.com/v1/search/webkr.json"
    s = safe_session()
    headers = {
```

```
"X-Naver-Client-Id": cid, "X-Naver-Client-Secret": csec,
"User-Agent": USER_AGENT, "Accept": "application/json",
```

```
    }
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 76/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
    params = {"query": query, "display": min(100, n), "start": 1}
```

```
    r = s.get(url, headers=headers, params=params, timeout=REQUEST_TIMEOUT)
```

```
if r.status_code != 200: raise RuntimeError("네이버 API 오류")
    js = r.json()
```

```
return [{"title": clean_text(i.get("title","")), "url": i.get("link","")} for
```

```
# ── 기타함수 ─────────────────────────────────────────────
deffetch_description(url, s):
```

```
try:
        r = s.get(url, timeout=REQUEST_TIMEOUT)
if r.status_code != 200: return""
        soup = BeautifulSoup(r.content, "lxml")
        meta = soup.find("meta", attrs={"name":"description"}) or soup.find("meta
if meta and meta.get("content"): return clean_text(meta.get("content"))
        p = soup.find("p")
return clean_text(p.get_text()[:300]) if p else""
except Exception: return""
```

```
defsave_results(rows, path):
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
```

```
        w = csv.DictWriter(f, fieldnames=headers); w.writeheader()
for r in rows: w.writerow(r)
```

```
defsave_top(rows, path):
    agg = {}
for r in rows:
        d = r.get("domain","")
ifnot d: continue
if d notin agg: agg[d] = {"count":0, "best_rank": 10**9, "ex":[]}
        agg[d]["count"] += 1
        agg[d]["best_rank"] = min(agg[d]["best_rank"], r.get("rank", 10**9))
iflen(agg[d]["ex"]) < 3: agg[d]["ex"].append(f"{r.get('rank')}:{r.get('t
    out = [{"domain": k, "count": v["count"], "best_rank": v["best_rank"], "examp
    out.sort(key=lambda x: (-x["count"], x["best_rank"]))
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=["domain","count","best_rank","examples"
for r in out: w.writerow(r)
```

```
# ── 메인 ───────────────────────────────────────────────────
```

```
defcollect_links(query, n, prov, fetch_desc=False):
    s = safe_session()
if prov == "duckduckgo": hits = search_duckduckgo(query, n, s)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

77/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
elif prov == "google": hits = search_google(query, n)
elif prov == "naver": hits = search_naver(query, n)
else: raise RuntimeError("지원하지않는 provider")
```

```
    rows, seen, rank = [], set(), 0
for it in hits:
        u = (it.get("url") or"").strip()
        t = clean_text(it.get("title",""))
ifnot u or u in seen: continue
        seen.add(u); rank += 1
        desc = fetch_description(u, s) if fetch_desc else""
        rows.append({
"rank": rank, "title": t, "url": u, "source": prov,
"query": query, "fetched_at_kst": now_kst(),
"description": desc, "domain": domain_of(u)
        })
if rank >= n: break
return rows
defmain():
print("검색어를입력하세요:")
    q = input("> ").strip()
ifnot q: print("검색어없음. 종료."); return
    q_clean = safe_name(q)
try:
print("최대개수(기본 50):")
        n = int(input("> ").strip() or"50")
        n = max(1, min(MAX_SAFE_RESULTS, n))
except: n = 50
print("엔진선택: duckduckgo / google / naver (기본 duckduckgo)")
    prov = (input("> ").strip().lower() or"duckduckgo")
if prov notin ("duckduckgo","google","naver"):
print("지원하지않는엔진. duckduckgo 사용"); prov = "duckduckgo"
print("설명(meta)도모을까요? (y/N)")
    fetch_desc = (input("> ").strip().lower() == "y")
    OUT_CSV = os.path.join(OUT_DIR, f"{q_clean}_{prov}.csv")
    LOG_CSV = os.path.join(OUT_DIR, f"LOG_{q_clean}_{prov}.csv")
    TOP_CSV = os.path.join(OUT_DIR, f"TOP_{q_clean}_{prov}.csv")
```

```
    write_log(LOG_CSV, "run_request", f"q={q},n={n},prov={prov},desc={fetch_desc}
```

```
try:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

78/169 

25. 11. 30. 오후 1:05 

 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

```
        rows = collect_links(q, n, prov, fetch_desc)
        save_results(rows, OUT_CSV)
        save_top(rows, TOP_CSV)
ifnot rows:
"
print(결과없음. 검색어/엔진을바꿔보세요.")
            write_log(LOG_CSV, "no_results", q)
return
print(f"완료: {len(rows)}개저장 → {OUT_CSV}")
        write_log(LOG_CSV, "run_done", f"saved={len(rows)}")
except Exception as e:
print("오류:", str(e))
print(f"자세한내용은{LOG_CSV}확인")
        write_log(LOG_CSV, "run_exception", traceback.format_exc())
if __name__ == "__main__":
    main()
```

## ✅ 실행 예시 

powershell 코드 복사 `cd "C:\Users\psm23\OneDrive\Desktop\Search" python .\Search.py` 

## 입력: 

makefile 코드 복사 

```
검색어: 연기오디션
엔진: google
```

## 📂 결과: 

csharp 코드 복사 `out/연기오디션_google.csv out/TOP_연기오디션_google.csv out/LOG_연기오디션_google.csv` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

79/169 

25. 11. 30. 오후 1:05 

 

**==> picture [147 x 9] intentionally omitted <==**

**----- Start of picture text -----**<br>
프로그램[파이썬/코딩] - [Crawler] URL 수집<br>**----- End of picture text -----**<br>


이제 파일 이름만 봐도 어떤 검색어·포털인지 한눈에 구분돼. 

실행해보고 `out` 폴더 화면 한 번만 캡처해서 보내줘, 정리용 필터 자동 스크립트 만들어 줄게. 

**==> picture [489 x 39] intentionally omitted <==**

duckduckgo 어느순간 부터 검색결과가 계속 0으로 나와 구글은 나오는데 

좋은 포인트 짚었어 형. 👇 

이건 코드 문제가 아니라 **DuckDuckGo 쪽 봇방어 정책** 이 최근 강화돼서 생긴 현상이야. 지금처럼 `html.duckduckgo.com/html/` 을 직접 긁는 방식은 예전엔 잘 됐는데, 최근엔 1~2회 이후 바로 **HTTP 202 / 빈 결과 / 봇 차단 페이지** 가 돌아와. 

## 🔍 **원인 요약** 

**DuckDuckGo HTML Search** 는 Cloudflare 보호 적용됨. 

일정 시간 이상 요청하거나 IP 패턴이 일정하면 “빈 응답(0개)”만 줌. 

- VPN·공용 IP·같은 User-Agent 반복 시 더 빨리 막힘. 

## ✅ **해결 방법 3가지 (선택형)** 

## 🅰 **라이브러리 모드 강제 고정 (추천)** 

`duckduckgo-search` 라이브러리 버전 6.2.2 이상이면 

새 API 엔드포인트( `api.duckduckgo.com` )를 써서 안정적이야. 

PowerShell에서 아래 한 줄 실행: 

powershell 코드 복사 

```
python -m pip install -U duckduckgo-search==6.2.2
```

그 다음, `Search.py` 에서 **이 부분만 교체** 👇 

python 코드 복사 

```
defsearch_duckduckgo(query, n, s):
```

```
try:
```
