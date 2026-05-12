package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/insajin/autopus-adk/pkg/collector"
)

func main() {
	channelURL := "https://www.youtube.com/@cocojuantv/videos"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Hour) // 대량 수집이므로 넉넉하게 2시간
	defer cancel()

	fmt.Println("\n🎬 N-ONE 채널 전체 지식 도축 시작")
	fmt.Println("====================================================")
	fmt.Printf("🎯 대상 채널: %s\n", channelURL)

	// 1. 채널 내 모든 비디오 ID 추출
	fmt.Println("🔍 동영상 목록 가져오는 중...")
	cmd := exec.CommandContext(ctx, "yt-dlp", "--get-id", "--flat-playlist", channelURL)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ 목록 추출 실패: %v\n", err)
		return
	}

	videoIDs := strings.Split(strings.TrimSpace(string(out)), "\n")
	fmt.Printf("✅ 총 %d개의 영상을 찾았습니다.\n", len(videoIDs))

	// 2. 통합 수집기 가동
	uc := collector.NewUniversalCollector()

	for i, id := range videoIDs {
		videoURL := "https://www.youtube.com/watch?v=" + id
		fmt.Printf("\n🚀 [%d/%d] 도축 중: %s\n", i+1, len(videoIDs), videoURL)
		
		path, err := uc.Execute(ctx, videoURL)
		if err != nil {
			fmt.Printf("  ❌ 실패: %v\n", err)
			// 너무 많은 실패(429)가 발생하면 잠시 쉬어감
			if strings.Contains(err.Error(), "429") {
				fmt.Println("  ⏳ 차단 감지! 60초간 대기합니다...")
				time.Sleep(60 * time.Second)
			}
		} else {
			fmt.Printf("  ✅ 성공! 저장 완료: %s\n", path)
		}
		
		// 연속 요청 방지를 위한 최소한의 매너 타임 (5초로 증대)
		fmt.Println("  ⏳ 다음 도축 대기 중 (5초)...")
		time.Sleep(5 * time.Second)
	}

	fmt.Println("\n🎬 채널 도축 공정 완료!")
}
