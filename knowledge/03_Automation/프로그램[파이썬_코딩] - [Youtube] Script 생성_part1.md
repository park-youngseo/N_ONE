25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

유튜브 스크립트를 추출하는 파이썬 프로그램을 만드려면 어떻게 해야하지? 

나는 코딩을 전혀 할줄몰라 초등학교 수준으로 유튜브 스 크립트 추출을 위한 프로그램 제작을 설명해줘 

17s 동안 생각함 

형, N·ONE Assistant v1.0 모드로 바로 간다. (PARTICIPATION=Off / EVALUATION=Off) 

## **A단계｜전체 개요** 

- 유튜브 **자막(캡션)** 이 공개된 영상이면, 파이썬 라이브러리로 **그 자막을 그대로 받아 서** TXT/SRT로 저장하는 방식이 가장 쉽고 빠름. [사실] 

- 영상에 자막이 없거나 막혀 있으면, **음성 인식(Whisper)** 으로 추출하는 우회 경로가 필요함(속도는 느리지만 대체 가능). [사실] 

- 형이 **코딩 0레벨** 기준이면, “폴더 내려받기 → 더블클릭 → 주소 붙여넣기” 흐름으로 작동하는 왕초보용 패키지를 준비해뒀어. [사실] 

## **B1단계｜핵심 도출 (핵심만)** 

- **방법 1(권장됨)** : 자막 기반 추출 → 가장 간단/안정. [사실] 

- **방법 2(옵션)** : 자막 없음 → Whisper(음성 인식)로 보완. [사실] 

## **B2단계｜세부 브레인스토밍 (왕초보용 제작·사용 순서)** 

## **1) 형을 위한 “바로 쓰는” 폴더 준비됨 (다운로드 후 즉시 실행) [사실]** 

- 내가 형을 위해 **왕초보용 프로젝트 폴더** 를 만들어 뒀어: 

- `youtube_script_extractor` (안에 실행 파일과 코드 모두 포함) 

- `README_KR.txt` : 설치/사용법 한글 안내 

- `run.bat` : (Windows) 더블클릭 실행 

- `run.command` : (Mac) 터미널에서 실행 

- `extract_transcript.py` : 실제 추출 스크립트(한글 주석) 

- `requirements.txt` : 필요한 라이브러리 목록 

- `outputs/…` : 결과 파일이 떨어지는 폴더 

압축 내려받는 링크가 필요하면 말해줘. 바로 내려받기 링크 제공 가능. [추론] 

## **2) 설치 (한 번만) – 초등학생 모드 [사실]** 

**1. 파이썬 설치** (권장: 최신) → 설치 화면에서 **“Add Python to PATH” 체크** 권장됨. 

## **2.** `youtube_script_extractor` 폴더 안에서 **명령 프롬프트/터미널** 열기 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

1/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

Windows: 폴더 주소창에 `cmd` 입력 → Enter 

   - Mac: 터미널 열고 `cd` 로 폴더 이동 

**3.** 아래 명령 한 줄 실행: 

nginx 코드 복사 

```
pip install -r requirements.txt
```

(처음 1번만) 

## **3) 사용 – 진짜 쉬운 순서 [사실]** 

**Windows** : `run.bat` 더블클릭 → 유튜브 **URL 붙여넣기** → Enter 

**Mac** : 

arduino 코드 복사 

```
chmod +x run.command
./run.command
```

완료 후 `outputs` 폴더에 두 가지 파일 생성: 

- `… .txt` : 순수 대사 텍스트 

- `… .srt` : 자막 타임코드 파일(편집·자막툴 호환) 

## **4) 안에서 무슨 일이 일어나나? (이해용) [사실]** 

프로그램이 **영상 ID** 를 뽑고 → **가능한 자막 목록** 을 확인 → 

1순위 **한국어(ko)** , 2순위 **영어(en)** 자막을 찾음 → 없으면 **한국어 번역** 까지 시도 → 자막을 가져와 **TXT + SRT** 로 저장. 

## **5) 제한·예외 (중요 안내) [사실]** 

영상에 **자막이 전혀 없음/비공개/권한 제한** 이면 이 방식으로는 **불가** . 

→ 이때는 **Whisper(음성 인식)** 로 전환(폴더에 `README_Whisper.txt` 안내 포함). 일부 영상은 자동 자막의 품질이 낮을 수 있음(특히 고유명사·전문용어). 

→ 필요 시 수동 교정 권장됨(참고용). 

## **6) (옵션) Whisper 방식 개요 — 자막 없는 영상용 [사실]** 

추가 설치: `ffmpeg` , `yt-dlp` , `faster-whisper` 

- 흐름: **오디오만 다운로드 → 음성 인식 → TXT/SRT 저장** 

- 장점: 자막 없어도 됨 / 단점: **시간 오래 걸림 + PC 성능 의존** 

- 상세는 폴더의 `README_Whisper.txt` 에 요약. 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

2/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

## 형, 바로 써보자: 

- “ **유튜브 URL** ” 하나 복사 → `run.bat` 실행(또는 Mac은 `run.command` ) → 붙여넣기 → **끝** . [사실] 

## 필요하면: 

