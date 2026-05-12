    log(f"크롤링시작 - {site_name}")
    raw_posts = fetch_posts(urls, years)
    cleaned_rows = []
    seen = set()
for p in raw_posts:
if p.content_raw in seen:
continue
        seen.add(p.content_raw)
        row = clean_post(p)
if row isnotNone:
            cleaned_rows.append(row)
```

```
    filename = f"{site_slug}.{ 'csv'if save_format == 'csv'else'json'}"
    save_path = f"output/{filename}"
```

```
if save_format == "csv":
        save_csv(cleaned_rows, save_path)
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

597/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

 

```
else:
```

```
        save_json(cleaned_rows, save_path)
```

```
    log("크롤러종료")
```

```
if __name__ == "__main__":
```

```
    site_name, site_slug, urls, years, save_format = ask_user_options()
    run_crawler(site_name, site_slug, urls, years, save_format)
```

## ────────────────────────────────── 

## 형, 

이 문서는 지금까지 우리가 만든 스펙 전체를 **빠짐 없이** , **누락 0개** , **합쳐서 하나로** , **그래비티에 그대로 넣으면 바로 적용되는 완전체** 다. 

형이 찾던 “마지막 안내문”, “파일명 규칙”, “URL 기반 구조”, 

“광고 제거/심리 태그/중복 제거/스탭 제외”, 전부 포함된 상태다. 

## 필요하면 다음 단계에서 

## **실제 exe로 빌드할 때 생기는 문제 대비 패치 가이드** 도 만들어줄게. 

**==> picture [489 x 38] intentionally omitted <==**

패치 가이드 만들어 

8s 동안 생각함 

## 좋아 형, 

이제부터는 **욕 안 나오게** “문제 생겼을 때 어떻게 고치는지”를 

**딱 정리된 패치 가이드** 로 박아둘게. 

(앞에서 만든 v2.4 스펙은 “완성본”, 이건 “고장 났을 때 쓰는 매뉴얼” 느낌이야.) 

## **0. 패치 대원칙 (이 3개만 절대 안 깨면 됨)** 

## **1. 엔진 코어 6개는 건드리지 않는다.** 

```
cleaner.py / tagger.py / patterns.py / post.py / saver.py / logger.py
```

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

598/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

 

여기 수정하면 형이랑 나눈 200개 Q&A 규칙이 깨져버리는 구간이라, 

**패치 대상은 항상** `fetcher.py` **/** `main.py` **/ exe 쪽만** 으로 제한. 

## **2. 문제 = “어느 층에서 터졌는지”만 찾으면 된다.** 

UI/입력 문제 → `main.py` 

크롤링/사이트 구조 문제 → `fetcher.py` 

저장/파일 문제 → `saver.py` 호출, 파일 경로, 권한 

엔진 코어는 “정제/태그”만 담당하니까, 여기까지 안 들어오면 크롤링 쪽 문제인 거고. 

## **3. 한번 결정된 규칙은 패치로 바꾸지 않는다.** 

광고 제거, URL 제거, 스탭 제외, 파일명 규칙 등은 **“엔진 규칙”** 패치는 “안 돌아가는 걸 돌게 만드는 것”에만 집중. 

- “룰을 바꾸는 패치”는 나중에 **v3.0** 같은 버전 업으로 따로. 

## **1. 패치 전 공통 체크(형이 할 수 있는 최소 루틴)** 

문제 생기면, 무조건 이 순서로만 보면 돼: 

## **1. 증상 메모** 

예: “exe 누르면 바로 꺼짐”, “파일이 안 생김”, “파일은 있는데 비어있음”, “광고 글까지 같이 담김” 이런 식으로. 

## **2. logs/crawler.log 확인** 

로그 파일이 **생기냐/안 생기냐** 부터 확인 

생긴다면 마지막 줄이 뭐라고 찍히는지 확인 

예: `fetch_posts error` , `save_csv error` , `Permission denied` 등등 

## **3. output 폴더 확인** 

`output/사이트이름.csv` 파일이 생겼는지 

크기가 0인지 / 열면 내용이 있는지 

## **4. 환경 문제 체크** 

