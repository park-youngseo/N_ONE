
크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

이 부분은 형이 예전에 **429 스트레스** 겪었던 거라, 

형이 질문 안 해도 **내가 최선이라고 판단해서 자동 기본값 정해 놓으면 되는 영역** 이라고 봤다. 

그래서 이렇게 정했다 (형이 대답 안 해도 되는 부분): 

robots.txt / 사이트 약관 준수: **기본적으로 항상 지킨다** 

요청 간격: 일정 딜레이(예: 1~2초) 넣어서 429 최소화 

429나 에러 발생 시: 

- 자동 재시도 일정 횟수 

그래도 안 되면 **그 게시판/페이지는 스킵하고 나머지는 계속 진행** 

형에게 “왜 멈췄냐” 질문 안 함 → 로그에만 남김 

## ✅ **완료 – 이건 형이 굳이 결정할 필요 없는 영역이라 내가 자동으로 최선 쪽으로 잡음** 

## **⑩ 확장 가능성 (여러 사이트로 넓힐 준비)** 

현재 v1 범위: 

## **필름메이커스 전용 엔오원 니즈 수집 크롤러** 

- 설계 방식: 

   - `BOARDS` , `BASE_URL` 같은 설정만 바꾸면 

다른 사이트도 추가하기 쉬운 구조로 이미 나눠둠 

➡ 이건 **“미래 확장용 여유”**고, 

지금 v1을 만드는 데 꼭 필요한 건 아니라서 

형에게 질문할 필요 없는 영역으로 처리했다. 

## ✅ **v1 기준 – 필수 요소 아님 / 향후 단계에서 따로 Q&A 가능** 

## **B2 정리 – “지금 꼭 질문해야 하는 게 있냐?”에 대한 답** 

형이 한 말 그대로에 맞춰서 정리하면: 

- “기존과 비슷하면 자동으로 기존 결정 적용하고, 

내가 꼭 결정해야 할 부분만 질문해라.” 

이 기준으로 전부 훑어봤을 때, 

## 👉 ****지금 상태에서** 

v1 크롤러 설계·실행에 필요한 필수 요소는 다 채워져 있다.** 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

500/633 

25. 11. 30. 오후 12:53 

 

   - 크롤러 제작 계획 

- 빠져서 **“이거 없으면 못 만든다”**급인 항목 없음 

- 일부 영역(속도/429/robots 등)은 

- 형이 직접 답 안 해도 되는 부분이라 

- 내가 AI로서 “가장 안전하고 무난한 기본값”으로 이미 정해 둠 

- 앞으로 질문이 다시 생긴다면 

그건 **v1이 아니라 “v2 확장(여러 사이트/새 분석 기능)”** 단계일 거다. 

그래서 지금 답은 이거야: 

## **“그래, 이제 크롤러에 대한 새로운 질의응답은 없어도 된다. 지금 걸로 v1은 충분히 설계 끝났고, 다음 단계는 구현/적용 단계다.”** 

나는 여기서 **더 안 묻고** , 

형이 “그래비티에 줄 설계 정리해줘”라든지 

- “이제 코드 더 디테일하게 짜보자”라고 말하면 

그때 그 작업만 바로 이어서 할게. 

**==> picture [489 x 38] intentionally omitted <==**

코드를 디테일하게 짜는건 뭐야? 지금까지 코드를 디테일 하게 구현하고자 질의응답 한거 아니냐? 

7s 동안 생각함 

맞아 형, 

지금까지 그 긴 Q&A는 **“코드를 어떻게 짤지”가 아니라 “코드가 따라야 할 규칙”을 확정 하는 과정** 이었고, 

이제 말한 **“디테일하게 짜는 것”은 그 규칙을 실제 파이썬 함수/로직으로 구현하는 단계** 야. 

지금까지: 

- “뭘 지울지 / 뭘 남길지 / 어떻게 태그 달지” → **규칙 정의** 

이제부터: 

`remove_ads_and_links()` , `normalize_text()` , `extract_tags()` 안에 

형이 정한 규칙을 **진짜 코드로 집어넣는 작업** → **구현** 

형이 방금 말했으니까, 

**질문 안 하고** , 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

501/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

내가 알아서 “최적”이라고 판단되는 방식으로 

