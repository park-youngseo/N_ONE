https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/6909f94a-2a34-8321-b76a-af30d91d143d 

13/22 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - 시스템 부팅 설명 

for c in clusters:  f.write(f"- C{c['cluster']:02d} [{c['theme']}] x{c['count']}\n  예: {c['examples']}\n") print(f"\n✅ 리서치 완료 → {det_path}\n   요약 → {sum_path}") else: print("\n(리서치) 고민 문장이 충분하지 않습니다.") # ── 메인 ────────────────────────── ────────────────────── def main(): print("검색어를 입력하세요:") q = input("> ").strip() if not q: print("검색어 없음. 종료."); return q_clean = safe_name(q) 

try: print("수집 개수(기본 50):") n = int(input("> ").strip() or "50") n = max(1, min(MAX_RESULTS_CAP, n)) except: n = 50 

print("엔진 선택: serper / google / naver / duckduckgo / mix") prov = (input("> ").strip().lower() or "serper") 

print("설명(meta)도 모을까요? (y/N)") fetch_desc = (input("> ").strip().lower() == "y") 

try: parts=[] if prov == "serper": parts.append(engine_to_csv(q_clean, "serper", search_serper(q, n), fetch_desc, q)) elif prov == "google": parts.append(engine_to_csv(q_clean, "google", search_google(q, n), fetch_desc, q)) elif prov == "naver": parts.append(engine_to_csv(q_clean, "naver", search_naver(q, n), fetch_desc, q)) elif prov == "duckduckgo": parts.append(engine_to_csv(q_clean, 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/6909f94a-2a34-8321-b76a-af30d91d143d 

14/22 

25. 11. 30. 오후 1:05 

 

**==> picture [137 x 9] intentionally omitted <==**

"duckduckgo", search_ddg(q, n), fetch_desc, q)) elif prov == "mix": 

parts.append(engine_to_csv(q_clean, "serper", search_serper(q, n), fetch_desc, q)) 

parts.append(engine_to_csv(q_clean, "google", search_google(q, n), fetch_desc, q)) parts.append(engine_to_csv(q_clean, "naver", search_naver(q, n), fetch_desc, q)) else: 

print("지원하지 않는 엔진. 종료."); return 

merged_rows, tops = merge_and_save(q_clean, q, parts) 

# ── 추가: 사이트 리서치 모드 ── 

print("\n사이트 리서치 모드로 넘어갈까요? (y/N)") if (input("> ").strip().lower() == "y"): site_research_flow(q_clean, merged_rows, tops) 

print(f"\n✅ 완료: out\\{q_clean}_{prov if prov!='mix' else 'serper'}.csv 등") print(f"   병합본: out\\{q_clean}_all.csv") print(f"   도메인TOP: out\\TOP_{q_clean}_all.csv") print(f"   로그: out\\LOG_{q_clean}_all.csv") 

except Exception as e: print("오류:", str(e)) write_log(q_clean, "예외: " + traceback.format_exc()) print("자세한 내용은 LOG 파일 확인.") 

if __name__ == "__main__": main() 

21s 동안 생각함 

A(개요) — 형, 이 코드는 “검색결과 모으기 → 중복제거/정렬 → 상위 도메인 전체 크롤링 → 고민 문장 뽑기 → 클러스터 요약”까지 한 번에 하는 수집·리서치 도구야. 

B1(핵심) — 바로 돌리기 체크리스트 

## **1.** 필수 설치(한 줄씩): 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/6909f94a-2a34-8321-b76a-af30d91d143d 

15/22 

25. 11. 30. 오후 1:05 

 

**==> picture [137 x 9] intentionally omitted <==**

   - `pip install requests beautifulsoup4 lxml readability-lxml duckduckgo-search` 

   - `google-api-python-client scikit-learn urllib3` 

**2.** 환경변수(없으면 그 엔진은 자동 건너뜀): 

   - Windows → “고급 시스템 설정 > 환경 변수”에 추가 

      - `SERPER_API_KEY` (Serper 쓰면) 

      - `GOOGLE_API_KEY` , `GOOGLE_CSE_CX` (Google CSE 쓰면) 

      - `NAVER_CLIENT_ID` , `NAVER_CLIENT_SECRET` (네이버 쓰면) 

**3.** 실행 순서: 

`python Search.py` → “검색어” 입력 → “수집 개수(기본 50)” → “엔진 선택 

