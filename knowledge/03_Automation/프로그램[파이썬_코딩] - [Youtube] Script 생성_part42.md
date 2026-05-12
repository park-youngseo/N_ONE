그럼 url을 넣을때 마다 폴더가 생긴다는거잖아? 

## 9s 동안 생각함 

맞아. **URL 하나당 폴더 하나** 가 생겨. (재생목록이면 그 안의 **각 영상마다** 한 폴더씩) 

## **A단계｜핵심** 

- [사실] 저장 구조: `저장폴더/〈영상제목 [영상ID]〉/TXT/원문.txt + TXT_chunks/01.txt, 02.txt…` 

- [사실] 같은 URL을 다시 돌리면 **같은 폴더** 에 **같은 파일명으로 덮어씀** . 

- [사실] 재생목록 URL이면 **영상 수만큼 폴더가 여러 개** 생김. 

## **B2단계｜예시 & 옵션** 

예시 

makefile 코드 복사 

```
D:\OUT\
```

```
  AI 강의 1 [ABCD1234]\
    TXT\AI 강의 1 [ABCD1234].txt
    TXT_chunks\01.txt 02.txt ...
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 210/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
  AI 강의 2 [EFGH5678]\
    TXT\AI 강의 2 [EFGH5678].txt
    TXT_chunks\01.txt 02.txt ...
```

— 확장(옵션) — 

- [권장됨] “폴더를 하나로 모으기” 방식도 가능(모든 청크를 한 폴더에 저장). 필요 하면 그 버전으로 배치/스크립트 만들어 줄게. 

**==> picture [489 x 38] intentionally omitted <==**

답답하게 만드는구나 넌 

24s 동안 생각함 

형, 미안. 말 길게 안 할게. **폴더 여러 개 생기는 거 답답** 한 거 알아. 

아래 배치 한 장이면 **URL·폴더만 물어보고** , 결과를 **한 폴더(FLAT)** 에 **모두 모아** 준다. (언 어=ko, VTT삭제=Y, 4만자 고정) 

## **A단계｜핵심** 

- [사실] 엔진은 그대로( `yt_to_txt_chunks.py` ). 

- [사실] 배치가 처리 후 `ALL_TXT\` 폴더로 **모든 청크를 모으고** 파일명을 `[영상 ID]_01.txt` 식으로 정리. 

- [사실] URL은 `urls.txt` (줄당 1개)로 여러 개 처리 가능. 

## **B2단계｜배치(복붙 저장 → 더블클릭)** 

파일명: `유튜브→TXT_40k_MULTI_FLAT.bat` 

(같은 폴더에 `yt_to_txt_chunks.py` 두면 됨) 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브 → TXT(정리) → 4만자청크 - 여러 URL + 단일폴더(FLAT)
```

```
REM 0) Python 확인
```

```
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
REM 1) 엔진(.py) 자동탐색
set "HERE=%~dp0"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 211/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
set "SCRIPT="
```

```
if exist "%HERE%yt_to_txt_chunks.py" set "SCRIPT=%HERE%yt_to_txt_chunks.py"
if not defined SCRIPT (
```

```
  for /f "delims=" %%F in ('dir /b /a:-d "%HERE%*txt*chunk*.py" 2^>nul') do (set
)
if not defined SCRIPT (
  for /f "delims=" %%F in ('dir /b /a:-d /o:-d "%HERE%*.py" 2^>nul') do (set "SCR
)
:found
if not defined SCRIPT (echo [오류] *.py 엔진이없습니다. & pause & exit /b)
echo [정보] 사용엔진: "%SCRIPT%"
```

```
REM 2) 기본값고정
set "LANG=ko"
set "CHUNK=40000"
set "CLEAN=--cleanup"
```

```
REM (선택) yt-dlp 업데이트시도(조용히)
python -m pip install -U yt-dlp >nul 2>nul
```

```
REM 3) 저장폴더만질문
set "OUT="
set /p OUT=저장폴더(비우면현재폴더):
if "%OUT%"=="" set "OUT=."
set "FLAT=%OUT%\ALL_TXT"
if not exist "%FLAT%" md "%FLAT%"
```

