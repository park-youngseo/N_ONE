
프로그램[파이썬/코딩] - [Voice] Script 생성 

 

이타운_매매관련01.m4a" --strong --vad python: can't open file 

'C:\\Users\\psm23\\voice_script.py': [Errno 2] No such file or directory 

C:\Users\psm23>python "%USERPROFILE%\바탕화면 \Voice script\voice_script.py" "%USERPROFILE%\ [20230927]제이타운_매매관련01.m4a" --strong --vad python: can't open file 'C:\\Users\\psm23\\바탕화면 \\Voice script\\voice_script.py': [Errno 2] No such file or directory 

C:\Users\psm23>cd "%USERPROFILE%\바탕화면 \Voice script" 

C:\Users\psm23\바탕화면\Voice script>python voice_script.py "%USERPROFILE%\[20230927]제이타운_ 매매관련01.m4a" --strong --vad 

python: can't open file 'C:\\Users\\psm23\\바탕화면 \\Voice script\\voice_script.py': [Errno 2] No such file or directory 

C:\Users\psm23\바탕화면\Voice script>dir "%USERPROFILE%\[20230927]*.m4a" C 드라이브의 볼륨에는 이름이 없습니다. 볼륨 일련 번호: CC77-8F78 

C:\Users\psm23 디렉터리 

파일을 찾을 수 없습니다. 

C:\Users\psm23\바탕화면\Voice script>python "%USERPROFILE%\바탕화면\Voice 

script\voice_script.py" "D:\녹음\[20230927]제 이타운_ 매매관련01.m4a" --strong --vad 

python: can't open file 'C:\\Users\\psm23\\바탕화면 \\Voice script\\voice_script.py': [Errno 2] No such file or directory 

C:\Users\psm23\바탕화면\Voice script> 

형, 두 가지가 문제야: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

89/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

**1. 이 폴더에** `voice_script.py` **가 아직 없음** 

**2. 음성 파일 경로/철자** 가 헷갈림(“제 이타운”처럼 띄어쓰기 들어감) 

아래 순서대로 하면 바로 해결돼. 

## **1) 스크립트 파일 진짜 만들기(확실한 방법)** 

지금 창이 `C:\Users\psm23\바탕화면\Voice script>` 맞지? 거기서: 

**1. 메모장으로 새 파일 열기** 

nginx 코드 복사 `notepad voice_script.py` 뜬 메모장에 내가 준 **voice_script.py 코드 전체** 를 붙여넣고 → **파일 → 저장** (이름은 자동으로 `voice_script.py` , 위치는 지금 폴더) 저장 후, 파일이 생겼는지 확인: bash 코드 복사 `dir voice_script.py` 스크립트가 정상인지 빠르게 점검: nginx 코드 복사 `python voice_script.py -h` 

**2.** 뜬 메모장에 내가 준 **voice_script.py 코드 전체** 를 붙여넣고 → 

**3.** 저장 후, 파일이 생겼는지 확인: 

**4.** 스크립트가 정상인지 빠르게 점검: 

## → 옵션 도움말이 보이면 **정상 저장 완료** 야. 

만약 메모장이 “새 파일을 만들까요?”라고 묻는다면 **예** 눌러. 

## **2) 음성 파일이 어디 있는지 먼저 찾자** 

방금 보니까 `C:\Users\psm23\[20230927]…` 에도 없고, 마지막에 적은 경로는 

`D:\녹음\[20230927]제 이타운_…` 처럼 **띄어쓰기** 가 섞였어(오타일 가능성 높음). 

## **D: 드라이브 안에서 파일 찾기(가장 확실)** 

bash 

코드 복사 

```
dir /s /b "D:\*[20230927]*제이타운*.m4a"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

90/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

결과로 **정확한 전체 경로** 가 한 줄로 나와. 

- (예: `D:\녹음\[20230927]제이타운_매매관련01.m4a` ) 

만약 폴더 이름이 **‘녹음’**이 확실하면 먼저 거기로 가서: 

bash 코드 복사 `dir "D:\녹음\[20230927]*.m4a"` C: 안에 있을 수도 있으니(필요하면): bash 코드 복사 

- C: 안에 있을 수도 있으니(필요하면): 

```
dir /s /b "C:\*[20230927]*제이타운*.m4a"
```

## **3) 찾은 “정확한 경로”로 실행** 

예를 들어 찾은 경로가 

`D:\녹음\[20230927]제이타운_매매관련01.m4a` 라면: 

lua 코드 복사 

```
python voice_script.py "D:\녹음\[20230927]제이타운_매매관련01.m4a"--strong --vad
```

  

## **흔한 실수 3개** 

- **따옴표 빼먹음** → 대괄호/공백이 있으면 꼭 `"` 로 감싸기 

- **오탈자** (제이타운 ↔ 제 이타운) → 철자 1글자만 달라도 “파일 없음” 

- 현재 폴더 착각 → 우리는 이미 `Voice script` 폴더에 있음(OK) 

## **4) 그래도 헷갈리면, 드래그앤드롭 방식(.bat)** 

