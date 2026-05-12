
다만: 

- 쓸데없는 링크까지 다 요청하다 보니 속도가 느려지고, 

- 로그가 지저분해져서 형이 보기에는 “계속 에러나는 느낌”이었던 거지. 

지금처럼 필터 한 줄만 추가해 주면, 

**1.** 쓸모없는 링크들 안 타고 

**2.** 에러 메시지도 사라지고 

**3.** 속도도 약간 좋아지고 

이 세 가지가 동시에 정리된다. 

## 정리하면: 

- 그 에러는 `href="#"` **같은 가짜 링크를 진짜 URL로 착각해서** 생긴 거고, 

- 프로그램이 죽는 에러는 아니라서 내가 앞에서 결과/필터 쪽 위주로만 설명하다가 따로 못 짚은 거야. 

- `fetcher.py` 에 `if href == "#" or href.startswith("javascript:"): continue` 

- 이 한 줄만 넣어주면 깔끔하게 사라진다. 

이대로 수정하고 한 번 더 돌려보고, 

그래도 로그에 다른 이상한 에러가 새로 뜨면 그 문장 그대로 가져오면 

거기서부터 또 같이 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

그게 중요한게 아니라 왜 검색을 못해오느냐가 문제잖아 

59s 동안 생각함 

오케이, 핵심은 `# 에러` 가 아니라 

**“왜 배우 포럼 글을 제대로 못 긁어오냐”** 이거잖아. 그 얘기만 딱 짚을게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

62/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

## **1. 지금 로그 기준으로 보면, “멈춘 게 아니라 엉뚱한 걸 긁고 있다”** 

형이 준 `crawler.log` 보면 흐름이 이래: 

`actorsForum?page=1` 까진 잘 들어감 

- 그 다음에: 

   - `/community` , `/proCrewRecruiting` , `/notice` , `/event` , `/about` , `/adinfo` , 

   - `/privacypolicy` … 

   - 심지어 인스타/유튜브까지 전부 `글 요청:` 으로 돌고 있음 

즉, **배우 게시판 글만 긁는 게 아니라 사이트 메뉴/푸터/외부 링크까지 다 글로 착각** 해서 도는 상태야. 

거기서 두 가지가 동시에 터진다: 

## **1. 엉뚱한 링크까지 전부 도느라 시간만 쓰고** , 

**2.** 그 내용을 `process_post()` 에서 다시 필터링하다가 

   - 광고/스탭으로 오판되면 → 다 버림 

그래서 최종 결과가 **0개처럼 보이거나, 원하는 글이 하나도 없음** 

그래서 형이 느끼는 “검색을 못해온다”는 감각이 맞는 거야. 

**필름메이커스 배우 포럼 글을 타겟팅 못 해서 그렇다.** 

## **2. 왜 “배우 포럼 글”을 못 잡는지 구조를 보면** 

지금 fetcher는 대략 이런 구조야(형 코드 기준): 

python 코드 복사 

```
soup = BeautifulSoup(r.text, "html.parser")
items = soup.select("a")  # 페이지안의모든링크
```

```
for a in items:
    title = a.text.strip()
    href = a.get("href", "")
    ...
    d = requests.get(href, ...)
```

문제: 

`a` 태그면 **다 글이라고 착각** 

→ 메뉴, 탭, 푸터, 외부 사이트 전부 다 탐 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

63/114 

25. 11. 30. 오후 1:06 

 

- 프로그램[파이썬/코딩] - 크롤러 코드 설명 

배우 포럼 글만 골라내는 필터가 없음 

거기에 `post.py` 에서: 

python 코드 복사 

```
if is_staff_post(merged): returnNone
```

```
if is_ad_post(merged): returnNone
```

이런 필터로 또 한 번 싹 쓸어버림. 

결과적으로: 

## **배우 글만 콕 집어서 가져오는 게 아니라 사이트 전체를 넓게 긁고, 그중 대부분은 필터에 걸려서 버리는 구조** 

라서 형 입장에선 “검색 자체를 못한다”로 느껴지는 게 맞음. 

## **3. 해결 방향 – “배우 포럼 전용 + 필터 잠깐 풀기”** 

지금 당장은 **두 가지를 동시에 해야** 제대로 보일 거야. 

## **1.** `fetcher.py` 

→ **배우 포럼 글 링크만** 대상으로 좁히기 

- (예: `/actorsForum/숫자` 형태만 “글”로 인정) 

## **2.** `post.py` 

→ 필터를 잠깐 풀어서 

“일단 뭐라도 CSV에 쌓이는지” 확인하기 

일단 **“검색은 된다”**를 확인하는 단계니까, 광고/스탭 필터는 나중에 다시 조이면 돼. 

## **4. 지금 당장 바꿔야 하는 실제 코드** 

## **4-1.** `fetcher.py` **→ 배우 포럼 글만 긁는 버전** 

`fetcher.py` 전체를 **아래 코드로 갈아끼워 줘** (안티그래비티에서 열고 전부 교체 → 저 

장): 

python 코드 복사 `import re from urllib.parse import urljoin, urlparse` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

64/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
import requests
from bs4 import BeautifulSoup
import time
import random
from post import process_post
from logger import log
headers = {
"User-Agent": "Mozilla/5.0 (Windows NT 10.0)"
}
defis_actor_post_link(href: str) -> bool:
"""
필름메이커스배우포럼글링크만골라내는필터
예: /actorsForum/26973933?page=1 같은형태만 True
    """
ifnot href:
returnFalse
    href = href.strip()
# 페이지내이동 / 자바스크립트버튼제외
if href == "#"or href.startswith("javascript:"):
returnFalse
# 배우포럼이아니면제외
ifnot href.startswith("/actorsForum/"):
returnFalse
```

```
# 숫자 id 가들어있는글만허용
returnbool(re.search(r"/actorsForum/\d", href))
```

```
deffetch_generic_board(url: str, max_pages: int = 1, max_posts: int = 20) -> list
"""
필름메이커스배우포럼전용크롤러
    - 목록: https://www.filmmakers.co.kr/actorsForum?page=1,2...
    - 각페이지에서 /actorsForum/숫자형태링크만골라서상세를요청
    """
    posts = []
    log(f"크롤링시작: {url}")
print(f"[INFO] 크롤링시작: {url}")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 65/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    parsed = urlparse(url)
    base = f"{parsed.scheme}://{parsed.netloc}"
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
                log(f"[WARN] 상태코드{r.status_code}로중단")
print(f"[WARN] 상태코드{r.status_code}로중단")
break
```

```
            soup = BeautifulSoup(r.text, "html.parser")
```

```
# 1) 이페이지에서배우포럼글링크만수집
            links = set()
for a in soup.select("a[href]"):
                href = a.get("href", "")
if is_actor_post_link(href):
                    links.add(href)
ifnot links:
                log("[INFO] 이페이지에서배우포럼글링크를찾지못해중단")
print("[INFO] 이페이지에서배우포럼글링크를찾지못해중단")
break
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
                    title = ds.title.text.strip() if ds.title else""
                    body = ds.get_text(" ", strip=True)
                    comments = ""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 66/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
                    row = {
"title": title,
"body": body,
"comments": comments,
                    }
                    processed = process_post(row)
if processed:
                        posts.append(processed)
iflen(posts) >= max_posts: