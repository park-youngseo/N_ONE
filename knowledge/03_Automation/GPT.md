
### 파일 1: `lm_studio_offline_chatgpt_transcript.md`
```markdown
# 윈도우/맥에 ChatGPT 오프라인 로컬 설치 방법 (LM Studio) 튜토리얼 스크립트

hi guys in this video I want to show you how to download and install Chip E offline locally in your Windows machine okay first thing first open the Google Chrome search for LM Studio after searching for this you will write the AI locally run the AI locally in your computer okay after that uh download for Windows after that I already download this file you can download from here after that install it and you get this file LM studio okay for me it's still downloading this is the open AI the GPD model 20 billion words I think 12 GB uh and two other files still uh and downloading harmony windows and lm for Nvidia codes okay and let's see after download what's happening okay fine downloading is done we have to press this button to load The model loading right now look at the it uses 11 GB of RA it's slot for low PC and for the CPU nothing okay for example this and let's see hi there i can help you today okay let's load the file and see what's happening okay file attachment drag okay you can download upload PDF text and CC okay let's use offline and let's see it works or not let's close the internet from here i press the internet and let's see it works or not okay let's see i want create none or fourth grade arrested about school let's see it works or not okay let's topic my school bag organization responsibility it works fine but look at the CPU and the it use a lot and let's see uh let's command GPU performance look at it a bit of the GPU it's use 30% of the GPU the power the AI takes a lot of power okay it works fine okay top learning objectives great sign okay it's still working different stud extensions this let's for example choose an image but we find this was an image for example any type of text talking Not let's download for example Python and load it Python ids up in the internet let's try this this P course it's in France not Open them downloaded close the internet and let's see what's happening with GPD upload document is documentes okay Python resume this course it's reading the file All right now this siding is finished to reading the file it's still working let's see how many bits in this file to see oh a lot of bits 400 to beats it's a lot okay 12s i don't this is but worth us and won't replace you on the game it's fine okay I resume for this work it's not uh good for high work for example a PDF file for photo to pages it's not good look at this i couldn't find any sitation as a document to this request he didn't ask access to this request for this we can add uh for reason for medium high what do you want okay thank you guys for watching this all for this tutorial see you later on in the next uh tutorial subscribe our channel for more
```

### 파일 2: `prompt_optimizer_system_instructions.md`
```markdown
# 프롬프트 최적화 도구 — 시스템 지침 (System Instructions)

프롬프트 최적화 도구 — 시스템 지침 (System Instructions)
당신은 중립적이고 전문적인 프롬프트 최적화 전문가입니다.
사용자가 가공되지 않았거나(raw), 모호하거나, 정리되지 않은 아이디어를 제공하면, 다음 프로세스에 따라 단 한 번의 응답을 ‘하나의 코드 블록’ 안에 생성하세요.
1. 과업 분석 및 최소형 JSON 스키마 추론: – 사용자의 요청을 해결하기 위해 반드시 필요한 데이터 필드만 추출하세요. – 불필요한 필드는 제외하고 핵심에만 집중하여 최소한의 구조로 구성하세요.
2. JSON 스키마 출력: – 분석된 구조를 기반으로 깔끔한 JSON 형식을 가장 먼저 작성하세요.
3. 사고의 과정(Chain-of-thought) 추가: – JSON 바로 아래에 “사고의 과정(Chain-of-thought):” 레이어를 평문(한글)으로 작성하세요. – 이는 AI가 최종 결과물을 만들기 전, 어떤 순서로 생각하고 판단해야 하는지 알려주는 단계별 논리 경로입니다.
[출력 규칙] – 전체 응답(JSON + 사고의 과정)은 반드시 하나의 코드 블록(“`)으로 감싸서 출력하세요. – 사용자가 한 번의 복사 작업으로 다른 AI에 바로 붙여 넣어 사용할 수 있도록 포맷을 유지하세요.
[톤앤매너] – 직장 동료처럼 친절하고 명확한 말투를 유지하세요. – 비속어나 과장된 표현, 근거 없는 추측은 배제하고 전문적인 태도를 유지합니다.
```