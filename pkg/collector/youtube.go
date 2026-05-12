// Package collector는 외부 지식(YouTube, Web)을 하네스용 MD로 변환하는 기능을 제공한다.
package collector

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// YouTubeCollector는 유튜브 비디오에서 자막을 추출하여 MD로 변환한다.
type YouTubeCollector struct {
	outputDir string
}

// NewYouTubeCollector는 새로운 유튜브 수집기를 생성한다.
// 기본 저장소: data/inbox/youtube/
func NewYouTubeCollector() *YouTubeCollector {
	dir := filepath.Join("data", "inbox", "youtube")
	os.MkdirAll(dir, 0755)
	return &YouTubeCollector{outputDir: dir}
}

// Collect는 유튜브 URL을 받아 자막을 추출하고 .md 파일로 저장한다.
func (c *YouTubeCollector) Collect(ctx context.Context, videoURL string) (string, error) {
	// 1. 비디오 제목 및 ID 가져오기
	title, err := c.getVideoTitle(ctx, videoURL)
	if err != nil {
		return "", fmt.Errorf("youtube title: %w", err)
	}
	// 파일명 안전하게 변환
	safeTitle := c.makeSafeFilename(title)
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s_%s.md", timestamp, safeTitle)
	fullPath := filepath.Join(c.outputDir, filename)

	// 2. yt-dlp를 사용하여 자막 추출 (자동 번역 포함 한국어 우선)
	// --skip-download: 비디오 파일은 안 받음
	// --write-auto-subs: 자동 생성 자막 포함
	// --sub-lang: 한국어(ko) 우선
	// --convert-subs: srt 형식으로 변환
	tempPrefix := filepath.Join(c.outputDir, "temp_"+timestamp)
	args := []string{
		"--skip-download",
		"--write-auto-sub",
		"--sub-lang", "ko,en",
		"--convert-subs", "srt",
		"-o", tempPrefix,
		videoURL,
	}

	cmd := exec.CommandContext(ctx, "yt-dlp", args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("yt-dlp exec: %v (output: %s)", err, string(out))
	}

	// 3. 생성된 .srt 파일 찾기 및 읽기
	srtPath := tempPrefix + ".ko.srt"
	if _, err := os.Stat(srtPath); os.IsNotExist(err) {
		srtPath = tempPrefix + ".en.srt" // 한국어 없으면 영어 시도
		if _, err := os.Stat(srtPath); os.IsNotExist(err) {
			return "", fmt.Errorf("subtitles not found for: %s", videoURL)
		}
	}
	defer os.Remove(srtPath) // 작업 후 임시 파일 삭제

	srtContent, err := os.ReadFile(srtPath)
	if err != nil {
		return "", fmt.Errorf("read srt: %w", err)
	}

	// 4. SRT를 MD로 변환 및 저장
	mdContent := c.convertSrtToMd(title, videoURL, string(srtContent))
	if err := os.WriteFile(fullPath, []byte(mdContent), 0644); err != nil {
		return "", fmt.Errorf("save md: %w", err)
	}

	return fullPath, nil
}

func (c *YouTubeCollector) getVideoTitle(ctx context.Context, url string) (string, error) {
	cmd := exec.CommandContext(ctx, "yt-dlp", "--get-title", url)
	out, err := cmd.Output()
	if err != nil {
		return "unknown_video", nil
	}
	return strings.TrimSpace(string(out)), nil
}

func (c *YouTubeCollector) makeSafeFilename(title string) string {
	r := strings.NewReplacer("/", "_", "\\", "_", ":", "_", "*", "_", "?", "_", "\"", "_", "<", "_", ">", "_", "|", "_", " ", "_")
	return r.Replace(title)
}

func (c *YouTubeCollector) convertSrtToMd(title, url, srt string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# YouTube 지식 도축: %s\n\n", title))
	sb.WriteString(fmt.Sprintf("- **출처**: [%s](%s)\n", url, url))
	sb.WriteString(fmt.Sprintf("- **도축 일시**: %s\n\n", time.Now().Format("2006-01-02 15:04:05")))
	sb.WriteString("## 📜 스크립트 본문\n\n")

	// 단순화를 위해 타임라인 제거하고 텍스트만 추출 (AX-300 준수)
	lines := strings.Split(srt, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || c.isTimestamp(line) || c.isNumber(line) {
			continue
		}
		sb.WriteString(line + " ")
	}

	return sb.String()
}

func (c *YouTubeCollector) isTimestamp(line string) bool {
	return strings.Contains(line, "-->")
}

func (c *YouTubeCollector) isNumber(line string) bool {
	var n int
	_, err := fmt.Sscanf(line, "%d", &n)
	return err == nil
}
