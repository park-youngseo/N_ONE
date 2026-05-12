
- [사실] `python.exe` : 이름 바꾸면 `pip` , 파일 연결, `py` 런처, PATH가 깨질 수 있어. 

- [사실] `yt-dlp` **실행파일/모듈** : 이름을 바꾸면 업데이트/호출/ `import yt_dlp` 가 꼬인 다. 

- [사실] **우리 스크립트(.py)는 이름 바꿔도 가능** 하지만, 그걸 부르는 **배치(.bat)도 같이 수정** 해야 한다. 

- [권장됨] “이름을 바꾸고 싶다” → **래퍼(별명)만 만들자** : 원본은 그대로, 겉에서 새 이 름으로 호출. 

## **B2단계｜안전한 대안(두 가지)** 

## **1) 배치 래퍼(추천, 영구적)** 

**1.** 메모장에 ↓ 저장: `YouTube-script.bat` 

bat 

코드 복사 

```
@echo off
```

```
python -m yt_dlp %*
```

## **2.** 이제 이렇게 쓰면 됨: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

187/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

코드 복사 

vbnet 

 

```
YouTube-script --write-auto-subs --sub-langs "ko" --skip-download <URL>
```

→ 원본 `yt-dlp` 파일은 건드리지 않고, **새 이름(YouTube-script)** 으로 동일하게 동작. 

## **2) 임시 별명(doskey, 창 닫으면 사라짐)** 

makefile 코드 복사 

```
doskey YouTube-script=python -m yt_dlp $*
```

그 창에서: 

css 코드 복사 

```
YouTube-script --version
```

## **우리 파일들은?** 

`yt_to_txt_chunks.py` **(엔진)** : 이름 바꿔도 됨. 단, 배치에서 호출 줄도 같이 바꿔: 

bat 코드 복사 

```
python "내가_바꾼_이름.py" --url "..."
```

`유튜브→TXT_40k_원클릭.bat` **(버튼)** : 안에 적힌 스크립트 파일명을 **동일하게** 맞추면 OK. 

## **한 줄 정리** 

**바꾸면 안 됨** : `python.exe` , `yt-dlp` 실행파일/모듈 이름 

**바꿔도 됨** : 우리가 만든 `.py` 스크립트(배치의 호출도 함께 수정) 

- **원하는 이름으로 쓰기** : **배치/별명(래퍼)** 로 해결하는 게 정답 

원하면 내가 **배치에서 스크립트 이름을 자동 탐색** (예: 폴더 안의 `*_chunks*.py` 를 찾아 실 행)하도록 바꿔줄게. “자동 탐색 버전”이라고만 말해줘. 

**==> picture [489 x 39] intentionally omitted <==**

괜찮아 그냥니가 쓰는 파일로 만들어도 무방해 내가 정리 하기 편하기위해 이름 바꾸려고 했던거야 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

188/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

16s 동안 생각함 

- 형, 오케이. **내가 쓰는 파일명 그대로** 쓰자. 정리는 **폴더/별명(래퍼)**로 해결하면 깔끔 해. 

## **A단계｜핵심** 

- [사실] 실행에 필요한 건 ** `yt_to_txt_chunks.py` (엔진)**와 `유튜브→TXT_40k_원클 릭.bat` **(버튼)** 두 개면 끝. 이름은 **그대로** 쓰는 게 안전. 

- [사실] **정리 목적** 이면 파일명 대신 **폴더명** 을 바꾸거나, **별명 배치(래퍼)**를 만들어 쓰는 게 안정·유지보수에 유리. 

## **B2단계｜정리 요령(안전하게 깔끔 정리)** 

## **1. 폴더로 정리(추천)** 

예: 

css 코드 복사 `2025-10-24_오토코더_파이썬/ 입력URL_목록.txt 결과/ <제목 [ID]>/TXT/... <제목 [ID]>/TXT_chunks/01.txt, 02.txt...` 

[사실] 폴더명은 마음대로 바꿔도 실행 영향 없음. 

**2. 별명(래퍼)로 부르기** — 원본은 그대로, 겉이름만 바꾸기 

`YouTube-script.bat` 만들어서 원하는 이름으로 실행: 

bat 코드 복사 `@echo off python -m yt_dlp %*` 이후 `yt-dlp` 대신 `YouTube-script ...` 로 사용 가능. [사실] 

