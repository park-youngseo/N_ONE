package collector

import "context"

// Collector는 외부 지식을 수집하고 변환하는 공통 인터페이스다.
type Collector interface {
	// Name은 수집기의 이름을 반환한다 (예: "youtube", "web").
	Name() string
	// CanHandle은 해당 URL을 처리할 수 있는지 판단한다.
	CanHandle(url string) bool
	// Collect는 URL에서 데이터를 추출하여 마크다운 파일 경로를 반환한다.
	Collect(ctx context.Context, url string) (string, error)
}
