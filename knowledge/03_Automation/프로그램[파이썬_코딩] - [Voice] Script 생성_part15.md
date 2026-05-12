
## **3) (선택) 화자 구분까지** 

사전 준비(처음 1번만): 

nginx 코드 복사 `pip install pyannote.audio torch setx HUGGINGFACE_TOKEN "여기_본인_토큰" (창 닫고 다시 열기)` 실행: css 코드 복사 `python voice_script.py "회의.mp3" --diarize --vad --strong` → 문장 앞에 `[SPEAKER_00]` 식으로 **사람 태그** 가 붙어. **4) 안 될 때 3가지 체크 1.** `python` 이 안 되면 `py` 로: css 코드 복사 `py voice_script.py "회의.mp3" --strong --vad` **2.** 파일 못 찾음 → **따옴표** (경로/파일에 공백 있을 때 필수). **3.** 키 필요 에러(전사 단계) → 먼저: nginx 코드 복사 `setx OPENAI_API_KEY "여기_본인_API키" (창 닫고 다시 열기)` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

80/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

## **5) (원하면) 더블클릭용 배치 만들기** 

같은 폴더에 **run_voice_script.bat** 만들고 아래 붙여 넣어 저장: 

bash 

코드 복사 

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

이제 **배치 더블클릭** → 파일명 입력 또는 **드래그앤드롭** 으로도 실행 가능. 

실행해보고, 마지막 줄에 뜬 **에러 문구** 가 있으면 그대로 알려줘. 거기서 바로 잡아줄게! 

**==> picture [489 x 38] intentionally omitted <==**

Microsoft Windows [Version 10.0.26100.6899] (c) Microsoft Corporation. All rights reserved. 

C:\Users\psm23>cd "%USERPROFILE%\바탕화면 \Voice script" 

C:\Users\psm23\바탕화면\Voice script>dir C 드라이브의 볼륨에는 이름이 없습니다. 볼륨 일련 번호: CC77-8F78 

C:\Users\psm23\바탕화면\Voice script 디렉터리 

2025-10-28  오전 12:30    <DIR>          . 2025-10-28  오전 12:30    <DIR>          .. 0개 파일                   0 바이트 2개 디렉터리  953,307,123,712 바이트 남음 

C:\Users\psm23\바탕화면\Voice script> 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

81/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

19s 동안 생각함 

형, 지금 폴더가 비어 있으니까 **파일을 먼저 만들어** 넣자. (정말 쉬운 순서) 

## **1)** `voice_script.py` **만들기** 

## **1. 메모장** 열기 

**2.** 아래 코드를 **전부 복사 → 메모장에 붙여넣기** 

**3.** 메모장에서 **파일 → 다른 이름으로 저장** 

   - 파일 이름: `voice_script.py` 

   - 위치: `C:\Users\psm23\바탕화면\Voice script` 

   - 인코딩: **UTF-8** 선택 저장 

python 

코드 복사 

```
# -*- coding: utf-8 -*-
"""
voice_script.py  (원클릭파이프라인 v1.1)
```

## `기능 요약:` 

```
1) 긴녹음파일 '들리기좋게' 전처리(보통/강모드)
```

```
2) (선택) 대화만추출(VAD) → 음악/소음제외
```

```
3) (선택) 화자구분(pyannote.audio) → SPK_00, SPK_01 ...
```

```
4) 조각으로잘라 Whisper 전사 → TXT/SRT/CSV 생성
```

## `필요:` 

```
  pip install openai webrtcvad-wheels pandas
  (화자구분쓸때) pip install pyannote.audio torch
```

```
키:
```

```
  setx OPENAI_API_KEY "본인_API키"
  (화자구분) setx HUGGINGFACE_TOKEN "본인_토큰"
"""
import argparse, math, os, shutil, subprocess, sys, tempfile, wave
from dataclasses import dataclass
from pathlib import Path
from typing importList, Optional, Tuple
import pandas as pd
```

```
try:
from openai import OpenAI
    _OPENAI_OK = True
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

82/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
except Exception:
    _OPENAI_OK = False
```

```
try:
import webrtcvad
    _VAD_OK = True
except Exception:
    _VAD_OK = False
```

```
try:
from pyannote.audio import Pipeline as PyannotePipeline
    _PYANNOTE_OK = True
except Exception:
    _PYANNOTE_OK = False
```

```
defensure_ffmpeg():
ifnot shutil.which("ffmpeg"):
raise SystemExit("FFmpeg 먼저설치하세요. (ffmpeg -version 확인)")
```

```
defrun(cmd: list, quiet=False) -> str:
ifnot quiet: print("[CMD]", " ".join(cmd))
    r = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, check
return (r.stdout or r.stderr).decode(errors="ignore")
```

```
defsrt_time(t: float) -> str:
if t < 0: t = 0.0
```

```
int(t // 3600); m = int((t % 3600) // 60); s = int(t % 60); ms = int((t -
returnf"{h:02d}:{m:02d}:{s:02d},{ms:03d}"
```

```
defffprobe_duration(path: Path) -> float:
```

```
"ffprobe","-v","error","-show_entries","format=duration","-of","def
try: returnfloat(out)
except: return0.0
```

```
defbuild_filter(strong: bool) -> str:
```

```
"highpass=f=100","lowpass=f=8000"]
if strong:
```

```
        chain = base + ["acompressor=threshold=-22dB:ratio=3:attack=5:release=250
else:
        chain = base + ["dynaudnorm=f=150:g=15"]
return",".join(chain)
```

```
defpreprocess_to_wav16(src: Path, strong=False) -> Path:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

83/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
    out = Path(src.with_suffix("").as_posix() + "_clean.wav")
```

```
    run(["ffmpeg","-y","-i",str(src),"-ac","1","-ar","16000","-af",build_filter(st
print(f"[전처리] {src.name} → {out.name} ({'강모드'if strong else'보통'})")
return out
```

```
defvad_segments(wav16: Path, aggressiveness=3, frame_ms=30, min_speech=0.3, max_g
ifnot _VAD_OK: raise SystemExit("VAD 필요: pip install webrtcvad-wheels")
    wf = wave.open(str(wav16),"rb")
    sr = wf.getframerate(); assert sr in (8000,16000,32000,48000)
assert wf.getnchannels()==1and wf.getsampwidth()==2
    frame_bytes = int(sr*(frame_ms/1000.0))*2
    frames=[]; offset=0
whileTrue:
        b = wf.readframes(int(sr*(frame_ms/1000.0)))
ifnot b orlen(b) < frame_bytes: break
        start = offset/sr; end = (offset+(frame_bytes//2))/sr
        frames.append((b,start,end)); offset += (frame_bytes//2)
    wf.close()
    vad = webrtcvad.Vad(aggressiveness)
    segs=[]; cur_s=None; cur_e=None
for b,s,e in frames:
if vad.is_speech(b,sr):
if cur_s isNone: cur_s=s
            cur_e=e
else:
if cur_s isnotNone:
if (s-cur_e) <= max_gap: continue
if (cur_e-cur_s) >= min_speech: segs.append((max(0.0,cur_s-pad),
                cur_s=cur_e=None
```

```
if cur_s isnotNoneand (cur_e-cur_s) >= min_speech:
```

```
        segs.append((max(0.0,cur_s-pad), cur_e+pad))
    merged=[]
for s,e in segs:
ifnot merged or s>merged[-1][1]: merged.append([s,e])
else: merged[-1][1]=max(merged[-1][1],e)
return [(round(s,2),round(e,2)) for s,e in merged]
```

```
defdiarize_segments(wav16: Path):
```

```
ifnot _PYANNOTE_OK: raise SystemExit("화자구분: pip install pyannote.audio t
    token = os.environ.get("HUGGINGFACE_TOKEN")
```

```
ifnot token: raise SystemExit("HUGGINGFACE_TOKEN 환경변수필요")
```
