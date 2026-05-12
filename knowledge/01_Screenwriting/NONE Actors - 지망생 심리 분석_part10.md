"연습", "훈련", "제출", "참여", "숙제", "찍어오다", "준비해오다", "반복하라", "과제
```

https://chatgpt.com/g/g-p-68b4572c50c081918f5a18bb2ad93cc5-none-actors/c/68b90a57-cc04-832c-ba9c-a95f68318fb8 

43/205 

25. 11. 30. 오후 1:11 

NONE Actors - 지망생 심리 분석 

 

```
]
```

## `FORBIDDEN_EVALUATION = [` 

```
"평가", "심사", "콜백", "채점", "점수", "리뷰", "블라인드평가", "선호도", "완주율
]
```

```
# -----------------------------
```

## _`# Patterns`_ 

```
# -----------------------------
```

```
PATTERNS = {
```

- `"A_topic": re.compile(r"^다음\s*주제는\s*(?P<topic>[\d가-힣A-Za-z _-]+)", re.M)` 

- `"B_only": re.compile(r"(?<=,|\s)\s*(?P<sub>[\w가-힣]+)\s*만\s*브레인스토밍"),` 

- `"C_fields": re.compile(r"^\[필드\s*지정\]\s*(?P<fields>.+)$", re.M),` 

- `"D_core": re.compile(r"^핵심\s*[:=]\s*(?P<core>.+)$", re.M),` 

- `"E_core_candidates": re.compile(r"^핵심\s*후보", re.M),` 

- `"F_core_confirm": re.compile(r"^핵심\s*확정", re.M),` 

- `"G_core_only": re.compile(r"(?<=,|\s)\s*핵심\s*만\s*브레인스토밍"),` 

- `"H_participation": re.compile(r"(참여|숙제|과제|실습|연습 시켜|훈련 참여|숙제 제안` 

- `"I_evaluation": re.compile(r"(평가|심사|콜백|채점|점수|리뷰|블라인드 평가|선호도|` 

- `"J_psy": re.compile(r"(심리|심리적 해결|심리 해법)"),` 

- `"K_mkt": re.compile(r"(마케팅관점|마케팅 해법|마케팅 전환|마케팅)"),` 

```
}
```

```
# -----------------------------
```

```
# Utils
```

```
# -----------------------------
```

```
defnow_kst_str() -> str:
```

```
return datetime.now(KST).strftime("%Y-%m-%d %H:%M KST")
```

```
defsplit_fields(s: str) -> List[str]:
```

```
    parts = re.split(r"[·\./,\|]", s)
```

```
return [p.strip() for p in parts if p.strip()]
```

## `def role_round_robin(idx: int) -> str:` 

```
return ROLES[idx % len(ROLES)]
```

```
deffilter_body_and_isolate(state: EngineState, bullets: List[str]) -> Tuple[List
```

```
"""본문금지항목은확장으로격리"""
```

```
    body, isolated = [], []
```

```
for b in bullets:
```

```
        forbid = False
```

```
if state.switches.PARTICIPATION == "Off":
```

```
ifany(kw in b for kw in FORBIDDEN_PARTICIPATION):
```

```
                forbid = True
```

https://chatgpt.com/g/g-p-68b4572c50c081918f5a18bb2ad93cc5-none-actors/c/68b90a57-cc04-832c-ba9c-a95f68318fb8 

44/205 

25. 11. 30. 오후 1:11 

NONE Actors - 지망생 심리 분석 

 

```
if state.switches.EVALUATION == "Off":
```

```
ifany(kw in b for kw in FORBIDDEN_EVALUATION):
```

```
                forbid = True
```

```
if forbid:
            isolated.append(b)
else:
```

```
            body.append(b)
return body, isolated
```

```
# -----------------------------
```

## _`# Renderers`_ 

```
# -----------------------------
```

```
defrender_overview(state: EngineState) -> str:
```

```
    s = []
```

```
    s.append(f"[핵심결론 1] (예시) 세부주제 ‘{state.topic or'?'}’의핵심은우선 **
```

