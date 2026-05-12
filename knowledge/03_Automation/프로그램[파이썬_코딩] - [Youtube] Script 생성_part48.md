    OUT_02.txt
    OUT_03.txt
    ...
```

- 여기서 `OUT` 은 **작업 폴더 이름** (예: `Ai강의_01.txt` 형태가 될 수 있음). 

- **영상별 TXT** 와 **합쳐진 큰 파일** 은 **자동 삭제** 되어 깔끔하게 **최종본만 남음** . 

- 각 파일은 **4만 글자 기준** 으로 잘려 있어, 바로 나에게 올리면 돼. 

## **(선택) 명령줄로 돌리고 싶다면** 

- `D:\OUT` 폴더 창의 주소창에 `cmd` 입력 → Enter → 아래 한 줄: 

bat 코드 복사 `python finalize_txt_40k.py --root "D:\OUT" urls.txt` 순서를 반영하려면: bat 코드 복사 

`urls.txt` 순서를 반영하려면: 

```
python finalize_txt_40k.py --root "D:\OUT" --urls "D:\OUT\urls.txt"
```

## **자주 막히는 곳(초간단 체크)** 

- **파이썬 없음** : `python -V` 가 안 나오면 설치 필요. 

**정리 버튼이 에러** : `finalize_txt_40k.py` 와 `TXT_정리_한방.bat` 가 **같은 폴더** 

- **(D:\OUT)** 에 있는지 확인. 

**TXT가 없다** : 2단계(영상에서 TXT 뽑기)를 먼저 해야 한다. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

239/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **마지막 요약(딱 3줄)** 

**1.** `D:\OUT` 만들고, **영상별 TXT** 를 평소처럼 만든다. 

**2.** 같은 폴더에 `finalize_txt_40k.py` , `TXT_정리_한방.bat` (그리고 선택: `urls.txt` ) 넣기. 

**3.** `TXT_정리_한방.bat` **더블클릭** → `FINAL_TXT\폴더이름_01/02/03…` **최종본만 남음** . 

**==> picture [489 x 40] intentionally omitted <==**

다시 시작해 씨발 처음부터 파이썬 설치부터 이거 활용하 는 방법 모두 세세하게 만들어 

코드 입력 값까지 첫 시작부터 끝까지 세세하게 설명해 

21s 동안 생각함 

- 형, 알겠어. **처음 설치 → 실행 → 결과** 까지, 초보도 그대로 따라 할 수 있게 “완전 풀코 스”로 정리할게. 

- (지금은 “설명만”. 아무 것도 자동으로 만들지 않아.) 

## **A단계｜전체 개요 (뭘 하게 되는지)** 

- [사실] **1단계: 파이썬 설치 + 도구 설치** → 컴퓨터가 준비됨 

- [사실] **2단계: 영상에서 글(TXT) 뽑기** → 영상마다 TXT가 생김(형이 하던 방식 그대로 or 내가 준 버튼으로) 

- [사실] **3단계: 정리 한 방(더블클릭)** → TXT를 **모으고 → 하나로 합치고 → 4만자씩** `폴더이름_01/02/03…` **으로 나눠서 최종만 남김** 

## **B1｜핵심만 5줄 (명령 포함)** 

**1.** 파이썬 확인 

nginx 코드 복사 `python -V` 필수 도구 설치/업데이트 css 코드 복사 

**2.** 필수 도구 설치/업데이트 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

240/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m pip install --upgrade pip
python -m pip install -U yt-dlp
```

**3.** 작업 폴더 만들기(예: D:\OUT) 

arduino 코드 복사 `mkdir D:\OUT` **4.** (선택) 영상→TXT 뽑기 버튼으로 한꺼번에 scss 코드 복사 `유튜브→TXT_SIMPLE_MULTI.bat  (더블클릭)` 

**5.** 정리 한 방(필수) 

scss 코드 복사 `TXT_정리_한방.bat            (더블클릭)` 

## **B2｜처음부터 끝까지 “세세하게”** 

## **0) 파이썬 설치 (최초 1회)** 

**1. 공식 파이썬(3.10 이상)** 설치 → 설치 화면에서 **“Add python.exe to PATH” 체크** . 

**2.** 설치 뒤 검은창(cmd) 열고 아래 확인: 

nginx 코드 복사 `python -V` → `Python 3.11.x` 처럼 나오면 OK. 

## **pip가 안 된다?** 

“ `python -m pip --version` 입력. 버전이 나오면 정상.” 

- “안 나오면 파이썬 다시 설치하면서 “Add to PATH” 체크.” 

## **1) 필수 도구 설치/업데이트** 

아래 두 줄을 그대로 실행(여러 번 해도 안전): 

css 코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

241/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m pip install --upgrade pip
```

```
python -m pip install -U yt-dlp
```

## **2) 작업 폴더 만들기(예: OUT)** 

아무 폴더나 정해. 예시: `D:\OUT` 

arduino 

코드 복사 

```
mkdir D:\OUT
```

이 **폴더 이름이 최종 파일명 앞부분** 으로 쓰인다. (예: `OUT_01.txt` ) 

## **3) 파일 배치 (두 가지 중 택1)** 

## **방법 A) 형이 하던 방식 그대로 영상별 TXT 만들기** 

평소처럼 유튜브 URL 넣어서, 아래처럼 **영상마다 TXT** 가 생기게 해: 

makefile 코드 복사 

```
D:\OUT\
```

```
제목1 [ABCD1234]\TXT\제목1 [ABCD1234].txt
제목2 [EFGH5678]\TXT\제목2 [EFGH5678].txt
  ...
