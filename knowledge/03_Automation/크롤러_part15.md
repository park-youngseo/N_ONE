continue
# deep이별로면단일페이지만긁는다
        single_text = crawl_single_page(url)
if single_text:
            texts.append(single_text)
# 너무빠르게연속요청하지않도록잠깐쉰다.
        time.sleep(0.5)
    joined = "\n".join(texts)
return joined
```

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
    input_file = os.path.join(base_dir, "input_urls.txt")
    output_file = os.path.join(base_dir, "output_top20.txt")
```

```
    urls = read_input_urls(input_file)
ifnot urls:
```

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

67/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
print("[ERROR] input_urls.txt 에주소가없습니다.")
return
```

```
print("[INFO] URL 목록불러옴. 크롤링시작합니다...")
```

```
# 1) 전체텍스트긁어오기
    raw_all_text = crawl_all(urls)
```

```
ifnot raw_all_text:
print("[ERROR] 수집된텍스트가없습니다.")
return
```

```
# 2) 텍스트정리
    cleaned = clean_text(raw_all_text)
```

```
# 3) 상위 20 단어뽑기
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

형 다음 순서 다시 정리해줄게. [사실] 

**1.** `crawler.py` 에 지금 코드 그대로 붙여넣고 저장. 

**2.** 같은 폴더에 `input_urls.txt` 넣어두고 (주소 줄마다 하나씩). 

**3.** cmd 열고 폴더로 이동해서 

cmd 코드 복사 

```
python crawler.py
```

**4.** 검은 창에 `[INFO] 게시판 페이지 진입:` / `글 1/50 수집 중:` 이런 식으로 진행 상황 나 오는지 확인. 

**5.** 끝나면 `output_top20.txt` 열어서 TOP 20 단어 보면 된다. 

- 형, 이렇게 맞춰 둘게. 

**==> picture [489 x 38] intentionally omitted <==**

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

68/85 

25. 11. 30. 오후 12:58 

크롤러 

 

[INFO] 수집 중: https://www.filmmakers.co.kr/board [INFO] 게시판 페이지 진입: https://www.filmmakers.co.kr/board [WARN] 게시판 페이지 로드 실패 

형, 지금 상황은 코드 구조 문제라기보다 사이트가 막은 거라고 보면 된다. [추론] 

──────────────── 

A단계 (전체 개요) 

메시지 흐름을 보면 

```
[INFO] 게시판페이지진입: https://www.filmmakers.co.kr/board
```

```
[WARN] 게시판페이지로드실패
```

이 순서로 나왔지? [사실] 

## 이건 뭐냐면: 

**1.** 내가 `fetch_url()` 로 해당 주소에 접속을 시도했다. 

**2.** HTTP 응답 코드(사이트에서 준 답)가 200이 아니었다. → 즉 “정상 페이지를 안 돌려줌”. [사실] 

**3.** 그래서 `main_html` 이 비어서 WARN을 찍은 거야. [사실] 

이건 우리 코드가 잘못된 문법 때문에 터진 게 아니고, 그 사이트가 “이건 봇(자동스크립 트) 같네? 안 줄래” 하고 빈 응답(또는 403 같은 차단 응답)을 준 거다. [추론] 

filmmakers.co.kr 같은 곳은 업계 공고/게시판이라서, 자동 수집(크롤링)에 민감할 수 있다. 그래서 기본 헤더만 보내면 차단하는 경우가 많다. [추론] 

즉: 사이트가 우리를 내쫓은 거라서 텍스트를 못 가져온 거다. [사실] 

──────────────── 

B1단계 (핵심) 

지금 당장 형이 할 수 있는 해법은 2갈래야. [사실] 

**1.** 해결책 A (빠르게 확인용): 

   - 막지 않는 다른 공개 커뮤니티 / 블로그 / 유튜브 주소를 `input_urls.txt` 에 넣고 돌 려서 

프로그램 자체가 정상 동작하는지 먼저 확인한다. 

→ 이걸로 “내 파이썬, 내 폴더 구조, 실행법은 문제 없음”을 검증할 수 있어. [사실] 

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

69/85 

25. 11. 30. 오후 12:58 

크롤러 

 

**2.** 해결책 B (조금 더 공격적으로 filmmakers 같은 곳도 시도): 

우리가 보내는 요청을 진짜 사람 브라우저처럼 더 위장하도록 `fetch_url()` 과 

