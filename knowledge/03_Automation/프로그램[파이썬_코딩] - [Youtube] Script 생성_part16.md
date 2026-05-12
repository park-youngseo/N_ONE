```
#   python organize_outputs.py                  # 채널/제목/날짜기준정리
```

```
#   python organize_outputs.py --topics topics.json  # + 주제폴더분류
# 전제:
```

```
#   1) urls.txt 에유튜브 URL이줄마다있음 (방법 C)
```

```
#   2) 같은폴더에 *.clean.txt 가생성되어있음 (subs_cleaner_plus.py --no-srt 후)
```

```
import os, re, json, glob, shutil, subprocess, sys
from datetime import datetime
```

```
DEST_ROOT = "out"# 최종정리폴더
defsafe(name: str) -> str:
    name = re.sub(r'[\\/:*?"<>|]+', '_', name)
    name = re.sub(r'\s+', ' ', name).strip()
return name
defyt_info(url: str, use_cookies: bool):
# yt-dlp 로메타데이터 JSON 받아오기
    cmd = ["yt-dlp", "-J", url]
if use_cookies and os.path.exists("cookies.txt"):
        cmd = ["yt-dlp", "--cookies", "cookies.txt", "-J", url]
try:
        p = subprocess.run(cmd, capture_output=True, text=True, check=False)
if p.returncode != 0:
return {}
import json as _json
        data = _json.loads(p.stdout)
ifisinstance(data, dict) and"entries"in data:
# 재생목록/라이브등일때첫항목
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 82/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
            data = data["entries"][0] if data["entries"] else {}
return {
"id": data.get("id") or"",
"title": data.get("title") or"",
"channel": data.get("channel") or data.get("uploader") or"",
"upload_date": data.get("upload_date") or"",  # 'YYYYMMDD'
        }
except Exception:
return {}
deffmt_date(yyyymmdd: str) -> str:
ifnot yyyymmdd orlen(yyyymmdd) != 8ornot yyyymmdd.isdigit():
return""
try:
        dt = datetime.strptime(yyyymmdd, "%Y%m%d")
return dt.strftime("%Y-%m-%d")  # KST 표기는파일명에그대로사용
except Exception:
return""
```

```
defmove_clean_txt(video):
    vid = video.get("id","")
    title = video.get("title","")
    channel = video.get("channel","")
""
    date_str = fmt_date(video.get("upload_date",))
    channel_dir = safe(channel) or"UnknownChannel"
    title_part = safe(title) or"Untitled"
    id_part = f"[{vid}]" if vid else""
    date_part = f"{date_str}_" if date_str else""
```

```
# 대상폴더: out/채널명/날짜_제목 [ID]/
    dest_dir = os.path.join(DEST_ROOT, channel_dir, f"{date_part}{title_part}{id_
    os.makedirs(dest_dir, exist_ok=True)
# 옮길파일찾기: *[ID]*.clean.txt 우선, 없으면제목으로도탐색
    candidates = []
if vid:
        candidates = glob.glob(f"*{vid}*.clean.txt")
ifnot candidates and title_part:
        candidates = [p for p in glob.glob("*.clean.txt") if title_part.lower() i
    moved = []
for src in candidates:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

83/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
try:
```

```
            dst = os.path.join(dest_dir, os.path.basename(src))
if os.path.abspath(src) != os.path.abspath(dst):
                shutil.move(src, dst)
            moved.append(dst)
except Exception as e:
print(f"  [경고] 이동실패: {src} -> {dest_dir} ({e})")
return dest_dir, moved
deftopic_classify(dest_dir, moved_files, topics):
# topics = {"Acting":["연기","오디션"], "Coding":["코딩","프로그래밍","Cursor"]}
# 파일명/내용에서키워드매칭되면 out_by_topic/주제/ 로도복사
ifnot topics:
return
    TOP_ROOT = "out_by_topic"
    os.makedirs(TOP_ROOT, exist_ok=True)
for f in moved_files:
try:
            txt = open(f, "r", encoding="utf-8", errors="ignore").read()
except Exception:
            txt = ""
        name = os.path.basename(f)
        low = (name + "\n" + txt).lower()
        hits = []
for topic, keys in topics.items():
for kw in keys:
if kw.strip() and kw.lower() in low:
                    hits.append(topic); break
for topic inset(hits):
            tdir = os.path.join(TOP_ROOT, safe(topic))
            os.makedirs(tdir, exist_ok=True)
# 복사(원본은채널폴더에유지)
try:
                shutil.copy2(f, os.path.join(tdir, os.path.basename(f)))
except Exception as e:
print(f"  [경고] 주제복사실패: {topic} ({e})")
defload_topics(path: str):
ifnot path ornot os.path.exists(path):
return {}
try:
        data = json.load(open(path, "r", encoding="utf-8"))
# {"주제":["키워드1","키워드2",...]}
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 84/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
return {str(k): [str(x) for x in v] for k,v in data.items()}
except Exception:
return {}
```

