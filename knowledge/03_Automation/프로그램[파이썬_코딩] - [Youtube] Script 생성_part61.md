
```
# 1) urls.txt의링크(영상/재생목록)를순서대로처리
# 2) 자막(.vtt) → 깨끗한 TXT로변환
# 3) 모두이어붙여 4만자씩분할저장(FINAL_TXT\폴더이름_01.txt …)
# 4) 임시파일은정리후삭제
import argparse, re, html, shutil
from pathlib import Path
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 299/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
DEFAULT_LANG = "ko"
```

```
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
        s = html.unescape(re.sub(r"<.*?>","", s))
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

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 300/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
defdownload_texts_from_url(url: str, tmp_dir: Path, lang: str, cookies_from_brow
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
"quiet": True, "no_warnings": True,
    }
if cookies_from_browser:
        ydl_opts["cookiesfrombrowser"] = (cookies_from_browser,)
```

```
    results = []  # [(vid,title,text)]
with YoutubeDL(ydl_opts) as ydl:
        info = ydl.extract_info(url, download=True)
        items = (info.get("entries") ifisinstance(info, dict) elseNone) or [info
for it in items:
```

```
ifnot it: continue
"id")), safe(it.get("title"))
f"{title} [{vid}]"
list(tmp_dir.glob(f"{base}.ko.vtt")) \
orlist(tmp_dir.glob(f"{base}.*.vtt")) \
```

```
orlist(tmp_dir.glob(f"{base}.vtt"))
ifnot cand:
print(f"[경고] 자막없음 → 건너뜀: {base}")
continue
            txt = vtt_to_text(cand[0])
if txt: results.append((vid, title, txt))
return results
```

```
defmain():
```

```
    ap = argparse.ArgumentParser()
```

```
""
    ap.add_argument("--root", required=True, help=결과저장상위폴더)
"
    ap.add_argument("--urls", required=True, help="urls.txt 경로)
"
    ap.add_argument("--lang", default=DEFAULT_LANG, help=자막언어(기본 ko)")
""
    ap.add_argument("--chunk-size", type=int, default=40000, help=분할글자수)
""
    ap.add_argument("--cookies-from-browser", default=, help="chrome/edge/firefo
    args = ap.parse_args()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 301/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    root = Path(args.root).expanduser().resolve()
    urls_file = Path(args.urls).expanduser().resolve()
assert root.exists(), f"[오류] 폴더가없습니다: {root}"
assert urls_file.exists(), f"[오류] urls.txt가없습니다: {urls_file}"
```

```
    urls = read_urls(urls_file)
ifnot urls:
print("[오류] urls.txt가비어있습니다. (#로시작하는줄은무시)")
return
    tmp_dir = root / "__vtt_tmp"
    all_parts = []
try:
for u in urls:
print(f"[처리] {u}")
            items = download_texts_from_url(u, tmp_dir, args.lang, args.cookies_f
for _, _, t in items:
                all_parts.append(t); all_parts.append("\n\n")
```

```
        combined = "".join(all_parts).strip()
ifnot combined:
print("[오류] 합칠텍스트가없습니다. (자막없음/실패)")
return
        out_dir = root / "FINAL_TXT"
        out_dir.mkdir(parents=True, exist_ok=True)
        base = safe(root.name) or"OUTPUT"
```

```
for i, c inenumerate(split_by_chars(combined, args.chunk_size), 1):
            (out_dir / f"{base}_{i:02d}.txt").write_text(c, encoding="utf-8")
```

```
print(f"[완료] {out_dir}\\{base}_01.txt ... (각{args.chunk_size}자)")
finally:
```

```
if tmp_dir.exists():
            shutil.rmtree(tmp_dir, ignore_errors=True)
```

```
if __name__ == "__main__":
    main()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

302/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **3) 쓰는 법(딱 4줄)** 

**1.** 두 파일 **같은 폴더** 에 둔다: `Script.bat` , `Youtube.txt.py` 

**2.** `Script.bat` **더블클릭** 

**3.** 폴더 경로 입력 → **열린 메모장에 링크 붙여넣고 저장(Ctrl+S) 후 닫기** 

**4.** (비공개 있으면 Y → 브라우저 `chrome` 등 입력) → 끝나면 `FINAL_TXT\폴더이름 _01.txt, _02.txt …` 

## **막히면 여기만 봐** 

- `Private video` 에러 → 다음 실행 때 **Y** 선택 후 **chrome/edge** 입력. 

- `No module named yt_dlp` → `python -m pip install -U yt-dlp` 한 번 실행. 

- 결과가 없음 → `urls.txt` 에서 `#` 지우고 **링크만** 넣었는지 확인(저장 필수). 

대가리 쳐벌라 니 멋대로 바꿔서 입력하고 지랄인데 

8s 동안 생각함 

형, 미안. 이제 **아무것도 안 바꿔** . 딱 이대로만 해. 

## **파일 이름 (고정)** 

- 엔진: `Youtube.txt.py` 

- 버튼: `Script.bat` 

## **정말 그대로 따라 하기 (6단계)** 

**1.** 두 파일을 **같은 폴더** (예: `C:\Users\엔오원\Desktop\yotube_script` )에 둔다. 

**2.** 그 폴더 창 주소줄에 `cmd` 입력 → Enter (검은창이 그 폴더에서 열림). 

**3.** (최초 1회) 도구 설치: 

nginx 코드 복사 

```
python -m pip install -U yt-dlp
```

**4.** `Script.bat` 더블클릭. 

   - “정리할 폴더 경로”가 뜨면 **그 폴더 경로** 붙여넣고 Enter (혹은 그냥 Enter = 현재 폴더) 

**5.** 메모장( `urls.txt` )이 뜸 → **링크를 한 줄에 하나씩** 붙여넣기 → **Ctrl+S(저장)** → **창 닫 기** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

303/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

주의: 줄 맨 앞에 `#` 있으면 **무시됨** . 반드시 **링크만** . 

## **6.** (물으면) **비공개/연령제한 있나요?** 

없으면 **N** . 

있으면 **Y** → 브라우저 입력: `chrome` (또는 `edge` ) → 자동 실행. 

끝나면 여기 생성됨: 

makefile 

코드 복사 

```
C:\Users\엔오원\Desktop\yotube_script\FINAL_TXT\
  yotube_script_01.txt
  yotube_script_02.txt
  ...
```

(각 4만자, 중간 파일은 정리됨) 

## **막히는 포인트 3개 (바로 해결)** 

`urls.txt 비어있음` → 링크 앞의 `#` 지우고 **저장 후 닫기** 다시 실행. 

`Private video…` → 6단계에서 **Y → chrome** (또는 edge). 

`No module named yt_dlp` → 3단계 명령 다시 실행. 