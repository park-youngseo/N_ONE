
```
defrun_crawler():
"""
필름메이커스오디션게시판을크롤링해서
결과를 list[dict] 형태로반환하는간단버전.
지금은테스트용으로 MAX_PAGES까지만돌게만든다.
    """
    results = []
# TODO: total_pages는나중에실제페이지수를계산하는로직으로교체
    total_pages = MAX_PAGES
```

```
for page inrange(1, min(total_pages, MAX_PAGES) + 1):
        list_url = f"{BASE_URL}?page={page}"
print(f"[INFO] Fetching list page {page}: {list_url}")
```

```
        html = fetch_list_page(list_url)
ifnot html:
print(f"[WARN] Failed to fetch list page {page}")
continue
```

```
        items = parse_list(html)
print(f"[INFO] Page {page}: {len(items)} items")
```

```
for item in items:
```

```
# detail_url이있으면상세요청도해볼수있지만,
# 일단은리스트정보만저장하는테스트버전으로둔다.
            results.append(item)
```

```
return results
```

## ⚠ **중요한 포인트** 

`def run_crawler():` 줄은 **맨 왼쪽(0칸 들여쓰기)** 에 있어야 해. 

- 그 아래 줄들은 

`results = []` 부터 시작해서 전부 **한 번 들여쓰기(스페이스 4칸 또는 탭 1개)** 

- 그 안의 `for page ...:` 줄도 한 번 들여쓰기 

`list_url = ...` , `html = ...` 같은 줄은 그보다 한 단계 더 들여쓰기 ( **두 번** ) 

https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

50/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

 

## VSCode에서 보기 쉽게 하려면: 

상단 메뉴 `View > Render Whitespace` 켜놓으면 

들여쓰기 칸이 점으로 보여서 맞추기 편해. 

## **3단계 – 맨 아래 메인 블록 확인** 

`run_crawler` 바로 아래가 이렇게 되어 있으면 좋아: 

python 코드 복사 `if __name__ == "__main__": data = run_crawler() print(f"[RESULT] 총 {len(data)}개 수집") print(data[:3])` 

이 블록도 **맨 왼쪽(0칸 들여쓰기)** 이어야 하고, 

안쪽 줄 두 개는 한 번 들여쓰기. 

## **4단계 – 저장하고 다시 실행** 

**1.** `Ctrl + S` 로 저장. 

**2.** 터미널에서: 

bash 코드 복사 

```
python filmmakers_crawler.py
```

## 또 에러가 나오면, 

이번에는 `def run_crawler` **부터** `if __name__ == "__main__":` **까지** 전체를 복사해서 여기 채팅에 붙여줘. 

그 상태 그대로 보고 한 줄씩 다시 맞춰줄게. 

**==> picture [489 x 38] intentionally omitted <==**

https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

51/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

 

**==> picture [260 x 193] intentionally omitted <==**

## 좋아, 이제 화면이 어떻게 돌아가는지 딱 보이네. 

지금 상황 정리부터 할게. 

- 위에 파란 줄에 `Accept Changes / Reject` 떠 있는 건 

**안티그래비티 에이전트가** `filmmakers_crawler.py` **맨 위 부분을 이렇게 고치는 게 어 떻냐고 제안한 상태** 야. 

내용 자체는 

- BASE_URL 

- LIST_URL 

- HEADERS 

- ONE_YEAR_AGO(최근 1년 기준 날짜) 

- 같은 “설정값”이라서 그대로 써도 문제 없어. 

## **1. 우선 이 변경은** `Accept` **눌러도 된다** 

**1.** 지금 뜬 파란 바에서 `Accept Changes (Ctrl+Alt+])` 눌러. → 그러면 저 Config 블록이 파일에 확정으로 들어가. 

**2.** 그 다음에 우리가 테스트용으로 쓸 `MAX_PAGES` 상수를 하나만 더 추가하면 돼. 

## **MAX_PAGES 넣는 위치** 

지금 위쪽 코드가 대략 이렇게 보일 거야: 

python 

코드 복사 

https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

52/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

 

```
import requests
from bs4 import BeautifulSoup
import time
import random
from datetime import datetime, timedelta
```

