
**A**: bubbletea의 `tea.ProgramTest` 또는 `teatest` 패키지를 활용하여 키 입력 시뮬레이션 기반 통합 테스트를 작성한다. 각 step의 Model은 순수 함수(`Update`)로 구현되므로 unit test도 가능하다. non-TTY fallback 경로는 기존 테스트 패턴 유지.

---

## Appendix: Current Dependency Tree (Charmbracelet)

```
현재:
  github.com/charmbracelet/lipgloss v1.1.0

추가 예정:
  github.com/charmbracelet/bubbletea  (lipgloss 의존)
  github.com/charmbracelet/bubbles    (bubbletea 의존)
  github.com/charmbracelet/huh        (bubbletea + bubbles 의존, 검토 후 결정)
```

## Appendix: Current Init Flow vs Target Flow

```
현재 (bufio.Reader):
  ┌──────────────┐    번호 입력     ┌──────────────┐    번호 입력
  │  Step 1      │ ──────────────→ │  Step 2      │ ──────────────→ ...
  │  텍스트 출력  │   "1" + Enter   │  텍스트 출력  │   "2" + Enter
  └──────────────┘                 └──────────────┘

목표 (bubbletea/huh):
  ┌──────────────────────────────────────┐
  │  [1/5] Language Settings             │
  │                                      │
  │  Code comments language?             │
  │                                      │
  │    ● English           ← selected    │
  │    ○ Korean (한국어)                  │
  │    ○ Japanese (日本語)               │
  │    ○ Chinese (中文)                  │
  │                                      │
  │  ↑↓ navigate  Enter confirm          │
  └──────────────────────────────────────┘
```