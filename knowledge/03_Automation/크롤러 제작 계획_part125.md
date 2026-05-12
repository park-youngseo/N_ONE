returnany(word in text for word in STAFF_HINTS)
defis_ad_post(text: str) -> bool:
returnany(word in text for word in AD_HINTS)
```

```
[patterns.py 파일내용끝]
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

615/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

 

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
defprocess_post(post: dict) -> dict:
    title = clean_text(post.get("title", ""))
""
    body = clean_text(post.get("body", ))
    comments = clean_text(post.get("comments", ""))
```

```
    merged = f"{title}{body}{comments}".strip()
```

```
ifnot merged:
returnNone
if is_staff_post(merged):
returnNone
if is_ad_post(merged):
returnNone
```

```
    merged = apply_tags(merged)
```

```
return {
"title": title,
"body": body,
"comments": comments,
"fulltext": merged
    }
```

```
[post.py 파일내용끝]
```

```
# ============================================
# 5. saver.py
# ============================================
```

```
[saver.py 파일내용시작]
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

616/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

 

```
import csv
import os
```

```
defsave_csv(filename: str, rows: list):
    os.makedirs("output", exist_ok=True)
    filepath = os.path.join("output", filename)
```

```
withopen(filepath, "w", encoding="utf-8-sig", newline="") as f:
        writer = csv.writer(f)
        writer.writerow(["title", "body", "comments", "fulltext"])
for r in rows:
            writer.writerow([r["title"], r["body"], r["comments"], r["fulltext"]]
return filepath
```

```
[saver.py 파일내용끝]
```

```
# ============================================
# 6. logger.py
# ============================================
```

```
[logger.py 파일내용시작]
```

```
import os, datetime
```

```
deflog(message: str):
```

```
    os.makedirs("logs", exist_ok=True)
    now = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
withopen("logs/crawler.log", "a", encoding="utf-8-sig") as f:
        f.write(f"[{now}] {message}\n")
```

```
[logger.py 파일내용끝]
```

```
# ============================================
# 7. fetcher.py
# ============================================
```

```
[fetcher.py 파일내용시작]
```

```
import requests
from bs4 import BeautifulSoup
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

617/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

 

```
import time
import random
from post import process_post
from logger import log
headers = {
"User-Agent": "Mozilla/5.0 (Windows NT 10.0)"
}
deffetch_generic_board(url: str) -> list:
    posts = []
    log(f"크롤링시작: {url}")
try:
        page = 1
whileTrue:
            page_url = f"{url}?page={page}"
            r = requests.get(page_url, headers=headers, timeout=10)
if r.status_code != 200:
break
            soup = BeautifulSoup(r.text, "html.parser")
            items = soup.select("a")  # 모든링크우선수집
ifnot items:
break
for a in items:
                title = a.text.strip()
                href = a.get("href", "")
ifnot href:
continue
if href.startswith("/"):
                    href = url.split("/")[0] + "//" + url.split("/")[2] + href
# 상세페이지
try:
                    d = requests.get(href, headers=headers, timeout=10)
                    ds = BeautifulSoup(d.text, "html.parser")
                    body = ds.get_text(" ", strip=True)
                    comments = ""
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

618/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

 

```
                    row = {
"title": title,
"body": body,
"comments": comments
                    }
                    processed = process_post(row)
if processed:
                        posts.append(processed)
except:
continue
```

```
            page += 1
            time.sleep(random.uniform(1, 3))
```

```
except Exception as e:
        log(f"오류발생: {e}")
```

```
    log(f"수집완료: {len(posts)}개")
return posts
```

```
[fetcher.py 파일내용끝]
```

```
# ============================================
# 8. main.py
```

```
# ============================================
```

```
[main.py 파일내용시작]
```

```
from fetcher import fetch_generic_board
from saver import save_csv
from logger import log
```

```
defask_user():
print("=============================")
print("  엔오원니즈수집크롤러 v2")
print("=============================\n")
```

```
    site_name = input("[1] 사이트이름입력: ").strip()
ifnot site_name:
```

```
print("사이트이름은필수입니다.")
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

619/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

```
        exit()
```
