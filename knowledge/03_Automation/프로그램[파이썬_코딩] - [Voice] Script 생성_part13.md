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
        diar_rows = diarize_segments(wav)     # [(s,e,spk)]
        segments = [(s,e) for s,e,_ in diar_rows]
elif args.vad:
        segments = vad_segments(wav)          # [(s,e)]
else:
        dur = ffprobe_duration(wav)
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
withopen(txt_path,"w",encoding="utf-8") as f:
for seg in items:
            prefix = f"[{seg.speaker}] " if args.diarize else""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

70/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
            f.write(prefix + seg.text.strip() + "\n")
```

```
    pd.DataFrame([{
```

```
"start":seg.start,"end":seg.end,
"start_srt":srt_time(seg.start),"end_srt":srt_time(seg.end),
```

```
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
```

```
            line = (f"[{seg.speaker}] " if args.diarize else"") + (seg.text.stri
            f.write(line + "\n\n")
```

```
print("\n[완료]")
print("TXT :", txt_path.name)
print("SRT :", srt_path.name)
print("CSV :", csv_path.name)
print("팁) 음악가사가섞이면 --vad, 작은소리는 --strong, 화자구분은 --diarize
```

```
if __name__ == "__main__":
    main()
```

## **B) 질문 1–4에 대한 답** 

## **1) 이름 바꿔도 돼?** 

**돼.** 예: `stt_pipeline_my.py` 

실행만 이름에 맞춰: 

css 코드 복사 

```
python stt_pipeline_my.py회의.mp3--strong--vad
```

이름만 바꿔도 **기능은 같아** . (파이썬 파일 안의 경로·이름 의존 없음) 

## **2) 다른 폴더로 옮겨도 돼? 옮길 때 주의는?** 

**돼.** 단, 실행할 때 **작업 위치** (cd)만 맞추면 돼. 

방법 A) 오디오가 있는 폴더로 이동해서 실행 

mathematica 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

71/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
cdD:\meeting\2025-10
```

```
pythonD:\tools\stt_pipeline.py회의.mp3
```

방법 B) 파이프라인이 있는 폴더로 이동해서 **파일 경로를 풀로 지정** 

mathematica 코드 복사 

```
cdD:\tools
```

```
pythonstt_pipeline.pyD:\meeting\2025-10\회의.mp3
```

**문제 날 때 대부분** : 파일 경로를 못 찾음 → 위처럼 **cd** 또는 **풀 경로** 로 해결. 

FFmpeg는 **PATH** 에만 들어 있으면 **어느 폴더** 에서 실행해도 OK. (형은 이미 PATH 설 정 완료) 

## **3)** `HUGGINGFACE_TOKEN` **의 본인 토큰은?** 

**의미** : 화자 구분 모델을 내려받게 해주는 **허깅페이스 계정의 비밀키** (읽기용). 

- **만드는 법** : 

**1.** huggingface.co 회원가입 → 로그인 

**2.** 우상단 프로필 → **Settings** → **Access Tokens** 

**3. New token** → 이름 아무거나 → 권한은 **Read** → 생성 

**4.** 나온 긴 문자열이 **본인 토큰** 

**5.** 윈도우에서 저장: 

nginx 코드 복사 `setx HUGGINGFACE_TOKEN "여기_복사한_토큰"` 

창 닫고 새 창 열기(중요) 

## **확인** : 

bash 코드 복사 `echo %HUGGINGFACE_TOKEN%` 

(토큰 앞 몇 글자만 보여도 OK) 

## **4)** `OPENAI_API_KEY` **의 본인 API키는?** 

- **의미** : Whisper 전사를 쓰게 해주는 **오픈AI 계정의 비밀키** . **만드는 법** : 

**1.** OpenAI 계정 만들기 → 로그인 

**2. API Keys** (대시보드) → **Create new secret key** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

72/175 

25. 11. 30. 오후 1:03 

 

**==> picture [145 x 9] intentionally omitted <==**

**----- Start of picture text -----**<br>
프로그램[파이썬/코딩] - [Voice] Script 생성<br>**----- End of picture text -----**<br>


**3.** 나온 긴 문자열이 **본인 API키** → 바로 복사(다시 못 봄) 

**4.** 윈도우에서 저장: 

nginx 코드 복사 `setx OPENAI_API_KEY "여기_복사한_API키"` 

창 닫고 새 창 열기 

**확인** (간단): 

nginx 코드 복사 

```
python -c "import os;print('OK' if os.getenv('OPENAI_API_KEY') else 'NO')"
```

→ `OK` 뜨면 준비 완료. 

## **C) 설치·실행 한 줄 가이드(정리)** 

**1.** 필수 설치(한 번만) 

nginx 코드 복사 

```
pip install openai webrtcvad-wheels pandas
```

## (화자 구분 쓸 때) 

nginx 코드 복사 `pip install pyannote.audio torch` 키 저장(있는 경우) nginx 코드 복사 `setx OPENAI_API_KEY "본인_API키" setx HUGGINGFACE_TOKEN "본인_토큰"` 실행 예시 css 코드 복사 

## **2.** 키 저장(있는 경우) 

## **3.** 실행 예시 

```
python stt_pipeline.py회의.mp3--strong--vad
python stt_pipeline.py회의.mp3--diarize
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

73/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

필요하면 `stt_pipeline.py` 를 **“번호 고르기 메뉴형”** (완전 쉬운 모드)으로도 바꿔줄 수 있 어. 

지금 코드로 한 번 돌려보고, **어디에서 막혔는지** 만 알려줘—그 지점부터 딱 잡아줄게! 

**==> picture [489 x 39] intentionally omitted <==**

다시 적용해줘 나는 stt_pipeline.py의 이름을 voice_script 로 교체했고 bat의 내용중 voice_script.py  (원클릭 파이프라인 v1.1) 이 렇게 교체를 했어 

지정된 파일의 폴더는 바탕화면>Voice script>voice_script.py 로 저장했어 

19s 동안 생각함 

좋아 형. 이름을 **voice_script.py** 로 바꾸고, 폴더가 

**바탕화면 → Voice script → voice_script.py** 맞지? 그 기준으로 **다시 적용** 방법을 딱 맞 춰 줄게. 

## **1) 실행(명령창 버전) — 제일 확실** 

## **1. 명령 프롬프트(cmd)** 열기 

**2.** 폴더 이동(공백이 있어서 **따옴표** 꼭!) 

bash 코드 복사 

```
cd"%USERPROFILE%\Desktop\Voice script"
```

**3.** 바로 실행(소리 보정 + 대화만) 

css 

코드 복사 