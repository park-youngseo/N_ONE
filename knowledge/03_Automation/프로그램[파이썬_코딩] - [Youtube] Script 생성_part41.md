```
if exist "%HERE%urls.txt" (
```

```
  echo [정보] urls.txt 감지: "%HERE%urls.txt"
```

```
  call :process_file "%HERE%urls.txt"
  echo [완료] 처리종료.
  pause & exit /b
)
```

```
echo [입력] 한줄에한 URL씩붙여넣고, 끝낼때는빈줄로 Enter:
:collect_loop
set "LINE="
set /p LINE=
if "%LINE%"=="" goto :done_collect
call :run_url "%LINE%"
goto :collect_loop
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

205/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
:done_collect
echo [완료] 처리종료.
pause & exit /b
```

```
:process_args
:args_loop
if "%~1"=="" goto :after_args
set "ITEM=%~1"
if exist "%ITEM%" (
  call :process_file "%ITEM%"
) else (
  call :run_url "%ITEM%"
)
shift
goto :args_loop
:after_args
echo [완료] 처리종료.
pause & exit /b
REM ===== 파일에서 URL을읽어처리(줄당 1개, #으로시작하면주석) =====
:process_file
set "FILE=%~1"
if not exist "%FILE%" goto :eof
for /f "usebackq delims=" %%L in ("%FILE%") do (
  set "LINE=%%L"
  if not "!LINE!"=="" (
    echo !LINE! | findstr /b /c:"#">nul
    if errorlevel 1 (
      call :run_url "%%L"
    )
  )
)
goto :eof
REM ===== 단일 URL 처리 =====
:run_url
set "URL=%~1"
if "%URL%"=="" goto :eof
echo.
echo [RUN] python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-si
python "%SCRIPT%" --url "%URL%" --out "%OUT%" --lang %LANG% --chunk-size %CHUNK% %
goto :eof
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 206/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **어떻게 쓰나? (3가지 중 아무거나)** 

- **방법 A | 붙여넣기** : 배치 더블클릭 → **폴더만 입력** → URL들을 여러 줄로 붙여넣고 **빈 줄로 종료** . 

- **방법 B |** `urls.txt` : 배치와 같은 폴더에 `urls.txt` (한 줄에 한 URL) → 더블클릭(폴더 만 입력)하면 **자동 처리** . 

## **방법 C | 인자** : 

arduino 코드 복사 

```
유튜브→TXT_40k_MULTI.bat "https://youtu.be/AAA""https://youtu.be/BBB"
```

또는 `유튜브→TXT_40k_MULTI.bat urls.txt` 

## **결과(항상 동일)** 

css 코드 복사 

```
<저장폴더>/
제목[영상ID]/
    TXT/제목[ID].txt
    TXT_chunks/01.txt, 02.txt, ...
```

[사실] 재생목록 URL이면 **포함된 모든 영상** 에 대해 위 구조가 각각 생성된다. 

원하면 **OUT 기본값** 을 특정 경로로 고정하거나, **동시(병렬) 처리** 가 필요할 때 안전선(최대 동시 수)도 추가해 줄 수 있어. 

**==> picture [489 x 38] intentionally omitted <==**

초등학생도 이해하기 편하게 설명해 

35s 동안 생각함 

## 형, 진짜 쉽게 말해볼게. 

**여러 유튜브 주소를 한 번에** 넣으면, 컴퓨터가 **윗줄부터 차례대로** → **글(TXT)** 만들고 → 

**4만 글자씩 `01.txt, 02.txt…` **로 나눠 줘. [사실] 

## **A단계｜한눈에 이해** 

- [사실] 링크들을 “ **목록** ”으로 모아 두면, 컴퓨터가 **1번 링크 → 2번 링크 → 3번 링크** 순서로 자동으로 처리한다. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 207/341 

25. 11. 30. 오후 1:07 

 

**==> picture [154 x 9] intentionally omitted <==**

**----- Start of picture text -----**<br>
프로그램[파이썬/코딩] - [Youtube] Script 생성<br>**----- End of picture text -----**<br>


- [사실] 결과는 각 영상마다 폴더가 생기고, 그 안에 **TXT(전체 글)** + 

- **TXT_chunks(01.txt, 02.txt…)**가 들어간다. 

- [사실] **언어=한국어(ko)** , **자막 파일(VTT)=자동 삭제** , **조각 크기=4만 글자** 로 고정돼 있어 따로 묻지 않는다. 

## **B2단계｜아주 쉬운 사용법(3가지 중 하나만 하면 끝)** 

**방법 1)** `urls.txt` **로 한꺼번에 (가장 쉬움)** 

**1.** 메모장 열기 → **한 줄에 한 개 링크** 씩 적기 

arduino 코드 복사 `https:` _`//youtu.be/AAAA...`_ `https:` _`//youtu.be/BBBB...`_ `https:` _`//www.youtube.com/playlist?list=PL...`_ 

