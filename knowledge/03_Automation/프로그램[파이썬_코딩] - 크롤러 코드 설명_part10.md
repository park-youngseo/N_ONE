
우리가 원하는 건 사실: 

`actorsForum` 안에 있는 

- “글 상세 페이지” URL ( `/actorsForum/숫자` ) 만 가져오는 것. 

## 그래서 **필름메이커스 전용으로 필터를 좁혀야** 한다. 

## **3. 해결:** `fetcher.py` **를 “actorsForum 글만 긁는 버전”으로 교체** 

이제부터는: 

## `href` 가 **/actorsForum/로 시작하고, 뒤에 숫자(id)가 붙은 것만** 글로 인정 

- `#` , `?sort_index=…` , `/write` , `/signup` 같은 건 전부 무시 

- 그 링크들만 디테일 요청 

이렇게 만들면, 

## **필름메이커스 전체 사이트가 아니라 “배우 포럼 글들만”** 짧게 돌게 된다. 

## **3-1.** `fetcher.py` **전체를 아래 코드로 갈아끼워줘** 

Antigravity에서 `fetcher.py` 열고, 안에 있는 코드 전부 지우고 

아래 코드 그대로 붙여 넣고 **저장(Ctrl+S)** 해. 

python 

코드 복사 

```
import re
from urllib.parse import urljoin, urlparse
```

```
import requests
from bs4 import BeautifulSoup
import time
import random
```

```
from post import process_post
from logger import log
```

