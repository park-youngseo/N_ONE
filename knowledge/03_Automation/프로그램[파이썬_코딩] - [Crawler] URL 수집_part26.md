
```
# ===== 폴더준비 =====
```

```
base_dir = os.path.dirname(os.path.abspath(__file__))
out_dir = os.path.join(base_dir, "out")
```

```
os.makedirs(out_dir, exist_ok=True)
```

```
# ===== 네이버 API =====
```

```
NAVER_ID = os.getenv("NAVER_CLIENT_ID")
NAVER_SECRET = os.getenv("NAVER_CLIENT_SECRET")
```

```
# ===== Serper API =====
```

```
SERPER_KEY = os.getenv("SERPER_API_KEY")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 129/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
# ===== 구글 (비공식 DuckDuckGo 방식) =====
```

```
defgoogle_search(query, n=30):
from duckduckgo_search import DDGS
    results = []
```

```
with DDGS() as ddgs:
for r in ddgs.text(query, max_results=n):
            results.append((r["title"], r["href"], r["body"]))
return results
```

```
# ===== 네이버검색 =====
```

```
defnaver_search(query, n=30):
```

```
    headers = {"X-Naver-Client-Id": NAVER_ID, "X-Naver-Client-Secret": NAVER_SECR
    url = f"https://openapi.naver.com/v1/search/webkr.json?query={query}&display=
    res = requests.get(url, headers=headers)
```

```
    data = res.json()
    items = []
for i in data.get("items", []):
        items.append((i["title"], i["link"], i["description"]))
return items
```

```
# ===== Serper 검색 (Google 기반 API) =====
```

```
defserper_search(query, n=30):
```

```
    headers = {"X-API-KEY": SERPER_KEY, "Content-Type": "application/json"}
```

```
    payload = {"q": query, "num": n}
```

```
    r = requests.post("https://google.serper.dev/search", headers=headers, json=pa
    data = r.json()
    results = []
for i in data.get("organic", []):
        results.append((i["title"], i["link"], i.get("snippet", "")))
return results
```

```
# ===== 도메인추출 =====
defget_domain(url):
try:
return urlparse(url).netloc.replace("www.", "")
except:
return"unknown"
```

```
# ===== CSV 저장 =====
```

```
defsave_csv(path, rows):
withopen(path, "w", encoding="utf-8-sig", newline="") as f:
        writer = csv.writer(f)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 130/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
        writer.writerow(["title", "url", "desc"])
        writer.writerows(rows)
```

```
# ===== 로그저장 =====
defwrite_log(keyword, msg):
    log_path = os.path.join(out_dir, f"LOG_{keyword}_all.csv")
withopen(log_path, "a", encoding="utf-8-sig", newline="") as f:
        w = csv.writer(f)
        w.writerow([datetime.now().strftime("%Y-%m-%d %H:%M"), msg])
```

```
# ===== 병합로직 =====
defmerged_search(keyword, n=30):
    all_results = []
# Serper
try:
        serper = serper_search(keyword, n)
        save_csv(os.path.join(out_dir, f"{keyword}_serper.csv"), serper)
        all_results.extend(serper)
        write_log(keyword, "Serper 성공")
except Exception as e:
        write_log(keyword, f"Serper 실패: {e}")
# Google
try:
        google = google_search(keyword, n)
        save_csv(os.path.join(out_dir, f"{keyword}_google.csv"), google)
        all_results.extend(google)
        write_log(keyword, "Google 성공")
except Exception as e:
        write_log(keyword, f"Google 실패: {e}")
# Naver
try:
        naver = naver_search(keyword, n)
        save_csv(os.path.join(out_dir, f"{keyword}_naver.csv"), naver)
        all_results.extend(naver)
        write_log(keyword, "Naver 성공")
except Exception as e:
        write_log(keyword, f"Naver 실패: {e}")
# 중복제거
    unique = {}
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

131/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
for t, u, d in all_results:
if u notin unique:
            unique[u] = (t, d)
```

```
    merged = [(v[0], k, v[1]) for k, v in unique.items()]
    save_csv(os.path.join(out_dir, f"MERGED_{keyword}_all.csv"), merged)
```

```
# 도메인별 TOP 10
    domain_count = {}
for _, url, _ in merged:
        dom = get_domain(url)
        domain_count[dom] = domain_count.get(dom, 0) + 1
```

```
    top_domains = sorted(domain_count.items(), key=lambda x: x[1], reverse=True)[
    top_path = os.path.join(out_dir, f"TOP_{keyword}_all.csv")
withopen(top_path, "w", encoding="utf-8-sig", newline="") as f:
        w = csv.writer(f)
        w.writerow(["domain", "count"])
        w.writerows(top_domains)
```

`write_log(keyword, f"병합 완료: {len(merged)}개 결과") print(f"\n` ✅ `병합 완료 ({len(merged)}개 결과) → {top_path}")` 

_`# =====`_ `실행` _`=====`_ `if __name__ == "__main__": keyword = input("검색어를 입력하세요: ").strip() count = input("수집 개수 (기본 30): ").strip() count = int(count) if count else 30 print(f"\n` 🔍 `검색 시작: {keyword}") merged_search(keyword, count)` 

## ⚙ **실행 방법** 

## 1️⃣ 폴더로 이동 

powershell 코드 복사 

```
cd"C:\Users\psm23\OneDrive\Desktop\Search"
```

## 2️⃣ 실행 

powershell 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

132/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
python .\Search_v2.7.py
```

## 형이 이대로 실행하면 결과 폴더에 

`MERGED_검색어_all.csv` , `TOP_검색어_all.csv` , `LOG_검색어_all.csv` 

이 세 개가 추가될 거야. 

## ✅ **확인 요청:** 

이대로 바로 적용해줄까, 

아니면 `meta(설명문)` 도 같이 수집하도록 옵션을 추가해줄까? (둘 다 가능해 — meta는 결과가 2~3배 풍부해져.) 

**==> picture [489 x 39] intentionally omitted <==**

야 파일이름 바꼈자나 그냥 통일해 짜증나니까 

## ㅇㅋ 형 😤 완전 이해했어. 

파일명 통일해서 깔끔하게 맞출게 — 진짜 혼란 없게 딱 정리. 

## 🔧 **수정방향 (형 요청 반영)** 

## 지금처럼 뒤죽박죽인 

python 코드 복사 `MERGED_드라마오디션_all.csv TOP_드라마오디션_all.csv LOG_드라마오디션_all.csv` 

## 이런 거 전부 **형의 기존 패턴으로 통일** 한다 👇 

## ✅ **새 규칙 (엔오원 표준 파일명 통일)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 133/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

**역할 새 파일명 형식 예시** 본 결과(병합 전체) `드라마오디션_all.csv` ✅ 로그 `LOG_드라마오디션_all.csv` ✅ 도메인 상위 `TOP_드라마오디션_all.csv` ✅ 개별 엔진 `드라마오디션_google.csv` , `드라마오디션` ✅ 그대로 유지 `_naver.csv` , `드라마오디션_serper.csv` 

## 🧩 **통일 버전 코드 (Search_v2.7_fixed.py)** 

python 

코드 복사 

```
# Search_v2.7_fixed.py
```

```
# 파일명통일형 (N·ONE 표준)
```
