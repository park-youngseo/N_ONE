```go
// auto go 완료 핸들러에 추가 (1줄)
if pipelineSuccess {
    hint.RecordSuccess(projectPath)
    hint.CheckAndShow(projectPath, cfg.UsageProfile.Effective(), cfg.Hints.IsPlatformHintEnabled(), os.Stdout)
}
```

### hint.CheckAndShow 로직

```go
func CheckAndShow(projectPath string, profile config.UsageProfile, hintsEnabled bool, w io.Writer) bool {
    // Guard: only for developer profile with hints enabled
    if profile != config.ProfileDeveloper || !hintsEnabled {
        return false
    }

    store, err := NewStateStore()
    if err != nil {
        return false // graceful degradation
    }

    state, err := store.Load(projectPath)
    if err != nil {
        return false
    }

    var shown bool

    if state.GoSuccessCount >= 1 && !state.FirstHintShown {
        fmt.Fprintln(w, Hint1)
        state.FirstHintShown = true
        shown = true
    } else if state.GoSuccessCount >= 3 && !state.SecondHintShown {
        fmt.Fprintln(w, Hint2)
        state.SecondHintShown = true
        shown = true
    }

    if shown {
        _ = store.Save(projectPath, state) // best-effort
    }

    return shown
}
```

### hint.RecordSuccess 로직

```go
func RecordSuccess(projectPath string) error {
    store, err := NewStateStore()
    if err != nil {
        return err
    }

    state, err := store.Load(projectPath)
    if err != nil {
        state = &ProjectState{}
    }

    state.GoSuccessCount++
    return store.Save(projectPath, state)
}
```

## 6. auto config set hints.platform 지원

### 명령어 형태

```bash
auto config set hints.platform false   # 힌트 영구 비활성화
auto config set hints.platform true    # 힌트 재활성화
```

### 구현 방식

기존 `auto config set` 커맨드의 키 매핑에 `hints.platform` 경로를 추가한다.

```go
// config set 핸들러 내부
case "hints.platform":
    val := args[1] == "true"
    cfg.Hints.Platform = &val
```

변경 후 autopus.yaml을 다시 직렬화하여 저장한다.

## 7. 파일 구조 요약

```
autopus-adk/
├── pkg/
│   ├── config/
│   │   ├── schema.go              # 기존 (HarnessConfig에 2 필드 추가)
│   │   └── schema_profile.go      # 신규 (UsageProfile, HintsConf 타입)
│   └── hint/
│       ├── state.go               # 신규 (StateStore, ProjectState)
│       ├── state_test.go          # 신규 (StateStore 단위 테스트)
│       ├── hint.go                # 신규 (CheckAndShow, RecordSuccess)
│       └── hint_test.go           # 신규 (힌트 로직 단위 테스트)
├── internal/
│   └── cli/
│       ├── tui/
│       │   ├── wizard_steps.go    # 기존 (Opts/Result 확장, buildStepList 수정)
│       │   └── wizard_profile.go  # 신규 (buildProfileStep)
│       └── config_cmd.go          # 기존 또는 신규 (hints.platform set 지원)
```

## 8. Backward Compatibility

| Scenario | Behavior |
|----------|----------|
| 기존 autopus.yaml에 `usage_profile` 없음 | `developer`로 기본 동작 |
| 기존 autopus.yaml에 `hints` 없음 | 힌트 활성화(기본값) |
| `~/.autopus/state.json` 없음 | 자동 생성, 초기 상태 사용 |
| `~/.autopus/state.json` 읽기 실패 | 힌트 건너뜀, 에러 없음 |
| `~/.autopus/` 디렉토리 생성 실패 | 힌트 건너뜀, 에러 없음 |

## 9. 보안 고려사항

- `state.json`은 사용자 홈 디렉토리에 저장되며 프로젝트 코드에 포함되지 않음
- 프로젝트 경로는 SHA-256 해시로 저장되어 경로 노출 없음
- `state.json`에 민감한 정보(토큰, 인증 정보)를 저장하지 않음
- 파일 권한은 `0o644`(기본), 디렉토리는 `0o755`