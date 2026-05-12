
## **최후 수단(배치 없이 바로 실행)** 

검은창에서(폴더가 yotube_script인 상태): 

arduino 코드 복사 

`python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size 40000`   비공개가 섞였으면: 

css 코드 복사 

`python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size 40000`   어디서 멈췄는지 **검은창 마지막 3~4줄** 만 복사해서 주면, 그 지점부터 바로 고쳐줄게. 

**==> picture [489 x 39] intentionally omitted <==**

## 야.. 씨발놈아 아무것도 안되잖아 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 304/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

9s 동안 생각함 

형, 미안. 지금 바로 **되는 루트** 만 딱 적을게. 그대로 복붙해서 해봐. 

## **0) 폴더 이동** 

bash 코드 복사 

```
cd"C:\Users\엔오원\Desktop\yotube_script"
```

## **1) 파일/파이썬 점검 (3줄)** 

bash 코드 복사 

```
dir /b Youtube.txt.py Script.bat
python -V
python -c "import yt_dlp; print('yt_dlp OK')"
```

`Youtube.txt.py` 가 안 보이면 **파일명이 틀린 거** 야(.py.txt 아님 확인). `yt_dlp OK` 가 안 뜨면: nginx 코드 복사 

```
python -m pip install -U yt-dlp
```

## **2) urls.txt 만들고 저장** 

nginx 코드 복사 

```
notepad urls.txt
```

메모장에 **유튜브 링크를 한 줄에 하나** 씩 넣고 **Ctrl+S(저장) → 창 닫기** 줄 맨 앞에 `#` 있으면 **무시** 됨(지워야 함). 플레이리스트도 한 줄이면 됨. 

## **3) 바로 실행 (공개 영상만 있을 때)** 

arduino 

코드 복사 

```
python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size 40000
```

  

## **4) 비공개/연령제한이 섞였으면 (쿠키 사용)** 

코드 복사 

css 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

305/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size40000
```

**==> picture [477 x 36] intentionally omitted <==**

**----- Start of picture text -----**<br>
 <br>(엣지 쓰면  chrome   대신  edge  )<br>**----- End of picture text -----**<br>


## **5) 결과 위치** 

makefile 코드 복사 

```
C:\Users\엔오원\Desktop\yotube_script\FINAL_TXT\
  yotube_script_01.txt
  yotube_script_02.txt
  ...
```

## **그래도 “아무 것도 안 됨”이면** 

바로 이 3줄만 복붙해서 결과를 보여줘. 거기서 한 방에 잡아줄게. 

bash 코드 복사 `cd "C:\Users\엔오원\Desktop\yotube_script" dir /b` 

```
python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size 40000
```

**==> picture [473 x 8] intentionally omitted <==**

**----- Start of picture text -----**<br>
 <br>**----- End of picture text -----**<br>


(에러가 뜨면 **마지막 3~4줄** 만 복사해서 줘.) 

**==> picture [489 x 39] intentionally omitted <==**

야 배치 파일을 씨발 새꺄 만들게했으면 그대로 하라고 했 지? 근데 결국 씨발 새꺄 니가 또 수정하는데로 한장짜리 로 바꿔? 배치파일이 왜 안되는지 개새끼야 문제를 해결해 야할거 아냐 

30s 동안 생각함 

형, 미안. **배치가 안 도는 문제만** 정확히 잡자. 

아래 `Script.bat` 로 **갈아끼우고** 다시 실행해. (이름/구조 절대 안 바꿈: `Script.bat` + `Youtube.txt.py` 딱 2개) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 306/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **A단계｜핵심 (딱 2줄)** 

- [사실] `Script.bat` 더블클릭 → 폴더 경로 입력 → 메모장에 **링크 붙여넣고 저장 후 닫기** → 자동 실행. 

- [사실] 결과는 `정리폴더\FINAL_TXT\폴더이름_01.txt, _02.txt…` (각 4만자). 중간파일은 정리. 

## **B1｜문제 재발 방지용** `Script.bat` **(그대로 덮어쓰기)** 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title Script — 폴더지정→urls.txt 편집→자동실행(자막→TXT→합치기→4만자분할)
```

```
rem 0) Python 확인
where python >nul 2>nul || (
```

