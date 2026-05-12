EXCLUDE_PAT = re.compile(r"(readme|urls\.txt|cookies\.txt|\.srt|\.vtt)$", re.I)
```

## `def reg_font():` 

```
for p in [r"C:\Windows\Fonts\malgun.ttf", r"C:\Windows\Fonts\Malgun.ttf",
r"/System/Library/Fonts/AppleSDGothicNeo.ttc"]:
if os.path.exists(p):
try:
                pdfmetrics.registerFont(TTFont("KR", p)); return"KR"
except Exception: pass
return"Helvetica"
defwrap(text, maxw, c, font, size):
    out=[]
for raw in text.splitlines():
        line=raw.rstrip()
ifnot line: out.append(""); continue
        words=line.split(" "); cur=""
for w in words:
            t=(cur+" "+w).strip() if cur else w
if c.stringWidth(t,font,size)<=maxw: cur=t
else:
if cur: out.append(cur)
if c.stringWidth(w,font,size)>maxw:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

108/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
                    buf=""
for ch in w:
if c.stringWidth(buf+ch,font,size)<=maxw: buf+=ch
else: out.append(buf); buf=ch
                    cur=buf
else: cur=w
if cur: out.append(cur)
return out
```

```
deftxt2pdf(path, font):
    base,_=os.path.splitext(path); pdf=base+".pdf"
    PAGE_W,PAGE_H=A4; m=20*mm; box=PAGE_W-2*m; y=PAGE_H-m
    title=os.path.basename(base)
    title=re.sub(r"\.clean$","",title,flags=re.I).replace("_"," ").strip()
    body=open(path,"r",encoding="utf-8",errors="ignore").read()
    c=canvas.Canvas(pdf,pagesize=A4)
    c.setFont(font,16); c.drawString(m,y,title); y-=12*mm
    c.setFont(font,12)
for ln in wrap(body,box,c,font,12):
if y < m+6*mm: c.showPage(); y=PAGE_H-m; c.setFont(font,12)
        c.drawString(m,y,ln); y-=6*mm
```

```
    c.setFont(font,8)
