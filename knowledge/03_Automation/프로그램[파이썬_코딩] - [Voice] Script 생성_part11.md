**1.** PowerShell을 **관리자** 로 열고: bash 코드 복사 

```
winget install --id Microsoft.VisualStudio.2022.BuildTools -e
```

**2.** 설치 창이 뜨면 체크: **C++ build tools(데스크톱 개발용 C++) MSVC v143 Windows 10/11 SDK** → 설치 완료 후 **재부팅** (권장) **3.** 명령창에서: 

nginx 코드 복사 `pip install webrtcvad` **4.** 이제 실행: css 코드 복사 `python stt_pipeline.py 회의.mp3 --strong --vad` 

## **급하게 지금 당장 돌리고 싶으면(임시)** 

`--vad` 없이도 돌아가니까 일단 이렇게: css 코드 복사 `python stt_pipeline.py 회의.mp3 --strong` 

음악 가사 섞임을 줄이려면, 나중에 **VAD** 를 꼭 켜는 걸 추천! 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

62/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **(선택) pip 업데이트 메시지가 보였을 때** 

필수는 아니지만 깔끔하게 하려면: 

css 

코드 복사 

```
python -m pip install --upgrade pip
```

막히면 **마지막 10줄 에러** 를 그대로 붙여줘. 

지금은 **방법 ①** 가 제일 빠르니까, 그걸로 먼저 해보고 결과 알려줘! 

**==> picture [489 x 38] intentionally omitted <==**

C:\Users\psm23> 

C:\Users\psm23>pip uninstall -y webrtcvad WARNING: Skipping webrtcvad as it is not installed. 

C:\Users\psm23>pip install webrtcvad-wheels Collecting webrtcvad-wheels 

Downloading webrtcvad_wheels-2.0.14-cp311-cp311win_amd64.whl.metadata (5.6 kB) 

Downloading webrtcvad_wheels-2.0.14-cp311-cp311win_amd64.whl (19 kB) 

Installing collected packages: webrtcvad-wheels Successfully installed webrtcvad-wheels-2.0.14 

[notice] A new release of pip is available: 24.0 -> 25.3 [notice] To update, run: python.exe -m pip install -- upgrade pip 

C:\Users\psm23> 

18s 동안 생각함 

좋아! **webrtcvad 설치 완료** 야. 이제 바로 돌려보자. (초간단 단계) 

## **1) 파일 받기** 

stt_pipeline.py 다운로드 

(다운로드 폴더에 저장) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

63/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **2) 확인(선택)** 

명령창에서 아래 한 줄로 VAD 잘 깔렸는지 확인: 

scss 코드 복사 

```
python -c "import webrtcvad; print('VAD OK')"
```

`VAD OK` 가 나오면 준비 끝. 

## **3) 실행(음악/소음 제외 + 작은 소리 보정)** 

다운로드 폴더에 녹음파일(예: `회의.mp3` )을 넣고: 

css 코드 복사 

```
cd %USERPROFILE%\Downloads
python stt_pipeline.py회의.mp3--strong--vad
```

`--strong` : 작은 소리 자동 보정(압축+리미터) 

- `--vad` : **대화만** 뽑아서 가사/배경음 쳐냄 

## **4) 결과(같은 폴더에 생성)** 

`회의_transcript.txt` → 줄글 

`회의.srt` → 자막(시간 포함) 

- `회의_segments.csv` → 시작/끝 시간 + (화자) + 문장 

## **(선택) 화자 구분까지 하고 싶을 때** 

조금 무겁지만 가능해. 먼저 설치: 

nginx 코드 복사 `pip install pyannote.audio torch setx HUGGINGFACE_TOKEN "본인_토큰"` css 코드 복사 

## 그 다음: 

```
python stt_pipeline.py회의.mp3--diarize--vad--strong
```

→ 결과 문장 앞에 `[SPEAKER_00]` , `[SPEAKER_01]` 처럼 **사람별 태그** 가 붙음. 

## **5) 문제 생기면 이렇게** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

64/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

