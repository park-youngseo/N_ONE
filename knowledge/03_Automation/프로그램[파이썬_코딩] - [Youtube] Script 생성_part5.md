**3. 실행** : `python extract_simple.py` → 유튜브 주소 붙여넣기. [사실] 

## **B2｜손잡고 따라하기 (Windows, 순서대로)** 

## **0) 폴더 준비** 

`youtube_script_extractor` 폴더를 **바탕화면/바로 가기 등 편한 곳** 에 둔다. [사실] 

## **1) 코드 파일 만들기** 

**1. 메모장** 열기. 

**2.** 형이 보낸 **코드 전체** 를 그대로 **복사 → 메모장에 붙여넣기** . 

**3. 파일 > 다른 이름으로 저장…** 클릭. 

**4.** 저장 위치를 `youtube_script_extractor` **폴더** 로 지정. 

**5.** 파일 이름에 `extract_simple.py` 입력. 

**6.** (중요) **파일 형식** 을 “모든 파일(.)”로 바꾸고, **인코딩은 UTF-8** 선택. 

**7. 저장** 클릭. 

   - → 폴더 안에 `extract_simple.py` 가 생기면 성공. [사실] 

## **2) “검은 창(CMD)”를 그 폴더에서 열기** 

**1.** `youtube_script_extractor` 폴더 창에서, **맨 위 주소창** (경로 보이는 칸)을 **한 번 클릭** . 

**2.** 키보드로 `cmd` 입력 → **Enter** . 

   - → 검은 창이 뜨고, 줄 앞에 폴더 경로가 보이면 성공. [사실] 

      - 붙여넣기는 **마우스 우클릭** 으로 가능. 

## **3) 필요한 패키지 설치(처음 1번만)** 

검은 창에 아래 **그대로 복사→붙여넣기→Enter** : 

nginx 코드 복사 

```
python -m pip install youtube-transcript-api
```

## 설치가 쭉 진행되면 **성공** . [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 23/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

만약 글자가 빨간색으로 오류가 보이면, 그 **문구 전체를 복사해서** 보내줘. [사실] 

## **4) 프로그램 실행** 

검은 창에서 아래 **입력 후 Enter** : 

nginx 코드 복사 

```
python extract_simple.py
```

화면에 **“유튜브 URL을 입력하세요:”** 라고 뜨면, 

유튜브 주소를 **붙여넣고 Enter** . 예) `https://youtu.be/CF4Fb9vkrik?...` [사실] 

## **5) 결과 확인** 

같은 폴더에 `output.txt` (대사만)와 `output.srt` (자막 파일) 생성되면 완료. [사실] 

## **(딱 필요한 만큼) 문제 해결 빠른표** 

`python` 이 안 된다 → 파이썬 설치/경로 문제. **새 창에서** `python -V` 가 버전 나오게 먼저 잡기. [사실] 

`ModuleNotFoundError: youtube_transcript_api` → 3단계(설치) 다시: 

nginx 코드 복사 

```
python -m pip install youtube-transcript-api
```

- **“자막을 찾지 못했습니다”** → 그 영상은 공개 캡션이 없음. 

- → 나중에 **Whisper(음성 인식)** 방식이 필요. 우선 다른 영상으로 테스트해봐. [사실] 

- `>>>` 가 보이는 창에 명령이 안 먹는다 → 그건 **파이썬 창** . `exit()` 후 닫고 **반드시 검** 

- **은 창에서** 명령 입력. [사실] 

원하면 이 코드를 **자동으로 SRT까지** 만드는 내 “안전판 버전(v1.1)”으로도 바로 교체해줄 수 있어. 

지금은 위 순서 그대로 해보고, 화면에 뜬 글자를 복사해 주면 다음 단계 바로 이어갈게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

텍스트로 그냥저장해? 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

24/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

형, 맞아. **메모장에 “파이썬 코드”를 텍스트로 저장** 하면 돼. 다만 **파일 확장자만** `.py` 로 저장해야 실행된다. 그리고 결과(스크립트)는 `output.txt` 로 텍스트로 저장되게 할 수 있 어. 

## **A단계｜핵심 한 줄** 

메모장에 코드 붙여넣고 `extract_simple.py` 로 저장 → 실행하면 `output.txt` (텍스 트만) 만들어진다. [사실] 

## **B1｜형이 원하는 “텍스트만 저장” 최소 버전** 

아래 코드 그대로 메모장에 붙여넣고 `extract_text_only.py` 로 저장해. 

python 코드 복사 

```
from youtube_transcript_api import YouTubeTranscriptApi, NoTranscriptFound, Trans
from urllib.parse import urlparse, parse_qs
```

