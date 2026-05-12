`폴더이름_TXT.txt` (모든 자막 통합·중복문장 제거·중복제목 통합) 

`폴더이름_01.txt` , `_02.txt` … (4만자 단위 분할)만 남고, 원본 `.vtt/.srt/.txt` 는 **성공 시 자동 삭제** . 

오류나 중간 실패가 있으면 **삭제하지 않고 보존** , `Youtube_log.txt` 에 원인 기록. 원하는 규칙 더 있으면 말만 해. (예: 괄호 태그 보존, 분할 크기 30,000자, 특정 문장 필터 등) 

**==> picture [489 x 39] intentionally omitted <==**

내가 이야기하는건 유튜브영상의 제목이 동일 할때 하나 로 통합하라는 이야기야 

7s 동안 생각함 

맞아—“유튜브 **영상 제목이 동일할 때** 하나로 합치는” 거 말하는 거지? 

이미 내가 준 `Scirpt.bat` + `Youtube.py` 는 제목을 `%(title)s [%(id)s]` 형태로 저장하 고, `Youtube.py` 에서 **대괄호 앞(=제목 부분)**을 키로 묶어서 같은 제목이면 **하나로 통 합** 해. 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

22/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

그래도 혹시 대소문자/공백 차이 같은 **사소한 표기 차이** 때문에 따로 잡히는 걸 더 확실히 막으려면, 아래처럼 **제목 정규화** 를 살짝 강화해 주면 된다. 

## **수정 포인트 (Youtube.py만)** 

`extract_title_key` 를 아래 코드로 교체하면 돼. (함수 하나만 바꾸면 끝) 

python 코드 복사 `import unicodedata import re import os` 

```
defnormalize_title(title: str) -> str:
"""
```

```
제목동등성강화를위한가벼운정규화:
    - 유니코드 NFKC 정규화 (풀/반각, 합성문자통일)
    - 앞뒤/다중공백정리
    - 대소문자무시(영문만영향)
    - 눈에띄는의미없는제어·제로폭문자제거
    """
    t = unicodedata.normalize("NFKC", title)
# 제로폭/제어문자제거
''
    t = re.sub(r'[\u200B-\u200D\uFEFF]', , t)
# 다중공백 -> 한칸
    t = re.sub(r'\s+', ' ', t).strip()
# 영문만소문자화(국문영향없음)
    t = t.lower()
return t
```

```
defextract_title_key(path):
"""
파일명예시:
      - 제목 [abc123].ko.vtt
      - 제목 [abc123].vtt
대괄호 [id] 앞부분을 '제목'으로간주하고정규화해키로사용.
    """
    name = os.path.basename(path)
    base_no_ext = os.path.splitext(name)[0]
# 언어코드점표기제거 (예: .ko, .ko-KR 등) -> os.path.splitext가앞에서한번만
''
    base = re.sub(r'\.(ko|ko-KR|kor|ko_.*)$', , base_no_ext, flags=re.I)
```

```
    m = re.match(r'^(?P<title>.+?)\s*\[[^\]]+\]$', base)
if m:
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

23/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
        raw_title = m.group('title').strip()
else:
```

```
        raw_title = base.strip()
```

```
return normalize_title(raw_title)
```

## 이렇게 바꾸면: 

`같은 제목` 인데 **공백이 다르거나** `Half/Full width` , `제로폭 문자` , `영문 대소문자` 만 다 른 경우도 **동일 제목으로 인식** 돼서 하나로 통합돼. 

그중 **가장 긴 텍스트** (보통 정식/수동 자막)를 대표로 채택하는 기존 로직은 그대로 유 지돼. 

원래 파일에서 바뀌는 건 이 함수들뿐이야. 나머지는 그대로 쓰면 돼. 

더 강한 통합(예: `"[자막]" 같은 말머리/말미 제거` , `이모지 제거` )도 원하면 규칙 추가해 줄 게. 

**==> picture [489 x 38] intentionally omitted <==**

이게 적용이된 py 파일 bat 파일 만들어줘 

형, 요청대로 **제목 정규화(동일 제목 통합 강화)** 적용해서 두 파일 완전체로 다시 줬어. 파일명은 고정: 배치 = **Scirpt.bat** , 파이썬 = **Youtube.py** . 

## **Scirpt.bat (이름 그대로 저장)** 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브목록처리(폴더지정) - TXT 통합/분할전용
```

```
rem 현재배치파일이있는폴더경로(절대경로)
```

```
set "ROOT=%~dp0"
```

```
echo 저장할폴더이름을입력하세요. (빈칸이면날짜_시각으로자동생성)
set "FOLDER="
set /p FOLDER=폴더명:
if "%FOLDER%"=="" (
```

```
  for /f %%i in ('powershell -NoProfile -Command "(Get-Date).ToString(\"yyyy-MM-dd
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

24/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
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
pushd "%DEST%"
rem 기존 urls.txt 정리
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
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

25/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
)
```

```
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
)
popd
pause
exit /b %ERR%
```

## **Youtube.py (이름 그대로 저장)** 

python 코드 복사 

```
# -*- coding: utf-8 -*-
"""
```

```
요구사항:
```

- `1) 폴더 내 VTT/SRT/TXT 자막을 통합: 타임스탬프/태그/줄단위 잘림 제거, 자연 연결.` 

- `2) 중복 문장 제거(등장 순서 유지).` 

- `3) 통합 텍스트를 4만자 기준으로 분할: 폴더이름_01.txt, _02.txt ...` 

- `4) 통합본 "폴더이름_TXT.txt" 생성 후, 원본 VTT/SRT/TXT(중간산출물)는 삭제.` 

- `5) 오류 발생 시 삭제 중단(보존), 로그 파일 남김.` 

- `6) 동일 제목(중복 URL)의 자막은 "제목" 단위로 하나로 통합(표기차 무시; 가장 긴 텍스트 채 """` 

```
import sys
import re
import os
import glob
import unicodedata
```

```
CHUNK_SIZE = 40000# 4만자
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

26/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
deflog_write(dest, msg):
withopen(os.path.join(dest, "Youtube_log.txt"), "a", encoding="utf-8") as log
        log.write(msg + "\n")
# ===== 제목정규화(동일제목통합강화) =====
defnormalize_title(title: str) -> str:
"""
    - 유니코드 NFKC 정규화(풀/반각·합성문자통일)
    - 제로폭/제어유니코드제거
    - 다중공백 -> 한칸
    - 영문만소문자화(한글영향없음)
    """
    t = unicodedata.normalize("NFKC", title)
    t = re.sub(r'[\u200B-\u200D\uFEFF]', '', t)  # zero-width
    t = re.sub(r'\s+', ' ', t).strip()
    t = t.lower()
return t
defextract_title_key(path):
"""