실행 명령은 `python 실제파일명.py` 다. 예시 글자(파일이름.py)를 쓰면 안 됨. [사실] 

## **B2단계｜지금 그대로 따라하기(Windows)** 

## **1) 내 폴더에 어떤 파일이 있는지 확인** 

검은 창(CMD)에 입력: 

bash 코드 복사 

```
dir
```

→ 목록에서 `.py` **로 끝나는 파일명** 을 찾아. (예: `extract_transcript.py` 또는 `extract_text_only.py` ) [사실] 

## **2) 실제 파일로 실행** 

목록에서 본 이름으로: 

nginx 코드 복사 `python extract_transcript.py` 

또는 

nginx 코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

28/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python extract_text_only.py
```

→ “유튜브 URL …” 뜨면 주소 붙여넣고 Enter. [사실] 

## **3) 만약** `.py` **가 안 보이면 (메모장이 .txt로 저장했을 때)** 

**1.** 탐색기에서 상단 **보기 → 파일 이름 확장명** 체크 켜기. 

**2.** 파일이 `extract_text_only.py.txt` 처럼 되어 있으면 **이름 바꾸기** : CMD에서: 

powershell 코드 복사 `ren "extract_text_only.py.txt" "extract_text_only.py"` 

## 그런 다음 다시 실행: 

nginx 코드 복사 

```
python extract_text_only.py
```

[사실] 

## **4) 프로젝트 압축판을 썼다면 더 쉬움** 

압축판( `youtube_script_extractor_v1_1.zip` )을 풀어 썼다면: 

가장 쉬운 방법: 

arduino 코드 복사 `run.bat` nginx 코드 복사 

## 또는 직접: 

```
python extract_transcript.py
```

[사실] 

## **5) 결과 위치** 

성공하면 같은 폴더에 

`output.txt` (텍스트만) — **간단 스크립트 버전** `outputs\제목-영상ID.txt / .srt` — **압축판 프로젝트 버전** 이 만들어진다. [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

29/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **빠른 체크** 

- `python 파일이름.py` → ❌ (예시 문구) 

- `python extract_transcript.py` → ⭕ 

- `python extract_text_only.py` → ⭕ 

- `dir` 로 파일명 먼저 확인 → **보이는 이름 그대로** 입력. [사실] 

막히면 `dir` 찍힌 목록이랑 실행했을 때 뜬 **문구 전체** 를 복사해서 보내줘. 거기서 바로 이 어서 잡아줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

야이 새끼야 처음 부터 다시해 저장부터 

형, 알겠어. **완전 처음부터, “저장”부터** 한 줄씩 간다. (PARTICIPATION=Off / EVALUATION=Off) 

## **A단계｜전체 개요** 

- 목표: **코드 파일을 메모장에 저장(.py) → 설치 1줄 → 실행 1줄 → output.txt 생성** . [사실] 

## **B2단계｜손잡고 처음부터 (Windows, 초등학생 코스)** 

## **0) 준비(파일 확장자 보이게)** 

**1.** 아무 폴더나 열기 → 상단 **보기** 클릭 → **파일 확장명** 체크 켜기. → 이유: **.py** 로 정확히 저장됐는지 눈으로 확인하려고. [사실] 

## **1) 코드 “저장” (메모장 → .py 만들기)** 

## **1. 메모장** 열기. 

**2.** 아래 코드를 **전부 복사** 해서 **메모장에 붙여넣기** : 

python 

코드 복사 

```
from youtube_transcript_api import YouTubeTranscriptApi, NoTranscriptFound, Trans
from urllib.parse import urlparse, parse_qs
```

```
defextract_video_id(url: str) -> str:
```

```
    q = urlparse(url.strip())
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

30/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

`host = (q.hostname or "").lower() if q.hostname else "" if host == "youtu.be": return q.path.lstrip("/") if host in ("www.youtube.com", "youtube.com", "m.youtube.com"): if q.path == "/watch": return parse_qs(q.query).get("v", [""])[0] if q.path.startswith("/embed/") or q.path.startswith("/v/"): return q.path.split("/")[2] return url.strip()` _`#`_ `그냥` _`ID`_ `만 넣어도 동작 if __name__ == "__main__": url = input("유튜브 URL 또는 영상 ID: ").strip() try: vid = extract_video_id(url)` _`#`_ `한글 우선` _`,`_ `없으면 영어 data = None for lang in ["ko", "en"]: try: data = YouTubeTranscriptApi.get_transcript(vid, languages=[lang]) break except NoTranscriptFound: continue if data is None: raise NoTranscriptFound("공개 자막 없음") text = "\n".join([(d.get("text") or "").replace("\n"," ").strip() for d i with open("output.txt", "w", encoding="utf-8") as f: f.write(text) print("완료! 같은 폴더에 output.txt 가 저장되었습니다.") except TranscriptsDisabled: print("이 영상은 자막이 비활성화되어 있습니다. (Whisper 등 음성 인식 필요)") except VideoUnavailable: print("영상을 찾을 수 없습니다. URL/ID를 확인하세요.") except NoTranscriptFound: print("자막이 없어 텍스트를 만들 수 없습니다. (공개 캡션 없음일 수 있음)") except Exception as e: print("오류:", e)`   

## **3. 파일 → 다른 이름으로 저장…** 클릭. 

**4.** 저장 위치를 형이 작업할 폴더(예: 바탕화면의 `youtube_script_extractor` )로 지정. 

**5. 파일 이름** : `extract_text_only.py` 

**6. 파일 형식** : “모든 파일(.)” 

- **7. 인코딩** : **UTF-8** 선택 → **저장** . 

 

 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

31/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**8.** 폴더에서 `extract_text_only.py` 로 끝나면 성공(뒤에 `.txt` 붙으면 실패). [사실] 

## **2) 설치(처음 1번만)** 

**1.** 방금 저장한 **폴더 창에서** 상단 주소창을 **한 번 클릭** → `cmd` 입력 → **Enter** . → 검은 창(CMD)가 **그 폴더** 에서 열린다. [사실] 

**2.** 아래 **그대로 입력 → Enter** : 

nginx 

코드 복사 

```
python -m pip install youtube-transcript-api
```

- → 설치 글자가 주르륵 나오면 끝. [사실] 

## **3) 실행(매번 이거)** 

검은 창에서: 

nginx 코드 복사 

```
python extract_text_only.py
```

- → “유튜브 URL 또는 영상 ID:” 라고 뜨면 **유튜브 주소를 붙여넣고 Enter** . [사실] 

## **4) 결과 확인** 

- **같은 폴더** 에 `output.txt` 파일이 생긴다. 

- 열어보면 **영상 스크립트(자막 텍스트)** 가 줄바꿈으로 저장됨. [사실] 

## **5) 안 되면 여기만 보면 됨(초간단 트러블슈팅)** 

`python: can't open file '...py'` → **파일 이름이 다르거나 폴더가 다름** 

**1.** CMD에서 `dir` 치고 **정확한 파일명** 확인 

**2.** 그 이름으로 실행: `python 정확한파일명.py` [사실] 

- 명령이 파란/빨간 오류로 끝남 → **오류 문구 전체 복사** 해서 보내줘. 그 자리에서 수정 포인트 찍어줄게. [사실] 

- “자막 없음” 메시지 → 해당 영상에 **공개 캡션이 없다** 는 뜻. 다른 영상으로 테스트하 거나, 필요하면 **Whisper(음성 인식)** 버전으로 안내 가능. [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

32/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **(선택) 더 쉬운 통합 패키지 쓰고 싶으면** 

내가 만든 **원클릭 배치 포함 버전** 도 있어. 압축 풀고 `run.bat` 만 더블클릭하면 된다: ▶ **youtube_script_extractor_v1_1.zip** (원하면 다시 링크 줄게). [추론] 

필요하면 지금 화면 캡처나 **CMD에 나온 글자 그대로** 보내. 거기서 바로 이어서 정리한 다, 형. [사실] 

**==> picture [489 x 39] intentionally omitted <==**

## 3) 실행(매번 이거) 

검은 창에서: 

python extract_text_only.py 

→ “유튜브 URL 또는 영상 ID:” 라고 뜨면 유튜브 주소를 붙여넣고 Enter. [사실] 

2번 끝내고 새로 cmd창을 열어야 하는거야? 아님 바로 실 행해도 되는거야? 

10s 동안 생각함 

형, **새 창 안 열어도 돼.** 2번(설치) 끝난 그 **같은 CMD 창에서 바로 실행** 하면 된다. [사실] 

## **A단계｜핵심** 

`python -m pip install youtube-transcript-api` 가 **성공했다면** , 바로 이어서 `python extract_text_only.py` 치면 됨. **새 CMD 불필요.** [사실] 

## **B2단계｜바로 실행 체크(예외 포함)** 

**1.** (선택) 확인: `python -V` → 버전 보이면 OK. [사실] 

**2.** 실행: `python extract_text_only.py` → “유튜브 URL…” 뜨면 주소 붙여넣기. [사실] 

**3. 새 창이 필요한 예외(드뭄)** [사실] 

방금 **처음으로 Python을 설치했고** , 그 이전에 열어둔 CMD에서 `python` 이 **인식** 

## **안 됐던 경우** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

33/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

또는 `pip` 만 쓰다가 **PATH 갱신이 필요한 환경** 

→ 이런 경우에만 **CMD를 닫고 다시 열기** 후 위 1)~2) 진행. 

문구가 뜨는데 막히면, 그 **메시지 그대로** 붙여줘. 바로 이어서 잡을게. 

**==> picture [489 x 39] intentionally omitted <==**

[notice] A new release of pip is available: 24.0 -> 25.2 [notice] To update, run: python.exe -m pip install -- upgrade pip 

C:\Users\엔오원 

\Desktop\youtube_script_extractor>python extract_text_only.py 

유튜브 URL 또는 영상 ID: https://youtu.be/ollMbEOzofc? si=sA8na-IHMHrpYwxP 