  exit /b 1
)
```

```
echo.
echo [1/3] 자막받는중 (세이프모드: 속도제한 + 랜덤딜레이 + 재시도)...
set "COOK="
if exist "%ROOT%cookies.txt" (
  copy /Y "%ROOT%cookies.txt" "cookies.txt" >nul
  set "COOK=--cookies cookies.txt"
)
set "UA=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like
set "LIMIT=1.5M"   rem 요청속도상한(대역폭낮추기)
set /a COUNT=0
for /f "usebackq delims=" %%U in ("urls.txt") do (
  set "ONEURL=%%U"
  set /a COUNT+=1
  call :DOWNLOAD "!ONEURL!"
  if errorlevel 1 (
    echo [경고] %%U 처리실패(2회시도후) → 건너뜀
  )
  rem 각영상사이 5~12초랜덤대기
  call :JITTER 5 13
  rem 15개처리마다 60~120초휴식
  if !COUNT! GEQ 15 (
    echo [세이프] 15개처리완료 → 길게휴식
    set /a COUNT=0
    call :JITTER 60 121
  )
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
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

38/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
) else (
```

```
  echo [경고] 처리중오류가발생했지만, 기존파일은보존되었습니다.
  echo 상세는 "%DEST%\Youtube_log.txt" 확인.
)
popd
pause
exit /b %ERR%
```

```
rem ================== 함수들 ==================
```

```
:DOWNLOAD
rem 입력: %1 = URL
setlocal
set "U=%~1"
echo [DL] %U%
rem 수동자막(ko) 우선, 없으면자동자막(ko). 파일명은 "제목 [id].vtt"
yt-dlp %COOK% --user-agent "%UA%" --write-subs --sub-langs "ko" ^
```

```
  --skip-download --convert-subs vtt --limit-rate %LIMIT% ^
```

```
  -o "%%(title)s [%%(id)s].%%(ext)s" "%U%"
