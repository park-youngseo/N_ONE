    **Round 2 (인정/통합/리스크 원문)**:
    {cleaned full output — 축약 금지}

    ### 토론자 B
    **Round 1 (독립 발산 원문)**: {원문}
    **Round 2 (인정/통합/리스크 원문)**: {원문}

    ### 토론자 C
    **Round 1 (독립 발산 원문)**: {원문}
    **Round 2 (인정/통합/리스크 원문)**: {원문}

    ## 과제: 합의 기반 통합 판정

    ### 1. 합의 영역 추출
    2명 이상의 토론자가 Round 2에서 **공통으로 통합한** 아이디어를 추출하세요.
    공통 통합 = 높은 확신도 신호입니다.

    ### 2. 고유 인사이트 식별
    1명만 제안하고 다른 토론자가 통합하지 않은 독창적 아이디어를 식별하세요.
    독창적이지만 위험할 수 있으므로 별도 표기합니다.

    ### 3. 교차 리스크 분석
    복수 토론자가 Round 2 Step 3에서 **공통으로 지적한 리스크**를 정리하세요.
    공통 리스크 = 실재할 가능성이 높은 위험입니다.

    ### 4. 최종 통합 ICE 스코어링
    위 분석을 기반으로 Top 5 아이디어를 선정하고 ICE 스코어링하세요:
    - Impact (1-10): 프로젝트 컨텍스트를 고려한 실질적 영향력
    - Confidence (1-10): **합의 수준 반영** — 3명 합의 > 2명 > 1명
    - Ease (1-10): 현재 코드베이스에서의 구현 용이성, 교차 리스크 반영
    - Score = (Impact × Confidence × Ease) / 100

    나머지 아이디어는 부록에 포함하세요.
    아이디어의 내용만으로 평가하세요. 토론자의 정체는 알 수 없으며 알 필요도 없습니다.
  """
)
```

#### 4.3: 결과 수신 및 매핑 복원

서브에이전트의 판정 결과를 수신한 후, 메인 세션이 익명 매핑을 복원하여 BS 파일에 기록합니다:
- 토론자 A → {실제 프로바이더 이름} (BS 파일의 프로바이더별 발산 결과 섹션용)
- ICE Top N은 익명 상태 그대로 기록 (어떤 프로바이더가 제안했는지는 부차적)
- 합의 영역 / 고유 인사이트 / 교차 리스크를 BS 파일에 별도 섹션으로 기록

#### Assumption Risk Overlay

ICE Top N 아이디어 각각에 대해 Step 2에서 식별한 가정의 위험도를 오버레이:

| Rank | Idea | ICE Score | Top Risk Assumption | Risk Level |
|------|------|-----------|---------------------|------------|
| 1 | ... | 7.2 | "사용자가 X를 원한다" (Value) | HIGH |
| 2 | ... | 6.8 | "API 지연 < 500ms" (Feasibility) | MEDIUM |

HIGH 위험 가정이 있는 아이디어는 `/auto plan` 전에 검증 실험을 권장합니다.

### Step 5: Save and Guide Next Steps

BS-{ID} 파일 저장 후 Workflow Lifecycle 바 표시 및 다음 단계 안내.

**ID 자동 증분**: `{target-module}/.autopus/brainstorms/BS-{ID}.md` 파일이 이미 존재하면 ID를 증분합니다. 전체 프로젝트 스캔으로 ID 유일성을 보장합니다.

## BS 파일 형식

`{target-module}/.autopus/brainstorms/BS-{ID}.md`:

```markdown
# BS-{ID}: {title}

**Created**: {date}
**Strategy**: {strategy}
**Providers**: {provider list}
**Status**: active

## 원본 아이디어
- What: {description}
- Why: {motivation}
- Who: {target users}
- When: {timeline}

## 프로바이더별 발산 결과
{raw brainstorm output}

## ICE 스코어링 — Top N
| Rank | Idea | Impact | Confidence | Ease | Score |
|------|------|--------|------------|------|-------|

## 추천 방향
{judge's recommendation}

## 다음 단계
`/auto plan --from-idea BS-{ID} "feature description"`
```

## 완료 후 출력

```
🐙 Workflow: BS-{ID}
  ● idea  →  ○ plan  →  ○ go  →  ○ sync
```

`--auto` 플래그가 있으면 자동으로 `/auto plan --from-idea BS-{ID}`로 체이닝합니다.