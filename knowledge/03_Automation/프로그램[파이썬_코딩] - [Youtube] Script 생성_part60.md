```
defmain_interactive():
ifnot ensure_yt_dlp():
        sys.exit(1)
```

```
# 1) 폴더물어보기
    here = Path.cwd()
    root_in = input(f"정리(저장)할폴더경로를입력하세요(엔터=현재폴더): ").strip()
    root = Path(root_in).expanduser().resolve() if root_in else here
ifnot root.exists():
print(f"[안내] 폴더가없어새로만듭니다: {root}")
        root.mkdir(parents=True, exist_ok=True)
```

```
# 2) urls.txt 편집
    urls_file = root / "urls.txt"
    edit_urls_file(urls_file)
    urls = read_urls(urls_file)
ifnot urls:
print("[오류] urls.txt에처리할링크가없습니다. (한줄=한링크)")
return
```

```
# 3) 비공개·연령제한대비(쿠키)
    cookies_from_browser = ""
if ask_yesno("로그인필요한(비공개/연령제한) 영상이있나요?", default="N"):
        b = input("사용할브라우저를입력( chrome / edge / firefox / brave ) [기본 c
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 295/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
ifnot b:
            b = "chrome"
        cookies_from_browser = b
```

```
# 4) 처리시작
    tmp_dir = root / "__vtt_tmp"
    all_parts = []
try:
for u in urls:
print(f"[처리] {u}")
            triples = download_one_url(u, tmp_dir, DEFAULT_LANG, cookies_from_brow
for _, _, txt in triples:
                all_parts.append(txt)
                all_parts.append("\n\n")  # 라벨없이공백만삽입
```

```
        combined = "".join(all_parts).strip()
ifnot combined:
print("[오류] 합칠텍스트가없습니다. (자막이없었거나모두실패)")
return
```

```
        out_dir = root / "FINAL_TXT"
        out_dir.mkdir(parents=True, exist_ok=True)
        base = safe(root.name) or"OUTPUT"
```

```
        chunks = split_by_chars(combined, size=DEFAULT_CHUNK)
for i, c inenumerate(chunks, 1):
            (out_dir / f"{base}_{i:02d}.txt").write_text(c, encoding="utf-8")
```

```
print(f"[완료] 결과: {out_dir}\\{base}_01.txt, {base}_02.txt ... (각{DEFA
finally:
# 5) 임시파일정리
if tmp_dir.exists():
            shutil.rmtree(tmp_dir, ignore_errors=True)
```

```
if __name__ == "__main__":
# 인자없이실행하면완전대화식(질문→엔터)으로동작
    main_interactive()
```

## **C단계｜문제 생기면 딱 이것만 확인** 

`python -V` 가 **버전 표시** 되나요? 아니면 파이썬 PATH 문제. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

296/341 

25. 11. 30. 오후 1:07 

 

## 프로그램[파이썬/코딩] - [Youtube] Script 생성 

- `python -m pip install -U yt-dlp` 실행했나요? **없으면 다운로드가 안 돼요.** 

- `urls.txt` 에 `#` **으로 시작하는 줄** 은 **무시** 돼요. 반드시 **링크만** 한 줄씩. 

- 비공개/연령제한 에러 나오면 실행 중 **쿠키 사용(Y)** 선택 후 **브라우저 이름** 입력. 

## 이 버전은 **배치 없이 이 파일 하나만** 으로 끝납니다. 

지금 폴더에서 바로 실행해보고, 마지막에 뜨는 메시지(에러라면 그 줄)만 보내줘. 바로 잡 아줄게. 

**==> picture [489 x 38] intentionally omitted <==**

## 미치겠네 갑자기 방법을 쳐바꿔? 

25s 동안 생각함 

## 형, 미안. **방식 안 바꾼다. 딱 2개 파일 고정** 해서 끝낸다. 

파이썬 엔진: `Youtube.txt.py` 

실행 버튼: `Script.bat` 

아래 코드 그대로 저장만 하면 된다. (다른 파일 절대 추가/변경 안 함) 

## **1)** `Script.bat` **(그대로 복붙해서 저장)** 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal
title Script — 폴더지정→urls.txt 편집→자동실행(자막→TXT→합치기→4만자분할)
```

```
REM 0) Python 확인
```

```
where python >nul 2>nul || (
```

```
  echo [오류] Python이인식되지않습니다. 설치후 "Add python.exe to PATH" 체크하세요
  pause & exit /b
```

```
)
```

```
REM 1) 엔진확인 (이름고정)
set "HERE=%~dp0"
```

```
if not exist "%HERE%Youtube.txt.py" (
```

```
  echo [오류] Youtube.txt.py 가이배치와같은폴더에없습니다.
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 297/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
  pause & exit /b
)
```

```
REM 2) 저장(정리)할폴더지정
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
REM 3) urls.txt 편집(저장후닫으면자동진행)
set "URLS=%ROOT%\urls.txt"
if not exist "%URLS%" (
```

```
  > "%URLS%" echo # 한줄에한개의유튜브링크(재생목록가능)
  >> "%URLS%" echo # 예시:
  >> "%URLS%" echo # https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX
  >> "%URLS%" echo # https://youtu.be/AAAAAAAAAAA
  >> "%URLS%" echo # 저장(Ctrl+S) 후창을닫으면계속진행됩니다.
)
echo.
echo [열기] 메모장이열리면링크를붙여넣고저장(Ctrl+S) 후창을닫으세요.
start /wait notepad "%URLS%"
```

```
REM 4) urls.txt 비었는지확인(첫유효줄검사)
set "FIRST="
for /f "usebackq tokens=1 delims= " %%L in ("%URLS%") do (
  set "LINE=%%L"
  if not "!LINE:~0,1!"=="#" (
    set "FIRST=%%L"
    goto :gotline
  )
)
:gotline
if not defined FIRST (
  echo [오류] urls.txt에처리할링크가없습니다. (# 로시작하는줄은무시)
  pause & exit /b
)
```

```
REM 5) 비공개/로그인영상대비(브라우저쿠키선택)
set "COOKIES_ARG="
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

298/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
echo.
```

```
echo [질문] 비공개/연령제한영상이있나요? 브라우저쿠키를사용할까요? (Y/N, 기본 N)
set "ANS="
```

```
set /p ANS=입력:
if /I "%ANS%"=="Y" (
  echo 사용할브라우저(chrome/edge/firefox/brave) [엔터=chrome]:
  set "BROWSER="
  set /p BROWSER=
  if "%BROWSER%"=="" set "BROWSER=chrome"
  set "COOKIES_ARG=--cookies-from-browser %BROWSER%"
)
REM 6) 도구업데이트(조용히시도)
python -m pip install -U yt-dlp >nul 2>nul
```

```
REM 7) 실행
echo.
echo [실행] 자막→TXT 생성및정리(4만자분할)를시작합니다...
python "Youtube.txt.py" --root "%ROOT%" --urls "%URLS%" --lang ko --chunk-size 400
if errorlevel 1 (
```

```
  echo [오류] 실행중문제가발생했습니다. 위메시지를확인하세요.
  pause & exit /b
)
echo.
```

```
echo [완료] 결과: "%ROOT%\FINAL_TXT\폴더이름_01.txt, _02.txt ..." (각 4만자)
pause
```

## **2)** `Youtube.txt.py` **(그대로 복붙해서 저장)** 

python 

코드 복사 

```
# Youtube.txt.py — 고정엔진
```