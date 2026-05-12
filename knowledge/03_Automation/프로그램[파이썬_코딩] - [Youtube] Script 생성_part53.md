```
python "finalize_txt_40k.py" --root "%ROOT%" --urls "%ROOT%\urls.txt"
```

```
echo.
```

```
echo [완료] 결과: "%ROOT%\FINAL_TXT\폴더이름_01.txt, _02.txt ..." 만남았습니다.
echo        (원본영상별 TXT는삭제되었습니다)
pause
```

## **2) 정리 엔진:** `finalize_txt_40k.py` **(이 내용으로 저장)** 

python 

코드 복사 

```
# finalize_txt_40k.py
```

```
# 동작:
```

```
# 1) ROOT 아래의영상별 TXT(*/TXT/*.txt) 수집
```

```
# 2) 라벨없이내용만이어붙여큰텍스트생성(파일로저장하지않음)
# 3) 4만자씩 '폴더이름_01.txt, _02.txt, ...' 저장 (FINAL_TXT/)
# 4) 영상별 TXT는삭제, 합본파일은애초에만들지않음
```

```
import argparse, re
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
    out=[]
for line in urls_path.read_text(encoding="utf-8", errors="ignore").splitlines
        s=line.strip()
ifnot s or s.startswith("#"): continue
        out.append(s)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 263/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
return out
```

```
deffind_video_txts(root: Path):
# */TXT/*.txt 만수집
returnsorted(root.rglob("TXT/*.txt"))
```

```
defvideo_id_from_filename(name: str) -> str:
    m=re.search(r"\[([A-Za-z0-9_\-]{6,})\]\.txt$", name)
return m.group(1) if m else""
```

```
defsplit_by_chars(text: str, size: int = 40000):
# 문단(빈줄) 경계우선, 너무큰문단은강제분할
    paras=[p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks=[]; cur=""
for p in paras:
iflen(cur)+len(p)+2<=size:
            cur+=(("\n\n"if cur else"")+p)
else:
if cur: chunks.append(cur); cur=""
whilelen(p)>size:
                chunks.append(p[:size]); p=p[size:]
            cur=p
if cur: chunks.append(cur)
return chunks
```

```
defmain():
```

```
    ap=argparse.ArgumentParser(description="모으기→합치기(라벨X)→4만자분할→원본TXT
""
    ap.add_argument("--root", required=True, help=정리할상위폴더경로)
""
    ap.add_argument("--urls", default=, help="(선택) URL 목록파일경로 — 순서반
"
    ap.add_argument("--chunk-size", type=int, default=40000, help=분할글자수 (기본
    args=ap.parse_args()
```

```
    root=Path(args.root).expanduser().resolve()
ifnot root.exists():
raise SystemExit(f"[오류] 폴더가없습니다: {root}")
```

```
    txt_files=find_video_txts(root)
ifnot txt_files:
```

```
print("[오류] 수집할 TXT가없습니다. 먼저영상별 TXT를만들어두세요.")
return
```

```
# URL 순서반영
    url_ids=[]
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 264/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
if args.urls:
```

```
        ufile=Path(args.urls)
        urls=load_urls(ufile) if ufile.exists() else []
        url_ids=[extract_video_id_from_url(u) for u in urls if extract_video_id_f
```

```
    id_to_path={}
```

```
for p in txt_files:
        vid=video_id_from_filename(p.name)
        key=vid if vid else p.name
if key notin id_to_path:
            id_to_path[key]=p
```

```
    ordered=[]
    used=set()
for vid in url_ids:
if vid in id_to_path and vid notin used:
            ordered.append(id_to_path[vid]); used.add(vid)
for k,p insorted(id_to_path.items(), key=lambda kv: kv[0].lower()):
if k notin used:
            ordered.append(p); used.add(k)
```

```
    parts=[]
for p in ordered:
        parts.append(p.read_text(encoding="utf-8", errors="ignore"))
        parts.append("\n\n")
    combined="".join(parts).strip()
```

