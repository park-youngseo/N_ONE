[사실] 

## **1) 스크립트 파일 만들기** 

메모장 열고 아래 전부 복사→붙여넣기 → `reclass_from_txt_pdf.py` 라는 이름으로 **TXT/PDF가 모여 있는 최상위 폴더** 에 저장. 

python 

코드 복사 

- _`# reclass_from_txt_pdf.py — N·ONE TXT/PDF`_ `자동 분류기` _`(`_ `채널` _`/`_ `날짜` _`/`_ `제목` _`/ID)`_ 

- _`#`_ `사용` _`:`_ 

- _`#   python reclass_from_txt_pdf.py`_ 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 118/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# 결과:
```

```
#   out/<채널명>/<YYYY-MM-DD>_<제목> [ID]/  아래에 *.clean.txt 와같은이름의 .pdf 함
#
# 규칙:
```

```
# - 파일명에 [영상ID]가포함되어있으면가장정확 (예: 제목 [CF4Fb9vkrik].clean.txt)
```

```
# - ID가없으면 yt-dlp 메타에서제목부분으로추정이동(정확도↓ → Unmatched로보관가능
# - 하위폴더까지재귀로모두처리
```

```
import os, re, json, subprocess, shutil
from datetime import datetime
ROOT = os.path.abspath(".")
DEST_ROOT = os.path.join(ROOT, "out")
UNMATCHED = os.path.join(DEST_ROOT, "_Unmatched")
```

```
ID_RE = re.compile(r"\[([A-Za-z0-9_-]{6,})\]")  # [VideoID]
BAD = re.compile(r'[\\/:*?"<>|]+')
```

```
defsafe(s:str) -> str:
    s = BAD.sub("_", s or"")
return" ".join(s.split()).strip()
defyt_info_by_id(video_id: str):
"""yt-dlp -J 로메타데이터조회"""
ifnot video_id:
return {}
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
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

119/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
return {
```

```
"title": data.get("title") or"",
"channel": data.get("channel") or data.get("uploader") or"",
"date": date,
"id": video_id
        }
except Exception:
return {}
```

```
defguess_id_from_name(name:str) -> str:
    m = ID_RE.search(name)
return m.group(1) if m else""
```

```
deffind_pdf_sibling(txt_path:str) -> str:
    base = os.path.splitext(txt_path)[0]
    cand = base + ".pdf"
return cand if os.path.exists(cand) else""
```

```
defmove_pair(txt_path:str, pdf_path:str, meta:dict):
    title = meta.get("title","") or os.path.splitext(os.path.basename(txt_path))[0
    channel = meta.get("channel","") or"UnknownChannel"
    date = meta.get("date","")
    vid = meta.get("id","")
# 대상폴더
    leaf = f"{date+'_'if date else''}{safe(title)}{(' ['+vid+']') if vid else'
    dest_dir = os.path.join(DEST_ROOT, safe(channel), leaf)
    os.makedirs(dest_dir, exist_ok=True)
# 이동
for p in [txt_path, pdf_path]:
if p and os.path.exists(p):
            dst = os.path.join(dest_dir, os.path.basename(p))
if os.path.abspath(p) != os.path.abspath(dst):
                shutil.move(p, dst)
return dest_dir
```

```
defmain():
    moved = 0
    os.makedirs(DEST_ROOT, exist_ok=True)
for root, _, files in os.walk(ROOT):
# out/ 아래는재귀스킵
if os.path.abspath(root).startswith(os.path.abspath(DEST_ROOT)):
continue
for fn in files:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 120/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
ifnot fn.endswith(".clean.txt"):
continue
            txt_path = os.path.join(root, fn)
            pdf_path = find_pdf_sibling(txt_path)
            base = os.path.basename(fn)
# 1) 파일명에서 ID
            vid = guess_id_from_name(base)
            meta = {}
if vid:
                meta = yt_info_by_id(vid)
# 2) 실패시 Unmatched로
ifnot meta:
                os.makedirs(UNMATCHED, exist_ok=True)
# 그래도 TXT/PDF는보존이동
for p in [txt_path, pdf_path]:
if p and os.path.exists(p):
                        shutil.move(p, os.path.join(UNMATCHED, os.path.basename(p
print(f"[UNMATCHED] {base} → {UNMATCHED}")
else:
                dest = move_pair(txt_path, pdf_path, meta)
print(f"[OK] {base} → {dest}")
                moved += 1
print(f"\n[완료] 이동/정리: {moved}개항목. out/ 폴더에서확인하세요.")
if os.path.isdir(UNMATCHED):
```

```
print("[안내] ID를못찾은파일은 out/_Unmatched 에보관했습니다. 파일명에 [영
```

```
if __name__ == "__main__":
    main()
```

## **2) 실행** 

작업 폴더에서: 

nginx 코드 복사 

```
python reclass_from_txt_pdf.py
```

- 결과: `out/채널명/YYYY-MM-DD_제목 [ID]/` 아래로 TXT/PDF가 **쌍으로 이동** . [사실] 

- **ID를 못 찾은 파일** 은 `out/_Unmatched` 에 임시 보관(파일명에 `[영상ID]` 추가 후 스크 립트를 다시 돌리면 제자리 이동). [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

121/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **(옵션) 주제(토픽)별 사본도 만들고 싶다면** 

원하는 키워드로 `topics.json` 을 만들어두면, **사본을 따로** 정리하는 스텝을 추가해줄 수 있어. 

예시: 

json 코드 복사 

```
{
"Python/기초": ["파이썬", "변수", "리스트", "딕셔너리"],
"자동화": ["셀레니움", "Selenium", "플레이라이트", "Playwright"],
"AI도구": ["Cursor", "커서", "GPT", "프롬프트"]
}
```

→ “토픽 사본도”라고 하면, 위 스크립트에 `topics.json` 읽어서 `out_by_topic/<주제>/` 아래로 **복사 사본** 까지 생성하도록 붙여줄게. [추론] 

## **엔오원 실무 규격 재확인** 

- **폴더** : `out/채널명/날짜_제목 [ID]/` ← 지금 스크립트 결과와 동일. [사실] 

- **파일** : TXT=학습/검색용, PDF=배포/인쇄용 (둘 다 한 폴더에). [권장됨] **운영 루틴** : 

**1.** 재생목록 처리 → `*.clean.txt` 생성 

**2. reclass_from_txt_pdf.py** 실행 → out/ 구조 완성 

**3.** 폴더 단위로 업로드/질의(50~150쪽 단위) [권장됨] 

막히는 지점이 있으면, **오류 메시지 그대로** 보내줘. 바로 그 자리에서 이어서 수정해줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

Microsoft Windows [Version 10.0.19045.6456] (c) Microsoft Corporation. All rights reserved. 

C:\Users\엔오원\Desktop\youtube_script_extractor\Ai 강의>python -m pip install yt-dlp 

Requirement already satisfied: yt-dlp in c:\users\엔오원 \appdata\local\programs\python\python311\lib\sit e-packages (2025.10.14) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 122/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

C:\Users\엔오원\Desktop\youtube_script_extractor\Ai 강의>python reclass_from_txt_pdf.py 

[완료] 이동/정리: 0개 항목. out/ 폴더에서 확인하세요. 

C:\Users\엔오원\Desktop\youtube_script_extractor\Ai 강의> 

형, 0개라고 뜬 건 **폴더 안에** `*.clean.txt` **를 못 찾았기 때문** 이야. (TXT가 그냥 `.txt` 로만 있거나, 다른 폴더에 있을 가능성) 

바로 해결 루트 두 가지를 같이 줄게. 

## **A단계｜바로 점검 (10초)** 

검은창(지금 폴더)에서 아래 두 줄만 실행해봐: 

bash 코드 복사 

```
dir /s *.clean.txt
dir /s *.pdf
```

- `*.clean.txt` 목록이 비면 → TXT가 **.clean.txt가 아님** 또는 **다른 폴더** 에 있음. 

## **B2단계｜그대로 정리(확장판 스크립트 사용)** 

- `*.clean.txt` 가 아니어도, 모든 `.txt` 를 잡아서 **같은 이름의 .pdf와 함께** 분류하도록 스 

- 크립트를 바꿨다. 

(파일명에 `[영상ID]` 가 있으면 채널/제목/날짜까지 자동 정렬, 없으면 `_Unmatched` 로 보 관) 

**1.** 메모장 열고 아래를 그대로 붙여넣어 `reclass_from_any_txt_pdf.py` 로 저장 (지금 폴더): 

python 코드 복사 