
## **3.** 필요 패키지 설치 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 100/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

코드 복사 

nginx 

```
pip install pyannote.audio torch
```

## **실행 예시(화자 구분 켠 버전)** 

css 

## 코드 복사 

```
cd "%USERPROFILE%\OneDrive\Desktop\Voice script"
```

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --diarize--vad--strong
```

  

→ 결과 문장 앞에 `[SPEAKER_00]` , `[SPEAKER_01]` 처럼 사람 태그가 붙어. 

## **굳이 화자구분 안 써도 되면** 

지금 당장 급하면 **이것만으로도 충분히 정확** 해: 

css 

코드 복사 

```
python voice_script.py "[...].m4a" --strong--vad
```

(대화만 추출 + 볼륨 보정) 

막히면 화면 캡처 한 장 더 주고, **마지막 10줄 에러** 같이 붙여줘. 바로 거기서 이어서 잡아 줄게. 

**==> picture [489 x 39] intentionally omitted <==**

**==> picture [289 x 143] intentionally omitted <==**

**==> picture [147 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
대시보드가 어디 있어?<br>**----- End of picture text -----**<br>


34s 동안 생각함 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

101/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

형, **대시보드 = OpenAI Platform** 안에 있어. 거기서 **API Keys** 화면만 찾으면 돼. 

## **바로 가는 길(두 가지)** 

**1.** 화면 오른쪽 위 **프로필/톱니(** ⚙ **) → Settings → API keys** 들어가서 **Create new secret key** . OpenAI Platform 

**2.** 주소창에 **API Keys 페이지** 로 바로 들어가기. OpenAI Platform 

## **헷갈릴 때(스크린샷 화면 기준)** 

왼쪽 메뉴는 문서 메뉴고, 실제 키는 **오른쪽 위 계정 메뉴(또는** ⚙ **Settings)** 쪽이야. 거기서 **API keys** 를 찾으면 됨. OpenAI Platform 

## **키 만들고 PC에 저장(윈도우)** 

키를 만들면 한 번만 보이니까 복사해서: 

nginx 코드 복사 

```
setx OPENAI_API_KEY "여기_복사한_API키"
```

## 창 닫고 다시 열고, 확인: 

scss 코드 복사 

```
python -c "import os;print(bool(os.getenv('OPENAI_API_KEY')))"
```

참고: 공식 도움말에도 “API key는 API Keys 페이지에서 찾는다”라고 되어 있어. OpenAI Help Ce… 

필요하면 지금 화면에서 어디를 눌러야 할지 캡처에 표시해서 더 짧게 알려줄게. 

**==> picture [12 x 11] intentionally omitted <==**

**==> picture [11 x 11] intentionally omitted <==**

**출처** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

102/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

**==> picture [193 x 212] intentionally omitted <==**

**==> picture [107 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
뭐를 써야하지?<br>**----- End of picture text -----**<br>


## 형, 저 창은 **API 키 만들 때 붙이는 라벨** 같은 거야. 딱 이렇게만 하면 돼. 

## **뭐라고 쓰면 되나?** 

   - **Owned by** : `You` 그대로 둬. (서비스 계정 필요 없음) 

   - **Name** : 마음대로. 예) `voicescript-whisper` (나중에 구분용) 

   - **Project** : `Default project` 그대로 OK. 

   - **Permissions** : `All` 그대로 둬. (우린 음성→텍스트만 쓸 거라 문제 없음) 

- 이제 **Create secret key** 눌러! 

## **만든 다음 꼭 해야 할 것 (윈도우)** 

**1. 방금 뜬 키를 바로 복사** (다시 못 봄) 

**2.** 명령창에 붙여넣기: 

nginx 

코드 복사 

```
setx OPENAI_API_KEY "여기_복사한_API키"
```

**3. 창 닫고 새로 열기** → 확인: 

nginx 

코드 복사 

```
python -c "import os;print('OK' if os.getenv('OPENAI_API_KEY') else 'NO')"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 103/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

`OK` 나오면 준비 완료. 

## **이제 실행** 

bash 코드 복사 

```
cd"%USERPROFILE%\OneDrive\Desktop\Voice script"
```

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --vad
```

## **팁** 

키는 **절대 공유 금지** . 잘못 노출되면 **대시보드→API keys** 에서 **Revoke(삭제)** 후 새로 만들면 돼. 

이름(Name)은 나중에 키를 여러 개 만들 때 구분만 해주는 “메모”라고 생각하면 편 해. 

**==> picture [489 x 39] intentionally omitted <==**

@echo off 

REM 1) 이 배치는 voice_script.py 와 같은 폴더에 둬야 해 요. 

cd /d "%~dp0" 

REM 2) 파이썬이 py로만 될 때는 아래 줄에서 python 을 py 로 바꿔도 돼요. 

if "%~1"=="" ( 

echo [안내] 음성 파일을 이 창에 "끌어넣거나", 파일명을 입력하세요. 

set /p AUDIO=파일이름을 입력하세요(예: 회의.mp3): python voice_script.py "%AUDIO%" --strong --vad ) else ( 

python voice_script.py --strong --vad %* 

) 

echo. 

echo [완료] 창을 닫으려면 아무 키나 누르세요. 

pause >nul 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 104/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

형, 배치파일 잘 만들었어! 그대로 써도 돼. 

다만 **초보자 보호(오류 덜 나게)** 몇 가지만 더 넣은 “안전버전”을 줄게. (키 확인·파일 확인 · `py` 자동대체) 

## **안전버전 run_voice_script.bat (복붙해서 기존 파일 덮어쓰기)** 

bash 

코드 복사 

```
@echo off
REM === Voice Script 원클릭실행(안전버전) ===
REM 이배치파일은 voice_script.py 와같은폴더에둡니다.
```

```
REM 0) 배치가있는폴더로이동
cd /d "%~dp0"
```

```
REM 1) Python 실행파일자동선택 (python 없으면 py 사용)
where python >nul 2>nul
if errorlevel 1 (
set"_PY=py"
) else (
set"_PY=python"
)
```

```
REM 2) OpenAI 키체크 (없으면안내후종료)
```

```
for /f "usebackq delims=" %%K in (`%_PY% -c "import os;print(os.getenv('OPENAI_AP
""
if"%_OPENAI%"== (
```

```
echo [에러] OPENAI_API_KEY 가없습니다.
echo아래명령으로먼저저장하고창을닫았다가다시여세요:
echo        setx OPENAI_API_KEY "여기_본인_API키"
  pause
exit /b 1
)
```

```
REM 3) 파일입력받기 (드래그앤드롭지원)
""
if"%~1"== (
echo [안내] 음성파일을이창에끌어넣거나, 파일명을직접입력하세요.
set /p AUDIO=파일이름(예: [20230927]제이타운_매매관련01.m4a):
) else (
set"AUDIO=%*"
)
```

## `REM 4) 파일 존재 확인` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

105/175 
