try:
import webrtcvad  # webrtcvad-wheels 도동일모듈명
    _VAD_OK = True
except Exception:
    _VAD_OK = False
```

```
# --- Diarization ---
try:
```

```
from pyannote.audio import Pipeline as PyannotePipeline
    _PYANNOTE_OK = True
except Exception:
    _PYANNOTE_OK = False
```

```
defensure_ffmpeg():
ifnot shutil.which("ffmpeg"):
```

```
raise SystemExit("FFmpeg 먼저설치하세요: https://ffmpeg.org (설치후새창에
```

```
defrun(cmd: list, quiet=False) -> str:
ifnot quiet:
print("[CMD]", " ".join(cmd))
    r = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, check
return (r.stdout or r.stderr).decode(errors="ignore")
```

```
defsrt_time(t: float) -> str:
if t < 0: t = 0.0
    h = int(t // 3600); m = int((t % 3600) // 60); s = int(t % 60); ms = int((t -
returnf"{h:02d}:{m:02d}:{s:02d},{ms:03d}"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 67/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
defffprobe_duration(path: Path) -> float:
```

```
    out = run(["ffprobe","-v","error","-show_entries","format=duration","-of","def
try: returnfloat(out)
except: return0.0
```

```
defbuild_filter(strong: bool) -> str:
    base = ["highpass=f=100","lowpass=f=8000"]
if strong:
        chain = base + ["acompressor=threshold=-22dB:ratio=3:attack=5:release=250
else:
        chain = base + ["dynaudnorm=f=150:g=15"]
return",".join(chain)
```

```
defpreprocess_to_wav16(src: Path, strong=False) -> Path:
```

```
    out = Path(src.with_suffix("").as_posix() + "_clean.wav")
```

```
    af = build_filter(strong)
```

```
    run(["ffmpeg","-y","-i",str(src),"-ac","1","-ar","16000","-af",af,str(out)])
print(f"[전처리] {src.name} → {out.name} ({'강모드'if strong else'보통'})")
return out
```

```
defvad_segments(wav16: Path, aggressiveness=3, frame_ms=30, min_speech=0.3, max_g
ifnot _VAD_OK: raise SystemExit("VAD 필요: pip install webrtcvad-wheels")
    wf = wave.open(str(wav16),"rb")
```

```
    sr = wf.getframerate(); assert sr in (8000,16000,32000,48000)
```

```
assert wf.getnchannels()==1and wf.getsampwidth()==2
    frame_bytes = int(sr*(frame_ms/1000.0))*2
    frames = []; offset = 0
whileTrue:
```

```
        b = wf.readframes(int(sr*(frame_ms/1000.0)))
```

```
ifnot b orlen(b) < frame_bytes: break
```

```
        start = offset/sr; end = (offset+(frame_bytes//2))/sr
```

```
        frames.append((b,start,end)); offset += (frame_bytes//2)
    wf.close()
```

```
    vad = webrtcvad.Vad(aggressiveness)
```

```
    segs=[]; cur_s=None; cur_e=None
for b,s,e in frames:
if vad.is_speech(b,sr):
if cur_s isNone: cur_s=s
            cur_e=e
else:
if cur_s isnotNone:
if (s-cur_e) <= max_gap: continue
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

68/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
if (cur_e-cur_s) >= min_speech: segs.append((max(0.0,cur_s-pad),
```

```
                cur_s=cur_e=None
```

```
if cur_s isnotNoneand (cur_e-cur_s) >= min_speech:
```

```
        segs.append((max(0.0,cur_s-pad), cur_e+pad))
    merged=[]
```

```
for s,e in segs:
```

```
ifnot merged or s>merged[-1][1]: merged.append([s,e])
```

```
else: merged[-1][1]=max(merged[-1][1],e)
```

```
return [(round(s,2),round(e,2)) for s,e in merged]
```

```
defdiarize_segments(wav16: Path):
```

- `if not _PYANNOTE_OK: raise SystemExit("화자 구분: pip install pyannote.audio t` 

- `token = os.environ.get("HUGGINGFACE_TOKEN")` 

- `if not token: raise SystemExit("HUGGINGFACE_TOKEN 환경변수 필요")` 

```
print("[화자구분] 모델로드중…(최초 1회는다운로드)")
```

```
    pipe = PyannotePipeline.from_pretrained("pyannote/speaker-diarization", use_a
    diar = pipe(str(wav16))
```

```
    rows=[]
```

```
for turn,_,spk in diar.itertracks(yield_label=True):
```

```
        rows.append((round(turn.start,2), round(turn.end,2), spk))
```

```
return rows
```

```
defcut_segment(src_wav: Path, s: float, e: float, out_wav: Path):
```

```
    run(["ffmpeg","-y","-i",str(src_wav),"-ss",str(s),"-to",str(e),"-c","copy",st
```

- `def openai_transcribe(path: Path, lang: Optional[str]=None) -> str:` 

- `if not _OPENAI_OK: raise SystemExit("OpenAI 라이브러리 필요: pip install openai` 

- `client = OpenAI()` 

```
withopen(path,"rb") as f:
```

```
        r = client.audio.transcriptions.create(model="whisper-1", file=f, **({"la
```

```
returngetattr(r,"text","") or""
```

## `@dataclass` 

```
classSegment: start: float; end: float; speaker: str; text: str
```

```
defmain():
```

```
    ap = argparse.ArgumentParser()
```

```
"
    ap.add_argument("input", help=오디오파일경로(mp3/wav/m4a 등)")
```

```
""
    ap.add_argument("--strong", action="store_true", help=강한전처리)
```

```
"
    ap.add_argument("--vad", action="store_true", help=대화만추출(WebRTC VAD)")
```

```
""
    ap.add_argument("--diarize", action="store_true", help=화자구분)
```

```
    ap.add_argument("--segment", type=int, default=60, help="(VAD/화자구분없을때
```

```
"
    ap.add_argument("--lang", type=str, default=None, help=언어힌트(ko, en 등)")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

69/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
    args = ap.parse_args()
```

```