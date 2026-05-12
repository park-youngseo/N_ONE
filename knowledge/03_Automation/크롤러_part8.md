```

```
# - 로그인필요한글은못긁는다.
```

- _`# -`_ `어떤 사이트는 막을 수도 있다` _`.`_ 

- _`# -`_ `유튜브는` _`"`_ `설명란` _`"`_ `위주만 긁는다` _`(`_ `댓글` _`/`_ `자막` _`X).`_ 

```
#
```

```
# 윈도우 11 / Python3 기준.
```

```
import os
import re
import time
import requests
from urllib.parse import urljoin, urlparse
from bs4 import BeautifulSoup
```

```
from collections import Counter
```

## _`########################################`_ 

```
# 0. 기본설정
```

```
########################################
```

## _`# STOPWORDS:`_ 

- _`#`_ `분석에서 빼버릴 짧은 조사` _`/`_ `불필요 단어 목록 STOPWORDS = {` 

**==> picture [220 x 11] intentionally omitted <==**

**==> picture [241 x 11] intentionally omitted <==**

**==> picture [249 x 65] intentionally omitted <==**

```
}
```

```
# 너무짧은단어는버린다 (예: 한글자단어)
MIN_WORD_LEN = 2
```

- _`#`_ `일부 사이트가 봇 막는 걸 줄이기 위한 요청 헤더 HEADERS = {` 

```
"User-Agent": (
```

- `"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "` 

- `"AppleWebKit/537.36 (KHTML, like Gecko) "` 

- `"Chrome/120.0.0.0 Safari/537.36"` 

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

36/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
    )
```

```
}
```

```
########################################
# 1. HTML → 텍스트뽑기
########################################
```

```
defextract_visible_text(html):
```

```
# HTML 안에서 script/style/nav/footer 같은건버리고
# 사람이읽는텍스트만뽑는다.
    soup = BeautifulSoup(html, "html.parser")
```

```
for tag_name in ["script", "style", "noscript", "header", "footer", "nav", "fo
for t in soup.find_all(tag_name):
            t.decompose()
    text = soup.get_text(separator="\n")
    text = re.sub(r"\s+", " ", text)
return text.strip()
```

```
defextract_youtube_description(html):
# 유튜브설명란(영상아래 "더보기" 영역비슷한곳) 추출시도
    soup = BeautifulSoup(html, "html.parser")
```

```
# 1) og:description 메타태그먼저시도
    og_desc = soup.find("meta", {"property": "og:description"})
if og_desc and og_desc.get("content"):
return og_desc.get("content").strip()
```

```
# 2) shortDescription라는부분을정규식으로찾기 (백업방식)
    m = re.search(r'"shortDescription":"(.*?)"', html, flags=re.DOTALL)
if m:
        desc_raw = m.group(1)
# \n, \uXXXX 등특수표현정리
try:
            desc_raw = desc_raw.encode("utf-8").decode("unicode_escape", errors="
except Exception:
pass
        desc_raw = desc_raw.replace("\\n", " ")
        desc_raw = re.sub(r"\s+", " ", desc_raw)
return desc_raw.strip()
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

37/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
return""
```

```
########################################
# 2. 커뮤니티게시판 (목록 → 글본문들까지)
########################################
```

```
defis_same_domain(url_a, url_b):
```

```
# 두 URL이같은사이트(도메인)인지비교
    da = urlparse(url_a).netloc
    db = urlparse(url_b).netloc
return da == db
```

```
defcollect_links_from_board_page(base_url, html):
```

```
# 게시판목록페이지일수도있는화면에서
# 같은도메인안의글링크후보들을뽑는다.
    soup = BeautifulSoup(html, "html.parser")
    links = set()
```

```
for a in soup.find_all("a", href=True):
        abs_url = urljoin(base_url, a["href"])
```

```
# 같은도메인만허용
ifnot is_same_domain(base_url, abs_url):
continue
# '#'만있는링크는버림
if abs_url.endswith("#"):
continue
```

```
# 너무짧은경로는버림 (예: '/')
        path_len = len(urlparse(abs_url).path)
if path_len < 2:
continue
        links.add(abs_url)
```

```
returnlist(links)
```

```
deffetch_url(url):
# URL에서 HTML을가져온다. 실패하면 "" 리턴.
```

```
try:
```

```
        resp = requests.get(url, headers=HEADERS, timeout=10)
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

38/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
if resp.status_code == 200:
return resp.text
else:
return""
except Exception:
return""
defcrawl_single_page(url):
# 블로그글/커뮤니티글/유튜브영상등 "단일페이지"의
# 본문텍스트를뽑는다.
    html = fetch_url(url)
ifnot html:
return""
```

```
# 유튜브면설명란만뽑는다.
if"youtube.com/watch"in url or"youtu.be/"in url:
        desc = extract_youtube_description(html)
return desc
```

```
# 일반페이지면전체가시텍스트추출
    text = extract_visible_text(html)
return text
```

```
defcrawl_board_deep(url):
# 커뮤니티게시판 URL이라고가정:
# 1) 그페이지자체텍스트
# 2) 그페이지안의글링크들을따라가서
#    각각의본문텍스트까지긁는다.
# 전부합쳐서하나의큰문자열로준다.
    main_html = fetch_url(url)
ifnot main_html:
return""
    out_texts = []
```

```
# 현재페이지본문텍스트
    out_texts.append(extract_visible_text(main_html))
```

```
# 내부글링크들후보
    candidate_links = collect_links_from_board_page(url, main_html)
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

39/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
    seen = set()
for link in candidate_links:
if link in seen:
continue