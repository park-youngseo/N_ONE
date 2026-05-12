## **1) 창 재시작 → 값 확인** 

PowerShell을 **닫고 다시 열기** . 

아래로 값이 보이면 OK. 

powershell 코드 복사 `echo %NAVER_CLIENT_ID% echo %NAVER_CLIENT_SECRET%` 

## **2) 패키지 점검(한 번만)** 

powershell 코드 복사 `python -m pip install -U requests beautifulsoup4 lxml duckduckgo-search google-ap`   **3) 폴더 이동 → 실행** powershell 코드 복사 

## **3) 폴더 이동 → 실행** 

```
cd /d "C:\Users\psm23\OneDrive\Desktop\Search"
python Search.py
```

## 입력 예시: 

검색어: `연기 커뮤니티` 

개수: `50` 

엔진: `naver` 

설명 수집: `n` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

47/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

완료되면 폴더에 

`Search.csv` , `TOP_Search.csv` , `LOG_Search.csv` 가 생겨요. 

## **4) 자주 나는 오류 빠른 해결** 

```
네이버키없음
```

→ echo 값 비었는지 확인, 창 다시 열기. 

```
401/403
```

- → 네이버 콘솔에서 **사용 API=검색** 체크됐는지, 

**환경: 웹서버 /** http://localhost 등록됐는지 확인. 

결과가 적다 

→ 개수 100, 키워드 더 넓게. 

바로 실행해보고, 터미널에 뜬 마지막 두 줄만 보내줘. 거기서 이어서 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

## Windows PowerShell 

Copyright (C) Microsoft Corporation. All rights reserved. 

새로운 기능 및 개선 사항에 대 한 최신 PowerShell을 설 치 하세요! https://aka.ms/PSWindows 

PS C:\WINDOWS\System32> cd /d 

"C:\Users\psm23\OneDrive\Desktop\Search" 

>> python Search.py 

>> 

Set-Location : 

'C:\Users\psm23\OneDrive\Desktop\Search' 인수를 허용하는 위치 매개 변수를 찾을 수 없습니다. 위치 줄:1 문자:1 

+ cd /d "C:\Users\psm23\OneDrive\Desktop\Search" 

+ 

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ 

~~~~~~~~~ 

+ CategoryInfo          : InvalidArgument: (:) [SetLocation], ParameterBindingException 

+ FullyQualifiedErrorId : PositionalParameterNotFound,Microsoft.PowerShell.Com mands.SetLocationCommand 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

48/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

C:\Users\psm23\AppData\Local\Programs\Python\P ython311\python.exe: can't open file 

'C:\\WINDOWS\\System32\\Search.py': [Errno 2] No such file or directory PS C:\WINDOWS\System32> 

16s 동안 생각함 

형, 이건 **PowerShell 문법** 때문에 난 거야. 

- `/d` 는 CMD용이야. PowerShell에서는 그냥 경로만 쓰면 돼. 

## **1) 폴더로 제대로 이동** 

아래 중 **첫 줄** 부터 해봐. (안 되면 2·3번 순서로 시도) 

powershell 코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
cd"$env:OneDrive\Desktop\Search"
cd"$env:OneDrive\바탕화면\Search"   # 한글 Windows일때
```

## 폴더가 맞는지 확인: 

powershell 코드 복사 `pwd dir Search.py` 가 보여야 정상. **2) 바로 실행 (현재 폴더에서)** powershell 코드 복사 `python .\Search.py` 

여기서 `Search.py` 가 보여야 정상. 

## **2) 바로 실행 (현재 폴더에서)** 

## **3) 이동이 계속 안 될 때(우회 실행)** 

경로를 **직접 넣어서** 실행해도 돼: 

powershell 코드 복사 

```
python "C:\Users\psm23\OneDrive\Desktop\Search\Search.py"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 49/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## **4) 체크포인트** 

`Test-Path "C:\Users\psm23\OneDrive\Desktop\Search"` 결과가 `True` 여야 함. 

`Search.py` 파일이 그 폴더 안에 실제로 있어야 함. (없으면 메모장으로 만들어 저장) 해보고 콘솔 마지막 2~3줄만 보여줘. 거기서 바로 이어서 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

**==> picture [218 x 193] intentionally omitted <==**

**==> picture [212 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
이렇게 나오는데 문제가 생기는데?<br>**----- End of picture text -----**<br>


형, 원인은 **작업 폴더가 `C:\Windows\System32` **라서야. `Search.py` 는 잘 실행됐지만, 로그/CSV를 **System32에 쓰려다 권한 거절** 이 난 거야. 

## **바로 해결 (권장)** 

**1.** 폴더로 먼저 이동해서 실행해: 

powershell 코드 복사 `cd "C:\Users\psm23\OneDrive\Desktop\Search" python .\Search.py` 

(안 되면) 

powershell 코드 복사 `Set-Location "C:\Users\psm23\OneDrive\Desktop\Search" python .\Search.py` 

## **영구 해결(코드 패치) — 파일을 스크립트 옆 폴더에만 저장** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

50/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

[REPAIR-16] 

## -[-삭제] 

ini 

코드 복사 

```
OUT_CSV = f"{BASE_NAME}.csv"
LOG_CSV = f"LOG_{BASE_NAME}.csv"
TOP_CSV = f"TOP_{BASE_NAME}.csv"
```

## [+추가] 

python 

코드 복사 

```
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
OUT_CSV = os.path.join(BASE_DIR, f"{BASE_NAME}.csv")
LOG_CSV = os.path.join(BASE_DIR, f"LOG_{BASE_NAME}.csv")
TOP_CSV = os.path.join(BASE_DIR, f"TOP_{BASE_NAME}.csv")
```

## (파일 상단의 `import os, ...` 이미 있으니 `os.path` 만 쓰면 됨) 

이렇게 하면 **어디서 실행해도** 항상 

`C:\Users\psm23\OneDrive\Desktop\Search` 안에 `Search.csv / TOP_Search.csv / LOG_Search.csv` 가 저장돼. 

## 위 두 가지 중 하나로 실행해보고, 

마지막 줄 2~3줄만 다시 보여줘. 바로 다음 단계 이어간다. 

**==> picture [489 x 38] intentionally omitted <==**

**==> picture [289 x 148] intentionally omitted <==**

이렇게 나오고 top 써치는 생성되지도 않아 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

51/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

형, 로그 보면 원인 딱 두 개야. 

## **0) 무엇이 문제?** 

**DuckDuckGo 202** : 봇방어 걸림(HTML 페이지 파싱 시 자주 남). → 라이브러리 우선 사용으로 회피. 

**Naver 401** : 인증 실패. 거의 항상 **ID/SECRET/환경 설정** 문제. 

그리고 **TOP_Search.csv** 는 “결과가 1개도 없으면” 안 만들어지게 되어 있어. (예외로 종료 됨) 

## **1) 빠른 점검 (네이버 401 해결)** 

PowerShell에서 값 확인: 

powershell 

코드 복사 

```
echo %NAVER_CLIENT_ID%
echo %NAVER_CLIENT_SECRET%
```

## 둘 다 **값이 보여야** 함. 

네이버 콘솔에서 아래 3가지를 꼭 확인: 

**1. 사용 API = ‘검색’** 체크 ✅ 

## **2. 비로그인 오픈 API 서비스 환경** 

유형: **웹 서버** 

URL: http://localhost 

## **3. 애플리케이션 사용** 이 **활성** 상태 

그래도 401이면 **SECRET 오타** 일 확률이 큼. 

- → 비공개키 다시 발급 후 `setx` 로 다시 저장하고 창 재시작. 

## **2) 코드 패치 (짧게 · 안전)** 

아래 **차분만** Search.py에 적용해줘. 

## **(A) 저장 경로 고정 + 결과 0개여도 TOP 파일 생성** 

[REPAIR-16] 

diff 

코드 복사 