print("엔진선택: duckduckgo / google / naver (기본 duckduckgo)")
    prov = (input("> ").strip().lower() or"duckduckgo")
if prov notin ("duckduckgo","google","naver"):
print("지원하지않는엔진. duckduckgo 사용"); prov = "duckduckgo"
print("설명(meta)도모을까요? (y/N)")
    fetch_desc = (input("> ").strip().lower() == "y")
    write_log("run_request", f"q={q},n={n},prov={prov},desc={fetch_desc}")
try:
        rows = collect_links(q, n, prov, fetch_desc)
# 부분결과라도저장
        save_results(rows)
        save_top(rows)
ifnot rows:
"
print(결과없음. 검색어/엔진을바꿔보세요.")
            write_log("no_results", q)
return
print(f"완료: {len(rows)}개저장 → {OUT_CSV}, {TOP_CSV}")
        write_log("run_done", f"saved={len(rows)}")
except Exception as e:
print("오류:", str(e))
print("자세한내용은 out/LOG_Search.csv 확인")
        write_log("run_exception", traceback.format_exc())
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

63/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
if __name__ == "__main__":
    main()
```

실행: 

powershell 

코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
python .\Search.py
```

네이버 401 계속 뜨면 **SECRET 재발급 → setx 재입력 → PowerShell 재시작** 후 `naver` 로 다시 테스트해줘. 

**==> picture [489 x 39] intentionally omitted <==**

확인, 형. 

방금 준 **Search.py v2.2** 코드가 “out 폴더 자동 저장 + 구글/네이버/덕덕고 + TOP 생 성”까지 전부 포함이야. 

## **지금 할 일 (3줄)** 

**1.** 메모장 열고 방금 코드 **그대로** 붙여넣어 저장 → 

   - `C:\Users\psm23\OneDrive\Desktop\Search\Search.py` 

**2.** PowerShell: 

bash 

코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
python .\Search.py
```

**3.** 입력 예시 

   - 검색어: `연기 커뮤니티` 

   - 개수: `30` 

   - 엔진: `duckduckgo` (또는 `google` ) 

   - 설명 수집: `n` 

## **결과 위치** 

```
out/Search.csv
```

```
out/TOP_Search.csv
```

```
out/LOG_Search.csv
```

네이버는 401이면 **Secret 재발급 → setx 재입력 → PowerShell 재시작** 후 `naver` 로 재시 도하면 돼. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

64/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

돌려보고 `out/LOG_Search.csv` 마지막 5줄만 보내주면, 거기서 이어서 바로 손볼게. 

**==> picture [489 x 39] intentionally omitted <==**

형, 준비 끝. 

방금 준 **Search.py v2.2** 를 그대로 붙여 넣어 저장하면 돼. 

- (경로: `C:\Users\psm23\OneDrive\Desktop\Search\Search.py` ) 

## **실행 3줄** 

bash 

코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
```

```
python .\Search.py
```

입력 예시: `연기 커뮤니티` / `30` / `duckduckgo` / `n` 

## **파일 위치** 

- 결과: `out/Search.csv` 

- 도메인집계: `out/TOP_Search.csv` 

- 로그: `out/LOG_Search.csv` 

## **빠른 체크** 

- 구글 쓸 땐 `GOOGLE_API_KEY` , `GOOGLE_CSE_CX` 들어있는지 `echo` 로 확인. 

- 네이버 401 나오면 **SECRET 재발급 → setx → PowerShell 재시작** . 

돌려보고 `out/LOG_Search.csv` 마지막 5줄만 보내줘. 바로 이어서 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

그리고 한가지더 log,top 서치 파일 끝에 사용된 포털 이 름을 붙여줘 가령 구글포털을 이용하면 

TOP_Search_google 이런식으로 네이버를 이용하면 TOP_Search_naver 이런식으로 

그리고 한가지더 log,top 서치 파일 끝에 사용된 포털 이 름을 붙여줘 가령 구글포털을 이용하면 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

65/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

TOP_Search_google 이런식으로 네이버를 이용하면 TOP_Search_naver 이런식으로 

좋아 형 👍 완전 깔끔한 요청이야. 

**Search.py** 에 살짝만 수정해서 이렇게 저장되게 만들면 돼: 

```
out/Search_google.csv
```

```
out/TOP_Search_google.csv
out/LOG_Search_google.csv
```

(엔진 바뀌면 자동으로 이름 변경됨) 

## 🔧 **수정할 부분만 복사해서 덮어쓰기** 

`Search.py` 에서 아래 줄 찾아서 바꿔줘. 

## **① 기존** 

python 코드 복사 

```
OUT_CSV = os.path.join(OUT_DIR, f"{BASE_NAME}.csv")
LOG_CSV = os.path.join(OUT_DIR, f"LOG_{BASE_NAME}.csv")
TOP_CSV = os.path.join(OUT_DIR, f"TOP_{BASE_NAME}.csv")
```

## **② 수정** 

python 코드 복사 

```
# 엔진이름이붙은버전 (예: _google, _naver)
```

```
prov_suffix = ""
try:
    prov_suffix = "_" + prov
except NameError:
    prov_suffix = ""
```

```
OUT_CSV = os.path.join(OUT_DIR, f"{BASE_NAME}{prov_suffix}.csv")
LOG_CSV = os.path.join(OUT_DIR, f"LOG_{BASE_NAME}{prov_suffix}.csv")
TOP_CSV = os.path.join(OUT_DIR, f"TOP_{BASE_NAME}{prov_suffix}.csv")
```

## 이제 실행하면 자동으로 이렇게 저장돼 👇 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

66/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

코드 복사 

csharp 

 

```
out/Search_google.csv
out/LOG_Search_google.csv
out/TOP_Search_google.csv
```

다른 엔진 쓰면 이름만 바뀜: 

csharp `out/Search_naver.csv out/TOP_Search_naver.csv` 

코드 복사 

원하면 내가 이 수정 포함한 **v2.3 전체 코드** 완성본 통째로 만들어줄까? (붙여넣기만 하면 되게) 

**==> picture [489 x 39] intentionally omitted <==**

만들어줘 

## 좋아 형 👍 

요청한 대로 **포털 이름 자동 붙는 버전 (v2.3)** 완성본을 아래에 줬어. 

이걸 통째로 복사해서 `C:\Users\psm23\OneDrive\Desktop\Search\Search.py` 에 붙여넣으 면 돼. 

python 코드 복사 

```
# N·ONE Link Collector v2.3 / 2025-11-04 15:10 KST
```

```
# BASE=Search  (OUT: out/Search_[엔진].csv / LOG_Search_[엔진].csv / TOP_Search_[엔
```

```
import os, csv, time, math, traceback
from urllib.parse import urlparse
import requests
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
from bs4 import BeautifulSoup
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 67/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# DuckDuckGo 라이브러리(있으면사용)
```

```
try:
```