## **3. 바꾸면 안 되는 것 / 바꿔도 되는 것** 

바꾸면 안 됨: `python.exe` , `yt-dlp` **실행파일/모듈 이름** . [사실] 바꿔도 됨: **출력 파일/폴더 이름** (TXT, 청크 `01.txt` 등). 스크립트 이름을 바꾼다면, 배치의 호출 줄도 같이 바꾸기: 

bat 코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

189/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python "yt_to_txt_chunks.py" --url "..."   ← 이부분을새파일명으로
```

## **4. 이름 규칙 팁(정렬·재처리 쉬움)** 

청크: `01.txt, 02.txt, 03.txt …` (앞자리 0 채우기) [사실] 

- 금지문자 주의: `\ / : * ? " < > |` 사용 금지, 경로는 큰따옴표로 감싸기 예: 

nginx 코드 복사 

```
python"C:\작업\yt_to_txt_chunks.py" --url "https://youtu.be/..."
```

원하면 **배치가 스크립트 파일명을 자동 탐색** (폴더 안에서 찾아 실행)하도록 바꿔줄 수도 있어. “자동 탐색 버전”이라고 말해주면 바로 만들어 줄게, 형. 

**==> picture [489 x 38] intentionally omitted <==**

## 자동탐색버전이면 어떤 부분이 바뀐거지? 

16s 동안 생각함 

형, “자동탐색 버전”은 **배치가 엔진 스크립트 이름을 하드코딩하지 않고, 폴더에서 자동으 로 찾아 실행** 하는 거야. 그래서 **스크립트 이름을 바꿔도** 잘 돌아가. 

## **A단계｜핵심 변화** 

- [사실] 배치가 **실행할 .py 파일을 자동 검색** (우선순위 규칙) → 이름 바꿔도 OK. [사실] 찾은 엔진 경로를 화면에 보여주고 그걸로 실행. 

- [사실] 나머지 사용 흐름(URL 입력 → TXT 정리 → 4만자 조각)은 동일. 

## **B2단계｜자동탐색 배치(교체용)** 

아래를 메모장에 붙여 `유튜브→TXT_40k_원클릭_자동탐색.bat` 로 저장해서 써. 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title 유튜브 → TXT(정리) → 4만자조각(01/02/03…) 원클릭 - 자동탐색버전
```

```
REM 0) Python 확인
```

```
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

190/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
REM 1) 엔진(.py) 자동탐색 (우선순위: 정확한이름 > 패턴 > 최신 .py)
set "HERE=%~dp0"
```

```
set "SCRIPT="
```

```
REM 1-1) 정확한이름이있으면최우선
```

```
if exist "%HERE%yt_to_txt_chunks.py" set "SCRIPT=%HERE%yt_to_txt_chunks.py"
```

```
REM 1-2) 이름을바꿨더라도, 파일명에 txt/chunk 가들어가면후보로채택
if not defined SCRIPT (
```

```
  for /f "delims=" %%F in ('dir /b /a:-d "%HERE%*txt*chunk*.py" 2^>nul') do (
```

```
    set "SCRIPT=%HERE%%%F"
    goto :found_script
  )
)
```

```
REM 1-3) 그래도못찾으면, 폴더안의가장최근 .py 사용(마지막수단)
if not defined SCRIPT (
```

```
  for /f "delims=" %%F in ('dir /b /a:-d /o:-d "%HERE%*.py" 2^>nul') do (
    set "SCRIPT=%HERE%%%F"
    goto :found_script
  )
)
```

```
:found_script
```

```
if not defined SCRIPT (
```

```
  echo [오류] 실행할파이썬스크립트를찾지못했습니다. 폴더에 *.py 파일이있는지확인하
  pause & exit /b
```

```
)
```

```
echo [정보] 사용엔진: "%SCRIPT%"
```

```
REM (선택) yt-dlp 설치/업데이트시도 — 실패해도계속진행
python -m pip install -U yt-dlp >nul 2>nul
```

```
REM 2) 사용자입력
set "URL="
set /p URL=유튜브영상또는재생목록 URL:
if "%URL%"=="" echo [오류] URL이필요합니다.& pause & exit /b
```

```
set "OUT="
set /p OUT=저장폴더(비우면현재폴더):
if "%OUT%"=="" set "OUT=."
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 191/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```