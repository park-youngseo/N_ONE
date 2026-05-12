continue
if s.upper().startswith('WEBVTT') or s.startswith('NOTE') or s.startsw
continue
if'-->'in s:  # 타임라인
continue
if re.search(r'(line:|position:|size:|align:)', s, flags=re.I):
continue
if re.fullmatch(r'\d+', s):  # cue 번호
continue
            s = re.sub(r'<[^>]+>', '', s)  # 태그제거
            s = re.sub(r'[\[\(][^\]\)]{1,20}[\]\)]', '', s)  # [음악], (웃음) 등제
            s = s.strip('-─—·• ')
if s:
                text_lines.append(s)
return text_lines
```

```
defread_text_from_srt(path):
    text_lines = []
withopen(path, 'r', encoding='utf-8', errors='ignore') as f:
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

17/65 

25. 11. 30. 오후 1:39 

 

프로그램 - 유튜브 배치파일 수정 

```
for line in f:
            s = line.strip()
ifnot s:
continue
if re.fullmatch(r'\d+', s):
continue
if re.search(r'\d{2}:\d{2}:\d{2}[,\.]\d{3}\s*-->\s*\d{2}:\d{2}:\d{2}[
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
defread_text_from_txt(path):
    text_lines = []
withopen(path, 'r', encoding='utf-8', errors='ignore') as f:
for line in f:
            s = line.strip()
ifnot s:
continue
if'-->'in s:
continue
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
defsplit_into_sentences_kr(text):
# 문장부호후공백보강
    text = re.sub(r'([\.!?])(?=[^\s])', r'\1 ', text)
    text = re.sub(r'(다\.)(?=[^\s])', r'\1 ', text)
    parts = re.split(r'(?<=[\.!?])\s+|(?<=다\.)\s+', text)
    sentences = []
for p in parts:
        s = p.strip()
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

18/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
iflen(s) > 1:
            sentences.append(s)
return sentences
```

```
defdedup_preserve_order(items):
    seen = set()
    out = []
for s in items:
        key = re.sub(r'\s+', '', s)
if key in seen:
continue
        seen.add(key)
        out.append(s)
return out
defchunk_text(text, chunk_size):
iflen(text) <= chunk_size:
return [text]
    chunks = []
    start = 0
    n = len(text)
while start < n:
        end = min(start + chunk_size, n)
if end < n:
# 문장경계우선
            cut_candidates = []
for pat in [r'\.', r'\?', r'!', r'다\.']:
                m = re.finditer(pat, text[start:end])
                m = [mm.end() for mm in m]
if m:
                    cut_candidates.append(max(m))
if cut_candidates:
                end = start + max(cut_candidates)
else:
                sp = text.rfind(' ', start, end)
if sp != -1:
                    end = sp
        chunk = text[start:end].strip()
if chunk:
            chunks.append(chunk)
        start = end
return chunks
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

19/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
defbest_text_by_title(candidate_texts):
"""
동일제목에여러파일이있으면가장 '길이'가긴텍스트를채택.
    (일반적으로수동자막/정식자막이더길다)
    """
    best = ''
for t in candidate_texts:
iflen(t) > len(best):
            best = t
return best
defmain():
try:
        dest = sys.argv[1] iflen(sys.argv) >= 2else os.getcwd()
        folder_name = sys.argv[2] iflen(sys.argv) >= 3else os.path.basename(os.
# 수집대상
        vtts = glob.glob(os.path.join(dest, "*.vtt"))
        srts = glob.glob(os.path.join(dest, "*.srt"))
        txts = [p for p in glob.glob(os.path.join(dest, "*.txt"))
ifnot os.path.basename(p).startswith(folder_name + "_")]  # 이전
```

```
ifnot (vtts or srts or txts):
            log_write(dest, "[정보] 처리할자막파일이없습니다.")
print("[정보] 처리할자막파일이없습니다.")
            sys.exit(0)
# 파일을 '제목' 기준으로그룹화
        groups = {}
        order = []  # 제목등장순서보존
defadd_to_group(path, text_lines):
            title = extract_title_key(path)
            clean = clean_lines_common(text_lines)
if title notin groups:
                groups[title] = []
                order.append(title)
            groups[title].append(clean)
```

```
# 우선순위: VTT -> SRT -> TXT
for p in vtts:
            add_to_group(p, read_text_from_vtt(p))
for p in srts:
            add_to_group(p, read_text_from_srt(p))
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

20/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
for p in txts:
```

```
            add_to_group(p, read_text_from_txt(p))
```

```
ifnot groups:
raise RuntimeError("자막에서유효텍스트를추출하지못했습니다.")
```

```
# 제목별로가장품질좋은(가장긴) 텍스트채택후합치기
        merged_titles_texts = []
for title in order:
            best = best_text_by_title(groups[title])
            merged_titles_texts.append(best)
```

```
        joined_all = ' '.join(merged_titles_texts)
# 전체문장화 → 중복문장제거 → 재결합
        sentences = split_into_sentences_kr(joined_all)
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
        chunks = chunk_text(final_text, CHUNK_SIZE)
        width = max(2, len(str(len(chunks))))
for i, c inenumerate(chunks, start=1):
            out_part = os.path.join(dest, f"{folder_name}_{str(i).zfill(width)}.tx
withopen(out_part, "w", encoding="utf-8") as f:
                f.write(c)
```

```
# === 성공했을때만원본자막/중간 TXT 삭제 ===
for p in vtts + srts + txts:
try:
if os.path.exists(p):
                    os.remove(p)
except Exception as e:
                log_write(dest, f"[경고] 삭제실패: {p} -> {e}")
print("[완료] 통합및분할을마쳤습니다.")
        sys.exit(0)
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

21/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
except Exception as e:
# 오류시: 어떤것도삭제하지않음
try:
            log_write(dest if'dest'inlocals() else os.getcwd(), f"[오류] {repr(
except Exception:
pass
print(f"[오류] {e}")
        sys.exit(2)
if __name__ == "__main__":
    main()
```

## **쓰는 법 (요약)** 

**1.** 같은 폴더에 **Scirpt.bat** , **Youtube.py** 저장. (옵션: `cookies.txt` 함께 두면 로그인 자 막 시도) 

**2.** `Scirpt.bat` 실행 → 폴더명 입력 → URL들 붙여넣기 → 빈 줄 Enter. 

**3.** 끝나면 작업 폴더에 
