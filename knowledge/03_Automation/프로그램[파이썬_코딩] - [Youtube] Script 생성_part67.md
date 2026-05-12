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
                wait = 15 * attempt + random.randint(0, 10)
print(f"[대기] {wait}초후재시도…")
                time.sleep(wait)
else:
print("[포기] 최대재시도초과. 이 URL은건너뜀.")
return []
```

```
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
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 331/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    args = ap.parse_args()
```

```
    root = Path(args.root).expanduser().resolve()
    urls_file = Path(args.urls).expanduser().resolve()
assert root.exists(), f"[오류] 폴더가없습니다: {root}"
assert urls_file.exists(), f"[오류] urls.txt가없습니다: {urls_file}"
    urls = read_urls(urls_file)
ifnot urls:
print("[오류] urls.txt가비어있습니다. (#로시작하는줄은무시)")
return
    tmp_dir = root / "__vtt_tmp"
    all_parts = []
try:
for i, u inenumerate(urls, 1):
if i > 1:
                d = random.randint(*SLEEP_BETWEEN_URLS)
print(f"\n[대기] 다음 URL까지{d}초대기")
                time.sleep(d)
```

```
print(f"\n[처리] ({i}/{len(urls)}) {u}")
            items = download_with_retry(u, tmp_dir, args.lang, args.cookies_from_
ifnot items:
```

```
print("[경고] 이 URL에서가져올텍스트가없습니다. (건너뜀)")
continue
for _, _, t in items:
                all_parts.append(t)
                all_parts.append("\n\n")
```

```
        combined = "".join(all_parts).strip()
ifnot combined:
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

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 332/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
print(f"\n[완료] {out_dir}\\{base}_01.txt … (각{args.chunk_size}자)")
```

```
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

## **5. 사용 방법 (진짜 간단, 딱 4단계)** 

## **1.** `Script.bat` **더블클릭** 

**2.** “정리할 폴더 경로” → 비우면 현재 폴더 사용(그냥 Enter 눌러도 됨) 

**3.** 자동으로 **메모장(urls.txt)**가 열림 

**플레이리스트 URL 1줄** 만 붙여넣기 → **Ctrl+S(저장)** → **창 닫기** 

- 예: 

arduino 코드 복사 

```
https://www.youtube.com/playlist?list=PLxxxxxxxxxxxxxxxx
```

**주의:** `#` 로 시작하는 줄은 **무시** 됨. 링크만 넣기. 

**4.** (질문 나오면) 

비공개/연령제한 있나요? → **없으면 N** , 있으면 **Y → chrome** (또는 edge) 

그러면 자동으로 실행되고, 끝나면 결과가 생긴다: 

코드 복사 

```
...\FINAL_TXT\
폴더이름_01.txt
폴더이름_02.txt
  ...
```

## (각 4만자, 임시파일은 정리됨) 

## **6. 자주 막히는 지점 (바로 해결)** 

- **배치가 안 돎** → `Script.bat` 가 `Script.bat.txt` 로 저장된 것. 확장자 보이기 켜고 이름 수정. 

**python 인식 안 됨** → 파이썬 재설치 때 **Add python.exe to PATH** 체크. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 333/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

- **urls.txt 비어 있음** → 링크 붙여넣고 **반드시 저장(Ctrl+S)** 후 닫기. 

- **Private / 로그인 필요** → `Script.bat` 실행 시 **Y → chrome/edge** 입력. 

- **Rate limit(잠시 후 다시 시도)** → 코드가 자동으로 **랜덤 대기·재시도** 해줌. (조금 기다 리면 진행) 

**아무 것도 안 되는 느낌** → 이번 버전은 **진행 로그** 가 보이니, 마지막 3~4줄 보내주면 그 자리에서 잡아줄게. 

## 형, 이대로만 하면 돼: **두 파일 그대로, URL 한 줄, 더블클릭 한 번** . 

막히는 지점 생기면 **검은창 마지막 3~4줄** 만 보여줘. 바로 거기부터 고쳐줄게. 

**==> picture [489 x 39] intentionally omitted <==**

기존 .py, .bat 프로그램을 다른 버전으로 만들거야 

1, 새롭게 만드는 버전 엔진의 이름은 Youtube_solo.txt.py 로 고정한다 (선택을 물어보거나 하지 않는다) 2, 이름을 cript_solo.bat 으로 고정한다 (선택을 물어보거나 하지 않 는다) Script_solo.bat실행시 폴더지정을 묻는다 urls 메모 장 파일에 유튜브 링크를 넣고 저장한다  3, 저장을 하면 자동으로 txt가 지정된 폴더에 생성되고 url의 txt 제목과 내용을 그대로 출력한다. 

4, txt 내용이 4만자 이상일때 (ai가 최적으로 인식 할수 있 도록) url 제목 01.txt ,02 로 만들어진다 

## 형, 오케이. **새 버전** 은 딱 두 파일만, 이름도 고정이야. 

- 파이썬 엔진: `Youtube_solo.txt.py` 

- 실행 배치: `Script_solo.bat` 

동작 요약(형이 원하는 대로) 

- `Script_solo.bat` 실행 → **폴더 지정** → 자동으로 `urls.txt` 메모장 열림 → **유튜브** 

- **링크 붙여넣고 저장** 

- 저장 후 자동 실행 → 각 URL(영상/재생목록의 각 영상)마다 **제목 그대로의 .txt** 생성 한 파일이 **4만 자 초과** 면 자동으로 `제목_01.txt, _02.txt …` 로 분할 

## **1) Script_solo.bat (그대로 저장)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

334/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

메모장에 아래 내용을 복사 → 원하는 폴더에 `Script_solo.bat` 이름으로 저장. 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title Script_solo — 폴더지정→urls.txt 편집→자동실행(영상별 TXT, 4만자분할)
```

```
rem 0) Python 확인
where python >nul 2>nul || (
  echo [오류] Python이인식되지않습니다. 설치시 "Add python.exe to PATH" 체크가필요
  pause & exit /b
)
rem 1) 엔진확인 (고정이름)
set "HERE=%~dp0"
if not exist "%HERE%Youtube_solo.txt.py" (
  echo [오류] Youtube_solo.txt.py 가이배치와같은폴더에없습니다. 현재: "%HERE%"
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
  > "%URLS%" echo # 한줄에한개의유튜브링크(영상또는재생목록)
  >> "%URLS%" echo # 예시:
  >> "%URLS%" echo # https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX
  >> "%URLS%" echo # https://youtu.be/AAAAAAAAAAA
  >> "%URLS%" echo # 저장(Ctrl+S) 후창을닫으면계속진행됩니다.
)
echo.
echo [열기] 메모장이열리면링크를붙여넣고저장(Ctrl+S) 후창을닫으세요.
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

335/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
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