
**2.** 같은 방식으로 `part002` , `part003` … 반복. 

## **3.** 끝나면 아래로 **요약 합치기(최대 5만 자)** : 

python 코드 복사 `python - << "PY" import glob,re files=sorted(glob.glob("part*_summary.txt")) text="" for f in files: text+=open(f,"r",encoding="utf-8",errors="ignore").read().strip()+"\n\n" text=re.sub(r'\s+\n','\n',text) limit=50000 if len(text)>limit: cut=text[:limit] p=max(cut.rfind('\n\n'),cut.rfind('. '),cut.rfind('다.\n')) text=cut if p<2000 else cut[:p] open("SUMMARY_50k.txt","w",encoding="utf-8").write(text) print("DONE",len(text)) PY` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 138/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

- → `SUMMARY_50k.txt` = 최종 질의단위 통합 요약. [사실] 

## **4) 빠른 품질 체크(있으면 좋음)** 

- [권장됨] 각 파트 요약 받을 때 **누락 금지 키워드** 를 같이 전달: 

- `“아래 키워드는 반드시 다뤄라: Cursor, 단축키, 설정, 오류 3가지, 재시도 요령.”` 

- [권장됨] 파트마다 **3줄 회고** 를 메모: “배울 점·적용할 점·주의점” → 커리큘럼 합치기 쉬움. 

## **바로 실행하려면** 

- 지금 폴더에서 `split_text.py` 가 없으면, 내가 보낸 **분할 스크립트** 부터 저장→실행. 첫 파트( `part001.txt` ) 내용을 붙여주면, **그 자리에서** 위 템플릿대로 요약 시작해서 **샘플 1건** 바로 뽑아줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

뭐가 이렇게 많아.. 간단하게 생성된 텍스트 파일에 대한 내용이 훼손되지 않게 50~150쪽 분량으로 

설정해서 01로 나누고 나머지를 위와 같은 내용으로 02로 분량이 넘어가면 03으로 만들어 업로드 하면 되는거 아 냐? 

25s 동안 생각함 

형, 딱 원하는 방식으로 정리하자. 

**원문 훼손 없이** 텍스트를 **50~150쪽 분량** 으로 잘라 **01.txt, 02.txt, 03.txt…** 만 만들어 주 는 초간단 스크립트야. 

## **A단계｜핵심** 

- [사실] 내용은 그대로 두고 **문단 경계** 를 우선으로 잘라서 손실/요약 없이 분할. 

- [권장됨] “쪽”→문자수 환산은 **1쪽≈900자** 로 계산(필요하면 옵션으로 바꿀 수 있음). 

## **B2단계｜그대로 따라하기** 

## **1) 스크립트 만들기** 

메모장 열고 아래 전체를 붙여넣어 `split_simple.py` 라는 이름으로 저장 (텍스트 파일이 있는 폴더에 저장해). 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 139/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

코드 복사 

python 

 

```
# split_simple.py — 원문훼손없이 50~150쪽단위로 01.txt, 02.txt... 분할
# 사용법예)
#   python split_simple.py "ALL_LECTURES.txt" --pages 120
# 옵션)
#   --pages 120        : 한덩어리목표쪽수(기본 120)
#   --ppc 900          : 1쪽당문자수(Page Per Char, 기본 900)
#   --outdir chunks    : 출력폴더(기본 chunks_<원본파일명>)
import os, re, argparse
```

```
defsplit_by_paragraphs(text, limit):
    paras = [p.strip() for p in re.split(r"\n\s*\n", text) if p.strip()]
    chunks, cur = [], ""
for p in paras:
iflen(cur) + len(p) + 2 <= limit:
            cur += (("\n\n"if cur else"") + p)
else:
if cur:
                chunks.append(cur)
                cur = ""
# 문단하나가 limit보다길면문장기준으로한번더잘라줌
iflen(p) > limit:
# 마침표/물음표/느낌표/‘다.’/‘요.’ 등기준으로안전분할
                sent = re.split(r'(?<=[\.!?]|다\.|요\.)\s+', p)
                buf = ""
for s in sent:
iflen(buf) + len(s) + 1 <= limit:
                        buf += ((" "if buf else"") + s)
else:
if buf:
                            chunks.append(buf)
                        buf = s
if buf:
                    cur = buf
else:
                cur = p
if cur:
        chunks.append(cur)
return chunks
```

```
defmain():
    ap = argparse.ArgumentParser()
""
    ap.add_argument("input", help=분할할 TXT 파일경로)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

140/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
"
    ap.add_argument("--pages", type=int, default=120, help=한묶음목표쪽수(기본
"
    ap.add_argument("--ppc", type=int, default=900, help=쪽당문자수(기본 900)")
"""
    ap.add_argument("--outdir", default=, help=출력폴더명(미지정시자동)")
    args = ap.parse_args()
```

```
    src = args.input
```

```
ifnot os.path.exists(src):
print("[오류] 파일을찾을수없습니다:", src); return
```

