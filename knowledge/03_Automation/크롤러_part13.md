            out_texts.append(page_text)
```

```
    joined = "\n".join(out_texts)
print(f"[INFO] 게시판수집종료. 총{post_index}개글반영.")
return joined
```

```
########################################
# 3. 텍스트정리 / 토큰나누기 / 단어세기
########################################
```

```
defclean_text(raw_text):
```

```
# 원본모음텍스트에서필요없는것들제거:
# - URL
```

```
# - 이메일/광고코드처럼보이는것
# - 특수문자더미
ifnot raw_text:
return""
```

```
    text = raw_text
```

```
# URL 제거
    text = re.sub(
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

58/85 

25. 11. 30. 오후 12:58 크롤러 

 

```
r"https?://[A-Za-z0-9\.\-_/=\?&%#:]+",
" ",
        text
    )
# 이메일비슷한문자열제거
    text = re.sub(r"[A-Za-z0-9._%+-]+@[A-Za-z0-9._%+-]+",
```

**==> picture [489 x 39] intentionally omitted <==**

## 글을 50개까지 하자 

형, 확인. 이제 한 게시판에서 따라가는 글 수 한도를 20개 → 50개로 올린다. [사실] 

아래는 수정 반영된 최종 전체본 `crawler.py` 다. 이걸 통째로 저장해서 쓰면 된다. (기존 파일 전부 지우고 그대로 붙여넣기 → 저장) [사실] 

python 코드 복사 

```
# -*- coding: utf-8 -*-
```

```
# 형전용크롤러 v0 (REPAIR+LIMIT=50 버전)
```

```
#
```

```
# 기능흐름:
```

- _`# 1) input_urls.txt`_ `에 있는 주소들을 읽는다` _`.`_ 

- _`# 2)`_ `각 주소에서 텍스트` _`(`_ `본문` _`)`_ `를 가져온다` _`.`_ 

```
#    - 블로그: 페이지본문전체
```

```
#    - 유튜브영상: 설명란
```

- _`#    -`_ `커뮤니티` _`:`_ `게시판` _`URL`_ `이면 게시글 링크들을 따라가서 글 본문을 긁는다` 

- _`#      (`_ `최대` _`50`_ `개 글까지만 따라간다` _`.`_ `폭주 방지` _`)`_ 

- _`# 3)`_ `모은 텍스트를 정리` _`(`_ `클리닝` _`)`_ `한다` _`.`_ 

- _`# 4)`_ `단어를 세고 많이 나온 단어` _`TOP 20`_ `을 뽑는다` _`.`_ 

- _`# 5) output_top20.txt`_ `로 저장한다` _`.`_ 

```
#
```

## _`#`_ `필요 라이브러리` _`:`_ 

```
#   pip install requests beautifulsoup4
```

```
#
```

```
# 사용법:
```

- _`#   1.`_ `같은 폴더에` _`input_urls.txt`_ `만들고` _`,`_ `한 줄에 하나씩 주소` _`(URL)`_ `적는다` _`.`_ 

- _`#   2.`_ `같은 폴더에 이 파일을` _`crawler.py`_ `라는 이름으로 저장한다` _`.`_ 

- _`#   3.`_ `명령프롬프트` _`(cmd)`_ `에서 그 폴더로 이동 후` _`:`_ 

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

59/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
#        python crawler.py
#      라고실행한다.
#   4. 실행이끝나면 output_top20.txt 가생긴다.
#
# 제한:
# - 로그인필요한글은못긁음
# - 사이트가봇을막을수도있음
# - 유튜브는 "설명란"만긁음 (댓글/자막은안긁음)
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
STOPWORDS = {
```

```
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
```

```
"User-Agent": (
```

```
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) "
"Chrome/120.0.0.0 Safari/537.36"
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

60/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
}
```

```
########################################
# 1. HTML → 사람이읽는텍스트
########################################
```

```
defextract_visible_text(html):
# HTML에서 script/style/nav/footer 등쓸모없는부분제거하고
# 순수텍스트만뽑는다.
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
# 유튜브설명란(영상아래 "더보기"에있는설명)을추출시도
    soup = BeautifulSoup(html, "html.parser")
```

```
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
```

```
return""
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

61/85 

25. 11. 30. 오후 12:58 

크롤러 

 

## _`########################################`_ 

```
# 2. 커뮤니티게시판처리
########################################
```

```
defis_same_domain(url_a, url_b):
```

- _`#`_ `두` _`URL`_ `이 같은 도메인` _`(`_ `사이트` _`)`_ `인지 비교` 

```
    da = urlparse(url_a).netloc
    db = urlparse(url_b).netloc
return da == db
```

```
defcollect_links_from_board_page(base_url, html):
```