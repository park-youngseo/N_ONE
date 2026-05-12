## **B1｜사용 흐름(딱 3단계)** 

**1.** `Script.bat` **더블클릭** → **정리할 폴더 경로** 물어봄 → 입력(비우면 현재 폴더) 

**2. 메모장(urls.txt)이 자동으로 열림** → **유튜브 링크(영상·재생목록) 붙여넣기 → 저장 (Ctrl+S) → 닫기** 

   - ✅ **저장하고 닫으면 자동으로 실행** 시작됨 (추가 클릭 없음) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 276/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **3.** 자동으로 **자막→TXT 생성 → 전부 합치기 → 4만자 기준 분할** 

결과: `정리폴더\FINAL_TXT\폴더이름_01.txt, _02.txt, ...` 

중간파일은 남기지 않거나(임시) 정리 끝에 지움 → **최종본만 남음** 

## **B2｜파일 내용(복붙 저장)** 

## **1) 버튼:** `Script.bat` 

메모장 열고 아래 전부 붙여넣어 ** `Script.bat` **로 저장. 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal
title Script — 폴더지정→urls.txt 편집→자동실행(자막→TXT→합치기→4만자분할)
```

```
REM 0) Python 확인
where python >nul 2>nul || (
  echo [오류] Python이필요합니다. 설치후다시실행하세요.
  pause & exit /b
)
```

```
REM 1) 엔진확인 (이름고정: Youtube.txt.py)
set "HERE=%~dp0"
```

```
if not exist "%HERE%Youtube.txt.py" (
```

```
  echo [오류] Youtube.txt.py 가이배치와같은폴더에없습니다.
  echo 같은폴더에두고다시실행하세요.
  pause & exit /b
)
```

```
REM 2) 정리할폴더입력(비우면현재폴더)
set "ROOT="
```

```
set /p ROOT=정리(저장)할폴더경로를입력하세요(비우면현재폴더):
if "%ROOT%"=="" set "ROOT=%HERE%"
set "ROOT=%ROOT:"=%"
if not exist "%ROOT%" (
  echo [안내] 폴더가없어새로만듭니다: "%ROOT%"
  md "%ROOT%" >nul 2>nul
)
```

```
REM 3) urls.txt 열어편집(저장후닫으면자동진행)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

277/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
set "URLS=%ROOT%\urls.txt"
```

```
if not exist "%URLS%" (
```

```
  echo [안내] "%URLS%" 파일이없어새로만듭니다.
```

```
  > "%URLS%" echo # 한줄에한개의유튜브링크를적으세요. (재생목록도가능)
  >> "%URLS%" echo # 예시:
  >> "%URLS%" echo # https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX
  >> "%URLS%" echo # https://youtu.be/AAAAAAAAAAA
  >> "%URLS%" echo #
  >> "%URLS%" echo # 저장(Ctrl+S) 후메모장을닫으면자동실행됩니다.
)
```

```
echo.
```

```
echo [열기] 메모장이열리면링크를붙여넣고 **저장(Ctrl+S)** 한뒤창을닫으세요.
start /wait notepad "%URLS%"
```

```
REM 4) 실행(필요도구업데이트는조용히시도)
python -m pip install -U yt-dlp >nul 2>nul
```

```
echo.
```

```
echo [실행] 자막→TXT 생성및정리(4만자분할)를시작합니다...
python "%HERE%Youtube.txt.py" --root "%ROOT%" --urls "%URLS%" --lang ko --chunk-s
```

```
echo.
```

```
echo [완료] 결과: "%ROOT%\FINAL_TXT\폴더이름_01.txt, _02.txt ..." (각 4만자)
pause
```

## **2) 엔진:** `Youtube.txt.py` 

메모장 열고 아래 전부 붙여넣어 ** `Youtube.txt.py` **로 저장. 

## (이 스크립트가 **자막 다운로드→TXT화→합치기→4만자 분할** 까지 한 번에 처리) 

python 코드 복사 

```
# Youtube.txt.py
# 동작:
```

```
# 1) urls.txt의링크(영상/재생목록)를읽어자막(.vtt)을내려받고깨끗한 TXT로변환
```

```
# 2) TXT들을라벨없이쭉이어붙여하나로만들고(메모리)
```

```
# 3) 4만글자씩 '폴더이름_01.txt, _02.txt, ...' 로저장(FINAL_TXT/)
```

```
# 4) 임시파일은정리후삭제 → 최종본만남김
```

```
import argparse, re, html, shutil
from pathlib import Path
from yt_dlp import YoutubeDL
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 278/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
defread_urls(urls_path: Path):
    lines = urls_path.read_text(encoding="utf-8", errors="ignore").splitlines()
return [ln.strip() for ln in lines if ln.strip() andnot ln.strip().startswit
defsafe(s: str) -> str:
return" ".join(re.sub(r'[\\/:*?"<>|]+', '_', s or"").split()).strip()
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
        s = html.unescape(re.sub(r"<.*?>","", s))
if s:
            lines.append(s)
# 인접중복줄제거
    out, prev = [], ""
for s in lines:
if s != prev:
            out.append(s)
        prev = s
return"\n".join(out).strip()
```

```
defdownload_texts(url: str, tmp_dir: Path, lang: str = "ko"):
"""URL(단일/재생목록)에서텍스트목록을순서대로반환"""
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
"quiet": True, "no_warnings": True,
    }
    results = []  # [(vid, title, text)]
with YoutubeDL(ydl_opts) as ydl:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

279/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
        info = ydl.extract_info(url, download=True)
        items = (info.get("entries") ifisinstance(info, dict) elseNone) or [info
for it in items:
ifnot it:
continue
            vid   = safe(it.get("id"))
            title = safe(it.get("title"))
            base  = f"{title} [{vid}]"
# ko 우선, 없으면임의 vtt
            cand = list(tmp_dir.glob(f"{base}.ko.vtt")) \
orlist(tmp_dir.glob(f"{base}.*.vtt")) \
orlist(tmp_dir.glob(f"{base}.vtt"))
ifnot cand:
print(f"[경고] 자막없음 → 건너뜀: {base}")
continue
            text = vtt_to_text(cand[0])
            results.append((vid, title, text))
return results
defsplit_by_chars(text: str, size: int = 40000):
    paras = [p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks, cur = [], ""
for p in paras:
iflen(cur) + len(p) + 2 <= size:
            cur += (("\n\n"if cur else"") + p)
else:
if cur:
                chunks.append(cur); cur=""
whilelen(p) > size:
                chunks.append(p[:size]); p = p[size:]
            cur = p
if cur:
        chunks.append(cur)
return chunks
```

```
defmain():
    ap = argparse.ArgumentParser()
""
    ap.add_argument("--root", required=True, help=결과를저장할상위폴더)
"
    ap.add_argument("--urls", required=True, help="urls.txt 경로)
"
    ap.add_argument("--lang", default="ko", help=자막언어(기본 ko)")
"
    ap.add_argument("--chunk-size", type=int, default=40000, help=분할글자수(기본
    args = ap.parse_args()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 280/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 