25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

@echo off chcp 65001 > nul setlocal enabledelayedexpansion title 유튜브 목록 처리(폴더 지정) - TXT만 생성 

rem 현재 배치파일이 있는 폴더 경로(절대경로) set "ROOT=%~dp0" 

echo 저장할 폴더 이름을 입력하세요. (빈칸이면 날짜_시 각으로 자동 생성) set "FOLDER=" set /p FOLDER=폴더명: if "%FOLDER%"=="" ( for /f %%i in ('powershell -NoProfile -Command "(GetDate).ToString(\"yyyy-MM-dd_HHmm\")"') do set "FOLDER=%%i" ) rem 폴더 생성 set "DEST=%ROOT%%FOLDER%" if not exist "%DEST%" mkdir "%DEST%" echo. 

echo [안내] URL을 한 줄씩 붙여넣고 Enter 하세요. echo        아무것도 입력하지 않고 Enter 하면 다운로드 를 시작합니다. echo. 

rem 작업 폴더로 이동 pushd "%DEST%" rem 기존 urls.txt 정리 if exist urls.txt del /f /q urls.txt 

:INPUT_LOOP set "URL=" set /p URL=URL: if "!URL!"=="" goto START echo !URL!>>urls.txt goto INPUT_LOOP 

:START 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

1/10 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

if not exist urls.txt ( echo [오류] URL이 없어 작업을 종료합니다. popd pause exit /b ) echo. echo [1/3] 자막 받는 중(ko 자동자막)... set "COOK=" if exist "%ROOT%cookies.txt" ( copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul set "COOK=--cookies cookies.txt" ) 

for /f "usebackq delims=" %%U in ("urls.txt") do ( echo [DL] %%U yt-dlp %COOK% --write-auto-subs --sub-langs "ko" -- skip-download "%%U" ) 

echo. echo [2/3] TXT만 생성 중... python "%ROOT%subs_cleaner_plus.py" --no-srt 

echo. 

echo [3/3] 완료! 결과는 아래 폴더에 있습니다: echo   %DEST% 

echo   *.clean.txt 파일만 사용하시면 됩니다. 

popd pause 

이게 전에 만든 유튜브 스크립트 bat 이야 지금부터 내가 원하는 지시사항을 중점으로 새롭게 bat 파 일과 py 파일을 만들거야 

1) 지금 부터 수정된 bat 파일의 파일명을 Scirpt.bat으로 바꾼다. (대화를 통해 새롭게 수정된 내용이 있더라도 bat 파일의 이름은  Scirpt.bat 

으로 고정한다. gpt 멋대로 파일명을 수정하지 않는다) 2) 생성된 bat 파일을 기준으로 py파일을 만든다. 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

2/10 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

py파일의 이름은 Youtube로 고정한다. 

3) 현재 bat 파일의 내용을 기준 추가 사항 

1. 지정된 폴더를 통해 만들어진 모든 TXT 파일의 중복 문 장을 정리해서 한 문장처럼 이어지도록 교체하고 타임스 탬프 및 태그, 줄단위 잘림등은 삭제한다. 

2. 통합된 txt 파일은 4만자 기준으로 분류한다. 

3. 분류된 4만자 TXT 파일은 폴더이름_01,폴더이름_02,폴 더이름_03 등으로 분류한다. 

4. 통합된 폴더이름_TXT 파일이 생성되면 나머지 VTT, TXT 파일은 삭제 한다. 

## 이렇게 만들어줘 가능해? 

23s 동안 생각함 

- 형, 가능하지! 아래처럼 **파일명 고정 규칙** 을 지키면서 새로 설계했어. 

   - 배치파일 이름: `Scirpt.bat` (철자 그대로 고정) 

   - 파이썬 파일 이름: `Youtube.py` (확장자는 `.py` , 베이스 이름은 “Youtube”로 고정) 

- 이 조합은: 

**1.** URL들에서 **한글 자동자막(ko)**만 받아오고, 

**2.** 모든 자막 파일을 모아 **텍스트를 정리/통합(중복 제거·타임스탬프/태그 제거·줄 단 위 끊김 보정)**하고, 

**3. 40,000자 기준으로 분할** 하여 `폴더이름_01.txt, _02.txt…` 생성, 

**4.** 마지막에 **통합본과 분할본을 제외한 모든** `.vtt/.srt/.txt` **잔여물 삭제** 까지 자동으 로 처리해. 

## **Scirpt.bat (이름 고정)** 

bat 

코드 복사 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

