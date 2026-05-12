set "LANG="
set /p LANG=자막언어(기본 ko):
if "%LANG%"=="" set "LANG=ko"
set "CLEAN="
set /p CLEAN=처리후 .vtt 삭제? (Y/N, 기본 Y):
if /I "%CLEAN%"=="N" (set "CLEAN=") else (set "CLEAN=--cleanup")
```

```
REM 3) 실행 (TXT 4만자고정)
echo.
echo [실행] python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-si
python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-size 40000 %C
```

```
echo.
```

```
echo [완료] TXT_chunks 폴더의 01.txt부터순서대로올려주세요.
pause
```

## **무엇이 달라졌나? (요점 3개)** 

**1. 이름 바꿔도 작동** : `yt_to_txt_chunks.py` 를 `YouTube_to_Text.py` 로 바꿔도, **파일명에** `txt` **와** `chunk` **가 들어있으면 자동 인식** 해서 실행. 

**2. 최신본 자동 선택** : 위 패턴이 없으면 **가장 최근에 수정된 .py** 파일을 선택(마지막 수 단). 

**3. 가시성** : 실행 전에 **[정보] 사용 엔진: "….py"** 를 보여줘서 뭘 실행하는지 분명히 확인 가능. 

## — 사용법은 동일 — 

- 더블클릭 → URL/폴더/언어/삭제 여부만 입력 → **TXT 정리 + 4만자 조각** 자동 생성. 결과 구조: 

css 코드 복사 

```
<저장폴더>/
제목[ID]/TXT/제목[ID].txt
제목[ID]/TXT_chunks/01.txt, 02.txt …
```

원하면 이 자동탐색 배치에 **기본값 고정(언어=ko, VTT 삭제=Y)** 도 넣어서 질문 없이 바 로 실행되게 만들어줄 수도 있어. “질문 없이 바로”라고 하면 그 버전으로 만들어줄게, 형. 

**==> picture [489 x 38] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

192/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## 질문없어도 바로 

30s 동안 생각함 

## 형, “질문 없이 바로” 도는 **완전 자동 배치** 넣어줄게. 

**URL만 준비** 돼 있으면 더블클릭 한 번에 **자막→TXT 정리→4만 자 청크(01/02/03…)**까 지 끝. 

## **A단계｜핵심** 

- [사실] **엔진 스크립트 이름 자동탐색** → `yt_to_txt_chunks.py` 를 찾고, 없으면 같은 폴더의 최신 `.py` 를 사용. 

- [사실] **무질문·무프롬프트** : 기본값 고정( **언어=ko, 청크=40,000자, VTT 자동삭제=Y, 저장폴더=현재 폴더** ). 

- [사실] **URL 입력 방식 3종** 지원(질문 없음): ① `urls.txt` 파일 ② 배치에 드래그&드 롭(텍스트파일/URL) ③ 명령줄 인자. 

## **B2단계｜배치 파일 (그대로 저장)** 

메모장에 붙여넣고 `유튜브→TXT_40k_AUTO.bat` 로 저장(엔진 `.py` 랑 같은 폴더). 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브 → TXT(정리) → 4만자청크(01/02/03…) - 완전자동
```

```
REM ===== 0) Python 확인 =====
```

```
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
```

```
REM ===== 1) 엔진(.py) 자동탐색 =====
set "HERE=%~dp0"
set "SCRIPT="
```

```
REM 1-1) 권장이름우선
```

```
if exist "%HERE%yt_to_txt_chunks.py" set "SCRIPT=%HERE%yt_to_txt_chunks.py"
```

```
REM 1-2) 이름바뀌었어도 'txt' 'chunk' 포함한 .py 우선
```

```
if not defined SCRIPT (
```

```
  for /f "delims=" %%F in ('dir /b /a:-d "%HERE%*txt*chunk*.py" 2^>nul') do (
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

193/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    set "SCRIPT=%HERE%%%F"
    goto :found_script
  )
)
```

```
REM 1-3) 마지막수단: 최신 .py
if not defined SCRIPT (
  for /f "delims=" %%F in ('dir /b /a:-d /o:-d "%HERE%*.py" 2^>nul') do (
    set "SCRIPT=%HERE%%%F"
    goto :found_script
  )
)
```

```
:found_script
if not defined SCRIPT (
```

```
  echo [오류] 실행할파이썬스크립트를찾지못했습니다. 폴더에 *.py 파일을두세요.
  exit /b 2
```

```
)
echo [정보] 사용엔진: "%SCRIPT%"
```

```
REM ===== 2) 기본값(질문없이고정) =====
set "OUT=."            REM 저장폴더 = 현재폴더
set "LANG=ko"          REM 자막언어 = 한국어
set "CHUNK=40000"      REM TXT 조각크기 = 4만자
set "CLEAN=--cleanup"  REM 처리후 VTT 자동삭제
```

```
REM (선택) yt-dlp 설치/업데이트시도(조용히). 실패해도계속.
python -m pip install -U yt-dlp >nul 2>nul
```

```
REM ===== 3) URL 소스수집(무프롬프트) =====
REM 우선순위: (A) 인자 → (B) 드롭된파일/URL → (C) HERE\urls.txt
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
  echo [오류] URL 입력이없습니다.
```

```
  echo        방법 A) 이배치에 URL을인자로주기:  유튜브→TXT_40k_AUTO.bat "https:/
  echo        방법 B) 이배치위로 urls.txt(한줄에한개 URL) 드래그&드롭
  echo        방법 C) 같은폴더에 urls.txt 생성후더블클릭
  exit /b 3
)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 194/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
REM ===== 4) 실행루틴 =====
```

```
REM 4-A) 인자로 URL/파일이온경우: 하나씩처리
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
REM 4-B) urls.txt 파일경로가지정된경우
if defined URLS_FILE call :process_file "%URLS_FILE%"
echo [완료] 처리종료.
exit /b 0
REM ===== 서브루틴: 파일에서 URL 읽기 =====
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
REM ===== 서브루틴: 단일 URL 실행 =====
:process_url
set "URL=%~1"
if "%URL%"=="" goto :eof
echo.
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 195/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

```
echo [RUN] python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-si
python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-size %CHUNK% %
goto :eof
```

## **어떻게 쓰냐? (질문 없이 3가지 방식)** 

## **1. urls.txt 방식(추천, 완전 자동)** 

배치랑 같은 폴더에 `urls.txt` 만들고, **한 줄에 하나** 씩 URL 적기. 예) 

arduino 코드 복사 `https:` _`//youtu.be/AAAA...`_ `https:` _`//youtu.be/BBBB...`_ `https:` _`//www.youtube.com/playlist?list=PL...`_ 

배치 **더블클릭** → 자동 처리 시작(아무 질문 없음). **2. 드래그&드롭** `urls.txt` 를 이 배치 **아이콘 위로 끌어다 놓기** → 바로 실행. **3. 명령줄 인자** arduino 코드 복사 `유튜브→TXT_40k_AUTO.bat "https://youtu.be/AAAA..." "https://youtu.be/BBBB..."`   또는 `urls.txt` 경로를 인자로 줘도 됨. **결과(항상 동일 구조)** css 코드 복사 `<현재폴더 또는 OUT>/ 제목 [영상ID]/ TXT/제목 [ID].txt TXT_chunks/01.txt, 02.txt, 03.txt ...` 

## **결과(항상 동일 구조)** 

[사실] 각 청크는 **4만 자** 기준. **VTT는 자동 삭제** 됨(내부에서 사용만). 

## **설치(최초 1회만)** 

css 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

196/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 
