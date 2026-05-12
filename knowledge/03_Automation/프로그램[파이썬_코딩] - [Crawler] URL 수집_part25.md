## 즉 — **Search.py 코드가 멈춘 이유는 네이버 API 자체 오류가 아니라,** 

아마도 이전 버전 실행 중 키 인식이 꼬이거나 

PowerShell 세션이 오래 열려 있어서 환경 변수가 즉시 적용되지 않았던 거야. 

## 🔧 **해결 절차 (이번 한 번만 하면 끝)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 123/169 

 

25. 11. 30. 오후 1:05 프로그램[파이썬/코딩] - [Crawler] URL 수집 1️⃣ **PowerShell 전부 닫기 (창 완전 종료)** 2️⃣ 새 창 열기 3️⃣ 경로 이동 powershell 코드 복사 `cd "C:\Users\psm23\OneDrive\Desktop\Search"` 4️⃣ 다시 실행 powershell 코드 복사 `python .\Search.py` 5️⃣ 엔진 선택 단계에서 이번엔 이렇게 입력 👇 makefile 코드 복사 `엔진: naver` 6️⃣ “설명(meta)”은 N으로 두고 scss 코드 복사 `설명(meta)도 모을까요? (y/N): n` 

## 🔍 **결과 확인** 

## 정상이라면 

csharp 코드 복사 ✅ `완료: 10개 저장 → out\연기오디션_naver.csv` 

## 이 문장이 뜰 거야. 

## 형이 이대로 해보고 

- ✅ `완료` 문장 

- `out/` 폴더 안의 파일 목록 

- 두 개를 캡처해서 보여주면 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 124/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

이제 “자동 키워드 요약 + 도메인별 분류 모듈” 바로 붙일게. 

**==> picture [489 x 38] intentionally omitted <==**

**==> picture [97 x 96] intentionally omitted <==**

**==> picture [97 x 96] intentionally omitted <==**

좋아 형 👏 완벽하게 돌아갔어. 이건 지금까지 했던 세팅 중 가장 깨끗한 실행 결과야. 

## ✅ **결과 분석** 

## **항목** 

## **상태** 

`드라마오디션_naver.csv` 본 검색 결과 (50개 링크) `TOP_드라마오디션_naver.csv` 도메인별 상위 5개 자동 정리 `LOG_드라마오디션_naver.csv` 실행 로그 (시간·엔진·상태 기록) 

저장 위치 

```
C:\Users\psm23\OneDrive\Desktop\Search\out
```

## 💡 **지금 구조의 장점** 

- ✅ 검색어별 자동 폴더 정리됨 ( `out/드라마오디션_naver.csv` ) 

- ✅ 포털별 이름 자동 포함 ( `_naver` , `_google` , `_serper` ) 

- ✅ 로그 자동 기록 (언제, 어떤 엔진, 몇 건) 

- ✅ 네이버 API 정상 통신 (401/403 없음) 

- ✅ 검색 결과 누락 0건 

## 🔧 **다음 단계 제안** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 125/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## 이제 완전히 정상화됐으니까, 

형이 원하면 **“Serper + Naver + Google 통합 병합 버전(v2.7)”**으로 업그레이드할 수 있어. 

이 버전은 👇 

csharp 코드 복사 

```
out/
 ├── 드라마오디션_serper.csv
 ├── 드라마오디션_naver.csv
 ├── 드라마오디션_google.csv
 └── TOP_드라마오디션_all.csv  ← 3개엔진통합랭킹
```

형이 “그래, 병합 버전으로 가자”라고 하면 

지금 이 상태 그대로 쓰면서, **엔진 통합+도메인 통계 모듈** 까지 자동으로 추가해줄게. 할까? 

**==> picture [489 x 38] intentionally omitted <==**

병합 버전으로 

좋아 형 👍 

이제 `v2.7 — Search_병합버전` 으로 업그레이드 들어간다. 이 버전은 세 가지 검색 엔진을 **한 번에 병합** 해서 결과를 만들어줘. 

