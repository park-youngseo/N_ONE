
실제 사이트의 인증서는 `www.filmmakers.co.kr` 쪽에 맞춰져 있어서 

“야, 이 인증서는 `filmmakers.kr` 용이 아닌데?” 하고 막는 상황이야. 

## 1️⃣ **해결 방법 – URL만 제대로 바꿔주면 된다** 

**1. 브라우저에서 필름메이커스 들어가서 실제 주소 확인** 

주소창에 지금 열려 있는 필름메이커스 메인 보면 이렇게 돼 있을 거야: 

text 코드 복사 `https://www.filmmakers.co.kr/` 배우/지망생 게시판 들어가서 **그 게시판 첫 페이지 주소 전체** 를 복사해 둬. VSCode에서 `config.py` **열기** 안에 이런 줄 있을 거야 (대충 이런 느낌): python 코드 복사 `TARGET_BOARDS = [ "https://filmmakers.kr/actorCommunity" ]` **도메인 부분을 전부 교체** 해 줘: python 코드 복사 `TARGET_BOARDS = [ "복사해온_실제_게시판_URL_여기에_붙여넣기"` _`#`_ `예시` _`) # "https://www.filmmakers.co.kr/actors"`_ `]` 

배우/지망생 게시판 들어가서 **그 게시판 첫 페이지 주소 전체** 를 복사해 둬. 

**2.** VSCode에서 `config.py` **열기** 

안에 이런 줄 있을 거야 (대충 이런 느낌): 

## **3.** 여기서 **도메인 부분을 전부 교체** 해 줘: 

핵심은: 

> `filmmakers.kr` ❌ `https://www.filmmakers.co.kr/....` ⭕ (브라우저에서 복사한 것 그대로) 

**4.** 저장(Ctrl+S) 하고, 터미널에서 다시: 

powershell 코드 복사 `python crawler.py` 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

76/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

## 2️⃣ **혹시 여러 군데에 잘못된 주소가 들어갔을 수도 있을 때** 

- `crawler.py` 안에서도 `https://filmmakers.kr` 가 박혀 있을 수 있어. 

- VSCode에서 `Ctrl + Shift + F` 누르고 

- `filmmakers.kr` **로 전체 검색** 해서 

전부 `www.filmmakers.co.kr` / 실제 주소로 바꿔주면 깔끔하다. 

이 오류는 코드 구조 문제라기보다 

- **“잘못된 도메인 + SSL 인증서 불일치”**라서 

도메인만 제대로 맞춰주면 해결되는 타입이야. 

일단 `config.py` URL 바꾸고 한 번 더 돌려보고, 

그래도 다른 에러 뜨면 그 화면 캡처해서 다시 보여줘. 

**==> picture [489 x 39] intentionally omitted <==**

**==> picture [193 x 201] intentionally omitted <==**

왜 이러는거냐 

## 형 이거 이유 딱 하나야. 

## **config.py 는 제대로 고쳤는데,** 

## **crawler.py 안에 아직도** `filmmakers.kr` **가 박혀 있어서** 

상세글/다음 페이지 갈 때 계속 `https://filmmakers.kr/...` 로 요청하는 중이야. 그래서 터미널 에러에도 전부 `https://filmmakers.kr/board/...` 가 찍히는 거고. 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

77/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

## **1. 그래서 왜 에러가 뜨냐?** 

에러 메시지 그대로 보면: 

certificate is not valid for 'filmmakers.kr' 

- 실제 사이트는 `https://www.filmmakers.co.kr` 가 진짜 도메인인데 

- 코드가 `https://filmmakers.kr` 로 접속하려고 하니까 

- “이 인증서는 그 도메인용이 아닌데?” 하고 SSL이 막는 거야. 

즉, 

- **config.py** → `www.filmmakers.co.kr` ✅ 

- **crawler.py 내부** → 아직 `filmmakers.kr` ❌ 

- 그래서 계속 저 에러. 

## **2. 지금 해야 할 것 (아주 구체적으로)** 

## **① 전체 프로젝트에서** `filmmakers.kr` **싹 찾기** 

**1.** VSCode에서 `Ctrl + Shift + F` 눌러. 

**2.** 검색창에 `filmmakers.kr` 입력. 

**3.** Enter 치면, 

`crawler.py` 안에 최소 몇 군데 나올 거야 

(BASE_URL 같은 거, 페이지 만들 때 쓰는 주소 등) 

## **② 전부 진짜 도메인으로 바꾸기** 

각 줄을 이렇게 고치면 돼: 

만약 이렇게 돼 있으면: 

python 코드 복사 `BASE_URL = "https://filmmakers.kr"` python 코드 복사 `BASE_URL = "https://www.filmmakers.co.kr"` 

이걸 

로 바꾸기. 

또는: 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

78/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

코드 복사 

python 

```
requests.get("https://filmmakers.kr" + path, ...)
```

