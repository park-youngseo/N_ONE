try:
        df = pd.read_csv(filepath)
print(f"[+] Loaded {len(df)} posts from {filepath}")
return df
except FileNotFoundError:
print(f"[!] File not found: {filepath}")
returnNone
```

```
defsplit_sentences(text):
"""
한글(text)을문장단위로대략분리합니다.
    - 마침표/물음표/느낌표기준으로단순분리
    - 내용은바꾸지않고, 나누기만합니다.
    """
ifnotisinstance(text, str) ornot text.strip():
return []
# 줄바꿈을공백으로바꾸기
    text = text.replace("\n", " ")
# ., ?, ! 기준으로문장분리 (아주단순한버전)
    sentences = re.split(r"[\.?!]", text)
# 공백제거후빈문자열버리기
return [s.strip() for s in sentences if s.strip()]
```

```
defsummarize_post(text, max_sentences=2):
"""
글의앞부분문장 1~2개를가져와서
    '이글이주로무슨얘기인지'를간단히보여주는요약을만듭니다.
    - 해석/추측없이, 앞에서부터자르는방식입니다.
    """
    sentences = split_sentences(text)
ifnot sentences:
return""
# 앞에서부터 max_sentences개만이어붙이기
return" ".join(sentences[:max_sentences])
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

92/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

```
# 아주간단한한국어불용어(버릴단어) 리스트
KOREAN_STOPWORDS = {
"그리고", "그러나", "하지만", "그래서",
"입니다", "있습니다", "합니다",
"이번", "저는", "제가", "우리",
"정도", "조금", "또는", "또", "있는",
}
defextract_keywords(text, top_n=10):
"""
글에서자주나온단어를뽑아냅니다.
    - 조사/기능어/너무짧은단어는대략적으로제거합니다.
    - '심리해석'이아니라, 그냥 '많이나온단어리스트'입니다.
    """
ifnotisinstance(text, str):
return []
# 한글/영문/숫자만남기고나머지기호는공백으로
    cleaned = re.sub(r"[^0-9A-Za-z가-힣\s]", " ", text)
    tokens = cleaned.split()
# 너무짧은단어, 불용어제외
    words = [
        tok for tok in tokens
iflen(tok) >= 2and tok notin KOREAN_STOPWORDS
    ]
ifnot words:
return []
```

```
    counter = Counter(words)
    common = [w for (w, _) in counter.most_common(top_n)]
return common
```

```
defanalyze_all_posts(input_path="posts.csv", output_path="posts_analyzed.csv"):
"""
    posts.csv를읽어서:
    - summary_auto: 자동요약한줄
    - keywords_auto: 자주등장한단어리스트(문자열)
를추가한뒤, 새파일(posts_analyzed.csv)로저장합니다.
    """
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

93/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

```
    df = load_posts(input_path)
if df isNone:
return
# 요약생성
""
    df["summary_auto"] = df.get("text_raw", ).apply(summarize_post)
# 키워드생성 (CSV에넣기위해쉼표로이어줌)
""
    df["keywords_auto"] = df.get("text_raw", ).apply(
lambda t: ", ".join(extract_keywords(t))
    )
# 새파일로저장 (엑셀에서한글안깨지게 utf-8-sig 사용)
    df.to_csv(output_path, index=False, encoding="utf-8-sig")
print(f"[+] Saved analyzed posts to {output_path}")
```

```
defaggregate_keyword_stats(
    input_path="posts_analyzed.csv",
    output_path="keyword_stats.csv",
):
"""
    posts_analyzed.csv 안의 keywords_auto 컬럼을모아서
키워드별등장횟수를집계한뒤 keyword_stats.csv 로저장합니다.
```

```
    - 아직동의어묶기, 카테고리나누기같은해석은하지않습니다.
    - '같은글자'인단어끼리만묶어서몇번나왔는지카운팅합니다.
    """
try:
        df = pd.read_csv(input_path)
