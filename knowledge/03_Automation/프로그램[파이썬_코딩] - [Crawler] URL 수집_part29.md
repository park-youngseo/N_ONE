    tops.sort(key=lambda x:(-x["count"], x["best_rank"]))
    top_fn = os.path.join(OUT_DIR, f"TOP_{keyword}_all.csv")
withopen(top_fn, "w", encoding="utf-8-sig", newline="") as f:
        w=csv.DictWriter(f, fieldnames=["domain","count","best_rank"]); w.writehea
for t in tops: w.writerow(t)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 141/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# 로그
    write_log(keyword, f"병합저장{len(merged)}개")
```

```
# ───────────────────── 메인 ─────────────────────
defmain():
print("검색어를입력하세요:")
    q = input("> ").strip()
ifnot q: print("검색어없음. 종료."); return
    q_clean = safe_name(q)
try:
print("수집개수(기본 50):")
        n = int(input("> ").strip() or"50")
        n = max(1, min(MAX_RESULTS_CAP, n))
except: n = 50
```

```
print("엔진선택: serper / google / naver / duckduckgo / mix")
    prov = (input("> ").strip().lower() or"serper")
```

```
print("설명(meta)도모을까요? (y/N)")
    fetch_desc = (input("> ").strip().lower() == "y")
```

```
try:
        parts=[]
