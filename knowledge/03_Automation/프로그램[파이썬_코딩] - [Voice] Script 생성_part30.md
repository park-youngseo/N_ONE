
162/175 

25. 11. 30. 오후 1:03 

 

프로그램[파이썬/코딩] - [Voice] Script 생성 

Requirement already satisfied: fsspec>=0.8.5 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (2025.9.0) Requirement already satisfied: mpmath<1.4,>=1.1.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from sympy>=1.13.3>torch) (1.3.0) 

Requirement already satisfied: MarkupSafe>=2.0 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from jinja2->torch) (3.0.3) 

[notice] A new release of pip is available: 24.0 -> 25.3 

50s 동안 생각함 

형, 지금 **GPU가 안 쓰이는 진짜 이유** 는 이거야 👉 

## **Windows에서 faster-whisper(=CTranslate2)로 GPU를 쓰려면 “CUDA 12.x 런타임 DLL”이 PC에 있어야 하는데, 그게 없음.** 

(Torch는 있어도 되고 없어도 돼. faster-whisper는 Torch 말고 **CTranslate2 + CUDA 12.x** 를 씀) 

아래 5줄만 그대로 진행하면 GPU 바로 잡혀. 

## **1) CUDA 12.4 툴킷 설치** 

**NVIDIA CUDA Toolkit 12.4 (Windows x64)** 설치. 

설치가 끝나면 자동으로 이 폴더가 생겨: 

- `C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v12.4\bin` 

## **2) PATH 자동 추가 확인 (안 돼 있으면 수동 추가)** 

PowerShell에 아래 입력: 

bat 코드 복사 

```
if exist "C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v12.4\bin\cudart64_1
```

  

## `OK` 면 다음 단계로. 

`NOT FOUND` 면 CUDA 설치가 제대로 안 된 거라 **다시 설치** . 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

163/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## (참고: 시스템 환경 변수 `PATH` 에 

- `C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v12.4\bin` 

- `C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v12.4\libnvvp` 

가 포함돼 있어야 함. 설치가 보통 자동으로 넣어줌. 안 들어갔으면 수동으로 추가.) 

## **3) ctranslate2(=GPU 지원 엔진) 재정비** 

bat 코드 복사 

```
pip uninstall -y ctranslate2
pip install --no-cache-dir "ctranslate2==4.6.0"
```

이 버전은 **CUDA 12.x + cuDNN 9** 기준. (형 드라이버 580.88이면 OK) 

## **4) GPU 되는지 3초 테스트** 

bat 코드 복사 `python - << "PY" from faster_whisper import WhisperModel try: WhisperModel("tiny", device="cuda", compute_type="float16") print("GPU OK") except Exception as e: print("GPU FAIL:", e) PY` 

**GPU OK** 나오면 성공. 

FAIL이면, 에러 마지막 줄을 복붙해줘. 

## **5) 이제 실제 실행 (GPU)** 

bat 코드 복사 

`cd "%USERPROFILE%\OneDrive\Desktop\Voice script" python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --vad --engine`   첫 실행만 모델 다운로드, 그 다음부터는 **미친 듯이 빨라짐** . 

## **자주 묻는 포인트 (초간단)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 164/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **왜 Torch에서 CUDA: False 가 떠도 되나?** 

→ 응. **faster-whisper는 Torch가 아니라 CTranslate2** 로 돌려. 중요한 건 **CUDA 12.x DLL** 이 있는지야. 

## **어떤 조합이 안전해?** 

→ `CUDA 12.4 + ctranslate2 4.6.0 + faster-whisper 1.2.0` (형 PC OK) 

## **더 빠르게?** 

→ `--modelsize base` (조금 정확도↓, 속도↑) 

→ SSD 여유 있으면 모델 캐시가 빨리 뜸. 

딱 위 5단계만 하고, ❹ 테스트 결과( **GPU OK/FAIL** ) 한 줄만 알려줘. FAIL이면 그 줄 기준으로 바로 고쳐줄게. 

**==> picture [489 x 38] intentionally omitted <==**

PS C:\WINDOWS\system32> if exist "C:\Program Files\NVIDIA GPU Computing 

Toolkit\CUDA\v12.4\bin\cudart64_12.dll" (echo OK) else (echo NOT FOUND) 

>> 

위치 줄:1 문자:3 

+ if exist "C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v12.4\bi ... 

+   ~ 

if 문의 'if' 뒤에 '('가 없습니다. 

+ CategoryInfo          : ParserError: (:) [], ParentContainsErrorRecordException 

+ FullyQualifiedErrorId : MissingOpenParenthesisInIfStatement 