```
defmain():
    use_cookies = os.path.exists("cookies.txt")
    topics = load_topics("topics.json") if"--topics"in sys.argv else {}
# urls.txt 읽기
ifnot os.path.exists("urls.txt"):
print("[오류] urls.txt가없습니다. 방법 C로 URL을먼저입력하세요.")
return
    urls = [ln.strip() for ln inopen("urls.txt","r",encoding="utf-8",errors="igno
ifnot urls:
print("[오류] urls.txt에 URL이없습니다.")
return
print(f"[정보] 총{len(urls)}개 URL 정리시작")
for u in urls:
print(f"\n[메타] {u}")
        info = yt_info(u, use_cookies)
ifnot info:
print("  -> 메타데이터조회실패(쿠키필요/네트워크이슈가능). 그래도파일
# 최소한 ID 추출만시도
            m = re.search(r"(?:v=|youtu\.be/)([A-Za-z0-9_-]{6,})", u)
            vid = m.group(1) if m else""
            info = {"id": vid, "title":"", "channel":"", "upload_date":""}
        dest_dir, moved = move_clean_txt(info)
print(f"  -> 대상폴더: {dest_dir}")
if moved:
print(f"  -> 이동: {len(moved)}개")
else:
print("  -> 이동할 *.clean.txt를못찾음(아직생성되지않았거나이름불일
if moved and topics:
            topic_classify(dest_dir, moved, topics)
```

```
print("\n[완료] out/ 아래로정리되었습니다.")
if topics:
print("[참고] 주제별사본: out_by_topic/ 아래생성")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

85/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
if __name__ == "__main__":
    main()
```

- (선택) 주제 분류를 쓰고 싶다면, 같은 폴더에 `topics.json` 만들고 예시처럼 저장: 

json 코드 복사 `{ "Cursor/코딩": ["cursor", "코딩", "프로그래밍", "단축키"], "연기/오디션": ["연기", "오디션", "독백", "대본"], "홍보/운영": ["광고", "홍보", "학원", "수강"] }` 

## **2) 사용 순서** 

**1. 방법 C** 로 자막 받기 + TXT 만들기(이미 쓰던 흐름): 

vbnet 코드 복사 

```
yt-dlp --write-auto-subs --sub-langs "ko" --skip-download "<URL들>"
python subs_cleaner_plus.py --no-srt
```

→ 폴더에 `…clean.txt` 들이 있음. [사실] 

## **2. 자동 정리 실행** : 

nginx 코드 복사 

```
python organize_outputs.py
```

`out/채널명/날짜_제목 [ID]/…clean.txt` 구조로 이동 정리. [사실] (주제 분류까지 하고 싶으면 `topics.json` 만들어두고) 

css 코드 복사 

```
python organize_outputs.py--topics
```

→ `out_by_topic/주제/…clean.txt` 사본도 생성. [사실] 

## **3) 버튼 한 번(배치파일) — “목록 입력→다운로드→정리” 올인원** 

메모장에 아래를 붙여넣고 `유튜브List_정리.bat` 로 저장: 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

86/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

코드 복사 

bat 

```
@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion
title URL 입력→자막→TXT→정리 (채널/날짜/제목자동폴더)
```

```
echo [안내] URL을줄마다붙여넣고, 빈줄(그냥 Enter) 입력시시작합니다.
if exist urls.txt del /f /q urls.txt
:LOOP
set "U="
set /p U=URL:
if "!U!"=="" goto START
echo !U!>>urls.txt
goto LOOP
:START
set "COOK="
if exist cookies.txt set "COOK=--cookies cookies.txt"
```

```
echo.
echo [1/3] 자막받는중...
for /f "usebackq delims=" %%U in ("urls.txt") do (