`python` 이 안 되면: `py` 로 바꿔서 실행 

css 코드 복사 

```
py -m pip install openai webrtcvad-wheels pandas
py stt_pipeline.py회의.mp3--strong--vad
```

“키 없어요” 에러: 먼저 OpenAI 키 넣기 

nginx 

코드 복사 

```
setx OPENAI_API_KEY "본인_API키"
(창닫고다시열기)
```

바로 해보고, **마지막 10줄 에러 문구** 만 가져오면 거기서 이어서 잡아줄게! 

**==> picture [489 x 39] intentionally omitted <==**

stt_pipeline.py 다운로드 파일 다시 만들어줘 

## 분석 중지됨 

위에서 보내준 stt_pipeline 과 똑같은 파일이지? 그렇다면 보낼필요는 없어 

만약  stt_pipeline 이 다른 버전으로 업데이트 됐다면 새롭게 메모장으로 만들테니까 코드를 만들어줘 그리고  1, stt_pipeline 의 이름을 변경해도 되는지 

2, 파일을 다른 폴더로 옮겨서 실행해도 괜찮은지 옮겨서 문제가 생긴다면 어떻게 옮겨야 하는지 

3, pip install pyannote.audio torch setx HUGGINGFACE_TOKEN "본인_토큰" 에서 본인 토큰 은 무엇인지 

4. setx OPENAI_API_KEY "본인_API키" 에서 본인 api키는 무엇이며 

어떻게 만들고 어떻게 해야하는지 상세하게 설명해 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

65/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

35s 동안 생각함 

형, 하나씩 깔~끔하게 정리해줄게. 

## **A)** `stt_pipeline.py` **최신본 (메모장에 복붙해서 저장)** 

아래 **전부 복사 → 메모장 →** `stt_pipeline.py` **이름으로 저장** (UTF-8). 이 버전은 `webrtcvad-wheels` 와 잘 맞고, **옵션은 예전과 동일** 해. 

python 

코드 복사 

```
# -*- coding: utf-8 -*-
"""
```

```
stt_pipeline.py  (원클릭파이프라인 v1.1)
```

## `기능 요약:` 

- `1) 긴 녹음 파일을 '들리기 좋게' 정리(보통/강모드)` 

- `2) (선택) 대화만 추출(WebRTC VAD) → 음악/소음 구간 제외` 

- `3) (선택) 화자 구분(pyannote.audio) → SPK_00, SPK_01 ...` 

- `4) 구간별로 잘라 OpenAI Whisper로 전사 → TXT/SRT/CSV 생성` 

## `필요:` 

- `FFmpeg (명령창에서 ffmpeg -version 확인)` 

- `파이썬 패키지:` 

```
    pip install openai webrtcvad-wheels pandas
```

- `(화자 구분 쓸 때) pip install pyannote.audio torch` 

- `키:` 

```
    setx OPENAI_API_KEY "본인키"
```

- `(화자 구분 쓸 때) setx HUGGINGFACE_TOKEN "본인토큰"` 

## `예시:` 

```
  python stt_pipeline.py 회의.mp3 --strong --vad
  python stt_pipeline.py 회의.mp3 --diarize
```

```
  python stt_pipeline.py 통화.m4a --vad --diarize --lang ko
```

## `옵션:` 

```
  --strong      : 강한전처리(작은소리키움)
```

- `--vad         : 말소리 구간만 사용(WebRTC VAD)` 

- `--diarize     : 화자 구분(pyannote.audio 필요)` 

- `--segment N   : VAD/화자 구분 없을 때 N초 분할(기본 60)` 

- `--lang xx     : 언어 힌트(ko, en 등). 기본 자동` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

66/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
"""
```

```
import argparse, math, os, shutil, subprocess, sys, tempfile, wave
from dataclasses import dataclass
from pathlib import Path
from typing importList, Optional, Tuple
import pandas as pd
```

```
# --- OpenAI (Whisper) ---
try:
from openai import OpenAI
    _OPENAI_OK = True
except Exception:
    _OPENAI_OK = False
```

```
# --- VAD (speech only) ---