for /f "usebackq delims=" %%H in (`%_PY% -c "import os;print(os.getenv('HUGGINGFAC
set "DIAR_FLAG="
```

```
if not "%HF_TOKEN%"=="" (
  choice /M "화자구분(pyannote)도사용할까요?" /C YN
  if errorlevel 2 (
```

```
  ) else (
```

```
모델접근권한간이체크 (네트워크필요)
```

```
    curl -s -H "Authorization: Bearer %HF_TOKEN%" https://huggingface.co/api/mode
    if errorlevel 1 (
```

```
      echo [알림] pyannote 모델접근동의가필요합니다. 브라우저에서동의후다시시도
      echo         (모델페이지에서 "Agree and access" 클릭)
      set "DIAR_FLAG="
    ) else (
      set "DIAR_FLAG=--diarize"
    )
  )
)
```

```
REM 5) 필수패키지자동설치(없는것만)
echo [체크] 파이썬패키지확인/설치중...
%_PY% -m pip show openai >nul 2>nul || %_PY% -m pip install -q openai
%_PY% -m pip show pandas >nul 2>nul || %_PY% -m pip install -q pandas
%_PY% -m pip show webrtcvad-wheels >nul 2>nul || %_PY% -m pip install -q webrtcvad
```

```
if "%DIAR_FLAG%"=="--diarize" (
```

`%_PY% -m pip show huggingface_hub >nul 2>nul || %_PY% -m pip install -q huggingf` https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 110/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
  %_PY% -m pip show pyannote.audio >nul 2>nul || %_PY% -m pip install -q pyannote
  %_PY% -m pip show torch >nul 2>nul || %_PY% -m pip install -q torch --index-url
)
REM 6) 입력파일받기 (드래그앤드롭 or 수동입력)
if "%~1"=="" (
  echo.
  echo [안내] 음성파일을이창에끌어넣거나, 파일명을입력하세요.
  set /p AUDIO=파일경로(예: "[20230927]제이타운_매매관련01.m4a") :
) else (
  set "AUDIO=%*"
)
REM 7) 파일존재확인
if not exist "%AUDIO%" (
  echo [에러] 파일을찾을수없습니다: "%AUDIO%"
  echo       파일경로와철자(공백/대괄호는따옴표필수)를확인하세요.
  pause
  exit /b 1
)
REM 8) 실행옵션(기본: 소리보정+대화만)
set "BASE_OPTS=--strong --vad"
```

```
echo.
echo [실행] %_PY% voice_script.py "%AUDIO%" %BASE_OPTS% %DIAR_FLAG%
%_PY% voice_script.py "%AUDIO%" %BASE_OPTS% %DIAR_FLAG%
set "RC=%ERRORLEVEL%"
```

```
echo.
if "%RC%"=="0" (
  echo [완료] 같은폴더에 TXT / SRT / CSV가생성되었습니다.
) else (
  echo [실패] 오류코드: %RC%
)
echo 아무키나누르면닫힙니다.
pause >nul
endlocal
```

## **이 배치가 해주는 것** 

**FFmpeg·Python·키** 있는지 확인 → 없으면 바로 안내 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

111/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

- **HUGGINGFACE_TOKEN** 있으면 모델 접근도 **자동 체크** 하고 `--diarize` 옵션 **자동 ON** 

- 필요한 **pip 패키지 자동 설치** (없는 것만) 

- 파일 **드래그앤드롭/직접입력** 모두 지원 

기본은 `--strong --vad` (대화만 + 작은 소리 보정) 

## **3) 최종 사용법** 

**1.** `voice_script.py` 와 `run_voice_all.bat` 를 **같은 폴더** 

   - `%USERPROFILE%\OneDrive\Desktop\Voice script` 에 둔다. 

**2. 배치 더블클릭** → (원하면 파일 드래그) → 엔터. 

**3.** 끝나면 **같은 폴더** 에 

`…_transcript.txt` , `….srt` , `…_segments.csv` 생김. 

막히면 방금 실행 결과의 **마지막 10줄** 만 복붙해줘. 그 자리에서 바로 해결할게! 

**==> picture [489 x 39] intentionally omitted <==**

아무것도 생성되지 않아 

형, “아무것도 안 생김”이면 보통 **중간에 에러 나고 멈춘** 거야. 

딱 4단계만 따라해서 **원인 잡고 바로 생성** 까지 가자. 

## **0) 폴더 맞게 열기** 

bash 코드 복사 

```
cd"%USERPROFILE%\OneDrive\Desktop\Voice script"
dir
```

- (여기 `voice_script.py` 랑 음성파일이 보여야 해) 

## **1) 제일 단순하게 먼저 돌려보기 (VAD 끄고)** 

css 

코드 복사 

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--segment60
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

112/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

- 이건 **대화만(VAD) 없이** 60초씩 잘라 전사 → 결과 3개(TXT/SRT/CSV)가 **같은 폴더** 에 생겨야 해. 

이걸로 되면, 다음에 `--vad` 를 다시 켠다: 

css 

코드 복사 

```
python voice_script.py "[...].m4a" --strong--vad
```

## **2) 에러 내용 저장해서 보기 (무조건 원인 잡힘)** 

아래 한 줄을 그대로 실행: 

css 

코드 복사 

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--vad1>ok.txt
notepad err.txt
```

  **err.txt** 에 마지막 줄(10줄 정도)만 나에게 붙여줘. 일반적으로 여기서 나는 에러 TOP3 