3/10 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브목록처리(폴더지정) - TXT 통합/분할자동
```

```
REM 현재배치파일이있는폴더경로(절대경로)
set "ROOT=%~dp0"
echo 저장할폴더이름을입력하세요. (빈칸이면날짜_시각으로자동생성)
set "FOLDER="
set /p FOLDER=폴더명:
if "%FOLDER%"=="" (
  for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToString(\"yyyy-MM-dd
)
REM 폴더생성
set "DEST=%ROOT%%FOLDER%"
if not exist "%DEST%" mkdir "%DEST%"
echo.
echo [안내] URL을한줄씩붙여넣고 Enter 하세요.
echo        아무것도입력하지않고 Enter 하면자막수집을시작합니다.
echo.
REM 작업폴더로이동
pushd "%DEST%"
REM 기존 urls.txt 정리
if exist urls.txt del /f /q urls.txt
:INPUT_LOOP
set "URL="
set /p URL=URL:
if "!URL!"=="" goto START
echo !URL!>>urls.txt
goto INPUT_LOOP
:START
if not exist urls.txt (
  echo [오류] URL이없어작업을종료합니다.
  popd
  pause
  exit /b
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

4/10 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

```
)
```

```
echo.
echo [1/3] 자막수집중(ko 자동자막)...
set "COOK="
if exist "%ROOT%cookies.txt" (
  copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul
  set "COOK=--cookies cookies.txt"
)
for /f "usebackq delims=" %%U in ("urls.txt") do (
  echo [DL] %%U
  yt-dlp %COOK% --write-auto-subs --sub-langs "ko" --skip-download --convert-subs
)
echo.
echo [2/3] 텍스트정리/통합및 4만자분할중...
REM Youtube.py 호출 (인자: 작업폴더, 폴더이름)
python "%ROOT%Youtube.py" "%DEST%" "%FOLDER%"
if errorlevel 1 (
  echo [오류] Youtube.py 실행중문제가발생했습니다.
  popd
  pause
  exit /b
)
echo.
echo [3/3] 완료! 결과는아래폴더에있습니다:
echo   %DEST%
echo   -> %FOLDER%_TXT.txt (전체통합본)
echo   -> %FOLDER%_01.txt, %FOLDER%_02.txt ... (4만자분할본)
popd
pause
```

## **Youtube.py (이름 고정)** 

python 

코드 복사 

```
# -*- coding: utf-8 -*-
"""
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

5/10 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

```
Youtube.py
```

```
- 입력폴더(DEST) 안의자막파일(.vtt/.srt/.txt)을수집
```

```
- 타임스탬프/태그제거, 줄단위잘림보정, 연속중복문장제거
- 전체통합본: <FOLDER>_TXT.txt
- 40,000자기준분할본: <FOLDER>_01.txt, _02.txt, ...
- 통합/분할본을제외한나머지 .vtt/.srt/.txt는모두삭제
사용: python Youtube.py "<DEST>" "<FOLDER>"
"""
import os
import re
import sys
import glob
import shutil
CHUNK_SIZE = 40000# 4만자
```

```
defread_text_file(path):
try:
withopen(path, 'r', encoding='utf-8') as f:
return f.read()
except UnicodeDecodeError:
withopen(path, 'r', encoding='cp949', errors='ignore') as f:
return f.read()
```

```
defwrite_text_file(path, text):
withopen(path, 'w', encoding='utf-8') as f:
        f.write(text)
```

```
defstrip_tags_and_timestamps(text, is_vtt=False, is_srt=False):
# 공통: 웹/형식태그제거
    text = re.sub(r'<[^>]+>', ' ', text)  # <c.colorE5E5E5> 등
    text = re.sub(r'\{[^}]+\}', ' ', text)  # {\\an8} 등
    text = re.sub(r'\[.*?\]', ' ', text)    # [Music], [박수] 등 (보존원하면조정)
# VTT 헤더/스타일제거
if is_vtt:
        text = re.sub(r'^\s*WEBVTT.*$', ' ', text, flags=re.MULTILINE)
        text = re.sub(r'^\s*NOTE.*?(?:\r?\n\r?\n|$)', ' ', text, flags=re.DOTALL
        text = re.sub(r'^\s*STYLE.*?(?:\r?\n\r?\n|$)', ' ', text, flags=re.DOTALL
# VTT 타임스탬프라인
        text = re.sub(r'^\s*\d{2}:\d{2}:\d{2}\.\d{3}\s-->\s\d{2}:\d{2}:\d{2}\.\d{
        text = re.sub(r'^\s*\d{2}:\d{2}\.\d{3}\s-->\s\d{2}:\d{2}\.\d{3}.*$', ' ',
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

6/10 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

```
# SRT 번호/타임스탬프제거
if is_srt:
        text = re.sub(r'^\s*\d+\s*$', ' ', text, flags=re.MULTILINE)
        text = re.sub(r'^\s*\d{2}:\d{2}:\d{2},\d{3}\s-->\s\d{2}:\d{2}:\d{2},\d{3}
# 라인말미스타일태그/이모지등간단정리
    text = re.sub(r'♪+', ' ', text)
    text = re.sub(r'…+', '…', text)
# 잔여 \r\n -> \n
    text = text.replace('\r\n', '\n').replace('\r', '\n')
return text
```

```
defnormalize_lines_and_merge(text):
# 중복공백줄이기, 라인병합
    lines = [re.sub(r'\s+', ' ', ln).strip() for ln in text.split('\n')]
    lines = [ln for ln in lines if ln]  # 빈줄제거
    merged = []
    prev = None
for ln in lines:
# 연속중복문장제거
if prev isnotNoneand ln == prev:
continue
        merged.append(ln)
        prev = ln
# 줄단위끊김보정: 문장사이공백으로합치기
    merged_text = ' '.join(merged)
# 공백정리
    merged_text = re.sub(r'\s+', ' ', merged_text).strip()
return merged_text
```

```