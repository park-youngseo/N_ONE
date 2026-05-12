
## **2) 멀티 엔진만 강제로 테스트(네이버/다음만)** 

통합본 폴더에서: 

nginx 

코드 복사 

```
python scripts\autocomplete_harvester_multi.py . naver,daum
```

끝나면 `outputs\suggestions_all.csv` 열어서 **engine=naver/daum** 이 찍히는지 확인. 

## **3) 통합 실행도 멀티로 돌리기** 

통합본이면 **STEP1_AUTOCOMPLETE.cmd** 가 이미 멀티(google,youtube,naver,daum)로 호출한다. 

혹시 예전 러너를 쓰고 있으면, 아래 한 줄로 바로 멀티 실행: 

nginx 

코드 복사 

```
python scripts\autocomplete_harvester_multi.py . google,youtube,naver,daum
```

## **4) 그래도 naver/daum이 비면?** 

`outputs\suggestions_all.csv` 에 **engine=error:naver / error:daum** 같은 행이 있는 지 확인. 

있으면 잠깐 쉬었다(요청 속도↓) 다시 시도: 

nginx 코드 복사 

```
python scripts\autocomplete_harvester_multi.py . naver,daum
```

(엔진을 naver,daum만 두고 수집 → 성공 확인 후 전체 멀티로 확장) 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

28/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

## **5) 최종 체크 포인트** 

- `outputs\suggestions_all.csv` 의 **engine 컬럼** 에 최소 하나 이상 `naver` 또는 `daum` 

- 이 보여야 정상. 

- 이후 자막 단계는 유튜브에서만 가져오니( **transcripts.csv** 는 당연히 youtube 기반), **국내 의도 반영은 suggestions_all.csv** 에서 이미 확보됨. 

필요하면, 내가 **요청간 대기시간** (rate limit 완화)과 **엔진 우선순위** 를 조정한 CMD 한 장으 로 다시 묶어서 줄게. 

**==> picture [489 x 39] intentionally omitted <==**

## 여전히 씨발 개새끼야 똑같네 

14s 동안 생각함 

형, 오케이. **네이버/다음이 계속 안 찍히는 문제** , 내가 강제로 뚫는 **“엄격(ND-STRICT) 하 베스터”** 패치를 만들었다. 

(여러 엔드포인트·파라미터 조합 시도 + Referer/헤더 세팅 + 결과를 별도 파일로 분리 저 장) 

**다운로드:** n1_public_needs_nd_strict_patch.zip 

## **어떻게 쓰냐 (진짜 간단)** 

## **1.** 압축을 **기존 키트 폴더 안** 에 풀어. (예: 

   - `C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\` ) 

   - → 풀면 `scripts\autocomplete_nd_strict.py` 와 ** `RUN_ND_STRICT.cmd` **가 생김 

**2.** 같은 폴더에서 `RUN_ND_STRICT.cmd` 더블클릭 

**outputs\suggestions_nd_only.csv** 가 생성(엔진 컬럼이 **naver/daum** 인 행만) 

- 기존 `outputs\suggestions_all.csv` 가 있으면 **자동 병합** 까지 해둠 

## **확인 포인트** 

- `outputs\suggestions_nd_only.csv` 열었을 때 **engine=naver/daum** 이 보여야 정상. 

- 이후에는 평소대로 `RUN_ME.cmd` (또는 STEP2~STEP3)로 진행하면 됨. (유튜브 자막 단계는 당연히 유튜브만 쓰므로 정상) 

## **왜 이게 먹히냐** 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

29/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

**네이버/다음** 자동완성은 응답 구조가 케이스마다 다르고, 가끔 **Referer가 없으면** 빈 결과가 떨어진다. 

- 이 패치는 

   - 네이버: `st=111/100/1` 등 **여러 조합을 순차 시도** 

   - 다음: 파라미터 **두 가지 모드** 시도 

   - **Referer 헤더** 강제 설정 

결과 파싱을 **여러 형태** 로 대응 

으로 **빈 결과** 를 최대한 줄였다. 

만약 이걸 돌려도 `suggestions_nd_only.csv` 가 비어 있으면, 그때는 **네트워크(회사/공유** · **기 DNS, 방화벽)** 이 막는 경우가 많다. 그때는 내가 **요청 간 대기시간↑ 프록시/ DNS 우 회 없이 순정 요청** 버전으로 다시 줄게. 

**==> picture [489 x 40] intentionally omitted <==**

**==> picture [289 x 39] intentionally omitted <==**

이렇게 나온다 

좋아, 스샷 보니까 **Daum 쪽이 404** 로 떨어졌네. 그 엔드포인트가 지역/환경에 따라 막히 는 경우가 있어서, **대체 호스트(** `suggest-bar.daum.net` **)로도 시도** 하는 **FIX2 패치** 를 만들 었다. 실행하면 **naver/daum 결과를 별도 파일로 뽑고** , 기존 `suggestions_all.csv` 까지 자동 병합한다. 

**다운로드:** n1_public_needs_nd_fix2.zip 

## **쓰는 법 (복붙 필요 없음, 더블클릭만)** 

## **1.** 압축 풀기 

**2.** 폴더 안의 `RUN_ND_FIX2.cmd` 더블클릭 

   - 자동으로 **기존 키트** 

   - ( `C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1` )에 스크립트를 복 사하고 

## **네이버/다음** 제안 수집을 실행함 

**3.** 결과 확인 

   - `C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\outputs\suggestio` 

   - `ns_nd_fix2.csv` 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

30/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

## (여기 **engine=naver/daum** 행이 나와야 정상) 

기존 `suggestions_all.csv` 도 **자동 병합** 되어서 다음 단계에서 그대로 사용 가능 

## **만약 또 비면(또는 error:daum만 보이면)** 

공유기/네트워크가 `daum` 특정 호스트를 필터링하는 케이스가 있다. 그때는 **네이버만 먼저** 확인: 

swift 코드 복사 

`python "C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\scripts\nd_f`   

실행 후 `suggestions_nd_fix2.csv` 에서 **engine=naver** 가 찍히면 네이버는 OK. 

다음(카카오)만 막히면, 다음은 **생략** 하고 네이버+유튜브로 먼저 굴리고, 

필요 시 내가 **대기시간↑/재시도** 옵션을 추가한 버전으로 한 번 더 줄게. 

이걸로 **국내 자동완성(네이버/다음)**이 실제 결과로 붙을 거야. 이어서 분석·자막 파이 프라인도 그대로 진행하면 된다. 

**==> picture [489 x 38] intentionally omitted <==**

야이 개새끼야 지금 폴더를 바꿨는데.. 씨발놈아 기존에 폴 더 경로로 만들면 어떻게해 씨발 짜증나게!!! 

n1_public_needs_miner_v1_full 여기로 바꿨자나 씨발새 꺄!! 

자꾸 이름 바꾸고 옮기고 하니까 씨발놈아 안되는거 잖아 

## 형, 미안. 내 러너에 **옛 경로가 하드코딩** 돼 있었던 게 문제야. 

지금부터 **폴더 이름 / 위치가 무엇이든** 그대로 돌아가게 **“현재 폴더 기준( %~dp0 )”**으 로 맞춰줄게. 

## **0) 지금 폴더** 

형이 쓰는 폴더: 

```
C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1_full
```

## **1) 먼저, “통합본” 기본 러너로 멀티 엔진 실행** 

이 폴더 안에서 `STEP1_AUTOCOMPLETE.cmd` 더블클릭 

→ `google, youtube, naver, daum` 자동완성 수집(상태 확인: 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

31/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

`outputs\suggestions_all.csv` 의 **engine** 컬럼) 

여전히 naver/daum이 안 찍히면, 아래 **FIX 러너(경로없음 버전)**를 추가해. 

## **2) 경로 없이 도는 FIX 러너 추가 (복붙해서 파일 만들어)** 

## **(A)** `scripts\nd_fix2.py` **새로 만들기** 

**1.** 폴더: `...\n1_public_needs_miner_v1_full\scripts\` 열기 

**2.** 메모장으로 **아래 내용 그대로** 저장: 파일명 `nd_fix2.py` 

python 코드 복사 `import os, sys, time, requests, pandas as pd` 

```
UA = {
"User-Agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KH
"Accept":"application/json,text/plain,*/*"
}
def_get_json(url, params=None, headers=None, timeout=10):
```

```
    r = requests.get(url, params=params, headers=headers or UA, timeout=timeout)
    r.raise_for_status()
```

```
try:
return r.json()
except Exception:
return {}
```

```
defnaver_suggest(q):
```

```
    headers = dict(UA); headers["Referer"] = "https://www.naver.com/"
    variants = [
```

```
        ("https://ac.search.naver.com/nx/ac", {"st":"111","q":q,"r_format":"json"
        ("https://ac.search.naver.com/nx/ac", {"st":"100","q":q,"r_format":"json"
        ("https://ac.search.naver.com/nx/ac", {"st":"1"  ,"q":q,"r_format":"json"
    ]
    out=[]
for url, params in variants:
        data = _get_json(url, params, headers)
for grp in data.get("items", []):
ifisinstance(grp, list):
for cand in grp:
ifisinstance(cand, list) and cand: out.append(str(cand[0]))
elifisinstance(cand, str): out.append(cand)
```

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

32/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

```
if out: break
    res, seen=[], set()
for t in out:
if t notin seen: res.append(t); seen.add(t)
return res
```

```
defdaum_suggest(q):
    headers = dict(UA); headers["Referer"] = "https://www.daum.net/"
    out=[]
# 1) suggest-bar (우선시도)
try:
        data = _get_json("https://suggest-bar.daum.net/suggest",