```

**순서가 중요** 하면 `D:\OUT\urls.txt` 에 **한 줄=한 URL** 로 적어두면 나중에 **그 순서대로** 합친다. 

## **방법 B) (선택) **내가 준 “영상→TXT 뽑기 버튼”**으로 만들기** 

“내가 직접 만들 능력이 없어요. 버튼으로 한 번에 뽑고 싶어요”면 이 방법 사용. 이 방법은 **영상마다** `\TXT\*.txt` **구조** 로 저장되게 해줘서, 4단계 ‘정리 한 방’과 100% 호환됨. 

**▷ 스크립트 1:** `yt_to_txt_simple.py` **(영상/재생목록 →** `제목 [ID]\TXT\원문.txt` **)** 메모장 열고 아래를 **그대로** 복붙 → `D:\OUT\yt_to_txt_simple.py` 로 저장. 

python 코드 복사 

- _`# yt_to_txt_simple.py —`_ `유튜브` _`URL(`_ `단일` _`/`_ `재생목록` _`) →`_ `제목` _`[ID]\TXT\`_ `제목` _`[ID].txt`_ 

- _`#`_ `사용 예` _`) python yt_to_txt_simple.py --url "https://youtu.be/..." --out "D:\OUT"`_ `import re, html, argparse` 

```
from pathlib import Path
```

```
from yt_dlp import YoutubeDL
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 242/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
defsafe(s): return" ".join(re.sub(r'[\\/:*?"<>|]+', '_', s or"").split()).stri
```

```
defvtt_to_text(vtt_path: Path) -> str:
    raw = vtt_path.read_text(encoding="utf-8", errors="ignore")
    raw = re.sub(r'^\ufeff?WEBVTT.*?\n', '', raw, flags=re.S)
    lines=[]
for ln in raw.splitlines():
        s=ln.strip()
ifnot s or"-->"in s or re.fullmatch(r"\d+", s): continue
        s=html.unescape(re.sub(r"<.*?>","",s))
if s: lines.append(s)
    out, prev=[], ""
for s in lines:
if s!=prev: out.append(s)
        prev=s
return"\n".join(out).strip()
```

```
defdownload_txt(url, out_dir: Path, lang="ko", cookies=None, cleanup=True):
    out_dir.mkdir(parents=True, exist_ok=True)
    opts={
"skip_download": True,
"writesubtitles": True,
"writeautomaticsub": True,
"subtitleslangs": [lang],
"subtitlesformat": "vtt",
"outtmpl": str(out_dir/"%(title)s [%(id)s].%(ext)s"),
"paths": {"home": str(out_dir)},
"noplaylist": False,
"quiet": True, "no_warnings": True,
    }
if cookies and Path(cookies).exists():
        opts["cookiefile"]=str(cookies)
    vtts=[]
    metas=[]
with YoutubeDL(opts) as ydl:
        info=ydl.extract_info(url, download=True)
        items=(info.get("entries") ifisinstance(info,dict) elseNone) or [info]
for it in items:
ifnot it: continue
            vid=safe(it.get("id"))
            title=safe(it.get("title"))
            base=f"{title} [{vid}]"
# ko 우선
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 243/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
            cand=list(out_dir.glob(f"{base}.ko.vtt")) orlist(out_dir.glob(f"{base
ifnot cand:
print(f"[경고] 자막없음 → 건너뜀: {base}");
continue
            vtts.append(cand[0]); metas.append((vid,title))
for (vid,title),v inzip(metas,vtts):
        text=vtt_to_text(v)
        base_dir=out_dir/f"{title} [{vid}]"
        (base_dir/"TXT").mkdir(parents=True, exist_ok=True)
        (base_dir/"TXT"/f"{title} [{vid}].txt").write_text(text, encoding="utf-8"
if cleanup:
try: v.unlink()
except: pass
```

```
defmain():
    ap=argparse.ArgumentParser()
    ap.add_argument("--url", required=True)
    ap.add_argument("--out", default=".")
    ap.add_argument("--lang", default="ko")
    ap.add_argument("--cleanup", action="store_true")
    ap.add_argument("--cookies", default="")
    a=ap.parse_args()
    download_txt(a.url, Path(a.out).resolve(), lang=a.lang, cookies=a.cookies orN
print("[완료] 영상→TXT 변환")
if __name__=="__main__":
    main()
```

## **▷ 버튼 1:** `유튜브→TXT_SIMPLE_MULTI.bat` **(여러 URL 한꺼번에 처리)** 

메모장 열고 아래를 **그대로** 복붙 → `D:\OUT\유튜브→TXT_SIMPLE_MULTI.bat` 저장. 

bat 코드 복사 `@echo off chcp 65001 > nul setlocal title 유튜브 → TXT (여러 URL 한꺼번에)` 

```