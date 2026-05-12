
오류: type object 'YouTubeTranscriptApi' has no attribute 'get_transcript' 

C:\Users\엔오원\Desktop\youtube_script_extractor> 

18s 동안 생각함 

형, 오류 원인 딱 하나로 보면 **라이브러리 버전/충돌** 이야. `get_transcript` 메서드를 못 찾는 걸로 보이면, **깔린 패키지가 꼬였거나 다른 버전** 이 잡힌 거야. 

→ **해결: 깨끗이 갈아끼우기(3줄)** 

## **A단계｜핵심** 

지금 창(CMD)에서 **그대로 3줄** 실행 → 다시 돌리면 끝. [사실] 

## **B2단계｜손잡고 따라하기 (CMD에 한 줄씩)** 

## **1. pip 업데이트(선택이지만 권장)** 

css 

코드 복사 

```
python -m pip install --upgrade pip
```

## **2. 문제 패키지 제거** 

nginx 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

34/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m pip uninstall -y youtube-transcript-api
```

## **3. 검증된 버전 설치( get_transcript 지원 )** 

nginx 코드 복사 `python -m pip install youtube-transcript-api==0.6.1` 

## 이제 **바로 실행** : 

nginx 코드 복사 `python extract_text_only.py` 

→ “유튜브 URL 또는 영상 ID:” 뜨면 주소 붙여넣기. 

## **만약 또 에러가 뜨면 (초간단 점검 1줄)** 

아래를 실행해서 메서드가 보이는지 확인해줘. 

python 코드 복사 

```
python - << EOF
```

```
from youtube_transcript_api import YouTubeTranscriptApi
print('get_transcript?', hasattr(YouTubeTranscriptApi,'get_transcript'))
print('list_transcripts?', hasattr(YouTubeTranscriptApi,'list_transcripts'))
EOF
```

결과를 그대로 붙여주면, 그 상태에 맞춰 **코드/버전** 즉시 맞춰줄게, 형. [사실] 

**==> picture [489 x 38] intentionally omitted <==**

C:\Users\엔오원 \Desktop\youtube_script_extractor>python extract_text_only.py 유튜브 URL 또는 영상 ID: https://youtu.be/ollMbEOzofc? si=VzGvz7qn8BP2YmBS 오류: no element found: line 1, column 0 

C:\Users\엔오원\Desktop\youtube_script_extractor> 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

35/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

형, 이 에러 **“no element found: line 1, column 0”** 는 보통 

**1.** 그 영상에 **공개 자막이 없거나/차단됨** , 2) **로그인·연령 제한/멤버십/지역 제한** , 3) **유 튜브 응답이 비어옴(간헐적 네트워크/파싱 이슈)** 때 뜬다. [사실] 

아래대로 처리하면 쉬워. 

## **A단계｜핵심 처리** 

**대응 3단계:** (1) 코드 보강본으로 교체 → (2) 패키지 재설치 → (3) 자막 있는 테스트 영상으로 점검. [사실] 

## **B2단계｜손잡고 해결** 

## **1) 코드 교체(붙여넣기→저장)** 

메모장 열고 아래 **전부** 붙여넣어 `extract_text_only.py` 로 **덮어쓰기 저장** : 

python 코드 복사 

```
from youtube_transcript_api import (
    YouTubeTranscriptApi, NoTranscriptFound,
    TranscriptsDisabled, VideoUnavailable
)
from urllib.parse import urlparse, parse_qs
```

```
defextract_video_id(url: str) -> str:
    u = url.strip()
    q = urlparse(u)
    host = (q.hostname or"").lower() if q.hostname else""
