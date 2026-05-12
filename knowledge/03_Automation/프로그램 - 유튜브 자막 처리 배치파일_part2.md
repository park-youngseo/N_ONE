defsplit_to_chunks(text, base_name, out_dir, chunk_size=CHUNK_SIZE):
    paths = []
iflen(text) == 0:
return paths
    total = (len(text) + chunk_size - 1) // chunk_size
for i inrange(total):
        part = text[i*chunk_size:(i+1)*chunk_size]
        fn = f"{base_name}_{i+1:02d}.txt"
        out_path = os.path.join(out_dir, fn)
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

7/10 

25. 11. 30. 오후 1:39 프로그램 - 유튜브 자막 처리 배치파일 

```
        write_text_file(out_path, part)
        paths.append(out_path)
return paths
```

```
defmain():
iflen(sys.argv) < 3:
print("[오류] 인자가부족합니다. 사용: python Youtube.py \"<DEST>\" \"<FOLDE
        sys.exit(1)
    dest = sys.argv[1]
    folder_name = sys.argv[2]
```

```
ifnot os.path.isdir(dest):
print("[오류] 대상폴더가존재하지않습니다:", dest)
        sys.exit(1)
```

```
# 대상파일수집
    vtt_files = glob.glob(os.path.join(dest, "*.vtt"))
    srt_files = glob.glob(os.path.join(dest, "*.srt"))
    txt_files = glob.glob(os.path.join(dest, "*.txt"))
```

```
# urls.txt 등유틸텍스트는통합대상에서제외할목록
    exclude_exact = { "urls.txt", "cookies.txt" }
```

```
# 통합/분할출력파일명(보존대상)
    combined_name = f"{folder_name}_TXT.txt"
    combined_path = os.path.join(dest, combined_name)
```

```
# 먼저기존통합/분할본은삭제대상에서제외하기위해확보
# (이전실행잔존물이있다면덮어씌울예정)
    keep_prefix = f"{folder_name}_"
    keep_allow = set()
```

```
if os.path.exists(combined_path):
        keep_allow.add(os.path.basename(combined_path))
```

```
for p in glob.glob(os.path.join(dest, f"{keep_prefix}[0-9][0-9].txt")):
```

```
        keep_allow.add(os.path.basename(p))
```

```
# 텍스트통합
```

```
    all_text_parts = []
```

```
# 1) VTT
```

```
for p insorted(vtt_files):
        raw = read_text_file(p)
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

8/10 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

```
        cleaned = strip_tags_and_timestamps(raw, is_vtt=True, is_srt=False)
        merged = normalize_lines_and_merge(cleaned)
if merged:
            all_text_parts.append(merged)
```

```
# 2) SRT
```

```
for p insorted(srt_files):
        raw = read_text_file(p)
        cleaned = strip_tags_and_timestamps(raw, is_vtt=False, is_srt=True)
        merged = normalize_lines_and_merge(cleaned)
if merged:
            all_text_parts.append(merged)
```

```
# 3) TXT (단, 유틸성/출력물은제외)
for p insorted(txt_files):
        base = os.path.basename(p)
if base in exclude_exact:
continue
if base == combined_name:
continue
if re.fullmatch(rf"{re.escape(folder_name)}_\d{{2}}\.txt", base):
# 이전실행의분할본이면패스 (입력으로재흡수방지)
continue
        raw = read_text_file(p)
# TXT도동일한정리루틴(간혹타임스탬프/태그가섞여올수있음)
        cleaned = strip_tags_and_timestamps(raw, is_vtt=False, is_srt=False)
        merged = normalize_lines_and_merge(cleaned)
if merged:
            all_text_parts.append(merged)
    merged_all = ' '.join([t for t in all_text_parts if t]).strip()
    merged_all = re.sub(r'\s+', ' ', merged_all)
```

```
# 통합본작성
    write_text_file(combined_path, merged_all)
```

```
# 4만자분할
```

```
    split_to_chunks(merged_all, folder_name, dest, CHUNK_SIZE)
```

```
# 분할본목록갱신(보존대상)
```

```
    keep_allow.add(os.path.basename(combined_path))
for p in glob.glob(os.path.join(dest, f"{keep_prefix}[0-9][0-9].txt")):
        keep_allow.add(os.path.basename(p))
```

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

9/10 

25. 11. 30. 오후 1:39 

프로그램 - 유튜브 자막 처리 배치파일 

```
# 잔여물정리: 통합/분할본외모든 .vtt/.srt/.txt 삭제
for pattern in ("*.vtt", "*.srt", "*.txt"):
for p in glob.glob(os.path.join(dest, pattern)):
            base = os.path.basename(p)
if base in keep_allow:
continue
if base in exclude_exact:
# 요구사항: "나머지 VTT, TXT 파일삭제" — urls.txt도삭제하는게원문
# 안전을위해여기선남겨두지않고삭제대상에포함시킴. 필요시주석해
pass
try:
                os.remove(p)
except Exception:
# 삭제실패는무시(권한등)
pass
```

```
print("[완료] 통합및분할이끝났습니다.")
print(" - 통합본 :", combined_path)
print(" - 분할본 :", os.path.join(dest, f"{folder_name}_01.txt (이하연속번호)"
if __name__ == "__main__":
    main()
```

## **사용 방법 (요약)** 

**1.** 위 **두 파일** 을 같은 폴더에 저장: 

   - `Scirpt.bat Youtube.py` 

**2.** (선택) 같은 폴더에 `cookies.txt` 가 있으면 자동 사용. 

**3.** `Scirpt.bat` 더블클릭 → 폴더명 입력 → URL들을 줄마다 붙여넣기 → 빈 줄 Enter. 

**4.** 완료 후 결과: 

**==> picture [37 x 48] intentionally omitted <==**

`폴더명_TXT.txt` (전체 통합본) 

- `폴더명_01.txt` , `폴더명_02.txt` … (각 4만자 분할본) 

- 그 외 `.vtt/.srt/.txt` 잔여 파일은 자동 삭제됨. 

원하는 추가 규칙(예: 특정 단어 필터링, 더 강한 문장 병합 규칙 등) 있으면 바로 반영해서 

https://chatgpt.com/g/g-p-68f8d8277a8881919e041e0e2d2a7c9d-peurogeuraem/c/68fe3b72-662c-8324-b7d6-5e405caf2cc5 

10/10 
