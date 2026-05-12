
## **3-1. 다시** `python main.py` **실행** 

**1.** CMD 창에서 다시: 

bash 코드 복사 

```
cd C:\Users\psm23\Desktop\n1_crawler
python main.py
```

**2.** 아까랑 똑같이: 

사이트 이름: `필름메이커스` 

URL: `https://www.filmmakers.co.kr/actorsForum` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

38/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

## 그리고 **다음 줄에서 빈 줄 + 엔터** . 

이번에도 화면엔 중간 진행 상황은 안 보이지만, 

우리가 `logs/crawler.log` **에 엄청 많이 찍히게 해놨다.** 

## **3-2. 실행 도중 / 실행 후에 로그 파일 열어보기** 

안티그래비티로 돌아와서: 

**1.** 왼쪽에서 `logs` 폴더가 생겼는지 확인. 

**2.** 그 안에 `crawler.log` 더블클릭해서 열어봐. 

파일 안에 예를 들어 이런 줄들이 보이면: 

text 코드 복사 

- `[2025-.. ..:..:..] 크롤링 시작: https://www.filmmakers.co.kr/actorsForum` 

- `[2025-.. ..:..:..] [INFO] 1페이지 요청: https://www.filmmakers.co.kr/actorsForum?p [2025-.. ..:..:..] [DETAIL] 글 요청: ...` 

```
...
```

```
[2025-.. ..:..:..] 수집완료: 50개
```

  

## → 이러면 **정상적으로 돌고 있는 거** 야. 

`수집 완료: N개` 까지 찍혀 있으면 끝난 거고, 

- 그 시점에는 `output/필름메이커스.csv` 도 같이 생겨 있을 거야. 

만약 로그에 

- `[WARN] 상태코드 403` / `404` 

- `[ERROR] ...` 이런 게 찍혀 있으면 

- 그 줄 내용 기준으로 우리가 다음 수정을 하면 된다. 

## **4단계. 그래도 너무 오래 걸리면?** 

지금 코드는 **최대 3페이지, 최대 50개 글** 까지만 보도록 막아놨으니까 

- 이전처럼 끝도 없이 도는 일은 거의 없을 거야. 

- 그래도 느리면 `max_pages=2` , `max_posts=20` 이런 식으로 숫자만 줄여도 된다. 

## 예: 

python 코드 복사 

```
deffetch_generic_board(url: str, max_pages: int = 2, max_posts: int = 20) -> lis
```

 https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

 

39/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

## **5단계. exe 만드는 건 그대로 진행** 

## 이제 

**1.** `python main.py` 로 한 번 수집이 잘 돌아가는 것만 확인하고, 

**2.** 그다음: 

bash 코드 복사 

```
build_exe.bat
```

실행해서 `dist/n1_crawler.exe` 만들면 된다. 

## 형, 

일단 지금은 “왜 아무 반응이 없어 보이는지”를 

**로그 + 제한** 으로 해결한 거니까 

**1.** `fetcher.py` 교체 

**2.** `python main.py` 재실행 

**3.** `logs/crawler.log` 랑 `output` 폴더만 한 번 확인 

이 세 가지만 먼저 해보고, 

`crawler.log` 내용이 어떻게 찍혔는지 그대로 가져오면 

그 다음 수정(필름메이커스 전용 최적화)까지 이어서 바로 잡아줄게. 

**==> picture [489 x 38] intentionally omitted <==**

KeyboardInterrupt  전혀 먹히지 않고 아무것도 안눌려 

34s 동안 생각함 

형, 지금 상황 두 가지가 섞여 있을 수 있어: 

**1. 콘솔이 “일시정지 모드(텍스트 선택 상태)”라서** 아무 키도 안 먹는 경우 

**2.** 실제로 크롤링이 너무 오래 돌아서, 진행상황이 안 보여서 멈춘 것처럼 느껴지는 경 우 

둘 다 한 번에 정리해 줄게. 

