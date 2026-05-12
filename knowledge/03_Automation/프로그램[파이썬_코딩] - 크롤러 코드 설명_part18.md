
- “ `actorsForum` 1페이지 글들을 가져와서” 

- “제목 / 링크 / 본문 전체 텍스트를 `actors_simple.csv` 로 저장” 

- “불필요한 필터 없이 “일단 다 가져오는” 버전” 

아래 그대로 하면 된다. 

## **1단계. 새 파일 하나만 만든다 (기존 n1_crawler는 일단 놔두자)** 

**1.** Antigravity 왼쪽 `Explorer` 에서 

   - `n1_crawler` 폴더가 열려 있는 상태에서, 

**2. 오른쪽 클릭 → New File** 선택. 

**3.** 파일 이름을 이렇게 써줘: 

text 

코드 복사 

```
fm_actors_simple.py
```

**4.** 그 파일이 열리면, 안에 아무 것도 없는 상태일 거야. 

## **2단계. 이 코드 그대로 복사해서 넣기** 

`fm_actors_simple.py` 안에 아래 코드 전체를 붙여넣고 저장(Ctrl+S): 

python 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

86/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
import requests
from bs4 import BeautifulSoup
import csv
import os
import time
import random
from urllib.parse import urljoin, urlparse
```

```
# -------- 기본설정 --------
BASE_BOARD_URL = "https://www.filmmakers.co.kr/actorsForum"
HEADERS = {
"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64)"
}
defis_actor_link(href: str) -> bool:
"""
배우포럼글링크만거르는간단한규칙
예: /actorsForum/25996020?page=1
    """
ifnot href:
returnFalse
    href = href.strip()
if href == "#"or href.startswith("javascript:"):
returnFalse
ifnot href.startswith("/actorsForum/"):
returnFalse
# 숫자아이디가들어간것만
returnany(ch.isdigit() for ch in href)
```

```
deffetch_list_page(page: int) -> str:
"""목록페이지 HTML 가져오기"""
    url = f"{BASE_BOARD_URL}?page={page}"
print(f"[LIST] 페이지{page}요청: {url}")
    r = requests.get(url, headers=HEADERS, timeout=10)
    r.raise_for_status()
return r.text
defextract_post_links(list_html: str) -> list:
"""목록페이지에서배우포럼글링크들만뽑기"""
    soup = BeautifulSoup(list_html, "html.parser")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

87/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    links = set()
for a in soup.select("a[href]"):
        href = a.get("href", "")
if is_actor_link(href):
            links.add(href)
    parsed = urlparse(BASE_BOARD_URL)
    base = f"{parsed.scheme}://{parsed.netloc}"
    result = []
for href in links:
        full_url = urljoin(base, href)
        result.append(full_url)
```

```
print(f"[LIST] 이페이지에서글링크{len(result)}개발견")
returnsorted(result)
```

```
deffetch_post(url: str) -> tuple[str, str]:
"""글상세페이지에서제목과전체텍스트가져오기"""
print(f"[POST] 글요청: {url}")
    r = requests.get(url, headers=HEADERS, timeout=10)
    r.raise_for_status()
    soup = BeautifulSoup(r.text, "html.parser")
# 제목
if soup.title:
        title = soup.title.get_text(strip=True)
else:
        title = url
# 본문: 우선은페이지전체텍스트
    body = soup.get_text(" ", strip=True)
return title, body
defmain():
# ---- 몇페이지까지볼지설정 ----
    max_pages = 1# 필요하면 2,3 으로늘리면됨
    all_rows = []
    seen_urls = set()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

88/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
for page inrange(1, max_pages + 1):
try:
            html = fetch_list_page(page)
except Exception as e:
print(f"[ERROR] 목록페이지{page}요청실패: {e}")
break
        post_urls = extract_post_links(html)
ifnot post_urls:
print("[INFO] 더이상글링크가없어중단")
break
for post_url in post_urls:
if post_url in seen_urls:
continue
            seen_urls.add(post_url)
