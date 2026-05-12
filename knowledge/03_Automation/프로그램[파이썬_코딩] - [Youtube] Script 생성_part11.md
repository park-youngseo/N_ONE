
[사실] 

## **3) 변환(자막 → 텍스트) — 파일명 자동 탐색 버전** 

폴더에 생긴 **아무 .srt** 나 자동으로 찾아서 **output.txt** 로 바꿔준다: 

bat 코드 복사 

```
python -c "import sys,glob,re; sys.stdout.reconfigure(encoding='utf-8'); f=(glob.
```

  끝나면 **같은 폴더에** `output.txt` 생성. [사실] 

## **왜 실패했나? (짧게)** 

- `FileNotFoundError` = **명령에서 지목한** `CF4Fb9vkrik.ko.srt` **파일이 실제로 없음** . 

- 먼저 **자막 파일부터 받아야** 변환이 된다. [사실] 

막히면: 

`dir *.srt` 결과, 

- `yt-dlp …` 실행 후 나오는 메시지 

- 이 두 개를 **그대로 붙여줘** . 거기 맞춰 바로 다음 스텝 해줄게, 형. 

**==> picture [489 x 38] intentionally omitted <==**

vtt파일로 생성되는데? 

36s 동안 생각함 

형, 이건 **정상** 이야. `yt-dlp` 가 **자막 변환에** `ffmpeg` **를 쓰는데** , PC에 `ffmpeg` 가 없으면 **.vtt** 그대로 남겨둔다. 

→ 방법 두 개 중 하나 택하면 끝. 

## **A단계｜핵심** 