## `HEADERS` 를 강화하고, 

요청이 성공했는지 상태코드(예: 200인지 403인지)를 콘솔에 찍어서 확인하게 만든 다. [사실] 

지금부터는 해결책 B를 위한 REPAIR 패치만 줄게. 형은 그 부분만 복사해서 바꿔 넣으면 된다. [사실] 

──────────────── 

B2단계 (세부 REPAIR: fetch_url() 개선 + 헤더 강화 + 상태코드 출력) 

**1.** 먼저 `HEADERS` 를 더 사람처럼 보이게 바꾼다. 

코드 맨 위쪽에 있는 `HEADERS = { ... }` 부분을 아래로 교체해. 

python 코드 복사 `HEADERS = { "User-Agent": ( "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " "AppleWebKit/537.36 (KHTML, like Gecko) " "Chrome/120.0.0.0 Safari/537.36" ), "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", "Accept-Language": "ko-KR,ko;q=0.9,en-US;q=0.6,en;q=0.4",` 

```
"Connection": "keep-alive",
}
```

  

의도: 

일반 브라우저처럼 “나 웹페이지 볼 거야(html 좀 줘)”라고 말하게 해. 한국어 선호( `ko-KR` )도 넣었다. 많은 국내 사이트는 이걸 보고 ‘사람일 수도 있네’라고 조금 더 느슨해진다. [추론] 

**2.** `fetch_url()` 함수를 아래 버전으로 교체해. 

- (기존 `def fetch_url(url):` 부터 `return ""` 까지 싹 지우고 아래 내용으로 붙여넣어.) 

python 코드 복사 

- `def fetch_url(url):` _`# URL`_ `에서` _`HTML`_ `을 가져온다` _`. #`_ `이번 버전은 상태코드` _`(`_ `몇 번인지` _`)`_ `와 차단 여부를 콘솔에 찍어서` 

- _`#`_ `어디서 막히는지 형이 눈으로 확인할 수 있게 한다` _`.`_ 

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

70/85 

25. 11. 30. 오후 12:58 

크롤러 

 

```
try:
# 요청보낼때 Referer(이전페이지) 흉내도준다.
# 일부사이트는 Referer 없으면차단한다.
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
# 403, 404 등이라면빈문자열반환
return""
except Exception as e:
print(f"[ERROR] 요청실패: {url} ({e})")
return""
```

## 의도: 

   - 이제 크롤링 중에 콘솔에 `[DEBUG] 요청 https://... -> 상태코드 403` 이런 식으로 찍 힐 거야. 

   - → 이걸 보면 “아, 사이트가 우리를 막았네(403 금지)”인지, 

   - “주소가 잘못돼서 못 찾네(404)”인지, 

   - “200 정상인데도 무언가 다른 문제네”인지 바로 알 수 있다. [사실] 

   - Referer 넣은 이유는 “나 이 사이트 안에서 페이지 옮겨다니는 중이야” 라고 보이게 하는 단순 위장이다. 일부 게시판은 Referer 없으면 바로 차단한다. [추론] 

**3.** 실행 다시 하는 법은 똑같아. 

      - 수정 저장 

cmd에서 `python crawler.py` 

      - 이번엔 검은 창에 `[DEBUG] 요청 ... -> 상태코드 ...` 줄이 추가로 뜬다. 

      - 만약 filmmakers에서 전부 403 같은 코드가 뜨면 = 그 사이트는 기본 봇 접근을 막고 있어서 이 단순한 방식으로는 못 긁는 거다. [사실] 

- 그 경우엔 우리가 할 수 있는 선택지가 두 가지야. [추론] 

   - 선택지 1: filmmakers처럼 막는 사이트는 제외하고, 다른 공개 커뮤니티(로그인 없이 읽히는 자유게시판 등) / 뉴스형 커뮤니티 / 블로그 / 유튜브 설명란 위주로 먼저 TOP20 흐름을 만든다. → 빠르고 안정적. 

https://chatgpt.com/c/6903817a-fecc-8324-8682-88ef9509697b 

71/85 

25. 11. 30. 오후 12:58 

 

크롤러 

선택지 2: filmmakers만을 꼭 크롤링하고 싶으면, 거기 전용 규칙(쿠키나 로그인 흐름 등)을 짜야 하고, 그건 사실상 “사이트 맞춤형 전용 크롤러” 단계로 간다. (여긴 민감 하니까 강하게 막을 확률 높음) 