```
    s.append(f"[핵심결론 2] (예시) 확정된핵심을기준으로심리→마케팅순으로아이디어를
    s.append("[세부절차/설명] A(개요)→B1(핵심)→B2P(심리)→B2M(마케팅). PARTICIPATION
    s.append("[위험·한계·대안] 규칙충돌시 P1(락) 우선. 금칙어는확장섹션으로자동격
return"\n".join(s)
```

```
defrender_core_candidates() -> str:
```

```
    s = []
```

- `s.append("[핵심 후보 1] 머무르기=새 연기에 시간을 쓰지 않음(‘대사만 안 틀리면 된다` 

- `s.append("[핵심 후보 2] 암기 안전지대 유지=변형 비용을 의미 없이 회피. [추론] — (근` 

- `s.append("[권고 확정문] “핵심= 머무르기(시간 투입 회피)”")` 

```
    s.append("[핵심승인요청:{선택/수정/결합/폐기}]")
return"\n".join(s)
```

```
defrender_psych(state: EngineState) -> str:
```

```
    core = state.core or"핵심미확정"
```

## _`#`_ `예시 심리 해법` _`(`_ `참여` _`/`_ `평가 요구 없는 메시지` _`/`_ `환경` _`/`_ `정책 중심` _`)`_ 

```
    fr = [
```

```
"‘버림’이아니라 ‘조정’이라는명칭고정(내색 30% 보존, 적합 70%).",
```

```
"‘안전 vs 적합’ 대비문장 3개만사용(실수공포↓/재현성↑/배역적합↑).",
"‘한가지차이’ 원칙고지: 시선·리듬·구도중 1개만알아도달라짐.",
    ]
    cog = [
```

- `"손실회피 리프레이밍: ‘실패 회피’가 아니라 ‘적합 확률 상승’의 언어 사용.",` 

- `"즉시보상 대체: 읽기 전용 예시(3컷)로 ‘노력 없이 체감’ 제공.",` 

```
"자기확증편향완화: ‘내색 30% 보존’ 수치고정문구.",
    ]
    env = [
```

```
"모든자료에 ‘참고용·제출의무없음’ 라벨상시노출.",
```

https://chatgpt.com/g/g-p-68b4572c50c081918f5a18bb2ad93cc5-none-actors/c/68b90a57-cc04-832c-ba9c-a95f68318fb8 

45/205 

25. 11. 30. 오후 1:11 

NONE Actors - 지망생 심리 분석 

 

```
"등록전 ‘부적합안내’ 배지노출(공격적표현금지, 중립톤).",
```

```
"랜딩/FAQ에화면언어(시선·리듬·구도) 용어집합고정.",
    ]
defbullets(label, arr):
        body, iso = filter_body_and_isolate(state, [f"• {a}" for a in arr])
        state.expansions.extend(iso)
returnf"- {label}: " + (" ".join(body) if body else"• (금칙어격리됨)")
    out = [
f"<세부주제: {state.topic or'?'}>（핵심= “{core}” / PSY=On / PARTICIPATION
        bullets("프레임전환", fr),
        bullets("인지/손실회피", cog),
        bullets("환경/정책", env),
"[위험·한계·대안: • 과도한단호함은반발↑ • ‘조정’/‘참고용’ 톤유지로완화]"
    ]
return"\n".join(out)
defrender_marketing(state: EngineState) -> str:
    core = state.core or"핵심미확정"
    hooks = [
"“대사는완벽했는데, 카메라는왜고개를젓나.”",
"“암기=안전? 카메라는 ‘방향’을본다.”",
"“시선 3도, 인상 180도.”",
"“잘외웠다에서, 잘보이게로.”",
    ]
    msg = [
"버림이아닌 **조정**(내색 30%·적합 70%).",
"하는사람기준 → **보는사람기준**(화면언어).",
"‘한가지차이’ 원칙(시선/리듬/구도중 1개).",
    ]
    assets = [
"3컷비교정지이미지(암기/과장/적합) + 1줄설명(참고용).",
"1분읽기 PDF: ‘암기 vs 화면언어한눈표’(읽기전용).",
"릴스 15초세트: 시선3도/리듬1초/구도절반(텍스트팝업).",