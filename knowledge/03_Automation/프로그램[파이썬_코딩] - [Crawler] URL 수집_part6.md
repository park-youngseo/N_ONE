개수: 예) 50 

   - 엔진: `google` 또는 `naver` 

   - 설명 수집: `y` 또는 `n` 

**3.** 파일 확인. 

   - 결과: **Search.csv** 

TOP 도메인: **TOP_Search.csv** 

- 로그: **LOG_Search.csv** 

## **7) 자주 나는 오류와 해결** 

- “구글 키 없음” 문구가 보여요. 

- → `GOOGLE_API_KEY` , `GOOGLE_CSE_CX` 다시 넣어. 

- → 터미널 **재시작** 필수. 

- “네이버 키 없음” 이래요. 

- → `NAVER_CLIENT_ID` , `NAVER_CLIENT_SECRET` 확인. 

- → **검색 API** 권한 체크했는지 확인. 

- 값은 있는데 계속 실패해요. 

- → 키에 공백이 섞였나 봐. 

- → 큰따옴표 포함해 다시 입력. 

- → 회사·학교망이면 프록시 이슈일 수 있음. 

결과가 너무 적어요. 

- → 개수를 100으로 올려봐. 

- → 검색어를 더 넓게 써봐. 

## **8) 추천 사용 흐름(빨리 성공)** 

**1.** 키 없이 **duckduckgo** 로 먼저 테스트. 

**2.** 결과가 잘 나오면 **google** 연결. 

**3.** 지역 자료가 중요하면 **naver** 도 추가. 

**4.** 설명 수집은 **필요할 때만** `y` . 

   - (느리면 `n` 으로 꺼두기) 

형, 이대로 하면 연결 끝이야. 

바로 진행해보고, 막히는 지점을 말해줘. 

그 부분만 콕 집어 **즉시 수정** 해줄게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

27/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

폴더 만들기·이동 

C:\Users\psm23\OneDrive\Desktop\Search 여기로 

좋아 형 👍 

이제 **Search.py** 를 넣을 전용 폴더를 만들어보자. 윈도우 기준으로 아주 쉽게 설명할게. 

## 🧭 **1단계 — 폴더 만들기** 

- 1️⃣ **시작 메뉴** → **PowerShell** 실행 

**==> picture [481 x 92] intentionally omitted <==**

**----- Start of picture text -----**<br>
2️⃣ 아래 명령어 그대로 복사해서 붙여넣고 Enter<br>powershell 코드 복사<br>mkdir "C:\Users\psm23\OneDrive\Desktop\Search"<br>**----- End of picture text -----**<br>


- 👉 이걸 하면 **바탕화면의 OneDrive 안에 Search 폴더** 가 생겨. 

## 📂 **2단계 — 폴더로 이동** 

**==> picture [481 x 173] intentionally omitted <==**

**----- Start of picture text -----**<br>
계속 같은 PowerShell 창에서 아래를 입력해.<br>powershell 코드 복사<br>cd /d "C:\Users\psm23\OneDrive\Desktop\Search"<br>✔ 성공하면<br>파란 글씨로 이렇게 바뀌어 있을 거야:<br>powershell 코드 복사<br>**----- End of picture text -----**<br>


**==> picture [243 x 10] intentionally omitted <==**

**----- Start of picture text -----**<br>
PS C:\Users\psm23\OneDrive\Desktop\Search><br>**----- End of picture text -----**<br>


https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

28/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

## ⚙ **3단계 — 이제 여기서 모든 작업 진행** 

이 폴더 안에 

```
Search.py
```

```
Search.csv
```

```
TOP_Search.csv
```

```
LOG_Search.csv
```

이렇게 파일들이 자동으로 생길 거야. 

형, 확인차 물어볼게 👇 

지금 PowerShell 프롬프트가 

```
PS C:\Users\psm23\OneDrive\Desktop\Search>
```

이렇게 되어 있어? 

