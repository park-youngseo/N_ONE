
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
print("[오류] urls.txt에처리할링크가없습니다. (한줄=한링크)")
return
    tmp_dir = root / "__vtt_tmp"
    all_text_parts = []
try:
for u in urls:
print(f"[처리] {u}")
            triples = download_texts(u, tmp_dir, lang=args.lang)
for _, _, text in triples:
if text:
                    all_text_parts.append(text)
                    all_text_parts.append("\n\n")  # 라벨없이공백만
```

```
        combined = "".join(all_text_parts).strip()
ifnot combined:
print("[오류] 합칠텍스트가없습니다. (자막이없었을수있음)")
return
        out_dir = root / "FINAL_TXT"
        out_dir.mkdir(parents=True, exist_ok=True)
        base = re.sub(r'[\\/:*?"<>|]+', '_', root.name).strip() or"OUTPUT"
```

```
        chunks = split_by_chars(combined, size=args.chunk_size)
for i, c inenumerate(chunks, 1):
            (out_dir / f"{base}_{i:02d}.txt").write_text(c, encoding="utf-8")
```

```
print(f"[완료] 최종파일: {out_dir}\\{base}_01.txt ... (각{args.chunk_size
finally:
```

```
# 임시 VTT 폴더정리
if tmp_dir.exists():
            shutil.rmtree(tmp_dir, ignore_errors=True)
if __name__ == "__main__":
    main()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 281/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **C단계｜초등학생도 따라 하기(진짜 짧게)** 

**1. 같은 폴더** 에 두 파일 넣기: `Script.bat` , `Youtube.txt.py` 

**2.** `Script.bat` **더블클릭** → 폴더 경로 입력(비우면 현재 폴더) 

**3.** 열리는 **메모장(urls.txt)**에 **링크 붙여넣기 → 저장(Ctrl+S) → 닫기** 

**4.** 자동 실행됨 → 끝나면 `FINAL_TXT\폴더이름_01.txt, _02.txt ...` 만 생김 

질문했던 것 답: **네. urls.txt에 “저장만 하면” 자동으로 실행** 이어진다. 

**==> picture [489 x 39] intentionally omitted <==**

아무것도 실행이 안돼 

7s 동안 생각함 

형, 오케이. “아무것도 실행 안 됨”이면, 거의 항상 아래 네 가지 중 하나야: 

① 파일 이름/확장자 틀림 ② 파일 위치 틀림 ③ 파이썬 인식 안 됨 ④ 배치 파일이 멈춰 있음(안 보임). 

바로 고치는 **한 방 절차 + 진단용 배치** 넣어줄게. 딱 이것만 해봐. 

## **A단계｜딱 5분 ‘한 방’ 점검** 

## **1. 폴더 안에 딱 2개 파일만** 있는지 확인 

```
Script.bat
```

```
Youtube.txt.py
```

(둘 다 **확장자 보이기** 켜고 이름 정확히 확인: 탐색기 상단 → 보기 → **파일 확장 명** 체크) 

## **2. 같은 폴더에서 검은창 열기** 

폴더 빈 공간 클릭 → 주소창에 `cmd` 입력 → Enter 

## **3. 파이썬 되는지 확인** 

nginx 코드 복사 

```
python -V
```

예: `Python 3.11.9` 가 나오면 OK 

안 나오면: 파이썬 재설치(“Add python.exe to PATH” 체크) 후 다시 `python -V` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 282/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **4. 엔진 파일 실행 테스트 (직접)** 

bash 

코드 복사 

```
python "Youtube.txt.py" --help
```

사용법(help) 문구가 뜨면 OK 

에러면 파일명이 다르거나 손상됨 → 이름/경로 다시 확인 

## **5. 배치 파일 실행(로그 보이게)** 

코드 복사 

```
Script.bat
```

- 메모장(urls.txt)이 열리면 링크 붙여넣고 **Ctrl+S 저장 → 메모장 닫기** 

- 저장하고 닫으면 자동으로 진행돼서 `FINAL_TXT\폴더이름_01.txt …` 가 생겨야 정 상 

## **B1｜가장 흔한 원인 3가지 (바로 해결)** 

## **1) 파일 이름/확장자 문제** 

- `Script.bat` 가 ** `Script.bat.txt` **로 저장되어 있으면 실행 안 돼. 

- `Youtube.txt.py` 가 ** `Youtube.txt.py.txt` **로 저장되어 있으면 인식 안 됨. 

- → 확장자 보이기 켜고 **.bat / .py** 로 끝나는지 꼭 확인. 

## **2) 파일 위치 문제** 

- `Script.bat` 와 `Youtube.txt.py` 는 **반드시 같은 폴더** 에 있어야 해. 

- 더블클릭하는 `Script.bat` 와 같은 곳에 엔진(.py)이 없으면 “없다” 메시지 뜸. 

## **3) 파이썬 PATH 문제** 

`python -V` 가 안 나오면 배치가 아무 것도 못 함. 

- 해결: 파이썬 재설치 → “ **Add python.exe to PATH** ” 체크 → 다시 시도. 

## **B2｜교체용 배치(진단 메시지 강화 버전)** 

지금 쓰는 `Script.bat` 내용을 **아래로 갈아끼우기** . 

(실행 단계마다 **무엇이 안 되는지** 빨간 줄로 딱 알려줌. 저장만 하면 됨.) 

bat 코드 복사 `@echo off chcp 65001 > nul` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 283/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
setlocal
title Script — 폴더지정→urls.txt 편집→자동실행
```

```
echo [체크] Python…
where python >nul 2>nul || (
  echo [오류] Python이인식되지않습니다. (python -V 가안됩니다)
  echo        파이썬재설치후 "Add python.exe to PATH" 를체크하세요.
  pause & exit /b
)
```

```
set "HERE=%~dp0"
if not exist "%HERE%Youtube.txt.py" (
  echo [오류] Youtube.txt.py 가이배치와같은폴더에없습니다.
  echo        현재폴더: "%HERE%"
  echo        Youtube.txt.py 를같은폴더로옮긴뒤다시실행하세요.
  pause & exit /b
)
echo.
set "ROOT="
set /p ROOT=정리(저장)할폴더경로(엔터=현재):
if "%ROOT%"=="" set "ROOT=%HERE%"
set "ROOT=%ROOT:"=%"
if not exist "%ROOT%" (
  echo [안내] 폴더가없어새로만듭니다: "%ROOT%"
  md "%ROOT%" >nul 2>nul
)
set "URLS=%ROOT%\urls.txt"
if not exist "%URLS%" (
  echo [안내] urls.txt 를새로만듭니다: "%URLS%"
```

```
  > "%URLS%" echo # 한줄에한개의유튜브링크(재생목록가능)
  >> "%URLS%" echo # 예시:
```

```
  >> "%URLS%" echo # https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX
  >> "%URLS%" echo # https://youtu.be/AAAAAAAAAAA
  >> "%URLS%" echo # 저장(Ctrl+S) 후메모장을닫으면자동진행됩니다.
)
```

```
echo.
```

```
echo [열기] 메모장이열립니다. 링크붙여넣고 **저장(Ctrl+S)** 한다음창을닫으세요.
start /wait notepad "%URLS%"
```

```
echo.
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 284/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

```
echo [체크] urls.txt 내용확인…
for /f "usebackq tokens=1 delims= " %%L in ("%URLS%") do (
  set "FIRST=%%L"
  goto :gotline
)
:gotline
if not defined FIRST (
  echo [오류] urls.txt 가비어있습니다. (한줄=한링크)
  echo        다시실행하여링크를넣고저장하세요.
  pause & exit /b
)
echo [업데이트] yt-dlp 최신화(조용히시도)…
python -m pip install -U yt-dlp >nul 2>nul
```

```
echo.
echo [실행] python "Youtube.txt.py" --root "%ROOT%" --urls "%URLS%" --lang ko --ch
python "Youtube.txt.py" --root "%ROOT%" --urls "%URLS%" --lang ko --chunk-size 400
if errorlevel 1 (
```

```
  echo [오류] 엔진실행중문제가발생했습니다. 위에러메시지를캡처해주세요.
  pause & exit /b
)
echo.
```

```
if exist "%ROOT%\FINAL_TXT" (
```

```
  echo [완료] 결과: "%ROOT%\FINAL_TXT\폴더이름_01.txt, _02.txt ..."
) else (
```

```
  echo [경고] FINAL_TXT 폴더가보이지않습니다. (자막부재/URL 오류일수있음)