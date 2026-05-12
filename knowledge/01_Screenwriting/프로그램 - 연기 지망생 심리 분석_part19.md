
https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

86/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

38s 동안 생각함 

맞아, 지금까지는 

글 단위로 요약 / keywords_auto 를 “붙여놓기”만 한 상태고 

이제 **그 keywords_auto들을 전부 모아서 “몇 번 나왔냐”를 세야** 객관적인 지표가 되 지. 

## 여기서는 **아직 어떤 해석/카테고리도 붙이지 않고** , 

그냥 **“같은 단어끼리 묶어서 몇 번 나왔는지”**만 보는 걸로 할게. 

## **1. 로직부터 먼저 정리** 

**대상 파일:** `posts_analyzed.csv` 

**목표:** `keyword, count` 형태의 표( `keyword_stats.csv` ) 만들기 

로직은 진짜 단순하게 이렇게야: 

**1.** `posts_analyzed.csv` 의 각 행에서 `keywords_auto` 값을 가져온다. 

   - 예: `"학원, 수업, 고민, 배우, 시작"` 

**2.** 이 문자열을 `,` (쉼표) 기준으로 잘라서 리스트로 만든다. 

   - → `["학원", "수업", "고민", "배우", "시작"]` 

**3.** 각 단어에서 양쪽 공백을 제거하고 빈 문자열은 버린다. 

**4.** 전부 하나의 큰 카운터에 **단어별로 +1씩** 더한다. 

   - 예: 

- 학원: 12번 수업: 9번 

고민: 15번 

   - … 

**5.** 마지막에 `keyword, count` 형태로 정렬해서 `keyword_stats.csv` 로 저장한다. 

→ 이 파일이 “객관적인 키워드 분포표”가 되는 거야. 

## 여기까지는 **동의어 묶기, 카테고리 만들기 없음** . 

그냥 “글들에서 이 단어가 몇 번 나왔냐”만 보는 100% 숫자 단계. 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

87/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

## **2. analyzer.py 에 코드 추가하기** 

`analyzer.py` 맨 아래 쪽에 **이 함수 하나만 더 추가** 하자. 

python 코드 복사 

```
from collections import Counter
import pandas as pd
defaggregate_keyword_stats(
    input_path: str = "posts_analyzed.csv",
    output_path: str = "keyword_stats.csv",
):
"""
    posts_analyzed.csv 안의 keywords_auto 컬럼을모아서
키워드별등장횟수를집계한뒤 keyword_stats.csv 로저장합니다.
    - 아직동의어묶기, 카테고리나누기같은해석은하지않습니다.
    - '같은글자'인단어끼리만묶어서몇번나왔는지카운팅합니다.
    """
try:
        df = pd.read_csv(input_path)
except FileNotFoundError:
print(f"[!] File not found: {input_path}")
return
    counter = Counter()
for value in df.get("keywords_auto", []):
ifnotisinstance(value, str):
continue
# "학원, 수업, 고민" -> ["학원", "수업", "고민"]
        parts = [p.strip() for p in value.split(",")]
# 빈문자열제거하고카운팅
for kw in parts:
if kw:
                counter[kw] += 1
ifnot counter:
print("[!] No keywords found to aggregate.")
return
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

88/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

```
# Counter -> DataFrame
```

```
    items = sorted(counter.items(), key=lambda x: x[1], reverse=True)
    stats_df = pd.DataFrame(items, columns=["keyword", "count"])
```

```
# 저장 (엑셀에서깨지지않도록 utf-8-sig 사용)
```

```
    stats_df.to_csv(output_path, index=False, encoding="utf-8-sig")
print(f"[+] Saved keyword stats to {output_path}")
```

그리고 파일 맨 아래 `if __name__ == "__main__":` 부분을 

이렇게 바꿔두면 편해: 

python 코드 복사 

```
if __name__ == "__main__":
# 1) 글단위요약/키워드붙이기
    analyze_all_posts()
# 2) 전체키워드분포집계
    aggregate_keyword_stats()