**원인:** ffmpeg 미설치라서 `.srt` 변환이 안 됨. [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

55/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **해결:** 

- ① ffmpeg 설치 후 다시 받기(자동 `.srt` ) **또는** 

- ② 지금 생긴 `.vtt` 를 **바로** `.srt` **/텍스트로 변환** . [사실] 

## **B2단계｜손잡고 따라하기** 

## `.srt` **방법 ① (권장) ffmpeg 설치 → 자동** 

**1.** (CMD에) 우선 ffmpeg 설치 대신, 지금은 **설치 없이도 바로 끝내는 방법(방법 ②)** 를 먼저 써도 됨. ffmpeg 설치가 번거롭다면 **방법 ②** 로 바로 진행. [사실] 

**2.** ffmpeg 설치 후엔 아래처럼 다시 받으면 자동으로 `.srt` 생성됨: 

vbnet 코드 복사 

```
yt-dlp --write-auto-sub --sub-lang "ko,en" --convert-subs srt --skip-download
```

  

- → 같은 폴더에 `*.ko.srt` 또는 `*.en.srt` 생김. [사실] 

## **방법 ② (설치 없이 즉시 끝)** `.vtt` **→** `.srt` **& 텍스트** 

## **지금 폴더** (vtt가 있는 곳) CMD에서 그대로 복붙. 

## **1) 폴더에 있는 모든 .vtt → .srt로 변환** 

swift 

코드 복사 

```
python -c "import re,glob,sys; sys.stdout.reconfigure(encoding='utf-8');
files=glob.glob('*.vtt');
```

```
assert files,'no .vtt files';
```

```
for v in files:
```

```
    t=open(v,'r',encoding='utf-8',errors='ignore').read()
```

```
    t=re.sub(r'^\ufeff?WEBVTT.*?\n','',t,flags=re.S)    # 헤더제거
```

```
    blocks=[b for b in re.split(r'\n\n+', t) if '-->' in b]
    out=[]; i=1
```

```
    for b in blocks:
```

```
        lines=[ln for ln in b.strip().splitlines() if ln and not ln.startswith('NO
```

```
        ts=lines[0].replace('.',',')                     # 타임스탬프형식 VTT→SRT
        text=' '.join([re.sub('<.*?>','',ln).strip() for ln in lines[1:]])
        out += [str(i), ts, text, '']; i+=1
```

```
    open(v.rsplit('.',1)[0]+'.srt','w',encoding='utf-8').write('\n'.join(out))"
```

 

 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

56/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

→ 끝나면 같은 폴더에 **같은 이름의** `.srt` 가 생김. [사실] 

## **2)** `.srt` **나** `.vtt` **에서 대사만 뽑아** `output.txt` **만들기** 

swift 

코드 복사 

```
python -c "import glob,re,sys; sys.stdout.reconfigure(encoding='utf-8');
f=(glob.glob('*.srt') or glob.glob('*.vtt'))[0]
```

```
t=open(f,'r',encoding='utf-8',errors='ignore').read()
```

```
print('\n'.join([re.sub(r'<.*?>','',ln).strip() for ln in t.splitlines() if ln.st
```

  

- → 같은 폴더에 **output.txt** 생성. [사실] 

## **어느 방법이 좋나?** 

- **빨리 끝내기** : 방법 ② (지금 파일로 즉시 변환) [사실] 

- **앞으로 계속 자동 srt** : 방법 ① (ffmpeg 설치 후 `--convert-subs srt` ) [사실] 

필요하면 **ffmpeg 설치** 도 단계별로 잡아줄게. 지금은 위 ②번 두 줄만 실행해보고, 결과 (생성된 파일명) 알려줘, 형. 

**==> picture [489 x 38] intentionally omitted <==**

돌아버리겠네 씨발 새끼 몇번을 해야해? 노가다를 

33s 동안 생각함 

형, 미안. 진짜 **노가다 없이 끝** 내는 “3줄 완성”으로 간다. 그대로 복붙만 해. 

## **A단계｜핵심** 

- **원인:** 자막 변환에 ffmpeg/권한 이슈가 섞여서 막힘. [사실] 

- **해결:** ffmpeg 설치 → 자막 받기 → 텍스트 추출, **3줄** . [사실] 

## **B2｜딱 3줄 (Windows CMD에서 순서대로)** 

## **1. ffmpeg + yt-dlp 설치** (한 번만) 

css 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

57/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
winget install -e --id Gyan.FFmpeg & python -m pip install -q--upgrade pip yt-dl
```

  **2. 자막 받기** (쿠키 없이 먼저; URL만 바꿔서 붙여넣기) arduino 코드 복사 `yt-dlp --write-auto-subs --sub-langs "ko,en" --convert-subs srt --skip-download ht`   실패하거나 **권한/연령/지역 제한** 이면, 크롬 확장으로 `cookies.txt` 만든 뒤 **같은 폴 더에 놓고** 아래처럼: css 코드 복사 `yt-dlp --cookies cookies.txt --write-auto-subs --sub-langs "ko,en" --convert-subs`   [사실] **3. 텍스트만 추출 → output.txt** (파일명 자동 탐색) lua 코드 복사 

## [사실] 

`python -c "import glob,re,sys; sys.stdout.reconfigure(encoding='utf-8'); f=(glob.g`   

## 끝. **같은 폴더에** `output.txt` 생긴다. [사실] 

## **초간단 체크** 

   - `*.srt` 가 안 생기면 → ②에서 `--cookies cookies.txt` 버전으로 다시. [사실] 

   - 여전히 실패면 그 영상은 **공개 자막 없음/제한 강함** → **Whisper(음성 인식)** 로 바로 가자고 하면, 내가 명령 2줄로 정리해서 붙여줄게. [추론] 

- 딱 위 3줄만 해보고, 에러 나면 **그 줄 출력 그대로** 복사해서 보내줘. 바로 거기서 잡아줄 게, 형. [사실] 

**==> picture [489 x 39] intentionally omitted <==**

C:\Users\엔오원\Desktop\youtube_script_extractor>ytdlp --write-auto-subs --sub-langs "ko,en" --convert-subs srt --skip-download https://youtu.be/CF4Fb9vkrik 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

58/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

[youtube] Extracting URL: https://youtu.be/CF4Fb9vkrik 

[youtube] CF4Fb9vkrik: Downloading webpage 

[youtube] CF4Fb9vkrik: Downloading tv client config 

[youtube] CF4Fb9vkrik: Downloading tv player API JSON [youtube] CF4Fb9vkrik: Downloading web safari player API JSON 

[youtube] CF4Fb9vkrik: Downloading m3u8 information [info] CF4Fb9vkrik: Downloading subtitles: ko, en WARNING: ffmpeg not found. The downloaded format may not be the best available. Installing ffmpeg is strongly recommended: https://github.com/yt-dlp/ytdlp#dependencies 

[info] CF4Fb9vkrik: Downloading 1 format(s): 96 [info] Writing video subtitles to: AI 코딩 끝판왕 커서 AI 사용법 - 설치부터 상세하게 (비개발자 모두 주목!) [CF4Fb9vkrik].ko.vtt 

WARNING: The extractor specified to use impersonation for this download, but no impersonate target is available. If you encounter errors, then see 

https://github.com/yt-dlp/yt-dlp#impersonation  for information on installing the required dependencies [download] Destination: AI 코딩 끝판왕 커서 AI 사용법 - 설치부터 상세하게 (비개발자 모두 주목!) [CF4Fb9vkrik].ko.vtt 

[download] 100% of   93.41KiB in 00:00:00 at 870.96KiB/s 

[info] Writing video subtitles to: AI 코딩 끝판왕 커서 AI 사용법 - 설치부터 상세하게 (비개발자 모두 주목!) [CF4Fb9vkrik].en.vtt 

WARNING: The extractor specified to use impersonation for this download, but no impersonate target is available. If you encounter errors, then see https://github.com/yt-dlp/yt-dlp#impersonation  for information on installing the required dependencies ERROR: Unable to download video subtitles for 'en': HTTP Error 429: Too Many Requests 

형, 잘 됐어. **한글 자막(.ko.vtt)은 이미 받아졌고** , 영어(en)는 429(요청 과다)라 실패했지만 **한국어만으로도 바로 텍스트 추출 가능** 해. (PARTICIPATION=Off / EVALUATION=Off) 

## **A단계｜핵심** 

지금 폴더에 있는 `.ko.vtt` **를 바로** `output.txt` **로 변환** 하면 끝. [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

59/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

429 경고는 **영어 자막만 실패** 한 것이고, 한글 파일은 이미 내려받았으니 무시해도 됨. [사실] 

## **B2단계｜딱 2줄 (CMD에 그대로 붙여넣기)** 

## **1. .vtt → 텍스트(output.txt)** 한 방에: 

swift 코드 복사 `python -c "import glob,re,sys; sys.stdout.reconfigure(encoding='utf-8'); f=(glob.g`   → 같은 폴더에 `output.txt` 생기면 완료. [사실] **2.** (선택) 확인: lua 코드 복사 

```
notepad output.txt
```

— 확장(옵션) — 

앞으로 **자동 srt** 가 필요하면 ffmpeg 설치 후, yt-dlp가 **.srt로 변환** 해주게 하면 편함: 

css 코드 복사 
