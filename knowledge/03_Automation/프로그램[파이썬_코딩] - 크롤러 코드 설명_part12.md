## **3. 진짜 “0개” 실행이 따로 있었을 가능성** 

지금 내가 분석한 `crawler.log` 는: 

## **마지막 실행 1번의 결과만 있는 게 아니라** 

형이 여러 번 돌린 것들이 계속 이어 붙어 있는 파일이야. 

그 안에 어딘가에는 분명: 

text 코드 복사 

```
수집완료: 0개
```

```
=== 저장완료: output\필름메이커스.csv ===
```

## 같은 구간도 있었을 거고, 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

57/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

## 형은 **그 실행을 기준으로 “0개다”라고 기억** 하고 있을 수 있어. 

이번에 형이 올린 최신 파일의 맨 끝 부분만 보면, 

**마지막 실행은 10개 수집 / 저장 완료** 가 맞다. 

## **형이 다음에 해보면 좋은 체크 순서** 

지금 상태에서, 딱 네 가지만 해보면 돼: 

## **1. 정확한 파일 열기** 

- `C:\Users\psm23\Desktop\n1_crawler\output\필름메이커스.csv` 

이 파일을 메모장으로 열어보고 

줄 수가 얼마인지 직접 확인. 

## **2. fetcher가 아직 “전체 사이트”를 도는 구조인지 확인** 

- 로그에 `/notice` , `/event` , `/about` , `/adinfo` 같은 주소가 계속 뜨면 

→ 배우포럼 전용 필터 버전으로 아직 안 바뀐 거야. 

- 그럴 땐 내가 준 `is_actor_post_link()` + `/actorsForum/` 필터 버전으로 `fetcher.py` 를 한 번에 갈아끼워야 함. 

## **3. patterns.py 광고 필터 확인** 

- `AD_HINTS` 안에서 `"문의", "카톡", "카카오"` 세 개는 

진짜로 지워졌는지 한 번 더 확인. 

- 남은 건 일단 `"체험단", "DM주세요", "홍보"` 정도만 둬 보자. 

## **4. 새로 돌릴 때는 사이트 이름도 다르게** 

헷갈리지 않게, 다음 실행은 

사이트 이름: `필름메이커스_test2` 

- 그럼 `output\필름메이커스_test2.csv` 가 별도로 생겨서 

이전 실행 결과랑 구분하기 쉽다. 

## **한 줄로 정리하면** 

## **지금 로그에 적힌 실행은 “0개”가 아니라 “10개 수집 후 저장 완료”다.** 

- 형이 “0개 같다”고 느끼는 건 

   - (1) 다른 파일을 열고 있거나 

   - (2) 우리가 원하는 종류의 글이 아니라서 체감상 0개처럼 느껴지거나 

   - (3) 예전 0개 실행을 기준으로 기억하고 있어서야. 

이제: 

**1.** `output\필름메이커스.csv` 메모장으로 열어서 줄 수 한 번 보고 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

58/114 

25. 11. 30. 오후 1:06 

 

**==> picture [137 x 9] intentionally omitted <==**

**2.** 필요하면 `fetcher.py` 를 아예 **배우포럼 전용 필터 버전** 으로 다시 한 번 정확히 붙여 넣고 

**3.** 사이트 이름 바꿔서 새 CSV 생성 

이렇게만 해보면, 

진짜로 “0개인지 / 10개인지 / 원하는 글이 맞는지”가 완전 분명해질 거야. 

**==> picture [489 x 39] intentionally omitted <==**

페이지 오류: Invalid URL '#': No scheme supplied. Perhaps you meant 이게 계속 뜨는데 이건 왜 아무말안하 지? 

56s 동안 생각함 

아 그 에러 얘기구나, 이제 딱 이해했어. 

```
Invalid URL '#': No scheme supplied. Perhaps you meant https://#?
```

이 메시지는 **크롤러가** `href="#"` **같은 링크를 진짜 웹주소로 착각해서 요청을 보내려다 가, requests가 “이건 URL이 아니야”라고 욕하는 상황** 이야. 

