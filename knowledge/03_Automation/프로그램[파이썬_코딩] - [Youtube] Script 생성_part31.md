
원하면 위 12개 중 **딱 하나** 골라서, 형 PC 기준으로 **완성 스크립트** 를 바로 만들어줄게. 

예: “11번 링크점검”처럼 **번호만** 말해줘 — 바로 실행 가능한 코드로 뽑아줄게, 형. 

**==> picture [489 x 38] intentionally omitted <==**

우선 지금 유튜브 스크립트를 가져오는 파이썬 프로그램 을 만들었어 다시 한번 정리해줄래? 실행방법과 순서 그리 고 다른컴퓨터에서 이걸 어떻게 작동시킬수 있는지 또한 꼭 vvt파일을 만들고 그걸 텍스트로 변환해서 pdf파일로 변환하고 

그 분량을 니가 분석하기 쉽게 

특정페이지로 분류하는게 번거로움이 있어 출력된 유튜브 스크립트를 바로 특정 페이지를 pdf파일로 변환시켜서 예 시) 01-03 

처럼 분류할수 있는지 그리고 

지금은 유튜브 url을 넣고  그 유튜브의 스크립트를 지정된 폴더에 저장 까지가 첫번째 과정이라면 

출력된 파일을 텍스트로 그다음  pdf로 이후에 니가 인식 할수 있도록 특정페이지를 분류하는 작업을 했어 그걸 한 번에 할수 있는지? 그게 불가능하면 어디까지 한번클릭으 로 가능한지 알고싶어 

42s 동안 생각함 

형, 지금까지 만든 흐름을 **딱 1번 실행** 으로 끝내는 버전으로 정리해 줄게. 핵심은: **유튜브 URL → (자동) 자막 다운로드 → (자동) 텍스트 정리 → (자동) PDF 생성 → (자동) 120쪽 단위로 01/02/03… 분할** 까지를 한 번에 처리. [사실] 

## **A단계｜전체 개요** 

- [사실] 다른 PC에서도 동작: **Python 3.10+**만 있으면 됨. 처음 한 번 `pip` 로 3개 패 키지 설치(yt-dlp, reportlab, PyPDF2). 

- [사실] **VTT 파일을 눈에 보이게 처리할 필요 없음.** 내부에서 받아 바로 TXT로 변환하 고 PDF로 만든 뒤, **원하면 VTT 자동 삭제(--cleanup)** . 

- [사실] **원클릭** : 배치( `유튜브→PDF_원클릭.bat` ) 실행 → URL/폴더/파트쪽수 입력 → 끝. 

- 결과는 `…\영상제목 [ID]\PDF_parts\01.pdf, 02.pdf, …` 형식. 

- [추론] 120쪽 대신 50~150쪽 등으로 바꾸면 형이 올리기 좋은 질의 단위를 바로 얻을 수 있어. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 157/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **B2단계｜설치 · 실행 · 다른PC에서 쓰는 법** 

## **1) (최초 1회) 필수 설치** 

검은 창(cmd)에서: 

css 

코드 복사 

```
python -m pip install --upgrade pip
python -m pip install yt-dlp reportlab PyPDF2
```

다른 컴퓨터도 **동일** . Python만 있으면 위 2줄로 준비 끝. [사실] 

## **2) 파이썬 스크립트 저장:** `yt_to_pdf_parts.py` 

메모장 열고 아래 전부 복붙 → 작업 폴더에 `yt_to_pdf_parts.py` 로 저장. 

python 코드 복사 

```
# yt_to_pdf_parts.py — 유튜브 URL → (자막)TXT → PDF → 01/02/03 … 분할까지원샷
```

```
# 사용예) python yt_to_pdf_parts.py --url "https://youtu.be/...." --out "D:\OUT"
import os, re, json, argparse, html
from pathlib import Path
from datetime import datetime
```

```
# 1) yt-dlp로자막(vtt)만다운로드
```

```
defdownload_vtt(url:str, out_dir:Path, lang:str="ko", cookies:Path|None=None):
from yt_dlp import YoutubeDL
    out_dir.mkdir(parents=True, exist_ok=True)
    ydl_opts = {
```

```
"skip_download": True,
"writesubtitles": True,
"writeautomaticsub": True,
"subtitleslangs": [lang],
"subtitlesformat": "vtt",
"outtmpl": str(out_dir / "%(title)s [%(id)s].%(ext)s"),
"paths": {"home": str(out_dir)},
"noplaylist": False,
"quiet": True,
"no_warnings": True,
    }
if cookies and cookies.exists():
        ydl_opts["cookiefile"] = str(cookies)
    vtts = []
    metas = []
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 158/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
with YoutubeDL(ydl_opts) as ydl:
        info = ydl.extract_info(url, download=True)
        entries = info.get("entries") ifisinstance(info, dict) elseNone
        lst = entries or [info]
for it in lst:
ifnot it:
continue
            vid = it.get("id") or""
            title = it.get("title") or""
            channel = it.get("channel") or it.get("uploader") or""
            up = it.get("upload_date")  # YYYYMMDD
            date = ""
if up andlen(up)==8and up.isdigit():
                date = f"{up[:4]}-{up[4:6]}-{up[6:]}"
# 파일탐색: ko 우선, 없으면아무 vtt
            base = f"{title} [{vid}]"
            cand_ko = list(out_dir.glob(f"{base}.ko.vtt"))
            cand_any = list(out_dir.glob(f"{base}.*.vtt")) orlist(out_dir.glob(f
            use = cand_ko[0] if cand_ko else (cand_any[0] if cand_any elseNone)
if use:
                vtts.append(use)
```

```
                metas.append({"id":vid,"title":title,"channel":channel,"date":date
return vtts, metas
```

```
# 2) VTT → TXT(타임스탬프/태그제거)
defvtt_to_text(vtt_path:Path)->str:
    raw = vtt_path.read_text(encoding="utf-8", errors="ignore")
# 헤더/타임라인제거
    raw = re.sub(r'^\ufeff?WEBVTT.*?\n','', raw, flags=re.S)
    lines = []
for ln in raw.splitlines():
ifnot ln.strip():
continue
if"-->"in ln:
continue# 타임라인
if re.match(r"^\d+$", ln.strip()):
continue# cue 번호
        ln = re.sub(r"<.*?>", "", ln)          # 태그제거
        ln = html.unescape(ln).strip()
if ln: lines.append(ln)
# 연속중복줄간단정리
    out = []
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 159/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    prev = ""
for s in lines:
if s!=prev:
            out.append(s)
        prev=s
return"\n".join(out)
# 3) TXT → PDF (ReportLab) + 파트분할
from reportlab.lib.pagesizes import A4
from reportlab.lib.units import mm
from reportlab.pdfgen import canvas
from reportlab.pdfbase import pdfmetrics
from reportlab.pdfbase.ttfonts import TTFont
```

```
defregister_korean_font():
for p in [r"C:\Windows\Fonts\malgun.ttf",
r"C:\Windows\Fonts\Malgun.ttf",
r"/System/Library/Fonts/AppleSDGothicNeo.ttc"]:
if os.path.exists(p):
try:
                pdfmetrics.registerFont(TTFont("KR", p))
return"KR"
except Exception:
pass
return"Helvetica"
```

```
defwrap_lines(text, c, font, size, maxw):
    out=[]
for raw in text.splitlines():
        s = raw.rstrip()
ifnot s:
            out.append("")
continue
        words = s.split(" ")
        cur=""
for w in words:
            t = (cur+" "+w).strip() if cur else w
if c.stringWidth(t,font,size) <= maxw:
                cur = t
else:
if cur: out.append(cur)
if c.stringWidth(w,font,size) > maxw:
# 아주긴단어강제분리
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

160/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
""
```

```
for ch in w:
```

```
if c.stringWidth(buf+ch,font,size)<=maxw: buf+=ch
else: out.append(buf); buf=ch
                    cur=buf
else:
                    cur=w
if cur: out.append(cur)
return out
```

```
deftext_to_pdf_parts(text, out_dir:Path, base_name:str, pages_per_part:int=120):
    out_dir.mkdir(parents=True, exist_ok=True)
    font = register_korean_font()
    PAGE_W, PAGE_H = A4
    margin = 20*mm
    box_w = PAGE_W - 2*margin
    line_h = 6*mm
    title_size, body_size = 16, 12
```

```
# 줄나누기계산을위해가상캔버스
    c = canvas.Canvas(str(out_dir / "_measure.pdf"), pagesize=A4)
    c.setFont(font, body_size)
    lines = wrap_lines(text, c, font, body_size, box_w)
```

```
    parts = []
    y = PAGE_H - margin
    page_count = 0
    part_idx = 1
    canv = None
defnew_canvas(idx):
        p = out_dir / f"{idx:02d}.pdf"
        cv = canvas.Canvas(str(p), pagesize=A4)
# 머리표시(파일명) 1페이지에만간단히
return cv, p
    canv, part_path = new_canvas(part_idx)
    parts.append(part_path)
# 제목한줄(작게)
    canv.setFont(font, title_size)
    canv.drawString(margin, y, base_name)
    y -= 12*mm
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 161/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    canv.setFont(font, body_size)
```

```
for ln in lines: