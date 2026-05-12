https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

55/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

41s 동안 생각함 

 

그래, 이해했다. 

이제 **찾아서 고치는 거 다 버리고** , `filmmakers_crawler.py` 전체를 한 방에 갈아끼우는 코 드 줄게. 

형은 그냥 **파일 내용 전부 지우고 → 아래 코드 그대로 복붙** 만 하면 된다. 

## **1. 갈아끼우는 방법** 

**1.** 안티그래비티에서 `filmmakers_crawler.py` 파일 연다. 

**2.** `Ctrl + A` 눌러서 **내용 전부 선택** . 

**3.** `Delete` 눌러서 싹 지운다. 

**4.** 아래 코드 블록 **전체** 를 복사해서 붙여 넣는다. 

**==> picture [481 x 409] intentionally omitted <==**

**----- Start of picture text -----**<br>
5.  Ctrl + S   로 저장.<br>6.  터미널에서:<br>bash 코드 복사<br>python filmmakers_crawler.py<br>실행.<br>2. 통째로 덮어쓸  filmmakers_crawler.py   코드<br>python 코드 복사<br>import requests<br>from bs4 import BeautifulSoup<br>import time<br>import random<br>from datetime import datetime, timedelta<br>from typing import List, Dict, Optional<br># ============================================<br># [Config]  크롤러 기본 설정<br>**----- End of picture text -----**<br>


```
# ============================================
```

```
# 필름메이커스기본주소
```

```
BASE_URL = "https://www.filmmakers.co.kr"
```

https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

56/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

 

```
# 배우오디션게시판리스트주소
```

```
LIST_URL = f"{BASE_URL}/actorsAudition"
```

```
# HTTP 요청헤더 (그냥일반브라우저처럼보이게)
HEADERS = {
"User-Agent": (
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) "
"Chrome/120.0.0.0 Safari/537.36"
    )
}
```

```
# 테스트용: 최대몇페이지까지돌릴지
MAX_PAGES = 2# 나중에 5, 10 이런식으로늘리면됨
```

```
# 수집기간기준 (최근 1년)
```

```
ONE_YEAR_AGO = datetime.now() - timedelta(days=365)
```

```
# ============================================
```

```
# [Utils] 유틸함수
```

```
# ============================================
```

```
defrandom_sleep(min_sec: float = 1.0, max_sec: float = 3.0) -> None:
```

```
"""차단덜당하려고요청사이에랜덤딜레이줌."""
```

```
    delay = random.uniform(min_sec, max_sec)
```

```
print(f"[SLEEP] {delay:.2f}초대기")
```

```
    time.sleep(delay)
```

```
# ============================================
```

```
# [Network] 페이지가져오기
```

```
# ============================================
```

```
deffetch_list_page(page: int) -> Optional[str]:
```

```
"""리스트페이지 HTML을가져온다. 실패하면 None."""
```

```
try:
```

```
        params = {"page": page}
```

```
print(f"[REQ] 리스트{page}페이지요청: {LIST_URL} params={params}")
```

```
        resp = requests.get(LIST_URL, headers=HEADERS, params=params, timeout=10)
```

```
if resp.status_code == 429:
```

```
print("[WARN] 429 Too Many Requests - 잠깐쉬었다가다시시도필요")
```

https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

57/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

 

```
returnNone
```

```
if resp.status_code == 403:
print("[WARN] 403 Forbidden - 차단가능성, 헤더/딜레이조정필요")
returnNone
        resp.raise_for_status()
return resp.text
except Exception as e:
print(f"[ERROR] 리스트페이지요청실패 (page={page}): {e}")
returnNone
finally:
        random_sleep()
```

```
# ============================================
```

```
# [Parsing] 리스트페이지파싱
```

```
# ============================================
```

```
defparse_list(html: str) -> List[Dict]:
"""
```

```
리스트페이지 HTML에서공고목록을뽑는다.
    *주의*: 선택자(selector)는대략맞춰둔거라, 실제구조와다를수있음.
잘안나오면나중에 selecotor만수정하면된다.
    """
    soup = BeautifulSoup(html, "html.parser")
    items: List[Dict] = []
```

```
# 테이블구조예측 (필름메이커스게시판기준대략)
# 실제로는개발자도구열어서태그/클래스를보고맞추면됨.
    rows = soup.select("table.board_list tbody tr")
```

```
for row in rows:
        title_el = row.select_one("td.title a")
ifnot title_el:
continue
```

```
        title = title_el.get_text(strip=True)
```

```
# 상대경로일가능성있으니 BASE_URL 붙여줌
        href = title_el.get("href", "")
if href.startswith("http"):
            url = href
else:
            url = BASE_URL + href
```

https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

58/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

 

```
        date_el = row.select_one("td.date")
        posted_at = date_el.get_text(strip=True) if date_el else""
```

```
        items.append(
            {
"title": title,
"url": url,
"posted_at": posted_at,
            }
        )
```

```
return items
```

```
# ============================================
# [Main Logic] 크롤러본체
# ============================================
```

```
defrun_crawler(max_pages: int = MAX_PAGES) -> List[Dict]:
"""
필름메이커스배우오디션게시판을
앞쪽 max_pages 페이지까지긁어서 list[dict]로반환.
    """
    results: List[Dict] = []
```

```
for page inrange(1, max_pages + 1):
        html = fetch_list_page(page)
ifnot html:
print(f"[WARN] 리스트{page}페이지 HTML 없음, 건너뜀")
continue
```

```
        items = parse_list(html)
print(f"[INFO] {page}페이지: {len(items)}개항목수집")
```
