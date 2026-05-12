
형, 진짜 **초등학생 코스** 로 간다. 한 줄씩 그대로 따라만 해. 

## **A단계｜전체 개요** 

- **목표:** 유튜브 주소 붙여넣고 **텍스트/자막 파일** 받기. [사실] 

- **핵심 3줄:** `폴더에서 CMD 열기 → 설치 1줄 → 실행 1줄` 끝. [사실] 

## **B2단계｜초간단 따라하기(Windows)** 

## **0) 준비** 

**1. 열려 있는 파이썬 창** (>>> 보이는 창)은 닫기 → `exit()` 입력 후 Enter. [사실] 

**2.** 압축 풀었다면, 폴더 이름이 `youtube_script_extractor` 인지 확인. [사실] 

## **1) 폴더에서 “검은 창(CMD)” 열기** 

**1.** `youtube_script_extractor` 폴더를 **더블클릭** 해서 연다. 

**2. 창 맨 위 주소창** (경로 보이는 칸)을 **한 번 클릭** 한다. 

**3.** 키보드로 `cmd` 라고 치고 **Enter** . → 검은 창이 뜨면 **성공** . [사실] 

## **2) 설치(처음 1번만)** 

## 검은 창에 **아래 줄을 그대로 복사→붙여넣기→Enter** : 

nginx 코드 복사 

```
python -m pip install -r requirements.txt
```

## (끝날 때까지 기다리기) [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

17/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **3) 실행(매번 이거 하면 됨)** 

방법 A — 더 쉬운 방법: 

arduino 코드 복사 

```
run.bat
```

→ “영상 URL 또는 영상 ID”가 뜨면 **유튜브 주소를 붙여넣고** Enter. [사실] 

방법 B — 직접 실행(같은 결과): 

nginx 코드 복사 `python extract_transcript.py` 

→ 주소 붙여넣고 Enter. [사실] 

## **4) 결과 확인** 

같은 폴더에 `outputs` 라는 폴더가 생김. 

- 그 안에 

`.txt` = 대사만 모은 텍스트 

- `.srt` = 시간표시가 있는 자막 파일 

- 이 두 개가 만들어져 있으면 **성공** . [사실] 

## **자주 막히는 딱 두 가지** 

- **검은 창이 아니라 파이썬(>>>) 창이 뜬다** → 그건 파이썬 창. 닫고 **다시 1단계로** . [사 실] 

- **에러 문구가 나온다** → 창에 나온 글자를 **그대로 복사해서 보내줘** . 바로 해결해줄게. [사실] 

필요하면 최신 패치 버전 압축 다시 받을 수 있어: 

▶ youtube_script_extractor_v1_1.zip [사실] 

**==> picture [489 x 39] intentionally omitted <==**

from youtube_transcript_api import YouTubeTranscriptApi from urllib.parse import urlparse, parse_qs 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

18/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

def extract_video_id(url): # 유튜브 URL에서 영상 ID 추출 query = urlparse(url) if query.hostname == 'youtu.be': return query.path[1:] if query.hostname in ['www.youtube.com', 'youtube.com']: if query.path == '/watch': return parse_qs(query.query)['v'][0] elif query.path[:7] == '/embed/': return query.path.split('/')[2] elif query.path[:3] == '/v/': return query.path.split('/')[2] # 지원하지 않는 링크 raise ValueError('유효한 유튜브 URL이 아닙니다.') 

def get_transcript(youtube_url): try: video_id = extract_video_id(youtube_url) transcript = YouTubeTranscriptApi.get_transcript(video_id, languages= ['ko', 'en']) script = ' '.join([entry['text'] for entry in transcript]) return script except Exception as e: return f"스크립트 추출 실패: {e}" 

if __name__ == "__main__": youtube_url = input("유튜브 URL을 입력하세요: ") script = get_transcript(youtube_url) print("\n=== 스크립트 내용 ===\n") print(script) 

형, 잘했어. 이 코드는 “URL→영상ID 뽑기 → 자막 받아오기 → 한 줄 텍스트로 합치기”를 해. 

딱 **설치 1줄 + 실행 1줄** 만 하면 돌아가. 

## **A단계｜먼저 해야 할 2줄** 

**1.** (처음 1번) 패키지 설치: 

nginx 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

19/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m pip install youtube-transcript-api
```

## **2.** 실행: 

nginx 코드 복사 

```
python파일이름.py
```

→ 뜨는 물음에 유튜브 URL 붙여넣기 → 엔터. [사실] 

## **B1｜지금 코드로 충분히 되는 이유(핵심)** 

`extract_video_id()` 가 youtu.be / watch / embed / v 경로를 처리. [사실] 

`YouTubeTranscriptApi.get_transcript(..., languages=['ko','en'])` 로 **한글 우선, 없으면 영어** 자막을 요청. [사실] 

## **B2｜왕초보 안전판(오류 줄인 개선본 + 파일저장)** 

아래로 갈아끼우면, **자막 없을 때 안내** , **SRT 파일** 까지 같이 저장돼. 

python 코드 복사 

```
from youtube_transcript_api import YouTubeTranscriptApi, NoTranscriptFound, Trans
from urllib.parse import urlparse, parse_qs
```

```
defextract_video_id(url: str) -> str:
    q = urlparse(url.strip())
    host = (q.hostname or"").lower()
