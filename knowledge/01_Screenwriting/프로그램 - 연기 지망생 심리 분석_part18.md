
## 맞아 형, **지금 나온 엑셀은 말 그대로 “그냥 긁어온 것”이야.** 

1단계에서 한 건 딱 이 정도뿐이야: 

정해둔 게시판에 들어가서 

글/메뉴 링크를 **그대로** 가져오고 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

81/97 

25. 11. 30. 오후 1:44 프로그램 - 연기 지망생 심리 분석 

 

제목, 본문, URL, 날짜 같은 **메타데이터만 붙여서 표로 만든 것** 

즉, 

**심리 해석 X** 

**카테고리 분류 X 키워드 분석 X 패턴 묶기 X** 

아직은 진짜로 **“원문 텍스트 창고를 만든 것”** 단계라고 보면 돼. 

이 창고를 바탕으로 다음에 2단계(요약/키워드/클러스터)부터가 

- “내용을 다루는 분석 단계”가 되는 거고. 

**==> picture [489 x 38] intentionally omitted <==**

다음단계로 진행해봐 

## 1m 12s 동안 생각함 

그래, 이제 **2단계(분석)**로 한 번만 딱 넘어가 보자. 

지금 상태 = “그냥 긁어온 원문 창고”까지 만든 거고, 

이제 할 일 = **각 글에 “요약 + 자주 나온 단어”만 붙여주는 단계** 야. 

(심리/타입 이름 붙이기, 패턴 해석은 아직 X) 

## **1단계 ┃** `analyzer.py` **코드 먼저 채우기** 

## **1-1. 해야 할 일만 딱 두 가지** 

**1.** `text_raw` 를 이용해서 **글 요약 한 줄** 만들기 

**2. 자주 나온 단어 몇 개** 뽑아서 붙이기 

이 두 개 결과를 `posts_analyzed.csv` 같은 새 파일로 저장할 거야. 

## **1-2.** `analyzer.py` **에 넣을 코드 예시** 

VSCode에서 `analyzer.py` **파일을 열고** , 

지금 있는 함수들(load_posts, split_sentences, summarize_post…) 자리에 아래 코드로 갈아끼우면 된다. 

## ⚠ **중요** 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

82/97 

25. 11. 30. 오후 1:44 

 

프로그램 - 연기 지망생 심리 분석 “이 코드는 “요약/키워드”까지만 한다.” ““나이 불안형” 같은 심리 이름, 카테고리는 전혀 안 붙인다.” 

python 

코드 복사 

```
import pandas as pd
import re
from collections import Counter
```

```
# CSV 파일을읽어오는함수
defload_posts(filepath: str = "posts.csv") -> pd.DataFrame | None:
"""
    posts.csv 파일을읽어서 DataFrame으로반환합니다.
    """
try:
        df = pd.read_csv(filepath)
print(f"[+] Loaded {len(df)} posts from {filepath}")
return df
except FileNotFoundError:
print(f"[!] File not found: {filepath}")
returnNone
```

```
defsplit_sentences(text: str) -> list[str]:
"""
한글(text)을문장단위로대략분리합니다.
    - 마침표/물음표/느낌표기준으로만단순분리
    - 내용은바꾸지않고, 나누기만합니다.
    """
ifnotisinstance(text, str) ornot text.strip():
return []
# 줄바꿈을공백으로바꿔서한줄처럼만든다음
    text = text.replace("\n", " ")
```

```
# . ? ! 기준으로문장분리 (아주단순한버전)
    sentences = re.split(r"[\.?!]", text)
# 공백정리후빈문장은버리기
return [s.strip() for s in sentences if s.strip()]
```

```
defsummarize_post(text: str, max_sentences: int = 2) -> str:
"""
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

83/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

```
글의앞부분문장 1~2개를가져와서
```

```
    '이글이주로무슨얘기를하는지'를간단히보여주는요약을만듭니다.
    - 해석/추측없이, 앞에서부터자른다고생각하면됩니다.
    """
    sentences = split_sentences(text)
ifnot sentences:
return""
# 앞에서부터 max_sentences개만이어붙이기
return" ".join(sentences[:max_sentences])
```

```
# 아주간단한불용어(버릴단어) 리스트
KOREAN_STOPWORDS = {
"그리고", "그러나", "하지만", "그래서",
"입니다", "있습니다", "합니다",
"이번", "저는", "제가", "우리",
"정도", "조금", "또는", "또", "있는",
}
```

```
defextract_keywords(text: str, top_n: int = 10) -> list[str]:
"""
글에서자주나온단어를뽑아냅니다.
    - 조사/기능어/짧은단어는대충걸러냅니다.
    - '심리해석'이아니라, 그냥 '많이나온단어리스트'입니다.
    """
ifnotisinstance(text, str):
return []
# 문자만남기고나머지기호는공백으로
    cleaned = re.sub(r"[^0-9A-Za-z가-힣\s]", " ", text)
    tokens = cleaned.split()
# 너무짧은단어, 불용어제외
    words = [
        tok for tok in tokens
iflen(tok) >= 2and tok notin KOREAN_STOPWORDS
    ]
ifnot words:
return []
    counter = Counter(words)
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

84/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

```
    common = [w for (w, _) in counter.most_common(top_n)]
return common
```

```
defanalyze_all_posts(
    input_path: str = "posts.csv",
    output_path: str = "posts_analyzed.csv",
):
"""
    posts.csv를읽어서:
    - summary_auto: 자동요약한줄
    - keywords_auto: 자주등장한단어리스트(문자열)
를추가한뒤, 새파일로저장합니다.
    """
    df = load_posts(input_path)
if df isNone:
return
# 요약생성
    df["summary_auto"] = df["text_raw"].apply(summarize_post)
# 키워드생성 (CSV에넣기위해쉼표로이어줌)
    df["keywords_auto"] = df["text_raw"].apply(
lambda t: ", ".join(extract_keywords(t))
    )
# 새파일로저장 (엑셀에서한글안깨지게 utf-8-sig 사용)
    df.to_csv(output_path, index=False, encoding="utf-8-sig")
print(f"[+] Saved analyzed posts to {output_path}")
```

```
if __name__ == "__main__":
    analyze_all_posts()
```

## 이대로 전체 갈아끼우고 **Ctrl+S** 로 저장. 

## **2단계 ┃ 분석 코드 실행하기** 

## **1.** VSCode 아래 **Terminal** 확인 

(이미 `C:\Users\psm23\Desktop\None_Agent>` 라고 되어 있으면 그대로 사용) 

**2.** 터미널에 아래처럼 입력: 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

85/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

코드 복사 

powershell 

 

```
python analyzer.py
```

**3.** 잘 되면 터미널에: 

   - `[+] Loaded 00 posts from posts.csv` 

   - `[+] Saved analyzed posts to posts_analyzed.csv` 

비슷한 문구가 뜰 거야. 

## **3단계 ┃ 새로 생긴** `posts_analyzed.csv` **확인** 

**1.** 왼쪽 Explorer에 `posts_analyzed.csv` 파일이 생겼는지 확인. 

**2.** 더블 클릭해서 열어보면, 기존 컬럼들 뒤에: 

   - `summary_auto` 

   - `keywords_auto` 

두 칸이 새로 생겨 있을 거야. 

`summary_auto` → 그 글 맨 앞 1~2문장을 이어 붙인 요약 

`keywords_auto` → 그 글에서 자주 나온 단어를 쉼표로 이어 놓은 것 

- (예: `학원, 수업, 고민, 배우, 시작` 이런 식) 

여기까지가 **2단계: “있는 글이 무슨 얘기인지 구조만 붙여주기” 단계** 야. 

심리 이름, 타입, 카테고리 같은 건 여전히 안 붙인 상태고, 

그건 형이 말했던 **3단계(주관 패턴 적용)**에서 따로 다루면 된다. 

일단 여기까지 실행해보고: 

에러가 뜨거나 

   - `summary_auto` / `keywords_auto` 내용이 이상해 보이면 

- 그 부분 캡처해서 보여주면, 그 지점만 골라서 다시 다듬자. 

**==> picture [489 x 39] intentionally omitted <==**

지금은 모든게 단계별로 나열이 되는 형식으로 만들어졌 는데 

언급된 keywords_auto → 그 글에서 자주 나온 단어를 쉼 표로 이어 놓은 것 

(예: 학원, 수업, 고민, 배우, 시작 이런 식)을 카운팅해야해 어떤식으로 같은 분류에 넣어서 많이 언급 되었는지 