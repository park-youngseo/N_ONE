package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inboxDir := `c:\Users\psm23\Desktop\N_ONE\data\inbox`
	baseKnowledgeDir := `c:\Users\psm23\Desktop\N_ONE\knowledge`

	fmt.Println("🚀 재귀적 지식 이식 공정 가동 시작...")

	err := filepath.Walk(inboxDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		category := classify(info.Name())
		// 만약 파일이 특정 폴더 안에 있다면, 그 폴더 이름도 분류에 참고
		parentDir := filepath.Base(filepath.Dir(path))
		if parentDir != "inbox" {
			category = classify(parentDir + "_" + info.Name())
		}

		destDir := filepath.Join(baseKnowledgeDir, category)
		os.MkdirAll(destDir, 0755)

		fmt.Printf("📦 재귀 처리 중: [%s] -> %s\n", category, info.Name())
		processAndSplit(path, destDir, info.Name())
		return nil
	})

	if err != nil {
		fmt.Printf("❌ 이식 공정 중 오류 발생: %v\n", err)
	}

	fmt.Println("\n✅ 모든 지식 이식이 완료되었습니다!")
}

// 파일명 기반 자동 분류 로직
func classify(name string) string {
	n := strings.ToLower(name)
	switch {
	case containsAny(n, "actors", "none", "김주희", "연기"):
		return "01_Screenwriting"
	case containsAny(n, "마케팅", "start", "전략", "시장"):
		return "02_Marketing"
	case containsAny(n, "프로그램", "크롤러", "프롬포트", "ai", "gpt", "제미나이", "coding"):
		return "03_Automation"
	case containsAny(n, "주식", "거래량", "차트"):
		return "04_Finance"
	default:
		return "00_General"
	}
}

func containsAny(s string, keywords ...string) bool {
	for _, k := range keywords {
		if strings.Contains(s, k) {
			return true
		}
	}
	return false
}

// 파일을 300줄 단위로 쪼개서 저장 (AX-300 준수)
func processAndSplit(srcPath, destDir, originalName string) {
	file, err := os.Open(srcPath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	part := 1
	baseName := strings.TrimSuffix(originalName, ".md")

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) >= 300 {
			savePart(destDir, baseName, part, lines)
			lines = nil
			part++
		}
	}
	if len(lines) > 0 {
		savePart(destDir, baseName, part, lines)
	}
	
	// 원본 파일은 인박스에서 삭제 (선택 사항, 여기서는 안전을 위해 보존하거나 다른 곳으로 이동 가능)
	// os.Remove(srcPath)
}

func savePart(dir, baseName string, part int, lines []string) {
	filename := fmt.Sprintf("%s_part%d.md", baseName, part)
	if part == 1 && len(lines) < 300 {
		filename = baseName + ".md" // 300줄 미만이면 파트 번호 없이 저장
	}
	
	path := filepath.Join(dir, filename)
	content := strings.Join(lines, "\n")
	os.WriteFile(path, []byte(content), 0644)
}