if host == 'youtu.be':
return q.path.lstrip('/')
if host in ('www.youtube.com', 'youtube.com', 'm.youtube.com'):
if q.path == '/watch':
return parse_qs(q.query).get('v', [''])[0]
if q.path.startswith('/embed/') or q.path.startswith('/v/'):
return q.path.split('/')[2]
# 그냥 ID만넣은경우도허용
iflen(url.strip()) >= 10and'/'notin url:
return url.strip()
raise ValueError('유효한유튜브 URL/ID가아닙니다.')
```

```
deffetch_transcript(video_id: str):
# 1) ko → en 우선시도
for lang in ['ko', 'en']:
try:
return YouTubeTranscriptApi.get_transcript(video_id, languages=[lang]
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

20/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
except NoTranscriptFound:
continue
# 2) (있으면) 다른언어자막 → 한국어번역시도
try:
        list_fn = getattr(YouTubeTranscriptApi, "list_transcripts")
        transcripts = list_fn(video_id)
# 번역우선
for tr in transcripts:
try:
return tr.translate('ko').fetch()
except Exception:
pass
# 번역실패시원문
for tr in transcripts:
try:
return tr.fetch()
except Exception:
pass
except Exception:
pass
raise NoTranscriptFound("사용가능한자막이없습니다.")
```

```
defto_srt(items):
defts(sec):
        ms = int((sec - int(sec)) * 1000)
        s = int(sec) % 60
        m = (int(sec) // 60) % 60
        h = int(sec) // 3600
returnf"{h:02d}:{m:02d}:{s:02d},{ms:03d}"
    lines = []
for i, it inenumerate(items, start=1):
        start = float(it.get("start", 0.0))
        end = start + float(it.get("duration", 0.0))
        text = (it.get("text") or"").replace("\n", " ").strip()
        lines += [str(i), f"{ts(start)} --> {ts(end)}", text, ""]
return"\n".join(lines)
if __name__ == "__main__":
    url = input("유튜브 URL 또는영상 ID: ").strip()
try:
        vid = extract_video_id(url)
        data = fetch_transcript(vid)
# TXT
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 21/341 

25. 11. 30. 오후 1:07 

후 1:07 1:07 프로그램[파이썬/코딩] - [Youtube] Script 생성 `txt = "\n".join([(d.get("text") or "").replace("\n"," ").strip() for d in`  `with open("output.txt", "w", encoding="utf-8") as f: f.write(txt)` _`# SRT`_ `with open("output.srt", "w", encoding="utf-8") as f: f.write(to_srt(data)) print("완료! output.txt / output.srt 파일이 생성되었습니다.") except TranscriptsDisabled: print("이 영상은 자막이 비활성화되어 있습니다. (Whisper 같은 음성 인식이 필요) except VideoUnavailable: print("영상을 찾을 수 없습니다. URL/ID를 확인하세요.") except NoTranscriptFound as e: print("자막을 찾지 못했습니다. (공개 캡션 없음일 수 있음)") except Exception as e: print("스크립트 추출 실패:", e)` 

## **사용법(초간단 반복)** 

**1.** 폴더에서 주소창 클릭 → `cmd` → 엔터. 

**2.** 설치(처음만): 

nginx 코드 복사 

```
python -m pip install youtube-transcript-api
```

**3.** 실행: 

nginx 코드 복사 

```
python파일이름.py
```

**4.** 유튜브 주소 붙여넣기 → **같은 폴더에** `output.txt` **,** `output.srt` 생기면 성공. [사실] 

막히면 **화면에 나온 글자 그대로** 붙여줘. 바로 이어서 잡아줄게, 형. 

**==> picture [489 x 38] intentionally omitted <==**

이걸 적용시키는걸 초등학생도 알수 있게 쉽게 해줘 순서 대로 

형, **초등학생 코스** 로 “형이 올린 코드” 그대로 적용하는 법을 한 줄씩 안내할게. (검은 창=명령창, `>>>` 보이는 흰 창=파이썬 창. **명령은 검은 창에서만** 입력) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

22/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **A단계｜전체 개요** 

- 목표: **코드 파일 저장 → 패키지 1개 설치 → 실행 →** `output.txt` **/** `output.srt` **받기.** [사실] 

- 필요한 것: **Python 3 설치됨** , 인터넷 연결. [사실] 

## **B1｜핵심만 (3줄)** 

**1. 코드 저장** : 메모장에 붙여넣고 `extract_simple.py` 로 저장. [사실] 

**2. 설치** : `python -m pip install youtube-transcript-api` (처음 1번). [사실] 
