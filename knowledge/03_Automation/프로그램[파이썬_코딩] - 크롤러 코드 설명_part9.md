                    log(f"[ERROR] 상세페이지오류: {e}")
print(f"[ERROR] 상세페이지오류: {e}")
continue
```

```
            page += 1
            time.sleep(random.uniform(1, 3))
except Exception as e:
        log(f"[ERROR] fetch_generic_board 전체오류: {e}")
print(f"[ERROR] fetch_generic_board 전체오류: {e}")
```

```
    log(f"수집완료: {len(posts)}개")
print(f"[INFO] 수집완료: {len(posts)}개")
return posts
```

## **2-2. main.py도 “진행상황 보이게” 바꾸기** 

`main.py` 열고, 내용 전체를 아래 코드로 교체 → 저장(Ctrl+S): 

python 코드 복사 

```
from fetcher import fetch_generic_board
from saver import save_csv
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 43/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
from logger import log
```

```
defask_user():
print("=============================")
print("  엔오원니즈수집크롤러 v2")
print("=============================\n")
    site_name = input("[1] 사이트이름입력: ").strip()
ifnot site_name:
print("사이트이름은필수입니다.")
        exit()
print("\n[2] 크롤링할게시판 URL 입력")
print("입력이끝났으면빈줄에서 Enter")
    urls = []
whileTrue:
        u = input("> ").strip()
if u == "":
break
        urls.append(u)
ifnot urls:
print("URL이없습니다.")
        exit()
return site_name, urls
```

```
defrun():
    site_name, urls = ask_user()
    log(f"=== {site_name}크롤링시작 ===")
print(f"[INFO] {site_name}크롤링시작, URL 개수: {len(urls)}")
    all_rows = []
for url in urls:
print(f"[INFO] 게시판크롤링시작: {url}")
        rows = fetch_generic_board(url, max_pages=1, max_posts=10)
print(f"[INFO] 게시판크롤링완료: {url}, 수집{len(rows)}개")
        all_rows.extend(rows)
ifnot all_rows:
print("[WARN] 수집된글이없습니다.")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

44/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    filename = f"{site_name}.csv"
    path = save_csv(filename, all_rows)
    log(f"=== 저장완료: {path} ===")
print(f"\n[INFO] 저장완료: {path}")
```

```
if __name__ == "__main__":
    run()
input("\n종료하려면 Enter를누르세요.")
```

## 이제 **최대 10개 글만, 1페이지만** 보고 

진행 상황이 콘솔에 계속 찍힐 거라, 멈춘 것처럼 안 보일 거야. 

## **3단계. Antigravity 터미널에서 다시 실행하기** 

이번에는 윈도우 기본 CMD 말고, **Antigravity 내부 터미널** 에서 돌리면 Ctrl+C도 더 잘 먹고, 창 닫기도 쉽다. 

**1.** Antigravity 상단 메뉴 → **Terminal → New Terminal** . 

**2.** 아래쪽 터미널이 열리면 경로가 

   - `C:\Users\psm23\Desktop\n1_crawler>` 인지 확인. 

**3.** 거기서: 

bash 코드 복사 

```
python main.py
```

## **4.** 화면 흐름: 

- `[1] 사이트 이름 입력:` → `필름메이커스` 

```
[2] 크롤링할게시판 URL 입력
```

   - `https://www.filmmakers.co.kr/actorsForum` 

   - (아무것도 안 쓰고 엔터) 

**5.** 그러면 이제 이렇게 진행이 보여야 한다: 

text 코드 복사 

```
[INFO] 필름메이커스크롤링시작, URL 개수: 1
```

```
[INFO] 게시판크롤링시작: https://www.filmmakers.co.kr/actorsForum
```

```
[INFO] 1페이지요청: ...
```

```
[DETAIL] 글요청: ...
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 45/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
...
```

   - `[INFO] max_posts(10) 도달, 수집 중단` 

   - `[INFO] 수집 완료: 10개` 

   - `[INFO] 게시판 크롤링 완료: ..., 수집 10개` 

   - `[INFO] 저장 완료: output\필름메이커스.csv` 

