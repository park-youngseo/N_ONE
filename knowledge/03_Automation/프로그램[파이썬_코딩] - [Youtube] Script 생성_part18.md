
25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

- ① **재생목록 URL 붙여넣기 → Enter** 

- ② **폴더명 입력** (예: `오토코더_파이썬` ) → Enter 

- → 폴더 안에 ** `…clean.txt` **들이 생성된다. [사실] 

## **추가 팁** 

- **다른 채널/재생목록도 동일** : 재생목록 URL만 바꿔 실행하면 끝. [사실] 

- 영어 자막이 필요하면 `--sub-langs "ko,en"` 로 바꿀 수 있어. (429 뜰 수 있음) [사실] 

- 재생목록이 매우 크면(수백~수천) **중간에 429/끊김** 이 날 수 있다 → 몇 번 더 실행하 면 이어서 채워짐. [권장됨] 

- `subs_cleaner_plus.py` 가 만든 파일명은 **원본 자막 파일명을 기준** 이라, 필요하면 **정** 

- **리 스크립트(organize_outputs.py)** 로 **채널/날짜/제목 폴더** 로도 자동 이동 시켜줄 수 있어. 원하면 합쳐서 “원클릭 전체 정리판”으로 만들어줄게. [추론] 

원하는 **오토코더 “파이썬” 재생목록 URL** 을 바로 붙여 실행해보고, 폴더에 생성된 

