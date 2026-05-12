```
#
```

```
# 제한:
```

```
# - 로그인필요한글은못긁음
# - 사이트가막을수도있음
```

```
# - 유튜브는 "설명란" 위주만긁음 (댓글/자막은안긁음)
#
# Windows 11 / Python3 기준
```

```
import os
import re
import time
import requests
from urllib.parse import urljoin, urlparse
from bs4 import BeautifulSoup
from collections import Counter
```

```
########################################
# 0. 기본설정
########################################
```

```
# 분석에서제외할짧은조사/불필요단어들
```

```
STOPWORDS = {
"은", "는", "이", "가", "을", "를",
"에", "에서", "으로", "로", "과", "와",
"하다", "합니다", "했습니다", "것", "거",
"그리고", "하지만", "또", "더", "좀",
"저", "너", "나", "제", "제발", "우리",
"입니다", "합니다", "근데",
}
```

```
# 너무짧은단어(예: 한글자)는버린다
MIN_WORD_LEN = 2
```

```
# 요청보낼때같이보내는헤더 (사이트가덜의심하게)
HEADERS = {
"User-Agent": (
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) "
"Chrome/120.0.0.0 Safari/537.36"
    )
}
```

## _`########################################`_ 

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

54/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
# 1. HTML → 사람이읽는텍스트
```

```
########################################
```

```
defextract_visible_text(html):
# HTML에서 script/style/nav/footer 등쓸모없는부분제거하고
# 순수텍스트만뽑는다.
    soup = BeautifulSoup(html, "html.parser")
for tag_name in ["script", "style", "noscript", "header", "footer", "nav", "fo
for t in soup.find_all(tag_name):
            t.decompose()
    text = soup.get_text(separator="\n")
    text = re.sub(r"\s+", " ", text)
return text.strip()
```

```
defextract_youtube_description(html):
# 유튜브설명란(영상아래 "더보기" 영역비슷한곳)을추출시도
    soup = BeautifulSoup(html, "html.parser")
# 1) og:description 메타태그우선시도
    og_desc = soup.find("meta", {"property": "og:description"})
if og_desc and og_desc.get("content"):
return og_desc.get("content").strip()
# 2) shortDescription 패턴(백업시도)
    m = re.search(r'"shortDescription":"(.*?)"', html, flags=re.DOTALL)
if m:
        desc_raw = m.group(1)
try:
# 유튜브페이지텍스트는 \n, \uXXXX 같은이스케이프가많아서정리해준다
            desc_raw = desc_raw.encode("utf-8").decode("unicode_escape", errors="
except Exception:
pass
        desc_raw = desc_raw.replace("\\n", " ")
        desc_raw = re.sub(r"\s+", " ", desc_raw)
return desc_raw.strip()
return""
```

```
########################################
```

```
# 2. 커뮤니티게시판처리
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

55/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
########################################
```

```
defis_same_domain(url_a, url_b):
# 두 URL이같은도메인(사이트)인지비교
    da = urlparse(url_a).netloc
    db = urlparse(url_b).netloc
return da == db
defcollect_links_from_board_page(base_url, html):
# 게시판목록페이지에서내부글링크후보들을추출한다.
# 같은도메인이고, 경로가너무짧지않은것만사용.
    soup = BeautifulSoup(html, "html.parser")
    links = set()
for a in soup.find_all("a", href=True):
        abs_url = urljoin(base_url, a["href"])
# 같은도메인만
ifnot is_same_domain(base_url, abs_url):
continue
# '#'만있는링크제외
if abs_url.endswith("#"):
continue
# 경로가지나치게짧으면제외 (예: 그냥 '/')
        path_len = len(urlparse(abs_url).path)
if path_len < 2:
continue
        links.add(abs_url)
returnlist(links)
```

```
deffetch_url(url):
# URL에서 HTML을가져온다. 실패하면 "" 리턴.
try:
        resp = requests.get(url, headers=HEADERS, timeout=10)
if resp.status_code == 200:
return resp.text
else:
return""
except Exception:
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

56/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
return""
```

```
defcrawl_single_page(url):
# 단일페이지(블로그글, 커뮤니티글, 유튜브영상등)에서
# 본문텍스트를뽑는다.
    html = fetch_url(url)
ifnot html:
return""
```

```
# 유튜브면설명란만
if"youtube.com/watch"in url or"youtu.be/"in url:
        desc = extract_youtube_description(html)
return desc
```

```
# 일반페이지면전체가시텍스트
    text = extract_visible_text(html)
return text
```

```
defcrawl_board_deep(url):
# 커뮤니티게시판 URL이라고보고처리한다.
# 동작:
# 1) 그 URL(게시판목록페이지) 자체텍스트추출
# 2) 그페이지안에서찾아낸글링크들중일부만따라가서
#    각글의본문뽑기
# 시간을너무끌지않도록글개수를제한한다.
```

```
    MAX_POSTS = 20# 이숫자만큼만따라간다 (필요하면늘릴수있음)
```

```
print(f"[INFO] 게시판페이지진입: {url}")
    main_html = fetch_url(url)
ifnot main_html:
print("[WARN] 게시판페이지로드실패")
return""
```

```
    out_texts = []
```

```
# 현재(목록) 페이지자체텍스트도일단포함
    main_text = extract_visible_text(main_html)
if main_text:
        out_texts.append(main_text)
```

```
# 내부글링크들후보뽑기
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

57/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
    candidate_links = collect_links_from_board_page(url, main_html)
```

```
# 너무많으면앞부분만사용해서폭주방지
    candidate_links = candidate_links[:MAX_POSTS]
```

```
    seen = set()
    post_index = 0
```

```
for link in candidate_links:
if link in seen:
continue
        seen.add(link)
```

```
        post_index += 1
print(f"[INFO]  글{post_index}/{len(candidate_links)} 수집중: {link}")
```

```
# 사이트에너무빠르게연속으로요청하지않도록잠깐쉰다.
        time.sleep(0.5)
```

```
        page_text = crawl_single_page(link)
if page_text: