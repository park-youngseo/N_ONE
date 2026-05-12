if s != prev:
            out.append(s)
        prev = s
return"\n".join(out).strip()
```

```
defprocess_one(url, out_dir: Path, lang: str, cookies: Path|None, tmp_dir: Path,
    out_dir.mkdir(parents=True, exist_ok=True)
    tmp_dir.mkdir(parents=True, exist_ok=True)
```

```
    ydl_opts = {
"skip_download": True,
"writesubtitles": True,
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 215/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
"writeautomaticsub": True,
"subtitleslangs": [lang],
"subtitlesformat": "vtt",
"outtmpl": str(tmp_dir / "%(title)s [%(id)s].%(ext)s"),
"paths": {"home": str(tmp_dir)},
"noplaylist": False,
"quiet": True,
"no_warnings": True,
    }
if cookies and cookies.exists():
        ydl_opts["cookiefile"] = str(cookies)
    texts = []  # [(id, title, text)]
with YoutubeDL(ydl_opts) as ydl:
        info = ydl.extract_info(url, download=True)
        entries = info.get("entries") ifisinstance(info, dict) elseNone
        items = entries or [info]
for it in items:
ifnot it:
continue
            vid   = (it.get("id") or"").strip()
            title = (it.get("title") or"").strip()
            base  = f"{title} [{vid}]"
# ko 우선, 없으면아무 vtt
            cand_ko  = list(tmp_dir.glob(f"{base}.ko.vtt"))
            cand_any = list(tmp_dir.glob(f"{base}.*.vtt")) orlist(tmp_dir.glob(f
            use = cand_ko[0] if cand_ko else (cand_any[0] if cand_any elseNone)
ifnot use:
print(f"[경고] 자막없음 → 건너뜀: {title} [{vid}]")
continue
try:
                text = vtt_to_text(use)
except Exception as e:
print(f"[경고] VTT→TXT 변환실패: {use.name} ({e})")
continue
# 평면저장: [영상ID].txt (제목길이문제/경로길이회피)
            out_path = out_dir / f"{vid}.txt"
            out_path.write_text(text, encoding="utf-8")
            texts.append((vid, title, text))
if cleanup:
try: use.unlink()
except: pass
return texts
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 216/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
defmain():
```

```
    ap = argparse.ArgumentParser()
"""
    ap.add_argument("--urls", default=, help=줄당 1개 URL이들어있는텍스트파일
"""
    ap.add_argument("--url",  default=, help=단일 URL")
"
    ap.add_argument("--out",  default="ALL_TXT", help=출력폴더 (기본: ALL_TXT)")
"
    ap.add_argument("--lang", default="ko", help=자막언어 (기본: ko)")
""
    ap.add_argument("--cleanup", action="store_true", help=처리후 .vtt 파일삭제
"
    ap.add_argument("--combine", action="store_true", help=모든 TXT를하나의 ALL_C
""
    ap.add_argument("--cookies", default=, help="cookies.txt 경로 (선택)")
    args = ap.parse_args()
```

```
    urls = read_urls(args.urls, args.url)
ifnot urls:
print("[오류] 처리할 URL이없습니다. --urls 또는 --url 사용")
return
```

```
    out_dir = Path(args.out).expanduser().resolve()
    tmp_dir = out_dir / "__vtt_tmp"
    cookies = Path(args.cookies) if args.cookies elseNone
```

```
    combined_order = []  # [(vid, title)]
for u in urls:
print(f"[처리] {u}")
        texts = process_one(u, out_dir, args.lang, cookies, tmp_dir, args.cleanup
        combined_order.extend([(vid, title) for vid, title, _ in texts])
```

```
# 임시폴더정리
```

```
if args.cleanup and tmp_dir.exists():
        shutil.rmtree(tmp_dir, ignore_errors=True)
```

```
if args.combine:
```

```
        combo = out_dir / "ALL_COMBINED.txt"
withopen(combo, "w", encoding="utf-8") as f:
for vid, title in combined_order:
                p = out_dir / f"{vid}.txt"
if p.exists():
                    f.write(f"\n\n===== {title} [{vid}] =====\n\n")
                    f.write(p.read_text(encoding="utf-8", errors="ignore"))
print(f"[완료] 합본생성: {combo}")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

217/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
if __name__ == "__main__":
    main()
```

## **2) 버튼:** `유튜브→TXT_FLAT_AUTO.bat` 

같은 폴더에 아래 저장(더블클릭): 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal
title 유튜브 → TXT 평면저장 + 합본 (질문없음)
```

```
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
python -m pip install -U yt-dlp >nul 2>nul
```

```
set "HERE=%~dp0"
```

```
if not exist "%HERE%yt_to_txt_flat.py" echo [오류] yt_to_txt_flat.py가필요합니다.
if not exist "%HERE%urls.txt" echo [오류] urls.txt(줄당 1개 URL)가필요합니다. & ex
```

```
REM 출력폴더: 현재폴더\ALL_TXT (고정)
set "OUT=%HERE%ALL_TXT"
```

```
echo [RUN] python "yt_to_txt_flat.py" --urls "%HERE%urls.txt" --out "%OUT%" --lang
python "yt_to_txt_flat.py" --urls "%HERE%urls.txt" --out "%OUT%" --lang ko --clea
```

```
echo.
```

```
echo [완료] 결과폴더: "%OUT%"
echo         - 각영상 TXT: [영상ID].txt
echo         - 전체합본   : ALL_COMBINED.txt
pause
```

  

## **쓰는 법 (초간단)** 

**1. 같은 폴더** 에 아래 3개를 둔다: 

`yt_to_txt_flat.py` , `유튜브→TXT_FLAT_AUTO.bat` , `urls.txt` 

**2.** `urls.txt` 에 **한 줄=한 링크** 적기(재생목록도 OK). 

**3. 배치 더블클릭** → 끝. 

`ALL_TXT\` 안에 `[영상ID].txt` 들이 **평면으로** 쌓임. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

218/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## `ALL_COMBINED.txt` 에 **전부 합쳐서** 들어감. 

**청크/번호 라벨 없음** , **한국어 고정** , **VTT는 자동 삭제** . 

원하면 합본에서 머릿글( `===== 제목 [ID] =====` )도 빼줄 수 있고, 파일명을 **제목.txt** 로 바 꾸는 버전도 만들어줄 수 있어. 

**==> picture [489 x 38] intentionally omitted <==**

아직 실행해서 뭔가 나오지마 만들지마 일단은 제일 중요 한건 알고 움직여서 만들어라 파일을 그냥 만들어서 넣지 말고 헷갈리니까 기존에 있던거 다 그대로 진행을 하고 그 부분 영상 URL을 집어 넣고 여러 URL을 집어 넣잖아 그러 면 나오는 텍스트 파일들을 그냥 한 곳에 모아 모아서 그 걸 전부 다 하나로 합쳐 그리고 삼환자로 분류해서 나벨링 01, 02, 03 삼환자로 기준으로 분류를 해서 모든걸 합친거 를 그렇게 해서 01, 02, 03 순서대로 만들으라고 그걸 얘기 하는 건데 뭘 자꾸 이래저래 이걸 하고 저걸 하고 하면서 이걸 뜯어 고치고 저걸 뜯어 고치고 지금 몇 번째 얘기하 는 거야 

38s 동안 생각함 

형, 오케이. **지금 방식 그대로 유지** 하고, 끝에 **모으기 → 하나로 합치기 → 4만자 단위로 01/02/03** 만 해 주는 “후처리”만 붙이면 된다. 

여기선 **아무것도 실행 안 함** . 파일만 준비해 두고, 형이 원할 때 돌리면 돼. [사실] 

## **A단계｜핵심 개요** 

- [사실] 기존처럼 URL 넣어서 **영상별 TXT** 가 생기게 먼저 돌린다. 

- [사실] 아래 **후처리 스크립트 1개** 로: 

- **(1) TXT를 한 폴더로 모으고 → (2) 전부 합쳐 하나의 TXT → (3) 그 파일을 4만자씩 01/02/03…으로 분할** 한다. 

- [사실] **라벨(머릿글) 안 붙임.** 그냥 내용만 이어붙이고, 분할 파일명만 `01.txt,` 

- `02.txt…` 로 깔끔히 저장. 

## **B2단계｜후처리 스크립트(저장만 해 둬)** 

파일명: `combine_and_split_40k.py` (메모장에 붙여넣어 저장) 

python 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

219/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# combine_and_split_40k.py
```

```
# 목적: 기존파이프라인으로생성된 "영상별 TXT"를한폴더로모으고 → 전부합쳐하나의 TX
# 사용예:
```

```
#   python combine_and_split_40k.py --root "D:\OUT" --urls "urls.txt"
# 설명:
#   --root : 영상별결과폴더들이모여있는상위폴더(그아래에 <제목 [ID]>\TXT\*.txt
#   --urls : (선택) URL 목록파일. 있으면 "가능한한" URL 순서를반영해서합침. 없으면
# 주의:
```

```
#   - 합본안에머릿글/라벨을 "추가하지않음" (형요청대로).
```