이런 식이면 전부 

python 코드 복사 

```
requests.get("https://www.filmmakers.co.kr" + path, ...)
```

로 바꾸기. 

## **포인트:** 

검색 결과에 나오는 `filmmakers.kr` 는 **한 군데도 남기지 말고** 

전부 `www.filmmakers.co.kr` 로 바꿔줘. 

**3.** 전부 고쳤으면 `Ctrl + S` 해서 저장. 

## **3. 다시 실행** 

터미널에서 다시: 

powershell 코드 복사 

```
python crawler.py
```

이번에는 에러 URL이 

`https://www.filmmakers.co.kr/board/....` 이렇게 찍혀야 정상이고, 

더 이상 SSL hostname mismatch 에러는 안 떠야 한다. 

요약하면: 

왜 이러냐? 

→ **config.py 는 고쳤는데, crawler.py 안에 남아있는** 

`https://filmmakers.kr` **때문에 계속 잘못된 도메인으로 접속하는 중이라 그렇다.** 

검색 한 번으로 싹 바꾸고 다시 돌려보자. 에러 메시지나 새 로그 뜨면 캡처해서 또 보여줘. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

79/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

**==> picture [289 x 102] intentionally omitted <==**

이렇게 엑셀파일이 떴다 이걸 로직과 관련이 있는건가? 

27s 동안 생각함 

## 응, **이 엑셀 파일이 바로 우리가 짠 1단계 수집 로직 결과** 야. 

`posts.csv` 를 엑셀로 연 거라서 이렇게 보이는 거고, 컬럼 구성도 아까 설계한 그대로 나 왔어. 

## **1. 각 컬럼이 의미하는 것 (우리가 설계한 그대로야)** 

- `post_id` → 지금은 글 고유 ID 대신 **글 URL 자체** 가 들어가 있음 

- `source` → 어떤 사이트인지 (filmmakers) 

- `board_name` → 우리가 설정한 게시판 이름 ( `actors` ) 

- `url` → 그 글의 상세 페이지 주소 

- `created_at` → 작성 날짜 

- `title_raw` → 글 제목(원문) 

- `text_raw` → 글 본문(원문) – 초록 박스 안에 문장들 잘 들어가 있는 부분 

- `view_count` / `like_count` / `comment_count` → 아직 0으로만 들어와 있는데, 사이트 

- 에서 숫자를 못 읽어오거나 안 넣어둔 상태 

→ 이 구조 자체는 **우리가 1단계 크롤러 설계에서 말한 “원문 + 최소 메타데이터” 딱 그** 

**형태** 라서, 로직이 실행된 게 맞아. 

## **2. 위쪽에 community / crewCommunity가 보이는 이유** 

위쪽 몇 줄이 이런 식이잖아: 

- `https://filmmakers.co.kr/community` 

- `https://filmmakers.co.kr/crewCommunity` 

- `https://filmmakers.co.kr/actorCommunity` 

- `https://filmmakers.co.kr/production` 

- `https://filmmakers.co.kr/archive` 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

80/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

- `https://filmmakers.co.kr/academy` 

이건 **게시판의 실제 글이 아니라 상단/왼쪽 메뉴 링크** 를 글로 착각해서 같이 긁어온 거라 서, 

CSS 선택자가 메뉴까지 같이 물어온 상황이야. 

- 로직 입장에서는 “페이지 안에 있는 `<a>` 나 리스트 항목”을 **전부 글처럼 취급** 해 버린 것. 

- 그래서 첫 몇 줄은 “진짜 게시글”이 아니라 “메뉴 링크 정보”가 들어가 있는 거고, 그 뒤에 있는 긴 본문(일본 전시회 얘기)은 **실제 게시글** 이 맞는 상태야. 

## 이건 **로직이 틀렸다기보다는, selector를 조금 더 좁혀야 하는 문제** 라서 

다음 단계에서: 

- “게시글 리스트가 들어 있는 부분만 잡도록 CSS 선택자 수정” 

- 메뉴·광고·배너 영역은 무시 

이렇게 한 번 더 다듬어 주면 돼. 

## 정리하면: 

## **맞아, 이 엑셀은 우리가 설계한 1단계 크롤링 로직이 돌아가서 나온 결과물** 이고, 

- 구조(컬럼)는 설계와 일치 

- 다만 **맨 위 몇 줄에 메뉴 링크까지 같이 잡힌 건 selector 미세 조정으로 빼야 할 부분** 이야. 

## 원하면 다음 턴에서 

`crawler.py` 안에서 “게시글 리스트를 뽑는 CSS 부분”만 같이 보고 

메뉴는 빼고 **진짜 글만 잡도록** 딱 한 군데만 손보자. 

**==> picture [489 x 39] intentionally omitted <==**

자 그렇다면 첫번째 단계를 통해 그냥 긁어온거나 마찬가 지네? 아무것도 없이 그냥? 