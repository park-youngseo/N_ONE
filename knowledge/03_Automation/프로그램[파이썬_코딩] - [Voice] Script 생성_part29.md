
## **화자 구분은 나중에** 

필요하면 `--diarize` 를 붙여서 화자 태그도 달 수 있어. 지금은 결과부터 확인하고, 원하 면 이어서 세팅해줄게. 

혹시 폴더에 파일이 **안 보이거나** 싱크가 많이 틀리면, **문제 구간의 시간대(예: 00:12:00~00:14:00)**만 알려줘. 그 구간만 **다시 보정** 해서 뽑는 커맨드까지 바로 만들어줄 게. 

**==> picture [489 x 38] intentionally omitted <==**

일단 병신아 빠르게 만드려면 그레픽 카드로 빠르게 해야 한다며? 

## 맞아. **GPU로 돌리면 훨씬 빨라** . 

지금 딱 **3단계** 만 하면 된다. (복붙만) 

## **1) GPU 준비 확인** 

bat 

코드 복사 

```
nvidia-smi
```

숫자랑 그래픽카드 정보가 나오면 OK. 

## **2) GPU용 PyTorch 설치(감지용)** 

bat 

코드 복사 

```
pip install --upgrade "torch" --index-url https://download.pytorch.org/whl/cu121
python -c "import torch; print('CUDA:', torch.cuda.is_available(), 'CUDA Ver:', to
```

 

 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

157/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

`CUDA: True` 나오면 준비 완료. 

## **3) 스크립트 2줄만 바꾸기 (GPU 자동 사용)** 

파일 열기: 

bat 코드 복사 

```
notepad "%USERPROFILE%\OneDrive\Desktop\Voice script\voice_script.py"
```

`local_transcribe` 함수 안에서 아래 **두 줄** 을 찾아: python 코드 복사 `device = "cpu" compute_type = "int8"` 

## → 이렇게 **교체** 하고 저장: 

python 코드 복사 

```
import torch
```

```
device = "cuda"if torch.cuda.is_available() else"cpu"
compute_type = "float16"if device == "cuda"else"int8"
```

```
print(f"[Local] faster-whisper 모델로드: {model_size} ({device}, {compute_type})"
```

  **실행 (이제 GPU 사용)** bat 코드 복사 `cd "%USERPROFILE%\OneDrive\Desktop\Voice script" python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --vad --engine 제이타운_매매관련01.m4a" --strong --vad --engine _매매관련01.m4a" --strong --vad --engine 매매관련01.m4a" --strong --vad --engine 01.m4a" --strong --vad --engine`   

## **실행 (이제 GPU 사용)** 

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --vad --engine 제이타운_매매관련01.m4a" --strong --vad --engine _매매관련01.m4a" --strong --vad --engine 매매관련01.m4a" --strong --vad --engine 01.m4a" --strong --vad --engine
```

## **속도 더 올리고 싶으면** 

조금 더 가볍게: `--modelsize base` 

아주 빠르게(정확도↓): `--modelsize tiny` 

## **체크 포인트** 

여전히 느리면: `nvidia-smi` 가 떠 있는지, 위 2)에서 `CUDA: True` 가 나오는지 확인. 에러 나면 **마지막 10줄** 붙여줘. 거기서 바로 잡아줄게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

158/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

야 그레픽 카드를 사용하면 빠르다면서 cpu를 사용해? 

25s 동안 생각함 

형, 맞아. **GPU로 돌리면 훨씬 빨라** . 지금 CPU로 돈 이유는 **코드가 GPU를 쓰라고 못 지시** 해서야. 바로 고치자. (복붙만 하면 끝) 

## **1) GPU 준비 확인 (2줄)** 

scss 

코드 복사 

```
nvidia-smi
```

```
python -c "import torch; print('CUDA:', torch.cuda.is_available())"
```

둘 다 OK면 GPU 쓸 준비 끝. 

`CUDA: False` 면 아래 2-1 실행. 

## **2-1) PyTorch GPU 설치(감지용, 1번만)** 

perl 

코드 복사 

```
pip install --upgrade "torch" --index-url https://download.pytorch.org/whl/cu121
python -c "import torch; print('CUDA:', torch.cuda.is_available())"
```

  

`True` 나오면 OK. 

## **2) 코드 2줄 교체(“무조건 GPU 우선”)** 

파일 열기: 

nginx 

코드 복사 

```
notepad"%USERPROFILE%\OneDrive\Desktop\Voice script\voice_script.py"
```

`local_transcribe` 함수 안에 **아래 두 줄** 을 찾아: 

python 코드 복사 

```
device = "cpu"
```

```
compute_type = "int8"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 159/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## → **이렇게 교체하고 저장** : 

