
 

rem 폴더 생성 set "DEST=%ROOT%%FOLDER%" if not exist "%DEST%" mkdir "%DEST%" 

echo. echo [안내] URL을 한 줄씩 붙여넣고 Enter 하세요. echo        아무것도 입력하지 않고 Enter 하면 다운로드 를 시작합니다. echo. 

rem 작업 폴더로 이동 pushd "%DEST%" 

rem 기존 urls.txt 정리 if exist urls.txt del /f /q urls.txt 

:INPUT_LOOP set "URL=" set /p URL=URL: if "!URL!"=="" goto START echo !URL!>>urls.txt goto INPUT_LOOP 

:START 

if not exist urls.txt ( echo [오류] URL이 없어 작업을 종료합니다. popd pause exit /b ) 

echo. echo [1/3] 자막 받는 중(ko 자동자막)... set "COOK=" if exist "%ROOT%cookies.txt" ( copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul set "COOK=--cookies cookies.txt" ) for /f "usebackq delims=" %%U in ("urls.txt") do ( echo [DL] %%U yt-dlp %COOK% --write-auto-subs --sub-langs "ko" -- 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

12/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

skip-download "%%U" ) 

echo. echo [2/3] TXT만 생성 중... python "%ROOT%subs_cleaner_plus.py" --no-srt 

echo. echo [3/3] 완료! 결과는 아래 폴더에 있습니다: echo   %DEST% 

echo   *.clean.txt 파일만 사용하시면 됩니다. 

popd pause 

이게 전에 만든 유튜브 스크립트 bat 이야 지금부터 내가 원하는 지시사항을 중점으로 새롭게 bat 파 일과 py 파일을 만들거야 

1) 지금 부터 수정된 bat 파일의 파일명을 Scirpt.bat으로 바꾼다. 

(대화를 통해 새롭게 수정된 내용이 있더라도 bat 파일의 이름은  Scirpt.bat 

으로 고정한다. gpt 멋대로 파일명을 수정하지 않는다) 2) 생성된 bat 파일을 기준으로 py파일을 만든다. py파일의 이름은 Youtube로 고정한다. 

3) 현재 bat 파일의 내용을 기준 추가 사항 

1. 지정된 폴더를 통해 만들어진 모든 TXT 파일의 중복 문 장을 정리해서 한 문장처럼 이어지도록 교체하고 타임스 탬프 및 태그, 줄단위 잘림등은 삭제한다. 

2. 통합된 txt 파일은 4만자 기준으로 분류한다. 

3. 분류된 4만자 TXT 파일은 폴더이름_01,폴더이름_02,폴 더이름_03 등으로 분류한다. 

4. 통합된 폴더이름_TXT 파일이 생성되면 나머지 VTT, TXT 파일은 삭제 한다. 

5. 중간에 에러가 생기더라도 나머지 파일은 보존한다. 

6. url의 중복된 제목은 하나로 통합한다. 

## 이렇게 만들어줘 가능해? 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

13/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

**==> picture [337 x 82] intentionally omitted <==**

24s 동안 생각함 

형, 가능해. 파일명 고정 규칙(배치: `Scirpt.bat` , 파이썬: `Youtube.py` ) 그대로 반영했고, **중복 제목 통합(6번)**까지 처리했어. 아래 코드 그대로 저장해서 쓰면 된다. 