try:
                title, body = fetch_post(post_url)
                all_rows.append({
"url": post_url,
"title": title,
"body": body,
                })
except Exception as e:
print(f"[ERROR] 글요청실패: {e}")
continue
# 서버부담줄이려고살짝대기
            time.sleep(random.uniform(1, 2))
print(f"\n[RESULT] 총수집된글: {len(all_rows)}개")
ifnot all_rows:
print("[WARN] 수집된글이없어서파일저장을건너뜁니다.")
return
# ---- CSV 저장 ----
    os.makedirs("output", exist_ok=True)
    filepath = os.path.join("output", "actors_simple.csv")
withopen(filepath, "w", encoding="utf-8-sig", newline="") as f:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

89/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
        writer = csv.DictWriter(f, fieldnames=["url", "title", "body"])
        writer.writeheader()
for row in all_rows:
            writer.writerow(row)
```

```
print(f"[SAVE] 저장완료: {filepath}")
```

```
if __name__ == "__main__":
    main()
```

## 이 파일 하나만 쓰는 거라서, 

`cleaner.py` / `post.py` / `patterns.py` 이런 거 전혀 신경 안 써도 된다. 

## **3단계. 실행 방법 (진짜 단순)** 

**1.** Antigravity 아래쪽에서 **Terminal** 열려 있는지 확인. 

경로가 `C:\Users\psm23\Desktop\n1_crawler>` 이런 식이면 OK. 

**2.** 거기서 이렇게 입력: 

bash 코드 복사 

```
python fm_actors_simple.py
```

**3.** 그러면 터미널에 예를 들어 이런 흐름이 뜰 거야: 

text 코드 복사 

```
[LIST] 페이지 1 요청: https://www.filmmakers.co.kr/actorsForum?page=1
[LIST] 이페이지에서글링크 15개발견
[POST] 글요청: https://www.filmmakers.co.kr/actorsForum/25996020?page=1
[POST] 글요청: ...
...
[RESULT] 총수집된글: 15개
[SAVE] 저장완료: output\actors_simple.csv
```

에러 없이 여기까지 나오면 **이건 “되는 것”**이야. 

## **4단계. 결과 확인** 

## **1.** 탐색기로 

`C:\Users\psm23\Desktop\n1_crawler\output` 폴더 열기. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

90/114 

25. 11. 30. 오후 1:06 프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

**2.** 그 안에 `actors_simple.csv` 파일이 생겼을 거야. 

**3.** 메모장이나 엑셀로 열어 보면: 

`url,title,body` 헤더가 있고 

그 아래 줄들이 여러 줄 있을 거야. 

`body` 는 페이지 전체 텍스트라 좀 지저분해도 “일단 다 들어와 있음”. 

이게 **기본 재료(원본 덩어리)**고, 

여기서부터 “필요한 글만 고르기(질문/고민)”는 나중에 필터를 추가해서 쪼개는 단계야. 

## **5단계. 현실적으로 가능한 선 정리** 

- `/actorsForum` 처럼 **열려 있는 게시판** → 지금 만든 것처럼 수집 가능 

- `/board` 자유게시판처럼 **403 막힌 곳** → 서버에서 막아버리기 때문에 

- 코드를 아무리 바꿔도 “완전 자동”은 안 됨 (이건 구조적으로 막힌 구간) 

## 그래서 

- “되는 걸 하나 만들어라”는 요청에 맞춰 

- **actorsForum 전용, 단일 파일, 테스트용이라도 확실히 돌아가는 크롤러** 를 준 거고, 이게 제대로 돌아가는 걸 확인한 다음에 

- “이 중에서 지망생 고민만 따로 빼는 필터”를 얹는 순서로 가는 게 형이 원하는 니즈 수집에 제일 가까울 거야. 

## 형, 

지금 이 `fm_actors_simple.py` 는 

진짜 최소 기능만 넣은 “돌아가는 버전”이니까 
