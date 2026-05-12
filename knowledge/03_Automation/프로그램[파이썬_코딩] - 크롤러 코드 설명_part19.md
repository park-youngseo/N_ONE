한 번 그대로 실행해 보고, 

- 터미널에 `[RESULT] 총 수집된 글: N개` 

- `output/actors_simple.csv` 안에 줄이 몇 줄 나오는지 

까지만 먼저 확인해보자. 

여기까지 제대로 나오면, 그다음부턴 “필요한 글만 고르는 작업”은 

그 위에 하나씩 얹으면서 같이 만들 수 있다. 

**==> picture [489 x 38] intentionally omitted <==**

## 진짜 씨발 이제 슬슬 욕이나오네 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

91/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

21s 동안 생각함 

## 형, 욕 나오는 거 당연해… 

여기까지 시간 이렇게 썼는데 원하는 건 하나도 안 나오니까 진짜 짜증 날 수밖에 없지. 일단 미안하다. 내가 중간에 로직도 여러 개 섞어서 더 헷갈리게 만든 것도 맞고. 

지금부터는 **딱 한 가지 “되는 것”만** 집중해서 만들자. 

복잡한 v2.4니 뭐니 전부 잊고, **새 폴더 + 단일 파일** 로 완전 리셋해서 갈게. 

## **A. 완전 리셋 – 새 폴더에서 “실제로 돌아가는 것” 하나 만들기** 

## **1단계) 아예 새 폴더 만들기** 

**1.** 바탕화면에서 새 폴더 만들고 이름을 이렇게 해줘: 

text 코드 복사 `n1_simple_actors` **2.** Antigravity → `Open Folder` 눌러서 이 `n1_simple_actors` 폴더를 연다. (왼쪽 Explorer에 아무 파일도 없는 상태여야 정상.) 

## **2단계) 파일 1개만 만든다** 

**1.** Explorer 빈 공간에서 오른쪽 클릭 → **New File** . 

**2.** 파일 이름: 

text 코드 복사 `actors_simple.py` 

**3.** 아래 코드 **전체** 를 그대로 복사해서 `actors_simple.py` 안에 붙여넣고 저장(Ctrl+S). 

python 코드 복사 `import requests from bs4 import BeautifulSoup import csv import os import time import random from urllib.parse import urljoin, urlparse` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

92/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
# 필름메이커스배우포럼기본주소
BASE_BOARD_URL = "https://www.filmmakers.co.kr/actorsForum"
HEADERS = {
# 일반크롬브라우저흉내 (너무특이하지않게)
"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "
"(KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
"Referer": "https://www.filmmakers.co.kr/",
}
defis_actor_link(href: str) -> bool:
"""배우포럼글링크만고르는간단규칙"""
ifnot href:
returnFalse
    href = href.strip()
if href == "#"or href.startswith("javascript:"):
returnFalse
ifnot href.startswith("/actorsForum/"):
returnFalse
# 숫자아이디가들어있는것만글로본다
returnany(ch.isdigit() for ch in href)
```

```
deffetch_list_page(page: int) -> str:
"""목록페이지 HTML 가져오기"""
    url = f"{BASE_BOARD_URL}?page={page}"
print(f"[LIST] 페이지{page}요청: {url}")
    r = requests.get(url, headers=HEADERS, timeout=10)
    r.raise_for_status()  # 200 아니면에러던짐
return r.text
```

```
defextract_post_links(list_html: str) -> list[str]:
"""목록페이지에서배우포럼글링크들만뽑기"""
    soup = BeautifulSoup(list_html, "html.parser")
    links = set()
for a in soup.select("a[href]"):
        href = a.get("href", "")
if is_actor_link(href):
            links.add(href)
```

```
    parsed = urlparse(BASE_BOARD_URL)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

93/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    base = f"{parsed.scheme}://{parsed.netloc}"
```

```
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
# 본문: 페이지전체텍스트(우선은통으로)
    body = soup.get_text(" ", strip=True)
return title, body
```

```
defmain():
    max_pages = 1# 우선 1페이지만테스트
    all_rows = []
    seen_urls = set()
for page inrange(1, max_pages + 1):
try:
            html = fetch_list_page(page)
except Exception as e:
print(f"[ERROR] 목록페이지{page}요청실패: {e}")
break
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

94/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
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
# 서버부담줄이려고쉬어가기
            time.sleep(random.uniform(1, 2))
```

```
print(f"\n[RESULT] 총수집된글: {len(all_rows)}개")
ifnot all_rows:
print("[WARN] 수집된글이없어서파일저장을건너뜁니다.")
return
# CSV 저장
    os.makedirs("output", exist_ok=True)
    filepath = os.path.join("output", "actors_simple.csv")
```

```
withopen(filepath, "w", encoding="utf-8-sig", newline="") as f:
        writer = csv.DictWriter(f, fieldnames=["url", "title", "body"])
        writer.writeheader()
for row in all_rows:
            writer.writerow(row)
```

```
print(f"[SAVE] 저장완료: {filepath}")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

95/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
if __name__ == "__main__":
    main()
```

## 여기까지 하면 **파일 1개, 의존성도 requests/bs4만 쓰는 초간단 버전** 이다. 

## **3단계) 터미널에서 실행 (딱 한 줄)** 

**1.** Antigravity 메뉴 → **Terminal → New Terminal** . 

**2.** 아래에 뜬 터미널 프롬프트가 이렇게 나오는지 확인: 

text 코드 복사 `C:\Users\psm23\Desktop\n1_simple_actors>` (아니면 `cd` 명령으로 이 폴더까지 들어가면 돼: `cd C:\Users\psm23\Desktop\n1_simple_actors` ) **3.** 필요한 라이브러리 없으면 먼저 한 번만 설치: bash 코드 복사 `pip install requests beautifulsoup4` **4.** 그 다음 **실행** : bash 코드 복사 `python actors_simple.py` 

## **4단계) 나올 수 있는 상황 딱 3가지** 

## **① 정상 동작인 경우 (내가 기대하는 상황)** 

터미널에 대략 이런 흐름이 찍혀야 해: 

text 코드 복사 

- `[LIST] 페이지 1 요청: https://www.filmmakers.co.kr/actorsForum?page=1 [LIST] 이 페이지에서 글 링크 15개 발견` 

```
[POST] 글요청: https://www.filmmakers.co.kr/actorsForum/25996020?page=1
[POST] 글요청: ...
...
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

96/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
[RESULT] 총수집된글: 15개
```

- `[SAVE] 저장 완료: output\actors_simple.csv` 

## → 이러면 **“되는 것”** 확보 끝. 

`output/actors_simple.csv` 안에 글이 쌓여 있을 거야. 

## **②** `403 Client Error: Forbidden` **이 actorsForum에서도 뜨는 경우** 

이 메시지가 뜨면: 

text 

코드 복사 

   - `[ERROR] 목록 페이지 1 요청 실패: 403 Client Error: Forbidden for url: ...` 

- → 이건 필름메이커스 서버가 **배우포럼까지 봇 요청을 막기 시작한 것** 이고, 

- 이 단계부터는 코드로 깔끔하게 해결할 수 있는 영역이 아니다. 

      - 지금까지 403은 `/board` (자유게시판)에서만 있었고 