폴더엔 `data/form_template.csv` (입력 서식), `scripts/` (파이썬 스크립트), `outputs/` (결 과)가 있어. 

핵심 스크립트는 두 개: 

- `voc_cluster.py` : 비슷한 문장끼리 자동으로 묶음(클러스터). 

- `need_tags.py` : 문장에 **태그** 를 붙임(예: 카메라공포, 발성, 호흡, 시선, 표정, 대사암기, 

- 시간압박, 자신감, 진단요구). 

- 결과 파일은 `cluster_assignments.csv` , `cluster_top_terms.csv` , `needs_tagged.csv` , `needs_counts.csv` . 

원하는 단어는 `scripts/keywords_dict.json` 에 추가해서 정확도를 계속 올릴 수 있 어. 

다운로드: n1_needs_kit_v1.zip 

## **2) N·ONE Public Needs Miner v1 — 대중 노출 데이터 기 반** 

이건 **아무도 설문/DM 안 해도** 돌아가게 만든 “공개 데이터 채굴기”야. 씨앗 키워드( `data/seeds.txt` )를 넣으면 이렇게 움직여: 

- `autocomplete_harvester.py` : 구글/유튜브 **자동완성 제안문** 을 대량 수집(사람들이 

- 실제로 치는 문장). 

- `yt_search_and_transcribe.py` : 그 제안으로 유튜브 검색 → **공개 자막** 을 모음. 

- `analyze_needs.py` : 제안문+자막에서 **니즈 태그 집계** + **상위 n-그램** (자주 함께 쓰이 

- 는 단어 조합). 

- 결과 파일은 `suggestions_all.csv` , `transcripts.csv` , `suggestions_tagged.csv` , `transcripts_tagged.csv` , `needs_counts.csv` , `top_ngrams_2.csv` , 

`top_ngrams_3.csv` . 

태그 사전은 여기서도 `scripts/keywords_dict.json` 로 관리해. 필요 단어를 계속 추 가하면 됨. 

다운로드: n1_public_needs_miner_v1.zip 

요약하면, 첫 번째는 **우리 내부 말** 로 니즈를 수치화, 두 번째는 **대중이 공개적으로 남긴 흔적** (자동완성·자막)으로 니즈를 수치화한다 — 둘 다 결과를 **CSV 숫자** 로 뽑아줘서 바로 공략 포인트를 고를 수 있게 만든 구성이다. 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

17/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

야.. 폴더경로 C:\Users\psm23\OneDrive\Desktop n1_public_needs_miner_v1 이걸로바꿔 

11s 동안 생각함 

좋아. 경로를 **C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1** 로 고 정했다. 아래 중 편한 방법으로 쓰면 된다. 

## **1) 통합 배치파일(권장)** 

다운로드: n1_public_needs_winpath_bats.zip 

- 압축 풀고 `run_all_winpath.bat` 더블클릭 

## **2) 명령어 직접 실행** 

bat 코드 복사 

```
REM ① 의존성설치(최초 1회)
```

```
pip install -r "C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\requirem
```

## `REM ② 자동완성 수집` 

```
python "C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\scripts\autocomp
```

## `REM ③ 유튜브 검색+자막` 

```
python "C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\scripts\yt_searc
```

## `REM ④ 분석` 

```
python "C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\scripts\analyze_
```

  

## **3) 단계별 배치파일(원하면 하나씩 실행)** 

`step0_install.bat` → 설치 

- `step1_autocomplete.bat` → 자동완성 수집 

- `step2_yt_transcripts.bat` → 유튜브 검색+자막 

- `step3_analyze.bat` → 분석 

## **출력 위치(고정)** 

```
C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\outputs
```

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

18/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

필요하면 쿼리 개수(30)·영상 수(8) 바꿔서 더 크게 채굴하자. 

'Google'은(는) 내부 또는 외부 명령, 실행할 수 있는 프로 그램, 또는 

배치 파일이 아닙니다. 

'ublic_needs_miner_v1'은(는) 내부 또는 외부 명령, 실행할 수 있는 프로그램, 또는 배치 파일이 아닙니다. 'install'은(는) 내부 또는 외부 명령, 실행할 수 있는 프로그 램, 또는 배치 파일이 아닙니다. python: can't open file 'C:\\scripts\\autocomplete_harvester.py': [Errno 2] No such file or directory '깋+?먮쭑...'은(는) 내부 또는 외부 명령, 실행할 수 있는 프 로그램, 또는 배치 파일이 아닙니다. 'EBASE" 30 8'은(는) 내부 또는 외부 명령, 실행할 수 있는 프로그램, 또는 배치 파일이 아닙니다. '??..'은(는) 내부 또는 외부 명령, 실행할 수 있는 프로그램, 또는 배치 파일이 아닙니다. 'BASEBASE"'은(는) 내부 또는 외부 명령, 실행할 수 있는 프로그램, 또는 배치 파일이 아닙니다. '즺.'은(는) 내부 또는 외부 명령, 실행할 수 있는 프로그램, 또는 배치 파일이 아닙니다. 계속하려면 아무 키나 누르십시오 . . . 

죽여벌라 씨발놈아 제대로 안해? 

형, 에러는 **배치파일 인코딩/경로/변수 확장** 문제야. 내가 **안전한 .CMD** 세트로 다시 만들 어서, 경로 고정·따옴표· `pushd` 까지 깔끔히 처리했다. 이걸로 돌리면 된다. 

**다운로드(수정본):** n1_public_needs_runner_fix.zip 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

19/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

## **꼭 이렇게 해** 

**1.** 압축 풀어서 안에 있는 `RUN_ME.cmd` 만 더블클릭 내부에서 자동으로 

   - **기본 폴더로 이동** : 

   - `C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1` 

   - **pip 설치** (최초 1회만) 

**자동완성 수집 → 유튜브 자막 수집 → 분석** 순서로 실행 

   - 결과: `outputs\` 폴더에 `suggestions_all.csv` , `transcripts.csv` , `needs_counts.csv` 등 생성 

**2.** 여전히 막히면, 단계별로 실행 

`STEP0_INSTALL.cmd` → 의존성 설치 

`STEP1_AUTOCOMPLETE.cmd` → 자동완성 수집 

`STEP2_YT_TRANSCRIPTS.cmd` → 유튜브 자막 수집 

- `STEP3_ANALYZE.cmd` → 분석 

## **왜 전보다 안전하냐** 

**UTF-8 BOM/한글 코드페이지** 로 명령이 깨지는 이슈 방지(표준 .CMD + `pushd` + 완 전한 따옴표). 

`set "BASE=..."` 대신 ** `pushd "<절대경로>"` **로 현재 작업 폴더 고정 → 

- `C:\scripts\...` 류 잘못된 경로 발생 차단. 

각 단계마다 **오류 메시지 명확화** (Python/pip 미설치 여부도 바로 뜸). 

## **빠른 자가 점검(문제 계속 나오면)** 

시작 버튼 → “ **cmd** ” 입력 → 오른쪽 클릭 “관리자 권한으로 실행” 후 아래 명령 3개 확인: 

lua 코드 복사 

```
python --version
pip --version
```

```
dir "C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\scripts"
```

→ 버전이 안 뜨면 Python 설치부터. 

→ 마지막 줄에서 `autocomplete_harvester.py` 파일이 보이면 OK. 

필요하면, `yt_search_and_transcribe.py` 의 숫자(30, 8)를 바꿔서 더 크게/작게 긁게도 바 로 바꿔서 줄 수 있어. 

**==> picture [489 x 39] intentionally omitted <==**

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

20/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

[INFO] Current folder: 

C:\Users\psm23\OneDrive\Desktop\n1_public_needs_ miner_v1 Python 3.11.9 pip 25.3 from 

C:\Users\psm23\AppData\Local\Programs\Python\P ython311\Lib\site-packages\pip (python 3.11) ERROR: Could not open requirements file: [Errno 2] No such file or directory: 'requirements.txt' python: can't open file 

'C:\\Users\\psm23\\OneDrive\\Desktop\\n1_publi c_needs_miner_v1\\scripts\\autocomplete_harvester.p y': [Errno 2] No such file or directory python: can't open file 

'C:\\Users\\psm23\\OneDrive\\Desktop\\n1_publi c_needs_miner_v1\\scripts\\yt_search_and_transcribe.p y': [Errno 2] No such file or directory python: can't open file 

'C:\\Users\\psm23\\OneDrive\\Desktop\\n1_publi c_needs_miner_v1\\scripts\\analyze_needs.py': [Errno 2] No such file or directory 

[DONE] Check outputs folder: 

"C:\Users\psm23\OneDrive\Desktop\n1_public_needs _miner_v1\outputs" 계속하려면 아무 키나 누르십시오 . . . 

## 형, 오류 원인은 **그 폴더가 비어 있어서** 그래. 

지금 경로에 `requirements.txt` 랑 `scripts\*.py` 가 **없으니까** 배치가 전부 실패한 거야. 

바로 고치자 — **정확히 이 순서로** (복붙용): 

## **0) 지금 상태 확인** 

bash 

코드 복사 

```
dir"C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1"
dir"C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\scripts"
```

## → 둘 다 **파일이 안 보이면** = 압축을 **제대로 풀지 않은 상태** . 

https://chatgpt.com/g/g-p-6904d8274c80819180164fe83b2fa49f-maketing/c/6904d887-3c28-8324-b3a4-5c7d856aa039 

21/62 

25. 11. 30. 오후 1:33 

마케팅 - 시스템 부팅 완료 

 

## **1) 본체 압축 다시 내려받아 정확한 위치에 풀기** 

본체: n1_public_needs_miner_v1.zip 

바탕화면(또는 OneDrive\Desktop)에서 **우클릭 → 모두 추출** 

**대상 폴더** 를 꼭 이렇게 지정: 

makefile 코드 복사 `C:\Users\psm23\OneDrive\Desktop\` 추출 후 폴더가 아래처럼 생겨야 정상이야: kotlin 코드 복사 

추출 후 폴더가 아래처럼 생겨야 정상이야: 

```
C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\
  ├─ requirements.txt
  ├─ README_빠른시작.txt
  ├─ data\seeds.txt
  ├─ outputs\  (비어있어도됨)
  └─ scripts\
       autocomplete_harvester.py
       yt_search_and_transcribe.py
       analyze_needs.py
       keywords_dict.json
```

## **2) 경로 고정 러너 사용(내가 만들어 둔 실행기)** 

러너: n1_public_needs_runner_fix.zip 

압축 풀고 `RUN_ME.cmd` 더블클릭 

## **3) 혹시 또 막히면, 단계별로 실행** 

bat 코드 복사 

## `REM (A) 설치` 

```
pip install -r "C:\Users\psm23\OneDrive\Desktop\n1_public_needs_miner_v1\requirem
```

## `REM (B) 자동완성 수집` 

```