python "%ROOT%subs_cleaner_plus.py" --no-srt
```

```
echo.
echo [DONE] 결과폴더: %DEST%
echo       *.clean.txt 파일만사용하시면됩니다.
popd
pause
```

## **2. 실행 요령** 

꼭 **“재생목록 URL”**을 넣어줘야 해. (채널 홈(@autocoderIT) 말고, “파이썬” **플 레이리스트 페이지** 주소) 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 97/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

주소는 꺾쇠 없이, 그대로 혹은 따옴표로: 

```
https://www.youtube.com/playlist?list=PLxxxxxxxxxxxx
```

**3. 바로 테스트** (재생목록 URL이 아직 애매하면) 

위 배치 대신, 검은창에서 다음 한 줄로 **먼저 점검** : 

arduino 코드 복사 

`python -m yt_dlp --flat-playlist -i --print "%(webpage_url)s" "재생목록URL`   → 영상 주소들이 주르륵 나오면 OK → 배치 실행해도 100% 동작. [사실] 

## **안 되면 여기만 체크** 

- **“You must provide at least one URL.”** → URL 칸이 비었거나 **채널 URL** 을 넣음. **재 생목록 URL** 로 교체. [사실] 

- **‘… is not recognized …’** → yt-dlp 미설치/경로 문제. 이 배치는 자동으로 `python -m yt_dlp` 를 쓰니 보통 해결됨. 그래도 안 되면: 

nginx 코드 복사 

```
python -m pip install yt-dlp
```

**출력은 되는데 TXT 없음** → `subs_cleaner_plus.py` 가 폴더에 없을 수 있음. 같은 폴 더에 두고 다시 실행. [사실] 

원하면 **“채널 URL + 키워드(파이썬)”만 넣으면 해당 재생목록을 자동 탐색** 하는 버전으로 도 만들어줄게. 

지금은 위 **SAFE2 배치** 로 재생목록 URL로 먼저 확인하자. 출력(마지막 10줄)만 보내줘도 바로 이어서 잡아줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

98/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**==> picture [289 x 170] intentionally omitted <==**

**==> picture [159 x 28] intentionally omitted <==**

**----- Start of picture text -----**<br>
어디에 재생목록이 있지?<br>**----- End of picture text -----**<br>


## 형, 지금 올린 화면이 바로 **재생목록 페이지** 야. 

주소창에 보이는 링크가 **플레이리스트 URL** 이고, 왼쪽 큰 카드가 제목 **“파이썬(Python)”** 재생목록이야. [사실] 

## **A단계｜핵심** 

**주소창(URL)** = 재생목록 주소 → 이걸 복사해서 배치에 넣으면 된다. [사실] 

## **B2단계｜한 걸음씩** 

**1.** 유튜브에서 **오토코더 채널 → ‘재생목록’ 탭** 으로 이동. 

**2. “파이썬(Python)”** 카드(형 화면의 **왼쪽 큰 카드** )를 클릭. 

**3. 브라우저 맨 위 주소창** 에 있는 링크가 재생목록 URL이야. 

   - 모양 예: 

arduino 

코드 복사 

```
https://www.youtube.com/playlist?list=PL........
```

## 이걸 **그대로 복사** (Ctrl+C). 

**4.** 우리 만든 배치에 붙여넣기: 

`재생목록→TXT_SAFE2.bat` 실행 → URL 칸에 **따옴표 없이** 붙여넣고 Enter. 

- 폴더명 입력 → Enter → 같은 폴더에 * **.clean.txt** 생성. [사실] 

## **빠른 확인(리스트만 미리 보기)** 

검은창에서 아래 한 줄로 **URL이 맞는지** 체크도 가능: 

arduino 

코드 복사 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

99/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

```
python -m yt_dlp --flat-playlist -i --print "%(webpage_url)s""여기에_재생목록_URL"
```

  

→ 동영상 URL들이 주르르 뜨면 OK, 배치도 그대로 동작. [사실] 

필요하면 형이 지금 보고 있는 **정확한 주소** 를 복사해서 보내줘. 거기에 맞춰 명령어를 완 성형으로 바로 찍어줄게. 

**==> picture [489 x 39] intentionally omitted <==**

https://www.youtube.com/playlist?list=PLDAF4b3hGbHI3Emxb1MU6AvdomysgoQa 

## 45s 동안 생각함 

## 형, 좋아. **지금 플레이리스트 URL** 잘 받았어: 

arduino 

코드 복사 

```
https://www.youtube.com/playlist?list=PLDAF4b3hGb-HI3Emxb1MU6AvdomysgoQa
```

아래대로 하면 그 재생목록의 **모든 영상 → 한글 자동자막 → 깔끔한 TXT( `… clean.txt` )**까지 자동으로 나온다. 

