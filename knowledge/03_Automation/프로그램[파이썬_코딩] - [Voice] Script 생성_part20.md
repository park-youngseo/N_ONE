25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
if not exist "%AUDIO%" (
echo [에러] 파일을찾을수없습니다:
echo"%AUDIO%"
echo파일을이폴더로옮기거나, 전체경로를다시입력하세요.
  pause
exit /b 1
)
```

```
REM 5) 기본실행(작은소리보정 + 대화만추출)
echo.
echo [실행] %_PY% voice_script.py "%AUDIO%" --strong --vad
%_PY% voice_script.py "%AUDIO%" --strong --vad
set"_RC=%ERRORLEVEL%"
```

```
echo.
```

```
if"%_RC%"=="0" (
echo [완료] 같은폴더에 TXT / SRT / CSV 가생성되었습니다.
) else (
echo [실패] 오류코드: %_RC%
)
echo아무키나누르면닫힙니다.
pause >nul
```

## **왜 더 안전해?** 

`python` 이 없으면 자동으로 `py` 사용. 

- **OPENAI_API_KEY** 없는 상태를 미리 잡아줘서 헛수고 방지. 

- 파일 경로를 **반드시 확인** 해서 “파일 없음” 실수를 줄임. 

- 드래그앤드롭/직접입력 둘 다 지원. 

## **쓰는 법 (아주 간단)** 

**1.** 이 배치파일과 `voice_script.py` 를 **같은 폴더** (예: 

   - `%USERPROFILE%\OneDrive\Desktop\Voice script` )에 둬. 

**2.** 배치 **더블클릭** → 파일명 입력(또는 **음성파일을 배치 아이콘 위로 끌어넣기** ). 

**3.** 끝나면 같은 폴더에 `…_transcript.txt` , `….srt` , `…_segments.csv` 가 생겨. 

## **화자 구분까지 쓰고 싶으면** 

## 한 번만: 

nginx 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

106/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
pip install pyannote.audio torch
```

```
setx HUGGINGFACE_TOKEN "본인_토큰"
```

실행할 때는 배치 창에서 이 한 줄로: 

css 코드 복사 

```
python voice_script.py "파일.m4a" --diarize--vad--strong
```

(배치 기본은 `--strong --vad` 라서 화자 구분이 필요할 땐 이렇게 수동으로 쳐줘.) 

## **마지막 체크 3개** 

`ffmpeg -version` 숫자 보이면 OK (형은 이미 OK). 

`" setx OPENAI_API_KEY "키` 후엔 **창을 닫고 다시 열기** . 

- 파일/폴더 경로에 공백·대괄호 있으면 **항상 따옴표** 로 감싸기. 

원하면 **“화자 구분 ON” 버전 배치** 도 따로 만들어 줄게. 말만 해! 

**==> picture [489 x 40] intentionally omitted <==**

나는 토큰과 api 키 입력까지 전부 완료했어 

형, 좋아! **토큰/키까지 끝** 났으면 이제 바로 돌리면 된다. 딱 이대로 해봐. 

## **1) 폴더로 이동** 

bash 

코드 복사 

```
cd"%USERPROFILE%\OneDrive\Desktop\Voice script"
```

## **2) 키가 잡히는지 3초 점검** 

nginx 

코드 복사 

```
python -c "import os;print('OPENAI OK' if os.getenv('OPENAI_API_KEY') else 'NO')"
python -c "import os;print('HF OK' if os.getenv('HUGGINGFACE_TOKEN') else 'SKIP')
ffmpeg -version
```

  

`OPENAI OK` (필수), `HF OK` (화자구분 쓸 때), ffmpeg 버전이 보이면 준비 완료. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

107/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **3) 실행 (소리 보정 + 대화만)** 

css 

코드 복사 

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--vad
```

끝나면 같은 폴더에: 

`…_transcript.txt` (텍스트) 

- `….srt` (자막) 

- `…_segments.csv` (시간·문장 표) 

## **4) 화자 구분도 쓰려면** 

(모델 접근 동의까지 했다는 가정) 

css 

코드 복사 

```
pip install pyannote.audio torch
```

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --diarize--vad--strong
```

  

처음 한 번은 **모델 다운로드** 로 시간이 조금 걸릴 수 있어. 

## **5) 배치로 더 쉽게** 

`run_voice_script.bat` 더블클릭 → 파일명 입력(또는 드래그앤드롭)하면 자동으로 `--strong --vad` 모드로 돌아간다. 

## **에러 나면** 

마지막 **10줄** 만 붙여줘. 거기서 바로 잡아줄게. 

**==> picture [489 x 40] intentionally omitted <==**

토큰부터 마지막 까지 설치 완료 했고, pip install pyannote.audio torch 모델접근 동의를 확인할수 있는 방 법은? 그리고 이 모든걸 적용한 배치 파일을 만들어줘 

형, 굿! 이제 **토큰/키/패키지/동의** 가 제대로 먹었는지 확인→원클릭 배치로 실행까지 한 번에 가자. 

## **1) pyannote “모델 접근 동의” 확인하는 가장 쉬운 법** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

108/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **A. 명령 한 줄(HTTP로 바로 확인)** 

bat 

코드 복사 

```
curl -H "Authorization: Bearer %HUGGINGFACE_TOKEN%" https://huggingface.co/api/mo
```

  화면에 JSON이 뜨고 `"id":"pyannote/speaker-diarization"` 같은 정보가 보이면 **접** 

## **근 OK** 

`401/403` 이나 “not authorized/agree” 비슷한 문구가 보이면 → 모델 페이지에서 **[Agree and access]** 다시 눌러. 

## **B. 파이썬으로 확인(허깅페이스 SDK)** 

bat 코드 복사 

`python -c "import os; from huggingface_hub import HfApi; api=HfApi(token=os.geten`   

출력에 `HF ACCESS OK` 나오면 **동의+토큰 모두 정상** 이야. (필요시 먼저 `pip install huggingface_hub` ) 

## **2) “모든 옵션 자동 처리” 원클릭 배치파일** 

아래를 `run_voice_all.bat` 라는 이름으로 

`%USERPROFILE%\OneDrive\Desktop\Voice script` 폴더에 저장해. 

(파일 아이콘 더블클릭만 하면 됨. 드래그앤드롭도 지원) 

bat 

코드 복사 

```
@echo off
REM === Voice Script: 원클릭전체파이프라인(설치/체크/실행) ===
setlocal ENABLEDELAYEDEXPANSION
```

```
REM 0) 배치위치로이동
cd /d "%~dp0"
REM 1) Python 선택 (python 없으면 py)
where python >nul 2>nul && (set "_PY=python") || (set "_PY=py")
```

```
REM 2) FFmpeg 확인
ffmpeg -version >nul 2>nul
if errorlevel 1 (
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 109/175 

25. 11. 30. 오후 1:03 

 

프로그램[파이썬/코딩] - [Voice] Script 생성 

```
  echo [에러] FFmpeg 이보이지않습니다. 먼저설치하고다시실행하세요.
  echo       (명령)  winget install --id=FFmpeg.FFmpeg -e
  pause
  exit /b 1
)
```

```
REM 3) OpenAI 키확인(필수)
for /f "usebackq delims=" %%K in (`%_PY% -c "import os;print(os.getenv('OPENAI_AP
if "%OPENAI_KEY%"=="" (
```

```
  echo [에러] OPENAI_API_KEY 가설정되어있지않습니다.
```

```
  echo       명령창에서:  setx OPENAI_API_KEY "여기_본인_API키"  후, 창을닫았다가다
  pause
  exit /b 1
)
```

```
REM 4) HuggingFace 토큰확인(있으면 diarize 사용제안)