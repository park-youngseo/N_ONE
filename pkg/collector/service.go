package collector

import (
	"context"
	"fmt"
	"log/slog"
)

// UniversalCollector는 여러 수집기를 관리하고 URL에 따라 적절한 엔진을 선택한다.
type UniversalCollector struct {
	collectors []Collector
}

// NewUniversalCollector는 유튜브와 웹 엔진이 장착된 통합 수집기를 생성한다.
func NewUniversalCollector() *UniversalCollector {
	return &UniversalCollector{
		collectors: []Collector{
			NewYouTubeCollector(),
			NewWebScraper(),
		},
	}
}

// Execute는 URL을 분석하여 최적의 엔진으로 지식을 도축한다.
func (uc *UniversalCollector) Execute(ctx context.Context, url string) (string, error) {
	slog.Info("통합 지식 수집 가동", "url", url)

	for _, c := range uc.collectors {
		if c.CanHandle(url) {
			slog.Info("최적 엔진 선택됨", "engine", c.Name())
			return c.Collect(ctx, url)
		}
	}

	return "", fmt.Errorf("해당 URL을 처리할 수 있는 엔진이 없습니다: %s", url)
}
