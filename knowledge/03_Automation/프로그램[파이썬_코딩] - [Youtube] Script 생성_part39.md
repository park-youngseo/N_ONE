 

```
python -m pip install --upgrade pip
python -m pip install -U yt-dlp
```

이제부터는 **질문 없이** `urls.txt` 만 놔두고 배치 더블클릭하면 끝이야. 

필요하면 **저장폴더(OUT) 기본값** 이나 **언어/lang** 도 배치 맨 위 고정값을 바꿔줄 수 있어. 

**==> picture [489 x 39] intentionally omitted <==**

아니 예전방식은 동일하고 기본값 고정은 물어보지 말고 실행하면 되는거잖아 

17s 동안 생각함 

맞아. 그래서 **“예전 방식은 그대로”** 두고, **기본값(언어=ko, 청크=40,000자, VTT삭제=Y, 저장=현재 폴더)** 은 **묻지 않고 바로 실행** 하는 **세미AUTO 배치** 를 줬어. 

→ 즉, **URL만** 필요하면 물어보고, 나머지는 **고정값으로 바로 진행** . (인자나 `urls.txt` 가 있으면 질문도 없이 즉시 실행) 

## **A단계｜핵심** 

- [사실] 엔진 자동탐색 유지(예전 방식 동일): `yt_to_txt_chunks.py` 우선 → 

- `*txt*chunk*.py` → 최신 `.py` . 

- [사실] 기본값 고정: `LANG=ko` , `CHUNK=40000` , `CLEAN=--cleanup` , `OUT=.` 

- [사실] 동작: (1) 인자/ `urls.txt` 있으면 **무질문 자동 실행** , (2) 없으면 **URL만 한 번** 물 어본 뒤 즉시 실행. 

## **B2단계｜배치 파일 (복붙 저장)** 

파일명: `유튜브→TXT_40k_세미AUTO.bat` (엔진 `.py` 와 같은 폴더) 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브 → TXT(정리) → 4만자청크(01/02/03…) - 세미AUTO
```

```
REM 0) Python 확인
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

197/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
REM 1) 엔진(.py) 자동탐색 (예전방식유지)
set "HERE=%~dp0"
set "SCRIPT="
if exist "%HERE%yt_to_txt_chunks.py" set "SCRIPT=%HERE%yt_to_txt_chunks.py"
if not defined SCRIPT (
  for /f "delims=" %%F in ('dir /b /a:-d "%HERE%*txt*chunk*.py" 2^>nul') do (
    set "SCRIPT=%HERE%%%F"
    goto :found_script
  )
)
if not defined SCRIPT (
  for /f "delims=" %%F in ('dir /b /a:-d /o:-d "%HERE%*.py" 2^>nul') do (
    set "SCRIPT=%HERE%%%F"
    goto :found_script
  )
)
:found_script
if not defined SCRIPT (
  echo [오류] 실행할파이썬스크립트를찾지못했습니다. 폴더에 *.py 파일을두세요.
  pause & exit /b
)
echo [정보] 사용엔진: "%SCRIPT%"
```

```
REM 2) 기본값(고정, 질문없음)
set "OUT=."
set "LANG=ko"
set "CHUNK=40000"
set "CLEAN=--cleanup"
```

```
REM (선택) yt-dlp 설치/업데이트시도(조용히)
python -m pip install -U yt-dlp >nul 2>nul
```

```
REM 3) URL 소스: (A) 인자 → (B) 드롭된파일/URL → (C) HERE\urls.txt → (D) 마지막으로
set "HAVE_URLS="
```

```
if not "%~1"=="" set "HAVE_URLS=YES"
```

```
if "%HAVE_URLS%"=="" if exist "%HERE%urls.txt" set "URLS_FILE=%HERE%urls.txt" & s
```

```
if "%HAVE_URLS%"=="" (
```

```
  set "URL="
```

```
  set /p URL=유튜브영상또는재생목록 URL(기본값은고정이므로 URL만입력):
  if not "%URL%"=="" (
```

```
    call :process_url "%URL%"
완료] 처리종료.
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 198/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    pause & exit /b
  ) else (
    echo [안내] URL이없어실행을종료합니다. (인자/urls.txt/드롭지원)
    pause & exit /b
  )
)
:args_loop
if "%~1"=="" goto after_args
set "ITEM=%~1"
if exist "%ITEM%" (
  call :process_file "%ITEM%"
) else (
  call :process_url "%ITEM%"
)
shift
goto args_loop
:after_args
if defined URLS_FILE call :process_file "%URLS_FILE%"
echo [완료] 처리종료.
pause & exit /b
:process_file
set "FILE=%~1"
if not exist "%FILE%" goto :eof
for /f "usebackq delims=" %%L in ("%FILE%") do (
  set "LINE=%%L"
  if not "!LINE!"=="" (
    echo !LINE! | findstr /b /c:"#">nul
    if errorlevel 1 (
      call :process_url "%%L"
    )
  )
)
goto :eof
:process_url
set "URL=%~1"
if "%URL%"=="" goto :eof
echo.
echo [RUN] python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-si
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

