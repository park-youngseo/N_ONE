  echo [DL] %%U
  yt-dlp %COOK% --write-auto-subs --sub-langs "ko" --skip-download "%%U"
)
echo.
echo [2/3] TXT만생성중...
python subs_cleaner_plus.py --no-srt
echo.
echo [3/3] 폴더자동정리중...
python organize_outputs.py
echo.
echo [완료] out/ 아래에서확인하세요.
pause
```

## **작동 원리 요약** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

87/341 

25. 11. 30. 오후 1:07 

 

   - 프로그램[파이썬/코딩] - [Youtube] Script 생성 

- **yt-dlp -J** 로 **채널/제목/ID/날짜** 를 읽고, 

- 만들어진 `…clean.txt` 를 **ID/제목 매칭으로 찾아** 

`out/채널명/날짜_제목 [ID]/` 에 **이동** . (옵션: 주제별 사본 `out_by_topic/` ) [사실] 

- 원하면 **폴더 규칙** (예: `YYYY-MM/채널/제목` ) 이나 **파일명 규칙** (예: `YYYYMMDD_제목_ID.txt` )도 바로 맞춰 줄게. “폴더 규칙=원하는 형태”로 말해줘, 형. 

**==> picture [489 x 39] intentionally omitted <==**

위 유튜브List_프롬프트.bat를 실행하기전에 물어보고 폴 더 생성과 이름을 넣을수 있었으면 좋겠어 

15s 동안 생각함 

형, 딱 원하던 대로 **실행 전에 폴더 이름을 물어보고, 그 폴더를 만들어 그 안에 정리** 되게 배치 만들었다. (한글/띄어쓰기 OK) 

## **A단계｜뭐가 되나** 

실행하면 **폴더 이름 먼저 물음** → 폴더 자동 생성 → **그 폴더 안에서** 

- URL 입력(빈 줄로 종료) → 자막 받기 → **TXT만(** `*.clean.txt` **)** 생성. [사실] 

## **B2단계｜배치파일 만들기 (복붙)** 

메모장에 아래 전부 붙여넣고 `유튜브List_프롬프트_폴더.bat` 로 저장(지금 작업 폴더). 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브목록처리(폴더지정) - TXT만생성
```

```
rem 현재배치파일이있는폴더경로(절대경로)
set "ROOT=%~dp0"
```

```
echo 저장할폴더이름을입력하세요. (빈칸이면날짜_시각으로자동생성)
set "FOLDER="
set /p FOLDER=폴더명:
if "%FOLDER%"=="" (
  for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToString(\"yyyy-MM-dd
)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 88/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
rem 폴더생성
set "DEST=%ROOT%%FOLDER%"
if not exist "%DEST%" mkdir "%DEST%"
echo.
echo [안내] URL을한줄씩붙여넣고 Enter 하세요.
echo        아무것도입력하지않고 Enter 하면다운로드를시작합니다.
echo.
rem 작업폴더로이동
pushd "%DEST%"
rem 기존 urls.txt 정리
if exist urls.txt del /f /q urls.txt
:INPUT_LOOP
set "URL="
set /p URL=URL:
if "!URL!"=="" goto START
echo !URL!>>urls.txt
goto INPUT_LOOP
:START
if not exist urls.txt (
  echo [오류] URL이없어작업을종료합니다.
  popd
  pause
  exit /b
)
echo.
echo [1/3] 자막받는중(ko 자동자막)...
set "COOK="
if exist "%ROOT%cookies.txt" (
  copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul
  set "COOK=--cookies cookies.txt"
)
for /f "usebackq delims=" %%U in ("urls.txt") do (
  echo [DL] %%U
  yt-dlp %COOK% --write-auto-subs --sub-langs "ko" --skip-download "%%U"
)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

89/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
echo.
echo [2/3] TXT만생성중...
python "%ROOT%subs_cleaner_plus.py" --no-srt
```

```
echo.
echo [3/3] 완료! 결과는아래폴더에있습니다:
echo   %DEST%
echo   *.clean.txt 파일만사용하시면됩니다.
```

```
popd
pause
```

## **사용법 (아주 간단)** 

**1.** 같은 폴더에 `subs_cleaner_plus.py` (이미 있음) 두고, 방금 만든 `유튜브List_프롬프트_ 폴더.bat` 더블클릭. 

**2. 폴더명 입력** (예: `Cursor_강의_정리` ) → Enter. 

**3. URL을 줄마다** 붙여넣기 → **빈 줄** (그냥 Enter) 입력하면 시작. 

**4.** 끝나면 `작업폴더\폴더명\` 안에 `…clean.txt` 들이 정리되어 있음. 

## **팁** 

- **로그인/연령/지역 제한** 영상이면, 배치와 같은 곳에 `cookies.txt` 만 두면 **자동 적용** 된다. [사실] 

폴더명/URL에 한글·띄어쓰기 OK. 결과는 폴더 안에만 모인다. [사실] 

- 결과를 **채널/날짜/제목으로 더 깊게 분류** 하고 싶으면 말해줘. 이전에 준 `organize_outputs.py` 와 묶어서 “폴더 규칙 버전”으로 합쳐줄게. 

**==> picture [489 x 39] intentionally omitted <==**

특정 유튜브 채널의 url를 복사해서 자동으로 txt를 만들어 줄수 있어? 

가령) https://www.youtube.com/@autocoderIT 오토코더 의 채널의 재생목록중 파이썬의 모든 동영상의 url를 자동 으로 복사해서 텍스트화 시킬수 있는지 궁금해 

25s 동안 생각함 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

90/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

형, **가능해.** 특정 채널의 “파이썬” 재생목록 URL만 알면, 그 목록의 **모든 동영상 URL → TXT(중복 제거본)** 까지 자동으로 뽑을 수 있어. [사실] 

## **A단계｜전체 개요** 

**1. 재생목록 URL** 을 복사(예: 오토코더 채널의 “파이썬” 플레이리스트). 

**2.** 그 재생목록에서 **모든 동영상 URL을 한 번에 urls.txt로 추출** . 

**3. 자막 받기 → TXT만 정리** 를 자동 실행. [사실] 

## **B2단계｜초간단 방법(배치 1개로 끝)** 

## **1) 배치파일 만들기** 

메모장 열고 아래 전부 복사 → 작업 폴더( `subs_cleaner_plus.py` 있는 곳 )에 `재생목록→TXT.bat` 로 저장. 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 재생목록전체를 TXT로 (ko 자동자막 → clean.txt)
```

```
echo [1] 재생목록 URL을붙여넣고 Enter:
set "PL="
set /p PL=Playlist URL:
if "%PL%"=="" ( echo 재생목록 URL이없습니다.& pause & exit /b )
```

```
echo [2] 저장할폴더이름을입력(빈칸=날짜_시각자동):
set "FOLDER="
set /p FOLDER=Folder:
if "%FOLDER%"=="" for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToSt
```

```
rem 폴더준비
set "ROOT=%~dp0"
set "DEST=%ROOT%%FOLDER%"
if not exist "%DEST%" mkdir "%DEST%"
pushd "%DEST%"
```

```
echo.
echo [단계1] 재생목록에서모든영상 URL 추출중...
rem 쿠키자동사용(있으면)
set "COOK="
if exist "%ROOT%cookies.txt" (
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

91/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
  copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul
  set "COOK=--cookies cookies.txt"
)
```

```
rem ★ yt-dlp가재생목록의모든동영상 URL을그대로출력 → urls.txt 저장
rem   --flat-playlist : 목록만빠르게
rem   --print "%(webpage_url)s" : 각항목의웹 URL만프린트
yt-dlp %COOK% --flat-playlist -i --print "%(webpage_url)s" "%PL%" > urls.txt
```

```
if errorlevel 1 (
  echo [오류] URL 추출중문제가발생했습니다. (권한/네트워크/URL 확인)
  popd & pause & exit /b
)
echo.
echo [단계2] 자막받는중(ko 자동자막)...
for /f "usebackq delims=" %%U in ("urls.txt") do (
  echo [DL] %%U
  yt-dlp %COOK% --write-auto-subs --sub-langs "ko" --skip-download "%%U"
)
echo.
echo [단계3] TXT만생성(subs_cleaner_plus.py --no-srt)...
python "%ROOT%subs_cleaner_plus.py" --no-srt
```

```
echo.
echo [완료] 결과: %DEST%
echo        *.clean.txt 파일만사용하시면됩니다.
popd
pause
```

## 설명 

- “ **재생목록 URL** 만 넣으면 `urls.txt` 에 리스트가 생기고, 한글 자동자막으로 받아 서 **TXT만(** `*.clean.txt` **)** 남겨준다. [사실]” 

- “같은 폴더에 `cookies.txt` 가 있으면 **자동 적용** (연령/로그인/지역 제한 우회). [사 실]” 

## **2) 사용법** 

**1.** 유튜브에서 **오토코더(@autocoderIT) 채널 → 재생목록** 으로 들어가서, **“파이썬” 재 생목록을 클릭 → 그 주소(URL)를 복사** . 

**2.** 방금 만든 `재생목록→TXT.bat` **더블클릭** → 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 92/341 