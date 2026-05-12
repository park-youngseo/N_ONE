
```
#   - 4만자분할은 "문단경계"를우선고려하되, 너무크면강제분할.
```

```
import argparse, re, shutil
from pathlib import Path
```

```
defextract_video_id_from_url(u: str) -> str:
    u = u.strip()
# 유튜브기본패턴들 (최소한만)
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
ifnot urls_path ornot urls_path.exists():
return []
    out = []
for line in urls_path.read_text(encoding="utf-8", errors="ignore").splitlines
        s = line.strip()
ifnot s or s.startswith("#"):
continue
        out.append(s)
return out
deffind_all_txt(root: Path):
# <제목 [ID]>/TXT/*.txt 만수집, TXT_chunks 제외
    files = sorted(root.glob("*/*/TXT/*.txt"))
return files
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 220/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
defvideo_id_from_filename(name: str) -> str:
# "... [ID].txt" 형태에서 ID 추출시도
    m = re.search(r"\[([A-Za-z0-9_\-]{6,})\]\.txt$", name)
return m.group(1) if m else""
```

```
defsplit_by_chars(text: str, size: int = 40000):
# 문단경계(빈줄) 우선 → 너무크면강제분할
import re
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
    ap = argparse.ArgumentParser()
"
    ap.add_argument("--root", required=True, help=영상별결과상위폴더 (그아래 <제
""
    ap.add_argument("--urls", default=, help="(선택) URL 목록파일경로 (줄당 1개
    args = ap.parse_args()
```

```
    root = Path(args.root).expanduser().resolve()
assert root.exists(), f"[오류] 폴더가없음: {root}"
```

```
    urls = load_urls(Path(args.urls)) if args.urls else []
    url_order_ids = [extract_video_id_from_url(u) for u in urls if extract_video_
```

```
    txt_files = find_all_txt(root)
```

```
ifnot txt_files:
```

```
print("[오류] 수집할 TXT가없습니다. 기존파이프라인으로먼저생성하세요.")
return
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 221/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
# 평면수집폴더
    flat_dir = root / "ALL_TXT" / "raw"
    flat_dir.mkdir(parents=True, exist_ok=True)
```

```
# (1) TXT 평면복사 + 인덱스구축
    index = []  # [(id, path, title)]
for p in txt_files:
        vid = video_id_from_filename(p.name)
        title = p.stem
# 복사파일명: ID가있으면 ID.txt, 없으면원래이름.txt (공백안전)
if vid:
            dst = flat_dir / f"{vid}.txt"
else:
            safe = re.sub(r'[\\/:*?"<>|]+', '_', title)
            dst = flat_dir / f"{safe}.txt"
ifnot dst.exists():
            shutil.copy2(p, dst)
        index.append((vid or dst.stem, dst, title))
```

```
# (2) 합치기 (라벨없음, URL 순서를최대한반영)
# URL에 ID가있으면그순서대로, 나머지는이름순으로뒤에.
    ordered = []
    used = set()
if url_order_ids:
        id_to_item = {vid: (vid, path, title) for (vid, path, title) in index}
for vid in url_order_ids:
if vid in id_to_item and vid notin used:
                ordered.append(id_to_item[vid]); used.add(vid)
```

```
# 남은것들(매칭안된것)은이름순
for item insorted([x for x in index if x[0] notin used], key=lambda x: x[2]
        ordered.append(item); used.add(item[0])
```

```
    all_txt_path = flat_dir.parent / "ALL_COMBINED.txt"
withopen(all_txt_path, "w", encoding="utf-8") as out:
for _, path, _ in ordered:
            out.write(path.read_text(encoding="utf-8", errors="ignore"))
            out.write("\n\n")  # 라벨없이공백만
```

```
# (3) 4만자분할 (01/02/03…)
```

```
    combined = all_txt_path.read_text(encoding="utf-8", errors="ignore")
    chunks = split_by_chars(combined, size=40000)
    chunks_dir = flat_dir.parent / "COMBINED_chunks"
    chunks_dir.mkdir(parents=True, exist_ok=True)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 222/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
for i, c inenumerate(chunks, 1):
```

```
        (chunks_dir / f"{i:02d}.txt").write_text(c, encoding="utf-8")
```

```
print(f"[완료] 평면수집: {flat_dir}")
print(f"[완료] 합본: {all_txt_path}")
print(f"[완료] 4만자분할: {chunks_dir} (01.txt, 02.txt, ...)")
if __name__ == "__main__":
    main()
```

## **쓰는 순서(형이 원하는 그대로)** 

**1. 기존 방식 그대로** URL 넣어 **영상별 TXT** 를 만든다. (여기선 아무것도 바꾸지 않음) [사 실] 

**2.** 위 `combine_and_split_40k.py` 를 **영상별 결과 폴더들의 상위(OUT)** 에 저장한다. [사 실] 

**3.** (선택) URL 순서를 반영하고 싶으면, 같은 위치에 `urls.txt` (줄당 1개 URL)도 둔다. [사실] 

**4. 나중에** 형이 원할 때 아래 중 하나로 실행: 

```
python combine_and_split_40k.py --root "D:\OUT"
```

```
python combine_and_split_40k.py --root "D:\OUT" --urls "D:\OUT\urls.txt"
```

**5.** 결과: 

makefile 코드 복사 

```
D:\OUT\ALL_TXT\
  raw\               ← 영상별 TXT가평면복사
  ALL_COMBINED.txt   ← 전부합친하나 (라벨없음)
  COMBINED_chunks\
    01.txt 02.txt…   ← 4만자기준분할(이름만 01/02/03)
```

딱 이대로면 **안 건드리고** , **안 묻고** , **안 붙이고** , 끝에 **합치고 4만자 분할** 만 한다. 원하면 청크 크기(예: 60,000자)만 바꾸는 옵션도 넣어줄게. 

**==> picture [489 x 39] intentionally omitted <==**

## 초등학생도 이해가능하게 이야기해 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 223/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

10s 동안 생각함 

## 형, **진짜 쉽게** 설명할게. 

지금처럼 **링크 넣어서 TXT가 여러 개** 생긴 다음, **나중에 한 번에 정리** 만 해주면 돼. (지금 은 만들지 않음) 

## **A단계｜한눈에 보기** 

- [사실] **지금 방식 그대로** 영상마다 TXT가 생긴다. 

- [사실] 나중에 **정리 도우미** 로 딱 3가지만 자동으로 한다: 

**1. 모으기** : 여기저기 있는 TXT를 **한 폴더** 에 쫙 모으기 

**2. 합치기** : 모인 TXT를 **하나의 큰 TXT(ALL_COMBINED.txt)** 로 만들기 

**3. 나누기** : 그 큰 TXT를 **4만 글자씩** ** `01.txt, 02.txt, 03.txt…` **로 **숫자만 붙여** 나누기 

(본문 안에 제목 라벨은 안 붙임 → 내용만 깨끗하게 이어 붙음) 

## **B2단계｜조금만 자세히** 

## **왜 이렇게 하냐?** 

- [사실] 한 폴더에 모이면 **헷갈리지 않음** . 

- [사실] 하나로 합치면 **검색·읽기 쉽다** . 

- [사실] 4만 글자 조각은 **내가 가장 잘 읽고 생각하기 좋은 크기** 라서 **정확도가 올 라감** . 

- [사실] `01, 02, 03…` 처럼 **숫자만** 붙이면 **순서가 딱 맞게** 정렬된다. 

## **정리 후 폴더 모습(예시)** 

makefile 코드 복사 `D:\OUT\ALL_TXT\ raw\                ← 영상별 TXT가 평면으로 모여 있음 ALL_COMBINED.txt    ← 전부 합친 하나 COMBINED_chunks\ 01.txt 02.txt 03.txt ...` 

## **언제 쓰나?** 

- [사실] 영상 TXT가 어느 정도 **쌓였을 때** 한 번 돌려서 **모으기→합치기→나누기** 를 해두면 깔끔하다. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 224/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

형이 “지금 정리 시작”이라고 하면, 방금 설명 그대로 **모으기→합치기→4만자 분할** 만 딱 실행하는 스크립트로 맞춰줄게. 