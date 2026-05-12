withopen(path,'r',encoding='utf-8',errors='ignore') as f:
for line in f:
            s=line.strip()
ifnot s: continue
if'-->'in s: continue
if re.fullmatch(r'\d+', s): continue
if re.search(r'\d{2}:\d{2}:\d{2}[,\.]\d{3}', s): continue
''
            s=re.sub(r'<[^>]+>',,s)
''
            s=re.sub(r'[\[\(][^\]\)]{1,20}[\]\)]',,s)
            s=s.strip('-─—·• ')
if s: out.append(s)
return out
```

```
# ---------- 반복구절압축 ----------
defcollapse_repeated_words(text:str)->str:
return re.sub(r'(\b\S+\b)(?:\s+\1\b){1,}', r'\1', text)
```

```
defcollapse_repeated_ngrams(text:str, max_n=8)->str:
    t = text
for n inrange(max_n, 1, -1):
        pattern = r'((?:\b\S+\b(?:\s+|$)){%d})(?:\s*\1){1,}' % n
        t = re.sub(pattern, r'\1', t, flags=re.IGNORECASE)
return t
defanti_echo(text:str, aggressive=False)->(str, dict):
    before_len = len(text)
    t = text
    t = collapse_repeated_words(t)
    t = collapse_repeated_ngrams(t, max_n=8if aggressive else6)
    t = re.sub(r'\s+', ' ', t).strip()
return t, {
"chars_before": before_len,
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

51/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
"chars_after": len(t),
```

```
"reduced": before_len - len(t)
    }
```

```
# ---------- 문장/분할 ----------
defsplit_into_sentences_kr(text):
    text = re.sub(r'([\.!?])(?=[^\s])', r'\1 ', text)
    text = re.sub(r'(다\.)(?=[^\s])', r'\1 ', text)
    parts = re.split(r'(?<=[\.!?])\s+|(?<=다\.)\s+', text)
return [p.strip() for p in parts iflen(p.strip())>1]
```

```
defdedup_preserve_order(items):
    seen=set(); out=[]
for s in items:
        key=re.sub(r'\s+','',s)
if key in seen: continue
        seen.add(key); out.append(s)
return out
defchunk_text(text, chunk_size):
iflen(text)<=chunk_size: return [text]
    chunks=[]; start=0; n=len(text)
while start<n:
        end=min(start+chunk_size, n)
if end<n:
            cuts=[]
for pat in [r'\.', r'\?', r'!', r'다\.']:
                idx=[m.end() for m in re.finditer(pat, text[start:end])]
if idx: cuts.append(max(idx))
if cuts: end=start+max(cuts)
else:
                sp=text.rfind(' ', start, end)
if sp!=-1: end=sp
        chunk=text[start:end].strip()
if chunk: chunks.append(chunk)
        start=end
return chunks
```

```
defbest_text_by_title(cands):
    best=''
for t in cands:
iflen(t)>len(best): best=t
return best
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

52/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
# ---------- 메인 ----------
```

```
defmain():
    parser = argparse.ArgumentParser()
    parser.add_argument("dest"), parser.add_argument("folder_name")
    parser.add_argument("--safe-delete", type=int, default=0)  # 0 보존 / 1 성공시
    parser.add_argument("--report", action="store_true")
    parser.add_argument("--aggressive", action="store_true")   # anti_echo 강도
    args = parser.parse_args()
```

```
    dest = args.dest
    folder_name = args.folder_name
```

```
try:
        vtts = glob.glob(os.path.join(dest,"*.vtt"))
        srts = glob.glob(os.path.join(dest,"*.srt"))
        txts = [p for p in glob.glob(os.path.join(dest,"*.txt"))
ifnot os.path.basename(p).startswith(folder_name+"_")
and os.path.basename(p) != f"{folder_name}_REPORT.txt"
and os.path.basename(p) != f"{folder_name}_TXT.txt"]
```

```
ifnot (vtts or srts or txts):
            log_write(dest,"[정보] 처리할자막파일이없습니다.")
print("[정보] 처리할자막파일이없습니다.")
            sys.exit(0)
```

```
        groups={}; order=[]; per_title_info=[]
        total_raw_chars = 0
```

```
defadd_to_group(path, lines):
nonlocal total_raw_chars
            key=extract_title_key(path)
            clean=clean_lines_common(lines)
            total_raw_chars += len(clean)
if key notin groups:
                groups[key]=[]; order.append(key)
            groups[key].append(clean)
```

```
for p in vtts: add_to_group(p, read_text_from_vtt(p))
for p in srts: add_to_group(p, read_text_from_srt(p))
for p in txts: add_to_group(p, read_text_from_txt(p))
```

```
ifnot groups:
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

53/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
raise RuntimeError("자막에서유효텍스트를추출하지못했습니다.")
```

```
        merged=[]; title_count=0
for k in order:
            cands = groups[k]
            chosen = best_text_by_title(cands)
            title_count += 1
```

```
            per_title_info.append((k, len(cands), max(len(x) for x in cands), len
            merged.append(chosen)
```

```
        joined_all = ' '.join(merged)
```

```
        anti_echo_text, stats = anti_echo(joined_all, aggressive=args.aggressive)
```

```
        sentences = split_into_sentences_kr(anti_echo_text)
        before_sent = len(sentences)
        sentences = dedup_preserve_order(sentences)
        after_sent = len(sentences)
```

```
        final_text = ' '.join(sentences).strip()
        final_text = re.sub(r'\s+',' ', final_text)
```

```
        out_merged = os.path.join(dest, f"{folder_name}_TXT.txt")
withopen(out_merged,"w",encoding="utf-8") as f:
            f.write(final_text)
```

```
        chunks = chunk_text(final_text, CHUNK_SIZE)
        width = max(2, len(str(len(chunks))))
for i,c inenumerate(chunks,1):
```

```
withopen(os.path.join(dest, f"{folder_name}_{str(i).zfill(width)}.txt
                f.write(c)
```

```
# 삭제는옵션에따라
        deleted=[]
if args.safe_delete == 1:
for p in vtts+srts+txts:
try:
if os.path.exists(p):
                        os.remove(p); deleted.append(os.path.basename(p))
except Exception as e:
```

```
                    log_write(dest, f"[경고] 삭제실패: {p} -> {e}")
# 리포트
if args.report:
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3c2a-0bd4-8322-8a7a-205e82101d69 

54/65 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 배치파일 수정 

 

```
            rep = []
            rep.append(f"# REPORT — {folder_name}")
            rep.append("")
            rep.append(f"- Titles grouped: {title_count}")
            rep.append(f"- Raw chars (sum of chosen per title BEFORE anti_echo): ~
            rep.append(f"- Anti-echo reduced: {stats['reduced']} chars (from {stat
            rep.append(f"- Sentences after anti-echo: {before_sent}  -> After ded
            rep.append(f"- Output chunks (40,000 chars): {len(chunks)}")
            rep.append("")
            rep.append("## Per-title (key, files_in_group, max_candidate_chars, c
for k, cnt, mx, ch in per_title_info:
                rep.append(f"- {k} | {cnt} files | max={mx} | chosen={ch}")
if args.safe_delete == 1:
                rep.append("")
                rep.append("## Deleted originals")
if deleted:
for b in deleted: rep.append(f"- {b}")
else:
                    rep.append("- (none)")
withopen(os.path.join(dest, f"{folder_name}_REPORT.txt"), "w", encod
                rf.write('\n'.join(rep))
```

```
print("[완료] 통합및분할을마쳤습니다.")
        sys.exit(0)
```

```
except Exception as e:
try:
            log_write(dest, f"[오류] {repr(e)}")
except: pass