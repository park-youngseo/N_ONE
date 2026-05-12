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
```

```
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
# ───────────── Serper.dev (Google 대체) ───────────────
```

```
defsearch_serper(query, n):
```

```
    api_key = os.getenv("SERPER_API_KEY")
```

```
ifnot api_key:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 105/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

`raise RuntimeError( "Serper.dev API 키가 없습니다.\n" "` 👉 `https://serper.dev 에서 가입 후 아래 명령 입력:\n" "   setx SERPER_API_KEY \"발급받은_API키\"\n" "PowerShell을 닫고 다시 열어야 적용됩니다." ) url = "https://google.serper.dev/search" headers = {"X-API-KEY": api_key, "Content-Type": "application/json"} data = {"q": query, "num": n} try:` 

```
        r = requests.post(url, headers=headers, json=data, timeout=REQUEST_TIMEOUT
if r.status_code != 200:
raise RuntimeError(f"Serper API 오류 ({r.status_code})")
        js = r.json()
return [{"title": clean_text(i.get("title","")), "url": i.get("link",""),
for i in js.get("organic", [])]
except Exception as e:
print("Serper 오류:", e)
return []
```

```
# ───────────── Naver 검색 + 키검증 ───────────────
```

```
defcheck_naver_keys():
```

`cid, csec = os.getenv("NAVER_CLIENT_ID"), os.getenv("NAVER_CLIENT_SECRET") if not cid or not csec: print("\n[` 🚫 `네이버 API 키가 없습니다]") print("` 1️⃣ `https://developers.naver.com/apps 에 접속하세요.") print("` 2️⃣ `'애플리케이션 등록' → 검색(OpenAPI) 선택 → 환경: Web / URL: http:/ print("` 3️⃣ `발급 후 아래 명령 입력:\n" "   setx NAVER_CLIENT_ID \"발급받은_ID\"\n" "   setx NAVER_CLIENT_SECRET \"발급받은_SECRET\"") print("` 4️⃣ `PowerShell을 닫고 새로 열면 적용됩니다.\n") return False return True` 

```
defsearch_naver(query, n):
```

```
ifnot check_naver_keys(): return []
```

```
    cid, csec = os.getenv("NAVER_CLIENT_ID"), os.getenv("NAVER_CLIENT_SECRET")
    url = "https://openapi.naver.com/v1/search/webkr.json"
    headers = {
"X-Naver-Client-Id": cid,
"X-Naver-Client-Secret": csec,
"User-Agent": USER_AGENT
    }
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 106/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

`params = {"query": query, "display": min(100, n), "start": 1} try: r = requests.get(url, headers=headers, params=params, timeout=REQUEST_TIM if r.status_code != 200: print(f"[` ⚠ `네이버 API 오류] 코드: {r.status_code}") if r.status_code in (401,403): print("→ 키가 만료되었거나 잘못된 조합입니다. 재발급 후 다시 시도하세요 "   https://developers.naver.com/apps 에서 '새 애플리케이션 등 return [] js = r.json() return [{"title": clean_text(i.get("title","")), "url": i.get("link",""), for i in js.get("items", [])] except Exception as e: print("Naver 오류:", e) return []` _`# ─────────────`_ `데이터 처리` _`───────────────`_ `def fetch_description(url): try: s = safe_session() r = s.get(url, timeout=8) if r.status_code != 200: return "" soup = BeautifulSoup(r.text, "lxml") meta = soup.find("meta", attrs={"name":"description"}) or soup.find("meta if meta and meta.get("content"): return clean_text(meta["content"]) p = soup.find("p") return clean_text(p.get_text()[:200]) if p else "" except: return ""` 

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
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 107/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
defsave_csv(rows, path):
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=headers)
        w.writeheader()
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
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=["domain","count","best_rank"])
        w.writeheader()
for r in out: w.writerow(r)
# ───────────── 메인 ───────────────
defmain():
print("검색어를입력하세요:")
    q = input("> ").strip()
ifnot q: print("검색어없음. 종료."); return
    q_clean = safe_name(q)
try:
print("최대개수(기본 50):")
        n = int(input("> ").strip() or"50")
except: n = 50
```

```
print("설명(meta)도모을까요? (y/N)")
    fetch_desc = (input("> ").strip().lower() == "y")
    OUT_CSV = os.path.join(OUT_DIR, f"{q_clean}_mix.csv")
    LOG_CSV = os.path.join(OUT_DIR, f"LOG_{q_clean}_mix.csv")
    TOP_CSV = os.path.join(OUT_DIR, f"TOP_{q_clean}_mix.csv")
    write_log(LOG_CSV, "run_start", f"query={q}, n={n}")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

108/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

`try: s1 = search_serper(q, n) s2 = search_naver(q, n) rows = merge_and_clean(s1, s2) if fetch_desc: for r in rows: r["description"] = fetch_description(r["url"]) else: for r in rows: r["description"] = "" if not rows: print("결과 없음. API키 또는 검색어 확인.") write_log(LOG_CSV, "no_results", q) return save_csv(rows, OUT_CSV) save_top(rows, TOP_CSV) print(f"\n` ✅ `완료: {len(rows)}개 저장 → {OUT_CSV}") write_log(LOG_CSV, "done", f"saved={len(rows)}") except Exception as e: print("오류:", str(e)) write_log(LOG_CSV, "error", traceback.format_exc()) if __name__ == "__main__": main()` 

## 🧭 **작동 방식** 

실행 시 자동으로 환경변수 점검 

`NAVER_CLIENT_ID` 또는 `SECRET` 이 없으면 

- 👉 바로 친절한 안내문이 콘솔에 출력됨 

- 잘못된 키(401, 403)는 즉시 감지 후 

- “재발급 방법 링크”와 예시 명령어 표시 

## 🚀 **실행 순서** 

powershell 

코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
```

```
python .\Search.py
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

109/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## ⚙ **출력 구조** 

csharp 

코드 복사 

```
out/연기오디션_mix.csv
out/TOP_연기오디션_mix.csv
out/LOG_연기오디션_mix.csv
```

이제 더 이상 키가 없거나 틀려도 

형이 오류코드를 일일이 볼 필요 없이, 

자동으로 “어디서 재발급 / 어떻게 등록”까지 안내돼. 

형, 이거 적용하고 한 번 실행해봐. 

만약 안내문이 뜨면 그대로 따라하면 되고, 

정상 작동되면 `out` 폴더 캡처 한 번만 보여줘 — 

그다음엔 “자동 키워드 추출 + 요약 분석 모듈” 바로 붙일게. 

**==> picture [489 x 40] intentionally omitted <==**

아무런 결과도 나오지 않아 예전처럼 내가 선택할수 있게 만들어줘 

아무래도 꼬이는것 같다 

좋아 형 👌 지금 상황 정확히 짚었어. 

v2.7은 내부에서 자동으로 엔진을 선택(Serper+Naver 병합)하도록 고정돼 있었지? 이게 편하긴 한데, 

실행 환경이 꼬였거나 키가 부분적으로 등록돼 있을 땐 

**“조용히 실패”** 하면서 결과가 안 나오는 현상이 생겨. 

그래서 내가 지금부터 
