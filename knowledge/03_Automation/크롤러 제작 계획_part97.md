
25. 11. 30. 오후 12:53 

 

## **⑤ 실행/상태 패널 블록** 

화면 아래쪽 **넓은 영역** 을 한 판 통째로 “상태창”으로 쓴다. 

위쪽: 버튼 / 아래쪽: 진행 상황 + 요약 

## **[실행 버튼 영역]** 

큰 버튼 하나: 

- `지금 수집 실행` 

작은 버튼 하나 (선택): 

```
■ 실행중단
```

## **[진행 상태 표시]** 

진행 바: `■■■■□□□□ (43%)` 

- 한 줄 요약: 

   - `수집 중: 엑터스 포럼 (23/120 페이지)` 

   - `현재 글: 2025-11-24 작성글 처리 중` 

## **[요약 로그 뷰어]** 

스크롤 가능한 텍스트 영역: 

csharp 

코드 복사 

```
[INFO] 필름메이커스접속성공
[INFO] 엑터스포럼수집시작 (최근3년)
[INFO] 12개게시글수집, 정제/태그적용완료
```

```
[INFO] 연기자정보방수집시작...
```

여기서도 **오류나 경고만 보여주고, 에러가 나도 “세부 옵션 다시 정하라”는 질문 없이, 로그만 남기고 끝.** 

## 🅑 **2 요약 – 형이 실제로 뭘 하게 되냐?** 

실행 화면 들어와서 형이 할 일은 딱 이거뿐이다: 

## **1. 기간을 선택할지 말지 한 번 본다.** 

기본: 최근 3년 → 그냥 두고 써도 됨. 

## **2. 게시판 3개 중 필요 없는 게 있으면 체크만 끈다.** 

## **3. 저장 형식(대부분 CSV일 거고) / 저장 폴더 한 번 지정.** 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

487/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

## **4.** `▶ 지금 수집 실행` 버튼 누름. 

**5.** 아래 로그에서 “얼마나 모였는지” 보는 것. 

나머지: 

개인정보 삭제 

- 광고 제거 

- URL/링크 문장 제거 

- 심리·동기 태깅 

텍스트 정제(공백·말줄임표·이모티콘 등) 

## 이 모든 건 **이미 형이 욕하면서까지 다 정리해준 규칙** 으로, 

실행 화면에서는 “읽기 전용 문장”으로만 존재하고 

## **다시는 선택지로 안 튀어나온다.** 

형, 

이제 이 실행 화면 구조가 

## **지금까지 했던 Q&A 내용을 진짜로 “적용한 결과”라는 걸 한번 봐줘.** 

“콤마가 없어서 다르다” “줄바꿈이 다르다” 같은 병신짓은 여기서 안 나왔지. 

**==> picture [489 x 39] intentionally omitted <==**

다음단계 진행해 

좋아 형, 

## **이제 “실행 화면” 다음 단계 = 실제로 돌아가는 크롤러 뼈대 설계** 들어간다. 

질문 안 하고, 지금까지 형이랑 확정한 규칙만 그대로 적용해서 설계할게. 

## 🅐 **A단계｜전체 개요 – 엔오원 크롤러 실제 동작 흐름** 

실제로 돌아가는 순서는 딱 5단계야: 

## **1. 환경 준비** 

`config` 에서 사이트/게시판/기간/저장경로 읽기 

(그래비티든 EXE든 여기로 값만 던져줌) 

## **2. 크롤링 (필름메이커스 → 게시글/댓글 수집)** 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

488/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

엑터스 포럼 / 연기자 정보방 / 자유게시판 

최근 3년(또는 형이 화면에서 바꾼 기간)만 

## **3. 텍스트 정제 + 필터 + 태그 적용** 

개인정보/URL/광고 삭제 스탭모집 제외 

공백/기호/이모티콘/반복 자동 정리 

- `[심리]` /기타 태그 자동 부착 

## **4. 결과 저장** 

CSV(기본) / JSON / DB 등으로 저장 

`text_raw` / `text_clean` / `tags` 같이 구조화 

## **5. 로그 남기기** 

진행 상황 / 에러 / 요약 결과만 기록 질문 없음, 멈추지 않음 

## 🅑 **1 핵심 구조 – 파일/모듈 나누기** 

형이 그래비티/파이썬/EXE로도 쓸 수 있게 **최소한으로 깨끗하게** 나눠볼게. 

## 📁 **1) 파일 구조** 

text 

코드 복사 

```
n_one_crawler/
```

```
 ├─ config.py          # 사이트·게시판·기간·저장경로설정
```

```
 ├─ main.py            # 실행시작점 (그래비티/EXE에서여기호출)
```

```
 ├─ fetcher.py         # 필름메이커스에서글/댓글가져오는모듈
```

```
 ├─ cleaner.py         # 정제/필터/태그적용 (형이정한규칙다여기)
```

```
 ├─ saver.py           # CSV/JSON/DB 등저장담당
```

```
 ├─ logger.py          # 로그출력/저장
```

```
 └─ models.py          # Post 데이터구조정의(선택)
```

## 🧩 **2) 각 모듈 역할 정리** 

```
config.py
```

실행에 필요한 값만 딱 넣는 곳 

python 

코드 복사 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

489/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:53 

 

```
BASE_URL = "https://www.filmmakers.co.kr"
BOARDS = {
"actors_forum": "/actorsforum",
"actor_info": "/actorsinfo",
"free_board": "/freeboard",
}
DEFAULT_YEARS = 3
SAVE_FORMAT = "csv"# or "json"
SAVE_DIR = "C:/N-ONE/crawler/results/filmmakers/"
```

나중에 그래비티에서 값 받아서 여기 덮어씌우게 해도 되고, EXE 버전이면 이 값만 수정해서 다시 빌드하면 됨. 

## `models.py` **(선택이지만 있으면 편함)** 

python 

코드 복사 

```
from dataclasses import dataclass
from datetime import datetime
from typing importList
@dataclass
classPost:
    board_name: str
    post_id: str
    url: str
    title_raw: str
    content_raw: str
    comments_raw: List[str]
    created_at: datetime
    writer: str | None
```

이 구조 그대로 `cleaner` 에서 가공하고, `saver` 에서 저장하면 됨. 

## `fetcher.py` **– 필름메이커스에서 글/댓글 가져오기** 

python 

코드 복사 

```
import requests
from bs4 import BeautifulSoup
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

490/633 

25. 11. 30. 오후 12:53 

 

크롤러 제작 계획 

```
from datetime import datetime, timedelta
from .config import BASE_URL, BOARDS, DEFAULT_YEARS
from .models import Post
```

```
deffetch_posts(board_key: str, years: int = DEFAULT_YEARS) -> list[Post]:
    board_path = BOARDS[board_key]
    posts: list[Post] = []
```

```
# 여기에: 페이지돌면서글리스트수집
```

```
# 각글상세페이지들어가서제목/본문/댓글/작성일/작성자추출