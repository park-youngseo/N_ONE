
- `키워드(예: 나이, 돈, 오디션 등)로 글을 골라내지 말 것.` 

- `반드시 할 것:` 

- `"어디서(어떤 게시판인지) + 언제(작성일) + 무엇(원문 텍스트)"만` 

- `최대한 훼손 없이 수집한다.` 

- `2단계 분석(요약, 키워드, 자동 클러스터링)은` 

- `"나중에 붙일 수 있게" 함수와 파일 구조만 준비한다.` 

- `[기능 요구사항 – 1단계: 수집(크롤링) 모듈]` 

`1. 설정 파일` 

- `Python 코드 안에 URL을 박아넣지 말고,` 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

66/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

 

```
별도의설정파일또는 config 모듈에서관리하게만들어라.
```

- `예를 들어, config.py 또는 config.json 형태로 만들어라.` 

- `설정 항목:` 

- `TARGET_BOARDS: 수집할 게시판 URL 목록 (리스트)` 

- `예: ["https://filmmakers.kr/actors?page=1", ...]` 

- `MAX_PAGES: 기본값 100` 

- `PERIOD_OPTION: "3m" / "6m" / "12m" 중 하나 (최근 3개월, 6개월, 1년)` 

- `USER_AGENT: 요청에 사용할 기본 User-Agent 문자열` 

`2. 기간 로직` 

- `오늘 날짜를 기준으로:` 

- `"3m" → 최근 3개월` 

- `"6m" → 최근 6개월` 

- `"12m" → 최근 12개월(1년)` 

- `start_date = 오늘 - 선택한 기간` 

- `end_date = 오늘` 

- `글의 작성일이 start_date보다 이전이면:` 

- `그 글은 수집하지 않는다.` 

- `한 페이지 안의 글들이 전부 start_date보다 이전이면:` 

- `그 게시판에 대해서는 크롤링을 중단한다.` 

- `페이지는 항상 1페이지부터 시작해서,` 

```
  MAX_PAGES(기본 100페이지)까지순서대로내려간다.
```

`3. 글(POST) 수집` 

- `각 게시판 URL에 대해:` 

- `page=1 부터 시작해서 page=MAX_PAGES까지 순차적으로 요청.` 

- `각 페이지의 HTML에서 글 목록을 파싱해서 다음 정보를 추출:` 

- `post_id: 사이트에서 사용하는 글 고유 ID 또는 URL` 

- `source: 사이트 이름 (예: "filmmakers")` 

- `board_name: 게시판 이름 문자열` 

- `url: 글 상세 페이지 URL` 

- `created_at: 작성 시간 (datetime 또는 문자열, 나중에 분석 가능하게)` 

- `title_raw: 제목 원문 (HTML 태그 제거만)` 

- `text_raw: 본문 원문 (HTML 태그 제거만, 내용은 수정하지 않는다)` 

- `view_count: 조회수 (있으면)` 

- `like_count: 추천/좋아요 수 (있으면)` 

- `comment_count: 댓글 수 (있으면)` 

- `text_raw는 다음만 허용한다:` 

- `줄바꿈 정리, HTML 태그 제거, 인코딩 깨짐 복구` 

- `절대 하지 말 것:` 

- `문장 요약, 욕설 삭제, 감정 표현 삭제, 문장 재작성` 

## `4. 댓글(COMMENT) 수집` 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

67/97 

25. 11. 30. 오후 1:44 

 

프로그램 - 연기 지망생 심리 분석 

- `각 글 상세 페이지에서 댓글이 있다면:` 

- `comment_id: 댓글 고유 ID (또는 순번)` 

- `post_id: 어떤 글의 댓글인지` 

- `created_at: 댓글 작성 시간` 

- `text_raw: 댓글 내용 원문 (HTML 태그 제거만)` 

- `이 역시 내용은 수정하지 않는다.` 

`5. 저장 형식` 

- `결과를 CSV 또는 JSON 파일로 저장해라.` 

- `두 가지 파일을 만든다:` 

- `posts.csv (또는 posts.json)` 

- `comments.csv (또는 comments.json)` 

- `posts.csv 컬럼 예:` 

- `post_id, source, board_name, url, created_at,` 

```
    title_raw, text_raw, view_count, like_count, comment_count
```

- `comments.csv 컬럼 예:` 

- `comment_id, post_id, created_at, text_raw` 

`6. 에러 처리 및 로그` 

- `요청 실패(예: 403, 404, 타임아웃)가 발생하면:` 

- `해당 URL과 에러 내용을 로그로 남기고, 프로그램은 계속 진행한다.` 