같은 폴더에 **run_voice_script.bat** 만들고 아래 내용 저장: 

bash 코드 복사 `@echo off cd /d "%~dp0" "" if "%~1"== ( echo [안내] 음성 파일을 이 창에 끌어넣거나 파일명을 입력하세요. set /p AUDIO=파일이름(예: 회의.mp3):` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

91/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
  python voice_script.py "%AUDIO%" --strong --vad
```

```
) else (
  python voice_script.py --strong --vad %*
)
echo.
echo [완료] 아무키나누르면닫혀요.
pause >nul
```

이제 **음성 파일을 이 배치 아이콘 위로 끌어넣기** 만 하면 자동으로 실행돼. 

## **5) 실행이 되면 폴더에 생기는 것** 

`…_transcript.txt` (텍스트) 

- `… .srt` (자막) 

- `…_segments.csv` (시간·문장 표) 

## **지금 바로 할 일(복붙용 체크리스트)** 

**1.** `notepad voice_script.py` → 코드 붙여넣고 저장 

**2.** `dir voice_script.py` → 파일 보이는지 확인 

**3.** `python voice_script.py -h` → 도움말 뜨면 OK 

**4.** `dir /s /b "D:\*[20230927]*제이타운*.m4a"` → 정확한 경로 확인 

**5.** `python voice_script.py "정확한_경로.m4a" --strong --vad` 

나온 결과(특히 에러 **마지막 10줄** )를 붙여주면, 그 지점부터 바로 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

씨발 왜 컴퓨터가 C:\Users\psm23\OneDrive\Desktop 경로가 이렇게 되어 있네 

19s 동안 생각함 

괜찮아 형—그건 **원드라이브가 바탕화면을 ‘OneDrive\Desktop’으로 옮겨놓은** 거라서 그래. 경로만 제대로 잡아주면 끝! 

## **바로 해결 (복붙 순서)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

92/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **1) 진짜 바탕화면 폴더 열기** 

nginx 

코드 복사 

```
explorer"%USERPROFILE%\OneDrive\Desktop"
```

열린 창에서 **Voice script** 폴더가 없으면 만들어줘. 

## **2) 스크립트 파일 위치 맞추기** 

이미 만든 `voice_script.py` 가 예전 위치에 있으면(비어있을 수도 있음) 아래로 **이동** : 

powershell 코드 복사 

```
mkdir "%USERPROFILE%\OneDrive\Desktop\Voice script"
move"%USERPROFILE%\바탕화면\Voice script\voice_script.py""%USERPROFILE%\OneDrive
```

  파일이 아직 없다면, 새로 만들기: nginx 코드 복사 `notepad "%USERPROFILE%\OneDrive\Desktop\Voice script\voice_script.py"` → 메모장에 내가 준 코드 전부 붙여넣고 저장(UTF-8). **3) 그 폴더로 이동해서 실행** bash 코드 복사 `cd "%USERPROFILE%\OneDrive\Desktop\Voice script" python voice_script.py "D:\녹음\[20230927]제이타운_매매관련01.m4a" --strong --vad`   

→ 메모장에 내가 준 코드 전부 붙여넣고 저장(UTF-8). 

## **3) 그 폴더로 이동해서 실행** 

오디오 파일 경로는 **정확히** 적어야 해. 모르겠으면 먼저 찾기: 

bash 코드 복사 

```
dir /s /b "D:\*[20230927]*제이타운*.m4a"
```

## **자주 틀리는 포인트(5초 점검)** 

**경로에 공백/대괄호** 있으면 반드시 **따옴표** 로 감싸기. 지금 바탕화면은 여기임: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

93/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

복사 

bash 

 

```
echo %USERPROFILE%
echo %OneDrive%
dir"%OneDrive%\Desktop"
```

## 헷갈리면 아예 폴더명을 **공백 없는** 걸로 바꿔도 좋아: 

powershell 코드 복사 

```
ren"%USERPROFILE%\OneDrive\Desktop\Voice script" VoiceScript
cd"%USERPROFILE%\OneDrive\Desktop\VoiceScript"
python voice_script.py "..."--strong--vad
```

## **(선택) 드래그앤드롭 배치도 원드라이브 경로로** 

`%USERPROFILE%\OneDrive\Desktop\VoiceScript` 안에 `run_voice_script.bat` 만들고: 

bash 코드 복사 

```
@echo off
cd /d "%~dp0"
""
if"%~1"== (
echo [안내] 음성파일을이창에끌어넣거나파일명을입력하세요.
set /p AUDIO=파일이름(예: 회의.mp3):
  python voice_script.py "%AUDIO%" --strong --vad
) else (
  python voice_script.py --strong --vad %*
)
echo.
echo [완료] 아무키나누르면닫혀요.
pause >nul
```

## 이제 **OneDrive\Desktop 경로** 로만 생각하면 돼. 

위 순서대로 했는데도 막히면, 방금 친 명령과 **마지막 10줄 에러** 를 그대로 붙여줘—그 지 점부터 바로 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**