
```
# ── 사이트리서치(전체크롤링 → 고민추출/클러스터) ─────────────────────
QUESTION_HINTS = [
"어떻게", "왜", "무엇", "뭐가", "어떤", "언제", "어디", "얼마", "방법", "팁",
"문제", "고민", "어려움", "힘들", "실수", "주의", "해결", "비용", "기간", "준비"
]
defis_same_domain(seed, url):
return domain_of(seed) == domain_of(url)
```

```
defextract_questions(text):
# 문장분할후고민/질문패턴만뽑기
    sents = re.split(r"[?\!\.\n]+", text)
    out=[]
for s in sents:
        st = clean_text(s)
ifnot st: continue
ifany(h in st for h in QUESTION_HINTS) or"?"in s:
if5 <= len(st) <= 180:
                out.append(st)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 155/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
returnlist(dict.fromkeys(out))  # 중복제거
```

```
deffetch_main_text(url, s):
try:
        r = s.get(url, timeout=8)
if r.status_code != 200or"text/html"notin r.headers.get("Content-Type
return""
        doc = Document(r.text)
        body_html = doc.summary(html_partial=True)
        body = BeautifulSoup(body_html, "lxml").get_text(" ")
        title = clean_text(doc.short_title() or"")
        soup = BeautifulSoup(r.text, "lxml")
        meta = soup.find("meta", attrs={"name":"description"}) or soup.find("meta
        desc = clean_text(meta.get("content","")) if meta else""
        combo = " ".join([title, desc, body])
return clean_text(combo)
except:
return""
defcrawl_site(seed_url, max_pages=200, max_depth=2, delay=0.5):
    s = session()
    seen=set([seed_url]); q=[(seed_url,0)]
    texts=[]; urls=[]
    base_domain = domain_of(seed_url)
while q andlen(urls) < max_pages:
        u, d = q.pop(0)
        txt = fetch_main_text(u, s)
if txt:
            urls.append(u); texts.append(txt)
if d >= max_depth: continue
try:
            r = s.get(u, timeout=8)
            soup = BeautifulSoup(r.text, "lxml")
for a in soup.find_all("a", href=True):
                v = urljoin(u, a["href"])
if v.startswith("mailto:") or v.startswith("javascript:"): contin
ifnot is_same_domain(seed_url, v): continue
if v notin seen:
                    seen.add(v); q.append((v,d+1))
            time.sleep(delay)
except:
continue
return urls, texts
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 156/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
defcluster_questions(domain, qpairs, k=None):
# qpairs: [(url, question sentence), ...]
    texts = [q for _, q in qpairs]
ifnot texts: return []
# 한글용 char n-gram 벡터화
    vect = TfidfVectorizer(analyzer="char", ngram_range=(2,4), min_df=2)
    X = vect.fit_transform(texts)
if k isNone:
        k = max(2, min(10, len(texts)//20or2))  # 데이터양에따른자동 k
    km = KMeans(n_clusters=k, n_init=10, random_state=42)
    labels = km.fit_predict(X)
# 클러스터별대표키워드
    terms = vect.get_feature_names_out()
    centers = km.cluster_centers_
    top_terms = []
for ci inrange(k):
        idx = centers[ci].argsort()[::-1][:6]
        top_terms.append("".join(terms[i] for i in idx)[:24])
# 결과정리
    out=[]
for ci inrange(k):
        members = [(u,q) for (u,q),lab inzip(qpairs, labels) if lab==ci]
        ex = " | ".join([q for _,q in members[:3]])
        out.append({"cluster": ci, "theme": top_terms[ci], "count": len(members),
    out.sort(key=lambda x: -x["count"])
return out, labels
```

```
defsite_research_flow(keyword, merged_rows, tops):
# 1) 상위도메인후보보여주기
    cand = [t["domain"] for t in tops[:10] if t["domain"]]
ifnot cand:
print("\n(리서치) 도메인후보가없습니다."); return
print("\n[리서치대상도메인선택]")
for i, d inenumerate(cand, start=1):
print(f"{i}. {d}")
try:
        idx = int(input("> 번호선택: ").strip() or"1") - 1
except: idx = 0
    idx = max(0, min(len(cand)-1, idx))
    dom = cand[idx]
# 2) 시드 URL (해당도메인링크중하나) 고르기
    seeds = [r["url"] for r in merged_rows if r.get("domain")==dom][:5]
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 157/169 

25. 11. 30. 오후 1:05 

 

후 1:05 1:05 프로그램[파이썬/코딩] - [Crawler] URL 수집 `seed = seeds[0] if seeds else f"https://{dom}/" print(f"\n[크롤 시작] {dom} (seed: {seed})")` _`# 3)`_ `크롤링 urls, texts = crawl_site(seed, max_pages=200, max_depth=2, delay=0.6) print(f" - 수집 페이지: {len(urls)}")` _`# 4)`_ `고민` _`/`_ `질문 문장 추출 qpairs=[] for u, t in zip(urls, texts): for q in extract_questions(t): qpairs.append((u,q)) print(f" - 고민/질문 문장: {len(qpairs)}")` _`# 5)`_ `클러스터링 report_rows=[] if qpairs: clusters, labels = cluster_questions(dom, qpairs)` _`# CSV`_ `저장` _`(`_ `클러스터 상세` _`)`_ `det_path = os.path.join(OUT_DIR, f"리서치_{dom}.csv") with open(det_path, "w", encoding="utf-8-sig", newline="") as f: w = csv.writer(f); w.writerow(["cluster","question","url"]) for (u,q), lab in zip(qpairs, labels): w.writerow([lab, q, u])` _`# SUMMARY`_ `저장 sum_path = os.path.join(OUT_DIR, f"SUMMARY_{dom}.txt") with open(sum_path, "w", encoding="utf-8") as f: f.write(f"[사이트 리서치 요약] {dom}\n수집페이지: {len(urls)} / 고민문장 for c in clusters: f.write(f"- C{c['cluster']:02d} [{c['theme']}] x{c['count']}\n  예 print(f"\n` ✅ `리서치 완료 → {det_path}\n   요약 → {sum_path}") else: print("\n(리서치) 고민 문장이 충분하지 않습니다.")` _`# ──`_ `메인` _`────────────────────────────────────────────────`_ `def main(): print("검색어를 입력하세요:") q = input("> ").strip() if not q: print("검색어 없음. 종료."); return q_clean = safe_name(q) try: print("수집 개수(기본 50):") n = int(input("> ").strip() or "50") n = max(1, min(MAX_RESULTS_CAP, n)) except: n = 50` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 158/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
print("엔진선택: serper / google / naver / duckduckgo / mix")
    prov = (input("> ").strip().lower() or"serper")
```

```
print("설명(meta)도모을까요? (y/N)")
    fetch_desc = (input("> ").strip().lower() == "y")
```

```
try:
        parts=[]
if prov == "serper":
            parts.append(engine_to_csv(q_clean, "serper", search_serper(q, n), fet
elif prov == "google":
            parts.append(engine_to_csv(q_clean, "google", search_google(q, n), fet
elif prov == "naver":
            parts.append(engine_to_csv(q_clean, "naver",  search_naver(q, n), fet
elif prov == "duckduckgo":
            parts.append(engine_to_csv(q_clean, "duckduckgo", search_ddg(q, n), fe
elif prov == "mix":
```

```
            parts.append(engine_to_csv(q_clean, "serper", search_serper(q, n), fet
            parts.append(engine_to_csv(q_clean, "google", search_google(q, n), fet
            parts.append(engine_to_csv(q_clean, "naver",  search_naver(q, n), fet
else:
```

```
"
print(지원하지않는엔진. 종료."); return
```

```
        merged_rows, tops = merge_and_save(q_clean, q, parts)
```

```
# ── 추가: 사이트리서치모드 ──
print("\n사이트리서치모드로넘어갈까요? (y/N)")
if (input("> ").strip().lower() == "y"):
```

```
            site_research_flow(q_clean, merged_rows, tops)
```

`print(f"\n` ✅ `완료: out\\{q_clean}_{prov if prov!='mix' else 'serper'}.csv print(f"   병합본: out\\{q_clean}_all.csv") print(f"   도메인TOP: out\\TOP_{q_clean}_all.csv") print(f"   로그: out\\LOG_{q_clean}_all.csv")` 

```
except Exception as e:
print("오류:", str(e))
        write_log(q_clean, "예외: " + traceback.format_exc())
print("자세한내용은 LOG 파일확인.")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

159/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

```
if __name__ == "__main__":
```

```
    main()
```

## **결과물이 어떻게 생겨?** 

```
out/리서치_filmmakers.co.kr.csv
```

열: `cluster, question, url` 

같은 고민끼리 같은 cluster 번호로 묶여 있어. 

```
out/SUMMARY_filmmakers.co.kr.txt
```

“주요 고민 주제 x 건수 + 예시 문장” 한눈 요약. 

## 원하면 **클러스터 개수** , **크롤 깊이/페이지수** , **질문 단어 사전** 도 커스텀해줄게. 

일단 이 버전으로 한 번 돌려보고, `SUMMARY_*.txt` 내용 보여주면 거기서 더 정교하게 다 듬자! 

**==> picture [489 x 39] intentionally omitted <==**

초등학생도 이해가능하게 설명해줘 

형, 아주 쉽게 설명할게. 

## **이 프로그램은?** 

검색어로 링크를 모아. 사이트 글을 읽어줘. 사람들 고민을 찾아줘. 

## **어떻게 쓰면 돼?** 

**1.** 검색어를 적어. 

**2.** 엔진을 고른다. 

**3.** 설명을 모을지 고른다. 

**4.** 끝나면 “리서치: y”. 

## **뭘 해 주는데?** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/69078421-4cf4-8321-83dd-3aad9350d5b4 

160/169 

25. 11. 30. 오후 1:05 

프로그램[파이썬/코딩] - [Crawler] URL 수집 

 

사이트 여러 쪽을 읽어. 질문 같은 문장을 찾지. 비슷한 것끼리 묶어줘. 주제별로 정리해 줘. 

## **어디서 보면 돼?** 

out 폴더를 열어봐. 

`검색어_all.csv` 가 있어. `TOP_검색어_all.csv` 도 있어. `리서치_도메인.csv` 도 생겨. `SUMMARY_도메인.txt` 를 읽어. 한눈 요약이 들어있어. 

## **무엇에 쓸까?** 

사람들 고민을 알지. 필요한 글만 고른다. 아이디어를 만든다. 할 일을 정리한다. 

## **잘 쓰는 꿀팁** 

검색어를 구체적으로! 예) 연기 오디션 모집 예) 단역 캐스팅 공고 설명은 느릴 땐 끄기. 필요할 때만 켜기. 

## **기억하면 좋아** 

링크가 전부는 아냐. 중요한 건 “의미야”. 묶인 주제를 먼저 봐. 그다음 링크를 보자. 

궁금한 부분 있으면 딱 한 가지만 말해줘. 바로 더 쉽게 풀게! 