```
    base_name=re.sub(r'[\\/:*?"<>|]+','_',root.name).strip() or"OUTPUT"
    out_dir=root/"FINAL_TXT"
    out_dir.mkdir(parents=True, exist_ok=True)
```

```
    chunks=split_by_chars(combined, size=args.chunk_size)
for i,c inenumerate(chunks,1):
        (out_dir/f"{base_name}_{i:02d}.txt").write_text(c, encoding="utf-8")
```

```
# 원본영상별 TXT 삭제
for p in txt_files:
try: p.unlink()
except: pass
print(f"[완료] 최종파일: {out_dir}\\{base_name}_01.txt, _02.txt ... (각{args.
print("        원본영상별 TXT는삭제되었고, 합본파일은만들지않았습니다.")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 265/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
if __name__=="__main__":
```

```
    main()
```

## **사용 순서(간단)** 

**1.** `yotube_script` 에 **영상별 TXT** 가 있는 상태로 두기. 

**2.** 같은 폴더(또는 아무 곳)에서 `TXT_정리_한방.bat` **더블클릭** → **정리할 폴더 경로** 입력. `urls.txt` 가 없으면 자동 생성(원하면 바로 메모장 열어서 붙여넣기). 

**3.** 자동 실행 끝 → `정리폴더\FINAL_TXT\폴더이름_01.txt, _02.txt …` 만 남는다. 

이제부터 **파일명은 절대 안 바꿀게** : `finalize_txt_40k.py` + `TXT_정리_한방.bat` 고정. 막히는 지점이 있으면 **에러 문구 한 줄만** 알려줘. 거기서 바로 잡아줄게. 

**==> picture [489 x 39] intentionally omitted <==**

지금은 몆개 파일이야? 세개지? 근데 바로 앞엔 몇개였어? 4개지? 뭐야 도대체 

31s 동안 생각함 

형, 헷갈리게 해서 미안. 숫자 딱 정리할게. 

## **지금 기준(형이 “기존처럼 TXT는 이미 있다”면)** 

## **필수 2개** : 

**1.** `finalize_txt_40k.py` (정리 엔진) 

**2.** `TXT_정리_한방.bat` (버튼) 

**선택 1개** : `urls.txt` → 없으면 **버튼이 자동 생성** 

- 👉 그래서 보통 폴더에 **3개처럼 보일 수 있어요** (필수 2 + 자동/선택 1). 

## **바로 앞에 내가 4개라고 했던 경우는?** 

그땐 **“추출까지 자동”** 쓰자고 해서, **추출용 2개** 를 추가로 보여줬던 거야(옵션). 

- 3) `yt_to_txt_simple.py` (자막→TXT 추출) 

- 4) `유튜브→TXT_SIMPLE_MULTI.bat` (여러 URL 추출 버튼) 

- 👉 이 두 개는 **형이 이미 TXT를 만들어두는 방식이면 필요 없음** . 그래서 **지금은 빼고 2개(정리용)만 쓰는 게 정답** . 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

266/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **최종 결론(헷갈리지 않게)** 

## **형이 원한 최소 세팅 = 정리만:** 

- → `finalize_txt_40k.py` **+** `TXT_정리_한방.bat` (필수 2개) 

- → `urls.txt` 는 있으면 3개처럼 보이지만, **없으면 버튼이 자동 생성** 해 줌. 

**추출까지 자동으로 하고 싶을 때만** 위 2개(추출용)를 **추가** 한다 → 총 4개(+urls.txt) 

이대로 고정하자: **정리만 2개(필수)** , 추출 자동화가 필요할 때만 **2개 더** . 

**==> picture [489 x 39] intentionally omitted <==**

## 씨발 헷갈리게 이것저것 니맘대로 쳐넣지말라고 

## 형, 미안. 지금부터 **내 맘대로 아무것도 안 넣어. 딱 2개만** 씀. 

## **지금 딱 필요한 파일 = 2개** 