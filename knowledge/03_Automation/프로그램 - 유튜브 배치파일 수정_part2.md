if re.fullmatch(r'\d+', s):
continue
if re.search(r'\d{2}:\d{2}:\d{2}[,\.]\d{3}', s):
continue
''
            s = re.sub(r'<[^>]+>', , s)
''
            s = re.sub(r'[\[\(][^\]\)]{1,20}[\]\)]', , s)
            s = s.strip('-─—·• ')
if s:
                text_lines.append(s)
return text_lines
```

```
defnormalize_and_join(lines):
# 줄단위잘림을자연스럽게공백으로연결
    text = ' '.join(lines)
# 다중공백/탭정리
    text = re.sub(r'\s+', ' ', text).strip()
return text
```

```
defsplit_into_sentences_kr(text):
"""
한국어중심간단분리: 문장부호(., ?, !) 및한국어마침(다.)에근거.
문장부호뒤에공백을보장하여분리안정화.
    """
# 문장부호뒤에공백보강
    text = re.sub(r'([\.!?])(?=[^\s])', r'\1 ', text)
# 일반적종결표현들뒤공백
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

7/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
    text = re.sub(r'(다\.)(?=[^\s])', r'\1 ', text)
```

```
# 분리 (구두점/종결부호를보존)
    parts = re.split(r'(?<=[\.!?])\s+|(?<=다\.)\s+', text)
# 잔여소절정리
    sentences = []
for p in parts:
        s = p.strip()
ifnot s:
continue
# 한줄짜리구두점/잡음제거
iflen(s) <= 1:
continue
        sentences.append(s)
return sentences
```

```
defdedup_preserve_order(items):
```

```
    seen = set()
    out = []
for s in items:
        key = re.sub(r'\s+', '', s)  # 공백무시하고비교
if key in seen:
continue
        seen.add(key)
        out.append(s)
return out
```

```
defchunk_text(text, chunk_size, prefer_sentence_boundary=True):
ifnot prefer_sentence_boundary orlen(text) <= chunk_size:
return [text]
    chunks = []
    start = 0
    n = len(text)
while start < n:
        end = min(start + chunk_size, n)
if end < n:
# 문장경계(가장가까운오른쪽구두점)를찾아자연분할
            cut = text.rfind('.', start, end)
            cut_q = text.rfind('?', start, end)
            cut_e = text.rfind('!', start, end)
            cut_d = text.rfind('다.', start, end)
            candidate = max(cut, cut_q, cut_e, cut_d)
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

8/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
if candidate == -1or candidate < start + int(chunk_size*0.7):
# 적절한경계가없으면공백기준
                cut_spc = text.rfind(' ', start, end)
                candidate = cut_spc if cut_spc != -1else end
else:
                candidate += 1# 구두점포함
            end = candidate
        chunk = text[start:end].strip()
if chunk:
            chunks.append(chunk)
        start = end
return chunks
```

```
defmain():
try:
# 인자
        dest = sys.argv[1] iflen(sys.argv) >= 2else os.getcwd()
        folder_name = sys.argv[2] iflen(sys.argv) >= 3else os.path.basename(os.
```

```
# 수집대상확장자
        vtts = glob.glob(os.path.join(dest, "*.vtt"))
        srts = glob.glob(os.path.join(dest, "*.srt"))
        txts = [p for p in glob.glob(os.path.join(dest, "*.txt"))
ifnot os.path.basename(p).startswith(folder_name + "_")]  # 기존
```

```
ifnot (vtts or srts or txts):
```

```
# 자막이없어도종료는정상(0) 처리: 다운로드만한상황일수있음
withopen(os.path.join(dest, "Youtube_log.txt"), "a", encoding="utf-8
                log.write("[정보] 처리할자막파일이없습니다.\n")
print("[정보] 처리할자막파일이없습니다.")
            sys.exit(0)
```

```
        lines = []
```

```
# VTT 우선(보통더깨끗함), 그다음 SRT, 그다음 TXT
for p in vtts:
            lines.extend(read_text_from_vtt(p))
for p in srts:
            lines.extend(read_text_from_srt(p))
for p in txts:
            lines.extend(read_text_from_txt(p))
```

```
ifnot lines:
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

9/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
raise RuntimeError("자막에서유효텍스트를추출하지못했습니다.")
```

```
# 줄합치기 → 문장분리 → 중복제거 → 다시합치기
        joined = normalize_and_join(lines)
        sentences = split_into_sentences_kr(joined)
        sentences = dedup_preserve_order(sentences)
        final_text = ' '.join(sentences).strip()
        final_text = re.sub(r'\s+', ' ', final_text)
```

```
# 통합본저장
        out_merged = os.path.join(dest, f"{folder_name}_TXT.txt")
withopen(out_merged, "w", encoding="utf-8") as f:
            f.write(final_text)
```

```
# 4만자분할저장
        chunks = chunk_text(final_text, CHUNK_SIZE, prefer_sentence_boundary=True
        width = max(2, len(str(len(chunks))))
for i, c inenumerate(chunks, start=1):
            out_part = os.path.join(dest, f"{folder_name}_{str(i).zfill(width)}.tx
withopen(out_part, "w", encoding="utf-8") as f:
                f.write(c)
```

```
# === 여기까지성공시에만원본자막/중간 TXT 삭제 ===
# (요구 4: 통합본이생성되면나머지 VTT/TXT 삭제)
for p in vtts + srts + txts:
try:
# 분할본/통합본은제외규칙상걸리지않음
if os.path.exists(p):
                    os.remove(p)
except Exception as e:
# 삭제실패는치명아님. 로그만.
withopen(os.path.join(dest, "Youtube_log.txt"), "a", encoding="ut
                    log.write(f"[경고] 삭제실패: {p} -> {e}\n")
```

```
print("[완료] 통합및분할을마쳤습니다.")
        sys.exit(0)
```

```
except Exception as e:
# 오류시: 어떤것도삭제하지않도록여기서반환
try:
withopen(os.path.join(dest if'dest'inlocals() else os.getcwd(), "Y
                log.write(f"[오류] {repr(e)}\n")
except Exception:
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

10/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
pass
```

```
print(f"[오류] {e}")
        sys.exit(2)
```

```
if __name__ == "__main__":
    main()
```

## **사용법 (초간단)** 

**1.** 위 **두 파일을 같은 폴더** 에 저장 (이름 고정: `Scirpt.bat` , `Youtube.py` ). 

**2.** (선택) `cookies.txt` 를 같은 폴더에 두면 로그인 필요 채널 자막도 시도. 

**3.** `Scirpt.bat` 더블클릭 → 폴더명 입력 → URL들 붙여넣기 → 빈 줄 Enter. 

**4.** 완료 후 작업폴더에 

`폴더이름_TXT.txt` (통합본), 

- `폴더이름_01.txt` , `_02.txt` … (4만자 분할본)만 남고, 나머지 `.vtt/.srt/.txt` 는 

- 자동 삭제. 

오류 발생 시 **아무 것도 삭제하지 않고** `Youtube_log.txt` 에 이유 기록. 

추가로 바꾸고 싶은 규칙(예: 괄호 태그 보존/삭제, 분할 단위 30,000자 등) 있으면 바로 말 해 줘. 

**==> picture [489 x 38] intentionally omitted <==**

@echo off chcp 65001 > nul setlocal enabledelayedexpansion title 유튜브 목록 처리(폴더 지정) - TXT만 생성 

rem 현재 배치파일이 있는 폴더 경로(절대경로) set "ROOT=%~dp0" 

echo 저장할 폴더 이름을 입력하세요. (빈칸이면 날짜_시 각으로 자동 생성) set "FOLDER=" set /p FOLDER=폴더명: if "%FOLDER%"=="" ( 

for /f %%i in ('powershell -NoProfile -Command "(GetDate).ToString(\"yyyy-MM-dd_HHmm\")"') do set "FOLDER=%%i" 

) 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

11/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 