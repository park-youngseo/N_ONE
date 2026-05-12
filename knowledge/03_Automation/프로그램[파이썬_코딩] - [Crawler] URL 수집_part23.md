✅ Serper / Naver / Google / DuckDuckGo 중 원하는 걸 **직접 선택** 

✅ 각각 별도로 테스트 가능 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 110/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

✅ 결과파일 이름 자동 표시 

✅ 꼬임 복구 기능 포함 

이걸 전부 합친 **v2.8 “선택형 완전버전”** 으로 만들어줄게. 

엔진 구조 다시 열어두고, 실행할 때 

코드 복사 

```
엔진을선택하세요: serper / naver / google / duckduckgo / mix
```

이렇게 직접 입력할 수 있어. 

## 👇 **이대로 실행만 하면 되는 코드 (Search.py v2.8)** 

python 

코드 복사 

- _`# N·ONE Link Collector v2.8 / 2025-11-04 19:10 KST`_ 

- _`# BASE=Search (OUT: out/[`_ `검색어` _`]_[`_ `엔진` _`].csv)`_ 

- _`#`_ `특징` _`: Serper·Naver·Google·DuckDuckGo·Mix`_ `선택형` _`/`_ `자동검증 포함` 

```
import os, csv, time, re, requests, traceback
from urllib.parse import urlparse
```

```
from bs4 import BeautifulSoup
from requests.adapters import HTTPAdapter
```

```
from urllib3.util.retry import Retry
```

```
BASE_DIR  = os.path.dirname(os.path.abspath(__file__))
```

```
OUT_DIR   = os.path.join(BASE_DIR, "out")
```

```
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

## `def safe_session():` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 111/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
    s = requests.Session()
```

```
    s.headers.update({"User-Agent": USER_AGENT})
    retries = Retry(total=3, backoff_factor=0.5, status_forcelist=[500,502,503,504
    s.mount("https://", HTTPAdapter(max_retries=retries))
return s
```

_`# ────────────────────────────────────────────────`_ `def write_log(path, event, detail=""): new_file = not os.path.exists(path) with open(path, "a", newline="", encoding="utf-8") as f: w = csv.writer(f) if new_file: w.writerow(["time_kst","event","detail"]) w.writerow([now_kst(), event, detail])` _`# ───────────── Serper.dev (Google`_ `대체` _`) ─────────────`_ `def search_serper(query, n): api_key = os.getenv("SERPER_API_KEY") if not api_key: print("\n[` ❌ `Serper.dev 키 없음]") print("→ https://serper.dev 가입 후 아래 입력:") print('   setx SERPER_API_KEY "발급받은_API키"\n') return [] url = "https://google.serper.dev/search" headers = {"X-API-KEY": api_key, "Content-Type": "application/json"} data = {"q": query, "num": n} try: r = requests.post(url, headers=headers, json=data, timeout=REQUEST_TIMEOUT if r.status_code != 200: print(f"Serper API 오류: {r.status_code}") return [] js = r.json() return [{"title": clean_text(i.get("title","")), "url": i.get("link",""), for i in js.get("organic", [])] except Exception as e: print("Serper 오류:", e) return []` 

_`# ───────────── Naver API ─────────────`_ `def search_naver(query, n): cid, csec = os.getenv("NAVER_CLIENT_ID"), os.getenv("NAVER_CLIENT_SECRET") if not cid or not csec: print("\n[` ❌ `네이버 API 키 없음]") print("→ https://developers.naver.com/apps 에서 새 앱 등록:")` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 112/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

`print("   setx NAVER_CLIENT_ID \"ID\"") print("   setx NAVER_CLIENT_SECRET \"SECRET\"\n") return [] url = "https://openapi.naver.com/v1/search/webkr.json" headers = { "X-Naver-Client-Id": cid, "X-Naver-Client-Secret": csec, "User-Agent": USER_AGENT } params = {"query": query, "display": min(100, n), "start": 1} try: r = requests.get(url, headers=headers, params=params, timeout=REQUEST_TIM if r.status_code != 200: print(f"[` ⚠ `Naver API 오류] 코드: {r.status_code}") return [] js = r.json() return [{"title": clean_text(i.get("title","")), "url": i.get("link",""), for i in js.get("items", [])] except Exception as e: print("Naver 오류:", e) return []` _`# ───────────── DuckDuckGo ─────────────`_ `def search_duckduckgo(query, n): from duckduckgo_search import DDGS try: with DDGS() as ddgs: out = [] for r in ddgs.text(query, region="kr-kr", safesearch="off", max_result out.append({"title": clean_text(r.get("title","")), "url": r.get( return out except Exception as e: print("DuckDuckGo 오류:", e) return []` 

_`# ───────────── Google API ─────────────`_ `def search_google(query, n): key, cx = os.getenv("GOOGLE_API_KEY"), os.getenv("GOOGLE_CSE_CX") if not key or not cx: print("\n[` ❌ `Google API 키 또는 CX 없음]") print("→ Google Cloud Console에서 키 발급 후 아래 입력:") print('   setx GOOGLE_API_KEY "키"') print('   setx GOOGLE_CSE_CX "cx코드"\n') return []` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 113/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
from googleapiclient.discovery import build
try:
        service = build("customsearch", "v1", developerKey=key)
        results, start = [], 1