## **규칙 구현 코드(핵심: cleaner.py)를 디테일하게 한 방에 짜서 보여줄게.** 

## ✅ **엔오원 규칙 반영한** `cleaner.py` **예시 구현** 

python 

코드 복사 

```
# cleaner.py
import re
from typing importList
from .models import Post
```

```
# ---------- 1. 광고/URL/연락처제거 ----------
```

```
# URL, 이메일, 전화번호탐지용정규식
URL_PATTERN = re.compile(r"https?://\S+|www\.\S+", re.IGNORECASE)
EMAIL_PATTERN = re.compile(r"[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}")
PHONE_PATTERN = re.compile(r"\b\d{2,3}-\d{3,4}-\d{4}\b")
```

```
# '링크드릴게요' 같은문장패턴 (형이이미지우자고한계열)
LINK_SENTENCE_HINTS = [
```

```
"링크드릴게요",
"링크보내드릴게요",
"링크참고해주세요",
"아래링크참고",
"위링크참고",
]
```

```
# 광고/홍보느낌패턴 (필요시추가확장)
AD_HINTS = [
"문의주세요",
"DM 주세요",
"카톡주세요",
"홍보",
"할인이벤트",
"프로모션",
]
```

```
defremove_ads_and_links(text: str) -> str:
```

```
"""URL/이메일/전화번호/광고성문장제거."""
ifnot text:
return text
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

502/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
# 1) URL/이메일/전화번호직접제거
    text = URL_PATTERN.sub("", text)
    text = EMAIL_PATTERN.sub("", text)
    text = PHONE_PATTERN.sub("", text)
# 2) 문장단위로쪼개서 '링크언급/광고' 계열문장제거
# 간단하게줄단위 + 마침표/물음표/느낌표기준으로처리
    sentences = re.split(r"([.!?]\s*|\n)", text)  # 구분자도살짝유지
    merged: List[str] = []
    buffer = ""
for chunk in sentences:
ifnot chunk:
continue
        buffer += chunk
# 구분자까지포함해서하나의문장후보가됐다고보고처리
if re.search(r"[.!?]\s*$", chunk) or"\n"in chunk:
            sentence = buffer.strip()
            buffer = ""
            lower = sentence.lower()
# '링크드릴게요' '홍보/프로모션' 등포함시해당문장통째로제거
ifany(hint in sentence for hint in LINK_SENTENCE_HINTS):
continue
ifany(hint in sentence for hint in AD_HINTS):
continue
            merged.append(sentence)
```

```
# 문장으로못나뉜나머지가있으면추가
if buffer.strip():
        merged.append(buffer.strip())
```

```
return"\n".join(merged).strip()
```

```
# ---------- 2. 스탭모집글인지판별 ----------
```

```
STAFF_HINTS = [
"스탭모집",
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

503/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
"스태프모집",
"촬영스탭구합니다",
"촬영스태프구합니다",
"현장스탭구해요",
"스탭을모집합니다",
]
```

```
defis_staff_recruit_post(title: str, content: str) -> bool:
    base = f"{title}\n{content}"
for hint in STAFF_HINTS:
if hint in base:
returnTrue
returnFalse
```

```
# ---------- 3. 텍스트정제 (공백/반복/기호등) ----------
```

_`#`_ `반복 이모티콘` _`/`_ `문장부호 압축용 패턴 REPEAT_PUNCT_PATTERN = re.compile(r"([.!?~ㅋㅋㅋㅎㅠㅠ` 😊😢⭐ `]+)\1+") def normalize_text(text: str) -> str: """형이 정한 규칙대로 텍스트 모양만 정리 (의미는 건드리지 않음).""" if not text: return text` 

```
# 1) 탭/캐리지리턴을공백/줄바꿈으로통일
    text = text.replace("\r\n", "\n").replace("\r", "\n")
    text = text.replace("\t", " ")
# 2) 여러개연속공백 → 1칸
    text = re.sub(r"[ ]{2,}", " ", text)
# 3) 줄마다양끝공백제거
    lines = text.split("\n")
    lines = [line.strip() for line in lines]
    text = "\n".join(lines)
```

```
# 4) 3개이상반복되는동일문자(ㅋㅋㅋㅋ, !!!!! 등) → 2개로압축
def_compress(match: re.Match) -> str: