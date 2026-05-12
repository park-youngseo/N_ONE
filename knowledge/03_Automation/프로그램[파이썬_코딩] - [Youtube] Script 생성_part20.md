```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

102/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
from reportlab.lib.units import mm
from reportlab.pdfgen import canvas
from reportlab.pdfbase import pdfmetrics
from reportlab.pdfbase.ttfonts import TTFont
from PyPDF2 import PdfMerger
ROOT = os.path.abspath(".")
MERGED_NAME = "ALL_LECTURES.pdf"# 통합본파일명
```

```
defregister_korean_font():
# 윈도우기본: C:\Windows\Fonts\malgun.ttf
    candidates = [
r"C:\Windows\Fonts\malgun.ttf",
r"C:\Windows\Fonts\Malgun.ttf",
r"/System/Library/Fonts/AppleSDGothicNeo.ttc",  # mac
    ]
for p in candidates:
if os.path.exists(p):
try:
                pdfmetrics.registerFont(TTFont("KR", p))
return"KR"
except Exception:
pass
# 폰트를못찾으면기본 Helvetica 사용(영문위주)
return"Helvetica"
```

```
defwrap_text(text, max_width, c, font_name, font_size):
# 아주단순한줄바꿈(단어기준). reportlab은너비계산필요.
    out_lines = []
for raw in text.splitlines():
        line = raw.strip()
ifnot line:
            out_lines.append("")
continue
        words = line.split(" ")
        cur = ""
for w in words:
            test = (cur + " " + w).strip() if cur else w
if c.stringWidth(test, font_name, font_size) <= max_width:
                cur = test
else:
if cur:
                    out_lines.append(cur)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 103/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# 단어자체가매우길면잘라넣기
if c.stringWidth(w, font_name, font_size) > max_width:
                    tmp = w
                    buf = ""
for ch in tmp:
if c.stringWidth(buf + ch, font_name, font_size) <= max_w
                            buf += ch
else:
                            out_lines.append(buf)
                            buf = ch
if buf:
                        cur = buf
else:
                    cur = w
if cur:
            out_lines.append(cur)
return out_lines
```

```
deftxt_to_pdf(txt_path, font_name):
# PDF 경로
    base, _ = os.path.splitext(txt_path)
    pdf_path = base + ".pdf"
# 페이지설정
    PAGE_W, PAGE_H = A4
    margin = 20 * mm
    text_box_w = PAGE_W - 2 * margin
    y = PAGE_H - margin
# 제목추출(파일명 → [대괄호] 제거, .clean 제거)
    title = os.path.basename(base)
    title = re.sub(r"\.clean$", "", title, flags=re.I)
    title = title.replace("_", " ").strip()
# 본문읽기
    content = open(txt_path, "r", encoding="utf-8", errors="ignore").read()
# 캔버스
    c = canvas.Canvas(pdf_path, pagesize=A4)
# 폰트크기
    title_size = 16
    body_size = 12
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 104/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# 제목
    c.setFont(font_name, title_size)
    c.drawString(margin, y, title)
    y -= 12 * mm
# 본문줄바꿈
    c.setFont(font_name, body_size)
    lines = wrap_text(content, text_box_w, c, font_name, body_size)
    line_h = 6 * mm
for ln in lines:
if y < margin + line_h:
            c.showPage()
            y = PAGE_H - margin
            c.setFont(font_name, body_size)
        c.drawString(margin, y, ln)
        y -= line_h
# 꼬리표(생성일시)
    c.setFont(font_name, 8)
    c.drawRightString(PAGE_W - margin, margin/2,
                      datetime.now().strftime("Generated KST %Y-%m-%d %H:%M"))
    c.save()
return pdf_path
defmain():
    font_name = register_korean_font()
# 모든하위폴더포함하여 *.clean.txt 찾기
    files = glob.glob("**/*.clean.txt", recursive=True)
ifnot files:
print("[오류] *.clean.txt 파일을찾지못했습니다. 먼저자막을 TXT로변환하세요
return
print(f"[정보] 감지된 TXT: {len(files)}개")
    pdfs = []
for i, p inenumerate(sorted(files), 1):
try:
print(f"  ({i}/{len(files)}) {p}")
            pdf = txt_to_pdf(p, font_name)
            pdfs.append(pdf)
except Exception as e:
print(f"  [실패] {p}: {e}")
```

```
# 통합본병합
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

105/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
if pdfs:
print("[정보] 통합 PDF 병합중…")
        merger = PdfMerger()
for pdf in pdfs:
try:
                merger.append(pdf)
except Exception as e:
print(f"  [병합건너뜀] {pdf}: {e}")
        merger.write(MERGED_NAME)
        merger.close()
print(f"[완료] 개별 PDF + 통합본 '{MERGED_NAME}' 생성")
else:
print("[오류] 생성된 PDF가없습니다.")
if __name__ == "__main__":
    main()
```

## **3) 실행** 

작업 폴더에서: 

nginx 코드 복사 `python txt_to_pdf.py` 

- 그 폴더(및 하위 폴더)에 있는 **모든 `*.clean.txt` **가 **동일 위치에 PDF** 로 생기고, **최상위 폴더** 에는 `ALL_LECTURES.pdf` (전부 합친 하나)도 같이 생긴다. [사실] 

## **선택 옵션(원하면)** 

- **특정 폴더만** 대상으로 하려면, 그 폴더로 이동한 뒤 실행하면 됨. 제목·폰트 크기·여백 바꾸고 싶으면 알려줘. 바로 맞춰서 코드 수정해줄게. 

- 영상별 표지(제목/채널/업로드일) 넣고 싶다면, 전에 만든 `organize_outputs.py` 와 묶 어 **메타데이터 표지 페이지** 도 넣어줄 수 있어. 

필요하면 내가 **배치 파일(더블클릭 한 번)** 로도 만들어줄게. “배치도”라고 말해줘, 형. 

**==> picture [489 x 39] intentionally omitted <==**

Microsoft Windows [Version 10.0.19045.6456] 

(c) Microsoft Corporation. All rights reserved. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

106/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

C:\Users\엔오원\Desktop\youtube_script_extractor\Ai 강의>python txt_to_pdf.py 

[오류] *.clean.txt 파일을 찾지 못했습니다. 먼저 자막을 TXT로 변환하세요. 

C:\Users\엔오원\Desktop\youtube_script_extractor\Ai 강의>python txt_to_pdf.py 

[오류] *.clean.txt 파일을 찾지 못했습니다. 먼저 자막을 TXT로 변환하세요. 

C:\Users\엔오원\Desktop\youtube_script_extractor\Ai 강의> 

형, 지금 폴더( **Ai강의** )에 `*.clean.txt` **가 없어서** 스크립트가 못 찾은 거야. 해결은 둘 중 하나로 바로 돼. 

## **A단계｜핵심** 

- **방법 1:** `…clean.txt` 파일이 있는 폴더로 이동해서 실행. [사실] 

- **방법 2:** 스크립트를 “일반 `.txt` 도 허용” 버전으로 교체해서, 어떤 TXT든 바로 PDF로 만들기. [사실] 

## **B2단계｜손잡고 해결** 

## **방법 1 — 위치만 맞추기** 

**1.** 상위 폴더로 올라가서 전체에 `clean.txt` 가 있는지 찾기: 

bash 코드 복사 `cd .. dir /s *.clean.txt` 

**2.** 목록에 보이는 폴더로 들어가서 실행: 

bash 코드 복사 

```
cd"여기에_표시된_폴더경로"
```

```
python txt_to_pdf.py
```

- → 개별 PDF + `ALL_LECTURES.pdf` 가 생성. [사실] 

## **방법 2 — “모든 .txt 허용” 스크립트로 교체(추천)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

107/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

지금 폴더에 **일반 .txt** 만 있어도 전부 PDF로 바꿔줌(README/urls 같은 건 자동 제외). 

**1.** 메모장 열고 아래를 그대로 붙여넣어 `txt_to_pdf_any.py` 라는 이름으로 저장(현재 폴더): 

python 코드 복사 

```
# txt_to_pdf_any.py — N·ONE TXT → PDF (개별+통합, 모든 .txt 허용)
```

```
import os, glob, re
from datetime import datetime
from reportlab.lib.pagesizes import A4
from reportlab.lib.units import mm
from reportlab.pdfgen import canvas
from reportlab.pdfbase import pdfmetrics
from reportlab.pdfbase.ttfonts import TTFont
from PyPDF2 import PdfMerger
ROOT = os.path.abspath(".")
MERGED_NAME = "ALL_LECTURES.pdf"
```

```