## **Scirpt.bat (이름 그대로!)** 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브목록처리(폴더지정) - TXT 통합/분할전용
```

```
rem 현재배치파일이있는폴더경로(절대경로)
set "ROOT=%~dp0"
```

```
echo 저장할폴더이름을입력하세요. (빈칸이면날짜_시각으로자동생성)
set "FOLDER="
set /p FOLDER=폴더명:
if "%FOLDER%"=="" (
  for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToString(\"yyyy-MM-dd
)
```

```
rem 작업폴더생성
set "DEST=%ROOT%%FOLDER%"
if not exist "%DEST%" mkdir "%DEST%"
```

```
echo.
echo [안내] URL을한줄씩붙여넣고 Enter 하세요.
echo        아무것도입력하지않고 Enter 하면다운로드를시작합니다.
echo.
```

```
pushd "%DEST%"
rem 기존 urls.txt 정리
if exist urls.txt del /f /q urls.txt
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

14/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
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
  exit /b 1
)
echo.
echo [1/3] 자막받는중 (ko 우선: 수동자막있으면사용, 없으면자동자막)...
set "COOK="
if exist "%ROOT%cookies.txt" (
  copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul
  set "COOK=--cookies cookies.txt"
)
for /f "usebackq delims=" %%U in ("urls.txt") do (
  echo [DL] %%U
  rem 파일명: "제목 [영상ID].확장자" 로통일, 자막은 vtt로변환
  yt-dlp %COOK% --write-subs --write-auto-subs --sub-langs "ko,ko.*" --skip-downlo
)
echo.
echo [2/3] 통합/정리/분할처리중...
python "%ROOT%Youtube.py" "%DEST%" "%FOLDER%"
set "ERR=%ERRORLEVEL%"
echo.
if "%ERR%"=="0" (
  echo [3/3] 완료! 결과폴더:
  echo   %DEST%
  echo   - "%FOLDER%_TXT.txt" (통합본)
  echo   - "%FOLDER%_01.txt", "%FOLDER%_02.txt" ... (4만자분할본)
) else (
  echo [경고] 처리중오류가발생했지만, 기존파일은보존되었습니다.
  echo 상세는 "%DEST%\Youtube_log.txt" 확인.
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

15/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
)
```

```
popd
pause
exit /b %ERR%
```

## **Youtube.py (이름 그대로!)** 

python 

코드 복사 

```
# -*- coding: utf-8 -*-
"""
```

```
요구사항:
```

```
1) 폴더내 VTT/SRT/TXT 자막을통합: 타임스탬프/태그/줄단위잘림제거, 자연연결.
```

```
2) 중복문장제거(등장순서유지).
```

```
3) 통합텍스트를 4만자기준으로분할: 폴더이름_01.txt, _02.txt ...
```

```
4) 통합본 "폴더이름_TXT.txt" 생성후, 원본 VTT/SRT/TXT(중간산출물)는삭제.
```

```
5) 오류발생시삭제중단(보존), 로그파일남김.
```

```
6) 동일제목(중복 URL)의자막은 "제목" 단위로하나로통합(여러파일중가장긴텍스트를
"""
```

```
import sys
import re
import os
import glob
```

```
CHUNK_SIZE = 40000# 4만자
deflog_write(dest, msg):
withopen(os.path.join(dest, "Youtube_log.txt"), "a", encoding="utf-8") as log
        log.write(msg + "\n")
```

```
defextract_title_key(path):
"""
파일명예시:
      - 제목 [abc123].ko.vtt
      - 제목 [abc123].vtt
      - 제목 [abc123].ko.srt
      - 임의.txt
규칙:
      - 대괄호 [id] 앞부분을 '제목'으로간주
      - 없으면확장자제거한베이스네임사용
    """
    name = os.path.basename(path)
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

16/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
# .vtt/.srt/.txt 전언어코드점표기제거용
    base = re.sub(r'\.(ko|ko-KR|kor|ko_.*)$', '', os.path.splitext(name)[0], flag
```

```
    m = re.match(r'^(?P<title>.+?)\s*\[[^\]]+\]$', base)
if m:
        title = m.group('title').strip()
else:
        title = base.strip()
# 공백/기호정규화
    title = re.sub(r'\s+', ' ', title)
return title
defclean_lines_common(lines):
# 줄단위 -> 공백연결, 다중공백정리
    text = ' '.join(lines)
    text = re.sub(r'\s+', ' ', text).strip()
return text
defread_text_from_vtt(path):
    text_lines = []
withopen(path, 'r', encoding='utf-8', errors='ignore') as f:
for line in f:
            s = line.strip()
ifnot s: