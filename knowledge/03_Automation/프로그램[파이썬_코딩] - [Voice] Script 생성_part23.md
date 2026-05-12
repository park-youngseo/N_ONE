
```
defensure_ffmpeg():
ifnot shutil.which("ffmpeg"):
raise SystemExit("FFmpeg 먼저설치 (ffmpeg -version 확인)")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 120/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
defrun(cmd: list, quiet=False) -> str:
```

```
ifnot quiet: print("[CMD]", " ".join(cmd))
```

```
    r = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, check
```

```
return (r.stdout or r.stderr).decode(errors="ignore")
```

```
defsrt_time(t: float) -> str:
```

```
if t < 0: t = 0.0
```

```
    h = int(t // 3600); m = int((t % 3600) // 60); s = int(t % 60); ms = int((t -
```

```
returnf"{h:02d}:{m:02d}:{s:02d},{ms:03d}"
```

```
defffprobe_duration(path: Path) -> float:
```

```
    out = run(["ffprobe","-v","error","-show_entries","format=duration","-of","def
```

```
try: returnfloat(out)
```

```
except: return0.0
```

```
defbuild_filter(strong: bool) -> str:
```

- `base = ["highpass=f=100","lowpass=f=8000"]` 

```
if strong:
```

```
        chain = base + ["acompressor=threshold=-22dB:ratio=3:attack=5:release=250
```

```
else:
```

```
        chain = base + ["dynaudnorm=f=150:g=15"]
```

```
return",".join(chain)
```

```
defpreprocess_to_wav16(src: Path, strong=False) -> Path:
```

- `out = Path(src.with_suffix("").as_posix() + "_clean.wav")` 

- `run(["ffmpeg","-y","-i",str(src),"-ac","1","-ar","16000","-af",build_filter(st` 

- `print(f"[전처리] {src.name} → {out.name} ({'강모드' if strong else '보통'})") return out` 

- `def vad_segments(wav16: Path, aggressiveness=3, frame_ms=30, min_speech=0.3, max_g` 

- `if not _VAD_OK: raise SystemExit("VAD 필요: pip install webrtcvad-wheels")` 

- `wf = wave.open(str(wav16),"rb")` 

- `sr = wf.getframerate(); assert sr in (8000,16000,32000,48000)` 

- `assert wf.getnchannels()==1 and wf.getsampwidth()==2` 

- `frame_bytes = int(sr*(frame_ms/1000.0))*2` 

- `frames=[]; offset=0` 

```
whileTrue:
```

- `b = wf.readframes(int(sr*(frame_ms/1000.0)))` 

- `if not b or len(b) < frame_bytes: break` 

- `start = offset/sr; end = (offset+(frame_bytes//2))/sr` 

```
        frames.append((b,start,end)); offset += (frame_bytes//2)
```

- `wf.close()` 

- `vad = webrtcvad.Vad(aggressiveness)` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 121/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
    segs=[]; cur_s=None; cur_e=None
for b,s,e in frames:
```

```
if vad.is_speech(b,sr):
if cur_s isNone: cur_s=s
else:
if cur_s isnotNone:
if (s-cur_e) <= max_gap: continue
if (cur_e-cur_s) >= min_speech: segs.append((max(0.0,cur_s-pad),
None
```

```
if cur_s isnotNoneand (cur_e-cur_s) >= min_speech:
```

```
        segs.append((max(0.0,cur_s-pad), cur_e+pad))
    merged=[]
for s,e in segs:
```

```
ifnot merged or s>merged[-1][1]: merged.append([s,e])
else: merged[-1][1]=max(merged[-1][1],e)
return [(round(s,2),round(e,2)) for s,e in merged]
```

```
deflocal_transcribe(path: Path, lang: Optional[str], model_size: str="small"
ifnot _FW_OK:
```

```
raise SystemExit("로컬변환용 faster-whisper 설치필요: pip install faster-
    device = "cpu"
```

```
    compute_type = "int8"# CPU에서빠르고가볍게
print(f"[Local] faster-whisper 모델로드: {model_size} ({device}, {compute_typ
    model = WhisperModel(model_size, device=device, compute_type=compute_type)
# vad_filter는우리파이프라인에서이미처리가능하니 False
```

```
    segments, info = model.transcribe(str(path), language=lang, vad_filter=False)
    text_parts=[]
```

```
for seg in segments:
```

```
        text_parts.append(seg.text.strip())
```

```
return" ".join([t for t in text_parts if t])
```

```
defopenai_transcribe(path: Path, lang: Optional[str]) -> str:
```

```
ifnot _OPENAI_OK:
```

```
raise SystemExit("OpenAI 라이브러리설치필요: pip install openai")
    client = OpenAI()
```

```
withopen(path,"rb") as f:
return client.audio.transcriptions.create(
            model="whisper-1", file=f, **({"language":lang} if lang else {})
        ).text or""
```

```
defsmart_transcribe(path: Path, lang: Optional[str], engine: str, model_size: st
"""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 122/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
    engine:
      - 'openai' : 무조건 OpenAI
      - 'local'  : 무조건로컬(faster-whisper)
      - 'auto'   : OPENAI 키가있으면 OpenAI 시도, 429/키없음등실패시로컬로자동전
    """
if engine == "local":
return local_transcribe(path, lang, model_size)
if engine == "openai":
return openai_transcribe(path, lang)
# auto
try:
# 키가없으면여기서바로예외나올수있음
if os.getenv("OPENAI_API_KEY"):
return openai_transcribe(path, lang)
else:
print("[AUTO] OPENAI 키없음 → 로컬엔진으로전환")
return local_transcribe(path, lang, model_size)
except OpenAIError as e:
        msg = str(e)
print(f"[AUTO] OpenAI 실패 → {msg}")
if"insufficient_quota"in msg or"RateLimit"in msg or"401"in msg:
print("[AUTO] 로컬엔진으로자동전환합니다.")
return local_transcribe(path, lang, model_size)
raise
```

```
@dataclass
```

```
classSegment: start: float; end: float; speaker: str; text: str
```

```
defmain():
    ap = argparse.ArgumentParser()
```

```
"
    ap.add_argument("input", help=오디오파일경로(mp3/wav/m4a 등)")
""
    ap.add_argument("--strong", action="store_true", help=강한전처리)
"
    ap.add_argument("--vad", action="store_true", help=대화만추출(VAD)")
""
    ap.add_argument("--diarize", action="store_true", help=화자구분)
    ap.add_argument("--segment", type=int, default=60, help="(VAD/화자구분없을때
"
    ap.add_argument("--lang", type=str, default=None, help=언어힌트(ko, en 등)")
    ap.add_argument("--engine", type=str, choices=["auto","openai","local"], defa
"
    ap.add_argument("--modelsize", type=str, default="small", help=로컬모델크기
    args = ap.parse_args()
```

```
    ensure_ffmpeg()
```

```