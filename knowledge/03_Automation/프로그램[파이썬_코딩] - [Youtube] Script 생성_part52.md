pause
```

## **사용 순서 (초등학생 버전)** 

**1. 영상별 TXT** 가 만들어질 폴더를 정해(예: `D:\yotube_script` ). 

**2.** `Script_txt.py` 와 방금 만든 `Script_txt.bat` 를 아무 곳에 두고, `Script_txt.bat` **더 블클릭** . 

**3.** 물어보면 **정리할 폴더 경로** 를 입력 → 엔터. 

   - 이때 `urls.txt` 가 없으면 **자동으로 만들어** 줘. “Y”를 누르면 메모장이 열림(필요 한 URL을 한 줄씩 넣고 저장). 

**4.** 나머지는 자동: **모으기 → 합치기(라벨X) → 4만자 분할 → 원본 삭제** 

**5.** 결과 확인: `정리한_폴더\FINAL_TXT\폴더이름_01.txt, _02.txt …` 만 남아 있다. 

- 원하면 `urls.txt` 를 **반드시 채웠을 때만 실행** 하도록 바꾸거나, **분할 글자수(기본 40,000)** 를 바꾸는 옵션도 더해줄 수 있어. 

**==> picture [489 x 39] intentionally omitted <==**

기존처럼 플레이리스트의 유튜브 영상또한 모두 가져올수 있는거지? 

맞아, 형. **플레이리스트(URL 하나)** 넣으면 그 안의 **모든 영상** 도 그대로 처리돼. 

## **한눈에** 

- [사실] **영상→TXT 뽑기 단계** 에서 플레이리스트 주소를 넣으면, **각 영상별로 TXT** 가 만들어져. 

- [사실] 그다음 **Script_txt.bat** (정리 버튼)은, 만들어진 TXT들을 **모으고→합치고→4만 자 분할** 해서 최종본만 남겨. 

- [사실] 즉, **플레이리스트 전체 → 영상별 TXT 생성 → 최종 TXT(01/02/03…)**까지 한 흐름으로 된다. 

## **어떻게 넣나 (두 가지 중 편한 걸로)** 

## **1. urls.txt에 플레이리스트 주소 한 줄로** 

arduino 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

258/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
https://www.youtube.com/playlist?list=PLxxxxxxxxxxxxxxxx
```

- `유튜브→TXT_SIMPLE_MULTI.bat` 더블클릭 → 폴더 물으면 입력 → 플레이리스트 

- 안의 **모든 영상 TXT 자동 생성** 

끝났으면 `Script_txt.bat` 더블클릭 → **정리 완료** 

## **2. 배치가 주소 직접 물을 때** 

- `유튜브→TXT_SIMPLE_MULTI.bat` 실행 → 주소 입력 칸에 **플레이리스트 URL** 그대로 

- 붙여넣기 → 엔터 

생성 끝나면 `Script_txt.bat` 실행 

## **알아두면 좋은 점** 

- [사실] 플레이리스트 안에 **자막이 없는 영상** 은 건너뛸 수 있어. (그 영상은 TXT가 안 생김) 

- [사실] **비공개/연령제한** 영상은 실패할 수 있어(필요하면 `cookies.txt` 로 해결 가능). 

- [사실] 정리 단계(Script_txt.bat)는 **영상별 TXT만 있으면** 출처가 뭐든 상관없이 잘 합 치고 분할해. 

필요하면 지금 쓰는 배치/스크립트에서 **플레이리스트 전용 상태표시(총 N개, 진행률)** 도 추가해줄 수 있어. 

**==> picture [489 x 38] intentionally omitted <==**

씨발 열받게 하네 

형, 미안해. 많이 답답했지. 지금 바로 되는 **최소 절차** 만 딱 정리해줄게. 이대로만 해봐. **A단계｜핵심(3줄)** 

- [사실] **플레이리스트 URL** 한 줄 넣은 `urls.txt` 만든다. 

- [사실] `유튜브→TXT_SIMPLE_MULTI.bat` 실행 → 영상별 **TXT 생성** . 

- [사실] `Script_txt.bat` 실행 → **모으기→합치기→4만자 분할** , 최종본만 남김. 

## **B1｜초간단 체크리스트** 

**1.** 파이썬·도구 확인 (한 번만) 

nginx 코드 복사 

```
python -V
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 259/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m pip install -U yt-dlp
```

**2.** 작업폴더(예: `D:\yotube_script` )에 파일 4개가 있어야 해 

코드 복사 

```
yt_to_txt_simple.py
유튜브→TXT_SIMPLE_MULTI.bat
Script_txt.py
Script_txt.bat
```

**3.** `D:\yotube_script\urls.txt` 만들기 (플레이리스트 1줄만 적어도 됨) 

arduino 코드 복사 

```
https://www.youtube.com/playlist?list=PLxxxxxxxxxxxxxxxx
```

## **B2｜실행 순서 (진짜 2번 클릭)** 

① **유튜브→TXT_SIMPLE_MULTI.bat** 더블클릭 → (필요 시 폴더 물으면 

`D:\yotube_script` ) → **영상별 TXT 생성** 끝. 

② **Script_txt.bat** 더블클릭 → **정리 폴더 물으면** `D:\yotube_script` 입력 → **최종본만 남 김** . 

## 완료 후 폴더: 

makefile 

코드 복사 

```
D:\yotube_script\FINAL_TXT\
```

```
  yotube_script_01.txt
  yotube_script_02.txt
  yotube_script_03.txt
  ...
```

## **급할 때 트러블슈팅(바로 해결)** 

`python 인식 안 됨` → 파이썬 재설치 시 **Add to PATH** 체크. 

- `자막 없음/429` → 그 영상은 건너뜀(나중에 다시), 플레이리스트의 다른 영상은 계속 

- 진행. 

- `권한/경로 오류` → 경로에 **큰따옴표** 사용: `"D:\yotube_script"` 

- `정리했는데 파일 안 생김` → 영상별 `\TXT\*.txt` 가 있는지 먼저 확인 후 `Script_txt.bat` 다시 실행. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

260/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

형, 지금은 **이 2클릭 루트** 가 제일 안정적이야. 그대로 따라 하고, 어디서 막혔는지 **딱 한 줄** 만 알려줘. 거기부터 바로 풀어줄게. 

**==> picture [489 x 38] intentionally omitted <==**

야 니가 자꾸 파일명을 바꾸면서 다른걸 만드니까 헷갈리 잖아 

30s 동안 생각함 

형, 미안. **이름 바꾸지 말고** 딱 원래대로 고정하자. 아래 그대로만 하면 돼. 

## **A단계｜전체 개요** 

- [사실] 실행 파일명은 `finalize_txt_40k.py` (정리 엔진) + `TXT_정리_한방.bat` (버튼)로 **고정** . 

- [사실] 버튼을 누르면 **정리할 폴더 경로** 를 물어보고, 그 폴더에 `urls.txt` **가 없으면 자동 생성** 후 실행. 

- [사실] 동작: **모으기 → 합치기(라벨 없음) → 4만자 분할(폴더이름_01/02/03…) → 영상별 TXT/합본 삭제** . 

## **B1｜핵심(지금 당장 할 일 3개)** 

**1.** `yotube_script` 폴더에 두 파일 이름을 **고정** 

```
finalize_txt_40k.py
```

```
TXT_정리_한방.bat
```

(다른 이름 쓰지 말기) 

**2.** 아래 코드로 **배치 내용 교체** (폴더 물어보고 `urls.txt` 자동 생성). 

**3.** 아래 코드로 **파이썬 스크립트** (정리 엔진) 내용을 맞춰두기. 

## **1) 버튼:** `TXT_정리_한방.bat` **(이 내용으로 저장)** 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal
```

```
title TXT 정리한방 — 폴더지정 → urls.txt 생성 → 모으기·합치기·4만자분할·원본삭제
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

261/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
REM 0) Python 확인
where python >nul 2>nul || (
  echo [오류] Python이필요합니다. 설치후다시실행하세요.
  pause & exit /b
)
REM 1) finalize_txt_40k.py 확인
set "HERE=%~dp0"
if not exist "%HERE%finalize_txt_40k.py" (
  echo [오류] finalize_txt_40k.py 가이배치와같은폴더에없습니다.
  echo 같은폴더에두고다시실행하세요.
  pause & exit /b
)
```

```
REM 2) 정리할폴더경로입력(비우면현재폴더)
set "ROOT="
set /p ROOT=정리할폴더경로(예: C:\work\yotube_script):
if "%ROOT%"=="" set "ROOT=%HERE%"
set "ROOT=%ROOT:"=%"
```

```
if not exist "%ROOT%" (
  echo [안내] 폴더가없어새로만듭니다: "%ROOT%"
  md "%ROOT%" >nul 2>nul
)
```

```
REM 3) urls.txt 자동생성(없으면)
if not exist "%ROOT%\urls.txt" (
  echo [안내] "%ROOT%\urls.txt" 파일이없어새로만듭니다.
```

```
  > "%ROOT%\urls.txt" echo # 한줄에한개의유튜브링크를적으세요. (재생목록도가능)
  >> "%ROOT%\urls.txt" echo # 예시:
  >> "%ROOT%\urls.txt" echo # https://youtu.be/AAAAAAAAAAA
  >> "%ROOT%\urls.txt" echo # https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX
  >> "%ROOT%\urls.txt" echo #
  >> "%ROOT%\urls.txt" echo # 비워두면이름순으로합쳐집니다.
  echo [안내] 지금 urls.txt를편집하시겠습니까? (Y/N, 기본 N)
  set "ANS="
  set /p ANS=입력:
  if /I "%ANS%"=="Y" start notepad "%ROOT%\urls.txt"
)
REM 4) 정리실행 (urls.txt 유무와상관없이전달해도안전)
echo.
echo [실행] python "finalize_txt_40k.py" --root "%ROOT%" --urls "%ROOT%\urls.txt"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 262/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 