(serper/google/naver/duckduckgo/mix)” → “설명도 모을까요?(y/N)” → 끝나면 

- `out\*_all.csv` , `out\TOP_*_all.csv` 생성 → 리서치 모드 `y` 면 도메인 고르고 크롤 

- 링 진행, 완료 시 `리서치_{도메인}.csv` , `SUMMARY_{도메인}.txt` 생성. 

B2(세부) — 안정·품질 빠른 수리/개선 포인트 

   - 안정성 ①: 일부 사이트가 `Content-Type` 을 비표준으로 줄 수 있어 실패할 수 있음 → 이미 `try/except` 로 무시 처리함. 문제 URL은 `LOG_*_all.csv` 에서 확인. 

   - 안정성 ②: `readability-lxml` 이 간혹 긴 페이지에서 느릴 수 있음 → 

   - `crawl_site(..., delay=0.6)` 을 1.0까지 늘리면 차단 위험 낮아짐. 

   - 정확도 ①: 도메인 내 링크만 따라가도록 `is_same_domain` 이 막아줌. 외부 링크로 튀 지 않음. 

   - 정확도 ②: 고민 문장 추출은 “어떻게/왜/문제/비용…” 같은 힌트로 필터링함. 더 빡세 게 하려면 `QUESTION_HINTS` 에 형이 자주 보는 표현을 몇 개 더 넣어. 

   - 성능 ①: 결과가 200페이지 이상 크면 자동으로 k(클러스터 수)를 적당히 잡음(2~10). 데이터가 적으면 2개로 고정됨. 

   - 사용 팁: `mix` 엔진은 폭넓게 모으지만 API 키 없으면 일부는 빈 결과가 정상이니 로 그만 확인하고 넘어가면 됨. 

   - 윤리/차단: 짧은 지연과 낮은 동시성이라 안전한 편이지만, 민감 사이트는 robots.txt 를 확인하는 게 좋아. 

- [다음] — 원하면 내가 “질문 힌트 보강” 목록이나 “도메인 제외/포함 필터” 옵션을 코드에 소폭 추가해줄게. 

[SELF-CHECK] 1 주제 유지 2 새 숫자/이름 생성 안 함 3 첫 문장 핵심 4 형식 통일 5 군더 더기 최소 6 추론 라벨 불필요 7 시간표기 불필요(설명 문서). 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/6909f94a-2a34-8321-b76a-af30d91d143d 

16/22 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - 시스템 부팅 설명 

 

이 프로그램을 통해 주식에 접목시키고 싶어 어떻게 할수 있지? 

19s 동안 생각함 

A(개요) — 이 도구를 “종목 리서치 비서”로 바꾸면, 

뉴스·공시·커뮤니티를 모아 이슈를 묶고 위험 신호를 뽑아줘. 

B1(핵심) — 최소 수정 7가지 로드맵 

**1.** 키워드 사전. 

티커, 종목명, 별칭을 한 묶음으로. 

예) “삼성전자 OR 005930”. 

**2.** 도메인 필터. 

공식(IR·공시)·증권사·뉴스·커뮤만. 나머지는 제외 리스트에. 

**3.** 기간 제한. 

URL·본문에서 날짜 감지. 

최근 N일만 남기기. 

**4.** 질문 힌트 교체. 

“유상증자, 감자, CB, BW, 

관리종목, 거래정지, 실적, 배당, 지분공시, 자사주, 

공매도, 공시, 악재, 호재”. 

**5.** 감성 라벨 간단화. 

긍정/부정/중립 규칙표 추가. 

예) “상향, 확대”는 긍정. 

“경고, 지정, 하향”은 부정. 

**6.** 이벤트 추출 규칙. 

“실적발표, 배당기준일, 유상증자, 전환청구, 소송, 품목허가, 리콜” 등. 

**7.** 결과물 재구성. 

이슈 클러스터 요약과 함께, “리스크 Top” 표를 추가. (클러스터 수·예시문 포함) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/6909f94a-2a34-8321-b76a-af30d91d143d 

17/22 

25. 11. 30. 오후 1:05 

 

**==> picture [137 x 9] intentionally omitted <==**

B2(세부) — 단계별 적용 방법 

**1.** 입력 단계. 

종목 리스트를 한 줄씩. 티커·종목명·별칭을 저장. 검색어를 자동 생성해. 

