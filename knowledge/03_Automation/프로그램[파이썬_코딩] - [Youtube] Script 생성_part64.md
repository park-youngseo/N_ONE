
```
# time.sleep(random.randint(3, 8))
```

→ 저장 후 닫고 다시 실행해. (대기 없음) 

## **B) 질문 대기 없이 바로 돌리기(배치 없이 “직접 실행”)** 

## **1.** 폴더에서 cmd 열기: 

bash 코드 복사 `cd "C:\Users\엔오원\Desktop\yotube_script"` 

## **2. urls.txt** 열어 링크 1~2개만 넣고 저장: 

nginx 코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 315/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
notepad urls.txt
```

**3.** 바로 실행(질문 없이 즉시 진행): 

css 코드 복사 `python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size 4`   (엣지면 `chrome` → `edge` ) 

## **C) 진행 로그가 안 보여서 멈춘 것처럼 보일 때** 

`Youtube.txt.py` 에서 `ydl_opts` 설정의 

python 코드 복사 

```
"quiet": True,
```

를 

python 코드 복사 

```
"quiet": False,
```

로 바꾸고 저장 → 실행하면 **다운로드 진행 로그** 가 떠서 어디까지 갔는지 보인다. 

## **2) 1분짜리 “확실한” 테스트 루트** 

이대로만 복붙해서 실행해봐. (문제 지점 정확히 보자) 

bash 코드 복사 

```
cd"C:\Users\엔오원\Desktop\yotube_script"
python -m pip install -U yt-dlp
notepad urls.txt
```

메모장에 **공개 영상 1개 링크** 만 넣고 저장→닫기 (맨 앞에 `#` 금지) 

## 그 다음: 

css 

코드 복사 

```
python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size40000
```

 

 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

316/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**정상일 때** : 화면에 `[처리] https://...` 가 찍히고, 끝나면 

`FINAL_TXT\yotube_script_01.txt` 생김. 

## **3) 그래도 “아무 반응 없음”이면, 이 3줄 결과만 보내줘** 

bash 코드 복사 

```
cd"C:\Users\엔오원\Desktop\yotube_script"
dir /b
type urls.txt
```

`python "Youtube.txt.py" --root "." --urls "urls.txt" --lang ko --chunk-size 40000`   

마지막 커맨드의 **마지막 5줄** 을 그대로 주면, 그 자리에서 딱 고친다. 

- 형, 더는 방식 안 바꾼다. **배치 그대로 + 이 엔진 그대로** 만 손본다. 지금은 **대기 제거 / 질 문 없이 직접 실행** 두 가지 중 하나로 바로 풀 수 있어. 

**==> picture [489 x 39] intentionally omitted <==**

Pdf 150쪽 분량이 txt 4만자 분량이랑 비슷하냐? 

21s 동안 생각함 

## 형, **결론부터** : 

- [사실] **아니야. 150쪽 PDF ≠ 4만 자 TXT** . 보통 150쪽이 **훨씬 더 많아** . 

## **A단계｜한눈에** 

- [사실] 보통 A4(11~12pt, 일반 여백/줄간격) 기준 **1쪽 ≈ 1,500~2,000자** . 

- [추론] 그러면 **150쪽 ≈ 22.5만~30만 자** 쯤. 

- [사실] **4만 자** 는 같은 조건에서 **약 20~30쪽** 정도. 

## **B1｜핵심 숫자만** 

- [사실] 4만 자 → **약 20~30쪽** 

- [추론] 150쪽 → **약 22.5만~30만 자** 

- [가정] 폰트가 크거나 그림이 많으면 변동 있지만, 그래도 150쪽이 4만 자보다 **크다** 는 판단은 변하지 않음. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 317/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **B2｜손가락 규칙(암산용)** 

- [사실] **1쪽 ≈ 1,800자** 로 잡고 

페이지수 ≈ (총자수 ÷ 1,800) 

      - 총자수 ≈ (페이지수 × 1,800) 

   - [예시/추론] 4만 자 ÷ 1,800 ≈ **22쪽** 정도. 

   - [예시/추론] 150쪽 × 1,800 ≈ **27만 자** 정도. 

- 즉, 형이 40,000자 기준으로 나눠 놓은 건 **대략 20~30쪽짜리 묶음** 이야. 

- 150쪽짜리 텍스트라면 **대략 6~8묶음** (4만 자 단위)로 생각하면 돼. 

**==> picture [489 x 39] intentionally omitted <==**

그런데 왜 4만자를 기준으로 너에게 입력하라고해 pdf가 더 많은데 

## 형, 결론만 딱 말할게. 

## **A단계｜전체 개요** 

- [사실] **4만 자는 “내가 답을 가장 정확히 내기 좋은 작업 단위”**라서 그래. 한 번에 너무 긴 글(수십만 자)을 넣으면, 모델이 중요한 문장을 놓치거나 앞뒤 문맥을 섞어 **정확도가 떨어질 수 있음** . 

- [추론] PDF 150쪽은 글자 수가 훨씬 많을 때가 많아(보통 20만~30만 자대). 그래서 **큰 문서를 4만 자씩 쪼개 주면** 각 덩어리에서 핵심을 더 선명하게 잡고, 질의응답 품질 이 안정적이야. 