`urls.txt` **/** `*.clean.txt` 상황을 알려줘. 거기 맞춰 다음 커스터마이즈(파일명 규칙/주제 분류 등)도 바로 붙여줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

[단계1] 재생목록에서 모든 영상 URL 추출 중... '만' is not recognized as an internal or external command, 

operable program or batch file. 

The filename, directory name, or volume label syntax is incorrect. 

Usage: yt-dlp [OPTIONS] URL [URL...] 

yt-dlp: error: You must provide at least one URL. Type yt-dlp --help to see a list of all options. [오류] URL 추출 중 문제가 발생했습니다. (권한/네트워 크/URL 확인 

Press any key to continue . . . 

형, 오류 원인 딱 두 가지였어. 

**1. 배치에서** `%` **는 특수문자** 라 `%(webpage_url)s` 를 `%%(webpage_url)s` 로 써야 해. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

93/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**2.** 배치 주석(한글/특수기호) 줄이 섞여서 CMD가 **‘만’ 같은 토큰을 명령으로 오해** 했어. 

아래 **“안전판(주석·특수문자 없음)” 배치** 로 갈아끼우면 깔끔히 된다. 

## **A단계｜핵심** 

- `--print "%(webpage_url)s"` → **배치에서는** `--print "%%(webpage_url)s"` 로 써야 한 

- 다. [사실] 

## **B2단계｜새 배치로 교체 (복붙)** 

메모장에 아래를 **그대로** 붙여넣고, 지금 폴더에 

`재생목록→TXT_SAFE.bat` 라는 이름으로 저장: 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title Playlist -> TXT (safe)
```

```
set "ROOT=%~dp0"
echo Paste PLAYLIST URL (no <>). Press Enter:
set "PL="
set /p PL=URL:
if "%PL%"=="" ( echo [ERR] No URL & pause & exit /b )
echo Folder name (blank = timestamp):
set "FOLDER="
set /p FOLDER=Folder:
if "%FOLDER%"=="" for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToSt
set "DEST=%ROOT%%FOLDER%"
if not exist "%DEST%" mkdir "%DEST%"
pushd "%DEST%"
if exist urls.txt del /f /q urls.txt
```

```
set "COOK="
if exist "%ROOT%cookies.txt" (
  copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul
  set "COOK=--cookies cookies.txt"
)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 94/341 

25. 11. 30. 오후 1:07 

 

**==> picture [154 x 9] intentionally omitted <==**

```
:: 1) extract ALL video URLs in the playlist -> urls.txt
```

```
yt-dlp %COOK% --flat-playlist -i --print "%%(webpage_url)s" "%PL%" > urls.txt
if errorlevel 1 (
```

```
  echo [ERR] Failed to extract URLs. Check the playlist URL / network / cookies.
```

```
  popd & pause & exit /b
)
```

```
:: 2) download auto-subs (ko) for each URL (no media)
for /f "usebackq delims=" %%U in ("urls.txt") do (
  echo [DL] %%U
```

```
  yt-dlp %COOK% --write-auto-subs --sub-langs "ko" --skip-download "%%U"
)
```

```
:: 3) make TXT only (clean)
python "%ROOT%subs_cleaner_plus.py" --no-srt
```

```
echo.
```

```
echo [DONE] TXT files created here:
echo %DEST%
popd
pause
```

## **사용법** 

**1.** `재생목록→TXT_SAFE.bat` **더블클릭** 

**2. 재생목록 URL** 그대로 붙여넣기(꺾쇠 `< >` 금지) → Enter 

**3.** 폴더명 입력(빈칸 가능) → Enter 

**4.** 끝나면 그 폴더에 `*.clean.txt` 가 생겨 있어. 

## **빠른 점검** 

URL 붙여넣을 때 `<https://...>` **말고** `"https://..."` **또는 그냥 https…** 로. 

- 에러가 나면, 창에 나온 **마지막 10줄** 만 복사해서 보내줘. 거기서 바로 잡아줄게. 

실행이 안되는데? 

22s 동안 생각함 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

95/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

형, 오케이. **지금 당장 돌아가게** 제일 안전한 방법으로 잡아줄게. (PARTICIPATION=Off / EVALUATION=Off) 

## **A단계｜핵심** 

- 배치 안 돌아가면 보통 **yt-dlp 인식/따옴표/퍼센트(%) 처리** 문제다. 

- **“파워-세이프 배치”**로 교체하면 해결: `yt-dlp` 가 없어도 ** `python -m yt_dlp` **로 자동 대체, 퍼센트/따옴표도 안전 처리. [사실] 

## **B2단계｜교체 배치(복붙 → 더블클릭)** 

**1.** 메모장 열고 아래 전부 복붙 → 작업 폴더( `subs_cleaner_plus.py` 있는 곳 )에 `재생목록→TXT_SAFE2.bat` 로 저장. 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title Playlist -> TXT (ultra safe)
```

```
:: 0) yt-dlp 실행기탐지 (없으면 python -m yt_dlp 사용)
set "YTDLP=yt-dlp"
```

```
where yt-dlp >nul 2>nul || set "YTDLP=python -m yt_dlp"
```

```
:: 1) 재생목록 URL 입력
```

```
echo 재생목록(Playlist) URL을붙여넣고 Enter (채널 URL 말고, '재생목록' 주소):
set "PL="
```

```
set /p PL=URL:
```

```
if "%PL%"=="" ( echo [ERR] URL 없음 & pause & exit /b )
```

```
:: 2) 저장폴더명
echo 저장할폴더명(빈칸=자동날짜_시각):
set "FOLDER="
set /p FOLDER=Folder:
```

```
if "%FOLDER%"=="" for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToSt
```

```
set "ROOT=%~dp0"
```

```
set "DEST=%ROOT%%FOLDER%"
if not exist "%DEST%" mkdir "%DEST%"
pushd "%DEST%"
```

```
if exist urls.txt del /f /q urls.txt
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

96/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

```
:: 3) 쿠키자동(있으면사용)
set "COOK="
if exist "%ROOT%cookies.txt" (
  copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul
  set "COOK=--cookies cookies.txt"
)
echo.
echo [1/3] 재생목록 URL 추출중...
"%YTDLP%" %COOK% --flat-playlist -i --print "%%(webpage_url)s" "%PL%" > urls.txt
if %ERRORLEVEL% NEQ 0 (
  echo [ERR] URL 추출실패. 재생목록 URL이맞는지확인하세요.
  echo 예) https://www.youtube.com/playlist?list=PLxxxxxxxxxxxx
  popd & pause & exit /b
)
for %%A in (urls.txt) do if %%~zA lss 5 (
  echo [ERR] urls.txt가비어있습니다. 재생목록 URL을다시확인하세요.
  popd & pause & exit /b
)
echo.
echo [2/3] 자막(ko) 받는중...
for /f "usebackq delims=" %%U in ("urls.txt") do (
  echo [DL] %%U
  "%YTDLP%" %COOK% --write-auto-subs --sub-langs "ko" --skip-download "%%U"
)
echo.
echo [3/3] TXT만생성...