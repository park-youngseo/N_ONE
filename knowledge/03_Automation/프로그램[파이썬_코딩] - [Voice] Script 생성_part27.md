**1. 개발자 모드 켜기** : 설정 → 개인정보 및 보안 → **개발자용** → **개발자 모드 On** 

**2.** 또는 한 줄로 경고 끄기(기능 영향 없음): 

nginx 

코드 복사 

```
setx HF_HUB_DISABLE_SYMLINKS_WARNING 1
```

(창 닫고 다시 열면 적용) 

## **저장공간·속도 팁 (선택)** 

**캐시 위치** 를 여유 있는 드라이브로 옮기기(예: D:) 

nginx 코드 복사 

```
setx HF_HOME "D:\HFcache"
```

**더 가볍게** 받고 싶으면(다운로드/속도 ↑): `small` 대신 `base` 또는 `tiny` 

css 코드 복사 

`python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --segment 60 --`   (정확도는 `small > base > tiny` 순) 

## **다 받으면 보이는 것** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 146/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

- 콘솔에 `[1/N] 00:00:00 - 00:01:00 | … chars` 처럼 **구간 진행 로그** 가 줄줄 뜸 완료 후 같은 폴더에: 

   - `[... ]_transcript.txt` 

   - `[... ].srt` 

   - `[... ]_segments.csv` 

## **멈춘 것 같으면 이렇게 체크** 

**네트워크가 느린 것** 뿐인 경우가 많아. 그래도 의심되면: 

arduino 코드 복사 

```
tasklist | find /i "python"
```

## 떠 있으면 **다운로드 진행 중** . 

중단하고 모델 크기 낮춰 다시 하려면 `Ctrl+C` 후 위의 `--modelsize base` 로 실행. 

## 원래 그대로 두면 **끝까지 받아져서 자동으로 전사** 들어간다. 

만약 에러로 끊기면 **마지막 10줄** 붙여줘—그 지점에서 바로 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

파일에서 클린이라고 하는 wav 파일은 뭐냐? 

## 형, ** `…_clean.wav` **는 “ **깨끗하게 손질한 복사본** ”이야. 

원본(m4a/mp3)을 바로 쓰지 않고, **듣기·인식이 잘 되도록** 먼저 씻겨서 만든 **WAV 파일** 이 지. 

## **왜 만드느냐?** 

STT(음성→텍스트)가 **깨끗한 소리** 일수록 정확해. 그래서 먼저 이렇게 손질해: 

- **모노 16kHz** 로 변환: 인식기(및 VAD)가 가장 좋아하는 표준 형식. 

- **잡음·저음 웅웅이 줄이기** : highpass/lowpass 필터. 

- **조용한 말 키우고 큰 소리 눌러서 균일하게** : 컴프레서 + 자동 라우드니스(다이내믹 노말라이즈). 

**과도한 피크 방지** : 리미터. 

- (옵션) `--strong` 이면 이 손질을 **더 세게** 적용. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 147/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **한눈에: 필터 → 효과** 

- `highpass=f=100` → 100Hz 아래 **웅웅** 제거 

- `lowpass=f=8000` → 8kHz 위 **쉿쉿/찌릿** 줄임 

- `acompressor` → **큰 소리 눌러** 전체 균일화 

- `volume=6dB` → 전체 **볼륨 상승** 

- `dynaudnorm` → 조용한 구간 **자동 보정** 

- `alimiter` → **찢어짐(클리핑)** 방지 

- `-ac 1 -ar 16000` → **모노·16kHz** 표준화 

## **언제 생기고, 어디 있나?** 

명령을 돌리면 **항상** 원본 옆에 만들어짐. 

- 예) `[20230927]...01_clean.wav` (같은 폴더) 

보통 **원본보다 용량 큼** (WAV라서). 

- 전사(TXT/SRT/CSV) 끝나면 **지워도 됨** . (다시 돌리면 또 만들어짐) 

## **언제** `--strong` **을 쓰나?** 

말소리가 작거나, 멀리서 녹음되어 **크고 작음이 심할 때** → `--strong` 권장 

- 음악이 크거나 **펌핑** 느낌 나면 → `--strong` 없이 실행해봐 

css 코드 복사 

```
python voice_script.py "[파일].m4a" --segment60--engine local --modelsize sm
```

-  

## (원하면 `--vad` 로 **말소리 구간만** 골라서 더 깔끔하게: 

css 코드 복사 

