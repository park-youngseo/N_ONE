```
where python >nul 2>nul || (echo [오류] Python이필요합니다. 설치후다시실행하세요.
```

```
REM 1) 위치/스크립트확인
set "HERE=%~dp0"
set "ROOT=%HERE%"
if not exist "%HERE%Script_txt.py" (
  echo [오류] Script_txt.py 가이폴더에없습니다.
  echo 같은폴더(정리할상위폴더)에두고다시실행하세요.
  pause & exit /b
)
```

```
REM 2) urls.txt 있으면순서반영, 없으면이름순
if exist "%HERE%urls.txt" (
```

```
  echo [실행] python "Script_txt.py" --root "%ROOT%" --urls "%HERE%urls.txt"
  python "Script_txt.py" --root "%ROOT%" --urls "%HERE%urls.txt"
) else (
```

```
  echo [실행] python "Script_txt.py" --root "%ROOT%"
  python "Script_txt.py" --root "%ROOT%"
)
echo.
echo [완료] FINAL_TXT 폴더에서 '폴더이름_01.txt, _02.txt...' 만남았습니다.
pause
```

## **3) 실행** 

`yotube_script` 폴더 안에 

영상별 TXT 결과들(예전처럼) 

`Script_txt.py` , `Script_txt.bat` , (선택) `urls.txt` 

가 있는지 확인 → `Script_txt.bat` **더블클릭** . 

  

## **4) 결과(동일)** 

markdown 코드 복사 

```
yotube_script\
  FINAL_TXT\
    yotube_script_01.txt
    yotube_script_02.txt
    yotube_script_03.txt
    ...
```

## **영상별 TXT** 와 **합본** 은 자동 삭제, **최종본(각 4만자)** 만 남음. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

250/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **참고** 

스크립트 내용 자체를 바꾸고 싶다면(이름만 바꾼 버전 붙여달라면) 말만 해. 

- 지금은 **기존** `finalize_txt_40k.py` **내용을 그대로 쓰되, 파일명만** `Script_txt.py` 로 바꿔서 동작하도록 안내했어. 

**==> picture [489 x 39] intentionally omitted <==**

finalize_txt_40k.py 도 코드 내용을 적용해줘 

7s 동안 생각함 

형, 바로 적용 가능한 **최종 코드** 로 바꿔서 줄게. 

파일 이름은 요청대로 `Script_txt.py` (정리 도우미) + `Script_txt.bat` (더블클릭 버튼) 조 합이야. 

## **A단계｜무엇을 하는가 (요약)** 

- [사실] `Script_txt.py` 는 **지금까지 만든 영상별 TXT** 를 

- ① 한곳으로 모으고 → ② **라벨 없이** 하나로 이어 붙인 뒤(메모리에서만) → ③ **4만 자** 씩 `폴더이름_01/02/03…` 으로 저장 → ④ **영상별 TXT는 삭제** 한다. 

- [사실] **합본 파일은 따로 남기지 않음** (형 요청 반영). 최종적으로 `FINAL_TXT\폴더이름` 

- `_01.txt, 02.txt…` **만 남음.** 

## **B1｜저장 방법 (이름 고정)** 

**1.** `yotube_script` 폴더 안에 아래 파일로 저장: 

   - `Script_txt.py` (아래 코드 전체 복붙) 

   - `Script_txt.bat` (더블클릭용 버튼 — 이전 메시지에서 준 내용 그대로 사용) 

## **B2｜코드 본문 (복붙해서** `Script_txt.py` **로 저장)** 

python 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

251/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# Script_txt.py
```

```
# 목적:
```

- _`#   1) OUT(=`_ `작업 상위` _`)`_ `폴더 아래에 흩어진 영상별` _`TXT`_ `들을 찾아 모음` 

- _`#   2)`_ `모두 이어붙여 큰 텍스트를 만들되` _`,`_ `파일로 저장하지 않고 메모리에서 처리` 

- _`#   3) 4`_ `만 글자씩` _`'`_ `폴더이름` _`_01.txt, _02.txt, ...'`_ `로 분할 저장` 

- _`#   4)`_ `원본 영상별` _`TXT`_ `는 삭제` _`(`_ `합본 파일도 만들지 않으므로 삭제할 것 없음` _`)`_ 

```
#
```

```
# 사용예:
```

```
#   python Script_txt.py --root "D:\yotube_script"
```

- _`#   (`_ `선택` _`) URL`_ `순서대로 합치려면` _`:`_ 

- _`#   python Script_txt.py --root "D:\yotube_script" --urls "D:\yotube_script\urls.`_ 

```
#
```

```
# 주의:
```

```
#   - 영상별 TXT는기존파이프라인이만든구조(예: <제목 [ID]>\TXT\*.txt)라고가정
```

- _`#   -`_ `합칠 때` _`"`_ `제목 라벨` _`"`_ `은 넣지 않음` _`(`_ `내용만 이어 붙임` _`) #   -`_ `분할 크기는 기본` _`40,000`_ `자` _`(`_ `옵션으로 변경 가능` _`)`_ 

```
import argparse
import re
from pathlib import Path
```

```
defextract_video_id_from_url(u: str) -> str:
```

```
    u = u.strip()
if"youtu.be/"in u:
return u.split("youtu.be/")[1].split("?")[0].split("/")[0]
if"watch?v="in u:
return u.split("watch?v=")[1].split("&")[0]
if"/shorts/"in u:
return u.split("/shorts/")[1].split("?")[0].split("/")[0]
if"/embed/"in u:
return u.split("/embed/")[1].split("?")[0].split("/")[0]
return""
```

```
defload_urls(urls_path: Path):
```

```
ifnot urls_path ornot urls_path.exists():
return []
```

```
for line in urls_path.read_text(encoding="utf-8", errors="ignore").splitlines
        s = line.strip()
ifnot s or s.startswith("#"):
continue
        result.append(s)
return result
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 252/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
deffind_video_txts(root: Path):
# 예) <제목 [ID]>/TXT/*.txt 형태들수집 (깊이 3 수준: */TXT/*.txt)
returnsorted(root.rglob("TXT/*.txt"))
```

```
defvideo_id_from_filename(name: str) -> str:
# 파일명끝의 [...].txt 에서 ID 추출
    m = re.search(r"\[([A-Za-z0-9_\-]{6,})\]\.txt$", name)
return m.group(1) if m else""
```

```
defsplit_by_chars(text: str, size: int = 40000):
# 문단경계(빈줄) 우선 → 너무큰문단은강제로자름
    paras = [p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks, cur = [], ""
for p in paras:
iflen(cur) + len(p) + 2 <= size:
            cur += (("\n\n"if cur else"") + p)
else:
if cur:
                chunks.append(cur)
                cur = ""
whilelen(p) > size:
                chunks.append(p[:size])
                p = p[size:]
            cur = p
if cur:
        chunks.append(cur)
return chunks
```

```
defmain():
```

```
    ap = argparse.ArgumentParser(description="영상별 TXT 모으기 → 합치기(라벨X) → 4
"
    ap.add_argument("--root", required=True, help=작업상위폴더 (예: D:\\yotube_s
""
    ap.add_argument("--urls", default=, help="(선택) URL 목록파일경로(줄당 1개)
"
    ap.add_argument("--chunk-size", type=int, default=40000, help=분할글자수 (기본
    args = ap.parse_args()
```

```
    root = Path(args.root).expanduser().resolve()
ifnot root.exists():
raise SystemExit(f"[오류] 폴더가없습니다: {root}")
```

```
# 1) 영상별 TXT 찾기
    txt_files = find_video_txts(root)
ifnot txt_files:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 253/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
print("[오류] 수집할 TXT가없습니다. 먼저영상별 TXT를만들어두세요.")
return
```

```
# 2) URL 순서가있으면그순서를최대한반영
    url_ids = []
if args.urls:
        urls_path = Path(args.urls)
        urls = load_urls(urls_path) if urls_path.exists() else []
        url_ids = [extract_video_id_from_url(u) for u in urls if extract_video_id_
```
