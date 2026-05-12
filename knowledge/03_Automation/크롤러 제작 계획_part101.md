        token = match.group(1)
# 길이가긴이모티콘묶음이라도 2번만남김
return token * 2
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

504/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
    text = REPEAT_PUNCT_PATTERN.sub(_compress, text)
```

```
# 5) 여러줄공백연속 → 1줄만
    text = re.sub(r"\n{3,}", "\n\n", text)
```

```
return text.strip()
```

```
# ---------- 4. 심리/니즈태그추출 ----------
```

```
# 형이말한것처럼: "심리쪽은전부태그붙여"
# 심리관련/환경제약관련표현패턴들
PSYCH_HINTS = [
"불안", "초조", "두렵", "무섭", "걱정",
"포기하고싶", "그만두고싶", "좌절", "자신감이떨어",
"자신감이없", "의욕이없", "무기력",
"멘탈이", "멘탈이무너", "멘탈나갔",
"우울", "침체", "허무",
]
ENV_HINTS = [
"시간이없", "시간부족",
"집이멀", "거리가멀",
"돈이없", "비용이부담", "수업료가부담",
"직장때문에", "회사때문에", "스케줄때문에",
]
RISK_HINTS = [
"죽고싶", "자살", "사는게힘들",
"극단적인", "진짜끝내고싶", "인생을포기",
]
defextract_tags(title: str, content: str, comments: List[str]) -> List[str]:
"""
형기준:
    - 심리/환경제약/불안/포기계열 → 전부 [심리]
    - 그중에서도극단적표현 → [심리위험] 추가
    """
    tags: List[str] = []
    base_text = "\n".join([title, content, *comments])
```

```
    has_psych = any(h in base_text for h in PSYCH_HINTS)
    has_env = any(h in base_text for h in ENV_HINTS)
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

505/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
    has_risk = any(h in base_text for h in RISK_HINTS)
```

```
if has_psych or has_env or has_risk:
        tags.append("심리")
```

```
if has_risk:
        tags.append("심리위험")
```

```
# 필요하면여기서동기/목표등추가태그구현가능
```

```
# ex) "배우가되고싶", "연기를오래하고싶" 등 → tags.append("목표")
```

```
returnlist(dict.fromkeys(tags))  # 중복제거유지
```

```
# ---------- 5. Post → 최종 dict 변환 ----------
```

```
defclean_post(post: Post) -> dict | None:
"""
```

```
하나의게시글(Post)을엔오원분석용 dict로변환.
```

```
    - 스탭모집글이면 None 반환해서상위에서제외.
```

```
    - 형이정한모든정제/필터/태그규칙을여기서적용.
    """
```

```
# 1) 광고/링크/연락처제거
```

```
    title = remove_ads_and_links(post.title_raw)
```

```
    content = remove_ads_and_links(post.content_raw)
```

```
    comments = [remove_ads_and_links(c) for c in post.comments_raw]
```

```
# 2) 스탭모집글이면제외
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
return {
```

```
"board_name": post.board_name,
```

```
"post_id": post.post_id,
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

506/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
"url": post.url,
```

```
"title_raw": post.title_raw,
```

```
"content_raw": post.content_raw,
```

```
"comments_raw": " ||| ".join(post.comments_raw),
```

```
"text_clean": title_clean + ("\n\n" + content_clean if content_clean else
```

```
"tags": ",".join(tags),
```

```
"created_at": post.created_at.isoformat() if post.created_at else"",
```

```
"writer": post.writer or"",
```

```
    }
```

## 🔍 **이 코드에 형이 지금까지 말한 게 어떻게 들어가 있냐** 

## **URL/이메일/전화번호/광고/링크 문장 → 제거** 

→ `remove_ads_and_links()` 안에 전부 들어있다. 

## **스탭 모집 제외** 

- → `is_staff_recruit_post()` 에서 체크해서, 해당되면 `clean_post()` 에서 `None` 반 환. 

## **공백/줄바꿈/반복/기호/이모티콘 → 자동 정리, 질문 없음** 

- → `normalize_text()` 안에 몰빵. 

- → 여기서 처리 끝, 이 규칙에 대해 앞으로 다시 안 묻는다. 

## **심리/환경 제약/무기력/불안 → 전부 [심리] 태그** 

- → `extract_tags()` 안에서 `PSYCH_HINTS` , `ENV_HINTS` 로 묶어서 처리. 

## **극단적 표현 → [심리위험] 추가** 

- → `RISK_HINTS` 기준으로 별도 플래그. 

## **최종적으로 저장할 필드 구조** 

- → `clean_post()` 마지막 dict에 형이 원한 구조 그대로 들어있다. 

## 형 말대로, 

“기존과 비슷하면 알아서 최적 선택해라. 

꼭 내가 결정해야 할 것만 질문해라.” 

## 그래서 이 코드는 

## **형이 이미 Q&A로 확정한 규칙만 반영했고** , 

추가 질문 없이 내가 **AI로서 “이게 제일 무난하고 안정적인 구현”**이라고 판단해서 짠 버전이야. 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

507/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

이제 이 `cleaner.py` 는 그래비티에 그대로 넘겨서: 

- 나머지 `fetcher.py` (실제 HTML에서 제목/본문/댓글 긁어오기)만 필름메이커스 구조에 맞게 채우면 되고, 

- `main.py` 에서 `fetch_posts()` → `clean_post()` → `save_csv()` 순으로 돌리면 

- 엔오원 규칙대로 크롤링 결과가 나올 수 있어. 

형이 원하면 

다음엔 `fetcher.py` 쪽도 “질문 없이” 

내가 최대한 안전한 구조로 한 번에 짜볼게. 

**==> picture [489 x 38] intentionally omitted <==**

자 니가 구현코드를 만들었다고 했는데 이걸 그래비티에 게 다시 전달해서 새로운 코드를 만들어라? 인거냐 아니면 이걸 그냥그래비티보고 실행하라는거냐? 