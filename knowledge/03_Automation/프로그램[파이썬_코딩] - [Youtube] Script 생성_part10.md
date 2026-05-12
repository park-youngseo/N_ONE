
- → 같은 폴더에 `CF4Fb9vkrik.mp3` 생성. [사실] 

**3.** Whisper로 텍스트화(빠른 모델): 

python 코드 복사 

```
python - <<#PY
from faster_whisper import WhisperModel
model = WhisperModel("base")  # 더빠른건 "tiny", 더정확한건 "small/medium"
segments, info = model.transcribe("CF4Fb9vkrik.mp3", language="ko")  # 모르면 lang
withopen("output.txt","w",encoding="utf-8") as f:
for seg in segments:
        f.write(seg.text.strip()+"\n")
print("완료! output.txt 저장")
#PY
```

  

- → `output.txt` 생성. (속도는 PC 성능에 비례) [사실] 

## **추가 팁(선택)** 

- **재시도 로직** 넣으면 간헐 네트워크 에러에 강해져. 필요하면 코드를 3회 재시도하도 록 바꿔줄게. [추론] 

- 회사/학원 **방화벽·VPN** 이 있으면, 잠깐 끄고 재실행하거나 집/핫스팟에서 테스트해 봐. [추론] 

## 형, 지금은 **1) 테스트 영상** 먼저 해보고 결과 알려줘. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

50/341 

25. 11. 30. 오후 1:07 

 

   - 프로그램[파이썬/코딩] - [Youtube] Script 생성 

- 테스트 영상 **성공** → 대상 영상은 자막 없음/제한 ⇒ **쿠키 or Whisper** 로 가자. [추론] 테스트 영상도 **실패** → 내가 말한 **루트 A(쿠키)** 그대로 적용해 보고 결과를 보내줘. 바로 다음 스텝 이어서 마무리하자. 

**==> picture [489 x 38] intentionally omitted <==**

테스트 영상도 실패 

8s 동안 생각함 

- 형, **테스트 영상도 실패** 면 거의 확실히 

- ① 로그인·연령·지역 제한/쿠키 필요 **또는** ② 네트워크(방화벽/VPN) 영향이야. [사실] 

아래 둘 중 하나로 끝내자. **가장 쉬운 루트 A(권장)** → 그래도 안 되면 루트 B. 

## **A단계｜전체 개요** 

**루트 A(권장): yt-dlp로 자막 파일 뽑기 → 필요시 쿠키로 우회.** [사실] **루트 B: 파이썬 코드에 쿠키 적용(같은 원리).** [사실] 

## **B1｜핵심만 (명령 3줄)** 

**1.** `python -m pip install yt-dlp` 

**2.** (쿠키 없이 먼저) 

코드 복사 

vbnet 코드 복사 `yt-dlp --write-auto-sub --sub-lang "ko,en" --convert-subs srt --skip-download`   **3.** 같은 폴더에 `*.ko.srt` 또는 `*.en.srt` 생기면 성공. [사실] 

## **B2｜손잡고 따라하기 (루트 A 권장)** 

## **0) 준비(한 번만)** 

css 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

51/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m pip install --upgrade pip
```

```
python -m pip install yt-dlp
```

## **1) 쿠키 없이 먼저 시도** 

arduino 

코드 복사 

```
yt-dlp --write-auto-sub --sub-lang "ko,en" --convert-subs srt --skip-download htt
```

  

- 성공: 폴더에 `CF4Fb9vkrik.ko.srt` (또는 `.en.srt` ) 생김 → 끝. [사실] 실패(자막/권한 오류): **2) 쿠키 방식** 으로 진행. [사실] 

## **2) 쿠키 파일 만들기(초간단)** 

**1. 크롬 로그인 상태** 로 유튜브 영상 페이지를 연다. 

**2.** 크롬 웹스토어에서 **“Get cookies.txt(쿠키 내보내기)”** 같은 확장 설치. 

**3.** 영상 페이지에서 **확장 아이콘 클릭 → Export →** `cookies.txt` **저장** . 

**4.** 그 `cookies.txt` 를 **작업 폴더(youtube_script_extractor)에 놓기** . [사실] 

## **3) 쿠키로 다시 받기** 

css 

코드 복사 

```
yt-dlp --cookies cookies.txt--write-auto-sub--sub-lang "ko,en" --convert-subs s
```

  성공: `*.srt` 생성. 원하면 텍스트로도 바로 빼기(간단 변환): lua 코드 복사 

```
python -c "import re,sys;import io;sys.stdout.reconfigure(encoding='utf-8');t
```

**==> picture [449 x 8] intentionally omitted <==**

**----- Start of picture text -----**<br>
 <br>**----- End of picture text -----**<br>


(파일명이 다르면 그 이름으로 바꿔서 입력) [사실] 

## **루트 B — 파이썬 코드로 “쿠키 적용” (원하면 이걸로)** 

**1.** `cookies.txt` 를 **동일 폴더** 에 둔다. 

**2.** `extract_text_only.py` 의 **자막 요청 줄** 을 아래처럼 바꾼다: 

python 

**==> picture [68 x 19] intentionally omitted <==**

**----- Start of picture text -----**<br>
코드 복사<br>**----- End of picture text -----**<br>


https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

52/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# 기존
```

