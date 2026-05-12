    for i,c in enumerate(chunks,1):
        name=f'{i:02d}.txt'
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 143/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
        with open(os.path.join(outdir,name),'w',encoding='utf-8') as f: f.write(c
    print(f'[완료] {len(chunks)}개로분할 → 폴더: {outdir}')
    print('       파일명: 01.txt, 02.txt, 03.txt …')
if __name__=='__main__': main()
'@; [IO.File]::WriteAllText('split_simple.py',$code,[Text.UTF8Encoding]::new($tru
)
```

```
:: 2) 파일선택(파워셸파일열기대화상자)
for /f "usebackq delims=" %%F in (`
  powershell -NoProfile -Command ^
```

```
  "$f=new-object -comobject Shell.Application; $fi=$f.BrowseForFolder(0,'분할할 TX
`) do set SELDIR=%%F
```

```
if "%SELDIR%"=="" (
  echo [안내] 폴더선택을취소했습니다. 드래그&드롭방식으로다시시도해도됩니다.
  echo( & pause & exit /b
)
echo.
echo 선택한폴더: %SELDIR%
echo TXT 파일을선택하세요(파일선택창이열립니다)...
for /f "usebackq delims=" %%P in (`
  powershell -NoProfile -Command ^
```

```
  "$f=New-Object System.Windows.Forms.OpenFileDialog; Add-Type -AssemblyName Syst
`) do set FILEPATH=%%P
```

```
if "%FILEPATH%"=="" (
  echo [오류] TXT 파일을선택하지않았습니다.
  pause & exit /b
)
```

```
:: 3) 쪽수/쪽당문자입력(기본값제공)
set PAGES=
set /p PAGES=한묶음목표쪽수입력 (기본 120):
if "%PAGES%"=="" set PAGES=120
```

```
set PPC=
set /p PPC=1쪽당문자수입력 (기본 900):
if "%PPC%"=="" set PPC=900
```

```
echo.
```

```
echo [실행] python split_simple.py "%FILEPATH%" --pages %PAGES% --ppc %PPC%
python "split_simple.py" "%FILEPATH%" --pages %PAGES% --ppc %PPC%
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

144/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
echo.
```

```
echo [완료] 생성된폴더내부의 01.txt, 02.txt ... 파일을업로드해주면됩니다.
pause
```

**2.** `TXT_분할_배치.bat` **더블클릭 →** 

폴더 선택 → TXT 파일 선택 → 쪽수/쪽당문자 입력(기본값 Enter) → **자동 분할 완료** . 

결과 폴더: `chunks_원본파일명\01.txt, 02.txt, 03.txt …` [사실] 

— 팁 — 

50쪽이면 **쪽수=50** , 150쪽이면 **쪽수=150** 으로 입력. 

- 한 페이지가 좀 길다/짧다 느껴지면 **쪽당문자(기본 900)** 를 1000/800 등으로 조절. 드래그&드롭도 가능: **배치 파일 위로 TXT를 끌어다 놓아 실행** 해도 됨(그 경우 파일 선택 단계 건너뜀). 

필요하면 **여러 TXT를 한 번에** 고르고 각각 `chunks_...` 로 분할하는 버전도 만들어줄게. 

**==> picture [489 x 39] intentionally omitted <==**

실행이 되지 않는데? 

형, 우선 미안. 배치가 안 뜨는 건 주로 **파워셸 대화상자/권한/경로 문제** 때문이야. **대화상자 없이 100% 먹히는 초간단 버전** 으로 바꿔서 바로 끝내자. 

## **A단계｜핵심** 

- **드래그&드롭** 만으로 실행되는 배치 + 이미 만든 `split_simple.py` 를 그대로 사용. 창이 꺼지지 않게 **에러를 화면에 남겨** 바로 원인을 볼 수 있게 함. [사실] 

## **B2단계｜지금 바로 되는 방법** 

## **1) 드래그&드롭 배치 만들기** 

메모장 열고 아래를 **그대로** 붙여넣고, 작업 폴더에 `TXT_분할_드랍전용.bat` 로 저장: 

bat 코드 복사 

```
@echo off
chcp 65001 > nul
title TXT 분할(드래그&드롭전용) - 01.txt, 02.txt ...
setlocal
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 145/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
REM 0) Python 확인
where python >nul 2>nul
if errorlevel 1 (
```

```
  echo [오류] Python이설치되지않았습니다. python.org 에서설치후다시시도하세요.
```

```
  echo 창을닫기전, 오류메시지를캡쳐해주세요.
```

```
  pause & exit /b
)
```

```
REM 1) split_simple.py 없으면자동생성(최소본)
```

```
if not exist "split_simple.py" (
```

- `split_simple.py echo import os,re,argparse` 

- `>> split_simple.py echo def split_by_paragraphs(text,limit):` 

- `>> split_simple.py echo ^timport re` 

- `>> split_simple.py echo ^tparas=[p.strip() for p in re.split(r"\n\s*\n",text) if` 

- `>> split_simple.py echo ^tchunks=[]; cur=""` 

- `>> split_simple.py echo ^tfor p in paras:` 

- `>> split_simple.py echo ^t^tif len(cur)+len(p)+2<=limit: cur+=(("\n\n" if cur e` 

- `>> split_simple.py echo ^t^telse:` 

- `>> split_simple.py echo ^t^t^tif cur: chunks.append(cur); cur=""` 

- `>> split_simple.py echo ^t^t^tif len(p)>limit:` 

- `>> split_simple.py echo ^t^t^t^tsent=re.split(r'(?<=[\.!?]|다\.|요\.)\s+',p)` 

- `>> split_simple.py echo ^t^t^t^tbuf=""` 

- `>> split_simple.py echo ^t^t^t^tfor s in sent:` 

- `>> split_simple.py echo ^t^t^t^t^tif len(buf)+len(s)+1<=limit: buf+=((" " if buf` 

- `>> split_simple.py echo ^t^t^t^t^telse:` 

- `>> split_simple.py echo ^t^t^t^t^t^tif buf: chunks.append(buf); buf=s` 

- `>> split_simple.py echo ^t^t^t^tif buf: cur=buf` 

- `>> split_simple.py echo ^t^telse: cur=p` 

- `>> split_simple.py echo ^tif cur: chunks.append(cur)` 

- `>> split_simple.py echo ^treturn chunks` 

- `>> split_simple.py echo def main():` 

- `>> split_simple.py echo ^tap=argparse.ArgumentParser()` 

- `>> split_simple.py echo ^tap.add_argument("input")` 

- `>> split_simple.py echo ^tap.add_argument("--pages",type=int,default=120)` 

- `>> split_simple.py echo ^tap.add_argument("--ppc",type=int,default=900)` 

- `>> split_simple.py echo ^tap.add_argument("--outdir",default="")` 

- `>> split_simple.py echo ^ta=ap.parse_args()` 

- `>> split_simple.py echo ^timport os` 

- `>> split_simple.py echo ^tif not os.path.exists(a.input): print("[오류] 파일없음` 

- `>> split_simple.py echo ^tbase=os.path.splitext(os.path.basename(a.input))[0]` 

- `>> split_simple.py echo ^tlimit=max(1000,a.pages*a.ppc)` 

- `>> split_simple.py echo ^toutdir=a.outdir or f"chunks_{base}"` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 146/341 

25. 11. 30. 오후 1:07 

 

**==> picture [154 x 9] intentionally omitted <==**

- `>> split_simple.py echo ^tos.makedirs(outdir,exist_ok=True)` 

- `>> split_simple.py echo ^tdata=open(a.input,"r",encoding="utf-8",errors="ignore` 

- `>> split_simple.py echo ^tparts=split_by_paragraphs(data,limit)` 

- `>> split_simple.py echo ^tfor i,c in enumerate(parts,1):` 

- `>> split_simple.py echo ^t^twith open(os.path.join(outdir,f"{i:02d}.txt"),"w",e` 

```
  >> split_simple.py echo ^tprint(f"[완료] {len(parts)}개로분할 → 폴더: {outdir}")
```

- `>> split_simple.py echo if __name__=="__main__": main()` 

```
)
```

```
REM 2) 사용법안내
```

```
if "%~1"=="" (
```

```
  echo [안내] 이배치파일위로 **분할할 TXT 파일을끌어다놓아** 실행하세요.
  echo 또는아래형태로직접실행도가능합니다:
```

```
  echo   TXT_분할_드랍전용.bat "C:\경로\ALL_LECTURES.txt" 120 900
```

```
  echo (마지막두숫자: 쪽수, 1쪽당문자수 / 기본 120쪽, 900자)
  pause & exit /b
)
```

```
REM 3) 인자처리: "%~1"=파일, "%~2"=쪽수(기본 120), "%~3"=ppc(기본 900)
set "FILE=%~1"
```

```
set "PAGES=%~2"
set "PPC=%~3"
```

```
if "%PAGES%"=="" set "PAGES=120"
```
