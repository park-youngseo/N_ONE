        seen.add(link)
# 너무빨리돌지않게살짝쉰다.
        time.sleep(0.5)
        page_text = crawl_single_page(link)
if page_text:
            out_texts.append(page_text)
    joined = "\n".join(out_texts)
return joined
```

```
########################################
# 3. 텍스트정리 / 토큰나누기 / 단어세기
########################################
defclean_text(raw_text):
# 원본텍스트에서필요없는것들없애기:
# - URL
# - 이메일/광고코드류
# - 특수문자뭉치
ifnot raw_text:
return""
    text = raw_text
# URL 제거
    text = re.sub(
r"https?://[A-Za-z0-9\.\-_/=\?&%#:]+",
" ",
        text
    )
# 이메일처럼생긴것제거
    text = re.sub(r"[A-Za-z0-9._%+-]+@[A-Za-z0-9._%+-]+", " ", text)
# 특수문자등제거 (한글/영문/숫자/공백만남기기)
    text = re.sub(r"[^0-9A-Za-z가-힣ㄱ-ㅎㅏ-ㅣ\s]", " ", text)
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

40/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
# 공백정리
```

```
    text = re.sub(r"\s+", " ", text)
```

```
return text.strip()
```

```
deftokenize_korean(text):
# 텍스트를단어로쪼갠다.
# 규칙:
# - 한글/영문/숫자덩어리만뽑는다.
# - 한글자같은너무짧은건버린다.
# - STOPWORDS(은,는,이,가등) 버린다.
```

```
    tokens = re.findall(r"[0-9A-Za-z가-힣]+", text)
```

```
    cleaned_tokens = []
for tok in tokens:
```

```
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
```

```
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
```

```
# input_urls.txt 에서 URL 목록을읽는다.
    urls = []
ifnot os.path.exists(file_path):
return urls
```

```
withopen(file_path, "r", encoding="utf-8") as f:
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

41/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
for line in f:
            u = line.strip()
if u:
                urls.append(u)
return urls
```

```
defcrawl_all(urls):
# 모든 URL을돌면서텍스트를긁어온다.
# 규칙:
# - 유튜브: 설명란만
# - 나머지: 게시판일수있다고보고 crawl_board_deep 먼저시도
#           실패/짧으면그냥 crawl_single_page
    texts = []
for url in urls:
print(f"[INFO] 수집중: {url}")
# 유튜브
if"youtube.com/watch"in url or"youtu.be/"in url:
            t = crawl_single_page(url)
if t:
                texts.append(t)
continue
# 커뮤니티게시판처럼보고 deep 시도
        deep_text = crawl_board_deep(url)
```

```
# deep 결과가어느정도길이가나오면그걸사용
if deep_text andlen(deep_text.split()) > 20:
continue
```

```
# deep이별로면단일페이지만긁는다
        single_text = crawl_single_page(url)
if single_text:
            texts.append(single_text)
# 너무빠르게연속요청하지않도록대기
0.5)
```

```
    joined = "\n".join(texts)
return joined
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

42/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
defsave_top20(result_path, top_list):
# top_list = [(단어, 빈도), ...] 형태를
# output_top20.txt 에저장한다.
    lines = []
    rank = 1
for word, freq in top_list:
        line = f"{rank}. 단어: {word} / 빈도: {freq}"
        lines.append(line)
        rank += 1
```

```
    out_txt = "\n".join(lines)
withopen(result_path, "w", encoding="utf-8") as f:
        f.write(out_txt)
```

```
print(f"[INFO] 결과가저장되었습니다 -> {result_path}")
```

```
defmain():
# 현재 crawler.py 가있는폴더위치
    base_dir = os.path.dirname(os.path.abspath(__file__))
```

```
    input_file = os.path.join(base_dir, "input_urls.txt")
    output_file = os.path.join(base_dir, "output_top20.txt")
```

```
    urls = read_input_urls(input_file)
ifnot urls:
```

```
print("[ERROR] input_urls.txt 에주소가없습니다.")
return
```

```
print("[INFO] URL 목록불러옴. 크롤링시작합니다...")
# 1) 전체텍스트긁어오기
```

```
    raw_all_text = crawl_all(urls)
```

```
ifnot raw_all_text:
```

```
print("[ERROR] 수집된텍스트가없습니다.")
return
```

```
# 2) 텍스트정리
    cleaned = clean_text(raw_all_text)
```

```
# 3) 상위 20 단어뽑기
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

43/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
    top20 = count_top_words(cleaned, top_n=20)
```

```
# 4) 저장
    save_top20(output_file, top20)
```

```
print("[DONE] 완료되었습니다.")
```

```
if __name__ == "__main__":
    main()
```

──────────────── 

다음 사용 순서 (형이 지금 바로 해볼 수 있는 단계) [사실] 

**1.** 방금 준 코드로 `crawler.py` 다시 저장해. (기존 내용은 다 지우고 완전히 바꿔치기) 

**2.** 같은 폴더에 `input_urls.txt` 가 꼭 있어야 돼. 예: 

한 줄에 커뮤니티 게시판 주소 

다음 줄에 블로그 글 주소 