199/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-size %CHUNK% %
goto :eof
```

## **쓰는 법(요약)** 

## **그대로 더블클릭** 

   - 인자/ `urls.txt` 있으면 **바로 실행** (아무것도 안 물음). 

   - 없으면 **URL 한 번만** 물어보고 즉시 실행(기본값은 고정이라 더 안 물음). 

- **드래그&드롭** : `urls.txt` (줄당 1개 URL)를 배치 아이콘 위에 드롭 → 즉시 실행. **명령줄 인자** : 여러 URL을 한 번에 

vbnet 코드 복사 

```
유튜브→TXT_40k_세미AUTO.bat "https://youtu.be/AAA""https://youtu.be/BBB"
```

## **결과(항상 동일)** 

css 코드 복사 

```
<현재폴더>/
제목[영상ID]/
    TXT/제목[ID].txt
    TXT_chunks/01.txt, 02.txt, ...
```

원하면 동일 로직으로 **완전 AUTO판** (URL도 묻지 않고 `urls.txt` 만 읽는 버전)과 **완전 수 동판** (모두 묻는 버전)을 나란히 두게도 정리해줄게. 

**2/2** 

이전 버전은 뭐였어? 

9s 동안 생각함 

형, 바로 정리할게. **“이전 버전”**은 아래 두 가지였어. 

## **A단계｜핵심** 

- [사실] **원클릭판(유튜브→TXT_40k_원클릭.bat)** : URL·저장폴더·언어·VTT삭제 여부를 **순서대로 물어보는** 인터랙티브 방식. 엔진 파일명은 **고정(yt_to_txt_chunks.py)** 이었 어. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 200/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

- [사실] **자동탐색 원클릭판(유튜브→TXT_40k_원클릭_자동탐색.bat)** : 위와 똑같이 **질문 은 하되** , 엔진(.py) **파일명을 자동으로 찾아 실행** 했어. 

## **B2단계｜차이 한눈에** 

## **질문 수** 

- 이전(원클릭/자동탐색 원클릭): **여러 질문** (URL, 폴더, 언어, VTT삭제). 지금(세미AUTO): **URL만 물음** (인자·urls.txt 있으면 이것도 안 물음). 

## **엔진 파일명** 

원클릭: **고정 이름** 만 호출. 

- 자동탐색·세미AUTO: **자동 탐색(우선순위 규칙)** 으로 이름 바뀌어도 실행. 

## **기본값 고정** 

원클릭/자동탐색 원클릭: 매번 값 입력. 

세미AUTO: **LANG=ko / CHUNK=40000 / CLEAN=Y / OUT=현재폴더** 로 **묻지 않고 고정** . 

## **URL 입력 방식** 

이전판: **항상 대화형** 입력 중심. 

세미AUTO: **인자/드롭/urls.txt** 지원 → 있으면 **무질문 즉시 실행** . 

원하면 완전 자동판(질문 전혀 없음, `urls.txt` 만 읽는 버전)도 계속 같이 둘 수 있어. 

**==> picture [489 x 38] intentionally omitted <==**

내가 원하는건 이전 버전의 여러질문에서 (url,폴더)는 기 존처럼 진행하고 (언어,vtt삭제)는 자동으로 ko, vtt삭제를 원하는거야 언어는 당연하게 한국어로 해야하고 vtt파일은 필요가 없으니까 

16s 동안 생각함 

형, 딱 형이 원하는 버전으로 정리했어. 

**URL·폴더만 물어보고** , **언어=ko / VTT삭제=Y / 4만자 분할** 은 **자동 고정** 이야. 

## **A단계｜개요** 

- [사실] 아래 배치를 쓰면: **URL 입력 → 저장 폴더 입력** 만 하고 바로 실행. 

- [사실] 고정값: `LANG=ko` , `CLEAN=--cleanup` , `CHUNK=40000` , `OUT=입력값(기본 현재 폴 더)` . 

## **B2단계｜배치 파일 (복붙 저장)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

201/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

파일 이름: `유튜브→TXT_40k_원클릭_KO_CLEAN.bat` 

(같은 폴더에 `yt_to_txt_chunks.py` 가 있어야 하고, 이름을 바꿔도 자동 탐색됨) 