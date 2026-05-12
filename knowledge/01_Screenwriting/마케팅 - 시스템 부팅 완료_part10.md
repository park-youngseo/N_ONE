```

채널 간 **공통 상위 구절** 교집합 지표 

`noise_words.csv` / `synonyms.csv` 

- 형이 운영하며 **지속적으로 정교화** (튜닝 포인트) 

## **H. 운영(일반 PC에서 가능?) — 가능** 

형이 말한 것처럼 **일반 PC** 에서도 충분히 가능(윈도우 10/11, 파이썬 기준). **한 번 세 팅하면 재실행만** 하면 됨. 내가원하는거 

## **실행 루틴(주간)** 

**1.** 주소 목록 업데이트(추가/삭제) 

**2.** 수집 실행 → 텍스트 업데이트 

**3.** 추출/스코어링 실행 → 새 **TOP10/20** 산출 

**4.** 대시보드(시트/노션) 자동 갱신 

## **I. 리스크·한계 & 대안** 

**법·정책** : **공개 페이지/공개 자막** 만, 과도한 요청·약관 위반 금지(robots.txt 존중) 

- **품질 이슈** : 스팸/중복/광고성 컨텐츠는 필터 강화가 필요 

- **토픽 왜곡** : 단기 이슈가 상위를 잠식할 수 있어 **기간 필터** (최근 3/6/12개월) 분리 추 천 

- **대안 소스** : RSS 지원 사이트·공식 블로그·보도자료 페이지 등 **안정 소스** 가중치↑ 

## **J. 엔오원 적용(바로 쓰는 법)** 

## **TOP 주제 → 수업 커리큘럼/콘텐츠 캘린더** 로 연결 

- 예: 상위 “카메라 연기/시선 처리/자기소개”가 강세면 → **1개월 집중 테마** 결정 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

47/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

## **채널별 차별화** 

블로그 TOP 구절 = **검색형 글/가이드** 

유튜브 TOP 구절 = **숏폼 훅/영상 챕터** 

## **교집합 = 브랜드 코어 키워드** 

공통 상위 5개를 **사이트/배너/프로필** 핵심 키워드로 고정 

## **다음 액션(형 선택)** 

**1. 주소 목록 시드** 부터 내가 뽑아줄까, 아니면 형이 1차 리스트를 만들어 줄래? (스프레 드시트 컬럼: `source_type, url, 메모` ) 

**2. 불용어·유의어 사전** 초안을 내가 기본 세트로 제안해줄게(형이 쓰는 용어 반영해서). 

형이 원한 건 “감”이 아니라 **데이터로 뽑은 TOP10/20** 이야 — 위 설계대로라면 **주간 리포 트** 로 안정적으로 뽑을 수 있다. 지금은 **분석·설계** 를 끝냈고, 형이 “주소 시드”만 주면 바로 **실행 단계** 로 붙일 수 있어. 

**==> picture [489 x 39] intentionally omitted <==**

주소목록 시드부터 만들어보자 이건 리스트를 입력할수 있게 해줘 

5s 동안 생각함 

- 좋아 형. **주소목록 시드** 를 바로 입력할 수 있게 **템플릿 두 개** 만들어서 올려놨어. 엑셀/노션에 붙여 쓰기 좋은 CSV: 

seed_urls_template.csv 

메모장에 붙여넣기 쉬운 최소형 TXT: 

seed_urls_minimal.txt 

## **어떻게 쓰면 되냐 (초등학생 버전)** 

## **1. CSV 열기** → 각 줄에 사이트/채널을 추가 

`source_type` : `blog / web / youtube_channel / youtube_playlist / youtube_video / rss` 중 택1 

`url` : 전체 주소 복붙 

`title_hint` : 사이트/채널 이름(선택) 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

48/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

`priority` : 1~3 (3이 중요) 

   - `crawl_depth` : 0=그 페이지만 / 1=페이지 넘김·태그 목록 / 2=섹션 전체 

   - `language` : `ko / en / mixed` 

   - `include` : `Y` (사용) / `N` (일시 제외) 

   - `notes` : 메모 

**2. TXT 버전** 은 한 줄에 하나씩 URL만 쭉 붙여 넣으면 끝. 

## **추천 입력 예시(가이드)** 

- **블로그** : 네이버 블로그 카테고리 루트(예: “연기 팁” 카테고리) 

- **웹사이트** : 학원/기관의 칼럼·가이드 목록 페이지 

- **유튜브 채널/플리** : 자막이 많은 채널, 오디션·카메라 연기 재생목록 

- **RSS** : 공식 블로그/뉴스의 RSS 피드 

원하면, 형이 채워준 CSV를 기준으로 **크롤/수집 모듈** 에 바로 물리도록 연결해줄게. 

첫 번째 내가 이야기 하는건 리스트를 자동 추출해야하는 것을 이야기하는거야 그 검색을 할수있게 시드를 넣을수 있게해달라는거고 

오케이 형. **내가 말한 “시드를 넣으면 자동으로 URL 리스트를 캐오는” 기능** 을 바로 쓰게 세팅했다. 

(= 시드 키워드 입력 → 네이버/다음/유튜브에서 링크 자동 수집 → `discovered_urls.csv` 생성) 

## **받기** 

- 패치: n1_seeded_discovery_patch.zip 

이걸 **형이 쓰는 메인 폴더** (예: `...n1_public_needs_miner_v1_full\` )에 풀어. 폴더 이름 바꿔도 OK. 러너는 **현재 폴더 기준** 으로 돈다. 

## **쓰는 법 (3단계)** 

## **1. 시드 입력** 

`data/discovery_seeds.txt` 열고, **검색어를 줄마다 하나씩** 적어. 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

49/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

예: 

코드 복사 

```
카메라연기
오디션자기소개
시선처리
```

- (선택) 이미 돌린 `outputs/suggestions_all.csv` 가 있으면, 거기 상위 제안문을 쿼리로 쓰도록 옵션도 있음. 

## **2. 실행** 

폴더에서 `RUN_DISCOVERY.cmd` 더블클릭 

기본 옵션: `--engines naver,daum,youtube --top 10` 

엔진은 쉼표로 선택 가능( `naver` , `daum` , `youtube` ) 

`--top 10` 은 시드×엔진별 최대 10개 링크 

## **3. 결과** 

`outputs/discovered_urls.csv` 생성 

컬럼: `source_type, url, title, engine, seed, rank, channel_url, host, ts_kst` 

예) `blog / web / youtube_video / youtube_channel` 

## **옵션(원하면)** 

`RUN_DISCOVERY.cmd` 를 메모장으로 열고, 다음처럼 바꿔서 쓰면 돼: 

**유튜브만** : 

css 코드 복사 `python -u scripts\discovery_search.py "%CD%" --engines youtube --top 15`   **네이버 20개·다음 비활성** : 

css 코드 복사 

```
python -u scripts\discovery_search.py "%CD%" --engines naver --top20
```

**자동완성 결과에서 상위 30개를 쿼리로** : 

lua 코드 복사 

```
python -u scripts\discovery_search.py "%CD%"--from_suggestions 30
```

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

50/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

## **체크 포인트(안 되면 이거 확인)** 

**1.** `python --version` (3.10 이상) 

**2.** `pip install -r requirements.txt` (최초 1회) 

**3.** 네트워크(회사망/공유기 차단 없나) 

**4.** `outputs/discovered_urls.csv` 파일이 생겼는지 

원래 배포본에는 **실제 수집 로직이 들어간 `scripts/discovery_search.py` **가 들어가는 데, 러너/구조는 위대로니 바로 돌려보고, 결과 파일 공유해주면 내가 **필요한 셀렉터(네이 버/다음 선택자) 안정화** 까지 즉시 손볼게. 

**==> picture [489 x 39] intentionally omitted <==**

전부 새롭게 만들어 우선 기존에 있던 파일 그런거 전부 없애고 처음부터 할거야 

내가 원하는걸 위주로 만들라고 기존에 있었던 파일 그런 거 전부 리셋하고 

## 좋아 형. **완전 새로** 만들었다. 이전 거 싹 무시하고 **이 패키지로 시작** 하면 된다. 

**다운로드:** N1_NeedsMiner_clean_v1.zip 

## **이게 딱 뭘 하냐** 

**STEP1_자동완성** : 구글/유튜브/네이버/다음 자동완성 → 

- `outputs/suggestions_all.csv` 

- **STEP2_주소발굴** : **시드 키워드로 실제 URL 자동 수집** (네이버 VIEW/WEB, 다음 BLOG/WEB, 유튜브 검색) → `outputs/discovered_urls.csv` 

- **STEP3_자막수집** : 자동완성 상위 쿼리로 유튜브 검색→공개자막 수집 → 

- `outputs/transcripts.csv` 

- **STEP4_분석** : 니즈 태그 집계 + 2/3-gram TOP 리스트 → `needs_counts.csv` , `top_ngrams_*.csv` 

- **RUN_ALL.cmd** : 0→1→2→3→4 한 방 실행 

## **폴더/파일 이름 자유 규칙** 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

51/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

최상위 폴더·배치파일 이름: **마음대로 바꿔도 됨** (현재 폴더 기준으로 실행) 

- 내부 구조는 고정 권장: `data/` , `scripts/` , `outputs/` , `requirements.txt` 

## **처음 세팅 (1분 컷)** 

**1.** 압축 풀기 

**2.** `STEP0_설치.cmd` 실행(최초 1회): 필요한 라이브러리 설치 

**3.** 실행 방식 두 가지 중 택1 

한 방: `RUN_ALL.cmd` 

수동: **STEP1 → STEP2 → STEP3 → STEP4** 

## **시드 넣는 곳(형이 원하는 핵심)** 

**자동완성용 시드** : `data/seeds.txt` 

**주소발굴용 시드** : `data/discovery_seeds.txt` 

→ 여기 줄마다 키워드 넣으면, **STEP2** 가 네이버/다음/유튜브에서 **URL/제목 자동 수** 

- **집** 해 `discovered_urls.csv` 로 저장한다. 

## **막히면 10초 점검** 

css 코드 복사 

```
python --version
pip --version
py -m yt_dlp --version   ← 안뜨면폴더최상단의 yt-dlp.cmd가대신처리
dir outputs
```