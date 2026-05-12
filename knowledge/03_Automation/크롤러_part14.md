
- _`#`_ `게시판 목록 페이지에서 내부 글 링크 후보들을 추출한다` _`. #`_ `같은 도메인이고` _`,`_ `경로가 너무 짧지 않은 것만 사용` _`.`_ 

- `soup = BeautifulSoup(html, "html.parser") links = set()` 

```
for a in soup.find_all("a", href=True):
        abs_url = urljoin(base_url, a["href"])
```

```
# 같은도메인만
ifnot is_same_domain(base_url, abs_url):
continue
# '#'만있는링크제외
if abs_url.endswith("#"):
continue
```

```
# 경로가지나치게짧으면제외 (예: 그냥 '/')
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
```

```
# URL에서 HTML을가져온다. 실패하면 "" 리턴.
try:
        resp = requests.get(url, headers=HEADERS, timeout=10)
if resp.status_code == 200:
return resp.text
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

62/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
else:
```

```
return""
except Exception:
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
defcrawl_board_deep(url):
# 커뮤니티게시판 URL이라고보고처리한다.
# 동작:
# 1) 그 URL(게시판목록페이지) 자체텍스트추출
# 2) 그페이지안에서찾아낸글링크들중최대 50개만따라가서
#    각글의본문뽑기
# 너무오래안걸리게제한넣음.
    MAX_POSTS = 50# 한번에따라갈최대글개수
print(f"[INFO] 게시판페이지진입: {url}")
    main_html = fetch_url(url)
ifnot main_html:
print("[WARN] 게시판페이지로드실패")
return""
    out_texts = []
# 현재(목록) 페이지자체텍스트도일단포함
    main_text = extract_visible_text(main_html)
if main_text:
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

63/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
        out_texts.append(main_text)
```

```
# 내부글링크들후보뽑기
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
# - 이메일/광고코드처럼보이는것
# - 특수문자더미
ifnot raw_text:
return""
```

```
    text = raw_text
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

64/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
# URL 제거
    text = re.sub(
r"https?://[A-Za-z0-9\.\-_/=\?&%#:]+",
" ",
        text
    )
# 이메일비슷한문자열제거
    text = re.sub(r"[A-Za-z0-9._%+-]+@[A-Za-z0-9._%+-]+", " ", text)
# 특수문자제거 (한글/영문/숫자/공백만남기기)
    text = re.sub(r"[^0-9A-Za-z가-힣ㄱ-ㅎㅏ-ㅣ\s]", " ", text)
# 공백정리
    text = re.sub(r"\s+", " ", text)
return text.strip()
deftokenize_korean(text):
# 텍스트를단어로쪼갠다.
# 규칙:
# - 한글/영문/숫자덩어리만단어로본다.
# - 한글자등너무짧은건버린다.
# - STOPWORDS(은/는/이/가등) 버린다.
    tokens = re.findall(r"[0-9A-Za-z가-힣]+", text)
    cleaned_tokens = []
for tok in tokens:
iflen(tok) < MIN_WORD_LEN:
continue
if tok in STOPWORDS:
continue
        cleaned_tokens.append(tok)
```

```
return cleaned_tokens
```

```
defcount_top_words(all_text, top_n=20):
# 전체텍스트에서단어빈도를세고
# 상위 top_n개를 (단어, 빈도) 리스트로준다.
    tokens = tokenize_korean(all_text)
    counter = Counter(tokens)
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

65/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
    most_common = counter.most_common(top_n)
return most_common
```

```
########################################
# 4. 메인흐름 (input_urls.txt → output_top20.txt)
########################################
defread_input_urls(file_path):
# input_urls.txt 에서 URL 목록을읽는다.
    urls = []
ifnot os.path.exists(file_path):
return urls
withopen(file_path, "r", encoding="utf-8") as f:
for line in f:
            u = line.strip()
if u:
                urls.append(u)
return urls
defcrawl_all(urls):
# 모든 URL을돌면서텍스트를긁어온다.
#
# 규칙:
# - 유튜브: 설명란만
# - 나머지: 커뮤니티게시판일수도있다고보고 crawl_board_deep 먼저시도
#           deep 결과가충분히길면 deep 사용
#           deep 결과가비거나짧으면그냥해당페이지 1장만 crawl_single_page
    texts = []
for url in urls:
print(f"[INFO] 수집중: {url}")
# 유튜브
if"youtube.com/watch"in url or"youtu.be/"in url:
            t = crawl_single_page(url)
if t:
                texts.append(t)
continue
# 커뮤니티게시판비슷하다고보고 deep 시도
        deep_text = crawl_board_deep(url)
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

66/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
# deep 결과가어느정도길면그걸사용
if deep_text andlen(deep_text.split()) > 20:
            texts.append(deep_text)