**다운로드 링크 제공** / **아이콘 더블클릭 전용 런처** / **TXT 대신 워드/자막 합본** 등으로 확장 가능. “확장 승인”이면 추가 모듈 붙여서 정리해줄게. [추론] 

**==> picture [489 x 39] intentionally omitted <==**

## 링크가필요해 

형, 준비 완료. 

## **A단계｜링크** 

- youtube_script_extractor.zip 다운로드 (압축 풀고 바로 사용) [사실] 

## **B1｜사용 순서(초간단)** 

**1.** 압축 해제 → 폴더 들어가서 

**2.** 첫 1회: `pip install -r requirements.txt` 

**3.** 실행: 

   - Windows: `run.bat` 더블클릭 

   - Mac: 터미널에서 

arduino 코드 복사 

```
chmod +x run.command
```

```
./run.command
```

**4.** 유튜브 URL 붙여넣기 → `outputs` 폴더에 `.txt` / `.srt` 생성. [사실] 

**==> picture [289 x 100] intentionally omitted <==**

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

3/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

 

이것 밖에 안나오는데 뭘해 

형, 지금 보이는 파일 구성이 **정상** 이야. `outputs` 폴더는 **첫 추출이 성공하면 자동으로 생 겨** . 지금부터 그대로 따라 해봐. 

## **A단계｜무엇을 하면 되나 (핵심)** 

**1회 설치:** `pip install -r requirements.txt` 를 그 폴더에서 실행. [사실] 

**실행:** Windows는 `run.bat` 더블클릭 → 유튜브 URL 붙여넣기 → 완료되면 `outputs` 폴더 생성. [사실] 

## **B2단계｜손잡고 따라하기(Windows 기준)** 

## **1. 파이썬 설치 확인** 

시작 메뉴 → “cmd” 실행 → 아래 입력 

nginx 코드 복사 

```
python -V
```

버전이 나오면 OK. “인식 안 됨”이면 파이썬 먼저 설치(설치 시 **Add Python to PATH** 체크). [사실] 

## **2. 폴더에서 설치 명령 실행(처음 1번만)** 

파일이 보이는 그 폴더의 주소창에 `cmd` 입력 → Enter 

검은 창이 뜨면 아래 그대로 입력: 

nginx 코드 복사 

```
pip install -r requirements.txt
```

끝날 때까지 기다리기. [사실] 

## **3. 실행** 

같은 폴더에서 `run.bat` 더블클릭 

“영상 URL 또는 영상 ID”라고 뜨면 **유튜브 주소 붙여넣기** → Enter 성공하면 폴더 안에 `outputs` 가 생기고, 그 안에 

`… .txt` (텍스트) 

- `… .srt` (자막) 파일 생성. [사실] 

## **4. 자주 막히는 부분(빠른 해결)** 

`pip` 가 안 된다 → 파이썬 다시 설치(“Add Python to PATH” 체크) 후 창 닫고 **다 시** 2)부터. [사실] 

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

4/341 

25. 11. 30. 오후 1:07 

 

프로그램[파이썬/코딩] - [Youtube] Script 생성 

- “자막이 없다/권한 문제” 메시지 → 해당 영상엔 공개 캡션이 없음. **Whisper 방 식** 으로 해야 함(폴더의 `README_Whisper.txt` 참고). [사실] 

- 영어 자막만 있다 → 자동으로 한국어 번역 시도함. 실패 시 원문으로 저장될 수 있음. [사실] 

맥이라면: 

arduino 코드 복사 

```
chmod +x run.command
./run.command
```

(처음 1회 `pip3 install -r requirements.txt` 실행 필요) [사실] 

필요하면 지금 뜨는 오류 메시지 그대로 붙여줘. 바로 잡아줄게. 

**==> picture [489 x 38] intentionally omitted <==**

pip install -r requirements.txt 이게 어디 있는데? 

형, `pip install -r requirements.txt` 는 **명령어** 야. “어디에 치나요?”가 포인트지. **파일이 보이는 그 폴더에서 명령창을 열고** 치면 돼. 

## **A단계｜핵심** 

- **그 폴더(스크린샷에 보인 곳)에서 명령창을 열고** 아래 줄을 그대로 입력: `pip install -r requirements.txt` → 처음 1번만 실행. [사실] 

## **B2단계｜손잡고 따라하기** 

## **Windows** 

**1.** 폴더 열기 → 상단 **주소창** (경로 보이는 곳)을 클릭 → `cmd` 입력 → **Enter** . 검은 창(명령 프롬프트)이 **그 폴더 위치로** 열린다. [사실] 

   - (대안) 폴더 빈 곳에서 **Shift+마우스 우클릭 → “여기서 PowerShell 창 열기”** 도 가능. 

**2.** 아래 중 하나 입력(순서대로 시도): 

nginx 코드 복사 

```
pip install -r requirements.txt
```

https://chatgpt.com/g/g-p-68fe4fd9d960819191b1253984973ce2-peurogeuraem-paisseon-koding/c/68f87a3b-a034-8322-9f9b-c55f1958ad33 

5/341 

25. 11. 30. 오후 1:07 

프로그램[파이썬/코딩] - [Youtube] Script 생성 
