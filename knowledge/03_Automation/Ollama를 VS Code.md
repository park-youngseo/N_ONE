요청하신 AX(에이전트 경험) 표준에 따라 제공해주신 원본 소스들의 디테일을 100% 보존하여 요약이나 압축 없이 AI가 읽기 편한 지식 베이스 형태의 마크다운(.md) 파일 형식으로 분할 변환해 드립니다. 모든 파일은 300줄 이하로 구성되었습니다.

### 파일 1: `ollama_vscode_gemma4_transcript.md`
```markdown
# VS Code와 Ollama 연결 및 Gemma 4 로컬 실행 튜토리얼 스크립트

hi guys in this tutorial I want to show you how to connect Olama with VS Code and how to use local moders like GMA 4 with VS Code locally uh using a plugin for free Okay let's go Okay first thing first go here to Google Chrome and go to this link and download the VS Code After downloading the VS Code you can install it easily and go here to this OLAM website and go right here to download Click this button if you have Windows Click this button to download Linux or Mac OS Uh after that we have to download and install the GMF4 uh model How to do that we have to go here to this link I want to set in the description of this video in our channel How to download locally GMA 4 with Olama After completing those steps we have turn right here and open both of them Open lama like this and VS code and now I just compile next section Okay mark all Okay don't we need that okay we have to open a folder Let's create new folder right here and name it VS at some point Open a folder desktop Yes And select this folder Okay Trust the order Yes Yes I trust the order And now this welcome The next step is to go here to extensions and search for this extension to install a local model like JA 4 Continue the first one right here Continue operent Install this one Okay Trust publisher and install Yes Now it's installing right now Okay Fine The installation is done After installing this you can get this this extension right here Okay how to connect modus with this continue uh extension Okay Before uh we run I I I want to see you the model running Okay Okay Once open the command prompt as you can see and copy this comment Uh this is my okay as this is my model It works fine For example the tok say hi number Okay Fine Now let's close this and connect it Okay To connect it we have to select add chat model Use local config Okay Let's go here to this local config and let's reload and go here to add chat model For the chat model we have open AI Mr for us is select and the model is auto detected Hit connect Okay As you can see uh this is my config is worked successfully This is is detected the GIMA 4 local model and I have other model is queen for us We have to use this uh GMA 4 Okay You can use it uh like this Ask anything Before you do that we have to go here and add some rules Go here to config your tools Uh screen width was small Let's adjust the screen Make it bigger And this my tool space Okay Read file Make it automatic Create new file Make it automatic Give your cliff Read currently amplifier automatic Request rers school automatic edit existing file automatic and finally single find and replace automatic grip check automatic okay now let's return right here and ask our model okay let's say I want to create to-do list okay and let's send And let's see Okay it started working This index html style CSS and uh script gs Now it's continue with it's creating now the index.html Okay this is my index.html This my file Okay continue creating uh the style CSS It's generating right now Let's close this Let's close this Make the screen bigger Okay fine It's finished creating the sty CSS The body container Now it's creating our script JS okay fine it's generating now I think it's uh finished creating this now this is my three files generate successfully as you can see this index html tss okay and scriptgs okay I'm glad basic structure of HTML adding task is completing okay let's see let's return here for our VS folder and let's see open with the Chrome Okay this is my uh website Add new task for example play and editing Okay it works fine as you can see Okay let's return right here Okay this is all for this tutorial See you later on the next tutorial
```

### 파일 2: `prompt_optimizer_system_instructions.md`
```markdown
# 프롬프트 최적화 도구 — 시스템 지침 (System Instructions)

당신은 중립적이고 전문적인 프롬프트 최적화 전문가입니다. 
사용자가 가공되지 않았거나(raw), 모호하거나, 정리되지 않은 아이디어를 제공하면, 다음 프로세스에 따라 단 한 번의 응답을 ‘하나의 코드 블록’ 안에 생성하세요.

1. 과업 분석 및 최소형 JSON 스키마 추론: – 사용자의 요청을 해결하기 위해 반드시 필요한 데이터 필드만 추출하세요. – 불필요한 필드는 제외하고 핵심에만 집중하여 최소한의 구조로 구성하세요.
2. JSON 스키마 출력: – 분석된 구조를 기반으로 깔끔한 JSON 형식을 가장 먼저 작성하세요.
3. 사고의 과정(Chain-of-thought) 추가: – JSON 바로 아래에 “사고의 과정(Chain-of-thought):” 레이어를 평문(한글)으로 작성하세요. – 이는 AI가 최종 결과물을 만들기 전, 어떤 순서로 생각하고 판단해야 하는지 알려주는 단계별 논리 경로입니다.

[출력 규칙] 
– 전체 응답(JSON + 사고의 과정)은 반드시 하나의 코드 블록(“`)으로 감싸서 출력하세요. 
– 사용자가 한 번의 복사 작업으로 다른 AI에 바로 붙여 넣어 사용할 수 있도록 포맷을 유지하세요.