Python 버전 (3.10 전후면 무난) 

exe 위치(폴더 권한) 

호출 경로(바탕화면/다운로드에서 실행했는지 등) 

이 4개만 해두면, 

## **“어디 층에서 깨졌는지”** 거의 감이 온다. 

## **2. 상황별 패치 가이드** 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

599/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

 

## **2-1. exe 자체가 실행이 안 될 때** 

## **증상** 

더블클릭하면 아무 반응 없음 

혹은 “Windows에서 실행을 차단했다”류 메시지 

## **원인 후보** 

빌드 자체가 잘못됐거나 

권한 문제 / 보안 프로그램 차단 

## **패치 방향** 

exe 문제는 코드가 아니라 “빌드 설정+보안 정책” 이슈라서, 그래비티/파이썬 쪽에 이렇게 던지면 됨: 

“현재 main.py + fetcher.py는 그대로 두고, 

Windows 10/11 기준으로 잘 실행되는 exe 빌드 설정과 필요한 옵션(pyinstaller 옵션 포함)을 다시 설계해서 알려줘.” 

형이 여기까지 직접 하기 힘들면, 

## **py 상태로 그냥 사용하는 것도 선택지** 고, 

그때는 exe 패치 말고 **main.py/실행방법** 만 손보면 됨. 

## **2-2. UI까지 뜨는데, “실행” 누르면 바로 꺼짐** 

## **증상** 

사이트 이름/URL/기간/저장 형식까지 입력했는데 엔터 치자마자 창이 ‘툭’ 꺼짐 

## **원인 후보** 

`ask_user_options()` 에서 잘못된 입력 처리 

- 예외 발생 후 콘솔 창이 바로 닫힘 

- `run_crawler()` 내부에서 에러 발생 

## **패치 루틴** 

**1. main.py 마지막에** `input("종료하려면 Enter를 누르세요.")` **한 줄 추가** 

에러 메시지가 떠도 창이 바로 안 꺼지고 멈춤 → 원인 확인 가능 

**2.** 그래비티에게 이렇게 요청: 

text 

코드 복사 

```
방금만든 main.py에서
```

- `1) ask_user_options()` 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

600/633 

크롤러 제작 계획 

25. 11. 30. 오후 12:54 

 

```
2) run_crawler()
```

```
두함수각각 try/except로감싸서
에러발생시 logger.log로전체에러메시지를기록하고,
콘솔에 '어디서에러났는지' 한줄로보여줘.
```

```
엔진코어파일은절대수정하지말고,
main.py 안에서만처리해.
```

## → 이렇게 하면 다음부터는 **어디서 터졌는지 로그 보고 바로 알 수 있음.** 

## **2-3. 파일이 아예 안 생길 때** 

## **체크 순서** 

**1.** `logs/crawler.log` 에 

   - “크롤링 시작” 로그가 찍혔는지 

   - “CSV 저장 완료” 같은 로그가 있는지 확인 

## **2. 없다면 → fetcher 단계에서 죽은 것** 

`fetch_posts()` 내부에서 에러 

URL 접근 차단/HTML 파싱 실패 

**3.** 그래비티 패치 요청 템플릿: 

text 코드 복사 

```
지금 fetcher.py에서 fetch_generic_board(url, years) 부분이
어디에서에러나는지알수있도록
```

- `각 주요 단계(페이지 요청, HTML 파싱, 글 리스트 추출, 상세 글 요청, 댓글 추출)마다 logger.log로 단계 이름을 남기고,` 

- `예외 발생 시 logger.log에 에러 타입과 URL을 남겨줘.` 

```
엔진코어(cleaner, tagger, patterns, post, saver, logger)는건드리지말고
fetcher.py 안에서만패치해.
```

## → 이렇게 하면, 

`crawler.log` 보고 “어느 단계까지 갔다가 죽었는지” 한 눈에 보임. 

## **2-4. 파일은 생기는데 “비어있거나, 몇 줄만 있을 때”** 

## **원인 후보** 

https://chatgpt.com/c/69251d70-9f80-8324-ac8e-f14cdd82afa5 

601/633 

25. 11. 30. 오후 12:54 