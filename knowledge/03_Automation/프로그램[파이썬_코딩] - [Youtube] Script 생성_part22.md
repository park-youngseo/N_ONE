https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 113/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

코드 복사 

nginx 

```
python txt_to_pdf_any.py
```

결과: 각 TXT 옆에 **개별 PDF** + 최상위에 **ALL_LECTURES.pdf** (폴더별 통합) 생성. [사실] 

이 루트는 **파일명이 곧 메타** 라서 정확히 **강의/토픽 단위** 로 딱 맞게 떨어진다. [추론] 

## **방법 2)** `ALL_LECTURES.pdf` **만 있다 → “PDF 되돌리기 루트”** 

한 번에 따라하면 돼. (폴더=현재 `ALL_LECTURES.pdf` 있는 곳) 

## **1. 필요 도구 설치** (한 번만) 

nginx 코드 복사 

```
python -m pip install pdfplumber PyPDF2 reportlab
```

[사실] 

## **2. 분할+분류 스크립트 만들기** 

메모장 열고 아래를 **그대로** 붙여넣어 ** `pdf_split_reclassify.py` **로 저장: 

python 코드 복사 

```
# pdf_split_reclassify.py — ALL_LECTURES.pdf를제목단위로자동분할/정리
# 사용: python pdf_split_reclassify.py
# 결과:
```

```
#   recovered/sections/  -> 개별 PDF(제목_번호.pdf)
#   recovered/text/      -> 개별 TXT
```

```
#   recovered/ALL_BY_TOPIC/주제/ -> (선택) 키워드로주제사본정리
```

```
import os, re, json
import pdfplumber
from PyPDF2 import PdfReader, PdfWriter
```

```
PDF_NAME = "ALL_LECTURES.pdf"
OUT_ROOT = "recovered"
SECT_DIR = os.path.join(OUT_ROOT, "sections")
TXT_DIR  = os.path.join(OUT_ROOT, "text")
TOPIC_DIR= os.path.join(OUT_ROOT, "ALL_BY_TOPIC")
os.makedirs(SECT_DIR, exist_ok=True); os.makedirs(TXT_DIR, exist_ok=True)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

114/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# (선택) 주제키워드 — 필요시수정/추가
TOPICS = {
```

```
"Python/파이썬": ["python","파이썬","변수","리스트","딕셔너리","플레이라이트","셀
"크롤링/자동화": ["크롤링","crawler","selenium","playwright","스크롤","헤드리스"
"AI/도구활용": ["GPT","클로드","구글 AI","토큰","컨텍스트","Cursor","커서"],
}
```

```
IGNORE_FIRSTLINE_PAT = re.compile(r"^(Generated KST|Page\s+\d+)", re.I)
```

```
defis_title_line(line: str) -> bool:
ifnot line: returnFalse
if IGNORE_FIRSTLINE_PAT.search(line): returnFalse
```

```
# 간단휴리스틱: 한줄길이짧고, 문장부호과다아님, 공백으로만안됨
iflen(line.strip()) < 6:  returnFalse
iflen(line.strip()) > 120: returnFalse
# 제목스럽게: 마침표로끝나지않는경향
if line.strip().endswith("."): returnFalse
returnTrue
```

```
defextract_firstline(page) -> str:
    txt = (page.extract_text() or"").strip()
ifnot txt: return""
for ln in txt.splitlines():
        s = ln.strip()
if s:
return s
return""
```

```
defsplit_pdf_by_titles(pdf_path: str):
    reader = PdfReader(pdf_path)
with pdfplumber.open(pdf_path) as pdf:
        titles = []
for i, page inenumerate(pdf.pages):
            first = extract_firstline(page)
            titles.append(first)
```

```
# 섹션시작페이지찾기: "제목처럼보이는첫줄"이나타나는페이지
    start_pages = []
for i, first inenumerate(titles):
if is_title_line(first):
# 직전페이지의첫줄과다르고, 새섹션으로볼만한경우
if i == 0or first != titles[i-1]:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 115/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
                start_pages.append(i)
```

```
ifnot start_pages:
print("[경고] 제목후보를찾지못했습니다. 휴리스틱을조정하세요.")
return []
    start_pages.append(len(titles))  # 마지막경계
    sections = []
for idx inrange(len(start_pages)-1):
        s, e = start_pages[idx], start_pages[idx+1]  # [s, e)
        title = titles[s] orf"Section_{idx+1}"
        safe_title = re.sub(r'[\\/:*?"<>|]+', '_', title).strip()
ifnot safe_title: safe_title = f"Section_{idx+1}"
# PDF 저장
        writer = PdfWriter()
for p inrange(s, e):
            writer.add_page(reader.pages[p])
        out_pdf = os.path.join(SECT_DIR, f"{idx+1:03d}_{safe_title}.pdf")
withopen(out_pdf, "wb") as f:
            writer.write(f)
# TXT 저장
        text_buf = []
with pdfplumber.open(pdf_path) as pdf2:
for p inrange(s, e):
                text_buf.append((pdf2.pages[p].extract_text() or"").strip())
        out_txt = os.path.join(TXT_DIR, f"{idx+1:03d}_{safe_title}.txt")
withopen(out_txt, "w", encoding="utf-8", errors="ignore") as f:
            f.write("\n\n".join([t for t in text_buf if t]))
        sections.append({"start": s+1, "end": e, "title": title,
"pdf": out_pdf, "txt": out_txt})
print(f"[분할] p.{s+1}~{e} → {os.path.basename(out_pdf)}")
return sections
deftopic_copy(sections):
ifnot TOPICS: return
    os.makedirs(TOPIC_DIR, exist_ok=True)
import shutil
for s in sections:
try:
            txt = open(s["txt"], "r", encoding="utf-8", errors="ignore").read().lo
except Exception:
            txt = ""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 116/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
        name_low = os.path.basename(s["txt"]).lower()
        hits = set()
for topic, kws in TOPICS.items():
for kw in kws:
if kw.lower() in txt or kw.lower() in name_low:
                    hits.add(topic); break
for topic in hits:
            td = os.path.join(TOPIC_DIR, re.sub(r'[\\/:*?"<>|]+','_',topic))
            os.makedirs(td, exist_ok=True)
# 텍스트우선복사(필요시 PDF도복사)
            shutil.copy2(s["txt"], os.path.join(td, os.path.basename(s["txt"])))
print("[완료] 주제별사본정리:", TOPIC_DIR)
```

```
defmain():
ifnot os.path.exists(PDF_NAME):
print(f"[오류] {PDF_NAME}파일이현재폴더에없습니다."); return
print("[정보] 제목기반자동분할시작…")
    sections = split_pdf_by_titles(PDF_NAME)
ifnot sections:
print("[오류] 분할에실패했습니다. 휴리스틱을조정하세요."); return
print(f"[완료] 총{len(sections)}개섹션으로분할")
    topic_copy(sections)
```

```
print(f"[결과] {OUT_ROOT}/sections, {OUT_ROOT}/text 폴더를확인하세요.")
```

```
if __name__ == "__main__":
    main()
```

## **3. 실행** 

nginx 

코드 복사 

```
python pdf_split_reclassify.py
```

## 결과 폴더: 

`recovered/sections/` → **강의별 PDF** 

`recovered/text/` → **강의별 TXT** 

`recovered/ALL_BY_TOPIC/주제/` → (키워드 매칭 사본) 

- 첫 실행 후 **파일명이 어색한 구간** 만 제목을 살짝 손봐 주면 실무 정리 끝. [권장됨] 

## **실무 규격(엔오원 기준) 제안** 

**폴더 규칙:** `콘텐츠군/채널/플레이리스트/강의번호_강의제목/` [추론/권장됨] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

117/341 

25. 11. 30. 오후 1:07 

 

   - 프로그램[파이썬/코딩] - [Youtube] Script 생성 

- **파일 규칙:** `YYYY-MM-DD_강의제목 [영상ID].clean.txt` → PDF 동명 생성 [추론/권장됨] 

- **주제 사본:** `out_by_topic/연기/오디션/운영/AI도구/` (다중 소속 허용) [권장됨] 

- **질의 워크플로우:** “해당 폴더만 업로드 → 50~150쪽 단위로 질의” (정확도↑) [권장 됨] 

필요하면 위 스크립트에 **분권(예: 500쪽마다 자동 나누기)** , **파일명 규칙 커스텀** , **메타 표 지(채널/업로드일) 자동 생성** 도 바로 얹어줄게. 

**==> picture [489 x 38] intentionally omitted <==**

우선 텍스트가 남아 있고 변환된 pdf 파일또 한 남아 있어 

형, 딱 좋아. **TXT도 있고 PDF도 있다면** “파일명 속 영상ID → 메타데이터 조회 → 폴더 자 동정리(채널/날짜/제목)”로 한 번에 끝내자. 

## **A단계｜핵심** 

- [사실] 아래 스크립트 하나로 `*.clean.txt` **+ 같은 이름의 PDF** 를 **채널/날짜/제목/ID** 규칙에 맞춰 자동 이동 정리한다. 

- [권장됨] 정리 규칙: `out/채널명/YYYY-MM-DD_제목 [ID]/` 아래에 **TXT/PDF 같이 보관** . 

## **B2단계｜손잡고 실행** 

## **0) 준비 (한 번만)** 

작업 폴더(텍스트와 PDF가 모여 있는 최상위)에서: 

nginx 코드 복사 `python -m pip install yt-dlp` 
