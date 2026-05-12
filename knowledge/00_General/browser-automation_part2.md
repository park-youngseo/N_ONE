cmux browser --surface $SURFACE snapshot
# → @e3 switch "Auto Fallback" 이 존재하면 렌더링 성공
cmux browser --surface $SURFACE screenshot --out /tmp/ai-settings.png
```

### 패턴 2: 토글 동작 검증

```bash
cmux browser --surface $SURFACE snapshot
cmux browser --surface $SURFACE click --selector "@e3" --snapshot-after
```

### 패턴 3: 인증 후 테스트

```bash
SURFACE=$(cmux browser open https://example.com/login)
cmux browser --surface $SURFACE snapshot
cmux browser --surface $SURFACE fill --selector "@e2" --text "user@example.com"
cmux browser --surface $SURFACE fill --selector "@e3" --text "password"
cmux browser --surface $SURFACE click --selector "@e4"
cmux browser --surface $SURFACE wait --load-state complete
cmux browser --surface $SURFACE goto https://example.com/settings/ai
cmux browser --surface $SURFACE snapshot
```

## 판정 기준

| 판정 | 기준 |
|------|------|
| PASS | 기대 요소가 snapshot에 존재하고 올바른 상태 |
| WARN | 요소는 존재하나 상태가 예상과 다름 |
| FAIL | 기대 요소가 snapshot에 없거나 에러 발생 |

## 주의사항

- 백엔드는 `auto terminal detect`로 자동 선택됨 — 수동 지정 금지
- cmux 환경에서는 `cmux browser`가 최우선 — agent-browser를 사용하면 세션이 분리됨
- cmux `--surface` 핸들은 `open` 명령이 반환 — 후속 명령에 반드시 전달
- `snapshot`은 **접근성 트리**만 반환 — CSS 스타일은 포함 안 됨 (시각 확인은 `screenshot` 사용)
- 운영환경 테스트 시 **쓰기 작업**(삭제, 설정 변경)은 신중하게 — 되돌릴 수 없을 수 있음