프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

site_name, urls = ask_user() log(f"=== {site_name} 크롤링 시작 ===") 

all_rows = [] 

for url in urls: 

rows = fetch_generic_board(url) all_rows.extend(rows) 

filename = f"{site_name}.csv" 

path = save_csv(filename, all_rows) 

log(f"=== 저장 완료: {path} ===") print(f"\n저장 완료: {path}") 

if __name__ == "__main__": run() input("\n종료하려면 Enter를 누르세요.") 

[main.py 파일 내용 끝] 

그래비티에게 줄 파일이야 최종적으로 exe 파일을 만들수 있게 문서에 추가해 

24s 동안 생각함 

## text 

코드 복사 

```
# ============================================
```

- `# 엔오원 니즈 수집 크롤러 v2.4` 

- `# 그래비티 생성용 완성본(복붙용)` 

- `# ============================================` 

## `# 생성 규칙` 

- `# 1) 아래에 정의된 파일명을 그대로 생성한다.` 

- `# 2) cleaner.py / tagger.py / patterns.py / post.py / saver.py / logger.py 는 수정 # 3) fetcher.py / main.py 만 그래비티가 자동 생성.` 

- `# 4) output / logs 폴더 자동 생성.` 

- `# 5) 엔진 코어는 절대 변형하지 말고 그대로 만든다.` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 9/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
# 6) requirements.txt / build_exe.bat 파일도함께생성한다.
# 7) build_exe.bat 는 Windows에서 PyInstaller로 exe를빌드하는스크립트로만든다.
```

```
# ============================================
# 1. cleaner.py
# ============================================
[cleaner.py 파일내용시작]
import re
def clean_text(text: str) -> str:
    if not text:
        return ""
    # URL 제거
    text = re.sub(r"http[s]?://\S+", "", text)
    # HTML 태그제거
    text = re.sub(r"<[^>]+>", "", text)
    # 광고/홍보성패턴제거
    text = re.sub(r"(문의|카톡|오픈채팅|insta|인스타|DM주세요|구독|좋아요|체험단)", "
    # 이모지제거
    text = re.sub("[\U00010000-\U0010ffff]", "", text)
    # 반복문자 2개제한
    text = re.sub(r"(.)\1{2,}", r"\1\1", text)
    # 양쪽공백정리
    text = text.strip()
    return text
[cleaner.py 파일내용끝]
```

```
# ============================================
# 2. tagger.py
# ============================================
[tagger.py 파일내용시작]
from patterns import *
def apply_tags(text: str) -> str:
    tags = []
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

10/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    for word in AUDITION_WORDS:
        if word in text:
            tags.append("<심리_오디션>")
    for word in CASTING_WORDS:
        if word in text:
            tags.append("<심리_드라마캐스팅>")
    for word in START_WORDS:
        if word in text:
            tags.append("<심리_시작>")
    if tags:
        return text + " " + " ".join(set(tags))
    return text
[tagger.py 파일내용끝]
```

```
# ============================================
# 3. patterns.py
# ============================================
```

```
[patterns.py 파일내용시작]
```

```
# 스탭모집제외
STAFF_HINTS = [
    "스탭", "스태프", "촬영도와주실", "현장도와주실", "스탭모집"
]
# 광고제거
AD_HINTS = [
    "문의", "카톡", "카카오", "체험단", "DM주세요", "홍보"
]
# 심리태그용
AUDITION_WORDS = ["오디션", "프로필", "연기영상", "캐스팅디렉터"]
CASTING_WORDS  = ["드라마", "배역", "출연"]
START_WORDS    = ["시작", "입문", "고민"]
# 제거규칙
def is_staff_post(text: str) -> bool:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

11/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
    return any(word in text for word in STAFF_HINTS)
```

```
def is_ad_post(text: str) -> bool:
    return any(word in text for word in AD_HINTS)
```

```
[patterns.py 파일내용끝]
```

```
# ============================================
# 4. post.py
```

```
# ============================================
```

```
[post.py 파일내용시작]
```

```
from cleaner import clean_text
from tagger import apply_tags
from patterns import is_staff_post, is_ad_post
```

```
def process_post(post: dict) -> dict:
    title = clean_text(post.get("title", ""))
    body = clean_text(post.get("body", ""))
    comments = clean_text(post.get("comments", ""))
```

```
    merged = f"{title} {body} {comments}".strip()
```

```
    if not merged:
        return None
    if is_staff_post(merged):
        return None
    if is_ad_post(merged):
        return None
    merged = apply_tags(merged)
```

```
    return {
        "title": title,
        "body": body,
        "comments": comments,
        "fulltext": merged
    }
[post.py 파일내용끝]
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

12/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
# ============================================
# 5. saver.py
# ============================================
```

```
[saver.py 파일내용시작]
import csv
import os
```

```
def save_csv(filename: str, rows: list):
    os.makedirs("output", exist_ok=True)
    filepath = os.path.join("output", filename)
    with open(filepath, "w", encoding="utf-8-sig", newline="") as f:
        writer = csv.writer(f)
        writer.writerow(["title", "body", "comments", "fulltext"])
        for r in rows:
            writer.writerow([r["title"], r["body"], r["comments"], r["fulltext"]]
    return filepath
[saver.py 파일내용끝]
```

```
# ============================================
# 6. logger.py
# ============================================
[logger.py 파일내용시작]
```

```
import os, datetime
```

```
def log(message: str):
    os.makedirs("logs", exist_ok=True)
    now = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    with open("logs/crawler.log", "a", encoding="utf-8-sig") as f:
        f.write(f"[{now}] {message}\n")
```

```
[logger.py 파일내용끝]
```

```
# ============================================
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

13/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
# 7. fetcher.py
```
