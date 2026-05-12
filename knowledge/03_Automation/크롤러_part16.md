
정리하면 지금 파이프라인은 이미 제대로 짜여 있고, 형 컴퓨터나 파이썬은 고장 아냐. [사 실] 

지금 막힌 건 그 사이트가 우리를 막았기 때문이고, 내가 준 패치(fetch_url 강화 + 디버그 출력) 넣으면 

**1.** 어디서 막히는지 눈으로 확인 가능 

**2.** 막 안 막는 주소로는 바로 TOP20 결과 만들어 볼 수 있어. [사실] 

형 이대로 수정 적용하고 한 번 다시 돌려봐. 

콘솔에 찍힌 `[DEBUG] 요청 ... -> 상태코드 ...` 줄만 그대로 복사해서 보여주면, 그 다음 부터는 내가 거기 맞춰서 더 때리는 쪽으로 갈 수 있어. [사실] 

**==> picture [489 x 39] intentionally omitted <==**

전부 패치된 py파일 만들어줘 

형, 전체 다 합쳐서 한 번에 넣을 수 있는 최종본 줄게. [사실] 이번 버전은 아래까지 반영돼 있다. [사실] 

게시판 글 최대 50개만 따라가기 (무한 대기 방지) 

- 진행 상황 출력 ( `[INFO] 글 x/50 수집 중:` ) 

- 요청마다 상태코드 디버그 출력 ( `[DEBUG] 요청 ... -> 상태코드 ...` ) 

사이트가 막는지(403 등) 확인 가능하도록 `fetch_url()` 개선 

- 헤더 강화 (한국어 우선, Referer 흉내) 

이걸 그대로 `crawler.py` 로 저장하면 된다. 형은 기존 파일 전부 지우고 이걸 복사-붙여 넣기 하면 돼. [사실] 

python 코드 복사 

```
# -*- coding: utf-8 -*-
# 형전용크롤러 v0 (LIMIT=50 + DEBUG 버전)
#
# 기능흐름:
# 1) input_urls.txt 에있는주소들을읽는다.
# 2) 각주소에서텍스트(본문)를가져온다.
#    - 블로그: 페이지본문전체
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

72/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
#    - 유튜브영상: 설명란
```

- _`#    -`_ `커뮤니티` _`:`_ `게시판` _`URL`_ `이면 게시글 링크들을 따라가서 글 본문을 긁는다` _`#      (`_ `최대` _`50`_ `개 글까지만 따라간다` _`.`_ `폭주 방지` _`) # 3)`_ `모은 텍스트를 정리` _`(`_ `클리닝` _`)`_ `한다` _`.`_ 

- _`# 4)`_ `단어를 세고 많이 나온 단어` _`TOP 20`_ `을 뽑는다` _`.`_ 

- _`# 5) output_top20.txt`_ `로 저장한다` _`.`_ 

```
#
```

```
# 필요라이브러리 (한번만설치):
```

```
#   pip install requests beautifulsoup4
```

```
#
```

## _`#`_ `사용법` _`:`_ 

- _`#   1.`_ `같은 폴더` _`(crawler_project)`_ `에` _`input_urls.txt`_ `파일을 만든다` _`.`_ 

- _`#      ->`_ `한 줄에 하나씩 주소` _`(URL)`_ `적는다` _`.`_ 

- _`#   2.`_ `같은 폴더에 이 파일을` _`crawler.py`_ `라는 이름으로 저장한다` _`.`_ 

- _`#   3.`_ `명령프롬프트` _`(cmd)`_ `에서 그 폴더로 이동 후` _`:`_ 

- _`#        python crawler.py #`_ `실행한다` _`.`_ 

- _`#   4.`_ `실행이 끝나면` _`output_top20.txt`_ `가 생긴다` _`.`_ 

```
#
```

```
# 주의:
```

- _`# -`_ `로그인 필요한 글은 못 긁음` _`.`_ 

- _`# -`_ `어떤 사이트는 자동요청` _`(`_ `봇` _`)`_ `을 막아서 상태코드` _`403`_ `등으로 거절할 수 있음` _`.`_ 

- _`# -`_ `유튜브는` _`"`_ `설명란` _`"`_ `만 긁음` _`(`_ `댓글` _`/`_ `자막은 안 긁음` _`).`_ 

```
#
```

```
# Windows 11 / Python3 기준.
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

## _`########################################`_ 

```
# 0. 기본설정
########################################
```

```
# 분석에서제외할짧은조사/불필요단어들
```

```
STOPWORDS = {
```

**==> picture [249 x 47] intentionally omitted <==**

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

73/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
"그리고", "하지만", "또", "더", "좀",
```

```
"저", "너", "나", "제", "제발", "우리",
"입니다", "합니다", "근데",
}
```

```
# 너무짧은단어(예: 한글자)는버린다
MIN_WORD_LEN = 2
```

```
# 요청보낼때사이트가 "사람브라우저인가?" 라고느끼도록위장용헤더
HEADERS = {
```

```
"User-Agent": (
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) "
"Chrome/120.0.0.0 Safari/537.36"
    ),
"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
"Accept-Language": "ko-KR,ko;q=0.9,en-US;q=0.6,en;q=0.4",
"Connection": "keep-alive",
}
```

```
########################################
# 1. HTML → 사람이읽는텍스트
########################################
```

```
defextract_visible_text(html):
```

```
# HTML에서 script/style/nav/footer 등쓸모없는부분제거하고
# 순수텍스트만뽑는다.
    soup = BeautifulSoup(html, "html.parser")
```

```
for tag_name in ["script", "style", "noscript", "header", "footer", "nav", "fo
for t in soup.find_all(tag_name):
            t.decompose()
```

```
    text = soup.get_text(separator="\n")
    text = re.sub(r"\s+", " ", text)
return text.strip()
```

```
defextract_youtube_description(html):
```

```
# 유튜브설명란(영상아래 "더보기" 영역비슷한곳) 추출시도
    soup = BeautifulSoup(html, "html.parser")
```

```
# 1) og:description 메타태그우선시도
```

```
    og_desc = soup.find("meta", {"property": "og:description"})
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

74/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
if og_desc and og_desc.get("content"):
```

```
return og_desc.get("content").strip()
```

```
# 2) shortDescription 패턴(백업시도)
    m = re.search(r'"shortDescription":"(.*?)"', html, flags=re.DOTALL)
if m:
```

```
1)
try:
# 유튜브페이지텍스트는 \n, \uXXXX 같은이스케이프가많아서정리
            desc_raw = desc_raw.encode("utf-8").decode("unicode_escape", errors="
except Exception:
pass
```

```
        desc_raw = desc_raw.replace("\\n", " ")
r"\s+", " ", desc_raw)
return desc_raw.strip()
```

```
return""
```

```
########################################
```

```
# 2. 커뮤니티게시판처리
########################################
```

```
defis_same_domain(url_a, url_b):
```

```
# 두 URL이같은도메인(사이트)인지비교
    da = urlparse(url_a).netloc
    db = urlparse(url_b).netloc
return da == db
```

```
defcollect_links_from_board_page(base_url, html):
```

- _`#`_ `게시판 목록 페이지에서 내부 글 링크 후보들을 추출한다` _`.`_ 

```
# 같은도메인이고, 경로가너무짧지않은것만사용.