
 

프로그램 - AI 그래비티 활용법 

```
    soup = BeautifulSoup(html, "html.parser")
    results = []
# 실제로는사이트구조를보고올바른선택자로바꿔야함
# 예: 각공고가 li, tr, div 등에들어있을수있음
    article_rows = soup.select("/* 여기서각공고한줄을선택하는 CSS 선택자 */")
for row in article_rows:
try:
# 카테고리 (예: 유튜브, 광고/홍보등)
            category = row.select_one("/* 카테고리위치 */").get_text(strip=True)
# 제목과상세링크
            title_tag = row.select_one("/* 제목 a 태그위치 */")
            title = title_tag.get_text(strip=True)
            relative_link = title_tag.get("href", "")
            url = urljoin(BASE_URL, relative_link)
# 제작, 배역, 성별, 출연료, 등록일등
            producer = row.select_one("/* 제작정보 */").get_text(strip=True)
            role = row.select_one("/* 배역정보 */").get_text(strip=True)
            gender = row.select_one("/* 성별정보 */").get_text(strip=True)
            pay = row.select_one("/* 출연료정보 */").get_text(strip=True)
            created_at = row.select_one("/* 등록일정보 */").get_text(strip=True)
            text_for_filter = f"{title}{pay}"
            is_interesting = any(k in text_for_filter for k in KEYWORDS)
            results.append({
"created_at": created_at,
"category": category,
"title": title,
"producer": producer,
"role": role,
"gender": gender,
"pay": pay,
"url": url,
"is_interesting": "Y"if is_interesting else"N",
            })
except Exception as e:
# 한줄에서에러나도크롤러전체가멈추지않게
print("행파싱중에러:", e)
continue
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/69233b7e-1d14-832c-8b71-6d19cbe74ebf 

25/32 

프로그램 - AI 그래비티 활용법 

25. 11. 30. 오후 1:44 

 

```
return results
```

```
defsave_to_csv(rows, filename="filmmakers_performer_casting.csv"):
    fieldnames = ["created_at", "category", "title", "producer",
"role", "gender", "pay", "url", "is_interesting"]
```

```
withopen(filename, "w", newline="", encoding="utf-8-sig") as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
for r in rows:
            writer.writerow(r)
```

```
defmain():
    html = fetch_page(BASE_URL)
    rows = parse_list(html)
    save_to_csv(rows)
print(f"총{len(rows)}개의공고를저장했습니다.")
```

```
if __name__ == "__main__":
    main()
