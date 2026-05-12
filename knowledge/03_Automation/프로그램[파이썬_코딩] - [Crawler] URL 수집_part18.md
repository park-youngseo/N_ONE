```

```
    s = safe_session()
```

```
    headers = {
```

```
"X-Naver-Client-Id": cid, "X-Naver-Client-Secret": csec,
```

```
"User-Agent": USER_AGENT, "Accept": "application/json",
    }
```

```
    params = {"query": query, "display": min(100, n), "start": 1}
```

```
    r = s.get(url, headers=headers, params=params, timeout=REQUEST_TIMEOUT)
```

```
if r.status_code != 200:
```

```
raise RuntimeError(f"네이버 API 오류 ({r.status_code})")
```

```
    js = r.json()
```

```
return [{"title": clean_text(i.get("title","")), "url": i.get("link","")} for
```

```
# ── 기타 ───────────────────────────────────────────────────
```

```
deffetch_description(url, s):
```

```
try:
```

```
        r = s.get(url, timeout=REQUEST_TIMEOUT)
```

```
if r.status_code != 200: return""
```

```
        soup = BeautifulSoup(r.content, "lxml")
```

```
        meta = soup.find("meta", attrs={"name":"description"}) or soup.find("meta
```

```
if meta and meta.get("content"): return clean_text(meta.get("content"))
        p = soup.find("p")
```

```
return clean_text(p.get_text()[:300]) if p else""
```

```
except Exception: return""
```

```
defsave_results(rows, path):
```

```
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
```

```
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
```

```
        w = csv.DictWriter(f, fieldnames=headers)
```

```
        w.writeheader()
for r in rows: w.writerow(r)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 85/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

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
        w.writeheader()
for r in out: w.writerow(r)
```

```
# ── 수집 ───────────────────────────────────────────────────
```

```
defcollect_links(query, n, prov, fetch_desc=False):
    s = safe_session()
if prov == "duckduckgo": hits = search_duckduckgo(query, n)
elif prov == "google": hits = search_google(query, n)
elif prov == "naver": hits = search_naver(query, n)
elif prov == "brave": hits = search_brave(query, n)
else: raise RuntimeError("지원하지않는 provider")
```

```
    seen, rows, rank = set(), [], 0
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
```

```
# ── 메인 ───────────────────────────────────────────────────
defmain():
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

86/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
print("검색어를입력하세요:")
    q = input("> ").strip()
ifnot q: print("검색어없음. 종료."); return
    q_clean = safe_name(q)
try:
print("최대개수(기본 50):")
        n = int(input("> ").strip() or"50")
        n = max(1, min(MAX_SAFE_RESULTS, n))
except: n = 50
```

`print("엔진 선택: duckduckgo / brave / google / naver (기본 duckduckgo)") prov = (input("> ").strip().lower() or "duckduckgo") if prov not in ("duckduckgo","brave","google","naver"): print("지원하지 않는 엔진. duckduckgo 사용"); prov = "duckduckgo" print("설명(meta)도 모을까요? (y/N)") fetch_desc = (input("> ").strip().lower() == "y") OUT_CSV = os.path.join(OUT_DIR, f"{q_clean}_{prov}.csv") LOG_CSV = os.path.join(OUT_DIR, f"LOG_{q_clean}_{prov}.csv") TOP_CSV = os.path.join(OUT_DIR, f"TOP_{q_clean}_{prov}.csv") write_log(LOG_CSV, "run_request", f"q={q},n={n},prov={prov},desc={fetch_desc} try: rows = collect_links(q, n, prov, fetch_desc) save_results(rows, OUT_CSV) save_top(rows, TOP_CSV) if not rows: " print( 결과 없음. 검색어/엔진을 바꿔보세요.") write_log(LOG_CSV, "no_results", q) return print(f"` ✅ `완료: {len(rows)}개 저장 → {OUT_CSV}") write_log(LOG_CSV, "run_done", f"saved={len(rows)}") except Exception as e: print("오류:", str(e)) print(f"자세한 내용은 {LOG_CSV} 확인") write_log(LOG_CSV, "run_exception", traceback.format_exc())` 

```
if __name__ == "__main__":
    main()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 87/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## ✅ **실행 예시** 

powershell 코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
python .\Search.py
```

## 입력 예: 

makefile 코드 복사 

```
검색어: 연기오디션
엔진: brave
```

## → 결과: 

csharp 코드 복사 

```
out/연기오디션_brave.csv
out/TOP_연기오디션_brave.csv
out/LOG_연기오디션_brave.csv
```

## 💡 **요약** 

|💡**요약**|||
|---|---|---|
|**엔진**|**방식**|**비고**|
|duckduckgo|공식 API (DDGS)|빠름, 안정적|
|brave|공식 무료 API|백업용, 해외결과 다양|
|google|Custom Search API|키 필요|
|naver|OpenAPI|키 필요|



## 이제 DuckDuckGo 막힘 없이 돌아갈 거야. 

형, 이 버전 돌려보고 `out` 폴더 결과 캡처 한 번 보여줄래? 

(정상 결과 개수랑 파일 구조 확인 후 다음 단계 — 자동 필터/요약 기능 붙여줄게.) 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

88/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

DuckDuckGo 오류: https://duckduckgo.com/ 202 Ratelimit 결과 없음. 검색어/엔진을 바꿔보세요. PS C:\Users\psm23\OneDrive\Desktop\Search> 

## 좋아, 형. 

이건 **DuckDuckGo가 아예 서버단에서 “202 Ratelimit” (요청 제한)** 으로 막은 거야. 즉, 프로그램 오류가 아니라 DDG 서버가 형의 IP를 “너무 자주 요청했다”고 판단한 상태. 

## 🔍 **원인 정리** 

- DDG는 크롤링이나 반복 요청을 “봇”으로 인식하면 **202 코드(빈 응답)** 을 줘. 이건 브라우저에서도 같은 검색어를 짧은 간격으로 20회 이상 반복해도 발생함. 

- `duckduckgo-search` 라이브러리 내부 요청도 결국 `duckduckgo.com/api` 라서 

- IP 단위로 막히면 회피 불가능. 

## ✅ **해결책 세 가지** 

## 🅰 **가장 확실한 방법 — 엔진 Brave로 전환** 

Brave Search는 공식 무료 API가 열려 있어서 IP 차단이 거의 없어. DuckDuckGo 대신 아래처럼 선택하면 돼 👇 

코드 복사 

```
엔진선택: brave
```

- → `out/검색어_brave.csv` 로 동일하게 저장돼. 

- → 속도도 비슷하고 결과 품질은 Google과 DuckDuckGo 중간 수준이야. 

## 🅱 **완전 복구 (DuckDuckGo 다시 쓰고 싶을 때)** 

- 1️⃣ PowerShell 관리자 권한으로 열기 

- 2️⃣ 아래 입력 

powershell 
