    soup = BeautifulSoup(html, "html.parser")
    links = set()
```

```
for a in soup.find_all("a", href=True):
```

```
        abs_url = urljoin(base_url, a["href"])
```

## _`#`_ `같은 도메인만` 

```
ifnot is_same_domain(base_url, abs_url):
```

```
continue
```

- _`# '#'`_ `만 있는 링크 제외` 

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

75/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
if abs_url.endswith("#"):
continue
```

```
# 경로가지나치게짧으면제외 (예: 그냥 '/')
        path_len = len(urlparse(abs_url).path)
if path_len < 2:
continue
        links.add(abs_url)
returnlist(links)
```

```
deffetch_url(url):
# URL에서 HTML을가져온다.
# 이번버전은상태코드(몇번인지) 찍어서확인가능하게한다.
try:
# Referer 넣어서 "같은사이트에서이동중"처럼보이게
        local_headers = dict(HEADERS)
        local_headers["Referer"] = url
        resp = requests.get(url, headers=local_headers, timeout=10)
```

```
print(f"[DEBUG] 요청{url} -> 상태코드{resp.status_code}")
```

```
# 정상응답(200)만통과
if resp.status_code == 200:
return resp.text
else:
return""
except Exception as e:
print(f"[ERROR] 요청실패: {url} ({e})")
return""
```

```
defcrawl_single_page(url):
# 단일페이지(블로그글, 커뮤니티글, 유튜브영상등)에서
# 본문텍스트를뽑는다.
    html = fetch_url(url)
ifnot html:
return""
# 유튜브면설명란만
if"youtube.com/watch"in url or"youtu.be/"in url:
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

76/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
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
# 2) 그페이지안에서찾아낸글링크들중최대 50개만따라가서
#    각글의본문뽑기
# 너무오래안걸리게제한넣음.
```

```
    MAX_POSTS = 50# 한번에따라갈최대글개수
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
```

```
    main_text = extract_visible_text(main_html)
if main_text:
        out_texts.append(main_text)
```

```
# 내부글링크들후보뽑기
```

```
    candidate_links = collect_links_from_board_page(url, main_html)
```

```
# 너무많으면앞부분만사용해서폭주방지
```

```
    candidate_links = candidate_links[:MAX_POSTS]
```

```
    seen = set()
    post_index = 0
```

```
for link in candidate_links:
```

```
if link in seen:
```

```
continue
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

77/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
        post_index += 1
print(f"[INFO]  글{post_index}/{len(candidate_links)} 수집중: {link}")
```

```
# 사이트에너무빠르게연속으로요청하지않도록잠깐쉰다.
        time.sleep(0.5)
        page_text = crawl_single_page(link)
if page_text:
            out_texts.append(page_text)
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
# 원본모음텍스트에서필요없는것들제거:
# - URL
# - 이메일/광고코드처럼보이는것
# - 특수문자더미
ifnot raw_text:
return""
    text = raw_text
# URL 제거
    text = re.sub(
r"https?://[A-Za-z0-9\.\-_/=\?&%#:]+",
" ",
    text)
# 이메일비슷한문자열제거
    text = re.sub(r"[A-Za-z0-9._%+-]+@[A-Za-z0-9._%+-]+", " ", text)
# 특수문자제거 (한글/영문/숫자/공백만남기기)
    text = re.sub(r"[^0-9A-Za-z가-힣ㄱ-ㅎㅏ-ㅣ\s]", " ", text)
# 공백정리
    text = re.sub(r"\s+", " ", text)
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

78/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
return text.strip()
```

```
deftokenize_korean(text):
# 텍스트를단어로쪼갠다.
# 규칙:
# - 한글/영문/숫자덩어리만단어로본다.
# - 한글자등너무짧은건버린다.
# - STOPWORDS(은/는/이/가등) 버린다.
```

```
    tokens = re.findall(r"[0-9A-Za-z가-힣]+", text)
```

```
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
```

```
# 전체텍스트에서단어빈도를세고
# 상위 top_n개를 (단어, 빈도) 리스트로준다.
    tokens = tokenize_korean(all_text)
    counter = Counter(tokens)
    most_common = counter.most_common(top_n)
return most_common
```

```
########################################
# 4. 메인흐름 (input_urls.txt → output_top20.txt)
########################################
```

```
defread_input_urls(file_path):
# input_urls.txt 에서 URL 목록을읽는다.
    urls = []
ifnot os.path.exists(file_path):
return urls
withopen(file_path, "r", encoding="utf-8") as f:
for line in f:
            u = line.strip()
```
