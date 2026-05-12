
이제부터 **사이트/게시판 선택 = URL 입력** 으로 통일. 

## 🔧 **그래비티 패치용 문구 (이거 그대로 기존 문서 맨 아래에 붙여줘)** 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

568/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

형은 아래 블록을 **그대로 그래비티 지시문 끝에 복붙** 하면 돼. 

## **[스펙 수정 v2 — URL 기반 범용 크롤러로 변경]** 

이 수정은 이전 지시문에 있던 

“사이트 = 필름메이커스 고정 / 게시판 3개 고정 / 사이트 선택 메뉴” 내용을 **모두 대 체한다.** 

## **1. 사이트·게시판 하드코딩 금지** 

코드 어디에도 **“필름메이커스”** , **“엑터스 포럼/연기자 정보방/자유게시판”** 같은 **특정 사이트/게시판 이름을 하드코딩해 넣지 말 것.** 

크롤러는 **“URL 기준”으로만 동작** 해야 한다. 

## **2. exe 실행 시 첫 화면 = “URL 입력 화면”** 

`main.exe` 실행하면 가장 먼저 아래와 같은 화면(콘솔 기준)을 띄운다. 

text 코드 복사 

```
=============================
```

```
엔오원니즈수집크롤러 v2
```

```
=============================
```

```
[1] 크롤링할게시판 URL 입력
```

```
한줄에하나씩붙여넣어주세요.
```

```
예)
```

```
https://www.filmmakers.co.kr/filmmakerforum
https://www.filmmakers.co.kr/actorsinfo
https://어떤사이트.com/actor-board
```

```
입력이끝났으면빈줄에서 Enter를누르세요.
```

사용자는 **크롤링하고 싶은 게시판의 “목록 페이지 URL”**을 

- 한 줄에 하나씩 붙여넣는다. 

- 이 URL 목록이 곧 **사이트/게시판 정보 전체** 가 된다. 

- (사이트/게시판은 코드에 고정하지 않고, 전부 이 URL에서 파생) 

## **3. 이후 옵션 흐름** 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

569/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

URL 입력이 끝나면 그 다음 순서: 

## **1. 수집 기간 선택** 

   **1.** 최근 3년 

   **2.** 최근 1년 

   **3.** 기간 직접 입력 (YYYY-MM-DD ~ YYYY-MM-DD) 

**2. 저장 형식 선택** 

   **1.** CSV 

   **2.** JSON 

## **3. 실행 여부 확인** 

“입력한 URL N개 / 기간 / 저장 형식”을 다시 보여주고 

```
[Enter]를누르면크롤링을시작합니다.
```

→ 그 다음에 `run_crawler(urls, years, save_format)` 실행. 

## **4.** `fetcher.py` **인터페이스도 URL 기반으로 변경** 

기존(사이트/게시판 코드 기반) 설계는 버리고, 

다음과 같이 **URL 리스트를 직접 받는** 함수로 바꾼다. 

python 

코드 복사 

```
# fetcher.py
from typing importList
from post import Post
from logger import log
```

```
deffetch_posts(urls: List[str], years: int) -> List[Post]:
"""
```

```
    urls: 사용자가입력한게시판 '목록페이지' URL 리스트.
예) ["https://www.xxx.com/boardA", "https://www.yyy.com/free"]
```

```
    years: 최근몇년치까지수집할지.
    """
    all_posts: List[Post] = []
for url in urls:
        board_posts = fetch_generic_board(url=url, years=years)
        all_posts.extend(board_posts)
return all_posts
```

```
deffetch_generic_board(url: str, years: int) -> List[Post]:
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

570/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
"""
```

```
단일게시판 URL을크롤링하는범용함수.
```

```
    - url 기준으로페이지를넘기고
    - 각각글의제목/본문/댓글/작성자/작성일자를/본문/댓글/작성자/작성일자를본문/댓글/작성자/작성일자를/댓글/작성자/작성일자를댓글/작성자/작성일자를/작성자/작성일자를작성자/작성일자를/작성일자를작성일자를파싱해서 Post 리스트로 Post 리스트로리스트로
    - 사이트/게시판이름은 url에서유추하거나, 필요한최소수준만기록.
    """
    ...
```

```
    - 각각글의제목/본문/댓글/작성자/작성일자를/본문/댓글/작성자/작성일자를본문/댓글/작성자/작성일자를/댓글/작성자/작성일자를댓글/작성자/작성일자를/작성자/작성일자를작성자/작성일자를/작성일자를작성일자를파싱해서 Post 리스트로 Post 리스트로리스트로반환.
```

## 요구사항: 

`fetch_generic_board()` 는 

## **주어진 URL이 어떤 사이트든 간에** 

“게시글 목록 + 페이지 넘김 + 글 상세 + 댓글” 구조를 

일반적인 패턴(쿼리스트링 page, p, pageNo 등) 기준으로 처리한다. 

특정 사이트 전용 상수(도메인, path, 게시판명)을 넣지 말고, 

HTML 구조/패턴을 일반 로직으로 처리한다. 

(필요하면 내부적으로 “도메인별 분기” 정도는 가능하지만, 

지시문 상에서 특정 사이트 이름을 못 박지 않는다.) 

## **5.** `main.py` **도 URL 기준으로 변경** 

python 

코드 복사 

```
# main.py (개념)
```

```
from fetcher import fetch_posts
```

```
from cleaner import clean_post
from saver import save_csv, save_json
from logger import log
```

```
defask_user_options():
```

- _`# 1)`_ `크롤링할` _`URL`_ `들 입력` _`(`_ `여러 줄` _`)`_ 

- _`# 2)`_ `기간 선택` _`(3`_ `년` _`/1`_ `년` _`/`_ `직접 입력` _`) # 3)`_ `저장 형식 선택` _`(csv/json)`_ 

```
# → (urls, years, save_format) 반환
```

```
    ...
```

```
defrun_crawler(urls, years, save_format):
```

```
    log(f"엔오원크롤러시작 / URL {len(urls)}개 / 기간: 최근{years}년")
    raw_posts = fetch_posts(urls=urls, years=years)
```

```
    cleaned_rows = []
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

571/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
for p in raw_posts:
        row = clean_post(p)
if row isNone:
continue
        cleaned_rows.append(row)
```

```
# 저장형식분기
if save_format == "json":
        ...
else:
        ...
    log("엔오원크롤러종료")
```

```
if __name__ == "__main__":
    urls, years, save_format = ask_user_options()
    run_crawler(urls, years, save_format)
```

## **6. 게시판 이름(board_name) 처리** 

게시판 이름/사이트 이름도 **코드에 고정하지 않고** , 

각 Post의 `board_name` 은 다음과 같이 처리한다. 

**1.** 기본값: 

URL에서 도메인 + path 일부를 잘라서 자동 생성 

- 예: 

`https://www.filmmakers.co.kr/filmmakerforum` → `"filmmakers / filmmakerforum"` 
