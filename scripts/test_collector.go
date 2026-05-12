package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/insajin/autopus-adk/pkg/collector"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	fmt.Println("\n🐙 N-ONE 지식 수집기(Collector) 통합 테스트")
	fmt.Println("==========================================")

	// 1. 유튜브 수집 테스트
	fmt.Println("\n[1/2] 유튜브 지식 도축 테스트...")
	yt := collector.NewYouTubeCollector()
	ytURL := "https://www.youtube.com/watch?v=dQw4w9WgXcQ" // 테스트용 (Rick Astley - Never Gonna Give You Up)
	ytPath, err := yt.Collect(ctx, ytURL)
	if err != nil {
		fmt.Printf("  ❌ 유튜브 도축 실패: %v\n", err)
	} else {
		fmt.Printf("  ✅ 유튜브 도축 성공: %s\n", ytPath)
		content, _ := os.ReadFile(ytPath)
		fmt.Printf("     (내용 일부: %s...)\n", string(content[:100]))
	}

	// 2. 웹 수집 테스트
	fmt.Println("\n[2/2] 웹 지식 도축 테스트...")
	web := collector.NewWebScraper()
	webURL := "https://go.dev/blog/go1.21" // 테스트용 (Go 공식 블로그)
	webPath, err := web.Collect(ctx, webURL)
	if err != nil {
		fmt.Printf("  ❌ 웹 도축 실패: %v\n", err)
	} else {
		fmt.Printf("  ✅ 웹 도축 성공: %s\n", webPath)
		content, _ := os.ReadFile(webPath)
		fmt.Printf("     (내용 일부: %s...)\n", string(content[:100]))
	}

	fmt.Println("\n🐙 지식 수집기 테스트 완료!")
}