```
    base = os.path.splitext(os.path.basename(src))[0]
    limit = max(1000, args.pages * args.ppc)  # 안전하한
    outdir = args.outdir orf"chunks_{base}"
    os.makedirs(outdir, exist_ok=True)
```

```
    text = open(src, "r", encoding="utf-8", errors="ignore").read()
    chunks = split_by_paragraphs(text, limit)
```

```
for i, c inenumerate(chunks, 1):
# 01.txt, 02.txt, ...
        name = f"{i:02d}.txt"
withopen(os.path.join(outdir, name), "w", encoding="utf-8") as f:
            f.write(c)
```

```
print(f"[완료] {len(chunks)}개로분할 → 폴더: {outdir}")
print("       파일명: 01.txt, 02.txt, 03.txt …")
```

```
if __name__ == "__main__":
    main()
```

## **2) 실행 (매번 이거)** 

검은창을 **텍스트 파일이 있는 폴더** 에서 열고, 예를 들어 **120쪽** 단위로 자르려면: 

lua 코드 복사 

```
python split_simple.py "ALL_LECTURES.txt"--pages 120
```

- 결과: `chunks_ALL_LECTURES` 폴더 안에 `01.txt, 02.txt, 03.txt …` 생성. [사실] 

- 50쪽으로 자르고 싶으면 `--pages 50` , 150쪽이면 `--pages 150` . 

- 1쪽 환산이 다르다고 느끼면 `--ppc 1000` 처럼 조정 가능. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

141/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **업로드/질의 안내(형 루틴에 딱 맞게)** 

   - 이제 `01.txt → 질문` **, 끝나면** `02.txt → 질문` **, …** 이렇게 **순서대로 업로드** 하면 돼. 

   - 각 업로드 때 “ **핵심 10~15문장 + 실행 체크리스트** ” 템플릿으로 요청하면 품질 좋고, 나중에 원하면 내가 **요약 합본** 도 만들어줄게. 

- 필요하면 이걸 배치(더블클릭)로 바꿔서 `파일 선택 → 쪽수 입력 → 자동 분할` 까지 만들어 줄게. “배치도”라고 한마디만 해줘, 형. 

**==> picture [489 x 38] intentionally omitted <==**

배치도 

## 10s 동안 생각함 

## 형, **더블클릭 한 번** 으로 

**1.** TXT 파일 선택 → 2) 쪽수 입력 → 3) 자동 분할(01.txt, 02.txt …)까지 끝내는 **배치파일** 준비했어. 

## **A단계｜핵심** 

- 실행하면 **파일 선택 창** 이 뜨고, **쪽수(기본 120쪽)** 입력하면 `chunks_원본파일명 \01.txt, 02.txt …` 자동 생성. [사실] 

## **B2단계｜배치 만들기 & 사용** 

**1.** 메모장 열고 아래를 **그대로** 붙여넣고, 작업 폴더에 `TXT_분할_배치.bat` 로 저장. 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title TXT 분할(50~150쪽) - 01.txt, 02.txt ...
```

```
:: 0) Python 확인
where python >nul 2>nul
if errorlevel 1 (
```

```
  echo [오류] Python이필요합니다. https://www.python.org 에서설치후다시실행하세요
  pause & exit /b
```

```
)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 142/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
:: 1) split_simple.py 없으면자동생성(파워셸사용)
if not exist "split_simple.py" (
  powershell -NoProfile -Command ^
    "$code=@'
import os, re, argparse
def split_by_paragraphs(text, limit):
    paras=[p.strip() for p in re.split(r'\n\s*\n', text) if p.strip()]
    chunks=[]; cur=''
    for p in paras:
        if len(cur)+len(p)+2<=limit:
            cur+=(('\n\n' if cur else '')+p)
        else:
            if cur: chunks.append(cur); cur=''
            if len(p)>limit:
                sent=re.split(r'(?<=[\.!?]|다\.|요\.)\s+', p)
                buf=''
                for s in sent:
                    if len(buf)+len(s)+1<=limit: buf+=((' ' if buf else '')+s)
                    else:
                        if buf: chunks.append(buf)
                        buf=s
                if buf: cur=buf
            else:
                cur=p
    if cur: chunks.append(cur)
    return chunks
def main():
    ap=argparse.ArgumentParser()
    ap.add_argument('input', help='분할할 TXT 경로')
    ap.add_argument('--pages', type=int, default=120, help='한묶음목표쪽수(기본
    ap.add_argument('--ppc', type=int, default=900, help='쪽당문자수(기본 900)')
    ap.add_argument('--outdir', default='', help='출력폴더명(미지정시자동)')
    a=ap.parse_args()
    if not os.path.exists(a.input):
        print('[오류] 파일이없습니다:', a.input); return
    base=os.path.splitext(os.path.basename(a.input))[0]
    limit=max(1000, a.pages*a.ppc)
    outdir=a.outdir or f'chunks_{base}'
    os.makedirs(outdir, exist_ok=True)
    text=open(a.input,'r',encoding='utf-8',errors='ignore').read()
    chunks=split_by_paragraphs(text, limit)