**==> picture [489 x 39] intentionally omitted <==**

다음 단계 부터 새로만든 프롬포트로 적용시켜서 순서대 로 실행해 

14s 동안 생각함 

확인. **Search** 잠금 유지. 

이제 **순서대로 실행** 시작해. 

## **1) 설치** 

PowerShell(현재 폴더 유지)에서 한 줄씩 실행해. 

sql 코드 복사 

```
python -m pip install -U duckduckgo-search requests beautifulsoup4 lxml google-ap
```

  **2) (선택) 키 저장** 구글·네이버 키가 있으면 넣어. 없으면 건너뛰어도 작동해. nginx 코드 복사 `setx GOOGLE_API_KEY "여기_구글_API_KEY" setx GOOGLE_CSE_CX  "여기_구글_CX"` 

## **2) (선택) 키 저장** 

구글·네이버 키가 있으면 넣어. 없으면 건너뛰어도 작동해. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 29/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
setx NAVER_CLIENT_ID "여기_네이버_ID"
setx NAVER_CLIENT_SECRET "여기_네이버_SECRET"
```

→ 창을 **닫고 다시** 열어. 

## **3) 파일 만들기** 

메모장으로 **Search.py** 새로 만들고, 아래 내용을 **그대로** 붙여넣어 저장해. 

python 코드 복사 

```
# N·ONE Link Collector v2.0 / 2025-11-03 00:00 KST
```

```
# BASE=Search  (OUT: Search.csv, LOG: LOG_Search.csv, TOP: TOP_Search.csv)
```

```
import os, csv, time, math, traceback
from urllib.parse import urlparse
import requests
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry
from bs4 import BeautifulSoup
```

```
try:
```

```
from duckduckgo_search import ddg
    HAS_DDG_LIB = True
except Exception:
    HAS_DDG_LIB = False
```

```
BASE_NAME = "Search"
OUT_CSV = f"{BASE_NAME}.csv"
LOG_CSV = f"LOG_{BASE_NAME}.csv"
TOP_CSV = f"TOP_{BASE_NAME}.csv"
USER_AGENT = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0 Safari/537.36"
DEFAULT_PROVIDER = "duckduckgo"
MAX_SAFE_RESULTS = 500
REQUEST_TIMEOUT = 15
```

## `def now_kst():` 

```
return time.strftime("%Y-%m-%d %H:%M", time.localtime())
```

```
defwrite_log(event, detail=""):
```

```
    header = ["time_kst","event","detail"]
    new_file = not os.path.exists(LOG_CSV)
try:
withopen(LOG_CSV, "a", newline="", encoding="utf-8") as f:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 30/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
            w = csv.writer(f)
```

```
if new_file: w.writerow(header)
            w.writerow([now_kst(), event, detail])
except Exception as e:
print("로그기록실패:", e)
```

```
defsafe_session():
    s = requests.Session()
    s.headers.update({"User-Agent": USER_AGENT})
    retries = Retry(total=3, backoff_factor=0.8, status_forcelist=(500,502,503,504
    s.mount("https://", HTTPAdapter(max_retries=retries))
    s.mount("http://", HTTPAdapter(max_retries=retries))
return s
```

```
defdomain_of(url):
try:
return urlparse(url).netloc.lower().replace("www.", "")
except:
return""
defclean_text(s):
return" ".join((s or"").split()).strip()
# ---- 엔진들 ----
defsearch_duckduckgo_html(query, max_results=50, session=None):
    s = session or safe_session()
    results, base, start = [], "https://html.duckduckgo.com/html/", 0
whilelen(results) < max_results:
try:
            r = s.post(base, data={"q": query, "s": start}, timeout=REQUEST_TIMEOU
except Exception as e:
            write_log("ddg_html_error", str(e)); break
if r.status_code != 200:
            write_log("ddg_html_status", f"{r.status_code}"); break
        soup = BeautifulSoup(r.text, "html.parser")
        anchors = soup.select("a.result__a")