whilelen(results) < n:
            res = service.cse().list(q=query, cx=cx, start=start, num=min(10, n-le
for i in res.get("items", []):
                results.append({"title": clean_text(i["title"]), "url": i["link"]
if"nextPage"notin res.get("queries", {}): break
            start += len(res.get("items", []))
return results
except Exception as e:
print("Google 오류:", e)
return []
# ───────────── 보조기능 ─────────────
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
            seen.add(u); merged.append(it)
for i,r inenumerate(merged, start=1):
        r["rank"]=i; r["domain"]=domain_of(r["url"]); r["fetched_at_kst"]=now_kst
return merged
```

```
defsave_csv(rows, path):
    headers = ["rank","title","url","source","query","fetched_at_kst","descriptio
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=headers)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 114/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
        w.writeheader(); [w.writerow(r) for r in rows]
```

```
defsave_top(rows, path):
    agg = {}
for r in rows:
        d=r.get("domain","");
ifnot d: continue
if d notin agg: agg[d]={"count":0,"best":999}
        agg[d]["count"]+=1; agg[d]["best"]=min(agg[d]["best"],r["rank"])
    out=[{"domain":k,"count":v["count"],"best_rank":v["best"]} for k,v in agg.item
    out.sort(key=lambda x:(-x["count"],x["best_rank"]))
withopen(path,"w",newline="",encoding="utf-8-sig") as f:
        w=csv.DictWriter(f,fieldnames=["domain","count","best_rank"])
        w.writeheader(); [w.writerow(r) for r in out]
```

```
# ───────────── 메인 ─────────────
defmain():
print("검색어를입력하세요:")
    q = input("> ").strip()
ifnot q: return
    q_clean = safe_name(q)
try:
print("검색결과개수 (기본 50):")
        n = int(input("> ").strip() or"50")
except: n=50
```

`print("엔진을 선택하세요 (serper / naver / google / duckduckgo / mix):") prov = input("> ").strip().lower() or "serper" print("설명(meta)도 모을까요? (y/N):") fetch_desc = (input("> ").strip().lower() == "y") OUT_CSV = os.path.join(OUT_DIR, f"{q_clean}_{prov}.csv") TOP_CSV = os.path.join(OUT_DIR, f"TOP_{q_clean}_{prov}.csv") LOG_CSV = os.path.join(OUT_DIR, f"LOG_{q_clean}_{prov}.csv") print(f"\n[` 🔍 `검색 시작] {q} ({prov}) ...\n") write_log(LOG_CSV,"run_start",f"query={q},engine={prov}") try: if prov=="serper": rows=search_serper(q,n) elif prov=="naver": rows=search_naver(q,n)` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

115/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

`elif prov=="google": rows=search_google(q,n) elif prov=="duckduckgo": rows=search_duckduckgo(q,n) elif prov=="mix": rows=merge_and_clean(search_serper(q,n),search_naver(q,n)) else: print("지원하지 않는 엔진"); return if fetch_desc: for r in rows: r["description"]=fetch_description(r["url"]) else: for r in rows: r["description"]="" if not rows: print("` ❌ `결과 없음. API키 또는 검색어 확인.") write_log(LOG_CSV,"no_results",q) return save_csv(rows,OUT_CSV) save_top(rows,TOP_CSV) print(f"\n` ✅ `완료: {len(rows)}개 저장 → {OUT_CSV}") write_log(LOG_CSV,"done",f"saved={len(rows)}")` 

```
except Exception as e:
print("오류:",e)
        write_log(LOG_CSV,"error",traceback.format_exc())
if __name__=="__main__":
    main()
```

## 💡 **사용법 요약** 

powershell 코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
python .\Search.py
```

## 입력 예시 👇 

makefile 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 116/169 

25. 11. 30. 오후 1:05 

 

오후 1:05후 1:05 1:05 프로그램[파이썬/코딩] - [Crawler] URL 수집 `검색어: 연기 오디션 개수: 50 엔진: mix 설명: n` → 결과 csharp 코드 복사 `out/연기오디션_mix.csv out/TOP_연기오디션_mix.csv out/LOG_연기오디션_mix.csv` 

## ✅ **형이 직접 선택 가능** 

**입력값 기능** serper 구글 기반 (Serper.dev API) naver 네이버 검색 API google 구글 공식 Custom Search duckduckgo 덕덕고 API mix Serper + Naver 병합 

## 이제 형이 어떤 엔진이든 직접 골라서 테스트할 수 있어. 

만약 이 버전에서도 결과가 비면, 그 엔진 키 쪽만 점검하면 돼 — 다른 엔진으로는 그대로 수집 가능하거든. 

지금 바로 실행해봐. 

정상 출력되면 `out/` 폴더 스크린샷 한 번 보여줘 — 그다음엔 “자동 요약 + 키워드 정리 모듈” 붙일게. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 117/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 
