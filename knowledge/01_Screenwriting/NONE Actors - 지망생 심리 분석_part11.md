    ]
    channels = [
"인스타하이라이트묶음(안전vs적합, 시선3도) — CTA: ‘예시더보기(의무없음)’"
"피드고정포스트(3컷비교) — CTA: ‘읽기전용 PDF 열기’",
"유튜브쇼츠(12–15초) — 설명란에 PDF 링크",
    ]
# 역할태깅 & 필터
deftag_and_filter(arr, role_prefix):
        body, iso = filter_body_and_isolate(state, [f"• [역할: {role_prefix}] {a}"
        state.expansions.extend(iso)
```

https://chatgpt.com/g/g-p-68b4572c50c081918f5a18bb2ad93cc5-none-actors/c/68b90a57-cc04-832c-ba9c-a95f68318fb8 

46/205 

25. 11. 30. 오후 1:11 

NONE Actors - 지망생 심리 분석 

 

```
return" ".join(body) if body else"• (금칙어격리됨)"
    out = [
```

```
f"<세부주제: {state.topic or'?'}>（핵심= “{core}” / MKT=On / PARTICIPATION
f"- 훅/카피: {tag_and_filter(hooks, '광고카피라이터')}",
f"- 메시지/프레이밍: {tag_and_filter(msg, '브랜드전략가')}",
f"- 포맷/자산: {tag_and_filter(assets, '콘텐츠디자이너')}",
f"- 채널/CTA: {tag_and_filter(channels, '퍼포먼스마케터')}",
"[위험·한계·대안: • 부정톤과다시이탈↑ → ‘조정/차이 1개’ 중심중립톤유지]"
    ]
return"\n".join(out)
```

```
defrender_expansions(state: EngineState) -> str:
ifnot state.expansions:
return""
    s = ["— 확장(옵션) —"]
for e in state.expansions:
# 확장아이템에태그표시
        tag = "[확장·참여의존]"ifany(k in e for k in FORBIDDEN_PARTICIPATION) els
        s.append(f"{tag}{e}")
    s.append("[승인요청:{승인/수정/결합/보류/폐기}]")
return"\n".join(s)
```

```
# -----------------------------
# Engine
# -----------------------------
classOneShotEngine:
def__init__(self):
        self.state = EngineState()
```

```
defprocess(self, user_text: str) -> str:
        st = self.state
        txt = user_text.strip()
```

```
# Triggers / Locks
if m := PATTERNS["A_topic"].search(txt):
            st.topic = m.group("topic").strip()
            st.switches.SUBTOPIC_LOCK = True
            st.switches.BRAINSTORM = "Focus"
"A"
```

```
if PATTERNS["C_fields"].search(txt):
            fields_str = PATTERNS["C_fields"].search(txt).group("fields")
            st.fields = split_fields(fields_str)
```

https://chatgpt.com/g/g-p-68b4572c50c081918f5a18bb2ad93cc5-none-actors/c/68b90a57-cc04-832c-ba9c-a95f68318fb8 

47/205 

25. 11. 30. 오후 1:11 

NONE Actors - 지망생 심리 분석 

 

```
if m := PATTERNS["D_core"].search(txt):
```

```
            st.core = m.group("core").strip()
            st.switches.CORE_LOCK = True
            st.last_phase = "B1"
```

```
if PATTERNS["H_participation"].search(txt):
            st.switches.PARTICIPATION = "On"# 여전히본문금지, 확장격리만허용
if PATTERNS["I_evaluation"].search(txt):
            st.switches.EVALUATION = "On"# 본문금지, 확장격리만허용
```

```
if PATTERNS["J_psy"].search(txt):
```

```
            st.switches.PSY = "On"
```

```
if PATTERNS["K_mkt"].search(txt):
```

```
            st.switches.MKT = "On"
```

```
# Phases rendering
```

```
        out = []
        ts = now_kst_str()
```

```
if st.last_phase == "A"and (st.core isNone):
```

```
            out.append(render_overview(st))
            out.append(render_core_candidates())
return"\n".join(out) + f"\n— {ts}"
```

```
# B1 done → B2P (Psych-first)
```

```
if st.core and st.switches.PSY == "On"and st.last_phase in ("A", "B1", "
            st.last_phase = "B2P"
            out.append(render_psych(st))
```

```
# B2M after B2P if MKT=On
```

```
if st.core and st.switches.MKT == "On":
if st.last_phase != "B2P":
# 심리선행보장
                out.append(render_psych(st))
            st.last_phase = "B2M"
            out.append(render_marketing(st))
```

```
# 확장섹션(격리된금칙어)
```

```
        exp = render_expansions(st)
if exp:
```

https://chatgpt.com/g/g-p-68b4572c50c081918f5a18bb2ad93cc5-none-actors/c/68b90a57-cc04-832c-ba9c-a95f68318fb8 

48/205 

25. 11. 30. 오후 1:11 

NONE Actors - 지망생 심리 분석 

 

```
            out.append(exp)
```

```
return ("\n".join(out) if out else"(출력없음)") + f"\n— {ts}"
```

```
# -----------------------------
# Demo
# -----------------------------
if __name__ == "__main__":
    eng = OneShotEngine()
# 샘플시퀀스
print(eng.process("다음주제는지망생의심리 3번"))
```

```
print(eng.process("핵심: 머무르기(대사만안틀리면된다는인식으로새연기에시간을
print(eng.process("심리해법먼저보고싶어 (심리)"))
print(eng.process("마케팅관점으로전환해줘 (마케팅전환)"))
```

## 원하면 이걸 **웹(React/Next.js) UI** 나 **CLI** 로도 바로 포팅해줄게. 

**==> picture [489 x 39] intentionally omitted <==**

## 이건 어디에 사용할수 있지? 

20s 동안 생각함 

좋아, 요 엔진/프롬포트는 “핵심→심리 해법→마케팅 전환”을 강제하는 **브레인스토밍·콘텐 츠 생성 워크플로** 라서 아래에 쓸 수 있어. 

## **어디에 쓰나 (실전 용도)** 

- **카피·콘텐츠 제작** : 랜딩/인스타/릴스/쇼츠용 훅·카피·자산 초안 자동화(MKT=On, 역할 태그로 바로 분담). 

- **상담/세일즈 문구** : “안전 vs 적합” 등 **프레이밍 문장** 과 FAQ 마이크로카드(PSY=On) 빠르게 뽑기. 

- **교육자료/온보딩** : “참고용·의무 없음” 라벨의 PDF/카드뉴스 골조(참여·평가 요구 금지 규칙 포함). 

- **전략·기획 회의** : 핵심 후보→확정(B1) 산출물로 회의 출발점 정렬, 결정 기록(병합 로 그) 남기기. 

- **브랜드 톤 가이드** : “버림이 아닌 조정” 같은 톤을 일관되게 재사용(NonParticipation/Non-Evaluation 가드레일 유지). 

- **운영 정책 문안** : 부적합 안내/랜딩 경고 배지 등 **정책 문구** 시안 생성. 

https://chatgpt.com/g/g-p-68b4572c50c081918f5a18bb2ad93cc5-none-actors/c/68b90a57-cc04-832c-ba9c-a95f68318fb8 

49/205 

25. 11. 30. 오후 1:11 

NONE Actors - 지망생 심리 분석 

 

## **어디에 붙이나 (통합/배포)** 

- **로컬/CLI** : 제공한 `one_shot_v4_engine.py` 그대로 실행해 프롬프트-인→아웃 텍스트 파이프라인. 

- **슬랙/디스코드 봇** : 특정 트리거(“다음 주제는…”, “핵심: …”, “마케팅관점…”)에 엔진 연 결. 

- **웹 UI(Next.js/React)** : 텍스트 입력→A/B1/B2P/B2M 탭 뷰로 단계별 결과 노출, 역할 태그별 컬럼 정렬. 

- **문서 워크플로** : Notion/Google Docs에서 “복붙 프롬포트” 템플릿으로 팀이 동일 규 칙 하에 생성. 

- **자동화 도구** : Zapier/Make로 키워드 감지→엔진 호출→결과를 Slack 채널/노션 DB에 적재. 

- **API 마이크로서비스** : FastAPI 래핑해 `POST /process` 로 조직 내 어떤 툴에서든 호출. 

## **누구에게 유용한가 (역할 매핑)** 

- **광고카피라이터/SNS매니저/UX라이터** : 훅·CTA·마이크로카피 즉시 생산(MKT 역할 태 그). 

- **브랜드전략가/기획자** : 프레임·메시지 룰(PSY)을 먼저 고정한 뒤 자산으로 전환. **상담/입학 운영** : 참여/평가 요구 금지된 문구만 나오므로 리스크 최소화. 

## **빠른 사용 시나리오 (3단 락)** 

**1. B1 핵심 확정** : `핵심: 머무르기(대사만 안 틀리면 된다는 인식)` 

**2. B2P 심리 해법** (PSY=On): 프레임전환·인지/손실회피·환경/정책 아이디어만 출력(지망 생 참여/평가 요구 금지). 

**3. B2M 마케팅 전환** (MKT=On): 위 심리 해법을 훅/메시지/자산/채널로 변환, **[역할: …]** 자동 태깅. 

## **제약/주의** 

- 기본값이 **PARTICIPATION=Off / EVALUATION=Off** 라서, 지망생에게 **무언가 하라고 요구하거나 평가 전제** 하는 아이디어는 본문에 못 나오고 ‘확장(옵션)’으로 격리됨. 엔진은 **포맷·가드레일·역할 태깅** 을 담당하고, 실제 퍼블리싱(이미지/영상 제작·배포) 은 별도 툴과 연동 필요. 

- 최신성 검증이 필요한 정보(가격/정책 등)는 별도 소스와 교차검증해 문서화(프롬포 트 규칙에 명시). 

원하면 지금 **B2P(심리 해법)**부터 다시 뽑고, 이어서 **B2M(마케팅 전환)**으로 바로 넘길게. 

**==> picture [489 x 35] intentionally omitted <==**

https://chatgpt.com/g/g-p-68b4572c50c081918f5a18bb2ad93cc5-none-actors/c/68b90a57-cc04-832c-ba9c-a95f68318fb8 

50/205 

25. 11. 30. 오후 1:11 

NONE Actors - 지망생 심리 분석 

 

내용을 바꾸면 사업아이템 혹은 무엇인가 전문적으로 사 용하고 싶을때 유용한 프로그램 혹은 앱이 될수 있겠네? 

9s 동안 생각함 

맞아. 이 엔진은 “핵심→심리 해법→마케팅 전환” 파이프라인과 각종 락/필터를 이미 갖고 있어서, 도메인만 갈아끼우면 **전문용 B2B SaaS/내부툴/에이전시용 앱** 으로 충분히 제품화 가능해. 

## **무엇으로 만들면 좋은가** 

- **SaaS 웹앱** : 프로젝트 단위로 핵심 도출→심리 해법→마케팅 자산까지 생성·관리. **팀용 봇** : Slack/Discord/Notion에서 트리거 문구로 A/B1/B2P/B2M 출력. 