```

- 이 정도 코드가 한 번만 제대로 뽑혀 나오면, 

- 그 다음부터는 **그래비티 없어도 형 PC에서 계속 실행 가능** 한 상태가 되는 거야. 

## **4. 형이 실제로 할 수 있는 순서 (지금 파이썬 깔려 있으니까)** 

형 PC 기준으로, **필름메이커스 v0.1** 크롤러를 만들려면: 

**1.** 폴더 하나 만들기 

   - 예: `C:\dev\filmmakers_crawler` 또는 아무 곳이나 

**2.** 그 안에 `filmmakers_crawler.py` 파일 만들고 

   - 위에 예시 코드(일단 뼈대)를 붙여넣기 

**3.** 라이브러리 설치 

bash 

코드 복사 

```
pip install requests beautifulsoup4
```

## **4.** 실행 

bash 

코드 복사 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/69233b7e-1d14-832c-8b71-6d19cbe74ebf 

26/32 

프로그램 - AI 그래비티 활용법 

25. 11. 30. 오후 1:44 

 

```
python filmmakers_crawler.py
```

→ 같은 폴더에 `filmmakers_performer_casting.csv` 생기면 성공. 

**5.** 나중에 그래비티를 쓰게 되면 

   - 그래비티에서 HTML 구조를 분석해서 

   - 저기 주석으로 남겨둔 CSS 선택자 부분을 **자동으로 넣어주는 코드** 를 뽑게 만들 면 됨 

   - → 그럼 형은 “선택자 맞추기”까지도 그래비티에 넘겨버리는 구조. 

## **핵심 정리 한 줄** 

- **[사실]** 지금은 형이 파이썬 깔려 있으니까, 

## **필름메이커스용 v0.1 크롤러는 파이썬으로 만드는 게 제일 현실적이고 좋다.** 

- **[추론]** 나중에 그래비티를 같이 쓰게 되면, 

   - 이런 코드를 **자동 생성/보정해주는 도구** 로 쓰고 

   - 결과 `.py` 는 형이 계속 가져가서 쓰는 구조로 만들면 됨. 

## 형, 

여기까지가 **“필름메이커스 모델/출연자 모집 크롤러 v0.1 – 그래비티 코드 생성기 조합 을 염두에 둔 파이썬 버전 예시”**야. 

다음 단계는 내가 그냥 이어서 잡아서 가면, 

## **1. 이 크롤러에 “페이지 넘김(2페이지, 3페이지…)” 붙이기** 

**2. 출연료·키워드 필터를 좀 더 똑똑하게 만드는 법** 

**3. 노션/구글시트에 바로 보내는 버전 구조** 

## 중에서 하나를 더 얹을 수 있어. 

형이 굳이 고르지 않아도 되면, 나는 보통 **1번(페이지 넘김)**부터 붙여주는 게 자연스러 워. 

**==> picture [489 x 39] intentionally omitted <==**

다시한번 정확하게 이야기 한다. 제대로 내용을 파악하고 이야기해 

나는 파이썬으로 오류를 계속 체크 하고 수정하는 시간을 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/69233b7e-1d14-832c-8b71-6d19cbe74ebf 

27/32 

프로그램 - AI 그래비티 활용법 

25. 11. 30. 오후 1:44 

 

줄이고 싶은거다. 

따라서 내가 원하는 크롤러의 내용을 그래비티에 입력하 면 

그래비티는 그코드를 파이썬에서 돌아갈수 있게 만들어 주는 방식이다 

24s 동안 생각함 

형 말 정확히 이해했어. 

**“내가 하고 싶은 건 파이썬 코딩/디버깅이 아니라, 

내가 원하는 크롤러 설명을 그래비티에 넣으면 

그래비티가 ‘바로 돌아가는 .py 코드’를 만들어 주는 구조”**지? 

그걸 기준으로 다시 정리해볼게. 

## **A단계｜전체 개요 (형이 원하는 그림 정리)** 

**[사실]** 형이 원하는 구조는 이거야: 

**1.** 형이 **“어떤 사이트에서, 무엇을, 어떻게 긁을지”**를 말로 적음 

**2.** 그 설명을 **그래비티에 넣음** 

**3.** 그래비티가 거기서 **파이썬 코드(.py)를 자동 생성** 

**4.** 형은 그냥 `python xxx.py` **만 실행** 

   - → 중간에 오류 잡고 디버깅하는 시간은 **최소화** . 

- **[추론]** 즉, 그래비티의 역할은 

   - “크롤러 설계 + 코드 생성 + 기본 오류 처리까지 해놓고 넘겨주는 코드 공장” 이고, 

파이썬의 역할은 

- “그 코드를 실제로 실행하는 엔진” 

인 거야. 

## **B1｜핵심 정리 한 줄** 

- **[사실]** 형이 말한 방식은 **논리적으로 완전히 맞는 방향** 이야. 

- **[추론]** 형이 직접 파이썬 코드를 짜거나, 에러 메시지 보면서 수정하는 게 아니라, 

- **그래비티가 “에러 덜 나는 코드”를 대신 만들어주게 설계하는 게 목표** 인 거지. 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/69233b7e-1d14-832c-8b71-6d19cbe74ebf 

28/32 

프로그램 - AI 그래비티 활용법 

25. 11. 30. 오후 1:44 

 

## **B2｜세부 구조 – “그래비티 코드 생성기 조합”이 실제로 어떤 모양이냐** 

형이 머릿속으로 그리기 쉽도록, 

**실제 흐름을 4단계로 쪼개서** 설명할게. 

## **1) 형이 그래비티에 입력하는 것 (설명 템플릿)** 

형이 하는 일은 **코드가 아니라 사양(SPEC)**을 적는 거야. 

예를 들면 이런 식 템플릿을 만들어 둘 수 있어: 

## **1. 타겟 사이트** 

- 주소: `https://www.filmmakers.co.kr/performerCasting` 

로그인 필요 여부: 없음 / 있음(아이디/비번) 

## **2. 수집 범위** 

페이지: 1~3페이지 

기간: 최근 7일만 / 전체 

## **3. 뽑을 정보(필드)** 

카테고리 

- 제목 

- 제작 

- 배역 

- 성별 

- 출연료 등록일 글 링크(URL) 

## **4. 필터 조건** 

제목/내용에 “장편, 단편, 광고, OTT, 고페이” 들어간 것만 

출연료에 “만원, 원” 같은 금액 단위 포함된 것만 등 

## **5. 저장 방식** 

CSV 파일 

- 파일 이름: `filmmakers_캐스팅_YYYYMMDD.csv` 

나중에 노션이나 구글시트로 옮기기 쉽게 

## **6. 실행 방식** 