**6.** 실행 중에 또 멈춘 것 같으면, 

이 터미널 창에서는 **Ctrl+C** 가 거의 무조건 먹을 거야. 

그래도 안 되면 그냥 터미널 오른쪽 위에 있는 **쓰레기통 아이콘(휴지통)** 클릭해서 그 터미널 세션만 종료하면 된다. 

## **4단계. 결과 확인** 

## 실행이 끝나면: 

- `n1_crawler/output/필름메이커스.csv` 생겼는지 보고, 

- `n1_crawler/logs/crawler.log` 를 열어 보면 

- 어떤 페이지/글까지 봤는지 로그가 전부 들어 있다. 

## 형, 정리하면: 

**1.** 지금 멈춘 CMD 창은 **Esc → Ctrl+C** 시도, 안 되면 그냥 X 누르고 닫기. 

**2.** `fetcher.py` , `main.py` 를 위 코드로 갈아끼워서 

   - **1페이지 / 10개 글 / 진행상황 출력 모드** 로 만든다. 

**3.** Antigravity 터미널에서 `python main.py` 로 다시 실행한다. 

여기까지 하고 나서 

콘솔에 어떤 로그가 찍혔는지, 

- `output` / `logs` 안에 뭐가 생겼는지만 보면, 

- 다음 단계(필름메이커스 전용 파싱, 글 내용만 추리는 쪽)으로 이어갈 수 있다. 

**==> picture [489 x 39] intentionally omitted <==**

**crawler.log** 파일 

로그파일 전부를 넣어줄께 무엇때문인지 알려줘 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

46/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

- 형, 로그 보니까 **프로그램은 멈춘 게 아니라 아주 크게 한 바퀴 돌고, 이미 csv까지 저장하 고 끝난 상태** 야. 

## **1. 지금 무슨 일이 벌어진 건지** 

로그 마지막 부분 보면: crawler 

- `=== 필름메이커스 크롤링 시작 ===` 

- `크롤링 시작: https://www.filmmakers.co.kr/actorsForum` 

- 그 뒤로 **사이트 메뉴, 크루모집, 클래스, 약관, 파트너, 인스타, 유튜브까지** 

- 온갖 링크를 전부 `글 요청:` 하면서 돌고 있어. 

- 중간에 외부 사이트(영화진흥위원회 등) 접속하다가 SSL 에러도 남아 있고 

- ( `certificate verify failed` 라는 줄). 

- 맨 마지막에: 

   - `[INFO] max_posts(10) 도달, 수집 중단` 

   - `수집 완료: 10개` 

   - `=== 저장 완료: output\필름메이커스.csv ===` 

## 즉: 

**1.** `soup.select("a")` 로 **페이지 안에 있는 모든 a 태그(링크)** 를 긁어와서 

**2.** 거기 달린 URL을 **하나하나 전부 요청** 하고 있어서 

   - 필름메이커스 전체 메뉴, 카테고리, 파트너사, 외부기관까지 다 도는 상황. 

**3.** 그 작업이 끝나야 max_posts(10) 조건에 걸리고, 그제서야 csv 저장하고 종료. 

그래서 형이 보기에는: 

콘솔에 아무 글자도 안 뜨고 

- 키보드도 잘 안 먹는 것처럼 보이는 동안 

## → 사실은 **필름메이커스 전체 사이트를 한 번 싹 도는 중** 이었던 거야. 

좋은 소식은: 

- 이미 `output\필름메이커스.csv` 파일이 만들어졌고, 

- 프로그램 자체는 에러로 죽지 않고 정상 종료했다는 것. 

## **2. 진짜 문제는 “너무 많이, 엉뚱하게 크롤링하는 구조”** 

지금 fetcher 구조의 문제점은 딱 두 개야. 

## **1. 목록 페이지에서 모든 링크를 긁는다** 

메뉴 / 광고 / 외부 사이트까지 전부 포함. 

**2.** 그 링크들을 전부 `requests.get()` 해서 돌다 보니 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

47/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

시간이 오래 걸리고 

형 입장에서는 “아무 반응 없는 것처럼” 느껴지는 것. 