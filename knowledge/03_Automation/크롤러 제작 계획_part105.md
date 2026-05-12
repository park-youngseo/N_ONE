"집이멀", "거리가멀",
"돈이없", "비용이부담", "수업료가부담",
"등록비가부담", "학원비가부담",
"직장때문에", "회사때문에", "알바때문에",
"스케줄때문에", "시간조율이안", "쉬는날이안맞",
]
# 극단적인위험표현
RISK_HINTS = [
"죽고싶", "죽고싶",
"자살", "극단적인선택",
"사는게힘들", "사는게힘들",
"모든걸끝내고싶", "끝내버리고싶",
"인생을포기", "포기하고싶을정도",
]
```

## 🧩 **3.** `tagger.py` **— 태그 붙이는 로직** 

python 코드 복사 _`# tagger.py`_ `from typing import List from patterns import PSYCH_HINTS, ENV_HINTS, RISK_HINTS` 

```
defextract_tags(title: str, content: str, comments: List[str]) -> List[str]:
"""
엔오원규칙:
    - 심리/환경제약관련표현이있으면: '심리' 태그
    - 그중에서도극단표현이있으면: '심리위험' 태그추가
    """
    tags: List[str] = []
    base_text = "\n".join([title, content, *comments])
```

```
    has_psych = any(hint in base_text for hint in PSYCH_HINTS)
    has_env = any(hint in base_text for hint in ENV_HINTS)
    has_risk = any(hint in base_text for hint in RISK_HINTS)
if has_psych or has_env or has_risk:
        tags.append("심리")
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

523/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
if has_risk:
        tags.append("심리위험")
# 중복제거 (순서유지)
returnlist(dict.fromkeys(tags))
```

## 🧩 **4.** `cleaner.py` **— 정제 + 필터 + 태깅 + dict 변환** 

python _`# cleaner.py`_ `from typing import List, Optional import re` 

코드 복사 

```
from post import Post
from patterns import (
    URL_PATTERN,
    EMAIL_PATTERN,
    PHONE_PATTERN,
    LINK_SENTENCE_HINTS,
    AD_HINTS,
    STAFF_HINTS,
    REPEAT_PUNCT_PATTERN,
)
from tagger import extract_tags
```

```
# ---------- 1. 광고/URL/연락처제거 ----------
defremove_ads_and_links(text: str) -> str:
"""URL/이메일/전화번호/광고·링크언급문장제거."""
ifnot text:
return text
# 1) URL/이메일/전화번호직접제거
    text = URL_PATTERN.sub("", text)
    text = EMAIL_PATTERN.sub("", text)
    text = PHONE_PATTERN.sub("", text)
# 2) 문장단위로쪼개서 '링크언급/광고' 계열문장제거
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

524/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
# 간단하게줄단위 + 마침표/물음표/느낌표기준으로처리
# 구분자까지포함해서문장재구성
    chunks = re.split(r"([.!?]\s*|\n)", text)
    merged: List[str] = []
    buffer = ""
for chunk in chunks:
if chunk isNone:
continue
if chunk == "":
continue
        buffer += chunk
# 마침표/물음표/느낌표/줄바꿈으로문장단위구분
if re.search(r"[.!?]\s*$", chunk) or"\n"in chunk:
            sentence = buffer.strip()
            buffer = ""
            lower = sentence.lower()
# 링크언급힌트
ifany(hint in sentence for hint in LINK_SENTENCE_HINTS):
continue
# 광고/홍보힌트
ifany(hint in sentence for hint in AD_HINTS):
continue
            merged.append(sentence)
```

```
# 남은잔여버퍼처리
if buffer.strip():
        merged.append(buffer.strip())
```

```
return"\n".join(merged).strip()
```

```
# ---------- 2. 스탭모집글여부판별 ----------
```

```
defis_staff_recruit_post(title: str, content: str) -> bool:
"""제목/본문에스탭모집관련문구가있으면 True."""
    base = f"{title}\n{content}"
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

525/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
for hint in STAFF_HINTS:
if hint in base:
returnTrue
returnFalse
```

```
# ---------- 3. 텍스트정제 (공백/줄바꿈/반복/기호등) ----------
defnormalize_text(text: str) -> str:
"""
엔오원규칙:
    - 탭/캐리지리턴정리
    - 여러칸공백 → 1칸
    - 줄마다양끝공백제거
    - 3개이상반복되는기호/이모티콘 → 2개로압축
    - 3줄이상연속줄바꿈 → 2줄로압축
    - 문장의미는절대변경하지않음
    """
ifnot text:
return text
# 줄바꿈정규화
    text = text.replace("\r\n", "\n").replace("\r", "\n")
    text = text.replace("\t", " ")
# 여러개연속공백 → 1칸
    text = re.sub(r"[ ]{2,}", " ", text)
# 줄마다앞뒤공백제거
    lines = text.split("\n")
    lines = [line.strip() for line in lines]
    text = "\n".join(lines)
# 3개이상반복되는기호/이모티콘 → 2개로압축
def_compress(match: re.Match) -> str:
        token = match.group(1)
return token * 2
    text = REPEAT_PUNCT_PATTERN.sub(_compress, text)
```

```
# 3줄이상연속줄바꿈 → 2줄
    text = re.sub(r"\n{3,}", "\n\n", text)
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

526/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
return text.strip()
```

```
# ---------- 4. Post → 분석용 dict 변환 ----------
```

```
defclean_post(post: Post) -> Optional[dict]:
"""
하나의 Post 객체에엔오원규칙(정제/필터/태그)을모두적용해
최종 CSV/JSON 저장용 dict로변환.
    - 스탭모집글이면 None 반환.
    """
```

```
# 1) 광고/링크/연락처제거
```

```
    title = remove_ads_and_links(post.title_raw)
    content = remove_ads_and_links(post.content_raw)
    comments = [remove_ads_and_links(c) for c in post.comments_raw]
```

```
# 2) 스탭모집글은제외
```

```
if is_staff_recruit_post(title, content):
```

```
returnNone
```

```
# 3) 텍스트정제
```

```
    title_clean = normalize_text(title)
    content_clean = normalize_text(content)
    comments_clean = [normalize_text(c) for c in comments]
```

```
# 4) 태그추출
```

```
    tags = extract_tags(title_clean, content_clean, comments_clean)
```

```
# 5) 최종 dict 구성
```

```
    text_clean = title_clean
if content_clean:
```

```
        text_clean = f"{title_clean}\n\n{content_clean}"
```

```
return {
```

```
"board_name": post.board_name,
```

```
"post_id": post.post_id,
```

```
"url": post.url,
```

```
"title_raw": post.title_raw,
```

```
"content_raw": post.content_raw,
```