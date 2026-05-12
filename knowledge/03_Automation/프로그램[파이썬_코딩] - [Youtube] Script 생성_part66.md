
윈도우 컴퓨터 

- 인터넷 

- (선택) 유튜브에 로그인된 크롬/엣지 브라우저 

## **1. 파이썬 설치 (한 번만)** 

**1.** https://www.python.org/downloads/ 에서 **Python 3.x** 다운 → 설치할 때 **“Add python.exe to PATH” 체크** → Next → Install. 

**2.** 설치 후 검은창(명령 프롬프트) 열기: 시작 버튼 누르고 `cmd` 입력 → Enter. 

**3.** 아래를 쳐서 버전이 나오면 성공: 

nginx 코드 복사 

```
python -V
```

- (예: Python 3.11.9) 

## **2. 작업 폴더 만들기** 

**1.** 바탕화면에 폴더 하나 만들기: `yotube_script` (이름은 형 마음대로 해도 됨) 

**2.** 이 폴더를 **더블클릭해서 열어둔 상태** 에서, 위쪽 주소창에 `cmd` 입력 → Enter. → 검은창이 **그 폴더에서** 열린다. (중요) 

## **3. 도구 설치 (한 번만)** 

검은창에 아래를 입력: 

nginx 코드 복사 `python -m pip install -U yt-dlp` 

(몇 분 걸릴 수 있음. 끝나면 다음 단계.) 

## **4. 파일 2개 만들기 (이름 고정)** 

## **4-1) Script.bat 만들기** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

326/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **1.** 메모장 열기 → 아래 내용을 **그대로** 복사/붙여넣기 → 

`yotube_script` 폴더에 **이름을 꼭** `Script.bat` **로** 저장 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title Script — 폴더지정→urls.txt 편집→자동실행(자막→TXT→합치기→4만자분할)
```

```
rem 0) Python 확인
where python >nul 2>nul || (
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
```

```
rem 3) urls.txt 열기 (저장후닫으면자동진행)
set "URLS=%ROOT%\urls.txt"
if not exist "%URLS%" (
  > "%URLS%" echo # 한줄에한개의유튜브링크(재생목록가능)
  >> "%URLS%" echo # 예시:
  >> "%URLS%" echo # https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX
  >> "%URLS%" echo # https://youtu.be/AAAAAAAAAAA
  >> "%URLS%" echo # 저장(Ctrl+S) 후창을닫으면계속진행됩니다.
)
echo.
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 327/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
echo [열기] 메모장이열리면링크를붙여넣고저장(Ctrl+S) 후창을닫으세요.
start /wait notepad "%URLS%"
```

```
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
  echo [오류] 실행중문제가발생했습니다. 위메시지마지막줄을복사해서알려주세요.
  pause & exit /b
)
echo.
echo [완료] 결과: "%ROOT%\FINAL_TXT\폴더이름_01.txt, _02.txt ..." (각 4만자)
pause
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 328/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

저장할 때 **파일 확장자 보이기** 가 꺼져 있으면 `Script.bat.txt` 가 될 수 있어요. 탐색기 상단 “보기 → 파일 확장명” **체크** 하고 `Script.bat` 로 저장하세요. 

## **4-2) Youtube.txt.py 만들기** 

**1.** 메모장 새로 열기 → 아래 내용을 **그대로** 복사/붙여넣기 → 같은 폴더에 ** `Youtube.txt.py` **로 저장 

python 

코드 복사 

```
# Youtube.txt.py — 안정화버전 (파일명/흐름그대로)
# 기능: urls.txt(영상/재생목록) → 자막(vtt) → TXT → 모두합치기 → 4만자분할저장
# 안정화: 진행로그표시, URL별재시도, 레이트리밋대비랜덤대기, 부분실패건너뛰기, 임시
```

```
import argparse, re, html, shutil, time, random, sys
from pathlib import Path
```

```
DEFAULT_LANG = "ko"
DEFAULT_CHUNK = 40000
SLEEP_BETWEEN_URLS = (8, 15)      # 각 URL 사이랜덤대기(초) — 레이트리밋회피
SLEEP_BETWEEN_ITEMS = (2, 5)      # 같은재생목록내영상사이잠깐대기
MAX_RETRIES = 3# URL별재시도횟수
```

```
defsafe(s: str) -> str:
return" ".join(re.sub(r'[\\/:*?\"<>|]+', '_', s or"").split()).strip()
```

```
defvtt_to_text(vtt_path: Path) -> str:
    raw = vtt_path.read_text(encoding="utf-8", errors="ignore")
    raw = re.sub(r'^\ufeff?WEBVTT.*?\n', '', raw, flags=re.S)
    lines = []
for ln in raw.splitlines():
        s = ln.strip()
ifnot s or"-->"in s or re.fullmatch(r"\d+", s):
continue
        s = html.unescape(re.sub(r"<.*?>", "", s))
if s:
            lines.append(s)
    out, prev = [], ""
for s in lines:
if s != prev:
            out.append(s)
        prev = s
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 329/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
return"\n".join(out).strip()
```

```
defsplit_by_chars(text: str, size: int):
    paras = [p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks, cur = [], ""
for p in paras:
iflen(cur) + len(p) + 2 <= size:
            cur += (("\n\n"if cur else"") + p)
else:
if cur:
                chunks.append(cur); cur = ""
whilelen(p) > size:
                chunks.append(p[:size]); p = p[size:]
            cur = p
if cur:
        chunks.append(cur)
return chunks
```

```
defread_urls(urls_path: Path):
    data = urls_path.read_text(encoding="utf-8", errors="ignore").splitlines()
return [ln.strip() for ln in data if ln.strip() andnot ln.strip().startswith
```

```
defdownload_once(url: str, tmp_dir: Path, lang: str, cookies_from_browser: str):
"""한번시도: URL(단일/재생목록) → [(vid,title,text), ...] (성공분만)"""
from yt_dlp import YoutubeDL
    tmp_dir.mkdir(parents=True, exist_ok=True)
    ydl_opts = {
"skip_download": True,
"writesubtitles": True,
"writeautomaticsub": True,
"subtitleslangs": [lang],
"subtitlesformat": "vtt",
"outtmpl": str(tmp_dir / "%(title)s [%(id)s].%(ext)s"),
"paths": {"home": str(tmp_dir)},
"noplaylist": False,
"quiet": False,              # 진행로그보이게 (중요)
"no_warnings": True,
    }
if cookies_from_browser:
        ydl_opts["cookiesfrombrowser"] = (cookies_from_browser,)
```

```
    results = []  # [(vid,title,text)]
with YoutubeDL(ydl_opts) as ydl:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 330/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
        info = ydl.extract_info(url, download=True)
        items = (info.get("entries") ifisinstance(info, dict) elseNone) or [info
for it in items:
ifnot it:
continue
            vid, title = safe(it.get("id")), safe(it.get("title"))
            base = f"{title} [{vid}]"
            cand = list(tmp_dir.glob(f"{base}.ko.vtt")) \
orlist(tmp_dir.glob(f"{base}.*.vtt")) \
orlist(tmp_dir.glob(f"{base}.vtt"))
ifnot cand:
print(f"[경고] 자막없음 → 건너뜀: {base}")