```
data = YouTubeTranscriptApi.get_transcript(vid, languages=[lang])
# 변경(쿠키적용)
```

```
data = YouTubeTranscriptApi.get_transcript(vid, languages=[lang], cookies='cookie
```

  

(두 군데 있으면 두 군데 모두 적용) 

3) 실행: 

nginx 코드 복사 `python extract_text_only.py` 

→ `output.txt` 생성. [사실] 

## **추가 체크(문제 지속 시)** 

- **VPN/방화벽** 이 있으면 끄고 다시 시도(유튜브 응답이 비어오는 경우가 있음). [추론] **연령·멤버십·지역 제한** 영상은 **쿠키 없이는 실패** 가 정상. [사실] 

- 여전히 안 되면 **Whisper(음성 인식)** 로 전환하면 자막 없어도 추출 가능: 

scss 코드 복사 

```
python -m pip install yt-dlp faster-whisper
yt-dlp -x --audio-format mp3 -o "%(id)s.%(ext)s" <URL>
```

그 뒤 음성 → 텍스트 변환 스크립트 안내해줄게. [사실] 

## 형, 바로 **루트 A의 0→1→(필요시)2→3 순서** 로 해봐. 

생성된 파일명( `.srt` 또는 `output.txt` )이나 에러 메시지를 그대로 보내주면, 그 상태에서 다음 단계 이어서 딱 마무리하자. 

**==> picture [489 x 38] intentionally omitted <==**

## C:\Users\엔오원 

\Desktop\youtube_script_extractor>python -c "import re,sys;import io;sys.stdout.reconfigure(encoding='utf8');t=open(r'CF4Fb9vkrik.ko.srt','r',encoding='utf- 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

53/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

8',errors='ignore').read();print('\n'.join([re.sub(r'^\d+$|\ d\d:\d\d:\d\d,\d{3} --> .*','',line).strip() for line in t.splitlines() if line.strip() and '-->' not in line and not line.strip().isdigit()]))" > output.txt 

Traceback (most recent call last): File "<string>", line 1, in <module> 

FileNotFoundError: [Errno 2] No such file or directory: 'CF4Fb9vkrik.ko.srt' 

C:\Users\엔오원\Desktop\youtube_script_extractor> 

15s 동안 생각함 

형, **파일이 없어서** 저 변환 명령이 실패한 거야. 즉, `CF4Fb9vkrik.ko.srt` 가 폴더에 안 만 들어진 상태. [사실] 

## **A단계｜핵심** 

먼저 **자막 파일(.srt)을 실제로 받아오고** , 그다음 **텍스트로 변환** 해야 해. [사실] 

## **B2단계｜손잡고 해결 (그대로 복붙)** 

## **1) 지금 폴더에 자막이 있는지 확인** 

bat 

코드 복사 

```
dir *.srt
```

아무것도 안 나오면 **아직 자막이 없는 상태** . [사실] 

## **2) 자막 받기 (쿠키 없이 먼저)** 

bat 

코드 복사 

```
python -m pip install yt-dlp
```

```
yt-dlp --write-auto-sub --sub-lang "ko,en" --convert-subs srt --skip-download htt
```

-  

- 성공하면 `*.ko.srt` 또는 `*.en.srt` 가 생김. [사실] 실패(권한/연령/지역 등)하면 **쿠키 방식으로** : 

   **1.** 크롬에서 해당 영상 페이지 열기(로그인 상태) 

   **2.** 확장프로그램 “cookies.txt” 내보내기 설치 → **cookies.txt** 저장 

   **3.** 그 파일을 **현재 폴더** 에 복사 

   **4.** 다시 실행: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

54/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

코드 복사 

bat 

```
yt-dlp --cookies cookies.txt --write-auto-sub --sub-lang "ko,en" --conver
```

  