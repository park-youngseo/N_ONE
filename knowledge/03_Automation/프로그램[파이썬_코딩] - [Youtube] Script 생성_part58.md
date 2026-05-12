)
pause
```

이 배치는 **저장→닫기 후 자동 실행** 을 그대로 유지하면서, 어디서 막히는지 **한글 메시지로 정확히** 보여줘. 

## **C단계｜딱 한 줄로 요약(지금 바로 해볼 것)** 

**1. 확장자 보이기** 켜고, 파일명이 정확히 

`Script.bat` / `Youtube.txt.py` 인지 확인 

**2.** 같은 폴더에서 `python -V` 확인 

**3.** `Script.bat` 더블클릭 → 폴더 입력 → **메모장에 링크 붙여넣고 저장→닫기** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

285/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **4.** 결과는 `FINAL_TXT\폴더이름_01.txt …` 

여기서 멈추면, 검은창에 나온 **마지막 두세 줄** 을 그대로 붙여줘. 그 한 줄로 바로 해결해 줄게. 

**==> picture [489 x 39] intentionally omitted <==**

Microsoft Windows [Version 10.0.19045.6456] (c) Microsoft Corporation. All rights reserved. 

C:\Users\엔오원\Desktop\yotube_script>python -V Python 3.11.9 

C:\Users\엔오원\Desktop\yotube_script>python "Youtube.txt.py" --help 

usage: Youtube.txt.py [-h] --root ROOT --urls URLS [-lang LANG] [--chunk-size CHUNK_SIZE] 

options: 

-h, --help            show this help message and exit --root ROOT           결과를 저장할 상위 폴더 --urls URLS           urls.txt 경로 

--lang LANG           자막 언어(기본 ko) --chunk-size CHUNK_SIZE 분할 글자수(기본 40000) 

C:\Users\엔오원\Desktop\yotube_script> 

좋아 형, 지금 파이썬도 잘 나오고 **엔진 도움말도 정상** 이야. 이제 **실행만** 하면 된다. 아주 짧게, 그대로 따라 해. 

## **지금 위치** 

`C:\Users\엔오원\Desktop\yotube_script>` (여기에서 계속 진행) 

## **1) (한 번만) 도구 점검** 

bat 

코드 복사 

```
python -m pip install -U yt-dlp
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

286/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **2) urls.txt 만들고 저장** 

bat 코드 복사 

```
notepad urls.txt
```

메모장이 열리면 **한 줄에 한 개** 링크를 붙여넣고 **Ctrl+S(저장) → 창 닫기** 예) 

arduino 코드 복사 `https:` _`//www.youtube.com/playlist?list=PLxxxxxxxxxxxxxxxx`_ `https:` _`//youtu.be/AAAAAAAAAAA`_ 

## **3) 실행 (지금 폴더를 결과 폴더로 사용)** 

bat 코드 복사 `python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size 40000`   **4) 결과 확인** 끝나면 여기 생김: makefile 코드 복사 `C:\Users\엔오원\Desktop\yotube_script\FINAL_TXT\ yotube_script_01.txt yotube_script_02.txt ...` 

## **4) 결과 확인** 

각 파일은 **4만 글자** 기준으로 나뉨. 중간 임시파일은 정리됨. 

## **문제 생기면 이 두 가지만 확인** 

에러가 `No module named yt_dlp` → 1)번 명령 다시 실행. 

어떤 영상은 자막이 없을 수 있음(그 영상은 자동 건너뜀). 

- 여기까지 하고, 검은창에 뜬 **마지막 에러/안내 두 줄** 만 보여주면 바로 이어서 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

287/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**==> picture [289 x 102] intentionally omitted <==**

이건 뭐냐? 

## 형, 지금 화면은 `urls.txt` **가 비어 있다고 판단되어 종료된 상태** 야. 

왜냐면 메모장에 있는 링크 앞에 ** `#` 기호가 붙어 있어서 “주석(무시)”**로 처리됐어. 

## **바로 고치기 (3단계)** 

**1.** 메모장(열려있는 `urls.txt` )에서 

   - `# https://www.youtube.com/playlist?list=...` 

   - 이 줄의 `#` **랑 앞의 공백** 을 지워: 

arduino 

코드 복사 

```
https://www.youtube.com/playlist?list=PLDAF4b3hGb-GzXTqdgzHtz7QncEUniNmY
```

## (한 줄에 **딱 하나의 링크** 만) 

## **2. Ctrl + S** 로 저장 → 메모장 **닫기** . 

