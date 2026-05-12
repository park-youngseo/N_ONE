**3.** 엔진 생성 후 **제어판** 에서 **검색 엔진 ID(CX)** 복사. 

## **5) 윈도우에 키 저장** 

PowerShell(관리자)에서 아래 실행 → 창 **다시 열기** : 

powershell 코드 복사 `setx GOOGLE_API_KEY "여기_API키" setx GOOGLE_CSE_CX  "여기_CX값"` **6) 파이썬 설치 한 번에** powershell 코드 복사 

## **6) 파이썬 설치 한 번에** 

```
python -m pip install -U duckduckgo-search requests beautifulsoup4 lxml google-ap
```

  **7) 키 적용 확인** powershell 코드 복사 

## **7) 키 적용 확인** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

36/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
echo %GOOGLE_API_KEY%
echo %GOOGLE_CSE_CX%
```

값이 보이면 OK. 

## **8) 실행** 

powershell 

코드 복사 

```
cd /d "C:\Users\psm23\OneDrive\Desktop\Search"
python Search.py
```

- 엔진 물으면 `google` 입력. 

- 개수: 50 추천. 

- 설명 수집: 처음엔 `n` (빠름). 

## **9) 결과 파일** 

링크 목록: **Search.csv** 

- 도메인 집계: **TOP_Search.csv** 

- 로그: **LOG_Search.csv** 

## **10) 자주 막히는 곳(빠른 해결)** 

- “구글 키 없음” → 5단계 다시, 창 재시작. 

- 결과 적음 → 개수 100, 검색어 넓히기. 

- 403/429(요청 과다) → 잠시 후 재시도, 키 제한 완화 또는 쿼리 줄이기. 회사/학교망 이슈 → 집/모바일 핫스팟에서 테스트. 

원하면 지금 화면 어디까지 했는지 말해줘. 거기서 다음 클릭 위치까지 바로 이어서 안내 할게. 

**==> picture [489 x 39] intentionally omitted <==**

**==> picture [289 x 96] intentionally omitted <==**

**==> picture [65 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
안나와<br>**----- End of picture text -----**<br>


https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

37/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

17s 동안 생각함 

형, 지금 화면은 **키가 아직 없을 때** 가 맞아. 

아래 순서 그대로 하면 바로 생겨. 

## **1) API 켜기(반드시 먼저)** 

**1.** 왼쪽에서 **라이브러리** 클릭. 

**2.** 위 검색창에 **Custom Search API** 입력 → 엔터. 

**3.** 결과에서 **Custom Search API** 선택 → **사용 설정(Enable)** 버튼 누르기. 

   - 버튼이 **사용 중(Enabled)** 으로 바뀌면 성공. 

만약 검색이 안 보이면: 상단 프로젝트가 **My First Project** 로 되어 있는지 확인해줘. 

## **2) API 키 만들기** 

**1.** 왼쪽 **사용자 인증 정보** 로 돌아오기. 

**2.** 상단 파란색 **+ 사용자 인증 정보 만들기** 클릭 → **API 키** 선택. 

**3.** 팝업에 **API 키가 생성되었습니다** 뜸 → 키 복사. 

**4.** 같은 팝업에서 **키 제한** 클릭 → 

**API 제한** : **Custom Search API** 만 허용으로 선택 → 저장. 

- (앱 제한은 기본값 그대로 OK) 

## **3) CSE(검색 엔진) ID 만들기(CX)** 

**1.** 새 탭에서 **Programmable Search Engine** 를 검색해 접속. 

**2. 새 검색 엔진 만들기** 클릭. 

**3.** “검색할 사이트”는 일단 아무 사이트 넣고 생성 → 

   - 생성 후 **제어판** → **기본** 탭에서 

   - **“검색할 사이트: 전체 웹”** 으로 변경(스위치/옵션). 

**4.** 제어판 상단의 **검색 엔진 ID(CX)** 복사. 

## **4) 키 저장(윈도우)** 

PowerShell(관리자)에서: 

powershell 코드 복사 

```
setx GOOGLE_API_KEY "여기_API키"
setx GOOGLE_CSE_CX  "여기_CX값"
```

## → 창 **닫고 다시 열기** → 아래로 확인: 

powershell 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

38/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
echo %GOOGLE_API_KEY%
echo %GOOGLE_CSE_CX%
```

## **5) 실행 확인** 

powershell 

코드 복사 

```
cd /d "C:\Users\psm23\OneDrive\Desktop\Search"
```

