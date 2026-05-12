```
if abs_url.endswith("#"):
continue
# 너무짧은링크는버림 (ex: 그냥 "/")
        path_len = len(urlparse(abs_url).path)
if path_len < 2:
continue
        links.add(abs_url)
returnlist(links)
```

```
deffetch_url(url):
"""
    URL에서 HTML을받아온다.
실패하면 "" 반환.
    """
try:
        resp = requests.get(url, headers=HEADERS, timeout=10)
if resp.status_code == 200:
return resp.text
else:
return""
except Exception:
return""
```

```
defcrawl_single_page(url):
"""
단일페이지(블로그글본문, 커뮤니티글본문, 유튜브영상등)에서
텍스트(사람이읽는내용)를추출해서돌려준다.
    """
    html = fetch_url(url)
ifnot html:
return""
# 유튜브일경우: 설명란만추출
if"youtube.com/watch"in url or"youtu.be/"in url:
        desc = extract_youtube_description(html)
return desc
# 일반페이지(블로그글, 커뮤니티글등): 전체텍스트뽑기
    text = extract_visible_text(html)
return text
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

27/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
defcrawl_board_deep(url):
"""
커뮤니티게시판같은 URL이라가정하고:
    1) 해당 URL의 HTML을읽는다.
    2) 그페이지자체의텍스트도추출.
    3) 같은도메인내부글링크후보들을모은다.
    4) 후보링크들각각도방문해서본문텍스트를긁는다.
       (중복최소화를위해 set 사용)
최종적으로한덩어리문자열을리턴.
    """
    main_html = fetch_url(url)
ifnot main_html:
return""
```

```
    out_texts = []
```

```
# 현재페이지본문도넣는다.
    out_texts.append(extract_visible_text(main_html))
```

```
# 게시글후보링크들수집
```

```
    candidate_links = collect_links_from_board_page(url, main_html)
```

```
    seen = set()
for link in candidate_links:
if link in seen:
continue
        seen.add(link)
```

```
# 너무많이긁으면과부하위험 → 안전하게슬립조금씩
0.5)
```

```
        page_text = crawl_single_page(link)
if page_text:
            out_texts.append(page_text)
```

```
# 전부합쳐서한큰문자열로
    joined = "\n".join(out_texts)
return joined
```

```
###############################################################################
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

28/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
# 3. 본문정리(클리닝) / 단어쪼개기 / 단어빈도세기
```

```
###############################################################################
```

```
defclean_text(raw_text):
"""
원본텍스트에서불필요한것들제거:
    - URL
    - 특수기호반복
    - 너무이상한문자들
    """
ifnot raw_text:
return""
    text = raw_text
# URL 제거
    text = re.sub(
r"https?://[A-Za-z0-9\.\-_/=\?&%#:]+",
" ",
        text
    )
# 이메일/광고코드류대충제거 (단순처리)
    text = re.sub(r"[A-Za-z0-9._%+-]+@[A-Za-z0-9._%+-]+", " ", text)
# 특수문자뭉치정리
    text = re.sub(r"[^0-9A-Za-z가-힣ㄱ-ㅎㅏ-ㅣ\s]", " ", text)
# 공백정리
    text = re.sub(r"\s+", " ", text)
return text.strip()
```

```
deftokenize_korean(text):
"""
한/영단어를단순하게잘라내기.
고급형태소분석기는안썼다. (추가설치없이가려고함)
규칙:
    - 한글/영어/숫자덩어리를단어로본다.
    - 너무짧은건버린다 (예: "은","는"...)
    """
# 한글/영문/숫자덩어리추출
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

29/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
# \w 는 _ 까지포함하므로직접정의
    tokens = re.findall(r"[0-9A-Za-z가-힣]+", text)
```

```
    cleaned_tokens = []
for tok in tokens:
# 한글자밖에없는토큰은버린다.
iflen(tok) < MIN_WORD_LEN:
continue
# stopwords(조사등) 버리기
if tok in STOPWORDS:
continue
        cleaned_tokens.append(tok)
```

```
return cleaned_tokens
```

```
defcount_top_words(all_text, top_n=20):
"""
전체텍스트에서단어빈도를세고
상위 top_n개 (기본 20개)를 [(단어, 빈도), ...] 형태로돌려준다.
    """
    tokens = tokenize_korean(all_text)
    counter = Counter(tokens)
    most_common = counter.most_common(top_n)
return most_common
```

```
###############################################################################
# 4. 메인흐름:
#    - input_urls.txt 읽기
#    - 각 URL 크롤링
#    - 텍스트합치기
#    - 정리/클리닝
#    - 단어빈도 → TOP20
#    - output_top20.txt 저장
###############################################################################
```

## `def read_input_urls(file_path):` 

```
"""
    input_urls.txt 에서줄마다 URL 읽어서리스트로돌려준다.
빈줄은무시.
    """
    urls = []
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

30/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
ifnot os.path.exists(file_path):
return urls
```

```
withopen(file_path, "r", encoding="utf-8") as f:
for line in f:
            u = line.strip()
if u:
                urls.append(u)
return urls
```

```
defcrawl_all(urls):
"""
모든 URL을돌면서텍스트를긁어온다.
규칙:
    - youtube.com/watch 또는 youtu.be 인경우 → 유튜브설명란만뽑기
    - 그외:
        1) 커뮤니티게시판(목록)일수도있다고가정하고 crawl_board_deep 시도
        2) 너무결과가비었으면그냥단일페이지로 crawl_single_page
    """
    texts = []
for url in urls:
print(f"[INFO] 수집중: {url}")
# 유튜브면바로설명란
if"youtube.com/watch"in url or"youtu.be/"in url:
            t = crawl_single_page(url)
if t:
                texts.append(t)
continue
# 커뮤니티게시판일수도있다고보고 deep 시도
        deep_text = crawl_board_deep(url)
# deep_text가뭔가있으면 OK
if deep_text andlen(deep_text.split()) > 20:
            texts.append(deep_text)
continue
# deep 실패/너무짧으면그냥단일페이지본문만
        single_text = crawl_single_page(url)
if single_text:
            texts.append(single_text)
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

31/85 

25. 11. 30. 오후 12:58 

크롤러 