**1.** `OPENAI_API_KEY` 없음/인식 안 됨 → 키 다시 넣고 **창 닫았다 열기** 

**2.** `ffmpeg` 안 보임 → PATH 재확인(형은 이미 OK일 확률 큼) 

**3.** 네트워크/방화벽(회사망)로 API 호출 실패 → 집/핫스팟에서 시도 

## **3) 키/도구 3초 점검** 

nginx 

코드 복사 

```
python -c "import os;print('OPENAI OK' if os.getenv('OPENAI_API_KEY') else 'NO')"
ffmpeg -version
```

  

`OPENAI OK` 가 나오고, ffmpeg 버전이 보이면 기본 준비는 끝. 

(화자구분도 쓸 거면) 

powershell 코드 복사 

```
python -c"import os;print('HF OK' if os.getenv('HUGGINGFACE_TOKEN') else 'SKIP')
curl-s-H"Authorization: Bearer %HUGGINGFACE_TOKEN%" https://huggingface.co/api/
```

 

 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 113/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

줄이 나오면 **모델 접근 동의 OK** . 아니면 브라우저에서 “Agree and access” 다시 클 릭. 

## **4) 배치(원클릭)로 전부 자동 실행** 

`run_voice_all.bat` 을 이렇게 고쳐 저장(덮어쓰기)하면, 

**키 체크/모델 동의 체크/필수 pip 설치/실행** 을 한 번에 처리해줘. 

bash 

코드 복사 

```
@echo off
setlocal ENABLEDELAYEDEXPANSION
cd /d "%~dp0"
```

```
where python >nul 2>nul && (set"_PY=python") || (set"_PY=py")
```

```
ffmpeg -version >nul 2>nul || (echo [에러] FFmpeg 미설치. winget install --id=FFmp
```

```
for /f "usebackq delims=" %%K in (`%_PY% -c "import os;print(os.getenv('OPENAI_AP
""""
if"%OPENAI_KEY%"== (echo [에러] OPENAI_API_KEY 없음. setx OPENAI_API_KEY 키후
```

```
for /f "usebackq delims=" %%H in (`%_PY% -c "import os;print(os.getenv('HUGGINGFAC
set"DIAR_FLAG="
```

```
""
if not "%HF_TOKEN%"== (
```

```
  curl -s -H "Authorization: Bearer %HF_TOKEN%" https://huggingface.co/api/models/
)
```

```
echo [체크] 필요한패키지설치/업데이트중...
```

```
%_PY% -m pip show openai >nul 2>nul || %_PY% -m pip install -q openai
```

```
%_PY% -m pip show pandas >nul 2>nul || %_PY% -m pip install -q pandas
```

```