if host in ("youtu.be",):
return q.path.lstrip("/")
if host in ("www.youtube.com", "youtube.com", "m.youtube.com"):
if q.path == "/watch":
return parse_qs(q.query).get("v", [""])[0]
if q.path.startswith("/embed/") or q.path.startswith("/v/"):
return q.path.split("/")[2]
# 그냥 ID만넣어도허용
return u
```

```
deffetch_transcript(video_id: str):
# 1) ko → en 우선시도 (신버전/구버전공통)
for lang in ["ko", "en"]:
try:
return YouTubeTranscriptApi.get_transcript(video_id, languages=[lang]
except NoTranscriptFound:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

36/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
continue
except TranscriptsDisabled:
raise
# 2) 가능하면목록조회후번역(구버전안전호출)
try:
        list_fn = getattr(YouTubeTranscriptApi, "list_transcripts")
if list_fn:
            trs = list_fn(video_id)
# 번역우선
for tr in trs:
try:
return tr.translate("ko").fetch()
except Exception:
pass
# 원문이라도
for tr in trs:
try:
return tr.fetch()
except Exception:
pass
except Exception:
pass
raise NoTranscriptFound("사용가능한공개자막이없습니다.")
if __name__ == "__main__":
    url = input("유튜브 URL 또는영상 ID: ").strip()
try:
        vid = extract_video_id(url)
        data = fetch_transcript(vid)
        text = "\n".join(
            [(d.get("text") or"").replace("\n"," ").strip()
for d in data if d.get("text")]
        )
withopen("output.txt", "w", encoding="utf-8") as f:
            f.write(text)
print("완료! output.txt 저장됨.")
except TranscriptsDisabled:
print("이영상은자막이비활성화(비공개)입니다. (Whisper 같은음성인식필요)"
except VideoUnavailable:
print("영상을찾을수없습니다. URL/ID 확인필요.")
except NoTranscriptFound as e:
print("자막없음/번역실패:", e)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 37/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
except Exception as e:
```

```
"
print(네트워크/파싱문제가능성. 잠시후다시시도하거나다른영상으로테스트:
```

## **2) 패키지 정리(같은 CMD 창에서)** 

lua 

코드 복사 

```
python -m pip install --upgrade pip
python -m pip uninstall -y youtube-transcript-api
python -m pip install youtube-transcript-api==1.2.3
```

※ 이렇게 하면 최신 API와 호환. [사실] 

## **3) 실행 & 점검** 

nginx 

코드 복사 

```
python extract_text_only.py
```

- 먼저 **자막이 확실히 있는 영상** 으로 테스트해봐. 

- 문제가 없다면, 다시 형이 원하는 영상으로 시도. 

- 그래도 실패하면 그 영상은 **공개 자막이 없음/제한됨** 일 확률 높음 → 그땐 **Whisper(음성 인식)** 경로로 가야 해. [사실] 

## **체크리스트(짧게)** 

- 같은 CMD 창에서 바로 실행해도 됨. 새 창 불필요. [사실] 

- 특정 영상만 실패 → 영상 자체 이슈(자막 없음/연령·멤버십·지역 제한). [사실] 

- 여러 영상 전부 실패 → 네트워크(회사/학원 방화벽, VPN, 일시적 차단) 가능 → 잠시 후 재시도. [추론] 

오류가 다시 뜨면 **문구 그대로** 보내줘. 그 상태에 맞춰 바로 다음 스텝 안내할게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

C:\Users\엔오원 \Desktop\youtube_script_extractor>python extract_text_only.py 

유튜브 URL 또는 영상 ID: https://youtu.be/ollMbEOzofc? si=VzGvz7qn8BP2YmBS 

네트워크/파싱 문제 가능성. 잠시 후 다시 시도하거나 다 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

38/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

른 영상으로 테스트: type object 'YouTubeTranscriptApi' has no attribute 'get_transcript' 

## 형, 이건 **라이브러리 “이름 충돌/잘못된 버전”** 가능성이 높아. 

아래 4단계만 그대로 복붙해서 하면 바로 잡힌다. (같은 CMD 창에서 진행) 

## **A단계｜핵심** 

**무엇이 임포트되는지** 먼저 확인 → **충돌 있으면 치우고** → **검증 버전 재설치** → **다시** 

**실행** . [사실] 

## **B2단계｜한 줄씩 따라하기** 

## **1) 어떤 파일이 임포트됐는지 확인** 

아래를 그대로 붙여넣고 Enter: 

python 코드 복사 

```
python - << EOF
```

```
import youtube_transcript_api, sys
```