```

```
    c.drawRightString(PAGE_W-m,m/2,datetime.now().strftime("Generated KST %Y-%m-%d
    c.save()
return pdf
```

```
defmain():
    font=reg_font()
# 1) 우선 *.clean.txt, 없으면모든 .txt(제외규칙적용)
    files=glob.glob("**/*.clean.txt",recursive=True)
ifnot files:
        files=[p for p in glob.glob("**/*.txt",recursive=True) ifnot EXCLUDE_PAT
ifnot files:
print("[오류] 변환할 TXT가없습니다."); return
print(f"[정보] TXT 감지: {len(files)}개")
    pdfs=[]
for i,p inenumerate(sorted(files),1):
try:
print(f"  ({i}/{len(files)}) {p}")
            pdfs.append(txt2pdf(p,font))
except Exception as e:
print(f"  [실패] {p}: {e}")
if pdfs:
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 109/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

```
print("[정보] 통합 PDF 병합중…")
        merger=PdfMerger()
for f in pdfs:
try: merger.append(f)
except Exception as e: print(f"  [건너뜀] {f}: {e}")
        merger.write(MERGED_NAME); merger.close()
print(f"[완료] 개별 PDF + '{MERGED_NAME}' 생성")
else:
print("[오류] 생성된 PDF가없습니다.")
if __name__=="__main__": main()
```

**2.** 실행: nginx 코드 복사 `python -m pip install reportlab PyPDF2 python txt_to_pdf_any.py` 

→ **현재 폴더와 하위 폴더의 모든 .txt** (README/urls 등 제외)를 PDF로 만들고, **ALL_LECTURES.pdf** 도 생성. [사실] 

## **어디에 TXT가 있는지 모를 때(빠른 탐색)** 

bash 코드 복사 `cd C:\Users\엔오원\Desktop\youtube_script_extractor dir /s *.txt` 

보이는 경로로 `cd "경로"` 이동해서 위 **방법 2** 실행하면 한 번에 끝난다. 

어느 폴더에 TXT가 있는지 한 줄만 알려줘도, 거기에 맞춰 **완성 명령** 바로 찍어줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

니가 허용가능한 pdf는 몇 페이지의 분량까지 허용가능하 지? 

형, 결론만 말할게. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

110/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **A단계｜핵심** 

- [사실] 우리가 쓰는 **ReportLab/PyPDF2는 페이지 ‘하드 제한’이 없어** . 텍스트만이라 면 **수천~수만 페이지** 도 생성 가능해. 

- [추론] 다만 **PC 메모리·용량·뷰어 성능** 이 실질 제한이야. 보통 **5,000~10,000쪽 (50~200MB)** 까지는 무리 없이 쓰고, 그 이상은 열기/검색이 느려질 수 있어. 

- [권장됨] 실사용은 **권당 300~500쪽(또는 100MB 이하)** 로 **분권** 하는 걸 추천. 

## **B2｜운영 팁** 

- [권장됨] 강의별 TXT를 **개별 PDF + 통합본** 두 가지로 만들고, 통합본은 **권 수를 나눠** 저장(예: `ALL_LECTURES_01.pdf` , `..._02.pdf` ). 

- [권장됨] 이후 필요한 범위만 열어 쓰면 속도/안정성이 좋아져. 

원하면 지금 스크립트에 “N쪽마다 자동 분권” 옵션 넣어줄게. “분권=500쪽” 처럼 숫자만 알려줘, 형. 

**==> picture [489 x 38] intentionally omitted <==**

생성이 아니라 니가 인식하고 나에게 답해줄수 있는 범위 를 말하는거야 

이작업을 하는 이유는 ai를 활용하는데 아무것도 모르기 때문에 

보다 효율적으로 ai를 이용하기 위한 가이드 라인을 잡기 위해서야 

20s 동안 생각함 

형, 질문 요지가 **“내가(어시스턴트)가 한 번에 이해·처리해 줄 수 있는 PDF 분량 한계”** 맞 지? 결론만 현실적으로 잡아줄게. 

## **A단계｜핵심 범위(실무 기준)** 

- [사실] **짧을수록 정확도↑** . 한 번에 건네는 길이가 길수록 요약·검색 정확도가 떨어진 다. 

- [추론] **안정 구간:** “텍스트 기준” **약 5만~10만 자(대략 80~160쪽)** 정도면 한 번에 분 석·요약·질의 응답이 비교적 안정적이다. 그 이상은 **분할 권장** . 

- [권장됨] **대용량(수백~수천 쪽)** 은 **주제/단원 단위로 50~150쪽씩 쪼개서** 순차 질의 가 가장 효율적이다. 

## **B2단계｜가이드라인(형이 그대로 쓰면 됨)** 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

111/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## **1. 분할 원칙** 

[사실] **한 파일 ≤ 150쪽 / 10–20MB 이하** 로 쪼개 주면 응답 품질이 좋다. 

- [권장됨] 강의 기준(에피소드·챕터·플레이리스트)으로 **논리 단위 분할** → 질문도 그 단위로. 

## **2. 질문 템플릿(복붙용)** 

- “파일 A(1–3장)에서 **핵심 개념 5가지** 만 추려줘. [사실/추론/가정 라벨] 붙여.” 

“파일 B(4–6장)에서 **실습 단계** 만 목록화하고, **명령어/단축키** 는 별표로 표시.” 

- “파일 C에서 **헷갈리기 쉬운 개념 3쌍** 을 비교표로 만들어줘(정의/차이/실수 사례/ 체크포인트).” 

- “A·B·C를 통합해 **학습 로드맵(1주/2주/4주)** 초안을 만들어줘.” 

## **3. 전처리 팁(형 환경에 맞춰)** 

이미 만든 `*.clean.txt` → **PDF 변환 전에도 그대로 질문 가능** . 텍스트가면 분석 정확도↑. 

표·코드가 많은 자료는 **PDF + 원문 TXT를 함께 제공** 하면 좋다(문단·코드 인식 정확도 보완). 

## **4. 폴더/파일 규칙(질문 호환 최적화)** 

```
오토코더_파이썬/01_기초/01_환경설정.clean.txt
```

`오토코더_파이썬/02_문법/01_변수.clean.txt` … 

- [권장됨] **챕터당 하나씩** 질문: “02_문법 폴더 전체 핵심 10문장” → “03_자료구조 실습만 단계별로” 처럼. 

## **5. 초장문 자료 다루기(200쪽↑)** 

[권장됨] ① **목차만 먼저 보내고** “어떤 단원부터 요약?” 묻기 → ② 우선순위 단원 부터 50–150쪽씩 투입. 

- [사실] 한 번에 모든 걸 넣으면 **요약은 된다** 해도, 세부 질의 정밀도가 떨어진다. **작게, 빈번히** 가 정답. 

## **형에게 딱 맞춘 운영 루틴(복사해 써)** 

**1.** `urls.txt → *.clean.txt` 까지 만들기(지금까지 한 방식). 

**2. 단원 단위** 로 50–150쪽에 해당하는 `*.clean.txt` 묶음만 올리고 질문: 

   - “이 묶음의 배우·강사 실전 적용 체크리스트 10개만.” 

   - “엔오원 커리큘럼에 맞춰 **원데이/주2회/집중반** 으로 분류해줘.” 

**3.** 정리된 산출을 **엔오원 양식** 으로 변환(필요하면 내가 템플릿 붙여줄게). 

**4.** 다음 묶음으로 반복. 

원하면 내가 **분할 자동화** (예: `N자/문단/장 기준으로 TXT 자동 쪼개기 + 파일명 규칙 적용` ) 스 크립트도 바로 드릴게. “분할 기준=N자 혹은 N쪽”만 말해줘. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 112/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**ALL_LECTURES.pdf** PDF 

이미 이렇게 합쳐졌어 그렇다면 어떻게 이걸 다시 분류해 서 실무기준으로 맞출수 있지? 

형, **이미 합친** `ALL_LECTURES.pdf` **를 “실무 기준”으로 다시 분류하는 법** 을 두 가지로 딱 정 리해줄게. 

(업로드된 파일 확인: `ALL_LECTURES.pdf` ALL_LECTURES ) 

## **A단계｜전체 개요** 

- [사실] **가장 빠른 길** : 합치기 전 원본 `*.clean.txt` 들이 남아있다면, 그걸 **자동 분류 → 다시 PDF** 로 만드는 게 정확도·속도 모두 최고. 

- [추론] **PDF만 남았어도 가능** : `ALL_LECTURES.pdf` 를 **제목/구간 기준으로 자동 분할 → 폴더 규칙에 맞춰 재배치** 할 수 있어(완전 자동 90% + 눈으로 확인 10% 권장). 

## **B2단계｜실행 절차(둘 중 하나 선택)** 

## **방법 1) 원본 TXT가 남아있다 → “정답 루트”(권장)** 

## **1. 자동 정리(채널/날짜/제목 기준)** 

작업 폴더에서: 

nginx 코드 복사 

```
python organize_outputs.py
```

→ `out/채널명/날짜_제목 [ID]/...clean.txt` 구조로 이동 정리. [사실] 

## **2. (선택) 주제 폴더 사본까지** 

같은 폴더에 `topics.json` 만들고: 

css 

코드 복사 

```
python organize_outputs.py--topics
```

- → `out_by_topic/주제/…clean.txt` 사본 생성. [사실] 

## **3. PDF 재생성(개별 + 통합)** 

폴더별로 들어가서(또는 루트에서 한 번에): 