```
defextract_video_id(url: str) -> str:
    q = urlparse(url.strip())
    host = (q.hostname or"").lower() if q.hostname else""
if host == "youtu.be":
return q.path.lstrip("/")
if host in ("www.youtube.com", "youtube.com", "m.youtube.com"):
if q.path == "/watch":
return parse_qs(q.query).get("v", [""])[0]
if q.path.startswith("/embed/") or q.path.startswith("/v/"):
return q.path.split("/")[2]
# 그냥 ID만넣은경우도허용
return url.strip()
```

```
if __name__ == "__main__":
    url = input("유튜브 URL 또는영상 ID: ").strip()
try:
        vid = extract_video_id(url)
# 한글우선, 없으면영어
        data = None
for lang in ["ko", "en"]:
try:
                data = YouTubeTranscriptApi.get_transcript(vid, languages=[lang])
break
except NoTranscriptFound:
continue
if data isNone:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 25/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
raise NoTranscriptFound("공개자막없음")
        text = "\n".join([(d.get("text") or"").replace("\n"," ").strip() for d i
withopen("output.txt", "w", encoding="utf-8") as f:
            f.write(text)
print("완료! 같은폴더에 output.txt 가저장되었습니다.")
except TranscriptsDisabled:
print("이영상은자막이비활성화되어있습니다.")
except VideoUnavailable:
print("영상을찾을수없습니다.")
except NoTranscriptFound as e:
print("자막이없어텍스트를만들수없습니다.")
except Exception as e:
print("오류:", e)
```

## **B2｜저장·실행 순서 (초등학생 코스)** 

**1. 메모장 열기** → 위 코드 **전부 복사** 해 붙여넣기. [사실] 

**2. 파일 > 다른 이름으로 저장…** 

파일 이름: `extract_text_only.py` 

파일 형식: **모든 파일(.)** 

- 인코딩: **UTF-8** 

저장 위치: 작업 폴더(예: `youtube_script_extractor` ) [사실] 

**3. 검은 창(CMD) 열기** : 폴더 창의 **주소창에** `cmd` **입력 → Enter** . [사실] 

**4. 필수 설치(처음 1번)** : 

nginx 코드 복사 

```
python -m pip install youtube-transcript-api
```

(설치 글자가 쭉 나오면 성공) [사실] 

## **5. 실행** : 

nginx 코드 복사 `python extract_text_only.py` 

## → 물으면 **유튜브 주소 붙여넣기** → Enter. [사실] 

**6.** 같은 폴더에 `output.txt` 생기면 끝. [사실] 

## — 확장(옵션) — 

- 파일명을 영상 제목으로 저장하고 싶으면 말해줘. 자동으로 `제목.txt` 로 저장되게 딱 맞춰 줄게. [추론] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

26/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

Microsoft Windows [Version 10.0.19045.6456] 

(c) Microsoft Corporation. All rights reserved. 

## C:\Users\엔오원 

\Desktop\youtube_script_extractor>python -m pip install youtube-transcript-api 

Requirement already satisfied: youtube-transcript-api in c:\users\엔오원 

\appdata\local\programs\python\python311\lib\sit e-packages (1.2.3) 

Requirement already satisfied: defusedxml<0.8.0,>=0.7.1 in c:\users\엔오원 

\appdata\local\programs\python\python311\lib\sit e-packages (from youtube-transcript-api) (0.7.1) 

Requirement already satisfied: requests in c:\users\엔오 원 

\appdata\local\programs\python\python311\lib\sit 

e-packages (from youtube-transcript-api) (2.32.5) 

Requirement already satisfied: charset_normalizer<4,>=2 in c:\users\엔오원 

\appdata\local\programs\python\python311\lib\sit 

e-packages (from requests->youtube-transcript-api) (3.4.4) 

Requirement already satisfied: idna<4,>=2.5 in 

c:\users\엔오원 

\appdata\local\programs\python\python311\lib\sit 

e-packages (from requests->youtube-transcript-api) (3.11) 

Requirement already satisfied: urllib3<3,>=1.21.1 in c:\users\엔오원 

\appdata\local\programs\python\python311\lib\sit e-packages (from requests->youtube-transcript-api) (2.5.0) 

Requirement already satisfied: certifi>=2017.4.17 in c:\users\엔오원 

\appdata\local\programs\python\python311\lib\sit e-packages (from requests->youtube-transcript-api) (2025.10.5) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 27/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

[notice] A new release of pip is available: 24.0 -> 25.2 [notice] To update, run: python.exe -m pip install -- upgrade pip 

C:\Users\엔오원 

\Desktop\youtube_script_extractor>python 파일이름.py python: can't open file 'C:\\Users\\엔오원 \\Desktop\\youtube_script_extractor\\파일이름.py': [Errno 2] No such file or directory 

C:\Users\엔오원\Desktop\youtube_script_extractor> 

형, 지금은 **“파일이름.py”** 라고 예시 문구를 그대로 쳐서 생긴 오류야. 

→ **실제 파일 이름** 으로 실행해야 해. (예: `extract_transcript.py` 또는 `extract_text_only.py` ) [사실] 

## **A단계｜핵심** 
