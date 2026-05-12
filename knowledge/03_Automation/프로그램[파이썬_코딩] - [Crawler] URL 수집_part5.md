
25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
elif provider == "naver":
            hits = search_naver_api(query, max_n)
elif provider == "kakao":
            hits = search_kakao_api(query, max_n)
else:
raise RuntimeError("알수없는 provider")
except Exception as e:
        write_log("search_error", str(e))
raise
# 중복제거, 순서보존
    seen = set()
    rank = 0
for item in hits:
        u = item.get("url","").strip()
        t = clean_text(item.get("title",""))
ifnot u or u in seen:
continue
        seen.add(u)
        rank += 1
        desc = ""
if fetch_desc:
try:
                desc = fetch_description(u, session=session)
except Exception as e:
                desc = ""
        rows.append({
"rank": rank,
"title": t,
"url": u,
"source": provider,
"query": query,
"fetched_at_kst": now_kst(),
"description": desc,
"domain": domain_of(u)
        })
if rank >= max_n:
break
return rows
defmain():
print("검색어를입력하세요:")
    q = input("> ").strip()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 22/169 

25. 11. 30. 오후 1:05 

 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

```
ifnot q:
"
print(검색어가없습니다. 종료.")
return
try:
print("최대개수(권장 50). 아무것도엔터하면 50:")
        n = int(input("> ").strip() or"50")
        n = max(1, min(MAX_SAFE_RESULTS, n))
except:
        n = 50
```

```
print("검색엔진선택: duckduckgo / google / naver / kakao")
    prov = input("> ").strip().lower() or DEFAULT_PROVIDER
if prov notin ("duckduckgo","google","naver","kakao"):
print("지원하지않는엔진, duckduckgo 사용")
        prov = "duckduckgo"
print("각링크의설명(meta)도수집할까요? (y/N) - 속도느려질수있음")
    fetch_desc = input("> ").strip().lower() == "y"
    write_log("run_request", f"q={q},n={n},prov={prov},desc={fetch_desc}")
try:
        rows = collect_links(q, max_n=n, provider=prov, fetch_desc=fetch_desc)
ifnot rows:
"
print(결과가없습니다. 검색어를바꿔보세요.")
            write_log("no_results", q)
return
        save_results(rows)
        save_top(rows)
print(f"완료: {len(rows)}개수집. 파일: {OUT_CSV}, {TOP_CSV}")
        write_log("run_done", f"saved={len(rows)}")
except Exception as e:
print("오류발생:", str(e))
print("자세한에러는 LOG_Search.csv 를확인하세요.")
        write_log("run_exception", traceback.format_exc())
if __name__ == "__main__":
    main()
```

## 형, 지금 이 스크립트가 **오류 안 나게** 만든 포인트 정리 (짧게): 

## **1.** 기본은 **DuckDuckGo** (라이브러리 없으면 HTML 파싱) → 키 없어도 동작. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 23/169 

25. 11. 30. 오후 1:05 

 

   - 프로그램[파이썬/코딩] - [Crawler] URL 수집 

**2.** Google/Naver/Kakao는 **환경변수 키** 필요 → 없으면 친절한 에러 메시지. 

**3. 타임아웃·재시도** 설정으로 네트워크 불안정 완화. 

**4.** 설명 수집은 옵션으로 둬서 속도 조절 가능. 

**5.** 로그( `LOG_Search.csv` )에 모든 에러·시작·완료 기록. 

**6.** 도메인별 TOP 자동 생성( `TOP_Search.csv` ). 

**7.** 한 번에 과다 요청 방지(MAX_SAFE_RESULTS). 

마지막으로, 형이 지금 원하는 것 딱 하나 알려줘: 

**1.** 그대로 이 코드로 바로 줄까? 

**2.** 아니면 **구글/네이버 API 키 연결 방법 포함한 단계별 가이드** 부터 줄까? 

답은 하나만 말해. 난 지금 바로 진행해서 코드/가이드 전부 제공할게. 

**==> picture [489 x 39] intentionally omitted <==**

2번 

9s 동안 생각함 

좋아, 형. **2번 가이드** 로 간다. 

구글·네이버 **API 키 연결** 을 처음부터 끝까지. 

## **0) 준비 상태** 

파일 이름은 **Search** 고정. 

실행 파일은 **Search.py** . 

저장 파일은 **Search.csv** , **TOP_Search.csv** . 

로그는 **LOG_Search.csv** . 

## **1) 구글 설정(3단계)** 

## **1-1. 프로젝트 만들기** 

**1.** 브라우저에서 **Google Cloud** 열기. 

**2.** 상단 **프로젝트 만들기** 클릭. 

**3.** 이름은 아무거나. 만들기. 

## **1-2. API 키 만들기** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

24/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

**1.** 왼쪽 **API 및 서비스** 이동. 

**2. 사용자 인증 정보** 클릭. 

**3. + 사용자 인증 정보** → **API 키** . 

**4.** 생성된 **API 키** 복사. 

## **1-3. CSE(커스텀 검색) 엔진 만들기** 

**1. Programmable Search Engine** 열기. 

**2. 새 검색 엔진 만들기** 클릭. 

**3.** “검색할 사이트”에 `전체 웹` 선택(또는 빈칸 후 전체웹 허용). 

**4.** 엔진 만들기 완료 후 **제어판** 열기. 

**5. 검색 엔진 ID(CX)** 복사. 

- 👉 여기까지 **두 값** 이 필요해. 

   - `GOOGLE_API_KEY` 

   - `GOOGLE_CSE_CX` 

## **2) 네이버 설정(2단계)** 

## **2-1. 애플리케이션 만들기** 

**1. 네이버 개발자 센터** 접속. 

**2. 내 애플리케이션** → **애플리케이션 등록** . 

**3.** 이름은 자유. 사용 API는 **검색** 체크. 

**4.** 만들기 후 **Client ID** , **Client Secret** 복사. 

- 👉 필요한 값 두 개. 

   - `NAVER_CLIENT_ID` 

   - `NAVER_CLIENT_SECRET` 

## **3) 윈도우에 키 저장(영구)** 

PowerShell을 **관리자** 로 열어. 

아래를 그대로 붙여넣어. 

powershell 코드 복사 

```
setx GOOGLE_API_KEY "여기에_구글_API_KEY"
setx GOOGLE_CSE_CX  "여기에_구글_CX"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 25/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
setx NAVER_CLIENT_ID     "여기에_네이버_ID"
setx NAVER_CLIENT_SECRET "여기에_네이버_SECRET"
```

큰따옴표 유지해. 

완료 후 **터미널을 껐다가** 다시 켜. 

- 그래야 값이 적용돼. 

## **4) 설치 한 번에 끝내기** 

bash 

코드 복사 

```
python -m pip install -U duckduckgo-search requests beautifulsoup4 lxml google-ap
```

  

DuckDuckGo는 키 없이도 됨. 

구글·네이버는 **방금 키** 가 필요. 

## **5) 키가 잘 들어갔는지 확인** 

PowerShell에서 하나씩 찍어봐. 

powershell 코드 복사 `echo %GOOGLE_API_KEY% echo %GOOGLE_CSE_CX% echo %NAVER_CLIENT_ID% echo %NAVER_CLIENT_SECRET%` 

값이 보이면 **성공** . 

빈 줄이면 **다시 입력** 해. 

## **6) Search.py에서 쓰는 법** 

**1.** 실행해. 

bash 

코드 복사 

```
python Search.py
```

**2.** 안내에 답해. 

검색어: 예) 연기 커뮤니티 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 26/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 
