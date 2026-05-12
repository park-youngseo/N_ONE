**1.** 설치해요. 

**2.** 실행해요. 

**3.** 결과를 봐요. 

## **B2 세부** 

## **1) 설치** 

nginx 코드 복사 

```
python -m pip install requests beautifulsoup4
```

## **2) 실행** 

nginx 코드 복사 `python Search.py` 

화면 안내에 따라 입력해요. 예: 검색어, 개수. 

## **3) 코드 (Search.py)** 

python 코드 복사 

```
# N·ONE Link Collector v1.0 / 2025-11-02 KST / 파일: Search.py
```

```
# BASE=Search  (CSV: Search.csv, LOG: LOG_Search.csv)
```

```
import csv, time, sys
```

```
from urllib.parse import urlencode
from datetime import datetime
from html import unescape
```

```
import requests
```

```
from bs4 import BeautifulSoup
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

7/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# ---- 설정값 ----
```

```
USER_AGENT = (
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) "
"Chrome/124.0 Safari/537.36"
)
BASE_NAME = "Search"# 파일명잠금
OUT_CSV = f"{BASE_NAME}.csv"             # 결과저장
LOG_CSV = f"LOG_{BASE_NAME}.csv"         # 로그저장
PROVIDER = "duckduckgo"# 검색소스
DDG_HTML = "https://html.duckduckgo.com/html/"
```

```
# ---- 공용함수 ----
defnow_kst_str():
# KST 고정표기
return time.strftime("%Y-%m-%d %H:%M", time.localtime())
```

```
defwrite_log(event, detail=""):
    header = ["time_kst","event","detail"]
    row = [now_kst_str(), event, detail]
    new_file = not os_path_exists(LOG_CSV)
withopen(LOG_CSV, "a", newline="", encoding="utf-8") as f:
        w = csv.writer(f)
if new_file: w.writerow(header)
        w.writerow(row)
```

```
defos_path_exists(path):
try:
import os
return os.path.exists(path)
except:
returnFalse
defclean_text(s):
return unescape(" ".join(s.split())).strip()
# ---- DuckDuckGo HTML 검색 ----
defsearch_duckduckgo(query, max_results=50, session=None):
    s = session or requests.Session()
    s.headers.update({"User-Agent": USER_AGENT})
    results = []
    start = 0
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

8/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
whilelen(results) < max_results:
        params = {"q": query, "s": start}
        resp = s.post(DDG_HTML, data=params, timeout=20)
if resp.status_code != 200:
break
        soup = BeautifulSoup(resp.text, "html.parser")
for a in soup.select("a.result__a"):
            url = a.get("href")
            title = clean_text(a.get_text(" ", strip=True))
ifnot url ornot title:
continue
            results.append({
"rank": len(results) + 1,
"title": title,
"url": url,
"source": PROVIDER,
"query": query,
"fetched_at_kst": now_kst_str(),
            })
iflen(results) >= max_results:
break
# 다음페이지추정이동
        start += 50
# 안전대기
        time.sleep(0.8)
# 더이상결과가없으면중단
ifnot soup.select("a.result__a"):
break
return results
```

```
defsave_csv(rows, path):
    headers = ["rank","title","url","source","query","fetched_at_kst"]
    new_file = not os_path_exists(path)
withopen(path, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.DictWriter(f, fieldnames=headers)
        w.writeheader()
for r in rows:
            w.writerow(r)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

9/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
defmain():
print("검색어를입력해요:")
    query = input("> ").strip()
ifnot query:
print("검색어가없어요.")
return
try:
"
print(최대개수를적어요. 예: 50")
        n_str = input("> ").strip() or"50"
        max_n = max(1, min(500, int(n_str)))
except:
        max_n = 50
    write_log("start", f"query={query}, max={max_n}")
try:
        sess = requests.Session()
        rows = search_duckduckgo(query, max_results=max_n, session=sess)
        save_csv(rows, OUT_CSV)
print(f"완료. {len(rows)}개저장.")
print(f"파일: {OUT_CSV}")
        write_log("done", f"saved={len(rows)}")
except Exception as e:
print("오류가났어요.")
print(str(e))
        write_log("error", str(e))
if __name__ == "__main__":
    main()
```

## **4) 테스트** 

예: `연기 커뮤니티` 로 해봐요. 개수는 50으로 해봐요. 끝나면 **Search.csv** 를 열어요. 

## **5) 문제해결** 

결과가 0개예요? → 단어를 더 넓게 바꿔봐요. → 예: 연기 → 연극, 배우. 너무 적어요? → 개수를 100으로 넣어요. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 10/169 

25. 11. 30. 오후 1:05 

 

**==> picture [147 x 9] intentionally omitted <==**

느려요? 

→ 다른 검색어로 먼저 점검해요. 

→ 잠깐 쉬고 다시 해요. 

## **6) 확장 예시** 

제목만 말고 설명도 모아요. 

중복 링크를 지워요. 

도메인별 TOP을 만들어요. 

→ **TOP_Search.csv** 로 저장해요. 

필요하면 확장 만들어 줄게요. 

“확장”이라고 말해요. 

**==> picture [489 x 38] intentionally omitted <==**

확장에서 1,설명을 모은다는건 뭐지? 2, 도메인별 top은 어떤기준의 top일까? 

3,링크를 모아오는 과정에서 DuckDuckGo만 사용되지? 다 른 검색 포털이 존재하는데? 

예) 네이버, 구글, 다음 등 

4, 여기에 필요한 설치 프로그램이 있지 않아? 그걸 왜 빼 먹고 혹은 방법을 이야기 하지 않고 넘어가지? 

16s 동안 생각함 

형, 질문 4개에 바로 답할게요. 

(모두 **Search.py / Search.csv / TOP_Search.csv / LOG_Search.csv** 유지) 

## **1) “설명을 모은다” 뜻** 

**뜻:** 링크마다 짧은 요약을 함께 저장. 

## **빠른 방법:** DuckDuckGo 결과에 있는 **요약칸(result snippet)** 을 같이 긁어와 

- `description` 칼럼으로 저장. 

**느린(정확) 방법:** 각 링크를 열어 **메타설명(** `<meta name="description">` **)** 이나 **본문** 

**첫 문단** 을 가져와 저장. 

장점: 더 정확함. 

단점: 느리고, 막힌 사이트도 있음. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 11/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## **2) “도메인별 TOP” 기준** 

**기본 기준(추천):** 많이 나온 순서. 

   - `domain` , `count` , `best_rank` 로 저장. 

- 정렬: `count` 내림차순 → `best_rank` 오름차순(동률일 때 상위 노출 승). 

- **가중치 기준(선택):** 상위 랭크 가중치. 

   - 예: 점수 = Σ (1 / rank). 

   - 정렬: `score` 내림차순. 

파일은 **TOP_Search.csv** 로 저장. 

## **3) 검색 포털 추가 가능?** 