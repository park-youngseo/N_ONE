
```
python voice_script.py "회의.mp3" --strong--vad
```

파일명이 다르면 `"파일이름.mp3"` 만 바꿔 적어. 

- 결과: 같은 폴더에 `회의_transcript.txt` , `회의.srt` , `회의_segments.csv` 

- 화자 구분 원하면: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 74/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

코드 복사 

 

css 

```
python voice_script.py "회의.mp3" --diarize--vad--strong
```

- (사전: `pip install pyannote.audio torch` + `setx HUGGINGFACE_TOKEN "토큰"` ) 

## **2) 실행(더블클릭/끌어넣기 .bat) — 가장 쉬움** 

**같은 폴더** (바탕화면\Voice script)에 아래 내용으로 파일을 만들고 이름을 **run_voice_script.bat** 로 저장: 

bash 

코드 복사 

```
@echo off
REM 1) 이배치는 voice_script.py 와같은폴더에둬야해요.
cd /d "%~dp0"
```

```
REM 2) 파이썬이 py로만될때는아래줄에서 python 을 py 로바꿔도돼요.
```

```
""
if"%~1"== (
```

```
echo [안내] 음성파일을이창에"끌어넣거나", 파일명을입력하세요.
```

```
set /p AUDIO=파일이름을입력하세요(예: 회의.mp3):
```

```
  python voice_script.py "%AUDIO%" --strong --vad
```

```
) else (
  python voice_script.py --strong --vad %*
```

```
)
```

```
echo.
```

```
echo [완료] 창을닫으려면아무키나누르세요.
```

```
pause >nul
```

## **쓰는 법** 

- **방법 A** : `run_voice_script.bat` 을 **더블클릭** → 파일명 입력 → 엔터 

- **방법 B** : 음성 파일을 **배치파일 위로 끌어넣기** (드래그앤드롭) → 자동 실행 

- 여러 파일도 한 번에 가능(여러 개 드래그) 

- `python` 이 안 되면 배치 안의 `python` 을 `py` 로 바꿔서 다시 저장. 

## **3) bat 안의 “(원클릭 파이프라인 v1.1)” 문구는 지워도 돼** 

배치파일은 **그냥 실행 명령** 만 있으면 돼. 

- 설명 문구를 적고 싶다면 `REM` 주석으로만 적어줘. (실행엔 영향 없음) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

75/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **4) 이름 바꾼 것 OK?** 

**파일명 바꾸기 OK** : `stt_pipeline.py` → **voice_script.py** 문제 없음. 

바꾼 뒤엔 실행할 때도 그 **새 이름으로** 불러주기만 하면 돼. 

## **5) 다른 폴더로 옮겨도 돼?** 

- **돼.** 단, 실행할 때 **경로만 맞추기** . 

방법 A: 그 폴더로 **cd** 한 뒤 실행 

- 방법 B: 경로를 **따옴표로 감싸서** 풀경로로 실행 

lua 코드 복사 

`python "C:\Users\psm23\Desktop\Voice script\voice_script.py" "D:\audio\회`   **공백** 이 있으면 항상 `"따옴표"` 로 감싸! 

## **6) 토큰/키 다시 한 번 정리** 

## **(1) HUGGINGFACE_TOKEN = 화자 구분용 (pyannote.audio)** 

**==> picture [481 x 188] intentionally omitted <==**

**----- Start of picture text -----**<br>
어디서: huggingface.co → 로그인 →  Settings → Access Tokens → New token<br>(Read)<br>저장:<br>nginx 코드 복사<br>setx HUGGINGFACE_TOKEN "여기_복사한_토큰"<br>창 닫고 새 창  열기 → 확인:<br>bash 코드 복사<br>**----- End of picture text -----**<br>


```
echo %HUGGINGFACE_TOKEN%
```

## **(2) OPENAI_API_KEY = 자막 변환(Whisper)용** 

- 어디서: OpenAI 계정 → **API Keys** → **Create new secret key** → 복사 저장: 

nginx 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

