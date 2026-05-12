package collector

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// WebScraper는 웹사이트의 본문을 추출하여 MD로 변환한다.
type WebScraper struct {
	outputDir string
}

// NewWebScraper는 새로운 웹 수집기를 생성한다.
// 기본 저장소: data/inbox/web/
func NewWebScraper() *WebScraper {
	dir := filepath.Join("data", "inbox", "web")
	os.MkdirAll(dir, 0755)
	return &WebScraper{outputDir: dir}
}

// Name returns "web".
func (s *WebScraper) Name() string { return "web" }

// CanHandle returns true for any non-YouTube URL (as a fallback).
func (s *WebScraper) CanHandle(url string) bool {
	return strings.HasPrefix(url, "http") && !strings.Contains(url, "youtube.com") && !strings.Contains(url, "youtu.be")
}

// Collect는 웹 URL의 본문을 긁어서 .md 파일로 저장한다.
func (s *WebScraper) Collect(ctx context.Context, targetURL string) (string, error) {
	c := colly.NewCollector()
	var title string
	var content strings.Builder

	// 1. 타이틀 추출
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = strings.TrimSpace(e.Text)
	})

	// 2. 본문 추출 (보통 article, main, 또는 p 태그들이 모인 곳)
	// 광고/메뉴를 피하기 위해 p, h1, h2, h3 위주로 수집
	c.OnHTML("article, main, div.content, div.post-body", func(e *colly.HTMLElement) {
		e.ForEach("p, h1, h2, h3, h4", func(_ int, el *colly.HTMLElement) {
			text := strings.TrimSpace(el.Text)
			if text == "" {
				return
			}
			switch el.Name {
			case "h1":
				content.WriteString("# " + text + "\n\n")
			case "h2":
				content.WriteString("## " + text + "\n\n")
			case "h3":
				content.WriteString("### " + text + "\n\n")
			default:
				content.WriteString(text + "\n\n")
			}
		})
	})

	// 폴백: 위 태그들이 없을 경우 모든 p 태그 수집
	if content.Len() == 0 {
		c.OnHTML("p", func(e *colly.HTMLElement) {
			content.WriteString(strings.TrimSpace(e.Text) + "\n\n")
		})
	}

	// 3. 실행
	if err := c.Visit(targetURL); err != nil {
		return "", fmt.Errorf("colly visit: %w", err)
	}

	if title == "" {
		title = "web_article"
	}

	// 4. 저장
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s_%s.md", timestamp, s.makeSafeFilename(title))
	fullPath := filepath.Join(s.outputDir, filename)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# Web 지식 도축: %s\n\n", title))
	sb.WriteString(fmt.Sprintf("- **출처**: [%s](%s)\n", targetURL, targetURL))
	sb.WriteString(fmt.Sprintf("- **도축 일시**: %s\n\n", time.Now().Format("2006-01-02 15:04:05")))
	sb.WriteString("--- \n\n")
	sb.WriteString(content.String())

	if err := os.WriteFile(fullPath, []byte(sb.String()), 0644); err != nil {
		return "", fmt.Errorf("save web md: %w", err)
	}

	return fullPath, nil
}

func (s *WebScraper) makeSafeFilename(title string) string {
	r := strings.NewReplacer("/", "_", "\\", "_", ":", "_", "*", "_", "?", "_", "\"", "_", "<", "_", ">", "_", "|", "_", " ", "_")
	return r.Replace(title)
}
