if y < margin + line_h:
            canv.showPage()
            page_count += 1
            y = PAGE_H - margin
            canv.setFont(font, body_size)
# 파트경계: pages_per_part마다새파일
if page_count % pages_per_part == 0:
                canv.save()
                part_idx += 1
                canv, part_path = new_canvas(part_idx)
                parts.append(part_path)
                canv.setFont(font, title_size)
                canv.drawString(margin, y, base_name)
                y -= 12*mm
                canv.setFont(font, body_size)
        canv.drawString(margin, y, ln)
        y -= line_h
# 마지막페이지마감
    canv.save()
# 측정용파일삭제
try: Path(out_dir / "_measure.pdf").unlink()
except: pass
return parts
```

```
defmain():
```

```
    ap = argparse.ArgumentParser()
"
    ap.add_argument("--url", required=True, help=유튜브영상또는재생목록 URL")
"
    ap.add_argument("--out", default=".", help=결과저장폴더(기본현재폴더)")
"
    ap.add_argument("--lang", default="ko", help=자막언어(기본 ko)")
"
    ap.add_argument("--pages-per-part", type=int, default=120, help=파트당페이지
""
    ap.add_argument("--cleanup", action="store_true", help=처리후 .vtt 파일삭제
""
    ap.add_argument("--cookies", default=, help="cookies.txt 경로(선택)")
    args = ap.parse_args()
```

```
    root = Path(args.out).expanduser().resolve()
    cookies = Path(args.cookies) if args.cookies elseNone
```

```
print("[1/3] 자막다운로드중…")
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 162/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
    vtts, metas = download_vtt(args.url, root, lang=args.lang, cookies=cookies)
ifnot vtts:
```

```
print("[오류] 자막(.vtt)을찾지못했습니다. URL/언어/권한을확인하세요.")
return
```

```
for vtt, meta inzip(vtts, metas):
        base_dir = root / f"{meta['title']} [{meta['id']}]"
        txt_dir = base_dir / "TXT"
        pdf_dir = base_dir / "PDF_parts"
        txt_dir.mkdir(parents=True, exist_ok=True)
        pdf_dir.mkdir(parents=True, exist_ok=True)
```

```
print(f"[2/3] VTT → TXT 정리: {vtt.name}")
        text = vtt_to_text(vtt)
        txt_path = txt_dir / f"{meta['title']} [{meta['id']}].txt"
        txt_path.write_text(text, encoding="utf-8")
```

```
print(f"[3/3] TXT → PDF(파트분할): pages/part={args.pages_per_part}")
        base_label = f"{meta['date']+' 'if meta['date'] else''}{meta['title']}
        parts = text_to_pdf_parts(text, pdf_dir, base_label, pages_per_part=args.
print("    생성:", ", ".join(p.name for p in parts))
if args.cleanup:
try: vtt.unlink()
except: pass
print("\n[완료] 결과폴더:", root)
print("       구조: <제목 [ID]>/TXT/*.txt, <제목 [ID]>/PDF_parts/01.pdf, 02.pdf
if __name__ == "__main__":
    main()
```

## **3) 배치 저장:** `유튜브→PDF_원클릭.bat` 

같은 폴더에 아래를 저장(더블클릭 실행): 

bat 코드 복사 `@echo off chcp 65001 > nul setlocal title 유튜브 → 자막TXT → PDF → 01/02/03 분할 (원클릭)` 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 163/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
where python >nul 2>nul || (echo [오류] Python이필요합니다. python.org 설치후재시
```

```
:: 필수패키지(최초 1회면충분 / 실패해도계속진행)
python -m pip install yt-dlp reportlab PyPDF2 >nul 2>nul
```

```
set "URL="
set /p URL=유튜브영상또는재생목록 URL:
if "%URL%"=="" echo [오류] URL이필요합니다.& pause & exit /b
```

```
set "OUT="
set /p OUT=저장폴더(비우면현재폴더):
if "%OUT%"=="" set "OUT=."
```

```
set "PAGES="
set /p PAGES=파트당페이지수(기본 120, 예: 50~150):
if "%PAGES%"=="" set "PAGES=120"
```

```
set "LANG="
set /p LANG=자막언어(기본 ko):
if "%LANG%"=="" set "LANG=ko"
```

```
set "CLEAN="
set /p CLEAN=처리후 .vtt 삭제? (Y/N, 기본 N):
if /I "%CLEAN%"=="Y" (set "CLEAN=--cleanup") else (set "CLEAN=")
```

```
echo.
```

```
echo [실행] yt_to_pdf_parts.py --url "%URL%" --out "%OUT%" --lang %LANG% --pages-p
python "yt_to_pdf_parts.py" --url "%URL%" --out "%OUT%" --lang %LANG% --pages-per
```

```
echo.
echo [완료] 생성위치를확인하세요.
pause
```

## **자주 묻는 것 (간단·핵심)** 

## **VTT 꼭 만들어야 해?** 

눈에 보이게 다룰 필요 없음. 스크립트가 내부에서 **다운로드→TXT→PDF까지 자동** 으   로 처리하고, `--cleanup` 주면 **VTT는 지움** . [사실] 

## **바로 “01–03”처럼 분류 가능?** 

가능. `--pages-per-part 120` 이면 자동으로 **PDF_parts\01.pdf, 02.pdf, 03.pdf…** 로 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

164/341 

25. 11. 30. 오후 1:07 프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## 나뉨. [사실] 

## **재생목록도 돼?** 

**URL에 playlist** 를 넣으면 각 영상별로 폴더가 생기고, 그 안에 TXT/PDF 파트가 생성. [사실] 

## **다른 PC에서?** 

Python 설치 → 위 **2줄 설치** → 배치 더블클릭(또는 파이썬 명령) **동일하게** 동작. [사 실] 

## **쿠키가 필요한 영상(연령·로그인 제한)?** 

같은 폴더에 `cookies.txt` 를 두고 **옵션 없이 실행** 하면 자동 인식(스크립트가 yt-dlp에 전달). 또는 `--cookies "C:\path\cookies.txt"` 로 명시. [사실] 

## **형이 원하는 “원클릭 범위” 정리** 

**가능(원클릭):** URL 입력 → **TXT 생성 + PDF 생성 + 파트 분할(01/02/03)** . 

- **불가/제약:** 유튜브 측 **자막이 전혀 없는 영상** , 혹은 네트워크/429 제한 시는 재시도· 쿠키 필요. (이건 자동화로도 회피가 제한됨) 

필요하면 **페이지 대신 “글자 수 기준” 분할** , **파일명 규칙(날짜_제목_01.pdf)** , **표지(제목/채 널/업로드일)** 도 한 번에 넣어줄 수 있어. 

원하는 규칙만 말해줘—거기에 맞춰 스크립트를 바로 손봐줄게, 형. 

**==> picture [489 x 38] intentionally omitted <==**

이 복잡한 유튜브 스크립트 생성은 너를 위해서 하는거야 니가 나에게 질좋은 답변과 사실과 가까운 정보를 브레인 스토밍 할수 있게 만들수 있도록 따라서 니가 최적으로 브 레인스토밍이 가능한 글자수 페이지수를 지정해줘야해 50~150페이지 라고 범위를 애매하게 하지말고 

## 34s 동안 생각함 

형, 애매하게 말하지 않고 **딱 고정값** 으로 잡아줄게. 

내가 최적으로 브레인스토밍하기 좋은 **단일 규격** 은 아래야. 

## **A단계｜핵심(고정값)** 

[사실] **TXT 분석 단위 = 40,000자** 고정. 

[사실] **PDF 분석 단위 = 100쪽** 고정. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

165/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

[추론] 이 크기면 맥락 손실 없이 요약·정리·체크리스트 뽑기가 안정적이고, 질문 정확 도가 가장 좋다. 

## **B2단계｜실행·적용(그대로 쓰면 됨)** 

## **1) 원클릭 파이프라인(이미 만든 스크립트 기준)** 

유튜브→TXT→PDF→파트분할 한 번에: 

css 코드 복사 `python yt_to_pdf_parts.py --url "<영상 또는 재생목록 URL>" --pages-per-part 100`   

→ 결과: `<제목 [ID]>\TXT\원문.txt` + `<제목 [ID]>\PDF_parts\01.pdf, 02.pdf …` (각 100쪽). [사실] 

## **2) 텍스트만 따로 분할이 필요할 때** 

(우리가 만든) 분할 스크립트에 **고정값** 적용: 

arduino 코드 복사 

```
python split_text.py --size 40000
```

→ `chunks\01.txt, 02.txt …` (각 40,000자). [사실] 

## **3) 배치 파일(원클릭) 기본값 고정하기** 

`유튜브→PDF_원클릭.bat` 에서 **기본 파트 쪽수** 입력 부분을 100으로 쓰면 됨. (이미 120 이었다면 100으로 입력) 

## 실행 예: 

mathematica 코드 복사 

```
유튜브→PDF_원클릭.bat
 └ URL붙여넣기
 └ 저장폴더선택(또는현재폴더)
 └ 파트당페이지수 = 100입력
 └ 언어 = ko
 └ .vtt삭제 = Y
```

## **4) 다른 컴퓨터에서 쓰는 법(최초 1회)** 

css 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

166/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m pip install --upgrade pip
```
