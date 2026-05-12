
bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브 → TXT(정리) → 4만자청크 - URL/폴더만질문, 언어=ko·VTT삭제=Y 자동
```

```
REM 0) Python 확인
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
```

```
REM 1) 엔진(.py) 자동탐색 (이름바뀌어도찾음)
set "HERE=%~dp0"
set "SCRIPT="
```

```
if exist "%HERE%yt_to_txt_chunks.py" set "SCRIPT=%HERE%yt_to_txt_chunks.py"
if not defined SCRIPT (
```

```
  for /f "delims=" %%F in ('dir /b /a:-d "%HERE%*txt*chunk*.py" 2^>nul') do (
    set "SCRIPT=%HERE%%%F"
    goto :found_script
  )
)
if not defined SCRIPT (
```

```
  for /f "delims=" %%F in ('dir /b /a:-d /o:-d "%HERE%*.py" 2^>nul') do (
    set "SCRIPT=%HERE%%%F"
    goto :found_script
  )
)
:found_script
if not defined SCRIPT (
```

```
  echo [오류] 실행할파이썬스크립트를찾지못했습니다. 폴더에 *.py(엔진) 파일을두세요
  pause & exit /b
)
echo [정보] 사용엔진: "%SCRIPT%"
```

```
REM 2) 기본값(자동고정)
set "LANG=ko"          REM 한국어자막고정
set "CHUNK=40000"      REM 4만자청크고정
set "CLEAN=--cleanup"  REM 처리후 .vtt 자동삭제
```

```
REM (선택) yt-dlp 설치/업데이트시도(조용히). 실패해도계속.
python -m pip install -U yt-dlp >nul 2>nul
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 202/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
REM 3) URL/폴더만질문
set "URL="
set /p URL=유튜브영상또는재생목록 URL:
if "%URL%"=="" echo [오류] URL이필요합니다.& pause & exit /b
```

```
set "OUT="
set /p OUT=저장폴더(비우면현재폴더):
if "%OUT%"=="" set "OUT=."
```

```
REM 4) 실행
echo.
```

```
echo [RUN] python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-si
python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-size %CHUNK% %
```

```
echo.
```

```
echo [완료] 결과: "<저장폴더>\제목 [ID]\TXT\원문.txt" 및 "TXT_chunks\01.txt, 02.txt
pause
```

## **쓰는 법 (아주 간단)** 

**1.** 위 배치를 **메모장에 붙여넣어 저장** → `yt_to_txt_chunks.py` 와 **같은 폴더** 에 두기. 

**2. 더블클릭** → 

**URL만 입력** , 

**저장 폴더만 입력** (비우면 현재 폴더). 

**3.** 끝! 결과는 `TXT` (전체 원문)와 `TXT_chunks\01.txt, 02.txt…` (각 **4만자** )로 생성. [사실] 

필요하면 **OUT 기본값을 특정 폴더로 고정** , 또는 **여러 URL을 한 번에 처리** 하는 옵션도 추 가해줄게. 

**==> picture [489 x 38] intentionally omitted <==**

여러url을 한번에 처리하는건 뭐지? 

23s 동안 생각함 

형, “여러 URL 한 번에 처리”는 **영상/재생목록 주소들을 한 큐에 넣고, 순서대로 자동 반 복 실행** 한다는 뜻이야. 한 주소 처리 끝나면 다음 주소로 넘어가서 **TXT → 4만자 청크 (01, 02, …)**까지 전부 만들어준다. [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 203/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **A단계｜개요** 

- [사실] 방식은 3가지: **①** `urls.txt` **(한 줄=한 URL)** , **② 배치에 인자로 여러 URL** , **③ 붙여넣기로 여러 줄 입력** . 

- [사실] “이전 버전” 흐름 유지: **폴더만 물어보고** , **언어=ko / VTT삭제=Y / 4만자** 는 **자 동 고정** . 

- [사실] 재생목록 URL도 넣으면 **영상별로 자동 분기** 해서 각각 결과 폴더가 생긴다. 

## **B2단계｜배치 파일(복붙해서 저장)** 

파일명 예: `유튜브→TXT_40k_MULTI.bat` 

- (같은 폴더에 `yt_to_txt_chunks.py` 가 있어야 함. 이름 바꿔도 자동 탐색함) 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브 → TXT(정리) → 4만자청크 - 여러 URL 한번에
```

```
REM 0) Python 확인
```

```
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
```

```
REM 1) 엔진(.py) 자동탐색
```

```
set "HERE=%~dp0"
```

```
set "SCRIPT="
```

```
if exist "%HERE%yt_to_txt_chunks.py" set "SCRIPT=%HERE%yt_to_txt_chunks.py"
```

```
if not defined SCRIPT (
```

```
  for /f "delims=" %%F in ('dir /b /a:-d "%HERE%*txt*chunk*.py" 2^>nul') do (
```

```
    set "SCRIPT=%HERE%%%F"
```

```
    goto :found
  )
```

```
)
```

```
if not defined SCRIPT (
```

```
  for /f "delims=" %%F in ('dir /b /a:-d /o:-d "%HERE%*.py" 2^>nul') do (
    set "SCRIPT=%HERE%%%F"
```

```
    goto :found
  )
)
:found
```

```
if not defined SCRIPT (
```

```
  echo [오류] 실행할파이썬스크립트를찾지못했습니다. 폴더에 *.py(엔진) 파일을두세요
  pause & exit /b
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 204/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
)
```

```
echo [정보] 사용엔진: "%SCRIPT%"
```

```
REM 2) 기본값(자동고정)
set "LANG=ko"          REM 한국어자막
set "CHUNK=40000"      REM TXT 4만자분할
set "CLEAN=--cleanup"  REM VTT 자동삭제
```

```
REM (선택) yt-dlp 업데이트시도(조용히). 실패해도계속.
python -m pip install -U yt-dlp >nul 2>nul
```

```
REM 3) 저장폴더만질문 (이전방식유지)
set "OUT="
set /p OUT=저장폴더(비우면현재폴더):
if "%OUT%"=="" set "OUT=."
```

```
echo.
```

```
echo [안내] 여러 URL 처리모드입니다.
```

```
echo  - 방법 A) 이창에 URL을여러줄로붙여넣고, 빈줄(Enter만)로종료
```

```
echo  - 방법 B) 같은폴더의 urls.txt(한줄=한 URL) 사용
```

```
echo  - 방법 C) 이배치를 "배치.bat URL1 URL2 ..." 처럼인자로실행
echo.
```

```
REM 4) 우선순위: (C) 인자 → (B) urls.txt → (A) 붙여넣기
set "HAS_ARG="
```

```
if not "%~1"=="" set "HAS_ARG=YES"
```

```
if defined HAS_ARG goto :process_args
```
