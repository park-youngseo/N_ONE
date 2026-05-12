# [Project: NONE_AI - Stock Trading Harness]

## 주식 매매 통합 관리 규범 (Stock Trading Harness)

이 문서는 거장들의 매매 로직을 NONE-OS 시스템에 이식하기 위한 최상위 지침이다. 모든 매매 전략 분석 및 제안은 아래의 **'생존 법칙'**을 1순위로 준수한다.

---

### 🛡️ 제1법칙: 생존과 리스크 관리 (Survival First)
1.  **자본 보호 (Capital Preservation)**: 수익보다 손실 방지가 우선이다. 100만 원의 시드가 0이 되면 게임은 끝난다.
2.  **1%의 룰 (The 1% Rule)**: 단일 매매에서 발생하는 손실은 전체 자산의 1~2%를 넘지 않도록 포지션 규모를 조절한다.
3.  **기계적 손절 (Hard Stop-Loss)**: 진입과 동시에 손절 주문을 예약하며, 어떤 이유로도 손절 라인을 뒤로 미루지 않는다. (최대 -7~8% 원칙)

### 📊 제2법칙: 통계적 우위와 근거 (Statistical Edge)
1.  **데이터 기반 진입 (Data-Driven Entry)**: "오를 것 같다"는 추측은 금지한다. 하네스에 정의된 패턴(VCP, 이격도 등)이 발생했을 때만 진입한다.
2.  **손익비 우위 (Asymmetric Payoff)**: 기대 수익이 기대 손실보다 최소 3배 이상 큰 (3:1) 자리에서만 싸운다.
3.  **추세 순응 (Trend Following)**: 시장의 대세 하락장(Bear Market)에서는 매매를 쉬는 것도 전략이다.

### 🤖 제3법칙: 감정 배제와 프로세스 (Process over Emotion)
1.  **체크리스트 강제 (Checklist Force)**: 모든 매매는 사전에 정의된 체크리스트를 100% 통과해야 한다.
2.  **복기 및 기록 (Journaling)**: 성공과 실패의 원인을 반드시 기록하여 시스템의 통계적 신뢰도를 높인다.
3.  **뇌동매매 금지 (No Revenge Trading)**: 손실 후 복구하려는 감정적 매매는 즉시 차단한다.

---

## [NONE-OS 매매 엔진 분류]
- **Type-A (Aggressive)**: 소액 계좌 급성장을 위한 모멘텀 돌파 매매 (Minervini, Zanger)
- **Type-S (Safety)**: 과매도 구간의 기술적 반등을 노리는 역추세 매매 (BNF)
- **Type-G (Growth)**: 펀더멘털과 기술적 분석의 조화 (O'Neil)

---

*하네스 확정: 2026-05-04 | 자본 시장 데이터 인출기 v1.0 작동 시작*

---
## 🔗 관련 시너지 지식
- **핵심 하네스**: [Agent_Self_Control_Protocol.md](../13_Agent_Brain_System/Agent_Self_Control_Protocol.md)
- **클러스터**: Business & Monetization