76/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
setx OPENAI_API_KEY "여기_복사한_API키"
```

**창 닫고 새 창** 열기 → 확인: 

nginx 코드 복사 

```
python -c "import os;print('OK' if os.getenv('OPENAI_API_KEY') else 'NO')"
```

→ `OK` 면 준비 끝. 

## **7) 흔한 막힘 3개(바로 해결)** 

## **1. ModuleNotFoundError: openai** 

nginx 

코드 복사 

```
pip install openai pandas webrtcvad-wheels
```

## **2. ffmpeg not found** 

`ffmpeg -version` 확인, 안 뜨면 PATH 재확인 

## **3. 파일을 못 찾음** 

경로/파일명에 공백→ **"따옴표"** , 

혹은 **bat와 같은 폴더** 로 파일을 옮겨서 실행 

일단 **bat 만들고 드래그앤드롭** 으로 한 번 돌려봐. 

에러 나면 **마지막 10줄** 만 붙여줘—그 지점부터 바로 잡아줄게! 

**==> picture [489 x 39] intentionally omitted <==**

Microsoft Windows [Version 10.0.26100.6899] (c) Microsoft Corporation. All rights reserved. 

C:\Users\psm23>cd "%USERPROFILE%\Desktop\Voice script" 지정된 경로를 찾을 수 없습니다. 

C:\Users\psm23> 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 77/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

오키 형, 지금은 **폴더 경로가 다르거나 폴더가 없는** 상태야. 아래 순서대로 확인→만들기 →이동 하자. (복붙만) 

## **1) 내 사용자 폴더와 데스크톱 경로 확인** 

bash 코드 복사 `echo %USERPROFILE% dir "%USERPROFILE%" dir "%USERPROFILE%\Desktop" dir "%USERPROFILE%\바탕화면"` 

위 네 줄을 **그대로** 치고, 어디가 **있는 폴더** 인지 본다. 

(한국어 윈도우는 화면엔 “바탕 화면”으로 보이지만, 실제 경로는 보통 `Desktop` 이 맞 음. PC마다 다를 수 있어 둘 다 확인하는 거야.) 

## **2) 폴더가 없으면 만들기** 

아래 둘 중 **존재하는 데스크톱 경로** 에 맞춰 한 줄만 실행: 

데스크톱 경로가 `... \Desktop` 일 때: 

arduino 코드 복사 `mkdir "%USERPROFILE%\Desktop\Voice script"` 데스크톱 경로가 `... \바탕화면` 일 때: arduino 코드 복사 `mkdir "%USERPROFILE%\바탕화면\Voice script"` 

## **3) 그 폴더로 이동(공백은 따옴표!)** 

`Desktop` 이 있는 경우: 

bash 코드 복사 `cd "%USERPROFILE%\Desktop\Voice script" 바탕화면` 이 있는 경우: bash 코드 복사 `cd "%USERPROFILE%\바탕화면\Voice script"` 

## **4) 파일 위치 확인** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

78/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

코드 복사 

bash 

```
dir
```

여기에 **voice_script.py** 가 보여야 해. 

안 보이면, 현재 폴더에 파일을 복사하거나, 파일이 있는 경로에서 **풀 경로** 로 실행하면 돼: 

lua 코드 복사 `python "C:\Users\psm23\Desktop\Voice script\voice_script.py" "회의.mp3"` _`--strong -`_   **5) 바로 실행 예시** (지금 폴더가 Voice script 라면) css 코드 복사 

## **5) 바로 실행 예시** 

(지금 폴더가 Voice script 라면) 

```
python voice_script.py "회의.mp3" --strong--vad
```

막히면 위 명령들을 쳤을 때 나온 **결과(특히 dir 결과)**를 붙여줘. 경로를 정확히 잡아줄 게! 

**==> picture [489 x 39] intentionally omitted <==**

Microsoft Windows [Version 10.0.26100.6899] (c) Microsoft Corporation. All rights reserved. C:\Users\psm23>cd "%USERPROFILE%\바탕화면 \Voice script" C:\Users\psm23\바탕화면\Voice script> 

굿! 이제 **정확한 폴더** 에 들어왔어. 바로 실행만 하면 된다. 

## **1) 파일 있는지 확인** 

bash 코드 복사 `dir` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 79/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

목록에 **voice_script.py** 랑 음성파일(예: `회의.mp3` )이 보이면 OK. 

## **2) 바로 실행(소리 보정 + 대화만)** 

css 코드 복사 `python voice_script.py "회의.mp3" --strong --vad` 파일명이 다르면 `"파일이름.mp3"` 로 바꿔 써. 결과는 **같은 폴더** 에 생겨: `회의_transcript.txt` , `회의.srt` , `회의_segments.csv` 