지금부터 **(1) 강제 종료 → (2) 코드 살짝 수정 → (3) 다시 짧게 테스트** 순서로 가자. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

40/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

## **1단계. 지금 멈춘 창 강제로 끊기** 

## **1-1. 콘솔이 “선택 모드”인지 먼저 풀기** 

윈도우 CMD는 글자 드래그하거나 클릭하면 프로그램이 **일시 정지** 돼서 

키보드가 하나도 안 먹을 때가 있어. 

**1.** 검은 창 안 아무 곳이나 **마우스로 한 번 클릭** 하고 

**2.** 키보드에서 `Esc` 한 번 눌러봐. 

**3.** 그다음에 `Ctrl + C` 한 번 더. 

그래도 안 되면 아래 방식으로 그냥 꺼버리자. 

## **1-2. 그냥 창 자체를 닫기** 

**1.** 검은 CMD 창 오른쪽 위에 **X 버튼** 클릭해서 창을 닫아. 

**2.** 또는 **Alt + F4** 눌러도 돼. 

완전히 닫히면, 그 파이썬 프로세스는 끝난 거야. 

**2단계.** `fetcher.py` **/** `main.py` **를 “짧고 눈에 보이는 버전”으로 바꾸기** 우선 **절대 오래 안 돌고, 진행이 콘솔에 찍히는 버전** 으로 바꾸자. 

Antigravity(코드 편집 화면)에서 작업하면 된다. 

## **2-1. fetcher.py 교체 (페이지 1장, 글 10개까지만)** 

Antigravity 왼쪽에서 `fetcher.py` 클릭 → 안의 코드 전부 지우고 

아래 코드 **그대로 복붙** 해서 넣고 저장(Ctrl+S): 

python 코드 복사 

```
import requests
from bs4 import BeautifulSoup
import time
import random
from post import process_post
from logger import log
```

```
headers = {
"User-Agent": "Mozilla/5.0 (Windows NT 10.0)"
}
```

```
deffetch_generic_board(url: str, max_pages: int = 1, max_posts: int = 10) -> list
"""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

41/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    url: 게시판기본주소 (예: https://www.filmmakers.co.kr/actorsForum)
    max_pages: 몇페이지까지만돌지 (안전장치)
    max_posts: 최대몇개글까지만수집할지 (안전장치)
    """
    posts = []
    log(f"크롤링시작: {url}")
print(f"[INFO] 크롤링시작: {url}")
try:
        page = 1
while page <= max_pages:
            page_url = f"{url}?page={page}"
            log(f"[INFO] {page}페이지요청: {page_url}")
print(f"[INFO] {page}페이지요청: {page_url}")
            r = requests.get(page_url, headers=headers, timeout=10)
if r.status_code != 200:
                log(f"[WARN] 상태코드{r.status_code}로중단")
print(f"[WARN] 상태코드{r.status_code}로중단")
break
```

```
            soup = BeautifulSoup(r.text, "html.parser")
            items = soup.select("a")  # 모든링크우선수집
ifnot items:
                log("[INFO] 링크가더이상없어중단")
print("[INFO] 링크가더이상없어중단")
break
for a in items:
                title = a.text.strip()
                href = a.get("href", "")
ifnot href:
continue
if href.startswith("/"):
                    href = url.split("/")[0] + "//" + url.split("/")[2] + href
```

```
try:
f"[DETAIL] 글요청: {href}")
print(f"[DETAIL] 글요청: {href}")
10)
"html.parser")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 42/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

`body = ds.get_text(" ", strip=True)`  `comments = ""` 

```
                    row = {
"title": title,
"body": body,
"comments": comments
                    }
                    processed = process_post(row)
if processed:
                        posts.append(processed)
iflen(posts) >= max_posts:
                            log(f"[INFO] max_posts({max_posts}) 도달, 수집중단")
print(f"[INFO] max_posts({max_posts}) 도달, 수집중단"
                            log(f"수집완료: {len(posts)}개")
print(f"[INFO] 수집완료: {len(posts)}개")
return posts
```

```
except Exception as e: