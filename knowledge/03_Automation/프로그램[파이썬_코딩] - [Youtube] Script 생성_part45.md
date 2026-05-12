
**==> picture [489 x 38] intentionally omitted <==**

영상별 txt 파일은 필요없으니까 지우고 합쳐진 파일 하나 를 폴더이름_01,02 순서대로 4만자 기준으로 분할해 다합 쳐진 파일도 이작업을 하면 필요없을거아냐? 그것도 지워 결국 폴더이름으로 만들어진 txt 01,02,03 등이 최종으로 남는거야 

36s 동안 생각함 

형, 딱 형 말대로 정리만 해주는 “후처리” 하나로 끝낼게. 

**영상별 TXT 싹 모으기 → 하나로 합치기 → 4만 글자씩** `폴더이름_01/02/03…` **로 나누기 → 원본 TXT랑 합본은 삭제** . (실행은 안 함) 

## **A단계｜핵심(쉬운 말)** 

- [사실] **지금 하던 대로** 영상마다 TXT를 먼저 만들어 둔다. 

- [사실] 아래 **후처리 파일 1개** 로 **모으고→합치고→4만자씩 나누고→원본/합본은 지운 다** . 

- [사실] 마지막에 `폴더이름_01.txt, _02.txt, _03.txt…` **만 남는다** . 

## **B2단계｜딱 3단계 사용법** 

**1.** 메모장 열고 아래 코드 전부 붙여넣기 → `finalize_txt_40k.py` 로 저장. 

**2.** 이 파일을 **결과들이 모여 있는 “상위 폴더(OUT)”** 에 둔다. 

**3.** 검은창에서 (예시): 

nginx 

코드 복사 

```
python finalize_txt_40k.py --root "D:\OUT"
```

(URL 넣은 **순서대로** 합치고 싶으면, 같은 곳에 `urls.txt` (줄당 1개 URL)도 두고:) 

nginx 코드 복사 

```
python finalize_txt_40k.py --root "D:\OUT" --urls "D:\OUT\urls.txt"
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 225/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**후처리 코드:** `finalize_txt_40k.py` 

python 

코드 복사 

```
# finalize_txt_40k.py
# 하는일:
```

```
# 1) OUT 아래에있는영상별 TXT들을찾아서한폴더로모음(임시)
# 2) 전부합쳐큰텍스트하나를만들고
# 3) 4만글자씩 '폴더이름_01.txt, _02.txt ...' 로나눔
# 4) 영상별 TXT와합쳐진파일은삭제 → 최종파일만남김
import argparse, re, shutil
from pathlib import Path
```

```
defextract_video_id_from_url(u: str) -> str:
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
ifnot urls_path ornot urls_path.exists():
return []
    out = []
for line in urls_path.read_text(encoding="utf-8", errors="ignore").splitlines
        s = line.strip()
ifnot s or s.startswith("#"):
continue
        out.append(s)
return out
```

```
deffind_video_txts(root: Path):
```

```
# 예: <제목 [ID]>/TXT/*.txt 형태들을전부수집
returnsorted(root.rglob("TXT/*.txt"))
```

```
defvideo_id_from_filename(name: str) -> str:
# "... [ID].txt" 형태에서 ID 추출
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 226/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    m = re.search(r"\[([A-Za-z0-9_\-]{6,})\]\.txt$", name)
return m.group(1) if m else""
```

```
defsplit_by_chars(text: str, size: int = 40000):
# 문단(빈줄) 경계우선 → 너무크면강제분할
    paras = [p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks, cur = [], ""
for p in paras:
iflen(cur) + len(p) + 2 <= size:
            cur += (("\n\n"if cur else"") + p)
else:
if cur:
                chunks.append(cur); cur=""
whilelen(p) > size:
                chunks.append(p[:size]); p = p[size:]
            cur = p
if cur:
        chunks.append(cur)
return chunks
```

```
defmain():
    ap = argparse.ArgumentParser()
"
    ap.add_argument("--root", required=True, help=영상별결과들의상위폴더(OUT)")
""
    ap.add_argument("--urls", default=, help="(선택) URL 목록파일(줄당 1개) — 순
"
    ap.add_argument("--chunk-size", type=int, default=40000, help=조각크기(기본 4
    args = ap.parse_args()
```

```
    root = Path(args.root).expanduser().resolve()
assert root.exists(), f"[오류] 폴더가없음: {root}"
```

```
# 1) 영상별 TXT 모으기
    txt_files = find_video_txts(root)
ifnot txt_files:
print("[오류] 모을 TXT가없습니다. 먼저영상별 TXT를만들어두세요.")
return
# URL 순서반영준비
    url_ids = []
if args.urls:
        urls = load_urls(Path(args.urls))
        url_ids = [extract_video_id_from_url(u) for u in urls if extract_video_id_
    id_to_path = {}
for p in txt_files:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 227/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
        vid = video_id_from_filename(p.name)
```

```
        id_to_path[vid or p.name] = p
```

```
# 2) 합치기(라벨 X, 공백만)
```

```
    ordered_paths = []
    used = set()
```

```
for vid in url_ids:
```

```
        key = vid if vid in id_to_path elseNone
```

```
if key and key notin used:
            ordered_paths.append(id_to_path[key]); used.add(key)
```

```
for k, p insorted(id_to_path.items(), key=lambda kv: kv[0].lower()):
if k notin used:
```

```
            ordered_paths.append(p); used.add(k)
```

```
    combined_text = []
```

```
for p in ordered_paths:
        combined_text.append(p.read_text(encoding="utf-8", errors="ignore"))
        combined_text.append("\n\n")
```

```
    combined_text = "".join(combined_text).strip()
```

```
# 3) 4만자분할 + 저장 (폴더이름_01.txt ...)
    base_name = re.sub(r'[\\/:*?"<>|]+', '_', root.name).strip() or"OUTPUT"
    out_dir = root / "FINAL_TXT"
    out_dir.mkdir(parents=True, exist_ok=True)
```

```
    chunks = split_by_chars(combined_text, size=args.chunk_size)
for i, c inenumerate(chunks, 1):
        (out_dir / f"{base_name}_{i:02d}.txt").write_text(c, encoding="utf-8")
```

```
# 4) 원본영상별 TXT와합본은삭제
#   - 영상별 TXT
for p in txt_files:
try: p.unlink()
except: pass
#   - 합본파일자체는만들지않았으므로삭제할것없음 (만약만들었다면여기서지우
```

```
print(f"[완료] 최종파일들: {out_dir}\\{base_name}_01.txt, _02.txt ... 만남겼습
if __name__ == "__main__":
    main()
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 228/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **결과 모습(정말 딱 이렇게 남음)** 

makefile 

코드 복사 

```
D:\OUT\
```

```
  FINAL_TXT\
    OUT_01.txt
    OUT_02.txt
    OUT_03.txt
    ...
```

여기서 `OUT` 은 **OUT 폴더 이름** 이 들어가요. (예: `Ai강의_01.txt` ) 

## **한 줄 요약(초등학생 버전)** 

**1. 예전처럼** 영상 TXT들을 만든다. 

**2. finalize_txt_40k.py** 를 OUT 폴더에서 실행한다. 

**3.** 그러면 **영상별 TXT는 지우고** , `폴더이름_01/02/03…` **만 남겨** 준다. (각 **4만 글자** ) 