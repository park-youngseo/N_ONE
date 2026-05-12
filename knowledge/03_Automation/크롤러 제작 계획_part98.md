```

```
# 최근 years년까지만수집 (날짜조건체크)
```

```
return posts
```

## 여기서는 **날짜 필터 / 게시판별 URL / 페이지 넘김** 만 처리. 

정제·태그는 여기서 절대 안 건드림. 

## `cleaner.py` **– 형이 욕하면서까지 만든 규칙 전부 적용하는 모듈** 

python 코드 복사 

```
from .models import Post
```

```
defclean_post(post: Post) -> dict:
```

```
# 1) 광고/URL/연락처삭제
```

```
    title = remove_ads_and_links(post.title_raw)
    content = remove_ads_and_links(post.content_raw)
    comments = [remove_ads_and_links(c) for c in post.comments_raw]
```

```
# 2) 스탭모집등제외여부판정
```

```
if is_staff_recruit_post(title, content):
returnNone# saver 쪽에서 None은건너뛰게처리
```

```
# 3) 텍스트정제(공백/반복/기호등자동정리)
```

```
    title_clean = normalize_text(title)
    content_clean = normalize_text(content)
    comments_clean = [normalize_text(c) for c in comments]
```

```
# 4) 심리/동기/환경관련태그달기
```

```
    tags = extract_tags(title_clean, content_clean, comments_clean)
```

```
return {
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

491/633 

25. 11. 30. 오후 12:53 

 

크롤러 제작 계획 

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

```
"comments_raw": " ||| ".join(post.comments_raw),
```

```
"text_clean": title_clean + "\n\n" + content_clean,
```

```
"tags": ",".join(tags),
"created_at": post.created_at.isoformat(),
"writer": post.writer or"",
```

```
    }
```

## 여기 안에서: 

```
remove_ads_and_links
```

- → 광고/URL/링크 언급 문장/홍보 문구 제거 

```
normalize_text
```

- → 공백/반복문자/기호/이모티콘 등 **형이 통으로 묶어버린 정제 규칙** 전부 자동 적용 

- → “한 칸, 줄바꿈, 콤마” 같은 걸 묻지 않고 여기서 그냥 처리 

```
extract_tags
```

- → 심리/동기/환경 고민 → 전부 `[심리]` 한 덩어리 기준 

- → 위험한 표현 있으면 `[심리위험]` 같이 추가 

- → 형이 말한 “심리 계열은 전부 태그 붙여” 그거 그대로 

## `saver.py` **– CSV/JSON으로 저장** 

python 

코드 복사 

```
import csv
import json
from pathlib import Path
```

```
defsave_csv(rows: list[dict], save_path: str) -> None:
ifnot rows:
return
    path = Path(save_path)
    path.parent.mkdir(parents=True, exist_ok=True)
with path.open("w", newline="", encoding="utf-8-sig") as f:
        writer = csv.DictWriter(f, fieldnames=rows[0].keys())
        writer.writeheader()
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

492/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
        writer.writerows(rows)
```

```
defsave_json(rows: list[dict], save_path: str) -> None:
    path = Path(save_path)
    path.parent.mkdir(parents=True, exist_ok=True)
with path.open("w", encoding="utf-8") as f:
        json.dump(rows, f, ensure_ascii=False, indent=2)
```

## `logger.py` **– 화면 & 파일 로그** 

python 

코드 복사 

```
from datetime import datetime
from pathlib import Path
```

```
deflog(message: str):
    ts = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    line = f"[{ts}] {message}"
print(line)
    log_dir = Path("logs")
    log_dir.mkdir(exist_ok=True)
with (log_dir / "crawler.log").open("a", encoding="utf-8") as f:
        f.write(line + "\n")
```

그래비티/EXE에서 이 `log()` 만 보여줘도 형이 실행 화면에서 본 “진행 상태” 그대로 나옴. 

## `main.py` **– 전체 조합 (실행 진입점)** 

python 

코드 복사 

```
from .config import BOARDS, DEFAULT_YEARS, SAVE_FORMAT, SAVE_DIR
from .fetcher import fetch_posts
from .cleaner import clean_post
from .saver import save_csv, save_json
from .logger import log
from pathlib import Path
from datetime import datetime
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

493/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
defrun_crawler(years: int = DEFAULT_YEARS):
    all_rows: list[dict] = []
```

```
for board_key, board_path in BOARDS.items():
        log(f"{board_key}수집시작 (최근{years}년)")
        posts = fetch_posts(board_key, years)
```

```
for post in posts:
            cleaned = clean_post(post)
if cleaned isNone:
continue
            all_rows.append(cleaned)
```

```
        log(f"{board_key}수집완료: {len(posts)}건중{len(all_rows)}건유효")
```

## _`#`_ `저장 파일명` 

```
    now_str = datetime.now().strftime("%Y%m%d_%H%M%S")
if SAVE_FORMAT == "csv":
        save_path = str(Path(SAVE_DIR) / f"filmmakers_{now_str}.csv")
        save_csv(all_rows, save_path)
        log(f"CSV 저장완료: {save_path}")
elif SAVE_FORMAT == "json":
        save_path = str(Path(SAVE_DIR) / f"filmmakers_{now_str}.json")
        save_json(all_rows, save_path)
        log(f"JSON 저장완료: {save_path}")
```

## 그래비티든 EXE든, 

**결국 이** `run_crawler()` **만 호출해주면 엔오원 규칙대로 전부 돌아가게 설계되어 있어.** 

## 🅑 **2 정리 – 지금 단계에서 중요한 포인트** 

## **질문 없음** 

→ 이 설계 안에는 더 고를 옵션/선택지가 없다. 

## **형이 이미 결정한 규칙만 반영** 

- → 정제/삭제/태그/필터 룰은 전부 `cleaner.py` 안에만 고정. 

## **콤마/공백/줄바꿈 관련 장난 없음** 

- → 그건 `normalize_text()` 내부에서 자동 처리, 더 이상 밖으로 안 나온다. 

## **심리 계열은 전부 태그** 

- → `[심리]` , `[심리위험]` 등 형이 말한 대로 한 덩어리 취급. 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

494/633 

크롤러 제작 계획 