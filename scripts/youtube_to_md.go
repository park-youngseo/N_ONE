package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	inboxDir := `c:\Users\psm23\Desktop\N_ONE\data\inbox\youtube`
	outDir := `c:\Users\psm23\Desktop\N_ONE\knowledge\03_Automation\YouTube`
	os.MkdirAll(outDir, 0755)

	files, _ := filepath.Glob(filepath.Join(inboxDir, "*.vtt"))
	fmt.Printf("📂 총 %d개의 유튜브 자막 파일을 발견했습니다.\n", len(files))

	reTimestamp := regexp.MustCompile(`(?m)^\d{2}:\d{2}:\d{2}.\d{3} --> \d{2}:\d{2}:\d{2}.\d{3}.*$`)
	reHeader := regexp.MustCompile(`(?m)^WEBVTT.*$|^Kind:.*$|^Language:.*$`)

	for _, f := range files {
		baseName := filepath.Base(f)
		mdName := strings.TrimSuffix(baseName, filepath.Ext(baseName)) + ".md"
		destPath := filepath.Join(outDir, mdName)

		content, err := os.ReadFile(f)
		if err != nil {
			continue
		}

		// VTT 정제 로직
		text := string(content)
		text = reTimestamp.ReplaceAllString(text, "")
		text = reHeader.ReplaceAllString(text, "")
		
		// 중복 줄바꿈 제거 및 정리
		lines := strings.Split(text, "\n")
		var cleanLines []string
		seen := make(map[string]bool)
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || seen[trimmed] {
				continue
			}
			cleanLines = append(cleanLines, trimmed)
			seen[trimmed] = true
		}

		finalText := fmt.Sprintf("# YouTube Knowledge: %s\n\n%s", baseName, strings.Join(cleanLines, " "))
		os.WriteFile(destPath, []byte(finalText), 0644)
		
		// 처리 완료 후 원본 삭제 (선택)
		// os.Remove(f)
		fmt.Printf("✅ 변환 완료: %s\n", mdName)
	}

	fmt.Println("🚀 모든 유튜브 지식이 MD 파일로 도축되어 수납되었습니다!")
}
