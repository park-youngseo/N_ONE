크롤러 

 

다음 

코드 

## 형, 이제 코드 단계 열었다. [사실] 

아래는 형 폴더 `crawler_project` 안에 넣고 돌릴 파이썬 기본 버전(v0) 전체 코드야. [사 실] 

출력 방식은 우리 약속 그대로다: 

`input_urls.txt` 읽는다. 

- 주소들 돌면서 글 내용을 싹 긁는다. (블로그 본문 / 유튜브 설명란 / 커뮤니티 글 본 문) 

- 전부 합친다. 쓰레기 치운다. 조사(은/는/이/가…) 버린다. 

많이 나온 단어 TOP 20 뽑는다. 

- `output_top20.txt` 로 저장한다. 

## [사실] 

이 코드는 “로그인 필요 없는 공개 페이지” 기준으로 만든 거다. [사실] 

사이트가 막아버리는 경우나 특이한 구조(자바스크립트로만 내용 뿌리는 페이지)는 안 될 수도 있다. 그건 나중 단계에서 사이트별 보정으로 개선 가능하다. [추론] 

형은 지금 그냥 이 파일 하나만 `.py` 로 저장하면 된다. 이름은 `crawler.py` 로 고정 권장. [사실] 

python 코드 복사 

```
# -*- coding: utf-8 -*-
"""
crawler.py
형전용크롤러 v0
```

## `기능 흐름:` 

- `1) input_urls.txt 에 있는 주소들을 읽는다.` 

- `2) 각 주소에서 텍스트(본문)를 가져온다.` 

- `블로그: 페이지 본문 전체` 

- `유튜브 영상: 설명란` 

- `커뮤니티: 게시글 본문 (또는 게시판 목록에서 따라간 글들도 시도)` 

- `3) 모은 텍스트를 정리(클리닝)한다.` 

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

23/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
4) 단어를세고많이나온단어 TOP 20을뽑는다.
```

```
5) output_top20.txt 로저장한다.
```

```
필요라이브러리:
pip install requests beautifulsoup4
```

```
Windows 11 / Python3 기준
"""
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
###############################################################################
# 0. 공통설정
```

```
###############################################################################
```

```
# 너무시끄러운조사/불필요단어(너무짧은말등)를버리기위한리스트.
# 이단어들은빈도계산에서제외된다.
STOPWORDS = {
"은", "는", "이", "가", "을", "를",
"에", "에서", "으로", "로", "과", "와",
"하다", "합니다", "했습니다", "것", "거",
"그리고", "하지만", "또", "더", "좀",
"저", "너", "나", "제", "제발", "우리",
"입니다", "합니다", "합니다.", "근데",
}
```

```
# 너무짧은단어(예: 한글자짜리)는버릴거다.
MIN_WORD_LEN = 2
```

```
# User-Agent: 사이트가 "이거봇이네?" 하고막는걸조금덜하게하는헤더.
HEADERS = {
```

```
"User-Agent": (
```

```
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) "
"Chrome/120.0.0.0 Safari/537.36"
    )
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

24/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
}
```

```
###############################################################################
# 1. 유틸: HTML에서사람이읽는본문텍스트만뽑기
###############################################################################
```

```
defextract_visible_text(html):
"""
    HTML 전체에서 script/style/nav/footer/header 같은
쓸모없는영역은빼고순수텍스트만모은다.
    """
    soup = BeautifulSoup(html, "html.parser")
# 보기불필요한태그싹제거
for tag_name in ["script", "style", "noscript", "header", "footer", "nav", "fo
for t in soup.find_all(tag_name):
            t.decompose()
```

```
# 텍스트만추출
    text = soup.get_text(separator="\n")
# 공백정리
    text = re.sub(r"\s+", " ", text)
return text.strip()
```

```
defextract_youtube_description(html):
"""
유튜브영상페이지 HTML에서설명란(Description)을추출하려고시도.
완벽하진않지만, 기본적으로 og:description 메타태그를먼저본다.
    """
    soup = BeautifulSoup(html, "html.parser")
```

```
# 1) og:description 시도
    og_desc = soup.find("meta", {"property": "og:description"})
if og_desc and og_desc.get("content"):
return og_desc.get("content").strip()
# 2) 위방식이안될경우: 페이지안에 "shortDescription" 같은패턴찾기 (단순정
    m = re.search(r'"shortDescription":"(.*?)"', html, flags=re.DOTALL)
if m:
        desc_raw = m.group(1)
# 유튜브는 \n, \u0026 등이스케이프문자가섞일수있다.
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

25/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
        desc_raw = desc_raw.encode("utf-8").decode("unicode_escape", errors="igno
        desc_raw = desc_raw.replace("\\n", " ")
        desc_raw = re.sub(r"\s+", " ", desc_raw)
return desc_raw.strip()
```

```
# 못찾으면빈문자열
return""
```

```
###############################################################################
# 2. 커뮤니티게시판처리: 목록페이지에서글상세들까지시도
###############################################################################
```

```
defis_same_domain(url_a, url_b):
"""
두 URL이같은도메인인지간단히비교
    """
    da = urlparse(url_a).netloc
    db = urlparse(url_b).netloc
return da == db
```

```
defcollect_links_from_board_page(base_url, html):
"""
```

```
커뮤니티게시판목록같은페이지일수있다고가정하고,
```

```
같은도메인내에서 '글상세로보이는링크들' 후보를긁어온다.
```

```
완벽자동탐지는어렵다. 기본전략:
```

```
    - <a> 태그전부훑는다.
```

```
    - 같은도메인이고, 너무짧지않은경로면후보로넣는다.
```

```
    - 중복제거.
```

```
나중에사이트별로더정확하게손볼수있다.
    """
```

```
    soup = BeautifulSoup(html, "html.parser")
    links = set()
```

```
for a in soup.find_all("a", href=True):
        abs_url = urljoin(base_url, a["href"])
```

```
# 같은사이트내부만본다
```

```
ifnot is_same_domain(base_url, abs_url):
continue
# Query나 anchor(#)만있는이상한건버림
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

26/85 

25. 11. 30. 오후 12:58 

크롤러 

 