**2.** 파일 이름을 ** `urls.txt` **로 저장(버튼 파일과 같은 폴더). 

**3. 버튼(배치) 더블클릭** → **저장 폴더만** 적으면 **자동으로 전부** 처리. [사실] 

## **방법 2) 여러 줄 붙여넣기** 

- 버튼 더블클릭 → **저장 폴더** 만 적기 → 창에 **링크들을 여러 줄로 붙여넣고** , 마지막에 **빈 줄(엔터 한번)** 넣으면 시작. [사실] 

## **방법 3) 끌어다 놓기/한 줄 실행** 

- `urls.txt` 파일을 **버튼 아이콘 위로 드래그&드롭** 하거나, 

- 명령창에서 

arduino 코드 복사 

```
(버튼이름).bat "https://youtu.be/AAA""https://youtu.be/BBB"
```

## 처럼 **여러 개를 한 줄에** 넣어도 된다. [사실] 

## **결과 폴더(모든 방법 동일)** 

css 코드 복사 `저장폴더/ 영상제목 [영상ID]/ TXT/영상제목 [ID].txt           ← 전체 글 TXT_chunks/01.txt, 02.txt, ...  ← 4만 글자씩 나눈 조각` 

[사실] 재생목록 링크를 넣으면, 안에 있는 **모든 영상이 각각** 위 구조로 만들어진다. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

208/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

형이 할 일은 딱 **링크만 준비** 하는 것뿐이야. 

나머지(한국어 선택, VTT 삭제, 4만 글자 나누기)는 **자동** 이니까 편하게 써! 

**==> picture [489 x 38] intentionally omitted <==**

1번으로 할께 

## 12s 동안 생각함 

좋아, **방법 1 (urls.txt로 한꺼번에)**로 바로 쓰기 쉽게 안내할게. 

말 그대로 “링크 목록 파일 하나”만 만들면 나머지는 자동이야. 

## **A단계｜한눈에** 

- [사실] 같은 폴더에 `yt_to_txt_chunks.py` **(엔진)** , `유튜브→TXT_40k_MULTI.bat` **(버튼)** , `urls.txt` **(링크 목록)** 세 개가 있으면 끝. 

- [사실] 실행하면 **언어=한국어(ko)** , **VTT 자동삭제** , **TXT 4만 글자씩 `01.txt,` 

- `02.txt…` **로 자동 분할된다. 

## **B2단계｜초간단 순서** 

## **1. 폴더 준비** (권장됨) 

한 폴더에 파일 2개를 넣어둔다: 

`yt_to_txt_chunks.py` , `유튜브→TXT_40k_MULTI.bat` 

**2.** `urls.txt` **만들기** (권장됨) 

   - 메모장 열고 **한 줄에 한 개** 유튜브 주소를 적는다. 재생목록도 가능. 예시(복붙 가능): 

arduino 코드 복사 

```
https://youtu.be/AAAAAAA
```

```
https://youtu.be/BBBBBBB
```

```
https://www.youtube.com/playlist?list=PLxxxxxxxxxxxxxxxx
```

파일 이름을 **urls.txt** 로 저장한다. (배치랑 같은 폴더) 

## **3. 실행** (참고용) 

`유튜브→TXT_40k_MULTI.bat` **더블클릭** → 

- “저장 폴더”만 입력(비우면 현재 폴더 사용) → **자동으로 전부 처리** . 

## **4. 결과 위치** (사실) 

css 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

209/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
저장폴더/
```

```
영상제목[영상ID]/
```

```
    TXT/영상제목[ID].txt           ← 전체글
    TXT_chunks/01.txt, 02.txt, ...  ← 4만글자씩조각
```

여러 URL이면 **위 폴더가 영상마다 하나씩** 생긴다. 

## **5. 작동 팁** (참고용) 

`urls.txt` 는 **한 줄=한 링크** , 빈 줄/ `#` 로 시작하는 줄은 무시. 

- 자막이 없는 영상은 건너뛰거나, `--lang en` 이 필요할 수 있음(지금 배치는 ko 고정). 

- 연령/로그인 제한 영상은 `cookies.txt` 가 필요할 수 있음(원하면 그 버전도 만들 어 줄 수 있음). 

이대로 하면 링크만 모아두고 **한 번에** 돌릴 수 있어. 

완료 후 `TXT_chunks/01.txt` 부터 순서대로 올려주면 내가 **바로 브레인스토밍** 들어갈게. 

**==> picture [489 x 39] intentionally omitted <==**
