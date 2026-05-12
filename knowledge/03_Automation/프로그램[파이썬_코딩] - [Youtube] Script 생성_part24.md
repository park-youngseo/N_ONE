
- _`# reclass_from_any_txt_pdf.py — TXT/PDF`_ `자동 분류기` _`(`_ `확장판` _`)`_ 

- _`# - *.clean.txt`_ `가 없어도 모든` _`.txt`_ `대상` _`(README/urls`_ `등은 제외` _`)`_ 

- _`# - txt/pdf`_ `둘 중 하나 파일명에` _`[VideoID]`_ `가 있으면` _`yt-dlp`_ `메타 조회로` 

- _`#   out/<`_ `채널` _`>/<YYYY-MM-DD>_<`_ `제목` _`> [ID]/`_ `아래로` _`TXT`_ `와` _`PDF`_ `를 함께 이동` 

- _`# - ID`_ `가 전혀 없으면` _`out/_Unmatched/`_ `로 안전 보관` 

```
import os, re, json, subprocess, shutil
from datetime import datetime
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 123/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
ROOT = os.path.abspath(".")
DEST_ROOT = os.path.join(ROOT, "out")
UNMATCHED = os.path.join(DEST_ROOT, "_Unmatched")
```

```
ID_RE = re.compile(r"\[([A-Za-z0-9_-]{6,})\]")
BAD = re.compile(r'[\\/:*?"<>|]+', re.U)
EXCLUDE = re.compile(r"(readme|urls\.txt|cookies\.txt)$", re.I)
defsafe(s:str) -> str:
    s = BAD.sub("_", s or"")
return" ".join(s.split()).strip()
defyt_info_by_id(video_id: str):
ifnot video_id: return {}
    url = f"https://youtu.be/{video_id}"
try:
        p = subprocess.run(
            ["python","-m","yt_dlp","-J", url],
            capture_output=True, text=True, check=False
        )
if p.returncode != 0ornot p.stdout.strip():
return {}
        data = json.loads(p.stdout)
ifisinstance(data, dict) and"entries"in data:
            data = (data["entries"] or [None])[0]
ifnot data:
return {}
        up = data.get("upload_date")  # YYYYMMDD
        date = ""
if up andlen(up)==8and up.isdigit():
            date = datetime.strptime(up, "%Y%m%d").strftime("%Y-%m-%d")
return {
"title": data.get("title") or"",
"channel": data.get("channel") or data.get("uploader") or"",
"date": date,
"id": video_id
        }
except Exception:
return {}
defguess_id(name:str) -> str:
    m = ID_RE.search(name)
return m.group(1) if m else""
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 124/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
defpair_pdf(txt_path:str) -> str:
    base, _ = os.path.splitext(txt_path)
    cand = base + ".pdf"
return cand if os.path.exists(cand) else""
```

```
defmove_pair(txt_path:str, pdf_path:str, meta:dict):
```

```
    title = meta.get("title","") or os.path.splitext(os.path.basename(txt_path))[0
    channel = meta.get("channel","") or"UnknownChannel"
    date = meta.get("date","")
    vid = meta.get("id","")
    leaf = f"{(date+'_') if date else''}{safe(title)}{(' ['+vid+']') if vid else
    dest_dir = os.path.join(DEST_ROOT, safe(channel), leaf)
    os.makedirs(dest_dir, exist_ok=True)
for p in [txt_path, pdf_path]:
```

```
if p and os.path.exists(p):
```

```
            dst = os.path.join(dest_dir, os.path.basename(p))
```

```
if os.path.abspath(p) != os.path.abspath(dst):
                shutil.move(p, dst)
return dest_dir
```

```
defto_unmatched(txt_path:str, pdf_path:str):
    os.makedirs(UNMATCHED, exist_ok=True)
for p in [txt_path, pdf_path]:
if p and os.path.exists(p):
```

```
            shutil.move(p, os.path.join(UNMATCHED, os.path.basename(p)))
```

```
defmain():
    moved = 0
    skipped = 0
    os.makedirs(DEST_ROOT, exist_ok=True)
for root, _, files in os.walk(ROOT):
if os.path.abspath(root).startswith(os.path.abspath(DEST_ROOT)):
continue
for fn in files:
ifnot fn.lower().endswith(".txt"):
continue
if EXCLUDE.search(fn):
continue
            txt_path = os.path.join(root, fn)
            pdf_path = pair_pdf(txt_path)
            base_txt = os.path.basename(txt_path)
```

