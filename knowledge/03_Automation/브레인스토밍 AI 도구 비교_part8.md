## **트랙 A: 노코드 MVP (빠른 적용, 유지 쉬움)** 

## **A-1. 목표** 

- 형이 말한 **3모드 사고 시스템** (Claude/ Gemini/ 통합)을 챗봇에 태우기 

- **카카오 채널·인스타 DM 유입 → 답변/가이드 → 내부 시트 기록** 자동화 스케줄·커리큘럼·홍보 카피 초안까지 즉시 뽑기 

## **A-2. 스택** 

**자동화** : Make(=Integromat) 또는 Zapier 

- **LLM** : ChatGPT (GPT-5 Thinking) 

- **DB** : Google Sheets(초기) → Notion DB(확장) 

- **인터페이스** : 웹 위젯(typedream/chatbase/landbot 중 택1) + 카카오 i 오픈빌더(후속) 

- **파일** : 구글 드라이브(콘텐츠/템플릿 저장) 

## **A-3. 플로우** 

**1.** 사용자가 질문/요청 입력 

## **2. 분류 라우팅** (규칙): 

키워드에 따라 **모드 자동 선택** 

   - 감정·심리·서사 → `/Claude` 

   - 통계·트렌드·리서치 → `/Gemini` 

   - 전략·통합·실행 설계 → `/ChatGPT` 

**3.** 선택된 모드 프롬프트로 LLM 호출 → 응답 생성 

**4.** 응답 + 메타데이터(이름/채널/의도/태그) **Sheets/Notion에 로깅** 

**5.** 필요 시 **후속 액션** : 스케줄 제안, 카드뉴스 문안, 체크리스트, To-Do 생성 

## **A-4. 라우팅 규칙(간단)** 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

40/72 

25. 11. 30. 오후 1:02 

 

브레인스토밍 AI 도구 비교 

- “심리/감정/의미/서사/톤/캐릭터/관계/연민/자기합리화” → `/Claude` 

- “데이터/검색/통계/트렌드/키워드/시장/사례” → `/Gemini` 

- “전략/구조/실행/로드맵/템플릿/체계/파이프라인” → `/ChatGPT` 

## **A-5. 모드별 시스템 프롬프트(복사해서 바로 쓰면 됨)** 

## **/Claude** 

- “감정·의미·윤리 중심. 연민과 통찰을 통해 원인→내면→사회구조→인간적 결론 순서. 문장은 간결, 단문. 과장 금지.” 

## **/Gemini** 

- “사실·데이터 중심. 근거 3개 이상, 최근성 우선. 용어 정의→수치→출처 요약. 불 확실은 명시.” 

## **/ChatGPT (통합)** 

- “전략·구조·실행. 문제→원인→전략옵션→선택기준→실행체크리스트→리스크→ 측정지표 순서. 한국어 쉬운 말.” 

## **고정 가드레일** (세 모드 공통): 

- ““모든 날짜·시간은 KST 기준. 불확실 시 ‘추정/확실치 않음’ 표기. 6개 고정 시선 (심리/전략/교육/연출/수익/행동패턴) 중 2개 이상 반영.”” 

## **A-6. 노코드로 바로 만드는 화면(샘플 슬롯)** 

**/start** : 기능 목록 + 모드 버튼 3개 

- **/schedule** : “이번 주 레슨 안내 자동 생성” 

- **/content** : “각성형 카드뉴스 5안” 

- **/curriculum** : “수강생 레벨별 피드백 루브릭” 

- **/promo** : “카카오 채널 공지 멘트 초안” 

## **트랙 B: 프로덕션(코드) — 카카오 채널/웹 챗 완전 연동** 

## **B-1. 목표** 

- **카카오 i 오픈빌더 + 웹훅** 으로 우리 백엔드 연결 

- **FastAPI** 서버에서 모드 라우팅·컨텍스트 관리·로그 저장 

- **Notion/DB** 에 영구 기록, 대시보드 

## **B-2. 아키텍처** 

## **Frontend** : 

카카오 채널 챗(오픈빌더), 웹 챗(Widget/Next.js) 

- **Backend** : FastAPI(Python) 

- **LLM** : OpenAI Chat Completions (모드별 system prompt) 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

41/72 

25. 11. 30. 오후 1:02 

브레인스토밍 AI 도구 비교 

 

- **DB** : PostgreSQL(운영), Notion(운영+가시화) 

- **Queue** : Celery(리서치/긴 작업 비동기) 

- **Storage** : GDrive/Cloud(Storage) 

- **Monitoring** : Sentry + 로그 대시보드(Superset/Metabase) 

## **B-3. 데이터 스키마(간단)** 

```
users(id, name, channel, tags)
```

- `sessions(id, user_id, mode, started_at)` 

- `messages(id, session_id, role, text, meta)` 

- `tasks(id, type, payload_json, status, created_at)` 

- `assets(id, type, url, owner_id)` 

## **B-4. 라우팅 로직(요지)** 

**1.** 인입 메시지 → 의도 분류(규칙 + 소형 LLM) 

**2.** 모드 결정( `/Claude` / `/Gemini` / `/ChatGPT` ) 

**3.** 세션 컨텍스트 불러와 system prompt 구성 

**4.** OpenAI 호출 → 응답 

**5.** 후처리(체크리스트/일정/파일 생성) → 응답 + 로깅 

## **B-5. 예시 API(요약 코드)** 

python 

코드 복사 

```
# FastAPI 요약
from fastapi import FastAPI
from pydantic import BaseModel
import openai
```

```
app = FastAPI()
```

```
classIn(BaseModel):
    user_id: str
    message: str
```

```
defchoose_mode(text:str)->str:
    t=text.lower()
ifany(k in t for k in ["감정","의미","서사","연민","관계","자기합리화"]):
return"claude"
ifany(k in t for k in ["데이터","통계","검색","트렌드","키워드","시장"]):
return"gemini"
return"chatgpt"
```

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

42/72 

25. 11. 30. 오후 1:02 

브레인스토밍 AI 도구 비교 

 

```
defsystem_prompt(mode:str)->str:
if mode=="claude":
return"감정·의미·윤리중심... (위 A-5 그대로)"
if mode=="gemini":
return"사실·데이터중심... (위 A-5 그대로)"
return"전략·구조·실행... (위 A-5 그대로)"
```

```
@app.post("/chat")
defchat(inp: In):
    mode = choose_mode(inp.message)
```

```
    sys = system_prompt(mode)
```

```
    msgs = [{"role":"system","content":sys},{"role":"user","content":inp.message}
    resp = openai.chat.completions.create(model="gpt-5-thinking", messages=msgs)
return {"mode":mode, "reply":resp.choices[0].message.content}
```

실제에선 세션/메모리/에러처리/재시도/토큰한도 관리가 추가됨. 

## **공통 모듈: 형 고정 규칙(“엔오원 가드레일”)** 

**시간/스케줄 포맷** : “PM 03-05 메커니즘 Class (TEAM.A)” 형식 

- **톤** : 쉬운 한국어, 과장 금지, 단정 금지(근거 제시) 

- **라벨** : [사실]/[추론]/[가정] 표시(민감 주제 시) 

- **출력 구조 기본형** : 

요약(2줄) → 핵심 체크리스트(최대 7개) → 다음 행동(3개) 

## **지금 바로 쓸 “모드 전환 명령어 세트”** 

- `/Claude` : 감정·의미 중심 분석 모드 

- `/Gemini` : 데이터·검색·트렌드 중심 모드 

- `/ChatGPT` : 전략·실행 통합 모드 

  `/Schedule` : 이번 주 스케줄 생성(고정 포맷) 

- `/Promo` : 카드뉴스/공지 멘트 3안 

- `/Curriculum` : 수강생 레벨 루브릭 템플릿 

- `/Log` : 대화+결정 로그 저장 

https://chatgpt.com/c/68e85323-7dc0-8322-b2b4-0227b9f70055 

43/72 

25. 11. 30. 오후 1:02 

브레인스토밍 AI 도구 비교 

 

## **보안·운영 체크** 

- 개인정보(수강생 DB) **분리 저장** , 최소 접근 원칙 

- 민감 대화/파일 **암호화 링크** 만 공유 

- 카카오 채널 연결 시 **오픈빌더 정책 준수** (홍보 과다·개인정보 수집 고지) 

## **형에게 추천 실행 순서(현실적인 7일)** 

**D1** : 노코드 MVP 배선(웹 폼/시트/LLM 연결) 

**D2** : 모드 프롬프트 고정 & 라우팅 규칙 튜닝 

- **D3** : `/Schedule` · `/Promo` · `/Curriculum` 3개 핵심 기능 완성 

**D4** : 카드뉴스/공지 실전 투입(소량) → 피드백 

**D5** : 로그 분석, 톤·가드레일 미세 조정 

**D6** : 카카오 채널 파일럿(소수 공개) 

**D7** : 유지 루틴 + 다음 스프린트 정의(자동화 확대) 

원하면, **트랙 A(MVP)**용으로 

- Google Sheets 구조(탭/컬럼 이름) 

- Make/Zapier 시나리오 블록 

- 모드 프롬프트 원문(복붙용) 

- 을 한 번에 만들어줄게. 

- 지금 바로 시작할까? 어느 트랙부터 할지만 골라줘 — **A(노코드) / B(프로덕션)** . 

**==> picture [489 x 38] intentionally omitted <==**

우선 니가 할수 있는건 알겠어 우선 내가 해야할일을 우선 순위로 둬야할게 무엇이지? 위에서 내가 해야될일이 많다 고 했는데 
