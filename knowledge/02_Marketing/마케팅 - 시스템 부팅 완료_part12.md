
**2.** `RUN_DISCOVERY_ONCE.cmd` 더블클릭 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

57/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

## **3.** 뜨는 **파이썬 입력창** 에: 

검색어: 예) `연기 오디션 팁` 

- 라벨: 예) `연기 커뮤니티` 

→ 네이버 VIEW/WEB + 다음 BLOG/WEB에서 **제목/URL** 자동 수집 

## **결과** 

```
outputs\discovered_once.csv
```

컬럼: `engine, rank, title, url, host, seed_query, label, ts_kst` 

(중복 URL 자동 제거, UTF-8-SIG로 저장 → 엑셀 바로 열림) 

## **안 될 때 10초 점검** 

파이썬 설치 여부: `python --version` 

최초 실행 시 자동으로 `requests / beautifulsoup4 / pandas` 설치함 

콘솔 한글 깨지면 정상(파일은 UTF-8). 필요하면 .cmd 맨 위 `chcp 65001 > nul` 넣어 둠. 

이 버전은 **“주소 발굴 1회 실행”**만 딱 한다. 

원하는 대로 동작 확인되면, 이후에 **엔진 추가(유튜브/카페 등), 개수조절, 중복제거옵션** 순서로 넓혀 가자. 

**==> picture [489 x 38] intentionally omitted <==**

**==> picture [289 x 59] intentionally omitted <==**

**==> picture [159 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
아무것도 나오지 않는데?<br>**----- End of picture text -----**<br>


오케이 형. **CSV 빈 이유는 ‘검색결과 파싱 선택자’가 페이지 버전에 안 맞아서** 링크를 못 집어온 거야. 

바로 고치는 **패치(붙여넣기)** 줄게. 이대로 바꾸면 네이버/다음 최신 UI에서도 잘 잡힌다. 

## **1) 파일 교체** 

메모장으로 `scripts\discover_once.py` **전체를 아래 코드로 통째로 교체** → 저장. 

python 

코드 복사 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

58/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

```
import os, sys, time
from datetime import datetime, timezone, timedelta
from urllib.parse import urlparse
import requests
from bs4 import BeautifulSoup as BS
import pandas as pd
```

```
UA = {
"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (
"Accept-Language": "ko-KR,ko;q=0.9,en-US;q=0.8",
}
KST = timezone(timedelta(hours=9))
```

```
defnow_kst(): return datetime.now(KST).strftime("%Y-%m-%d %H:%M")
defhost_of(u:str)->str:
try: return urlparse(u).netloc.lower()
except: return""
```

```
defget(url, params=None, headers=None, timeout=12):
    h = dict(UA)
if headers: h.update(headers)
    r = requests.get(url, params=params, headers=h, timeout=timeout)
    r.raise_for_status()
return r
```

```
# -------- NAVER --------
```

```
defnaver_block(query, where="view", top=20):
"""NAVER 검색(뷰/웹)에서타이틀/링크추출 — 다양한선택자대응"""
    rows = []
try:
        html = get("https://search.naver.com/search.naver",
                   {"where": where, "sm": "tab_jum", "query": query}).text
        soup = BS(html, "html.parser")
# UI 버전에따라번갈아나오는선택자풀
        candidates = [
"a.total_tit",                 # 기존
"a.title_link",                # 새 UI
"a.api_txt_lines.total_tit",   # 일부구버전
        ]
        rank = 0
        seen = set()
for sel in candidates:
for a in soup.select(sel):
```

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

59/62 

25. 11. 30. 오후 1:33 

 

마케팅 - 시스템 부팅 완료 

```
                href = (a.get("href") or"").strip()
                title = a.get_text(strip=True)
ifnot href.startswith("http"): continue
if href in seen: continue
                seen.add(href)
                rank += 1
                rows.append({
"engine": f"naver:{where}",
"rank": rank, "title": title, "url": href
                })
if rank >= top: break
if rank >= top: break
except Exception as e:
pass
return rows
```

```
# -------- DAUM --------
defdaum_block(query, w="blog", top=20):
"""
"""DAUM 검색(블로그/웹)에서타이틀/링크추출 — 다양한선택자대응
    rows = []
try:
        html = get("https://search.daum.net/search",
                   {"w": w, "q": query},
                   headers={"Referer":"https://www.daum.net/"}).text
        soup = BS(html, "html.parser")
        candidates = [
"a.f_link_b", "a.tit_main", "a.tit_item",  # 빈도가높은타이틀앵커
"a.link_coll", "div.wrap_tit a",           # 일부카드형
        ]
        seen = set(); rank = 0
for sel in candidates:
for a in soup.select(sel):
                href = (a.get("href") or"").strip()
                title = a.get_text(strip=True)
ifnot href.startswith("http"): continue
if"search.daum.net"in href: continue
if href in seen: continue
                seen.add(href); rank += 1
                rows.append({
"engine": f"daum:{w}",
"rank": rank, "title": title, "url": href
                })
if rank >= top: break
```

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

60/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

```
if rank >= top: break
except Exception:
pass
return rows
```

```
defmain():
    base = sys.argv[1] iflen(sys.argv)>=2else os.getcwd()
    outdir = os.path.join(base, "outputs"); os.makedirs(outdir, exist_ok=True)
print("검색어를입력하세요 (예: 연기오디션팁): ", end="", flush=True)
    query = input().strip()
print("라벨을입력하세요 (예: 연기커뮤니티) [엔터가능]: ", end="", flush=True)
    label = input().strip()
```

```
ifnot query:
print("[ERROR] 검색어가비었습니다."); sys.exit(1)
```

```
# 라벨이비어도 '커뮤니티'를자동부가 — 커뮤니티위주수집
    add = label if label else"커뮤니티"
    q_full = f"{query}{add}".strip()
```

```
    rows = []
    rows += naver_block(q_full, "view", top=30)
    rows += naver_block(q_full, "web",  top=30)
    rows += daum_block(q_full,  "blog", top=30)
    rows += daum_block(q_full,  "web",  top=30)
```

```
# 메타/중복처리
```

```
for r in rows:
```

```
        r["host"] = host_of(r["url"])
        r["seed_query"] = query
        r["label"] = label
        r["ts_kst"] = now_kst()
```

```
    df = pd.DataFrame(rows, columns=["engine","rank","title","url","host","seed_q
           .drop_duplicates(subset=["url"]).reset_index(drop=True)
    out = os.path.join(outdir, "discovered_once.csv")
dft(tidFldi"tf8i")
```

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

61/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

62/62 