except FileNotFoundError:
print(f"[!] File not found: {input_path}")
return
    counter = Counter()
# keywords_auto 컬럼에서키워드모으기
for value in df.get("keywords_auto", []):
ifnotisinstance(value, str):
continue
# "학원, 수업, 고민" -> ["학원", "수업", "고민"]
        parts = [p.strip() for p in value.split(",")]
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

94/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

```
# 빈문자열제거하고카운팅
for kw in parts:
if kw:
                counter[kw] += 1
ifnot counter:
print("[!] No keywords found to aggregate.")
return
```

```
# Counter -> DataFrame (많이나온순서대로정렬)
    items = sorted(counter.items(), key=lambda x: x[1], reverse=True)
    stats_df = pd.DataFrame(items, columns=["keyword", "count"])
```

```
# 저장 (엑셀에서깨지지않도록 utf-8-sig 사용)
    stats_df.to_csv(output_path, index=False, encoding="utf-8-sig")
print(f"[+] Saved keyword stats to {output_path}")
```

```
if __name__ == "__main__":
# 1) 글단위요약/키워드붙이기
    analyze_all_posts()
# 2) 전체키워드분포집계
    aggregate_keyword_stats()
```

## 🔁 **실행 순서 한 번만 다시 정리** 

## **1. 크롤러 돌려서 최신 posts.csv 만들기** 

powershell 

코드 복사 

```
python crawler.py
```

## **2.** 그다음 **analyzer 실행** 

powershell 

코드 복사 

```
python analyzer.py
```

## **3.** 실행 끝나면 폴더 안에: 

`posts_analyzed.csv` → 각 글에 summary_auto, keywords_auto 붙은 파일 

`keyword_stats.csv` → keyword, count 정리된 파일 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

95/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

두 개 생기면 로직까지 전부 제대로 적용된 거야. 

에러 뜨면 **에러 화면 캡처** 해서 보여주면, 그 지점만 다시 같이 고치자. 

**==> picture [489 x 38] intentionally omitted <==**

**==> picture [289 x 86] intentionally omitted <==**

## 1m 34s 동안 생각함 

- 응, 이건 **로직/코드 에러가 아니라 “사이트가 대답을 안 해서 시간 초과 난 것”**이야. 

## 로그를 보면: 

text 코드 복사 

```
[ERROR] Failed to fetch https://www.filmmakers.co.kr/board?page=2
(Caused by ConnectTimeoutError(... connect timeout=10))
```

## 뜻은 간단해: 

- 우리 코드: “page=2 주세요” 하고 요청 보냄 

- 필름메이커스 서버: 10초 안에 응답을 안 줌 

- `requests` 가 “10초 기다렸는데도 응답이 없네?” → **ConnectTimeoutError** 찍고 포 

- 기 

## 그래서: 

## **crawler.py / analyzer.py 로직이 틀려서가 아니고** , 

- “2,3,4페이지 서버 응답이 너무 느리거나, 잠깐 막혀서” 생긴 현상이라고 보면 돼. 

- (1페이지는 이미 잘 받아와서 `posts.csv` 가 만들어졌잖아.) 

## **형이 해볼 수 있는 조정 3가지** 

## **1) 타임아웃 늘리기 (10초 → 30초 정도)** 

**1.** `crawler.py` 열고 `Ctrl + F` 로 `timeout=` 검색해봐. 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

96/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

## **2.** 아마 이런 식으로 되어 있을 거야: 

python 코드 복사 

```
response = requests.get(url, headers=headers, timeout=10)
```

## **3.** 이걸 이렇게 바꿔: 

python 코드 복사 `response = requests.get(url, headers=headers, timeout=30)` 

→ 서버가 느려도 **최대 30초까지는 기다려준다** 는 뜻. 

## **2) 요청 간 딜레이 늘리기 (너무 빠르게 안 두드리게)** 

`config.py` 에 이런 줄 있지? 