```
python Search.py
```

- 엔진: `google` 

- 개수: `50` 

- 설명 수집: `n` (빠르게 테스트) 

## **막히면 여기를 체크** 

- **라이브러리에 Custom Search API가 안 보임** → 프로젝트 선택 잘못됨. 상단 프로젝 트 바꾸기. 

- **Enable 버튼이 회색** → 결제 계정 연결을 요구할 수 있어. 카드 등록만 하고 과금은 쿼터 내 무료. 

**Search.py에서 “구글 키 없음”** → 4단계 다시, 그리고 **PowerShell 재시작** . 

지금 어디 단계에서 멈췄는지(라이브러리? 키 생성? CX?) 한 단어로만 알려줘. 거기서 이 어서 눌러줄 버튼까지 정확히 짚어줄게. 

**==> picture [489 x 39] intentionally omitted <==**

**==> picture [263 x 192] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 39/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

좋아, 형. 지금 화면은 **Custom Search API가 ‘사용 설정됨’** 상태야. 

이제 딱 3가지만 하면 끝나. 

## **1) API 키 만들기** 

**1.** 왼쪽 **API 및 서비스 → 사용자 인증 정보** 이동 

**2.** 상단 파란 **+ 사용자 인증 정보 만들기** → **API 키** 

**3.** 뜬 팝업에서 키 복사 → **키 제한** 클릭 

   - **API 제한** 에서 **Custom Search API** 만 허용 → 저장 

## **2) CSE(검색 엔진) ID 만들기 = CX** 

**1.** 새 탭에서 **Programmable Search Engine** 접속 

**2. 새 검색 엔진 만들기** → 일단 아무 사이트로 생성 

**3.** 생성 후 **제어판** → **기본** 에서 “검색할 사이트”를 **전체 웹** 으로 변경 

**4.** 제어판 상단의 **검색 엔진 ID(CX)** 복사 

## **3) 윈도우에 키 저장 → 확인** 

PowerShell(관리자): 

powershell 코드 복사 `setx GOOGLE_API_KEY "여기_API키" setx GOOGLE_CSE_CX  "여기_CX값"` 창 닫고 다시 열고 확인: powershell 코드 복사 `echo %GOOGLE_API_KEY% echo %GOOGLE_CSE_CX%` 

## 창 닫고 다시 열고 확인: 

## **실행 체크** 

powershell 코드 복사 

```
cd /d "C:\Users\psm23\OneDrive\Desktop\Search"
python Search.py
```

엔진: `google` 

개수: `50` 

설명 수집: `n` (빠르게 테스트) 

## **막히면** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

40/169 

25. 11. 30. 오후 1:05 

 

   - 프로그램[파이썬/코딩] - [Crawler] URL 수집 

- **+ 사용자 인증 정보 만들기** 가 회색이면 → 상단 프로젝트가 맞는지 확인. 

- **전체 웹 전환 옵션이 안 보임** → 제어판의 **기본/세부정보** 탭 살펴보고 “검색할 사이 트” 항목에서 전환. 

지금 어디에서 멈췄는지(①키, ②CX) 한 단어만 말해줘. 거기서 다음 클릭 위치까지 딱 집 어줄게. 

**==> picture [489 x 38] intentionally omitted <==**

**==> picture [289 x 112] intentionally omitted <==**

좋아 형 👌 지금 완전 잘 왔어. 이제 이 화면에서 **CX (검색 엔진 ID)** 값을 복사하면 돼. 단계만 한 번 따라가면 끝난다. 

## ✅ **1단계 — 엔진 ID 확인하기 (CX 복사)** 

**1.** 목록에서 **none** 클릭 (파란 글씨 부분). 

**2.** 왼쪽 메뉴에서 **기본** (Basic) 탭 선택. 

**3.** 오른쪽에 “검색 엔진 ID” 또는 “Search engine ID” 항목이 있어. 

   - → 그 옆에 있는 긴 영문·숫자 코드 복사해. 

   - 예: `0123456789abcdefg:xyz123abcde` 이런 형태야. 

   - 👉 이게 바로 **CX 값** 이야. 

## ✅ **2단계 — 검색 범위 전체 웹로 바꾸기** 

**1.** 같은 **기본 탭** 에서 

   - “검색할 사이트”(Sites to search) 항목이 보일 거야. 

**2.** 그 옆에 **편집(Edit)** 또는 **전체 웹 포함(Include entire web)** 같은 옵션이 있으면 클릭해서 “전체 웹”으로 변경. 