```
# [Config] 크롤러설정
# ==============================
BASE_URL = "https://www.filmmakers.co.kr"
LIST_URL = f"{BASE_URL}/actorsAudition"
HEADERS = {
"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 .
}
# 수집기간설정 (예: 최근 1년)
ONE_YEAR_AGO = datetime.now() - timedelta(days=365)
```

  

## 여기서 **ONE_YEAR_AGO 바로 위나 아래에** 이 한 줄만 추가해 줘: 

python 코드 복사 

```
MAX_PAGES = 2# 테스트용: 최대 2페이지까지만크롤링
```

## 예시는 이렇게: 

python 코드 복사 

```
HEADERS = {
"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 .
}
```

```
MAX_PAGES = 2# 테스트용: 최대 2페이지까지만크롤링
# 수집기간설정 (예: 최근 1년)
ONE_YEAR_AGO = datetime.now() - timedelta(days=365)
```

  

## 여기까지가 상수 세팅. 

## **2.** `for page` **부분은 이 모양으로 통째로 맞춰** 

https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

53/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

 

이제 문제였던 `for page` 부분을 다시 한 번 **완성본 그대로** 줄게. 

`run_crawler()` 안에 있는 “페이지 돌리는 부분”을 아래 코드로 그냥 갈아끼우면 돼. 

## **run_crawler 안의 본문 예시** 

python 코드 복사 `def run_crawler(): """ 필름메이커스 오디션 게시판을 크롤링해서 결과를 list[dict] 형태로 반환하는 테스트 버전. """ results = []` _`# TODO:`_ `나중에 실제 total_pages 계산 로직으로 교체해도 됨 total_pages = MAX_PAGES for page in range(1, min(total_pages, MAX_PAGES) + 1): list_url = f"{LIST_URL}?page={page}" print(f"[INFO] Fetching list page {page}: {list_url}") html = fetch_list_page(list_url) if not html: print(f"[WARN] Failed to fetch list page {page}") continue items = parse_list(html) print(f"[INFO] Page {page}: {len(items)} items") for item in items:` _`#`_ `지금은 목록 정보만 저장` _`(`_ `상세 페이지는 나중에` _`)`_ `results.append(item) return results` 

## ❗ **들여쓰기 체크 포인트** 

`def run_crawler():` 는 맨 왼쪽(0칸). 

그 안에 있는 줄들( `results = []` , `total_pages = ...` , `for page ...` )은 

**한 번 들여쓰기 (탭 1번 또는 스페이스 4칸)** . 

`list_url = ...` , `html = ...` , `items = ...` 등은 `for page` 안쪽이니까 **두 번 들여쓰기** . 

## VSCode에서 들여쓰기 보기 쉽게 하려면: 

https://chatgpt.com/c/69234a95-8ccc-8322-a990-08f5b9cebecf 

54/61 

25. 11. 30. 오후 12:55 

안티그래비티 개요 정리 

 

상단 메뉴 → `View` → `Render Whitespace` 켜두면 

스페이스/탭이 표시돼서 맞추기 편해. 

## **3. 맨 아래 실행 부분도 이렇게 있는지 확인** 

파일 맨 아래쪽에 이 블럭이 있으면 테스트하기 좋다: 

python 

코드 복사 

```
if __name__ == "__main__":
    data = run_crawler()
print(f"[RESULT] 총{len(data)}개수집")
print(data[:3])
```

없으면 그대로 추가해. 

## **4. 다시 실행** 

**1.** `Ctrl + S` 로 저장. 

**2.** 터미널에서: 

bash 

코드 복사 

```
python filmmakers_crawler.py
```

실행해 보고, 

- 또 에러가 뜨면 **에러 메시지 전체 + run_crawler 함수 전체** 를 그대로 복사해서 보내줘. 

(이미지 말고 텍스트로 올려주면 내가 줄 단위로 바로 정리해줄 수 있어.) 

여기까지 하면 최소한 **IndentationError** 는 확실히 없어지고, 

그다음부터는 “실제로 필름메이커스 HTML 구조에 맞게 잘 긁히냐” 단계로 넘어갈 수 있 어. 

**==> picture [489 x 39] intentionally omitted <==**

야 그냥 다 갈아 끼울수 있게줘 씨발 찾아서 계속 에러 쳐 나잖아 
