
(영문으로는 “Search the entire web” 체크) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

41/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## ✅ **3단계 — CX & API 키 등록** 

PowerShell(관리자)에서 입력 👇 (앞에서 만든 API 키와 지금 복사한 CX 둘 다 넣기) 

powershell 

코드 복사 

```
setx GOOGLE_API_KEY "여기_API키"
setx GOOGLE_CSE_CX  "여기_CX값"
```

## → 창을 **닫고 다시 열어야 적용돼.** 

## 확인: 

powershell 

코드 복사 

```
echo %GOOGLE_API_KEY%
```

```
echo %GOOGLE_CSE_CX%
```

## 두 줄 다 값이 나오면 성공이야. 

## ✅ **4단계 — 테스트 실행** 

powershell 

코드 복사 

```
cd /d "C:\Users\psm23\OneDrive\Desktop\Search"
python Search.py
```

검색어: `연기 커뮤니티` 

개수: `50` 

엔진: `google` 

설명 수집: `n` 

## 형이 지금 **none 클릭 후 “기본 탭”에서 CX를 못 찾겠어** ? 

그럼 내가 딱 그 위치 캡처 기준으로 알려줄게. 

혹시 그 화면 보여줄래 (none 클릭 후 안쪽 화면)? 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

42/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

**==> picture [193 x 233] intentionally omitted <==**

## 완벽해 형 👍 

## 이제 진짜 **마지막 단계 바로 앞** 이야. 

## 지금 화면에서 

✅ “검색엔진 ID” → `1216bc4b1659641d4` 이게 바로 **CX 값** 이야. 

- ✅ “전체 웹 검색”도 체크되어 있어서 **완벽한 설정 상태** 야. 

## 이제 **키와 CX 등록만 하면 완료** 야. 👇 

- 1️⃣ PowerShell을 **관리자 권한** 으로 열어. 

- 2️⃣ 아래 그대로 입력 (형의 값 넣어서): 

powershell 코드 복사 

```
setx GOOGLE_API_KEY "여기_아까복사한_구글_API_키"
setx GOOGLE_CSE_CX  "1216bc4b1659641d4"
```

3️⃣ 창을 닫고 **새 PowerShell 열기** 

4️⃣ 확인: 

powershell 코드 복사 

```
echo %GOOGLE_API_KEY%
echo %GOOGLE_CSE_CX%
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 43/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

→ 둘 다 값이 나오면 적용 완료야. 

## 그다음 바로 실행: 

powershell 코드 복사 

```
cd /d "C:\Users\psm23\OneDrive\Desktop\Search"
python Search.py
```

## 입력 예시는 이렇게 👇 

- 검색어: `연기 커뮤니티` 

- 개수: `50` 

- 엔진: `google` 

- 설명 수집: `n` 

끝나면 **Search.csv** 파일이 바로 생성될 거야. 

## 형, 지금 단계에서 

- ✅ API 키 값은 이미 복사해뒀지? 

그럼 내가 **PowerShell 입력 예시 실제 형 키 포함한 형태** 로 정확히 써줄게 — 복사만 하 게. 

지금 키 복사해둔 거 알려줄 수 있어? (앞부분만 보여줘도 돼, 예: AIza…D3) 

**==> picture [489 x 39] intentionally omitted <==**

**==> picture [250 x 192] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

44/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

좋아 형 👌 

지금 화면은 **네이버 개발자 센터** 에서 “애플리케이션 등록” 중이야. 

## 이건 **NAVER 검색 API용 ID/Secret 만들기 과정** 맞고, 

이제 순서대로 이렇게만 해주면 돼 👇 

## ✅ **1단계 — 기본 정보** 

## **애플리케이션 이름** : 

- 👉 아무 이름이나 가능. 예: `SearchTest` 

## **사용 API 선택** 클릭 

- → 검색창에 **검색** 입력 

- → **검색** 항목 체크 ✅ 

- → 저장 

## ✅ **2단계 — 환경 추가** 

- “비로그인 오픈 API 서비스 환경” → **환경 추가** 버튼 클릭 

- → “웹 서버” 선택 

- → “서비스 URL”은 그냥 `http://localhost` 입력 

- → 확인 눌러. 

## ✅ **3단계 — 등록** 

이제 “등록하기” 버튼 눌러. 

완료되면 화면에 

## **Client ID** 

## **Client Secret** 

두 값이 나올 거야. 

## ✅ **4단계 — PowerShell에 등록** 

관리자 PowerShell에서 👇 

(형의 실제 ID/Secret 넣어야 돼) 

powershell 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

45/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
setx NAVER_CLIENT_ID "여기_아이디"
setx NAVER_CLIENT_SECRET "여기_시크릿"
```

## → 창 닫고 **새 PowerShell 열기** 

확인: 

powershell `echo %NAVER_CLIENT_ID% echo %NAVER_CLIENT_SECRET%` 

코드 복사 

둘 다 값이 보이면 성공이야 ✅ 

## ✅ **5단계 — Search.py 실행** 

powershell 코드 복사 

```
cd /d "C:\Users\psm23\OneDrive\Desktop\Search"
python Search.py
```

## 입력 예시: 

- 검색어: `연기 학원` 

- 개수: `50` 

- 엔진: `naver` 

- 설명 수집: `n` 

이제 형은 구글(API 키 + CX)도 완성했고 

네이버(Client ID + Secret)도 거의 끝이야. 

## 즉, **Search.py가 완전한 2엔진 모드로 작동 가능** 상태야 🔥 

다음 단계로 넘어가기 전에 — 

지금 “등록하기” 눌러서 **Client ID/Secret** 화면이 나왔어? 

그 창 캡처나 확인된 값(앞 몇 글자만 보여줘도 돼) 알려줘야 

바로 PowerShell 명령어 예시로 완성해줄게. 

**==> picture [489 x 38] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

46/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

**==> picture [289 x 80] intentionally omitted <==**

9s 동안 생각함 

좋아, 값 잘 넣었어. ✅ 이제 바로 돌려보자. 