PS C:\WINDOWS\system32> if exist "C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v12.4\bin\cudart64_12.dll" (echo OK) else (echo NOT FOUND) 

21s 동안 생각함 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

165/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

형, 지금 창은 **PowerShell** 이야. 방금 쓴 `if exist (...)` 는 **CMD 문법** 이라서 에러가 난 거야. 

PowerShell에선 이렇게 확인해야 해. (그대로 복붙) 

## **1) CUDA 런타임 DLL 있는지 PowerShell 방식으로 확인** 

powershell 

코드 복사 

```
if (Test-Path'C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA\v12.4\bin\cudar
```

  `OK` → 다음 단계(ctranslate2 재설치)로 

`NOT FOUND` → **CUDA Toolkit 12.4** 를 설치해야 해 

참고: CMD 창에서 하고 싶으면 

```
if exist "C:\Program Files\NVIDIA GPU Computing
```

```
Toolkit\CUDA\v12.4\bin\cudart64_12.dll" (echo OK) else (echo NOT FOUND)
```

## **2) (NOT FOUND일 때) CUDA Toolkit 12.4 설치** 

NVIDIA 공식 설치 프로그램으로 **CUDA Toolkit 12.4 (Windows x64)** 설치 설치 후 **새 PowerShell 창을 열고** 다시 1번 테스트 

PATH 자동 등록도 확인(선택): 

powershell 코드 복사 

```
$env:Path -split ';' | ? { $_ -like '*CUDA\v12.4\bin*' }
```

하나라도 나오면 OK. 

## **3) ctranslate2/faster-whisper 깨끗이 재설치 (GPU용)** 

powershell 코드 복사 

```
pip uninstall -y ctranslate2 faster-whisper
pip cache purge
pip install --no-cache-dir"ctranslate2==4.6.0"
```

```
pip install --no-cache-dir faster-whisper
```

## **4) GPU 사용 가능 3초 테스트** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

166/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

코드 복사 

powershell 

```
python -c"from faster_whisper import WhisperModel; WhisperModel('tiny', device='
```

  

`GPU OK` 뜨면 성공 

에러면 그 **마지막 10줄** 붙여줘 

## **5) 이제 실제 작업을 GPU로 실행** 

powershell 코드 복사 

```
cd"$env:USERPROFILE\OneDrive\Desktop\Voice script"
python voice_script.py "[20230927]제이타운_매매관련01.m4a"--strong--vad--engine
```

  더 빠르게: `--modelsize base` 아주 빠르게(정확도↓): `--modelsize tiny` 

## **왜 Torch에서** `CUDA: False` **가 떠도 되나?** 

`faster-whisper` 는 **PyTorch가 아니라 CTranslate2** 로 GPU를 씀. 

중요한 건 **CUDA 12.x 런타임 DLL** 이 설치돼 있고, `ctranslate2` 가 그걸 찾는 거야. 그래서 Toolkit 12.4 + ctranslate2 4.6.0 조합을 맞추는 게 핵심. 

필요하면 설치 화면에서 어디를 눌러야 하는지까지 **그림 없이 글로** 단계별로 더 풀어줄 게. 

**==> picture [489 x 38] intentionally omitted <==**

NVIDIA CUDA Toolkit 12.4 (Windows x64) 설치 이걸 어디 서 설치하는지 알려줘야 할거 아냐 씨발새꺄!! 

28s 동안 생각함 

오케이 형. 바로 **공식 사이트** 에서 받는 길 안내할게. (초간단, 그대로 따라만 해) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

167/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **어디서 받냐?** 

**1. CUDA 12.4 다운로드 아카이브** 열기 → 거기서 Windows 선택하면 돼. NVIDIA Develop… 

**2.** 필요하면 **CUDA 툴킷 아카이브 목록** 에서 12.4.1 또는 12.4.0도 고를 수 있어. 

   - NVIDIA Develop… 

**3.** 설치법은 **공식 Windows 가이드** 에 그대로 나와 있어. (우린 핵심만 뽑아 아래 적어 

   - 줌) NVIDIA Docs +1 

## **클릭 순서(그대로)** 

**1.** 위 “CUDA 12.4 아카이브” 페이지 들어가서: 

   - **Operating System:** `Windows` 

   - **Architecture:** `x86_64` 

**Version:** `12.4.x` (가능하면 `12.4.1` ) 

   - **Installer Type:** `exe (local)` 추천 (인터넷 느릴 때 안정적) **Download** 클릭. NVIDIA Develop… 

**2.** 다운받은 `cuda_12.4.x_windows.exe` 실행 → **Express(표준)** 설치 선택. 

**3.** 설치 끝나면 **컴퓨터 재부팅** . 