if errorlevel 1 (
```

```
  echo [경고] 1차실패 → 60~120초대기후자동자막으로재시도
  endlocal & call :JITTER 60 121 & setlocal
```

```
  yt-dlp %COOK% --user-agent "%UA%" --write-auto-subs --sub-langs "ko" ^
```

```
    --skip-download --convert-subs vtt --limit-rate %LIMIT% ^
```

```
    -o "%%(title)s [%%(id)s].%%(ext)s" "%U%"
```

```
)
endlocal & exit /b %errorlevel%
```

```
:JITTER
```

```
rem 사용법: call :JITTER 최소초최대초(최대초는포함되지않음)
```

```
powershell -NoProfile -Command "$s=Get-Random -Minimum %1 -Maximum %2; Start-Slee
exit /b 0
```

## **Youtube.py (변경 없음 – 동일 제목 통합/중복문장 제거/4 만자 분할/성공 시 원본 삭제)** 

지금 사용 중인 최신 버전 그대로 쓰면 돼. (이미 `ko` / `ko-orig` 같은 꼬리표 처리 + 동일 제목 하나로 통합 + 가장 긴 텍스트 선별 로직 포함) 

## **왜 이게 효과 있나?** 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

39/65 

25. 11. 30. 오후 1:39 

 

프로그램 - 유튜브 배치파일 수정 

**속도 제한 + 랜덤 대기** : 봇 패턴(규칙적·고빈도 요청)으로 보이는 걸 피함. 

- **수동→자동 2단계 재시도** : 첫 요청 실패(429/일시 차단) 시 바로 연속 재요청하지 않 고 **길게 휴식 후** 모드 바꿔 재시도. 

**쿠키 사용** : 자신의 세션으로 정상이용처럼 보이도록(필요 시). 

**요청 최소화** : 우리는 **자막만** 받고, 통합 성공 시 원본을 삭제해 **재요청** 가능성을 낮춤. 

원하면 “1회 실행 당 최대 N개만 처리” 옵션도 넣어줄게(예: 밤에 30개만, 새벽에 30개). 또, `Youtube.py` 에 **제목·영상ID·자막길이·대표선택 로그** 남기는 기능을 추가해 원인 파악 을 더 쉽게 만들 수도 있어. 

**==> picture [489 x 39] intentionally omitted <==**

Kind: captions Language: ko 안녕하세요 오토 코드입니다 이번 안녕하세요 오토 코드입니다 이번 안녕하세요 오토 코드입니다 이번 시간에는 동영상 자동 재생 웹앱을 시간 에는 동영상 자동 재생 웹앱을 시간에는 동영상 자동 재생 웹앱을 한번 만들어 보도록 하겠습니다 우리가 한번 만들 어 보도록 하겠습니다 우리가 한번 만들어 보도록 하겠습 니다 우리가 지난 시간까지는 오픈 AA 채지 PT 지난 시간 까지는 오픈 AA 채지 PT 지난 시간까지는 오픈 AA 채지 PT 같은 어 랭 체인이라는 프레임워크를 같은 어 랭 체인 이라는 프레임워크를 같은 어 랭 체인이라는 프레임워크 를 이용해서 우리가 오픈의 API 활용해 이용해서 우리가 오픈의 API 활용해 이용해서 우리가 오픈의 API 활용해 가 지고 어 우리가 엑셀 그니까 가지고 어 우리가 엑셀 그니 까 가지고 어 우리가 엑셀 그니까 여러분들의 자신의 엑셀 데이터에 대한 여러분들의 자신의 엑셀 데이터에 대한 여 러분들의 자신의 엑셀 데이터에 대한 질문에 대한 답변을 받아보는 걸로 질문에 대한 답변을 받아보는 걸로 질문에 대한 답변을 받아보는 걸로 했거든요 자 퀘스천에 대한 답 변을 했거든 

txt 파일에 중복된 내용을 하나로 통합하라고 했더니 지금 3번이나 반복되는 형태로 만들면 어떻게 하자는 거지? 

21s 동안 생각함 

형, 저런 “세 번씩 우르르 반복”은 **문장 전체가 합쳐진 뒤 내부에서 구간(단어·구절) 자체 가 반복** 된 케이스야. 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

40/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

그래서 문장 단위 중복 제거만으로는 못 잡혀. → **연속 반복 구절(단어/2~8-gram) 압축 단계** 를 하나 더 넣어 해결했어. 

아래처럼 **Youtube.py만 교체** 하면 끝. (파일명은 그대로 `Youtube.py` 유지) 

## **Youtube.py (업데이트: 반복 구절 압축 추가)** 

python 코드 복사 

```
# -*- coding: utf-8 -*-
```

```
import sys, re, os, glob, unicodedata
```

```
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
# ===== 제목정규화 =====
```

```
defnormalize_title(title: str) -> str:
```

```
    t = unicodedata.normalize("NFKC", title)
```

```
''
    t = re.sub(r'[\u200B-\u200D\uFEFF]', , t)
```

```
    t = re.sub(r'\s+', ' ', t).strip()
```

```
    t = t.lower()
return t
```

```
defextract_title_key(path):
```

```
    name = os.path.basename(path)
```

```
    base_no_ext = os.path.splitext(name)[0]
```

```
# ko / ko-KR / ko-orig / en / zh-Hans 등꼬리표제거
```

```
''
    base = re.sub(r'\.(?:[a-z]{2,3}(?:-[A-Za-z0-9]+)*)(?:-orig)?$', , base_no_ex
    m = re.match(r'^(?P<title>.+?)\s*\[[^\]]+\]$', base)
```

```
    raw_title = (m.group('title') if m else base).strip()
return normalize_title(raw_title)
```

```
# ======================
```

```
defclean_lines_common(lines):
```

```
    text = ' '.join(lines)
    text = re.sub(r'\s+', ' ', text).strip()
return text
```

```
defread_text_from_vtt(path):
    out=[]
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

41/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
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
```

```