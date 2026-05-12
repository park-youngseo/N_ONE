
```
for p in paras:
iflen(cur) + len(p) + 2 <= chunk_size:
            cur += (("\n\n"if cur else"") + p)
else:
if cur:
                chunks.append(cur); cur = ""
# 단일문단이 chunk_size보다길면강제분할
whilelen(p) > chunk_size:
                chunks.append(p[:chunk_size])
                p = p[chunk_size:]
            cur = p
if cur: chunks.append(cur)
# 비어있는조각은제거
return [c for c in chunks if c.strip()]
```

```
defmain():
```

```
    ap = argparse.ArgumentParser()
"
    ap.add_argument("--url", required=True, help=유튜브영상또는재생목록 URL")
"
    ap.add_argument("--out", default=".", help=저장폴더(기본현재폴더)")
"
    ap.add_argument("--lang", default="ko", help=자막언어(기본 ko)")
    ap.add_argument("--chunk-size", type=int, default=40000, help="TXT 조각크기(기
""
    ap.add_argument("--cleanup", action="store_true", help=처리후 .vtt 삭제)
""
    ap.add_argument("--cookies", default=, help="cookies.txt 경로(선택)")
    args = ap.parse_args()
```

```
    root = Path(args.out).expanduser().resolve()
```

```
    cookies = Path(args.cookies) if args.cookies elseNone
```

```
print("[1/2] 자막다운로드중…")
    vtts, metas = download_vtt(args.url, root, lang=args.lang, cookies=cookies)
ifnot vtts:
print("[오류] 자막(.vtt)을찾지못했습니다. URL/언어/권한을확인하세요."); ret
for vtt, meta inzip(vtts, metas):
        base_dir  = root / f"{meta['title']} [{meta['id']}]"
        txt_dir   = base_dir / "TXT"
        chunk_dir = base_dir / "TXT_chunks"
        txt_dir.mkdir(parents=True, exist_ok=True)
        chunk_dir.mkdir(parents=True, exist_ok=True)
```

```
print(f"[2/2] VTT → TXT 정리및 4만자조각분할: {vtt.name}")
        text = vtt_to_text(vtt)
# 전체 TXT 저장
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 177/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

```
        whole_txt = txt_dir / f"{meta['title']} [{meta['id']}].txt"
        whole_txt.write_text(text, encoding="utf-8")
# 4만자조각
        parts = split_text_by_chars(text, chunk_size=args.chunk_size)
for i, c inenumerate(parts, 1):
            (chunk_dir / f"{i:02d}.txt").write_text(c, encoding="utf-8")
