    src = Path(args.input)
ifnot src.exists(): raise SystemExit(f"파일없음: {src}")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

123/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
    wav = preprocess_to_wav16(src, strong=args.strong)
```

```
    diar_rows=[]
if args.diarize:
```

```
raise SystemExit("화자구분(--diarize)은 OpenAI가아니라 pyannote를쓰며, 본
elif args.vad:
```

```
        segments = vad_segments(wav)
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
```

```
    items: List[Segment] = []
for idx,(s,e) inenumerate(segments, start=1):
        part = tmp / f"part_{idx:04d}.wav"
```

```
        run(["ffmpeg","-y","-i",str(wav),"-ss",str(s),"-to",str(e),"-c","copy",st
        text = smart_transcribe(part, args.lang, args.engine, args.modelsize)
        items.append(Segment(start=s, end=e, speaker="UNK", text=text))
print(f"[{idx}/{len(segments)}] {srt_time(s)} - {srt_time(e)} | {len(text
```

```
    base = src.with_suffix("")
    txt_path = Path(str(base) + "_transcript.txt")
    srt_path = Path(str(base) + ".srt")
    csv_path = Path(str(base) + "_segments.csv")
```

```
withopen(txt_path,"w",encoding="utf-8") as f:
for seg in items: f.write(seg.text.strip() + "\n")
```

```
    pd.DataFrame([{
```

```
"start":seg.start,"end":seg.end,
```

```
"start_srt":srt_time(seg.start),"end_srt":srt_time(seg.end),
"speaker":"", "text": seg.text
```

```
    } for seg in items]).to_csv(csv_path,index=False,encoding="utf-8-sig")
```

```
withopen(srt_path,"w",encoding="utf-8") as f:
for i,seg inenumerate(items,start=1):
```

```
            f.write(f"{i}\n{srt_time(seg.start)} --> {srt_time(seg.end)}\n{(seg.te
```

```
print("\n[완료]")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 124/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
print("TXT :", txt_path.name)
print("SRT :", srt_path.name)
print("CSV :", csv_path.name)
"
print(팁) 무료로쓸땐 --engine local --modelsize small/base/medium 중선택. 정
```

```
if __name__ == "__main__":
    main()
```

## **3) 로컬 엔진으로 바로 실행** 

css 코드 복사 

```
python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong--segment60--
```

  되면 다음엔 `--vad` 켜서: css 코드 복사 

```
python voice_script.py "[...].m4a" --strong--vad--engine local --modelsize
```

  더 정확히: `--modelsize medium` (속도 느려짐) 

최고 정확도: `--modelsize large-v3` (다운로드/속도 가장 무거움) 

## **(선택) 배치 파일도 “자동 전환” 지원으로 업데이트** 

`run_voice_all.bat` 안의 실행 줄을 이렇게 바꿔줘요. (로컬 엔진 설치도 추가) 

**1.** 패키지 설치 파트에 추가: 

matlab 코드 복사 

```
%_PY% -m pip show faster-whisper >nul 2>nul || %_PY% -m pip install -q faster-whis
%_PY% -m pip show torch >nul 2>nul || %_PY% -m pip install -q torch --index-url h
```

  

## **2.** 실행 줄을 **자동엔진(auto) + 로컬모델 크기 지정** 으로: 

arduino 

코드 복사 

```
set "BASE_OPTS=--strong --vad --engine auto --modelsize small"
```

## 이제 배치를 더블클릭하면: 

## 오픈AI 크레딧이 있으면 **클라우드** 로, 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 125/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

없으면 자동으로 **로컬(faster-whisper)** 로 돌아가서 결과(TXT/SRT/CSV)를 만들어. 

## **한 줄 요약** 

- 지금 에러는 **크레딧 없음(429)** . 

- **유료로 계속 쓰려면** 결제정보/크레딧 추가. 

- **무료로 바로 돌리려면** 위 코드로 **로컬 엔진** 붙이고 `--engine local` 로 실행! 

해보고 **마지막 10줄** 만 보내줘. 거기서 바로 이어서 잡아줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

C:\Users\psm23>python voice_script.py "[20230927]제 이타운_매매관련01.m4a" --strong --segment 60 --engine local --modelsize small 

python: can't open file 

'C:\\Users\\psm23\\voice_script.py': [Errno 2] No such file or directory 

C:\Users\psm23> 

형, 지금 에러는 **현재 폴더가 달라서** 야. `voice_script.py` 가 있는 폴더로 가서 실행하면 끝! 

## **1) 스크립트 있는 폴더로 이동** 

bat 

코드 복사 

```
cd "%USERPROFILE%\OneDrive\Desktop\Voice script"
```

## **2) (한 번만) 로컬 엔진 설치** 

bat 

코드 복사 

```
pip install faster-whisper
```

```
pip install torch --index-url https://download.pytorch.org/whl/cpu
```

## **3) 파일 확인** 

bat 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

126/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

 

```
dir
```

여기서 `voice_script.py` 랑 `[20230927]제이타운_매매관련01.m4a` 가 보여야 해. 

## **4) 실행 (로컬 엔진)** 

bat 코드 복사 `python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --segment 60 --`   잘 되면 다음엔 VAD도 켜서: bat 코드 복사 `python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --vad --engine`   **만약 또 “파일 없음”이 뜨면 항상 폴더 먼저 이동** 후 실행하거나, **풀 경로** 로 실행: bat 코드 복사 `python "%USERPROFILE%\OneDrive\Desktop\Voice script\voice_script.py" "%USERPROFIL`   **만약 “ModuleNotFoundError: faster_whisper”가 뜨면** bat 코드 복사 

잘 되면 다음엔 VAD도 켜서: 

## **만약 또 “파일 없음”이 뜨면** 

## **만약 “ModuleNotFoundError: faster_whisper”가 뜨면** 

```
pip install faster-whisper
```

한 번 돌려보고, 마지막에 나온 **메시지 10줄** 만 붙여줘. 거기서 바로 이어서 잡아줄게! 

**==> picture [489 x 39] intentionally omitted <==**

[전처리] [20230927]제이타운_매매관련01.m4a → [20230927]제이타운_매매관련01_clean.wav (강모드) 로컬 변환용 faster-whisper 설치 필요: pip install faster- 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68fe515d-dbb8-8323-8992-21c0c50c0e9b 

127/175 

25. 11. 30. 오후 1:03 

프로그램[파이썬/코딩] - [Voice] Script 생성 

whisper 

 

## C:\Users\psm23\OneDrive\Desktop\Voice script 

## 10s 동안 생각함 

형, 딱 이것만 하면 돼. (로컬 엔진 설치가 안 돼서 멈춘 거라 `faster-whisper` 만 깔면 끝) 

## **1) 로컬 엔진 설치 (CPU 버전)** 

bat 코드 복사 `pip install faster-whisper pip install torch --index-url https://download.pytorch.org/whl/cpu` 

## **2) 설치 확인 (1초 테스트)** 

bat 코드 복사 

```
python -c "from faster_whisper import WhisperModel; print('FW OK')"
```

`FW OK` 뜨면 준비 완료. 

## **3) 다시 실행 (로컬 엔진)** 

bat 코드 복사 `python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --segment 60 --`   잘 되면 다음엔 대화만 추출도 켜서: bat 코드 복사 `python voice_script.py "[20230927]제이타운_매매관련01.m4a" --strong --vad --engine`   **참고(선택)** 느리면 `--modelsize base` 더 정확히 원하면 `--modelsize medium` (조금 느림) 

## **에러 나면** 