python 

코드 복사 

```
import torch
device = "cuda"if torch.cuda.is_available() else"cpu"
compute_type = "float16"if device == "cuda"else"int8"
print(f"[Local] faster-whisper 모델로드: {model_size} ({device}, {compute_type})"
```

  만약 torch 없이도 GPU를 바로 쓰고 싶다면(ctranslate2 직접 감지): 

python 코드 복사 

`try: import ctranslate2 use_cuda = getattr(ctranslate2, "get_cuda_device_count", lambda:0)() > 0 except: use_cuda = False device = "cuda" if use_cuda else "cpu" compute_type = "float16" if device == "cuda" else "int8"`   

## **3) GPU 테스트(3초)** 

scss 코드 복사 

```
python -c "from faster_whisper import WhisperModel; WhisperModel('tiny', device='
```

  `GPU OK` 뜨면 준비 완료. 에러가 “No GPU devices”면 **NVIDIA 드라이버** 먼저 업데이트. 

## **4) 이제 진짜 실행 (GPU 사용)** 

css 

코드 복사 

```
cd "%USERPROFILE%\OneDrive\Desktop\Voice script"
```

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--vad--engine
```

  더 빠르게: `--modelsize base` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

160/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

초고속(정확도↓): `--modelsize tiny` 

- 정확도↑: `--modelsize medium` (조금 느림) 

## **안 되면(딱 한 번만 점검)** 

- 드라이버 확인: `nvidia-smi` 가 떠야 함 

- 토치 CUDA: `python -c "import torch; print(torch.cuda.is_available())"` → True 그래도 오류면 **마지막 10줄** 붙여줘. 그 라인에서 바로 잡아줄게. 

Microsoft Windows [Version 10.0.26100.6899] 

(c) Microsoft Corporation. All rights reserved. 

C:\Users\psm23>nvidia-smi Tue Oct 28 03:02:35 2025 

+------------------------------------------------------------- 

----------------------------+ 

| NVIDIA-SMI 580.88                 Driver Version: 580.88 CUDA Version: 13.0     | 

+-----------------------------------------+------------------ 

------+----------------------+ 

| GPU  Name                  Driver-Model | Bus-Id Disp.A | Volatile Uncorr. ECC | 

| Fan  Temp   Perf          Pwr:Usage/Cap | Memory-Usage | GPU-Util  Compute M. | |                                         |                        | MIG M. | 

|==================================== =====+========================+====== ================| 

|   0  NVIDIA GeForce RTX 3060      WDDM  | 00000000:01:00.0 Off |                  N/A | | 39%   29C    P8              8W /  170W |       0MiB / 12288MiB |      0%      Default | |                                         |                        | N/A | +-----------------------------------------+-----------------------+----------------------+ 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 161/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

+------------------------------------------------------------- 

----------------------------+ 

| Processes: | |  GPU   GI   CI              PID   Type   Process name GPU Memory | |        ID   ID Usage      | 

|==================================== ===================================== ================| 

|  No running processes found | 

+----------------------------------------------------------------------------------------+ 

C:\Users\psm23>python -c "import torch; print('CUDA:', torch.cuda.is_available())" CUDA: False 

C:\Users\psm23>pip install --upgrade "torch" --indexurl https://download.pytorch.org/whl/cu121 Looking in indexes: https://download.pytorch.org/whl/cu121 Requirement already satisfied: torch in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (2.9.0) Requirement already satisfied: filelock in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (3.20.0) Requirement already satisfied: typingextensions>=4.10.0 in 

c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (4.15.0) Requirement already satisfied: sympy>=1.13.3 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (1.14.0) Requirement already satisfied: networkx>=2.5.1 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (3.5) Requirement already satisfied: jinja2 in c:\users\psm23\appdata\local\programs\python\py thon311\lib\site-packages (from torch) (3.1.6) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 