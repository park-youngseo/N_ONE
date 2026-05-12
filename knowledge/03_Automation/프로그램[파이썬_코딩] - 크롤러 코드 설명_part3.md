```
# ============================================
```

```
[fetcher.py 파일내용시작]
```

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
def fetch_generic_board(url: str) -> list:
    posts = []
    log(f"크롤링시작: {url}")
    try:
        page = 1
        while True:
            page_url = f"{url}?page={page}"
            r = requests.get(page_url, headers=headers, timeout=10)
            if r.status_code != 200:
                break
            soup = BeautifulSoup(r.text, "html.parser")
            items = soup.select("a")  # 모든링크우선수집
            if not items:
                break
            for a in items:
                title = a.text.strip()
                href = a.get("href", "")
                if not href:
                    continue
                if href.startswith("/"):
                    href = url.split("/")[0] + "//" + url.split("/")[2] + href
```

```
상세페이지
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

14/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
            page += 1
            time.sleep(random.uniform(1, 3))
```

```
    except Exception as e:
        log(f"오류발생: {e}")
    log(f"수집완료: {len(posts)}개")
    return posts
```

```
[fetcher.py 파일내용끝]
```

```
# ============================================
# 8. main.py
# ============================================
```

```
[main.py 파일내용시작]
```

```
from fetcher import fetch_generic_board
from saver import save_csv
from logger import log
```

```
def ask_user():
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

15/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    print("=============================")
    print("  엔오원니즈수집크롤러 v2")
    print("=============================\n")
    site_name = input("[1] 사이트이름입력: ").strip()
    if not site_name:
        print("사이트이름은필수입니다.")
        exit()
    print("\n[2] 크롤링할게시판 URL 입력")
    print("입력이끝났으면빈줄에서 Enter")
    urls = []
    while True:
        u = input("> ").strip()
        if u == "":
            break
        urls.append(u)
    if not urls:
        print("URL이없습니다.")
        exit()
    return site_name, urls
```

```
def run():
    site_name, urls = ask_user()
    log(f"=== {site_name} 크롤링시작 ===")
    all_rows = []
    for url in urls:
        rows = fetch_generic_board(url)
        all_rows.extend(rows)
```

```
    filename = f"{site_name}.csv"
    path = save_csv(filename, all_rows)
    log(f"=== 저장완료: {path} ===")
    print(f"\n저장완료: {path}")
```

```
if __name__ == "__main__":
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

16/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    run()
```

```
    input("\n종료하려면 Enter를누르세요.")
```

```
[main.py 파일내용끝]
```

```
# ============================================
```

```
# 9. requirements.txt
```

```
# ============================================
```

```
[requirements.txt 파일내용시작]
```

```
requests
beautifulsoup4
```

```
[requirements.txt 파일내용끝]
```

```
# ============================================
```

- `# 10. build_exe.bat (Windows용 exe 빌드 스크립트)` 

```
# ============================================
```

```
[build_exe.bat 파일내용시작]
```

```
@echo off
```

```
REM 엔오원니즈수집크롤러 v2.4 exe 빌드스크립트
REM 1) 먼저 PyInstaller를설치합니다.
REM    pip install pyinstaller
REM 2) 이스크립트를프로젝트폴더( main.py 가있는위치 )에서실행하면
REM    단일 exe 파일(n1_crawler.exe)을생성합니다.
```

```
pyinstaller --noconfirm --onefile --name n1_crawler main.py
```

```
echo.
```

```
echo 빌드가완료되면 dist 폴더안에 n1_crawler.exe 가생성됩니다.
pause
```

```
[build_exe.bat 파일내용끝]
```

```
# ============================================
```

## `# 그래비티를 위한 추가 안내` 

```
# ============================================
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

17/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

- `# - 위에 정의된 모든 .py 파일, requirements.txt, build_exe.bat 를 같은 프로젝트 루트에` 

- `# - output / logs 폴더는 코드 실행 시 자동 생성되도록 되어 있으므로, 별도 수동 생성은 필 # - 사용자는:` 

- `#   1) python -m pip install -r requirements.txt` 

- `#   2) build_exe.bat 실행` 

- `#   을 통해 Windows용 단일 exe (n1_crawler.exe)를 만들 수 있어야 한다.` 

**==> picture [489 x 39] intentionally omitted <==**

지금부터 그래비티에 이파일을 넣어서 exe파일 까지 뽑아 내는 작업을 할거야 

하나씩 하나씩 순차적으로 무엇을 어떻게 해야하고 설정 해야 하며 완성까지 할수 있는지 뽑아줘 

오케이 형, 이제 **이 코드 묶음을 그래비티에 넣어서** `n1_crawler.exe` **까지 뽑는 전체 과정** 을, 

진짜 하나씩 “이 버튼 누르고 → 여기 붙여넣고” 수준으로 끝까지 깔끔하게 정리해볼게. 

형이 이미 갖고 있는 건: 

- 각 파일 내용( `cleaner.py` , `tagger.py` , `patterns.py` , `post.py` , `saver.py` , 

- `logger.py` , `fetcher.py` , `main.py` ) 

- 거기에 내가 추가한 

   - `requirements.txt` 

   - `build_exe.bat` 

지금부터는 **“그래비티 내부 작업(코드 생성)” + “형 컴퓨터에서 하는 작업(실행/빌드)”** 두 덩어리로 보면 편해. 

## **0단계. 오늘 최종 목표 정리** 

**1.** 그래비티 안에서 

   - 이 프로젝트 구조/설명 + 각 파일 내용을 **한 번에 먹여서** 

   - `fetcher.py` , `main.py` , `requirements.txt` , `build_exe.bat` 까지 **코드 생성 흐름** 

   - **을 자동화하는 프롬프트** 를 만든다. 