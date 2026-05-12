
https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

563/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
"label": "필름메이커스",
    },
"2": {
"name": "site2",
"label": "추가예정사이트 1",
    },
}
```

## **2. 게시판 선택도 사이트별로 선택하도록 변경** 

기존: 

“엑터스 포럼 / 연기자 정보방 / 자유게시판” → 고정 3개 

## 수정 후: 

**선택한 사이트에 따라 게시판 목록이 달라지는 구조** 로 바꾼다. 

지금은 필름메이커스만 있으니까, 필름메이커스 선택 시에만 아래 메뉴를 띄운다. 

예시: 

text 코드 복사 

```
[1] 사이트선택
```

```
    1) 필름메이커스
```

```
    2) (추가예정사이트 1)
    3) (추가예정사이트 2)
```

```
→ 1 선택시:
```

```
[2] 게시판선택 (필름메이커스)
```

```
    1) 엑터스포럼
```

```
    2) 연기자정보방
```

```
    3) 자유게시판
```

```
    4) 전체 (1~3 전부)
```

```
    5) 번호를콤마로입력 (예: 1,3)
```

## 요구사항: 

## **필름메이커스 전용 게시판 목록** 은 위 3개로 고정. 

- “전체” 선택 시 → 1~3 모두 

- “1,3” 입력 같은 조합 선택도 가능. 

- 나중에 다른 사이트를 추가할 때는, 그 사이트만의 게시판 목록을 정의해서 같은 구 조로 쓰면 된다. 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

564/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

의사코드: 

python 

코드 복사 

```
SITE_BOARDS = {
"filmmakers": {
"1": {"code": "actors_forum", "label": "엑터스포럼"},
"2": {"code": "actor_info", "label": "연기자정보방"},
"3": {"code": "free", "label": "자유게시판"},
    },
# "site2": { ... }  # 나중확장용
}
```

## **3.** `fetcher.py` **구조도 “사이트 → 게시판 리스트”로 바꾸기** 

기존: 

`fetch_posts(years=3)` → 내부에서 필름메이커스 + 3개 게시판을 고정으로 돌게 설 계 

## 수정 후: 

`fetch_posts(site: str, boards: list[str], years: int)` 형식으로 바꾼다. 

예시(그래비티가 구현해야 할 형태): 

python 코드 복사 _`# fetcher.py`_ `from typing import List from post import Post from logger import log` 

```
deffetch_posts(site: str, boards: list[str], years: int = 3) -> List[Post]:
"""
    site: 'filmmakers' 등사이트코드
    boards: 게시판코드리스트 (예: ['actors_forum', 'actor_info'])
    years: 최근몇년치를가져올지
    """
if site == "filmmakers":
return fetch_filmmakers(boards=boards, years=years)
else:
        log(f"아직구현되지않은사이트: {site}")
return []
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

565/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
deffetch_filmmakers(boards: list[str], years: int) -> List[Post]:
"""
필름메이커스전용크롤링로직.
    boards에들어있는게시판만크롤링한다.
    """
    ...
```

## **4.** `main.py` **의 옵션 입력도 “사이트 → 게시판” 순서로 변경** 

기존 v1.1 의사코드: 

python 코드 복사 `years, boards, save_format = ask_user_options() run_crawler(years, boards, save_format)` python 코드 복사 `site, boards, years, save_format = ask_user_options() run_crawler(site, boards, years, save_format)` 

## 수정 후: 

## 의사코드 예시: 

python 코드 복사 

```
from fetcher import fetch_posts
from cleaner import clean_post
from saver import save_csv, save_json
from logger import log
```

```
defask_user_options():
```

```
# 1) 사이트선택
# 2) 해당사이트의게시판선택
# 3) 기간선택 (3년/1년/직접입력)
# 4) 저장형식선택 (csv/json)
# 을받아서 (site_code, board_codes, years, save_format) 반환
    ...
```

```
defrun_crawler(site: str, boards: list[str], years: int, save_format: str) -> No
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

566/633 

25. 11. 30. 오후 12:53 

 

크롤러 제작 계획 

```
    log(f"사이트: {site}, 게시판: {boards}, 기간: 최근{years}년, 저장: {save_format
    posts = fetch_posts(site=site, boards=boards, years=years)
```

```
    cleaned_rows = []
for post in posts:
        row = clean_post(post)
if row isNone:
continue
        cleaned_rows.append(row)
# 저장형식분기
if save_format == "json":
        ...
else:
        ...
```

## **5. 요약 (그래비티가 꼭 지켜야 할 변경점)** 

**1. 사이트 고정 X → exe 실행 시 사이트 선택 메뉴 추가** 

**2. 게시판 고정 X → 선택한 사이트에 따라 게시판 선택 메뉴 출력** 

**3.** `fetch_posts()` 는 `(site, boards, years)` 인자 기반으로 동작 

**4.** `main.py` 는 `ask_user_options()` 에서 

site boards years save_format 을 받아서 `run_crawler()` 에 넘긴다. 

- **5.** 현재 실제 구현은 **site='filmmakers'만 완성** 해도 되고, 

 

나머지 사이트는 “준비중” 안내만 해도 된다. 

중요한 건 구조가 **“확장 가능한 형태”**로 짜여 있어야 한다는 점이다. 

## 이렇게 바꾸면: 

- 지금은 **필름메이커스 1개 사이트만 구현해도** 

- 구조는 이미 “멀티 사이트/멀티 게시판 선택형”으로 잡혀 있고, 나중에 형이 “또 다른 커뮤니티 추가해보자” 할 때 → **사이트 목록·게시판 목록만 추가하면 바로 확장 가능** 한 상태가 된다. 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

567/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

이제 이 v1.2 수정 블록을 

**그래비티 지시문 맨 뒤에 그대로 붙여넣으면** , 

더 이상 “필름메이커스 고정 크롤러”가 아니라 

- **“사이트/게시판 선택형 크롤러”**로 스펙이 바뀌는 거야. 

**==> picture [489 x 38] intentionally omitted <==**

아니 내말은 필름메이커스를 고정으로 박지 말라는 이야 기야 

게시판도 마찬가지고 

url을 통해 프로그램을 그 사이트로 적용시키라는 이야기 잖아 

18s 동안 생각함 

형 말 이제 정확히 알겠어. 

**“필름메이커스라는 이름 자체를 코드에 못 박지 마라. 게시판 이름도 못 박지 말고,** 

**내가 URL을 넣으면 그 URL 기준으로 돌게 만들어라.”** 이거잖아. 

지금까지 설계는 “엔진은 범용인데, 예시 사이트로 필름메이커스 3게시판을 먼저 쓴다”였 고, 

형이 원하는 건 **“엔진 100% 범용 + 사이트/게시판은 전부 URL 입력”**이야. 

## 그래서, 아래처럼 **그래비티용 스펙을 갈아끼운다.** 

- (= 더 이상 `필름메이커스` / `엑터스 포럼` 이런 이름은 코드에 안 박힘) 

## ✅ **핵심 수정 2줄** 

**1. 코드 안에 “필름메이커스” / “엑터스 포럼” 같은 하드코딩 전부 금지.** 

**2. exe 실행 시, 형이 직접 “크롤링할 게시판 URL 목록”을 붙여넣고 → 그 URL들 기준 으로 크롤링.** 