```
python voice_script.py "[파일].m4a" --vad--engine local --modelsize small
```)
```

## **활용 팁** 

- **듣기 비교** : 원본 vs `_clean.wav` 들어보고 더 선명하면 성공. 

- **다른 프로그램 재활용** : `_clean.wav` 는 어떤 편집기에도 잘 먹힘(프리미어, 오더시티 등). 

- **공간 아끼기** : 결과(TXT/SRT/CSV) 확인 후 `_clean.wav` 삭제해도 전사 결과는 그대로 남아. 

필요하면 “손질 약하게/강하게” **두 버전** 을 자동으로 만드는 배치도 만들어줄게. 

**==> picture [489 x 38] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

148/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

warnings.warn(message) vocabulary.txt: 460kB [00:00, 4.29MB/s] 

tokenizer.json: 2.20MB [00:00, 5.54MB/s] | 0.00/484M [00:00<?, ?B/s] model.bin: 100%| ████████████████████████████████████ ████████████████████████████████| 484M/484M [07:44<00:00, 1.04MB/s] 

## 더럽게 느리구나 

21s 동안 생각함 

형, 맞아. **처음 1번만** 모델(484MB)을 내려받느라 엄청 느려 보여. 다운이 끝나면 같은 모델은 **캐시에 저장** 돼서 다음부터는 훨씬 빨라져. 지금 당장 빨라지게 하는 방법만 딱 찍어줄게: 

## **지금 당장 속도 올리는 4가지** 

## **1. 작은 모델로 실행** (정확도↓, 속도↑) 

css 코드 복사 `python voice_script.py "[파일].m4a" --strong --vad --engine local --modelsize`   또는 css 코드 복사 `--modelsize tiny tiny < base < small` (작을수록 빠름) **2. 말 부분만 처리** (쓸데없는 구간 건너뛰기) — 이미 `--vad` 쓰면 OK scss 코드 복사 `--vad  (유지)` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 149/175 

25. 11. 30. 오후 1:03 

 

오후 1:03후 1:03 1:03 프로그램[파이썬/코딩] - [Voice] Script 생성 **3. 윈도우 개발자 모드 켜기** → 캐시 경고 없어지고 파일 복사 덜 함 설정 → “개발자용” → **개발자 모드 On** (선택) 경고만 끄기: nginx 코드 복사 `setx HF_HUB_DISABLE_SYMLINKS_WARNING 1` **4. 캐시를 넉넉한 드라이브로 옮기기** (다운/로드 체감 개선) nginx 코드 복사 `setx HF_HOME "D:\HFcache"` 새 창 열고 다시 실행. **NVIDIA 그래픽카드 있으면(진짜 빠름)** GPU가 있으면 **몇 배 빨라져** . 확인: 코드 복사 `nvidia-smi` 뜬다면, 아래 2줄만 더: perl 코드 복사 `pip install torch --index-url https://download.pytorch.org/whl/cu121` 그리고 `voice_script.py` 에 **device 자동 설정** 한 줄만 바꾸자. 파일 열기: nginx 코드 복사 `notepad voice_script.py local_transcribe` 함수 안에 있는 이 부분을 python 코드 복사 `device = "cpu" compute_type = "int8"` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 150/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

이렇게 교체: 

python 

코드 복사 

```
import torch
device = "cuda"if torch.cuda.is_available() else"cpu"
compute_type = "float16"if device=="cuda"else"int8"
```

저장 후 실행: 

css 코드 복사 

`python voice_script.py "[파일].m4a" --strong --vad --engine local --modelsize`   

## **참고** 

## **한 번 받은 모델은 재사용** → 다음부터 훨씬 빠름. 

더 빠른 체감 원하면 `--segment 30` 처럼 **짧게 자르기** 도 좋아(진행률이 자주 뜸). 

- 원하면 내가 `GPU 자동 감지 버전` 으로 전체 파일을 깔끔히 수정해서 다시 붙여줄게. (말만 해!) 

**==> picture [489 x 38] intentionally omitted <==**

개새끼야 그래픽카드를 처음부터 했어야지 cpu로 왜 설정 을 만들어놔서 이지랄이야 

9s 동안 생각함 

형, 미안해. 처음부터 **GPU 자동** 으로 잡아줬어야 했는데 CPU 기본이라 느렸지. 지금 바로 **GPU(엔비디아)로 돌리게** 바꾸자. 아주 간단해. 

## **1) GPU용 토치 설치 (1번만)** 

perl 

코드 복사 

```
pip install --upgrade "torch" --index-url https://download.pytorch.org/whl/cu121
```

확인: 

  

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

151/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

코드 복사 

 

nginx 

```
python -c "import torch; print('CUDA:', torch.cuda.is_available(), 'Ver:', torch.v
```

  `CUDA: True` 면 준비 완료. **2) 스크립트 2줄만 수정 (GPU 자동 인식)** 파일 열기: nginx 코드 복사 

→ `CUDA: True` 면 준비 완료. 

## **2) 스크립트 2줄만 수정 (GPU 자동 인식)** 

## 파일 열기: 

nginx 

```
notepad"%USERPROFILE%\OneDrive\Desktop\Voice script\voice_script.py"
```

`local_transcribe` 함수 안의 이 부분을 찾고: 

python 코드 복사 `device = "cpu" compute_type = "int8"` 

## 아래처럼 **교체** 하고 저장: 

python 코드 복사 

## `import torch` 

```
device = "cuda"if torch.cuda.is_available() else"cpu"
compute_type = "float16"if device == "cuda"else"int8"
```

```
print(f"[Local] faster-whisper 모델로드: {model_size} ({device}, {compute_type})"
```

  **3) 실행 (GPU 사용)** css 코드 복사 