```

이렇게 해두면 `python analyzer.py` 한 번만 실행해도 두 작업이 순서대로 같이 돌아. 

## **3. 실행 순서** 

## **1. 먼저 크롤러 돌려서** `posts.csv` 갱신 

powershell 코드 복사 

```
python crawler.py
```

## **2.** 그 다음 **분석 + 키워드 집계** 

powershell 코드 복사 

```
python analyzer.py
```

**3.** 실행이 끝나면 폴더 안에: 

`posts_analyzed.csv` → 각 글의 요약 + keywords_auto `keyword_stats.csv` → `keyword, count` 형태의 전체 분포표 

두 개가 생겨 있을 거야. 

`keyword_stats.csv` 를 엑셀로 열면 

## A열: 키워드 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

89/97 

25. 11. 30. 오후 1:44 프로그램 - 연기 지망생 심리 분석 

 

B열: 그 키워드가 **몇 개 글에서/몇 번** 등장했는지 

이렇게 숫자로만 정리된 **객관 지표** 가 되는 거고, 

이건 아직 아무 해석 안 붙인 상태라 나중에 3단계에서 형이 직접 “어떤 무리가 무슨 의미냐”를 볼 때 기초 데이터로 쓰면 된다. 

**==> picture [489 x 38] intentionally omitted <==**

if __name__ == "__main__": 이걸 어떻게 찾냐? 

형, 복잡하게 생각 안 해도 돼. 

`if __name__ == "__main__":` 는 **파일 안에서 한 줄짜리 문장** 일 뿐이야. 

## **1. VSCode에서 찾는 방법** 

**1.** 왼쪽에서 `analyzer.py` 탭 눌러서 연다. 

**2.** 키보드에서 `Ctrl + F` 누르기. 

**3.** 검색창에 그대로 이렇게 입력: 

text 코드 복사 

```
if __name__ == "__main__":
```

**4.** Enter 치면 

   - 파일 안에 이 줄이 있으면, 거기로 바로 커서가 이동한다. 없으면 “0/0” 이런 식으로 아무 것도 안 나와. 

## **2. 만약 검색해도 안 나오면?** 

## 그럼 **형이 직접 맨 아래에 만들어주면 되는 줄** 이야. 

**1.** `analyzer.py` 제일 아래로 스크롤 내려가. 

**2.** 마지막 줄 아래에 엔터 한 번 치고, 이걸 그대로 붙여 넣어: 

python 코드 복사 `if __name__ == "__main__": analyze_all_posts() aggregate_keyword_stats()` 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

90/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

**3.** 저장(Ctrl + S). 

이렇게만 해두면, 

powershell 

코드 복사 

```
python analyzer.py
```

라고 실행했을 때 

```
analyze_all_posts()
aggregate_keyword_stats()
```

- 이 두 함수가 **자동으로 순서대로 실행되는 구조가 되는 거야. 

**==> picture [489 x 38] intentionally omitted <==**

오류나지 않게 적용된 모든걸 전부 다시 뽑아 

20s 동안 생각함 

좋아, 형. 

지금까지 말한 기능(요약 + keywords_auto + keyword_stats) 전부 한 파일에 깔끔하게 합 친 **최종 analyzer.py 전체 코드** 뽑아줄게. 

## 👉 **방법:** 

VSCode에서 `analyzer.py` 열고, 

**안에 있는 내용 전부 삭제 → 아래 코드 그대로 복붙 → 저장(Ctrl+S)** 하면 된다. 

## ✅ **최종 analyzer.py (오류 안 나게 전체 통합 버전)** 

python 코드 복사 `import pandas as pd import re from collections import Counter` 

```
defload_posts(filepath="posts.csv"):
"""
    posts.csv 파일을읽어서 DataFrame으로반환합니다.
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

91/97 

25. 11. 30. 오후 1:44 

 

프로그램 - 연기 지망생 심리 분석 

```
파일이없으면 None을반환합니다.
    """