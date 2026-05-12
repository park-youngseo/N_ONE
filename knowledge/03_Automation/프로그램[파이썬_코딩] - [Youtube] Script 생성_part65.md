
25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
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
    results = []  # [(vid,title,text)]
with YoutubeDL(ydl_opts) as ydl:
        info = ydl.extract_info(url, download=True)
        items = (info.get("entries") ifisinstance(info, dict) elseNone) or [info
for it in items:
ifnot it:
continue
            vid, title = safe(it.get("id")), safe(it.get("title"))
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 321/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
            base = f"{title} [{vid}]"
# ko 우선, 없으면다른언어라도가져오기
            cand = list(tmp_dir.glob(f"{base}.ko.vtt")) \
orlist(tmp_dir.glob(f"{base}.*.vtt")) \
orlist(tmp_dir.glob(f"{base}.vtt"))
ifnot cand:
print(f"[경고] 자막없음 → 건너뜀: {base}")
continue
            txt = vtt_to_text(cand[0])
if txt:
                results.append((vid, title, txt))
            time.sleep(random.randint(*SLEEP_BETWEEN_ITEMS))
return results
```

```
defdownload_with_retry(url: str, tmp_dir: Path, lang: str, cookies_from_browser:
"""레이트리밋등오류대비재시도"""
for attempt inrange(1, MAX_RETRIES + 1):
try:
print(f"\n[시도{attempt}/{MAX_RETRIES}] {url}")
return download_once(url, tmp_dir, lang, cookies_from_browser)
except Exception as e:
            msg = str(e)
print(f"[오류] {msg}")
if attempt < MAX_RETRIES:
# 레이트리밋/429/일시적오류 → 점증대기
                wait = 15 * attempt + random.randint(0, 10)
print(f"[대기] {wait}초후재시도…")
                time.sleep(wait)
else:
print("[포기] 최대재시도초과. 이 URL은건너뜀.")
return []
defmain():
    ap = argparse.ArgumentParser()
""
    ap.add_argument("--root", required=True, help=결과저장상위폴더)
"
    ap.add_argument("--urls", required=True, help="urls.txt 경로)
"
    ap.add_argument("--lang", default=DEFAULT_LANG, help=자막언어(기본 ko)")
"
    ap.add_argument("--chunk-size", type=int, default=DEFAULT_CHUNK, help=분할글
""
    ap.add_argument("--cookies-from-browser", default=, help="chrome/edge/firefo
    args = ap.parse_args()
```

```
    root = Path(args.root).expanduser().resolve()
    urls_file = Path(args.urls).expanduser().resolve()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 322/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
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
for i, u inenumerate(urls, 1):
# URL 사이랜덤대기 (레이트리밋완화)
if i > 1:
                d = random.randint(*SLEEP_BETWEEN_URLS)
print(f"\n[대기] 다음 URL까지{d}초대기")
                time.sleep(d)
```

```
print(f"\n[처리] ({i}/{len(urls)}) {u}")
            items = download_with_retry(u, tmp_dir, args.lang, args.cookies_from_
ifnot items:
print("[경고] 이 URL에서가져올텍스트가없습니다. (건너뜀)")
continue
for _, _, t in items:
                all_parts.append(t)
                all_parts.append("\n\n")
```

```
        combined = "".join(all_parts).strip()
ifnot combined:
```

```
print("[오류] 합칠텍스트가없습니다. (자막없음/실패)")
return
```

```
        out_dir = root / "FINAL_TXT"
        out_dir.mkdir(parents=True, exist_ok=True)
        base = safe(root.name) or"OUTPUT"
```

```
for idx, chunk inenumerate(split_by_chars(combined, args.chunk_size), 1)
            (out_dir / f"{base}_{idx:02d}.txt").write_text(chunk, encoding="utf-8
```

```
print(f"\n[완료] {out_dir}\\{base}_01.txt … (각{args.chunk_size}자)")
finally:
```

```
if tmp_dir.exists():
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 323/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
            shutil.rmtree(tmp_dir, ignore_errors=True)
```

```
if __name__ == "__main__":
    main()
```

## **바로 확인(30초 스모크 테스트)** 

**1.** 배치 실행 → 폴더 경로 입력 → 메모장에 **공개 영상 1개** 만 넣고 저장/닫기 

**2.** “쿠키 쓰나요?” → **N** 

**3.** 화면에 진행 로그가 **줄줄 뜨면 정상** . 끝나면 `FINAL_TXT\…_01.txt` 생성. 

## **자주 나는 에러 즉시 처리** 

**Private/로그인 필요** → 배치 재실행 시 **Y → chrome** (또는 edge). 

   - **rate-limited** → 위 코드가 자동 재시도/대기함. 그래도 반복되면 **urls를 20~30개씩 쪼개서** 여러 번 돌려. 

   - **아예 멈춘 것처럼** → 지금은 `quiet=False` 라 진행 로그가 보임. 로그가 안 바뀌면 그 줄을 캡처해줘. 

- 형, 이제 **파일명/흐름 그대로** 야. 위 코드 덮어쓰고 실행하면, 에러는 건너뛰고 되는 것부터 끝까지 밀어줄 거야. 

**==> picture [489 x 38] intentionally omitted <==**

플레이리스트의 영상이 200개가 넘어 그걸 하나하나 입력 하기 불편하잖아? 

20s 동안 생각함 

## 형, **200개 영상을 하나씩 넣을 필요 전혀 없어.** 

플레이리스트 **URL 한 줄만** `urls.txt` 에 넣으면 **전부 자동 처리** 돼. 

## **A단계｜전체 개요** 

- [사실] `urls.txt` 에 **플레이리스트 주소 1줄** 만 적어도, 스크립트가 **재생목록의 모든 영상** 을 순서대로 내려받아 TXT로 만들고 합쳐서 **4만 자 기준으로 분할** 해 준다. 

- [사실] 비공개·연령제한이 섞이면 `Script.bat` 질문에 **Y → chrome/edge** 로 브라우 저 쿠키를 쓰면 된다. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 324/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **B1｜초간단 사용법 (그대로)** 

**1.** `Script.bat` 더블클릭 → 정리할 폴더 경로 입력(엔터=현재). 

**2.** 자동으로 열린 메모장( `urls.txt` )에 **플레이리스트 링크 1줄** 붙여넣기 → **저장 (Ctrl+S)** → **닫기** . 

예) 

arduino 코드 복사 

```
https://www.youtube.com/playlist?list=PLxxxxxxxxxxxxxxxx
```

**3.** (물으면) **비공개/연령제한 있나요?** → 없으면 **N** , 있으면 **Y → chrome** (또는 edge). 

**4.** 끝나면 `FINAL_TXT\폴더이름_01.txt, _02.txt …` 생성. 

## **B2｜선택 옵션(원하면만)** 

**여러 재생목록** 을 한 번에: `urls.txt` 에 **재생목록 주소를 여러 줄** 로 넣기. 

- **재생목록을 “개별 영상 URL들”로 풀어 적고 싶다** 면(필수 아님), 아래 한 줄로 `urls.txt` 를 자동 생성할 수 있어: 

arduino 코드 복사 

`yt-dlp -i --flat-playlist --print "%(webpage_url)s" "<플레이리스트_URL>" > url`   

그다음 `Script.bat` 실행하면 끝. 

참고: 형이 쓰는 `Youtube.txt.py` 는 이미 **재생목록 전체 처리 + 레이트리밋 회피(랜덤 대기/재시도)** 로 짜여 있어. 주소 한 줄이면 충분해. 

**==> picture [489 x 39] intentionally omitted <==**

다시 처음부터 초등학생도 설치와 사용이 가능하도록 세 세하고 디테일하게 나열해 

17s 동안 생각함 

형, 알겠어. 더 안 바꾸고, **딱 2개 파일( Script.bat + Youtube.txt.py )** 만 쓰는 **처음부터 끝까지** 설명할게. 

그대로만 따라 하면 돼. (플레이리스트 200개도 URL 한 줄이면 자동 처리) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

325/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **0. 준비물** 