```
headers = {
"User-Agent": "Mozilla/5.0 (Windows NT 10.0)"
}
defis_actor_post_link(href: str) -> bool:
"""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

48/114 

25. 11. 30. 오후 1:06 

 

후 1:06 1:06 프로그램[파이썬/코딩] - 크롤러 코드 설명 `필름메이커스 배우 포럼 글 링크만 골라내는 필터 예: /actorsForum/26973933?page=1  같은 형태만 True """ if not href: return False` _`#`_ `페이지 내 이동` _`(#) /`_ `글작성` _`/`_ `가입` _`/`_ `로그인 등은 버린다 if href == "#" or "write" in href or "signup" in href or "dispMember" in href return False` _`#`_ `배우포럼이 아닌 다른 게시판은 버린다 if not href.startswith("/actorsForum/"): return False` _`#`_ `숫자` _`id`_ `가 들어 있는 글만 통과 return bool(re.search(r"/actorsForum/\d", href))` 

```
deffetch_generic_board(url: str, max_pages: int = 3, max_posts: int = 50) -> list
"""
필름메이커스배우포럼전용크롤러
    - 목록페이지: https://www.filmmakers.co.kr/actorsForum?page=1,2,3...
    - 각페이지에서 /actorsForum/숫자형태링크만골라서글상세를가져온다.
    """
    posts = []
    log(f"크롤링시작: {url}")
print(f"[INFO] 크롤링시작: {url}")
```

```
    base = f"{urlparse(url).scheme}://{urlparse(url).netloc}"
```

```
try:
        page = 1
while page <= max_pages:
            page_url = f"{url}?page={page}"
            log(f"[INFO] {page}페이지요청: {page_url}")
print(f"[INFO] {page}페이지요청: {page_url}")
```

```
            r = requests.get(page_url, headers=headers, timeout=10)
if r.status_code != 200:
```

```
                log(f"[WARN] 상태코드{r.status_code}로중단")
print(f"[WARN] 상태코드{r.status_code}로중단")
break
```

```
            soup = BeautifulSoup(r.text, "html.parser")
```

```
# 1) 페이지안에서배우포럼글링크만고른다
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

49/114 

25. 11. 30. 오후 1:06 

 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

```
            links = set()  # 중복방지
for a in soup.select("a[href]"):
                href = a.get("href", "")
if is_actor_post_link(href):
                    links.add(href)
ifnot links:
                log("[INFO] 이페이지에서글링크를찾지못해중단")
print("[INFO] 이페이지에서글링크를찾지못해중단")
break
```

```
# 2) 골라낸글링크만상세요청
for href in links:
                detail_url = urljoin(base, href)
try:
                    log(f"[DETAIL] 글요청: {detail_url}")
print(f"[DETAIL] 글요청: {detail_url}")
```

```
                    d = requests.get(detail_url, headers=headers, timeout=10)
                    ds = BeautifulSoup(d.text, "html.parser")
```

```
# 제목은보통 <title> 이나 h1/h2 에서잡고,
# 본문은전체텍스트로단순수집 (필요하면나중에더정교하게)
                    title = ds.title.text.strip() if ds.title else""
                    body = ds.get_text(" ", strip=True)
                    comments = ""
```

```
                    row = {
"title": title,
"body": body,
"comments": comments,
                    }
                    processed = process_post(row)
```

```
if processed:
                        posts.append(processed)
iflen(posts) >= max_posts:
                            log(f"[INFO] max_posts({max_posts}) 도달, 수집중단")
print(f"[INFO] max_posts({max_posts}) 도달, 수집중단"
                            log(f"수집완료: {len(posts)}개")
print(f"[INFO] 수집완료: {len(posts)}개")
return posts
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 50/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
except Exception as e:
                    log(f"[ERROR] 상세페이지오류: {e}")
print(f"[ERROR] 상세페이지오류: {e}")
continue
```

```
            page += 1
            time.sleep(random.uniform(1, 2))
```

```
except Exception as e:
        log(f"[ERROR] fetch_generic_board 전체오류: {e}")
print(f"[ERROR] fetch_generic_board 전체오류: {e}")
```

```
    log(f"수집완료: {len(posts)}개")
print(f"[INFO] 수집완료: {len(posts)}개")
return posts
```

## 이 버전은: 

**필름메이커스 배우 포럼 글만** 타겟 

- 최대 `max_pages` 페이지 × `max_posts` 글까지만 

- 콘솔에도 `[INFO]` , `[DETAIL]` 이 계속 찍혀서 

- “지금 뭐 하는지” 눈으로 보인다. 

## **4. 형이 바로 확인할 수 있는 것** 

**1.** `fetcher.py` 위 코드로 저장 후, 

**2.** Antigravity 터미널에서: 

bash 코드 복사 

```
python main.py
```

**3.** 사이트 이름: `필름메이커스` 

URL: `https://www.filmmakers.co.kr/actorsForum` 

→ 빈 줄 엔터. 

## 그러면 콘솔에: 

- `[INFO] 1페이지 요청: ...` 

- `[DETAIL] 글 요청: https://www.filmmakers.co.kr/actorsForum/숫자...` 

- `[INFO] max_posts(50) 도달, 수집 중단` 

- `[INFO] 저장 완료: output\필름메이커스.csv` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 51/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

이런 흐름이 훨씬 빠르고 깔끔하게 나올 거야. 

## **5. 정리** 

## **지금까지 “멈춘 것처럼” 보였던 이유** 

- → 배우포럼 글만이 아니라 **사이트 전체 링크를 죄다 타고 돌고 있어서** 

네트워크 요청이 엄청 많았기 때문. crawler 

**에러 메시지들(SSL,** `Invalid URL '#'` **)은 전부 try/except로 잡혀 있어서 프로그램이 죽은 건 아니다.** 

이미 한 번은 `output\필름메이커스.csv` 가 생성된 상태고, 

앞으로는 `fetcher.py` 를 위처럼 좁혀서 

- **“필름메이커스 배우 포럼 전용 크롤러”** 느낌으로 쓰면 된다. 

이대로 갈아끼우고 한 번만 더 돌려보면, 

이번엔 실행 시간 / 로그 / csv 결과가 다 눈에 들어올 거야. 

**==> picture [489 x 40] intentionally omitted <==**

**crawler.log** 파일 

이번엔 결과가 0이야 

형, 로그 보니까 왜 0개 나왔는지 딱 보인다. 

## **1. 지금 무슨 일이 있었는지** 

로그를 보면: 

배우포럼 글 링크들은 제대로 타고 들어가고 있어. 

( `.../actorsForum/26973933?page=1` 이런 주소들 전부 요청) 

그런데 마지막에 

text 코드 복사 

```
수집완료: 0개
```