## **A단계｜핵심** 

- 우리가 만든 `유튜브List_프롬프트_폴더.bat` 로 진행 → 폴더명 먼저 정하고 → URL 붙 여넣고 → **빈 줄(Enter)** 로 시작. [사실] 

## **B2단계｜한 걸음씩 (가장 쉬운 방법)** 

**1.** 작업 폴더( `subs_cleaner_plus.py` 있는 곳 )에서 `유튜브List_프롬프트_폴더.bat` **더블클릭.** 

**2. 폴더명 입력** (예: `오토코더_파이썬_재생목록` ) → Enter. → 이 폴더가 자동 생성된다. 

**3.** 아래 **플레이리스트 URL을 그대로 붙여넣고 Enter** : 

arduino 

코드 복사 

```
https://www.youtube.com/playlist?list=PLDAF4b3hGb-HI3Emxb1MU6AvdomysgoQa
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

100/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

**4. 아무것도 입력하지 않고 한 번 더 Enter** (빈 줄) → 다운로드/변환 시작. 

      - 자막(.vtt/.srt)은 중간 산출물이고, 끝에 `…clean.txt` 만 쓰면 된다. [사실] 

**5.** 완료 메시지에 표시된 폴더(방금 정한 폴더)로 들어가서 

   - `*.clean.txt` 파일들을 열어 확인. 

## **만약 배치 없이 “바로 실행”하고 싶으면 (CMD에서 그대로 복붙)** 

작업 폴더( `subs_cleaner_plus.py` 있는 곳 )에서 검은창 열고: 

bat 코드 복사 

```
:: 1) URL 목록뽑기 (urls.txt 생성)
```

```
python -m yt_dlp --flat-playlist -i --print "%(webpage_url)s" "https://www.youtub
```

```
:: 2) 각영상의한글자동자막만다운로드(미디어는받지않음)
for /f "usebackq delims=" %U in ("urls.txt") do python -m yt_dlp --write-auto-sub
```

```
:: 3) TXT만생성(중복제거 + 문장합치기)
python subs_cleaner_plus.py --no-srt
```

  

끝나면 같은 폴더에 `*.clean.txt` 파일들이 생긴다. [사실] 

## **자주 막히는 포인트** 

- **채널 URL(@autocoderIT)** 이 아니라 꼭 **재생목록 URL(playlist?list=…)** 을 넣어야 한 다. [사실] 

- `yt-dlp` 에러가 나면: 

nginx 코드 복사 

```
python -m pip install yt-dlp
```

- 로그인/연령/지역 제한 재생목록이면 **같은 폴더에** `cookies.txt` 넣어두면 배치가 자 동 적용한다. [사실] 

진행하면 뜨는 마지막 몇 줄(또는 폴더에 생긴 파일명)만 보내줘. 거기서 바로 다음 손질 (폴더 규칙/파일명 규칙) 마무리해줄게, 형. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

101/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

나는 이 강의를 모두 한번에 pdf 파일로 바꾸고 싶어 

30s 동안 생각함 

형, **지금 폴더(또는 폴더 전체)**에 있는 `*.clean.txt` 들을 **한 번에 PDF** 로 바꾸고, **모두 합친 통합 PDF** 까지 자동으로 만들자. 

## **A단계｜핵심** 

준비: `reportlab` (PDF 생성), `PyPDF2` (병합) 설치 → 파이썬 스크립트 1개 실행. [사실] 결과: 각 TXT마다 **개별 PDF** + 전부 합친 **ALL_LECTURES.pdf** 생성. [사실] 

## **B2단계｜손잡고 그대로** 

## **1) 설치(한 번만)** 

검은창(cmd)을 `*.clean.txt` **가 모여있는 폴더** 에서 열고 아래를 붙여넣기: 

css 코드 복사 

```
python -m pip install --upgrade pip
python -m pip install reportlab PyPDF2
```

[사실] 

## **2) 파이썬 파일 만들기** 

메모장 열고 아래 전부 복사→붙여넣기 → `txt_to_pdf.py` 라는 이름으로 같은 폴더에 저 장. 

python 코드 복사 

```
# txt_to_pdf.py — N·ONE TXT → PDF 변환기 (개별 + 통합)
# 사용: 이파일을 *.clean.txt 파일들이있는루트폴더에두고
#       cmd에서:  python txt_to_pdf.py
# 특징:
#  - 현재폴더및모든하위폴더의 *.clean.txt를자동검색
```

```
#  - 각 TXT 옆에 PDF 생성 + 최상위폴더에 ALL_LECTURES.pdf 병합본생성
```

```
#  - 한글폰트: 윈도우 '맑은고딕' 자동등록 (없으면기본폰트)
```

```
import os, glob, io, re
from datetime import datetime
from reportlab.lib.pagesizes import A4