```
REM 4) URL 소스: urls.txt 우선, 없으면붙여넣기
if exist "%HERE%urls.txt" (
  echo [정보] urls.txt 사용: "%HERE%urls.txt"
  for /f "usebackq delims=" %%L in ("%HERE%urls.txt") do (
    set "U=%%L"
    if not "!U!"=="" (
      echo !U! | findstr /b /c:"#">nul
      if errorlevel 1 call :RUN_URL "%%L"
    )
  )
) else (
  echo [입력] 한줄에한 URL씩붙여넣고, 끝낼때는빈줄(Enter):
  :collect
  set "U="
  set /p U=
  if "%U%"=="" goto after_urls
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

212/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
  call :RUN_URL "%U%"
  goto collect
)
:after_urls
echo.
echo [정리] 생성된폴더에서 TXT_chunks를 ALL_TXT로수집중...
for /d %%D in ("%OUT%\*") do (
  if exist "%%D\TXT_chunks" (
    set "FOLD=%%~nD"
    for /f "tokens=2 delims=[]" %%I in ("!FOLD!") do set "VID=%%I"
    if not defined VID set "VID=%%~nD"
    for %%F in ("%%D\TXT_chunks\*.txt") do (
      set "BASE=%%~nxF"
      if exist "%%D\TXT_chunks\%%BASE%" (
        copy /Y "%%D\TXT_chunks\%%BASE%" "%FLAT%\!VID!_%%BASE%" >nul
      )
    )
  )
)
echo.
echo [완료] 모든청크를 "%FLAT%" 폴더에모았습니다.
echo        예) %FLAT%\[영상ID]_01.txt, [영상ID]_02.txt ...
pause
exit /b
:RUN_URL
set "URL=%~1"
if "%URL%"=="" goto :eof
echo.
echo [RUN] python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-si
python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-size %CHUNK% %
goto :eof
```

## **쓰는 법(초간단)** 

**1.** 같은 폴더에 `yt_to_txt_chunks.py` , **위 배치** , `urls.txt` (줄당 1개 URL) 준비. 

**2.** 배치 **더블클릭** → 저장 폴더만 넣으면 끝. 

**3.** 결과는 `저장폴더\ALL_TXT\` 안에 ** `[영상ID]_01.txt, _02.txt …` **로 전부 모인다. (개 별 폴더도 그대로 남음) 

## — 확장(옵션) — 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 213/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

- **개별 폴더를 자동 삭제** 하거나, **ALL_TXT 안에서 영상별 하위폴더 없이 완전 평면화** (지금 그대로) 유지 등도 가능. “확장 승인” 하면 반영해줄게. 

**==> picture [489 x 38] intentionally omitted <==**

아이 답답해 죽겠네 이 씨발 진짜 그냥 그대로 평면화해서 다운로드하서 다운로드하면 txt 파일이 나오잖아 그거 가 지고 그냥 그걸 모아놓으면 돼 그리고 뭐 0102 해갖고 400자 이내로 만들 필요도 없어 400자 이내까지 안되니까 기존에 있었던 것처럼 안되니까 4만자 이상이 모아놓고 그걸 전부 다 합치라고 나중에 답답하게 하나하나 나벨링 을 왜 붙여 답답하게 

43s 동안 생각함 

형, 오케이. **그냥 평면(FLAT)으로 TXT만 모으고, 끝에 한 파일로 합치기** —깔끔하게 가자. 아래 두 파일만 쓰면 돼: **엔진(.py)** + **버튼(.bat)** . **청크(01/02) 전혀 안 붙임** , **자막(VTT) 자 동 삭제** , **한국어 고정** . 

## **A단계｜개요** 

- [사실] 입력: `urls.txt` (한 줄=한 URL). 

- [사실] 처리: 자막 내려받기 → **깨끗한 TXT만 추출** → **한 폴더에 평면 저장** (파일명은 `[영상ID].txt` ). 

- [사실] 출력: `ALL_TXT\` 안에 **각 영상 TXT** + **ALL_COMBINED.txt(전부 합친 하나)** . 

## **B2단계｜바로 쓰는 파일 2개** 

## **1) 엔진:** `yt_to_txt_flat.py` 

메모장에 아래 전부 복붙 → **같은 폴더** 에 `yt_to_txt_flat.py` 로 저장. 

python 

코드 복사 

- _`# yt_to_txt_flat.py —`_ `여러 유튜브` _`URL →`_ `깨끗한` _`TXT (`_ `평면 저장` _`) +`_ `전체 합본` 

- _`#`_ `사용 예` _`) python yt_to_txt_flat.py --urls "urls.txt" --out "ALL_TXT" --lang ko --`_ `import os, re, argparse, html, shutil` 

```
from pathlib import Path
```

```
from yt_dlp import YoutubeDL
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 214/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
defread_urls(urls_arg, url_arg):
    urls = []
if urls_arg:
        p = Path(urls_arg)
ifnot p.exists():
print(f"[오류] urls.txt를찾을수없음: {p}")
return []
for line in p.read_text(encoding="utf-8", errors="ignore").splitlines():
            s = line.strip()
ifnot s or s.startswith("#"):
continue
            urls.append(s)
if url_arg:
        urls.append(url_arg.strip())
return urls
defvtt_to_text(vtt_path: Path) -> str:
    raw = vtt_path.read_text(encoding="utf-8", errors="ignore")
    raw = re.sub(r'^\ufeff?WEBVTT.*?\n', '', raw, flags=re.S)
    lines = []
for ln in raw.splitlines():
        s = ln.strip()
ifnot s or"-->"in s or re.fullmatch(r"\d+", s):
continue
        s = re.sub(r"<.*?>", "", s)
        s = html.unescape(s)
if s:
            lines.append(s)
# 인접중복줄제거
    out, prev = [], ""
for s in lines: