```
# 파일명에서 ID를뽑아매핑
    id_to_path = {}
for p in txt_files:
        vid = video_id_from_filename(p.name)
        key = vid if vid else p.name  # ID가없으면파일명기준
# 같은 key가여러개면먼저온걸우선(덮어쓰지않음)
if key notin id_to_path:
            id_to_path[key] = p
# 3) 합칠순서정하기
    ordered_paths = []
    used = set()
# URL 순서우선반영
for vid in url_ids:
if vid in id_to_path and vid notin used:
            ordered_paths.append(id_to_path[vid])
            used.add(vid)
```

```
# 나머지는이름순으로뒤에붙임
for key, p insorted(id_to_path.items(), key=lambda kv: kv[0].lower()):
if key notin used:
            ordered_paths.append(p)
            used.add(key)
```

```
# 4) 큰텍스트만들기(라벨없이내용만이어붙임)
    parts = []
for p in ordered_paths:
        parts.append(p.read_text(encoding="utf-8", errors="ignore"))
        parts.append("\n\n")  # 문단공백만추가
    combined_text = "".join(parts).strip()
```

```
# 5) 4만자분할 + 저장 (폴더명_01.txt …)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

254/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    base_name = re.sub(r'[\\/:*?"<>|]+', '_', root.name).strip() or"OUTPUT"
    out_dir = root / "FINAL_TXT"
    out_dir.mkdir(parents=True, exist_ok=True)
```

```
    chunks = split_by_chars(combined_text, size=args.chunk_size)
for i, c inenumerate(chunks, 1):
        (out_dir / f"{base_name}_{i:02d}.txt").write_text(c, encoding="utf-8")
# 6) 원본영상별 TXT 삭제
for p in txt_files:
try:
            p.unlink()
except Exception:
pass
```

```
print(f"[완료] 최종파일: {out_dir}\\{base_name}_01.txt, _02.txt ... (각{args.
print("        원본영상별 TXT는삭제되었고, 합본파일은따로만들지않았습니다.")
if __name__ == "__main__":
    main()
```

## **C단계｜쓰는 순서 (진짜 간단)** 

**1.** `yotube_script` (=OUT) 안에, **영상별 TXT** 가 예전처럼 이미 있어야 해. 

**2.** 같은 폴더에 `Script_txt.py` 와 `Script_txt.bat` 를 둔다. 

**3.** `Script_txt.bat` **더블클릭** → 끝. 

   - (선택) `yotube_script\urls.txt` 가 있으면, **그 순서** 대로 합친다. 

**4.** 결과: 

markdown 코드 복사 

```
yotube_script\
  FINAL_TXT\
    yotube_script_01.txt
    yotube_script_02.txt
    ...
```

원본 TXT는 삭제됨, **최종본만 남음** . 

필요하면 **분할 크기(기본 40,000자)** 를 바꿔서 `60,000자` 처럼 지정해주는 **배치(버튼)** 버 전도 바로 만들어줄게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

255/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**==> picture [489 x 32] intentionally omitted <==**

야 잠깐만! 내가 기존에 있는 방법그대로 하라고 했지? Script_txt.bat을 누르면 물어보라고 했잖아? 지정폴더를 물어보라고 했잖아 그리고 urls.txt 파일도 생성해 

9s 동안 생각함 

## 오케이 형, 요구한 대로 **버튼을 누르면 ① 정리할 폴더를 물어보고 ② 그 폴더에** 

`urls.txt` **도 자동 생성** 하게 바꿨다. (지금은 “설명+코드”만 제공, 자동 실행 없음) 

## **A단계｜뭐가 달라졌나** 

- [사실] `Script_txt.bat` 를 더블클릭하면 **“정리할 폴더 경로”**를 먼저 물어본다. 

- [사실] 지정한 폴더에 `urls.txt` **가 없으면 자동으로 만들어** 준다. (원하면 바로 메모 장으로 열어 수정 가능) 

- [사실] 그다음 `Script_txt.py` 를 **그 폴더 기준으로 실행** 해 최종 TXT만 남긴다. 

## **B2단계｜파일 2개만 준비** 

## **1)** `Script_txt.py` **(이전 답변에서 준 그대로 사용)** 

- 기능: **모으기 → 합치기(라벨 없음) → 4만자 분할 → 원본 TXT 삭제** 

- 파일 위치: 어디에 둬도 되지만 보통 **버튼과 같은 폴더** 에 둔다. 

## **2) 새 배치:** `Script_txt.bat` **(아래 내용으로 교체 저장)** 

메모장에 아래를 **그대로** 붙여넣고, `Script_txt.bat` 로 저장해. 

bat 코드 복사 

```
@echo off
```

```
chcp 65001 > nul
setlocal
```

```
title Script_txt — 폴더지정 → urls.txt 생성 → 정리(모으기/합치기/4만자분할/원본삭제
```

```
REM 0) 파이썬확인
```

```
where python >nul 2>nul || (
```

```
  echo [오류] Python이필요합니다. 설치후다시실행하세요.
```

```
  pause & exit /b
```

```
)
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 256/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
REM 1) Script_txt.py 확인 (배치와같은폴더에두는걸권장)
```

```
set "HERE=%~dp0"
```

```
if not exist "%HERE%Script_txt.py" (
```

```
  echo [오류] Script_txt.py 가이배치와같은폴더에없습니다.
  echo 같은폴더에두고다시실행하세요.
  pause & exit /b
)
```

```
REM 2) 정리할폴더(OUT) 물어보기
```

```
set "ROOT="
```

```
set /p ROOT=정리할폴더경로를입력하세요(비우면현재폴더사용):
```

```
if "%ROOT%"=="" set "ROOT=%HERE%"
```

```
REM 입력경로정리(따옴표제거및공백처리)
```

```
set "ROOT=%ROOT:"=%"
```

```
if not exist "%ROOT%" (
  echo [안내] 폴더가없어새로만듭니다: "%ROOT%"
  md "%ROOT%" >nul 2>nul
```

```
)
```

```
REM 3) urls.txt 자동생성(없으면)
```

```
if not exist "%ROOT%\urls.txt" (
```

```
  echo [안내] "%ROOT%\urls.txt" 파일이없어새로만듭니다.
```

- `"%ROOT%\urls.txt" echo # 한 줄에 한 개의 유튜브 링크를 적으세요. (재생목록도 가능) >> "%ROOT%\urls.txt" echo # 예시:` 

- `>> "%ROOT%\urls.txt" echo # https://youtu.be/AAAAAAAAAAA` 

- `>> "%ROOT%\urls.txt" echo # https://www.youtube.com/playlist?list=PLXXXXXXXXXXXX` 

- `>> "%ROOT%\urls.txt" echo #` 

- `>> "%ROOT%\urls.txt" echo # 비워 두면 이름순으로 합쳐집니다.` 

- `echo [안내] urls.txt를 지금 편집하시겠습니까? (Y/N, 기본 N) set "ANS="` 

```
  set /p ANS=입력:
  if /I "%ANS%"=="Y" start notepad "%ROOT%\urls.txt"
)
```

```
REM 4) 정리실행 (urls.txt 유무와상관없이순서파일을함께전달해도안전)
echo.
```

```
echo [실행] python "Script_txt.py" --root "%ROOT%" --urls "%ROOT%\urls.txt"
python "Script_txt.py" --root "%ROOT%" --urls "%ROOT%\urls.txt"
```

```
echo.
```

```
echo [완료] 결과: "%ROOT%\FINAL_TXT\폴더이름_01.txt, _02.txt ..." 만남았습니다.
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

257/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

```
echo        (원본영상별 TXT는삭제되었습니다)
```

 

```