print(f"    결과: {len(parts)}개조각 → {chunk_dir}")
if args.cleanup:
try: vtt.unlink()
except: pass
print("\n[완료] 저장위치:", root)
print("       구조: <제목 [ID]>/TXT/*.txt, <제목 [ID]>/TXT_chunks/01.txt, 02.tx
if __name__ == "__main__":
    main()
```

## **2) 원클릭 배치:** `유튜브→TXT_40k_원클릭.bat` 

같은 폴더에 아래 저장(더블클릭 실행): 

bat 코드 복사 `@echo off chcp 65001 > nul setlocal title 유튜브 → TXT → 4만자 조각(01/02/03…) 원클릭` 

```
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
python -m pip install yt-dlp >nul 2>nul
```

```
set "URL="
set /p URL=유튜브영상또는재생목록 URL:
if "%URL%"=="" echo [오류] URL이필요합니다.& pause & exit /b
```

```
set "OUT="
set /p OUT=저장폴더(비우면현재폴더):
if "%OUT%"=="" set "OUT=."
set "LANG="
set /p LANG=자막언어(기본 ko):
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 178/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
if "%LANG%"=="" set "LANG=ko"
```

```
set "CLEAN="
set /p CLEAN=처리후 .vtt 삭제? (Y/N, 기본 Y):
if /I "%CLEAN%"=="N" (set "CLEAN=") else (set "CLEAN=--cleanup")
```

```
echo.
```

```
echo [실행] yt_to_txt_chunks.py --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-
python "yt_to_txt_chunks.py" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-si
```

```
echo.
```

```
echo [완료] TXT_chunks 안의 01.txt부터순서대로나에게올려주면돼요.
pause
```

## **3) 사용법(정말 쉬운 버전)** 

## **1. 최초 1회** 

css 

코드 복사 

```
python -m pip install --upgrade pip
python -m pip install yt-dlp
```

**2.** `유튜브→TXT_40k_원클릭.bat` **더블클릭** → 

   - URL 붙여넣기 → (저장 폴더 선택) → 언어=ko → **VTT 삭제=Y** . 

**3.** 생성된 ** `TXT_chunks/01.txt` **부터 **한 파일씩** 여기 올리면, 

   - 내가 **바로** (형 규격대로) 

      - 핵심 10~15문장 

      - 실행 체크리스트 

      - 주의할 함정 3개 

- 로 브레인스토밍해줄게. [사실] 

 

## **덧붙임(문제 상황)** 

- 자막이 **없는 영상** 이면 TXT가 안 생긴다 → 유튜브 설정/다른 언어( `--lang en` )로 시 도 or 해당 영상 건너뛰기. 

- 로그인/연령 제한이면 `cookies.txt` 가 필요할 수 있어: 

nginx 코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

179/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python yt_to_txt_chunks.py --url "..." --cookies "C:\path\cookies.txt"
```

형, 이제부터는 **PDF 완전 배제** , **TXT 4만자 단위** 만 쓰자. 

`01.txt` 만 올려도 바로 처리 시작 가능해. 

**==> picture [489 x 39] intentionally omitted <==**

python -m pip install --upgrade pip python -m pip install yt-dlp 

기존에 있는 부분에 새로 설치 하는거지? 

13s 동안 생각함 

형, 맞아. 그 두 줄은 **새 Python을 또 깔라는 게 아니고** , 이미 있는 파이썬에 **필요한 도구 를 최신으로 맞추는** 거야. 안전하게 **여러 번 실행해도 괜찮아** . 

## **A단계｜핵심** 

- [사실] `python -m pip install --upgrade pip` → **현재 파이썬의 pip만 최신으로 업데 이트** . 

- [사실] `python -m pip install yt-dlp` → yt-dlp를 **현재 파이썬에 설치** (이미 있으면 “Requirement already satisfied”가 뜸). 

- [권장됨] **업그레이드** 까지 원하면 `-U` (= `--upgrade` )를 붙여: `python -m pip install` 

- `-U yt-dlp` . 

## **B2단계｜빠른 체크 & 실행 순서** 

## **1. 지금 어떤 파이썬인지 확인** 

sql 코드 복사 `python -V python -m pip` _`--version`_ `where python` 

(보이는 경로의 파이썬에 설치돼.) 

## **2. 업데이트 & 설치(또는 업그레이드)** 

css 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

180/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m pip install --upgrade pip
python -m pip install -U yt-dlp
```

## **3. 설치 확인** 

css 코드 복사 `python -m yt_dlp --version` 

(PATH 이슈를 피하려고 `python -m yt_dlp` 형태를 권장) 

## **4. 권한 오류가 나면** 

css 코드 복사 

```
python -m pip install -U yt-dlp --user
```

## **5. 가상환경(venv) 쓰는 중이면** 

먼저 활성화하고 같은 명령을 실행: 

코드 복사 

```
.\.venv\Scripts\activate
python -m pip install -U yt-dlp
```

요약: **기존 파이썬에 덮어 업데이트/설치** 하는 거라서 걱정 안 해도 돼. 필요한 건 저 2~3 줄이면 끝! 

**==> picture [489 x 39] intentionally omitted <==**

1,yt_to_txt_chunks.py 2,유튜브→TXT_40k_원클릭.bat 

는 왜 두개가 있는거야 각각 어떤 역할을해? 

14s 동안 생각함 

형, 딱 두 줄로 끝! 

## **A단계｜한눈에** 

**1)** `yt_to_txt_chunks.py` **= 엔진** 

유튜브 자막 → **깨끗한 TXT** → **4만 자씩** `01.txt, 02.txt…` 로 나누는 **실제 작업자** . 
