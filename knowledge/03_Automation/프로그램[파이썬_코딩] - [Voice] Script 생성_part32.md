    print("GPU OK")
except Exception as e:
    print("GPU FAIL:", type(e).__name__, e)
PY
```

**GPU OK** 나오면 성공. 

## **D. 스크립트가 GPU를 잡게 2줄만 바꾸기** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

173/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

코드 복사 

powershell 

```
notepad "$env:USERPROFILE\OneDrive\Desktop\Voice script\voice_script.py"
```

## 아래 두 줄을 찾아서 

python `device = "cpu" compute_type = "int8"` 

코드 복사 

## 이걸로 **교체** 하고 저장: 

python 코드 복사 

```
import torch
```

```
device = "cuda"if torch.cuda.is_available() else"cpu"
compute_type = "float16"if device=="cuda"else"int8"
print(f"[Local] faster-whisper 모델로드: {model_size} ({device}, {compute_type})"
```

  

## **E. 실행 (이제 GPU)** 

powershell 코드 복사 

```
cd"$env:USERPROFILE\OneDrive\Desktop\Voice script"
python voice_script.py "[20230927]제이타운_매매관련01.m4a"--strong--vad--engine
```

  더 빠르게 = `--modelsize base` / 아주 빠르게(정확도↓) = `--modelsize tiny` . 

## **안 되면(딱 이 두 개만 보내줘)** 

**1.** CUDA DLL 검색 결과(위 A-1 출력 줄) 

**2.** C 단계 테스트의 **“GPU FAIL: …” 마지막 10줄** 

그 두 줄만 보면 바로 잡는다. 

**==> picture [489 x 39] intentionally omitted <==**

## Windows PowerShell 

Copyright (C) Microsoft Corporation. All rights reserved. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

174/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

새로운 기능 및 개선 사항에 대 한 최신 PowerShell을 설 치 하세요! https://aka.ms/PSWindows 

PS C:\WINDOWS\system32> # 1) CUDA 런타임 DLL 위 치 자동 검색 

>> Get-ChildItem 'C:\Program Files\NVIDIA GPU Computing Toolkit\CUDA' -Recurse -Filter cudart64_12.dll -ErrorAction SilentlyContinue | SelectObject -First 1 -Expand FullName >> 

16s 동안 생각함 

형, 방금 명령에 **아무 출력이 없었다 = CUDA 12.x가 아직 설치 안 됨** 이야. 

이대로 따라 하면 GPU로 바로 돌 수 있어. 

## **1) CUDA 12.4 설치 (그림 그대로 고르기)** 

웹에서 **“CUDA Toolkit 12.4 archive”** 검색 → NVIDIA 페이지 열기 

- 화면에서 이렇게 선택 

   - **Operating System:** Windows 

   - **Architecture:** x86_64 

   - **Version:** 11 

   - **Installer Type: exe (local)** ← 인터넷 느릴 때 안정적 

- 내려받은 `cuda_12.4.x_windows.exe` 실행 → **Express(표준)** → 설치 후 **재부팅** 

- 드라이버는 이미 최신(580.88)이니, 설치 중 “드라이버” 체크가 보이면 유지해도 되고 생략해도 됨. 

## **2) 설치 확인 (PowerShell에 그대로 붙여넣기)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

175/175 