if prov == "serper":
            parts.append(engine_to_csv(q_clean, "serper", search_serper(q, n), fet
elif prov == "google":
            parts.append(engine_to_csv(q_clean, "google", search_google(q, n), fet
elif prov == "naver":
            parts.append(engine_to_csv(q_clean, "naver",  search_naver(q, n), fet
elif prov == "duckduckgo":
            parts.append(engine_to_csv(q_clean, "duckduckgo", search_ddg(q, n), fe
elif prov == "mix":
# 순서: serper → google → naver (중복제거병합)
            parts.append(engine_to_csv(q_clean, "serper", search_serper(q, n), fet
            parts.append(engine_to_csv(q_clean, "google", search_google(q, n), fet
            parts.append(engine_to_csv(q_clean, "naver",  search_naver(q, n), fet
else:
"
print(지원하지않는엔진. 종료."); return
# 병합본만들기 (단일엔진도 all 생성해서통일)
        merge_and_save(q_clean, q, parts)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 142/169 

25. 11. 30. 오후 1:05 

 

후 1:05 1:05 프로그램[파이썬/코딩] - [Crawler] URL 수집 `print(f"\n` ✅ `완료: out\\{q_clean}_{prov if prov!='mix' else 'serper'}.csv print(f"   병합본: out\\{q_clean}_all.csv") print(f"   도메인TOP: out\\TOP_{q_clean}_all.csv") print(f"   로그: out\\LOG_{q_clean}_all.csv")` 

```
except Exception as e:
print("오류:", str(e))
        write_log(q_clean, "예외: " + traceback.format_exc())
print("자세한내용은 LOG 파일확인.")
```

```
if __name__ == "__main__":
    main()
```

## **사용법** 

bash 코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
python .\Search.py
```

엔진: `serper / google / naver / duckduckgo / mix` 중 선택 

설명 수집: `y` 하면 각 URL을 열어 meta/본문 요약 일부를 채워줘(느려질 수 있음) 

이제 파일명/폴더 구조는 전부 **한 가지 규칙으로 통일** 됐고, 설명도 옵션으로 붙였다. 돌려보고 `out` 폴더 화면만 보내줘—필요하면 TOP 정렬 기준(개수→최상위순위)도 미세 조정해줄게. 

**==> picture [489 x 39] intentionally omitted <==**

이건 포털에서 사용할수 있는 기능들이 잖아? 나는 이걸 어떻게 의미있게 사용할수 있지? 

9s 동안 생각함 

형, 맞아. 이건 “포털에서 긁어온 링크”라서 그대로면 반쯤 완성이야. 

**의미 있게 쓰려면 → 목적별로 가공·자동화** 하면 된다. 내가 바로 쓸 수 있게 흐름을 정리 했어. 

## **A) 한 줄 개요** 

수집(csv) → 분류(TOP/도메인/키워드) → 선별(점수·필터) → 행동(지원/연락/기록/알림). 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

143/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## **B1) 바로 써먹는 6가지** 

**1. 오디션/공고 모니터링:** `*_all.csv` 를 날짜별로 쌓고, 새 도메인만 체크. 

**2. 리드 리스트 만들기:** 학원/기획사/커뮤니티 도메인만 화이트리스트로 추려서 연락처 페이지로 이동. 

**3. 경쟁/레퍼런스 연구:** TOP 파일에서 상위 도메인 5~10개만 집중 탐색. 

**4. 콘텐츠 리서치:** 동일 키워드로 주 단위 저장 → 변화(새 링크/사라진 링크)만 비교. 

**5. 퀵 요약:** `description` 켠 뒤(설명 y) 제목+요약 3줄로 훑기 → 후보만 북마크. 

**6. 일간 알림:** 예약 실행(작업 스케줄러)로 아침에 최신 CSV 자동 생성. 

## **B2) 워크플로우(오디션 예시)** 

## **1. 키워드 묶음으로 수집** 

- 예: `연기 오디션` , `드라마 오디션` , `배우 오디션` , `단역 캐스팅` , `영화 엑스트라` 

- 엔진: `mix` → 해외+국내 섞어 확보. 

## **2. 도메인으로 1차 선별** 

TOP_파일에서 상위 도메인 추출 → `audition` , `casting` , `notice` , `apply` 포함 URL만 클릭. 

## **3. 신선도/중복 제거** 

같은 도메인에서 제목 유사(“오디션”, “모집”)만 남기고 이전 주 CSV와 비교(파일 이름에 날짜 붙여 보관). 

## **4. 행동 리스트화** 

“지원 필요”, “문의 필요”, “보류” 세 컬럼 추가해서 Excel에서 체크. 

## **5. 달력화** 

마감일이 있는 글은 달력(Outlook/구글 캘린더)에 수동 입력 → 리마인드. 

## **쓸모 확 올리는 “빠른 확장 7가지” (원하면 바로 코드에 넣어줄게)** 

## **1. 화이트·블랙리스트 필터** 

- 예: 화이트( `audition.kakaoent.com` , `filmmakers.co.kr` ), 블랙( `namu.wiki` , 카페/ 카더라). 

결과 저장 전에 자동 필터링. 

## **2. 랭킹 점수화(우선순위)** 

규칙: 제목 키워드 가중치(“모집”, “지원”, “공지” +), 도메인 가중치(화이트 +, 블 랙 -), 중복 - → `score` 컬럼로 정렬. 

## **3. 도메인별 캡(편중 방지)** 

한 도메인 최대 N개만 저장(예: 5개). 

## **4. 중복 제거(누적 기준)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 144/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

이전 CSV들과 URL 비교해 “새 링크만” 별도 파일 `NEW_검색어_all.csv` 로 저장. 

## **5. 날짜 인식(간단형)** 

본문/요약의 `YYYY-MM-DD` , `~까지` 패턴 찾아 `deadline` 추출 시도. 

## **6. 요약 스니펫 길이 조절** 

`description` 0/100/300자 모드 토글(빠르게 훑기용 ↔ 심층 확인용). 

## **7. 예약 실행(매일 09:00)** 

작업 스케줄러 → “작업 만들기” → 트리거 매일 09:00 → 동작: 

`Program: python` , `Arguments: "C:\Users\...\Search\Search.py"` 

→ 입력 자동화를 원하면 기본 검색어 리스트를 코드에 박아두게 해줄 수 있어. 

## **검색 정확도 올리는 팁(입력 문구)** 

- **구체 키워드:** `연기 오디션 모집 공고` , `단역 캐스팅 공지` , `배우 오디션 접수` 

- **연도/분기:** `2025 연기 오디션` , `11월 오디션` 

- **범위 좁히기:** `site:filmmakers.co.kr 오디션` , `site:instagram.com 오디션` 

- (Serper/Google에서 효과) 

- **영문 병행:** `acting audition korea` , `casting call seoul` (해외 플랫폼까지 감지) 

## **바로 적용하면 좋은 “최소 세팅”** 

검색어 5개 즐겨찾기(메모장): 매일 같은 세트로 실행. 

- 엔진: 일반은 `mix` , 속도 필요하면 `serper` 단독. 

- 설명: 평소 `n` , 후보 좁혀졌을 때만 `y` . 

- 결과 폴더: 주차별 하위 폴더(예: `out\2025-W45\...` )로 이동해서 보관. 

원하면 위 “확장 1~4번(화이트/블랙리스트, 점수화, 캡, 누적 중복 제거)”를 **지금 코드에 바로 붙여서** 드릴게. 

말만 하면, 내가 `Search.py` 에 깔끔하게 추가해줄게. 

**==> picture [489 x 38] intentionally omitted <==**

초등학생도 이해하기 편하게 

형, 아주 쉽게 말할게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 145/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## **A. 이 프로그램은?** 

검색어로 링크를 모아. CSV 파일로 저장해. 

## **B. 어떻게 써?** 

**1.** 파일을 연다. 

**2.** 좋은 걸 고른다. 

**3.** 해야 할 일 적는다. 

**4.** 다시 확인한다. 

## **C. 어떤 파일 봐?** 

`검색어_all.csv` 봐. 

제목·링크가 다 있어. `TOP_검색어_all.csv` 봐. 자주 나온 사이트야. 

`LOG_검색어_all.csv` 는 실행 기록이야. 

## **D. 하루 루틴** 

**1.** 아침에 실행해. 

**2.** `*_all.csv` 를 연다. 

**3.** 제목만 훑어본다. 

**4.** 좋은 것 체크해. 

**5.** 북마크해 둔다. 

**6.** 오후에 다시 본다. 

## **E. 고르는 법** 

제목에 ‘모집’ 찾기. ‘공지’ ‘지원’도 좋아. 기관·회사 도메인 우선. 카더라는 뒤로 미뤄. 

## **F. 더 편하게 하기** 

설명은 나중에 켜. 느려져도 더 보여. 키워드는 구체적으로. 예) 연기 오디션 모집. 

예) 단역 캐스팅 공고. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

146/169 

25. 11. 30. 오후 1:05 
