46/65 

25. 11. 30. 오후 1:39 프로그램 - 유튜브 배치파일 수정 

 

## 파이썬: **Youtube.py** (고정) 

## 이 버전은 

- 반복 구절(단어/2~8-gram) 압축 + 문장 중복 제거를 **둘 다** 적용 

- 무엇이 얼마나 줄었는지 `폴더이름_REPORT.txt` 로 **증거 리포트** 생성 

- 첫 실행은 **삭제 안 함(세이프 모드)** → 결과 확인 후 `SAFE_DELETE=1` 로 바꾸면 원본 삭제 

## **Scirpt.bat** 

bat 

코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브목록처리(폴더지정) - TXT 통합/분할(세이프+리포트)
```

```
rem === 기본경로 ===
set "ROOT=%~dp0"
```

```
echo 저장할폴더이름을입력하세요. (빈칸이면날짜_시각으로자동생성)
set "FOLDER="
set /p FOLDER=폴더명:
if "%FOLDER%"=="" (
  for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToString(\"yyyy-MM-dd
)
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
if exist urls.txt del /f /q urls.txt
```

```
:INPUT_LOOP
set "URL="
set /p URL=URL:
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

47/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
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
echo [1/3] 자막받는중 (수동 ko 우선, 없으면자동 ko)...
set "COOK="
if exist "%ROOT%cookies.txt" (
  copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul
  set "COOK=--cookies cookies.txt"
)
set "UA=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like
set "LIMIT=1.5M"
```

```
for /f "usebackq delims=" %%U in ("urls.txt") do (
  echo [DL] %%U
```

```
  yt-dlp %COOK% --user-agent "%UA%" --write-subs --sub-langs "ko" ^
    --skip-download --convert-subs vtt -o "%%(title)s [%%(id)s].%%(ext)s" "%%U" |
    powershell -NoProfile -Command "Start-Sleep -Seconds (Get-Random -Minimum 60
    yt-dlp %COOK% --user-agent "%UA%" --write-auto-subs --sub-langs "ko" ^
      --skip-download --convert-subs vtt -o "%%(title)s [%%(id)s].%%(ext)s" "%%U"
  )
  powershell -NoProfile -Command "Start-Sleep -Seconds (Get-Random -Minimum 5 -Max
)
```

```
echo.
```

```
echo [2/3] 통합/정리/분할 + 리포트생성중...
```

```
rem === 안전삭제토글: 0=삭제안함(기본), 1=성공시원본삭제 ===
set SAFE_DELETE=0
```

```
python "%ROOT%Youtube.py" "%DEST%" "%FOLDER%" --report --aggressive --safe-delete
set "ERR=%ERRORLEVEL%"
```

```
echo.
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

48/65 

25. 11. 30. 오후 1:39 

 

**==> picture [109 x 8] intentionally omitted <==**

```
if "%ERR%"=="0" (
  echo [3/3] 완료! 결과폴더: %DEST%
  echo   - "%FOLDER%_TXT.txt"        (통합본)
  echo   - "%FOLDER%_01.txt" ...     (4만자분할본)
  echo   - "%FOLDER%_REPORT.txt"     (정리/중복/압축리포트)
  if "%SAFE_DELETE%"=="1" (
    echo   - 원본 VTT/SRT/TXT는성공후삭제됨
  ) else (
    echo   - 원본파일은보존됨 (확인후 SAFE_DELETE=1로변경가능)
  )
) else (
  echo [경고] 처리중오류. 파일은보존됨. "%DEST%\Youtube_log.txt" 확인.
)
popd
pause
exit /b %ERR%
```

## **Youtube.py** 

python 

코드 복사 

```
# -*- coding: utf-8 -*-
"""
```

```
통합파이프라인 (리포트/세이프삭제/반복구절압축포함)
```

```
- 제목정규화로동일제목묶기 (ko, ko-KR, ko-orig 등꼬리표무시)
```

```
- 줄병합, 태그/타임스탬프제거
- anti_echo: 단어및 2~8-gram 연속반복압축
```

```
- 문장중복제거
- 4만자분할저장
```

```
- --safe-delete 1일때만원본자막/중간 TXT 삭제
```

```
- --report 시폴더이름_REPORT.txt 생성
"""
```

```
import sys, re, os, glob, unicodedata, argparse
CHUNK_SIZE = 40000
```

```
deflog_write(dest, msg):
```

```
withopen(os.path.join(dest, "Youtube_log.txt"), "a", encoding="utf-8") as log
        log.write(msg + "\n")
```

```
# ---------- 제목정규화 ----------
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

49/65 

25. 11. 30. 오후 1:39 

 

프로그램 - 유튜브 배치파일 수정 

```
defnormalize_title(title: str) -> str:
    t = unicodedata.normalize("NFKC", title)
''
    t = re.sub(r'[\u200B-\u200D\uFEFF]', , t)
    t = re.sub(r'\s+', ' ', t).strip()
    t = t.lower()
return t
defextract_title_key(path):
    name = os.path.basename(path)
    base_no_ext = os.path.splitext(name)[0]
# .ko / .ko-KR / .ko-orig / .en / .zh-Hans ... 꼬리표제거
''
    base = re.sub(r'\.(?:[a-z]{2,3}(?:-[A-Za-z0-9]+)*)(?:-orig)?$', , base_no_ex
    m = re.match(r'^(?P<title>.+?)\s*\[[^\]]+\]$', base)
    raw = (m.group('title') if m else base).strip()
return normalize_title(raw)
# ---------- 공통정리 ----------
defclean_lines_common(lines):
    text = ' '.join(lines)
    text = re.sub(r'\s+', ' ', text).strip()
return text
defread_text_from_vtt(path):
    out=[]
withopen(path,'r',encoding='utf-8',errors='ignore') as f:
for line in f:
            s=line.strip()
ifnot s: continue
if s.upper().startswith('WEBVTT') or s.startswith(('NOTE','STYLE')):
if'-->'in s: continue
if re.search(r'(line:|position:|size:|align:)', s, re.I): continue
if re.fullmatch(r'\d+', s): continue
''
            s=re.sub(r'<[^>]+>',,s)
''
            s=re.sub(r'[\[\(][^\]\)]{1,20}[\]\)]',,s)
            s=s.strip('-─—·• ')
if s: out.append(s)
return out
defread_text_from_srt(path):
    out=[]
withopen(path,'r',encoding='utf-8',errors='ignore') as f:
for line in f:
            s=line.strip()
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

50/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
ifnot s: continue
if re.fullmatch(r'\d+', s): continue
if re.search(r'\d{2}:\d{2}:\d{2}[,\.]\d{3}\s*-->\s*\d{2}:\d{2}:\d{2}[
''
            s=re.sub(r'<[^>]+>',,s)
''
            s=re.sub(r'[\[\(][^\]\)]{1,20}[\]\)]',,s)
            s=s.strip('-─—·• ')
if s: out.append(s)
return out
```

```
defread_text_from_txt(path):
    out=[]