[톤앤매너] 
– 직장 동료처럼 친절하고 명확한 말투를 유지하세요. 
– 비속어나 과장된 표현, 근거 없는 추측은 배제하고 전문적인 태도를 유지합니다.
```

### 파일 3: `prompt_optimizer_example_pairs.md`
```markdown
# PROMPT OPTIMIZER — EXAMPLE PAIRS (프롬프트 최적화 예시 쌍)비즈니스, 생산성 소프트웨어

이 문서는 AI에게 드럼 학습 및 관련 과업에서 최적화된 출력의 표준을 보여주는 가이드라인입니다.

---
### EXAMPLE ONE (예시 1)

Raw input (가공되지 않은 입력): “드럼을 배우고 싶은데 집에 악기도 없고 완전 처음이에요. 시간이 별로 없는데 어떻게 시작하죠?”

{ “goal”: “create a beginner drum learning plan”, “inputs_needed”: [ “available practice time per week”, “access to drum set or practice pad”, “noise constraints (apartment or house)”, “music style preference” ], “output_format”: [ “weekly rudiment practice schedule”, “essential gear recommendation”, “rhythm milestones”, “first drum-along songs” ] }

사고의 과정 (Chain-of-thought):
1. 드럼은 소음과 장소 제약이 큰 악기이므로 연습 패드나 전자드럼 유무를 먼저 파악함.
2. 실력 향상 이전에 ‘시간’이 가장 큰 제약이므로 실질적인 주간 연습 스케줄을 우선함.
3. 초보자에게는 화려한 세트 연주보다 기본 루디먼트(Rudiment) 연습을 먼저 권장함.
4. 동기 부여를 위해 2주 안에 간단한 8비트 리듬으로 따라 칠 수 있는 곡을 선정함.

---
### EXAMPLE TWO (예시 2)

Raw input (가공되지 않은 입력): “교회 반주나 밴드에서 드럼을 치고 싶은데, 리듬감이 전혀 없어요. 5km 달리기처럼 단계별 계획이 있을까요?”

{ “goal”: “build a systematic drum rhythm training plan”, “inputs_needed”: [ “current sense of rhythm”, “preferred music genre (Worship, Rock, Pop)”, “available training days per week”, “metronome accessibility” ], “output_format”: [ “week by week rhythm progression”, “metronome-based interval training”, “limb independence exercises”, “performance target dates” ] }

사고의 과정 (Chain-of-thought):
1. 리듬감이 부족하다는 점을 고려하여 메트로놈을 활용한 기초 박자 훈련을 최우선으로 함.
2. 달리기 인터벌처럼 ‘느린 템포 정복 후 증속’ 방식을 단계별로 적용함.
3. 드럼의 핵심인 ‘사지 분리(Independence)’ 훈련을 초반에 배치하여 기본기를 다짐.
4. 특정 곡 완주를 목표로 삼아 성취감을 느낄 수 있는 타임라인을 설정함.

---
### EXAMPLE THREE (예시 3)

Raw input (가공되지 않은 입력): “드럼 연주 영상을 찍어서 유튜브를 시작해보고 싶어요. 뭐부터 준비해야 하죠?”TV, 동영상

{ “goal”: “create a drum cover channel launch plan”, “inputs_needed”: [ “recording equipment (mic or electronic drum)”, “video editing skill level”, “target audience or niche”, “budget for audio interface” ], “output_format”: [ “audio/video recording setup guide”, “content structure for drum covers”, “editing software recommendations”, “upload and growth strategy” ] }

사고의 과정 (Chain-of-thought):
1. 드럼은 수음이 까다로우므로 예산에 맞는 마이크 배치나 전자드럼 연결 방식을 먼저 정함.
2. 단순 연주 영상인지 레슨 영상인지 타겟 audience를 명확히 하여 콘텐츠 방향을 잡음.
3. 영상 편집보다 오디오 믹싱이 더 중요한 악기 특성을 고려하여 오디오 설정 가이드를 포함함.
4. 초기 구독자 확보를 위해 대중적인 곡과 매니아틱한 곡의 업로드 비율을 전략적으로 제안함.
```