## 🔥 **Search v2.7 (병합버전) 기능 요약** 

|**항목**|**설명**||
|---|---|---|
|검색 엔진|`naver` +<br>`serper` (=Google 기반)|+<br>`google`|
|저장 구조|`out/검색어_엔진.csv` 각각 생성 +|<br>`TOP_검색어_all.csv` 병합|
|로그 구조|실행 로그(<br>`LOG_검색어_all.csv` ) 자동 생성||
|검색 개수|엔진당 최대 50개 (합계 최대 150개)||
|필터|중복 URL 자동 제거||



https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 126/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

**항목 설명** 

병합 기준 동일 도메인 기준 그룹화 (예: naver.com, daum.net 등) 결과 전체 링크 수 + 도메인별 순위 TOP10 

## ⚙ **실행 흐름 (자동)** 

python 코드 복사 

```
python Search.py
 └▶검색어입력
 └▶각엔진순차호출 (serper → google → naver)
 └▶ out 폴더에저장
 └▶도메인별합산후 TOP_검색어_all.csv 생성
 └▶ LOG_검색어_all.csv 기록
```

## 📂 **결과 예시** 

python 코드 복사 

```
out/
```

```
 ├─ 드라마오디션_serper.csv
 ├─ 드라마오디션_google.csv
 ├─ 드라마오디션_naver.csv
 ├─ TOP_드라마오디션_all.csv
 └─ LOG_드라마오디션_all.csv
```

## 🚀 **설치 준비 (Serper 전용 키 1회 등록)** 

1️⃣ https://serper.dev 접속 

2️⃣ GitHub 계정으로 로그인 

3️⃣ Dashboard → API Key 복사 

- 4️⃣ PowerShell에서 아래 입력 👇 

powershell 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 127/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
setx SERPER_API_KEY "여기에_복사한_API키"
```

## 5️⃣ PowerShell 닫고 새로 열기 

powershell 

코드 복사 

```
echo %SERPER_API_KEY%
```

값이 나오면 준비 완료. 

## 🧠 **병합버전 코드 구성 요약** 

python 

코드 복사 

```
# Search_v2.7.py
```

```
# ver 2.7 (Merged Naver + Serper + Google)
```

```
# 자동병합 + 로그 + 도메인 TOP10
```

```
엔진리스트 = ["serper", "google", "naver"]
검색결과 = []
for엔진in엔진리스트:
    ...
```

```
중복제거후병합
도메인카운트
TOP10 저장
```

형, 이렇게 설명대로 ** `Search_v2.7.py` **를 만들어줄게. 

그 전에 확인 한 가지만 👇 

지금처럼 결과를 `out` 폴더에 계속 누적 저장(날짜 유지)할까, 

아니면 **매 실행 시 out 폴더 초기화(이전 파일 자동삭제)** 로 갈까? 

🔘 1. 누적 저장 (추천, 백업용) 

⚪ 2. 실행 시 out 폴더 자동 초기화 

## 번호만 말해줘. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

128/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

1 

완벽 👌 형은 역시 백업을 중요하게 생각하네 — 그게 정답이야. 그럼 **v2.7 누적저장형(병합버전)** 으로 바로 구성할게. 

## ✅ **Search_v2.7.py — 병합버전 (누적저장형)** 

## 📦 기능요약 

- `naver + serper + google` 3개 엔진 자동 수집 

- `out/` 폴더 유지 (이전 결과 덮어쓰지 않음) 

- 도메인별 TOP 10 + LOG 자동 생성 

- 엔진별 중복 제거 및 병합 

## 🧩 **코드 (이대로 저장)** 

python 

코드 복사 

```
# Search_v2.7.py
```

- _`# ver 2.7 — Merged Naver + Serper + Google`_ 

- _`#`_ `누적저장형` _`/ out`_ `폴더 유지` 

- _`#`_ `작성` _`: N·ONE Hybrid System v1.7 (KST`_ `고정` _`)`_ 

```
import os, csv, time, requests
from datetime import datetime
```

```
from urllib.parse import urlparse
```