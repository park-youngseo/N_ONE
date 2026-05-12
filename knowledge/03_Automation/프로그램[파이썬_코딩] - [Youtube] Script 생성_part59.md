**2. 비공개도 받아야 하면:** 위 수정 적용 후 

   - `Script.bat` 실행 → “쿠키 쓸까요?”에 **Y** → 브라우저 입력( **chrome** 또는 **edge** ) 

   - → 자동 실행. 

완료되면 결과는 그대로: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 290/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

코드 복사 

 

```
…\FINAL_TXT\
폴더이름_01.txt
폴더이름_02.txt
  …
```

검은창에 다른 에러가 뜨면 **마지막 3~4줄** 만 붙여줘. 그 지점부터 바로 잡을게. 

**==> picture [489 x 40] intentionally omitted <==**

초등학생도 할수 있게 쉽게 txt.py 를 수정해서 할수 있도 록 만들어줘 

6s 동안 생각함 

형, 알겠어. **완전 초간단 한 파일짜리** 로 고쳤어. 

이 파일 하나만 있으면, **물어보는 대로 엔터 → 링크 붙여넣고 저장** 하면 끝나. 

## **A단계｜쓰는 법(초등학생 버전)** 

**1.** 폴더에 `Youtube.txt.py` 하나 둔다. 

**2.** 검은창에서 그 폴더로 이동 → 아래 한 줄 실행: 

nginx 코드 복사 

```
python Youtube.txt.py
```

## **3.** 화면에 나오는 순서대로: 

   - **정리할 폴더 경로** 물으면 입력(그냥 엔터=현재 폴더) 

   - 자동으로 열리는 **urls.txt** 메모장에 **유튜브 링크(영상/재생목록)** 붙여넣기 → **Ctrl+S(저장)** → **창 닫기** 

   - **“로그인 필요한 영상 있나요? (Y/N)”** 나오면 보통 **N** , 

   - 비공개/연령제한 있으면 **Y** 입력 후 **브라우저 이름(chrome/edge 등)** 적기 

**4.** 기다리면 결과는 여기: 

php-template 

코드 복사 

```
<정리폴더>\FINAL_TXT\<폴더이름>_01.txt, _02.txt, ...
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 291/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

(각 40,000자. 중간파일은 지움) 

처음 실행 전에 한 번: 

nginx 

코드 복사 

```
python -m pip install -U yt-dlp
```

## **B단계｜한 파일짜리 완성 코드 (그대로 저장해서 사용)** 

파일명: `Youtube.txt.py` 

python 코드 복사 

```
# Youtube.txt.py  —  초등학생버전 (한파일로끝!)
```

```
# 흐름: 폴더물어보기 → urls.txt 자동열기(붙여넣기·저장) → (선택)쿠키사용 → 자막→TXT
# 결과: <정리폴더>\FINAL_TXT\<폴더이름>_01.txt, _02.txt, ...
```

```
import argparse, re, html, shutil, subprocess, sys
from pathlib import Path
```

```
# ---- 기본설정(그대로둬도됨) ----
DEFAULT_LANG = "ko"
DEFAULT_CHUNK = 40000
```

```
defsafe(s: str) -> str:
"""파일명안전하게"""
return" ".join(re.sub(r'[\\/:*?\"<>|]+', '_', s or"").split()).strip()
```

```
defvtt_to_text(vtt_path: Path) -> str:
"""VTT 자막을깨끗한텍스트로변환"""
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
# 인접중복줄제거
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 292/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    out, prev = [], ""
for s in lines:
if s != prev:
            out.append(s)
        prev = s
return"\n".join(out).strip()
```

```
defsplit_by_chars(text: str, size: int = DEFAULT_CHUNK):
"""빈줄(문단) 기준우선분할, 너무크면강제컷"""
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
defensure_yt_dlp():
"""yt-dlp 없는경우안내"""
try:
import yt_dlp  # noqa
returnTrue
except Exception:
print("[오류] yt-dlp가없습니다. 아래명령을먼저실행하세요:")
print("       python -m pip install -U yt-dlp")
returnFalse
```

```
defread_urls(urls_path: Path):
"""urls.txt 읽기 (주석 # 무시)"""
    data = urls_path.read_text(encoding="utf-8", errors="ignore").splitlines()
return [ln.strip() for ln in data if ln.strip() andnot ln.strip().startswith
```

```
defedit_urls_file(urls_path: Path):
```

```
"""메모장으로 urls.txt 열고저장하게대기"""
ifnot urls_path.exists():
        urls_path.write_text(
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 293/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
"# 한줄에한개의유튜브링크(재생목록가능)\n"
```

```
"# 예시:\n"
"# https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX\n"
"# https://youtu.be/AAAAAAAAAAA\n"
"# 저장(Ctrl+S) 후창을닫으면계속진행됩니다.\n",
            encoding="utf-8"
        )
print("\n[열기] 메모장이열리면링크를붙여넣고저장(Ctrl+S) 후창을닫으세요.")
    subprocess.run(["notepad", str(urls_path)], check=False)
```

```
defask_yesno(msg: str, default="N"):
"""Y/N 입력받기"""
    s = input(f"{msg} (Y/N, 기본{default}): ").strip().lower()
ifnot s:
        s = default.lower()
return s.startswith("y")
```

```
defdownload_one_url(url: str, tmp_dir: Path, lang: str, cookies_from_browser: st
"""단일 URL(영상/재생목록)에서자막→텍스트변환하여리스트반환"""
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
"quiet": True,
"no_warnings": True,
    }
if cookies_from_browser:
# chrome / edge / firefox / brave 중하나
        ydl_opts["cookiesfrombrowser"] = (cookies_from_browser,)
```

```
    results = []  # [(vid, title, text)]
```

```
try:
with YoutubeDL(ydl_opts) as ydl:
            info = ydl.extract_info(url, download=True)
            items = (info.get("entries") ifisinstance(info, dict) elseNone) or
for it in items:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 294/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
ifnot it:
```

```
continue
                vid = safe(it.get("id"))
                title = safe(it.get("title"))
                base = f"{title} [{vid}]"
# ko 우선, 없으면다른 vtt라도
                cand = list(tmp_dir.glob(f"{base}.ko.vtt")) \
orlist(tmp_dir.glob(f"{base}.*.vtt")) \
orlist(tmp_dir.glob(f"{base}.vtt"))
ifnot cand:
print(f"[경고] 자막없음 → 건너뜀: {base}")
continue
                text = vtt_to_text(cand[0])
if text:
                    results.append((vid, title, text))
except Exception as e:
print(f"[오류] 처리실패(건너뜀): {url}\n       {e}")
return results
```