**3.** 창이 이미 “Press any key to continue…”로 멈춰있으면 키를 눌러 닫고, **Script.bat 다 시 더블클릭** 해서 시작. 

시작되면 검은창에 **[처리] https://…** 같은 줄이 보이고, 끝나면: 

코드 복사 

```
…\FINAL_TXT\
폴더이름_01.txt
폴더이름_02.txt
  …
```

## 처럼 만들어질 거야. 

혹시 또 멈추면, 검은창 마지막 두 줄만 보내줘—그 자리에서 바로 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

288/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

raise DownloadError(message, exc_info) yt_dlp.utils.DownloadError: ERROR: [youtube] IysEDBRR3bA: Private video. Sign in if you've been granted access to this video. Use --cookies-frombrowser or --cookies for the authentication. See https://github.com/yt-dlp/yt-dlp/wiki/FAQ#how-do-ipass-cookies-to-yt-dlp  for how to manually pass cookies. Also see  https://github.com/yt-dlp/ytdlp/wiki/Extractors#exporting-youtube-cookies  for tips on effectively exporting YouTube cookies 

## 형, 에러 뜻은 간단해: **재생목록 안에 “비공개·로그인 필요한 영상”이 섞여 있어서 yt-dlp 가 막혔어.** 

고치는 방법은 두 가지 중 하나야. 

## **A단계｜빠른 해결 (선호 순서)** 

## **1. 비공개 영상 건너뛰기** 

`urls.txt` 열기 → 문제가 된 링크 줄 **앞에** `#` 붙여서 주석 처리하거나 줄을 지 워. 

저장(Ctrl+S) → `Script.bat` 다시 실행. 

(가장 간단, 로그인 불필요) 

## **2. 로그인 쿠키로 내려받기** 

유튜브에 로그인된 **브라우저 쿠키** 를 써서 비공개/연령제한 영상도 처리. 

아래 “B단계”대로 **파일에 옵션 2줄만 추가** 하면 끝. 

## **B단계｜쿠키 사용 가능하게 “아주 작은” 수정** 

## **1)** `Youtube.txt.py` **에 옵션 2개 추가** 

메모장으로 `Youtube.txt.py` 열고, 맨 위의 `argparse` 옵션 부분에 **두 줄** 추가: 

python 코드 복사 

```
"""
ap.add_argument("--cookies", default=, help="cookies.txt 경로)
"""
ap.add_argument("--cookies-from-browser", default=, help=브라우저이름(chrome/ed
```

  

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

289/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

그리고 `ydl_opts = {...}` 만든 뒤 아래 **두 줄** 을 그대로 넣어: 

python 

코드 복사 

```
if args.cookies:
```

```
    ydl_opts["cookiefile"] = str(Path(args.cookies))
```

```
if args.cookies_from_browser:
```

```
    ydl_opts["cookiesfrombrowser"] = (args.cookies_from_browser,)
```

이렇게 하면 **쿠키 파일** 을 주거나 **브라우저 이름** 만 줘도 로그인 세션을 써서 받는다. 

## **2)** `Script.bat` **에서 간단 질문 1개 추가 (선택)** 

`Script.bat` 열고, 실행 줄을 **아래처럼 바꿔** (없는 옵션은 그냥 무시됨): 

bat 코드 복사 

```
set "COOKIES_ARG="
```

```
echo [질문] 로그인필요한영상(비공개등)이있습니까? 브라우저쿠키를쓸까요? (Y/N, 기본
set /p USEC=입력:
```

```
if /I "%USEC%"=="Y" (
```

```
  echo 사용할브라우저를입력하세요 (chrome/edge/firefox/brave) 기본 chrome:
  set /p BROWSER=
```

```
  if "%BROWSER%"=="" set "BROWSER=chrome"
```

```
  set "COOKIES_ARG=--cookies-from-browser %BROWSER%"
```

```
)
```

```
echo.
```

```
echo [실행] 자막→TXT 생성및정리(4만자분할)를시작합니다...
```

```
python "%HERE%Youtube.txt.py" --root "%ROOT%" --urls "%URLS%" --lang ko --chunk-s
```

  **크롬** 쓰면 `chrome` , **엣지** 면 `edge` 입력. 

- 이 방법은 별도의 cookies.txt 만들 필요가 없어서 가장 쉬워. 

## **C단계｜다시 실행 순서 (초간단)** 

**1. 가장 쉬운 방법:** 비공개 영상 줄에 `#` 붙여 주석 처리 → 저장 → `Script.bat` 실행. 
