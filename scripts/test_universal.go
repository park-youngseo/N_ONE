package main

import (
	"context"
	"fmt"
	"time"

	"github.com/insajin/autopus-adk/pkg/collector"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	fmt.Println("\n🐙 N-ONE 통합 지식 수집기(Universal Collector) 테스트")
	fmt.Println("====================================================")

	// 통합 서비스 생성 (단 하나의 입구)
	uc := collector.NewUniversalCollector()

	// 테스트 URL 리스트 (유튜브와 웹이 섞여 있음)
	urls := []string{
		"https://go.dev/blog/go1.22",              // 웹사이트
		"https://www.youtube.com/watch?v=dQw4w9WgXcQ", // 유튜브
	}

	for _, url := range urls {
		fmt.Printf("\n🚀 수집 시작: %s\n", url)
		path, err := uc.Execute(ctx, url)
		if err != nil {
			fmt.Printf("  ❌ 실패: %v\n", err)
		} else {
			fmt.Printf("  ✅ 성공! 도축된 지식 저장 위치: %s\n", path)
		}
	}

	fmt.Println("\n🐙 통합 수집기 테스트 완료!")
}
