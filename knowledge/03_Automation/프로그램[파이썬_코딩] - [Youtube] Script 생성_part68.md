:_ok
if not defined HASURL (
  echo [오류] urls.txt에처리할링크가없습니다. (#로시작하는줄은무시)
  pause & exit /b
)
rem 5) yt-dlp 업데이트(조용히시도)
python -m pip install -U yt-dlp >nul 2>nul
rem 6) 실행 (언어 ko, 분할 40000 고정)
echo.
echo [실행] 영상별 TXT 생성(4만자분할) 시작...
python "Youtube_solo.txt.py" --root "%ROOT%" --urls "%URLS%" --lang ko --chunk-si
if errorlevel 1 (
  echo.
  echo [오류] 실행중문제가발생했습니다. 위메시지마지막줄을복사해서알려주세요.
  pause & exit /b
)
echo.
echo [완료] 결과: "%ROOT%\TXT_BY_URL\제목.txt" 또는 "제목_01.txt, _02.txt …"
pause
```

**주의** : “파일 확장명 보이기”가 꺼져 있으면 `Script_solo.bat.txt` 로 저장될 수 있어. 탐색기 상단 “보기 → 파일 확장명” 체크 후 `.bat` 로 저장해. 

## **2) Youtube_solo.txt.py (그대로 저장)** 

메모장 새로 열기 → 아래 코드 복사 → ** `Youtube_solo.txt.py` **로 저장(배치와 같은 폴 더). 

python 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

336/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# Youtube_solo.txt.py — 영상별 TXT 생성(4만자분할)
# urls.txt(영상또는재생목록) → 각영상자막(VTT) 추출 → 깨끗한 TXT 저장
# 한파일이 4만자초과면제목_01.txt, _02.txt ...로분할
# 결과: <ROOT>\TXT_BY_URL\
import argparse, re, html, shutil, time, random, sys
from pathlib import Path
DEFAULT_LANG = "ko"
DEFAULT_CHUNK = 40000
SLEEP_BETWEEN_URLS = (8, 15)   # URL 사이대기(레이트리밋회피)
SLEEP_BETWEEN_ITEMS = (2, 5)   # 같은재생목록내영상사이대기
MAX_RETRIES = 3# URL별재시도
defsafe(s: str) -> str:
return" ".join(re.sub(r'[\\/:*?\"<>|]+', '_', s or"").split()).strip()
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
return"\n".join(out).strip()
defsplit_by_chars(text: str, size: int):
    paras = [p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks, cur = [], ""
for p in paras:
iflen(cur) + len(p) + 2 <= size:
            cur += (("\n\n"if cur else"") + p)
else:
if cur:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 337/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
                chunks.append(cur); cur = ""
whilelen(p) > size:
```

```
            cur = p
if cur:
        chunks.append(cur)
return chunks
```

```
defread_urls(urls_path: Path):
```

```
    data = urls_path.read_text(encoding="utf-8", errors="ignore").splitlines()
return [ln.strip() for ln in data if ln.strip() andnot ln.strip().startswith
```

```
defextract_texts(url: str, tmp_dir: Path, out_dir: Path, lang: str):
"""URL(단일/재생목록) → 각영상의 TXT 파일생성"""
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
"quiet": False,
"no_warnings": True,
    }
    results = 0
with YoutubeDL(ydl_opts) as ydl:
```

```
        info = ydl.extract_info(url, download=True)
        items = (info.get("entries") ifisinstance(info, dict) elseNone) or [info
for it in items:
ifnot it:
continue
            vid, title = safe(it.get("id")), safe(it.get("title"))
            base = f"{title} [{vid}]"
# ko 우선, 없으면다른 vtt라도
            cand = list(tmp_dir.glob(f"{base}.ko.vtt")) \
orlist(tmp_dir.glob(f"{base}.*.vtt")) \
orlist(tmp_dir.glob(f"{base}.vtt"))
ifnot cand:
print(f"[경고] 자막없음 → 건너뜀: {base}")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 338/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
continue
            txt = vtt_to_text(cand[0])
ifnot txt:
print(f"[경고] 추출텍스트없음 → 건너뜀: {base}")
continue
# 4만자분할저장
            chunks = split_by_chars(txt, size=args.chunk_size)
iflen(chunks) == 1:
                (out_dir / f"{base}.txt").write_text(chunks[0], encoding="utf-8")
print(f"[저장] {base}.txt")
else:
for i, c inenumerate(chunks, 1):
(outdir/f"{base}{i:02d}.txt").writetext(c,encoding="utf
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

339/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

340/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

341/341 
