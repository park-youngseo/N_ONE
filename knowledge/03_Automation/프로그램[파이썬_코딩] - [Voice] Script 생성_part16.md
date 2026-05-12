```
print("[화자구분] 모델로드중…(최초 1회는다운로드)")
```

```
    pipe = PyannotePipeline.from_pretrained("pyannote/speaker-diarization", use_a
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

84/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
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
return rows
```

```
defcut_segment(src_wav: Path, s: float, e: float, out_wav: Path):
    run(["ffmpeg","-y","-i",str(src_wav),"-ss",str(s),"-to",str(e),"-c","copy",st
```

```
defopenai_transcribe(path: Path, lang: Optional[str]=None) -> str:
ifnot _OPENAI_OK: raise SystemExit("OpenAI 라이브러리필요: pip install openai
    client = OpenAI()
withopen(path,"rb") as f:
```

```
        r = client.audio.transcriptions.create(model="whisper-1", file=f, **({"la
returngetattr(r,"text","") or""
```

```
@dataclass
```

```
classSegment: start: float; end: float; speaker: str; text: str
```

```
defmain():
```

```
    ap = argparse.ArgumentParser()
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
    args = ap.parse_args()
```

```
    ensure_ffmpeg()
    src = Path(args.input)
ifnot src.exists(): raise SystemExit(f"파일없음: {src}")
```

```
    wav = preprocess_to_wav16(src, strong=args.strong)
```

```
    diar_rows=[]
if args.diarize:
```

```
        diar_rows = diarize_segments(wav)     # [(s,e,spk)]
        segments = [(s,e) for s,e,_ in diar_rows]
elif args.vad:
```

```
        segments = vad_segments(wav)          # [(s,e)]
else:
        dur = ffprobe_duration(wav)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 85/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
        step = max(10,int(args.segment))
        segments = [(i, min(i+step, dur)) for i inrange(0, int(math.ceil(dur)),
```

```
ifnot segments: raise SystemExit("처리할구간이없습니다. --segment 60 으로시도
```

```
    tmp = Path(tempfile.mkdtemp(prefix="stt_parts_"))
    items: List[Segment] = []
for idx,(s,e) inenumerate(segments, start=1):
        part = tmp / f"part_{idx:04d}.wav"
        cut_segment(wav, s, e, part)
        spk = "UNK"
if diar_rows:
for ds,de,dn in diar_rows:
ifabs(ds-s)<0.01andabs(de-e)<0.01: spk = dn; break
        text = openai_transcribe(part, args.lang)
        items.append(Segment(start=s, end=e, speaker=spk, text=text))
print(f"[{idx}/{len(segments)}] {srt_time(s)} - {srt_time(e)} | {spk} | {
```

```
    base = src.with_suffix("")
    txt_path = Path(str(base) + "_transcript.txt")
    srt_path = Path(str(base) + ".srt")
    csv_path = Path(str(base) + "_segments.csv")
```

```
withopen(txt_path,"w",encoding="utf-8") as f:
for seg in items:
            prefix = f"[{seg.speaker}] " if args.diarize else""
            f.write(prefix + seg.text.strip() + "\n")
```

```
    pd.DataFrame([{
"start":seg.start,"end":seg.end,
"start_srt":srt_time(seg.start),"end_srt":srt_time(seg.end),
"speaker": seg.speaker if args.diarize else"",
"text": seg.text
```

```
    } for seg in items]).to_csv(csv_path,index=False,encoding="utf-8-sig")
```

```
withopen(srt_path,"w",encoding="utf-8") as f:
for i,seg inenumerate(items,start=1):
```

```
            f.write(f"{i}\n{srt_time(seg.start)} --> {srt_time(seg.end)}\n")
            line = (f"[{seg.speaker}] " if args.diarize else"") + (seg.text.stri
            f.write(line + "\n\n")
```

```
print("\n[완료]")
print("TXT :", txt_path.name)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

86/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
print("SRT :", srt_path.name)
print("CSV :", csv_path.name)
```

```
print("팁) 음악가사가섞이면 --vad, 작은소리는 --strong, 화자구분은 --diarize
```

```
if __name__ == "__main__":
    main()
```

## **2) 음성 파일도 이 폴더에 넣기** 

예: `회의.mp3` 를 **같은 폴더** (Voice script)에 복사 

## **3) 실행** 

명령창에서 이미 폴더가 맞으니, 그냥: 

css 

코드 복사 

```
python voice_script.py "회의.mp3" --strong--vad
```

(화자 구분까지 원하면 `--diarize` 추가) 

## **4) 결과** 

같은 폴더에 

`회의_transcript.txt` , `회의.srt` , `회의_segments.csv` 가 생기면 성공! 

막히면, 방금 친 명령과 **마지막 10줄** 을 그대로 붙여줘. 거기서 딱 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

Microsoft Windows [Version 10.0.26100.6899] (c) Microsoft Corporation. All rights reserved. 

C:\Users\psm23>python voice_script.py "[20230927]제 이타운_매매관련01.m4a" --strong --vad python: can't open file 'C:\\Users\\psm23\\voice_script.py': [Errno 2] No such file or directory 

C:\Users\psm23> 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

87/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

24s 동안 생각함 

 

형, 지금 에러는 **현재 폴더에** `voice_script.py` **가 없어서** 야. 

두 가지 중 **아무거나** 그대로 복붙해서 해봐. 

## **방법 A) 풀경로로 바로 실행(추천)** 

bat 코드 복사 

```
python "%USERPROFILE%\바탕화면\Voice script\voice_script.py" "%USERPROFILE%\[20230
```

  **방법 B) 스크립트 폴더로 이동 후 실행** bat 코드 복사 `cd "%USERPROFILE%\바탕화면\Voice script" python voice_script.py "%USERPROFILE%\[20230927]제이타운_매매관련01.m4a" --strong -`   

## **방법 B) 스크립트 폴더로 이동 후 실행** 

## **혹시 파일 위치가 다른 경우 확인하는 법** 

bat 코드 복사 

```
dir "%USERPROFILE%\[20230927]*.m4a"
```

파일이 보이면 위 명령 그대로 OK. 

안 보이면 파일이 있는 폴더로 바꿔서, **오디오 파일도 풀경로** 로 적어줘: 

bat 코드 복사 

`python "%USERPROFILE%\바탕화면\Voice script\voice_script.py" "D:\녹음\[20230927]제`   

## 규칙: **공백·대괄호** 가 있으면 항상 `"따옴표"` 로 감싸기! 

실행하면 같은 폴더에 `…_transcript.txt` , `.srt` , `_segments.csv` 가 생겨. 

**==> picture [489 x 39] intentionally omitted <==**

Microsoft Windows [Version 10.0.26100.6899] 

(c) Microsoft Corporation. All rights reserved. 

C:\Users\psm23>python voice_script.py "[20230927]제 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

88/175 

25. 11. 30. 오후 1:03 