- `최소한 콘솔 로그(print)와 간단한 로그 파일 기록 기능을 넣어라.` 

- `[기능 요구사항 – 2단계: 분석(요약/키워드/클러스터) 기본 뼈대만]` 

- `주의: 이 단계에서 "심리 해석, 카테고리, 패턴 이름"은 쓰지 않는다. 단지 나중에 분석을 할 수 있게 함수 틀만 만든다.` 

`1. analyzer.py (또는 비슷한 이름의 파일)을 만들어라.` 

`2. 이 모듈에는 다음 함수의 뼈대만 구현해라.` 

- `load_posts(): posts.csv를 읽어서 파이썬 리스트/데이터프레임으로 불러오는 함수` 

- `split_sentences(text): 한 글을 문장 단위로 나누는 함수 (실제 로직 간단하게)` 

- `summarize_post(text): "이 글이 무슨 이야기를 하는지" 한 줄 요약을 만드는 함수` 

- `extract_keywords(text): 자주 나오는 단어를 뽑아내는 함수` 

- `cluster_posts(posts): 글들의 특징을 이용해서 비슷한 글끼리 자동으로 묶는 함수` 

`3. 지금 단계에서는:` 

- `summarize_post, extract_keywords, cluster_posts는` 

   - `아주 간단한 임시 구현으로만 만들고,` 

   - `나중에 내가 바꿀 수 있게 코드 안에 주석을 많이 달아라.` 

`4. 수집(크롤링) 코드와 분석 코드는 반드시 파일을 분리해서 만들어라.` 

- `(예: crawler.py, analyzer.py, config.py)` 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

68/97 

25. 11. 30. 오후 1:44 

프로그램 - 연기 지망생 심리 분석 

```
[요청]
```

 

- `위 요구사항에 맞는 파이썬 프로젝트 구조를 설계하고,` 

```
필요한파일들을모두만들어라.
```

- `각 파일의 코드 전체를 보여줘라.` 

- `각 주요 함수에는 주석을 충분히 달아서,` 

- `나중에 사람이 읽고 수정하기 쉽게 만들어라.` 

- `가상환경, 패키지 설치 목록(requirements.txt)도 함께 제안해라.` 

## **3단계｜안티그래비티에서 코드 생성 → 확인할 포인트** 

안티그래비티가 코드를 뱉으면, 형은 **이것만 먼저 확인** 하면 돼: 

## **1. 파일 구조가 나뉘었는지** 

`config.py` (또는 json) 

- `crawler.py` 

- `analyzer.py` 

```
requirements.txt
```

이런 식으로 나뉘어 있는지 체크. 

## **2. config에 우리가 말한 값이 들어갔는지** 

`TARGET_BOARDS` 리스트가 있는지 

`MAX_PAGES = 100` 같은 값이 있는지 

`PERIOD_OPTION` 이 "3m"/"6m"/"12m" 중 하나로 되어 있는지 

## **3. crawler.py 안에** 

페이지를 1→MAX_PAGES까지 도는 루프가 있는지 

글의 작성일을 보고 기간을 판단하는 부분이 있는지 

posts / comments를 CSV/JSON으로 저장하는 부분이 있는지 

## **4. analyzer.py 안에** 

`summarize_post` , `extract_keywords` , `cluster_posts` 함수가 

“뼈대”라도 만들어져 있는지 

주석으로 “이 부분은 나중에 심화할 수 있다” 정도 설명이 있는지 

이 네 가지만 맞으면, 

## 일단 **안티그래비티 버전 v1 뼈대는 성공** 이야. 

## **4단계｜실행은 “아주 작은 범위”부터** 

실행은 이렇게 테스트하는 게 좋아: 

**1.** `TARGET_BOARDS` 에 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/6923c490-83e8-8323-b837-adcc00804053 

69/97 

25. 11. 30. 오후 1:44 

 

프로그램 - 연기 지망생 심리 분석 

필름메이커스 배우 게시판 URL **딱 1개만** 넣어서 테스트. 

**2.** `PERIOD_OPTION = "3m"` (최근 3개월)로 가장 짧게. 

**3.** 안티그래비티에서 `crawler.py` 를 실행. 

**4.** 끝나면 프로젝트 폴더 안에: 

`posts.csv` , `comments.csv` 가 생겼는지 확인. 

**5.** 파일 열어서 **한두 줄만** 눈으로 확인: 

글/댓글 내용이 “수정되지 않고 그대로” 들어갔는지 

- 날짜, URL 등이 같이 들어갔는지 