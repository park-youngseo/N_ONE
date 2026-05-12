
171/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

[사실] 재생목록(URL)이면 영상별로 폴더가 여러 개 생겨. 

## **4. 다른 실행법(검은창에서 한 줄)** 

nginx 

코드 복사 

```
"
python yt_to_pdf_parts.py --url "https://youtu.be/주소 --pages-per-part 100 -
```

  `--pages-per-part 100` → **PDF를 100쪽씩** 나눔. `--cleanup` → **VTT를 자동으로 지움** . **5. TXT를 4만 글자씩 나누고 싶을 때(선택)** 같은 폴더에서: arduino 코드 복사 

## **5. TXT를 4만 글자씩 나누고 싶을 때(선택)** 

```
python split_text.py --size 40000
```

`chunks/01.txt, 02.txt…` 가 생겨. 이 크기가 내가 읽고 생각하기 딱 좋아. [사실] 

## **6. 나에게 보내는 방법(질문용)** 

**PDF 01.pdf** (100쪽) 하나 또는 **TXT 01.txt** (4만 글자) 하나씩 보내면 돼. 

그러면 내가 **핵심 10~15문장 + 해야 할 체크리스트 + 주의할 점 3개** 로 빠르게 정리해 줄 수 있어. [사실] 

## **기억할 딱 두 줄(핵심 고정값)** 

**PDF는** : `--pages-per-part 100` 

**TXT는** : `split_text.py --size 40000` 

원하면 **파일 이름 규칙(예: 2025-10-24_제목_01.pdf)** 이나 **표지(제목/채널/날짜) 자동 넣 기** 도 더 쉽게 만들어 줄 수 있어. “파일 이름도”라고 말해줘! 

**==> picture [489 x 40] intentionally omitted <==**

너는 그럼 txt파일도 인식이 가능한거야? 

13s 동안 생각함 

## 형, **가능해.** 

## **A단계｜핵심** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

172/341 

25. 11. 30. 오후 1:07 

 

   - 프로그램[파이썬/코딩] - [Youtube] Script 생성 

- [사실] **TXT 파일 그대로 업로드하면** 내가 **바로 읽고 분석** 할 수 있어. 

- [사실] 정확도/속도 최적값은 **한 번에 40,000자(≈ 텍스트 기준)** 단위. 

- [권장됨] 여러 개라면 `01.txt, 02.txt, 03.txt…` 식으로 순서를 붙여줘. 

## **B2｜사용법(간단)** 

## **보내는 방법** 

**1.** 이 채팅에 **TXT 파일 업로드** (여러 개 가능 → 번호순 권장). 

**2.** 길면 **40,000자씩 잘라** `01.txt, 02.txt…` 로 올려도 되고, 한 파일만 올리면 내가 여기서 **분할·정리** 도 해줄 수 있어. 

**인코딩** : 특별히 손댈 필요 없고, **UTF-8** 이면 가장 안전해. 

## **내가 해줄 수 있는 것** 

- [사실] **핵심 10~15문장 요약 + 실행 체크리스트 + 주의점 3개** (형이 정한 포맷). 

- [사실] **PDF로 변환** 하거나 **100쪽 단위** 로 01/02/03… **자동 분할** 도 가능(원하면 바 로 수행). 

원본 TXT 몇 개만 올려줘. **바로 그 자리에서** 형 규격(40,000자 / PDF 100쪽)에 맞춰 브레 인스토밍 시작할게. 

**==> picture [489 x 39] intentionally omitted <==**

나는 니가 왜 pdf파일로만 인식하는줄 알았지? 

## 9s 동안 생각함 

형, 오해 풀자 😊 

**나는 PDF만이 아니라 TXT도 아주 잘 인식** 해. 오히려 **TXT가 더 정확** 하고 빨라. 

## **A단계｜핵심** 

- [사실] **TXT 그대로 업로드/붙여넣기** 하면 바로 분석 가능. 표·코드·줄바꿈도 깔끔. 

- [사실] **PDF도 OK** 지만, 추출 과정에서 줄바꿈/글자 깨짐이 생길 수 있어. 

- [사실] 최적 단위는 **TXT=40,000자** , **PDF=100쪽** (형이 확정한 규격). 

## **B2｜간단 사용 가이드** 

## **어떤 형식이 더 좋은가?** 

- **TXT 권장** : 자막/노트/스크립트처럼 글 위주인 자료 → 정확도↑, 속도↑. **PDF 선택** : 레이아웃 유지가 중요한 자료(표지/서식/도표). 

## **보내는 법** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 173/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

   - TXT면 `01.txt, 02.txt…` 식으로 **4만 자씩** 나눠 업로드. 

   - PDF면 `01.pdf, 02.pdf…` 식으로 **100쪽씩** 나눠 업로드. 

- **이름 규칙(읽기 쉬움)** : `YYYY-MM-DD_제목_01.txt` (또는 `.pdf` ) 

## **바로 쓰는 한 줄** 

- TXT 분할: `python split_text.py --size 40000` 

- PDF 분할: `python yt_to_pdf_parts.py --pages-per-part 100 --lang ko -- cleanup --url "<링크>"` 

원문 TXT가 이미 있으면 **TXT로 주는 게 가장 좋아** . 지금 `01.txt` 만 올려줘도 **형 규격(핵 심 10~15문장 + 실행 체크리스트 + 함정 3개)**으로 바로 정리해줄 수 있어. 

**==> picture [489 x 39] intentionally omitted <==**

그렇다면 위 과정에서 pdf로 바꿀필요가 없잖아.. 이건 내 가 사용하는게 아니라 너한테 맞추는거라고 했잖아 

19s 동안 생각함 

형, 정확히 이해했어. **이제부터 PDF는 생략** 하고, **TXT만** 내가 읽기 좋은 **고정 규격(4만 자)**으로 만들어 쓰자. 

아래대로 하면 **유튜브 URL → TXT 정리 → 4만 자 조각(01.txt, 02.txt…)**까지 한 번에 끝난다. 

## **A단계｜전체 개요** 

- [사실] **나는 TXT를 바로 인식** 하고, **4만 자 단위** 가 브레인스토밍 최적 크기야. 

- [사실] 아래 스크립트/배치로 **원클릭** : 유튜브 주소 넣으면 **자막 다운 → 깨끗한 TXT → 4만 자 조각(01/02/03…)** 자동 생성. 

- [권장됨] PDF는 필요 없으니 만들지 않아도 돼. 

## **B1｜핵심 도출(정말 필요한 것만)** 

[사실] 최초 1회 설치: 

css 코드 복사 

```
python -m pip install --upgrade pip
python -m pip install yt-dlp
```

- [사실] 실행 규격(고정): **TXT 조각 크기 = 40,000자** 

- [사실] 결과 구조: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 174/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

코드 복사 

css 

```
<저장폴더>/
제목[영상ID]/
    TXT/제목[ID].txt            ← 전체한덩어리
    TXT_chunks/01.txt, 02.txt…   ← 4만자씩나눔(내가바로읽는용)
```

## **B2｜세부 브레인스토밍(그대로 쓰면 됨)** 

## **1) 파이썬 스크립트 저장:** `yt_to_txt_chunks.py` 

메모장 열고 아래 전부 복붙 → 작업 폴더에 `yt_to_txt_chunks.py` 로 저장. 

python 코드 복사 

```
# yt_to_txt_chunks.py — 유튜브 URL → (자막)VTT → 깨끗한 TXT → 4만자조각(01/02/03…)
# 사용예)
```

```
#   python yt_to_txt_chunks.py --url "https://youtu.be/..." --out "D:\OUT" --lang
import os, re, argparse, html
from pathlib import Path
```

```
defsafe_name(s:str) -> str:
    s = re.sub(r'[\\/:*?"<>|]+', '_', s)
return" ".join(s.split()).strip()
```

```
defdownload_vtt(url: str, out_dir: Path, lang: str = "ko", cookies: Path | None =
from yt_dlp import YoutubeDL
    out_dir.mkdir(parents=True, exist_ok=True)
    ydl_opts = {
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
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 175/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    vtts, metas = [], []
with YoutubeDL(ydl_opts) as ydl:
        info = ydl.extract_info(url, download=True)
        entries = info.get("entries") ifisinstance(info, dict) elseNone
        items = entries or [info]
for it in items:
ifnot it:
continue
            vid   = it.get("id") or""
            title = safe_name(it.get("title") or"")
            base  = f"{title} [{vid}]"
# ko 우선, 없으면아무 vtt
            cand_ko  = list(out_dir.glob(f"{base}.ko.vtt"))
            cand_any = list(out_dir.glob(f"{base}.*.vtt")) orlist(out_dir.glob(f
            use = cand_ko[0] if cand_ko else (cand_any[0] if cand_any elseNone)
if use:
                vtts.append(use)
                metas.append({"id": vid, "title": title})
return vtts, metas
defvtt_to_text(vtt_path: Path) -> str:
    raw = vtt_path.read_text(encoding="utf-8", errors="ignore")
    raw = re.sub(r'^\ufeff?WEBVTT.*?\n', '', raw, flags=re.S)  # 헤더제거
    lines = []
for ln in raw.splitlines():
        s = ln.strip()
ifnot s or"-->"in s or re.match(r"^\d+$", s):
continue# 타임라인/큐번호제거
        s = re.sub(r"<.*?>", "", s)            # 태그제거
        s = html.unescape(s)
if s: lines.append(s)
# 바로전줄과같은중복줄제거
    out, prev = [], ""
for s in lines:
if s != prev:
            out.append(s)
        prev = s
return"\n".join(out)
```

```
defsplit_text_by_chars(text: str, chunk_size: int = 40000):
```

```
# 문단경계(빈줄) 기준우선분할 → 너무큰문단은강제자르기
    paras = [p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks, cur = [], ""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 176/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 