## **B1｜핵심 도출 (어떻게 운용?)** 

- [사실] **기본은 4만 자** 로 두자 → 정확도·속도·안정성의 균형점. 

- [추론] 파일 개수가 너무 많으면 **6만 자** 로 높여도 됨(여전히 안전권). 

- [가정] 10만 자 이상으로 키우면, 특정 질문에서 **누락/혼동 위험** 이 눈에 띄게 늘 수 있 음. 

## **B2｜세부 브레인스토밍 (형이 바로 쓰는 방법)** 

바로 지금 쓰는 `Script.bat` 는 `--chunk-size 40000` 으로 고정돼 있어. 

- **그대로 두면** : 품질 최적(권장). 

- **덩어리 좀 키우고 싶으면** : `Script.bat` 안의 숫자 ** `40000` → `60000` **만 바꿔 저장. (다른 건 손대지 말고 숫자만) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 318/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

- 형이 **PDF만** 주고 싶다면 줘도 돼. 다만 내가 내부적으로는 결국 **비슷한 크기로 쪼개 서** 읽어. 

기억해 두면 좋은 범위(권장): 

**정확도 우선** : 30,000~40,000자 

- **파일 수 줄이기** : 50,000~60,000자 

**비권장 상한** : 최대 80,000자(그 이상은 품질 저하 가능성↑) 

요약: **PDF가 더 많아도, 나한테는 4만 자 덩어리로 주는 게 가장 안전하고 잘 맞아.** 덩어리 수가 너무 많아 귀찮으면, 숫자만 6만으로 올리는 건 괜찮아. 

**==> picture [489 x 38] intentionally omitted <==**

니가 파이썬 url스크립트 추출 명령어에 오류가 너무 많아 서 문제가 심해 지금 

26s 동안 생각함 

- 형, 미안. 이제 **파일 이름/흐름 바꾸지 않고** 딱 고쳐서 **안정적으로 도는 버전** 만 줄게. 

   - 딱 2개만 씀: `Script.bat` , `Youtube.txt.py` 

   - 배치파일은 그대로. **파이썬 파일 내용만 교체** 해서 오류(비공개/레이트리밋/멈춤) 버 텨내도록 만들었어. 

## **A단계｜전체 개요** 

- [사실] 새 `Youtube.txt.py` 는 **재시도 + 대기(랜덤) + 진행로그 표시 + 부분실패 건너 뛰기** 포함. 

- [사실] **쿠키 사용** (chrome/edge 등)을 전달하면 **비공개/연령제한** 도 처리. 

- [사실] 끝나면 `FINAL_TXT\폴더이름_01.txt, _02.txt… (각 4만자)` 만 남고 임시는 정리. 

## **B1｜딱 3줄로 실행 순서** 

**1.** 아래 **코드 전체** 를 `Youtube.txt.py` 에 **그대로 덮어쓰기→저장** . 

**2.** `Script.bat` **더블클릭** → 폴더 경로 입력 → urls.txt에 링크 넣고 **저장 후 닫기** . 

**3.** (물으면) **쿠키 사용 Y → chrome/edge** 입력 → 진행로그 보이면서 실행. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

319/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **B2｜오류방어 강화판** `Youtube.txt.py` **(이걸로 “전체 교체저장”)** 

python 코드 복사 

```
# Youtube.txt.py — 안정화버전 (파일명/흐름그대로)
# 기능: urls.txt(영상/재생목록) → 자막(vtt) → TXT → 모두합치기 → 4만자분할저장
# 안정화: 진행로그표시, URL별재시도, 레이트리밋대비랜덤대기, 부분실패건너뛰기, 임시
```

```
import argparse, re, html, shutil, time, random, sys
from pathlib import Path
```

```
DEFAULT_LANG = "ko"
DEFAULT_CHUNK = 40000
SLEEP_BETWEEN_URLS = (8, 15)      # 각 URL 사이랜덤대기(초) — 레이트리밋회피
SLEEP_BETWEEN_ITEMS = (2, 5)      # 같은재생목록내영상사이잠깐대기
MAX_RETRIES = 3# URL별재시도횟수
defsafe(s: str) -> str:
return" ".join(re.sub(r'[\\/:*?\"<>|]+', '_', s or"").split()).strip()
```

```
defvtt_to_text(vtt_path: Path) -> str:
    raw = vtt_path.read_text(encoding="utf-8", errors="ignore")
    raw = re.sub(r'^\ufeff?WEBVTT.*?\n', '', raw, flags=re.S)
    lines = []
for ln in raw.splitlines():
        s = ln.strip()
ifnot s or"-->"in s or re.fullmatch(r"\d+", s):
continue
        s = html.unescape(re.sub(r"<.*?>", "", s))
if s:
            lines.append(s)
    out, prev = [], ""
for s in lines:
if s != prev:
            out.append(s)
        prev = s
return"\n".join(out).strip()
```

```
defsplit_by_chars(text: str, size: int):
    paras = [p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks, cur = [], ""
for p in paras:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 320/341 