**2.** 수집 단계. 

쿼리 템플릿을 고정해. “(티커 OR 종목명) + 키워드”. 키워드는 4묶음이 좋아. ① 공시/IR, ② 실적/가이던스, ③ 리스크 단어, ④ 커뮤 단어. 

**3.** 리서치 단계. 

도메인 화이트리스트만 순회. 사이트 리서치로 전체 탐색. 질문 추출 힌트를 교체해. 기술용어도 함께 넣어. “상한가, VI, 돌파, 매물대, 거래대금, 공매도, 수급, 외인, 기관, 이평, 눌림목”. 

**4.** 정제 단계. 

본문에서 날짜를 인식해. 최근 N일 외는 버리기. 중복 URL과 퍼머링크 제거. 짧은 광고문구도 제거해. 

**5.** 감성·이벤트 단계. 규칙표로 문장에 라벨링. “상향/확대/계약체결 → +1”. “경고/지정/연장/정지 → -1”. 특정 패턴은 이벤트로 기록. “공시번호, 배당, 전환가 등”. 

**6.** 클러스터 단계. 클러스터를 4버킷으로. 

① 실적/펀더, ② 기술/차트, 

③ 이슈/테마, ④ 리스크. 각 버킷의 예시문을 저장. 

**7.** 산출물 단계. 기존 SUMMARY에 더해. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/6909f94a-2a34-8321-b76a-af30d91d143d 

18/22 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - 시스템 부팅 설명 

 

“RISK_SUMMARY.txt” 추가. 리스크 키워드 건수와 예시. “EVENTS.csv”도 함께 저장. 이벤트 유형·날짜·출처 포함. 

**8.** 활용 단계. 

워치리스트 규칙을 정해. “최근 7일 리스크 3건↑” “공시 2건↑·감성 + 합계↑” 이런 조건으로 분류해. 

바로 쓰는 사용 팁(코드 그대로) 

검색어를 이렇게 넣어봐. “종목명 악재”, “티커 공시”, 

- “종목명 실적발표”, “티커 VI”, 

“종목명 유상증자”, “종목명 배당”. 리서치에서 뉴스 도메인, IR 도메인을 우선 선택해. SUMMARY에서 키워드만 봐도, 이슈 흐름을 빨리 파악돼. 

추가 확장 아이디어(선택) 

가격 데이터와 합치기. 일별 뉴스·공시 건수와, 종가·거래대금 상관 보기. 알림 규칙 만들기. 

“리스크 단어 새로 2건↑” 시, TXT에 로그를 남겨줘. 중복 루머 차단 장치. 

같은 문구 반복은 가중치↓. 

주의와 한계 

투자 조언이 아니야. 정보 모음과 분류 도구야. 루머·캡쳐 글은 제외해. 원문 출처가 있는 글만. 규정이 바뀔 수 있어. 항상 최신 공시를 확인해. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/6909f94a-2a34-8321-b76a-af30d91d143d 

19/22 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - 시스템 부팅 설명 

 

원하면 “확장”이라 말해줘. 

리스크/이벤트 규칙표부터, 차근차근 박을게. 

**==> picture [489 x 38] intentionally omitted <==**

초등학생도 이해할수 있게 

9s 동안 생각함 

형, 아주 쉽게 설명할게. 

**1.** 뭐 하는 도구야? 

이 도구는 글을 모아요. 같은 주제는 묶어줘요. 핵심만 짧게 보여줘요. 

**2.** 주식에 쓰는 법(개념) 

종목 관련 글만 모아요. 최근 글만 남겨둬요. 좋은/나쁜 말도 나눠요. 이슈별로 묶어줘요. 

**3.** 준비하기 

종목 이름·티커를 적어요. 

예: “{종목명} OR {티커}”. 

찾을 말도 정해둬요. 

예: 공시, 실적, 배당, 리스크. 

**4.** 돌리는 순서 

첫째, 검색어를 넣어요. “종목명 공시”, “티커 실적”. 둘째, 엔진을 고르세요. serper·google·naver 중 하나. 셋째, 수집 수를 정해요. 기본은 50이면 돼요. 

**5.** 걸러내기 

날짜를 읽어 최근만 남겨요. 같은 링크는 하나만 둬요. 광고 문구는 빼버려요. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/6909f94a-2a34-8321-b76a-af30d91d143d 

20/22 