## **1. 왜 이게 계속 뜨는지** 

지금 `fetcher.py` (형이 실제로 돌리고 있는 버전)는 대충 이런 구조일 거야: 

python 코드 복사 

```
items = soup.select("a")  # 페이지안의모든 <a> 태그
```

```
for a in items:
    title = a.text.strip()
    href = a.get("href", "")
ifnot href:
continue
# 여기서바로 requests.get(href, ...) 하는구조
```

근데 사이트 안에는 이런 링크들이 많아: 

html 코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

59/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
<ahref="#">맨위로</a>
```

```
<ahref="#">댓글로이동</a>
```

즉, 

- `href="#"` → **그 페이지 안에서만 위치 이동하는 버튼** 이지, 

`http://…` 이런 “진짜 주소”가 아니라서 

여기에다 대고 `requests.get("#")` 를 날리니까 파이썬이: 

- “이건 스킴(http/https)이 없는데? URL 아님” 

→ `Invalid URL '#': No scheme supplied…` 

라고 에러를 뱉는 거야. 

## 다행인 건: 

- 이 에러는 우리가 `try/except` 로 잡아서 

## **로그에만 남기고 그냥 넘어가고 있음** 

(그래서 프로그램이 죽지는 않음) 

하지만 로그에 계속 찍히니까 형 눈에는 계속 거슬리는 거고. 

형이 “왜 아무 말 안하냐?” 느낀 건, 

내가 앞에서 **동작 전체(0개/10개, 필터링 문제)**에 집중해서 설명하면서 

이 개별 메시지까지 따로 짚어주질 못해서 그래. 여기선 분명히 짚고 갈게. 

## **2. 해결 방법 –** `href="#"` **같은 가짜 링크는 아예 건너뛰자** 

지금 쓰고 있는 `fetcher.py` 의 이 부분을 고쳐주면 돼. 

## 🔧 **2-1.** `fetcher.py` **에서 이 부분 찾기** 

대략 이런 구간이 있을 거야(형 코드랑 완전 똑같진 않아도 비슷할 거야): 

python 코드 복사 

```
for a in items:
    title = a.text.strip()
    href = a.get("href", "")
ifnot href:
continue
if href.startswith("/"):
        href = url.split("/")[0] + "//" + url.split("/")[2] + href
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

60/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

```
# 상세페이지
try:
        d = requests.get(href, headers=headers, timeout=10)
        ...
```

## 여기서 **href를 검사하는 부분을 조금만 늘려주면 끝** 이야. 

## 🔁 **2-2. 이걸로 교체해 줘** 

python 

코드 복사 

```
for a in items:
    title = a.text.strip()
    href = a.get("href", "")
# 1) href가비어있거나, '#' 이거나, 자바스크립트링크면버린다
ifnot href:
continue
    href = href.strip()
if href == "#"or href.startswith("javascript:"):
continue
# 2) 상대경로(/로시작)면도메인붙여서절대주소로만든다
if href.startswith("/"):
        href = url.split("/")[0] + "//" + url.split("/")[2] + href
# 상세페이지
try:
        d = requests.get(href, headers=headers, timeout=10)
        ...
```

## 포인트는 딱 이 줄들 추가야: 

python 코드 복사 

```
if href == "#"or href.startswith("javascript:"):
continue
```

## 이렇게 해두면, 

```
href="#"
href="javascript:void(0)"
```

## 같은 **“버튼용 링크”는 전부 스킵** 되니까 

```
Invalid URL '#': No scheme supplied…
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69284f80-fd58-8324-8d8f-787b4682c444 

61/114 

25. 11. 30. 오후 1:06 

프로그램[파이썬/코딩] - 크롤러 코드 설명 

 

이 메시지는 더 이상 안 뜰 거야. 

## **3. 이 에러가 크롤링 결과를 망치고 있었냐?** 

## **아니야.** 

- 이 에러들은 `try/except` 안에서 로그만 찍고 넘어가는 거라, 

크롤러 전체가 멈추거나 CSV 저장이 안 되는 원인은 아니었어. 