```
  echo [오류] Python이인식되지않습니다. 설치시 "Add python.exe to PATH" 체크가필요
  pause & exit /b
)
rem 1) 엔진확인 (고정이름)
set "HERE=%~dp0"
if not exist "%HERE%Youtube.txt.py" (
  echo [오류] Youtube.txt.py 가이배치와같은폴더에없습니다. 현재: "%HERE%"
  pause & exit /b
)
rem 2) 정리(저장)할폴더지정
set "ROOT="
set /p ROOT=정리할폴더경로(엔터=현재폴더):
if "%ROOT%"=="" set "ROOT=%HERE%"
set "ROOT=%ROOT:"=%"
if not exist "%ROOT%" (
  echo [안내] 폴더가없어새로만듭니다: "%ROOT%"
  md "%ROOT%" >nul 2>nul
)
rem 3) urls.txt 열기 (저장후닫으면자동진행)
set "URLS=%ROOT%\urls.txt"
if not exist "%URLS%" (
  > "%URLS%" echo # 한줄에한개의유튜브링크(재생목록가능)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 307/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
  >> "%URLS%" echo # 예시:
  >> "%URLS%" echo # https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX
  >> "%URLS%" echo # https://youtu.be/AAAAAAAAAAA
  >> "%URLS%" echo # 저장(Ctrl+S) 후창을닫으면계속진행됩니다.
)
echo.
echo [열기] 메모장이열립니다. 링크를붙여넣고 **저장(Ctrl+S)** 후창을닫으세요.
start /wait notepad "%URLS%"
rem 4) urls.txt 비어있는지검사(주석/빈줄제외)
set "HASURL="
for /f "usebackq tokens=* delims=" %%L in ("%URLS%") do (
  set "LINE=%%L"
  set "LINE=!LINE: =!"
  if not "!LINE!"=="" if not "!LINE:~0,1!"=="#" set "HASURL=1" & goto _ok
)
:_ok
if not defined HASURL (
  echo [오류] urls.txt에처리할링크가없습니다. (#로시작하는줄은무시)
  pause & exit /b
)
rem 5) 비공개/연령제한대비(브라우저쿠키선택)
set "COOKIES_ARG="
echo.
set /p ANS=[질문] 비공개/연령제한영상이있나요? 브라우저쿠키사용? (Y/N, 기본 N):
if /I "%ANS%"=="Y" (
  set /p BROWSER=브라우저(chrome/edge/firefox/brave) [엔터=chrome]:
  if "%BROWSER%"=="" set "BROWSER=chrome"
  set "COOKIES_ARG=--cookies-from-browser %BROWSER%"
)
rem 6) yt-dlp 업데이트(조용히시도)
python -m pip install -U yt-dlp >nul 2>nul
```

```
rem 7) 실행
echo.
echo [실행] python "Youtube.txt.py" --root "%ROOT%" --urls "%URLS%" --lang ko --ch
python "Youtube.txt.py" --root "%ROOT%" --urls "%URLS%" --lang ko --chunk-size 400
if errorlevel 1 (
  echo.
```

```
  echo [오류] 실행중문제가발생했습니다. 위메시지마지막줄을복사해서알려주세요.
  pause & exit /b
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 308/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
)
```

```
echo.
```

```
echo [완료] 결과: "%ROOT%\FINAL_TXT\폴더이름_01.txt, _02.txt ..." (각 4만자)
pause
```

## **B2｜실행 체크(초간단)** 

**1.** `Script.bat` 와 `Youtube.txt.py` **같은 폴더** 에 있는지 확인. 

**2.** `Script.bat` **더블클릭** → 폴더 경로 입력. 

**3.** 뜬 메모장에 **링크 붙여넣기 → Ctrl+S(저장) → 창 닫기** . 

**4.** 비공개 있으면 **Y → chrome/edge** . 

**5.** 창에 **[실행] python "Youtube.txt.py" …** 뜨고 진행되면 정상. 

여기까지 했는데도 안 돌면, 배치 창의 **마지막 3~4줄** 그대로 보내줘. 그 줄 기준으로 바로 고친다. 

**==> picture [489 x 39] intentionally omitted <==**

raise DownloadError(message, exc_info) yt_dlp.utils.DownloadError: ERROR: [youtube] zwNTiZr4jvI: Video unavailable. This content isn't available, try again later. The current session has been rate-limited by YouTube for up to an hour. It is recommended to use `-t sleep` to add a delay between video requests to avoid exceeding the rate limit. For more information, refer to  https://github.com/yt-dlp/ytdlp/wiki/Extractors#this-content-isnt-available-try-againlater 

C:\Users\엔오원\Desktop\yotube_script> 

19s 동안 생각함 

형, 이 에러는 **유튜브가 요청을 너무 자주 보냈다고 판단해서 “잠깐 멈춰!”(rate limit)** 한 상황이야. 

해결은 두 가지가 바로 먹혀: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 309/341 