```
# 파일명에서 ID 시도 (txt, pdf 둘다확인)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

125/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
            vid = guess_id(base_txt)
ifnot vid and pdf_path:
                vid = guess_id(os.path.basename(pdf_path))
            meta = yt_info_by_id(vid) if vid else {}
if meta:
                dest = move_pair(txt_path, pdf_path, meta)
print(f"[OK] {base_txt} -> {dest}")
                moved += 1
else:
                to_unmatched(txt_path, pdf_path)
print(f"[UNMATCHED] {base_txt} -> {UNMATCHED}")
                skipped += 1
print(f"\n[완료] 이동: {moved}개 / ID불명보관: {skipped}개")
if os.path.isdir(UNMATCHED):
```

```
print("[안내] out/_Unmatched 내파일명에 [영상ID]를추가하고다시실행하면제
```

```
if __name__ == "__main__":
    main()
```

## **2.** 실행: 

nginx 코드 복사 

```
python reclass_from_any_txt_pdf.py
```

결과 구조(엔오원 실무 규격): 

`out/채널명/YYYY-MM-DD_제목 [ID]/` 에 TXT와 PDF가 **같이 이동** 

- **ID가 없는 파일** 은 `out/_Unmatched/` 에 보관 (파일명에 `[영상ID]` 만 붙여주면 재 실행 시 자동 제자리 이동) 

## **바로 다음 액션(막힘 없이)** 

**1.** 먼저 `dir /s *.clean.txt` 결과에서 **있는 위치** 가 보이면 그 폴더에서 위 스크립트를 실행. 

**2.** 결과 폴더 `out/…` 스샷이나 마지막 10줄 로그만 주면, **토픽 사본/분권/파일명 규칙** 까 지 이어서 마무리해줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

126/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**==> picture [289 x 132] intentionally omitted <==**

뭐야? 어떤거야? 

27s 동안 생각함 

## 형, 지금 화면은 **작업 루트 폴더** 고, 여기서 우리가 쓸 핵심 파일은 딱 이거야: 

## **A단계｜지금 필요한 것** 

`subs_cleaner_plus.py` → 자막 정리해서 **TXT만** 만들기. 

- `유튜브List.bat` (또는 `run_prompt.bat` ) → URL 넣으면 자막 받고 TXT 생성(배치). 

- `txt_to_pdf_any.py` → 폴더/하위폴더 **모든 TXT → PDF + 통합본** . 

- **(아직 없음)** `reclass_from_any_txt_pdf.py` → TXT/PDF를 **채널/날짜/제목** 폴더로 자 동 분류. 

## **B2단계｜바로 정리(어떤 폴더가 맞는지 모를 때)** 

**1.** 이 화면에서 주소창 클릭 → `cmd` 입력 → Enter. 

## **2. TXT/PDF가 있는 위치 자동 찾기:** 

bash 코드 복사 `dir /s *.clean.txt dir /s *.pdf` 

   - → 경로에 **어느 폴더** 에 있는지 뜸(예: `오토코더_파이썬_재생목록\…` ). 

**3.** 그 폴더로 이동: 

bash 코드 복사 

```
cd"표시된\폴더경로"
```

**4. 분류 스크립트 만들기** (TXT/PDF를 채널/날짜/제목으로 자동 이동): 

메모장에 아래 두 줄 복붙 → `reclass_from_any_txt_pdf.py` 로 저장(현재 폴더): 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

127/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

scss 

코드 복사 

```
(내가이전에준 reclass_from_any_txt_pdf.py 내용그대로붙여넣기)
```

저장 끝나면 실행: 

nginx 코드 복사 

```
python -m pip install yt-dlp
python reclass_from_any_txt_pdf.py
```

→ 결과는 루트의 `out/채널명/날짜_제목 [ID]/` 아래에 TXT와 PDF가 **쌍으로 이동** . 

## **5. PDF 변환/통합이 필요하면** (TXT만 있을 때): 

nginx 

코드 복사 

```
python -m pip install reportlab PyPDF2
python txt_to_pdf_any.py
```

## **파일별 역할(헷갈릴 때)** 

`유튜브List.bat` = 재생목록/여러 URL → 자막 받고 TXT 만들기(배치). 

`subs_cleaner_plus.py` = 자막(.vtt/.srt) → **중복 제거 TXT** . 
