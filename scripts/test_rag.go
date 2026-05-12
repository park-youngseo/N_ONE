package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/insajin/autopus-adk/pkg/rag"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	fmt.Println("\n🐙 N-ONE RAG 지식창고 통합 테스트")
	fmt.Println("================================")

	// 1. 임베딩 엔진 가동 확인
	fmt.Println("\n[1/4] 올라마 임베딩 엔진 연결 확인...")
	embedder := rag.NewEmbedder()
	if err := embedder.Ping(ctx); err != nil {
		fmt.Printf("  ❌ 임베딩 서버 미응답: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("  ✅ 임베딩 엔진 가동 중")

	// 2. 벡터 스토어 생성
	fmt.Println("\n[2/4] 벡터 스토어 초기화...")
	storePath := filepath.Join(".autopus", "knowledge", "vectors.json")
	os.MkdirAll(filepath.Dir(storePath), 0755)
	store, err := rag.NewVectorStore(storePath)
	if err != nil {
		fmt.Printf("  ❌ 스토어 초기화 실패: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  ✅ 벡터 스토어 초기화 완료 (기존 문서: %d개)\n", store.Count())

	// 3. 테스트 문서 임베딩
	fmt.Println("\n[3/4] 테스트 문서 임베딩 중...")
	testDocs := []struct {
		id, content, source string
	}{
		{"test-1", "로버트 맥키의 스토리 이론에 따르면, 모든 장면에는 가치의 전환(Value Shift)이 반드시 존재해야 한다.", "mckee_theory.md"},
		{"test-2", "N-ONE 연기 아카데미는 메소드 연기법을 기반으로 한 실전 배우 양성 프로그램이다.", "n_one_academy.md"},
		{"test-3", "AX-300 프로토콜은 모든 소스 파일을 300줄 이하로 유지하고, 표와 목록 형식만 사용하도록 강제한다.", "ax300_protocol.md"},
	}

	for _, td := range testDocs {
		vec, err := embedder.Embed(ctx, td.content)
		if err != nil {
			fmt.Printf("  ❌ 임베딩 실패 (%s): %v\n", td.id, err)
			continue
		}
		store.Add(rag.Document{
			ID: td.id, Content: td.content, Source: td.source, Vector: vec,
		})
		fmt.Printf("  ✅ '%s' 임베딩 완료 (벡터 차원: %d)\n", td.source, len(vec))
	}

	if err := store.Save(); err != nil {
		fmt.Printf("  ❌ 스토어 저장 실패: %v\n", err)
	} else {
		fmt.Println("  ✅ 벡터 스토어 디스크 저장 완료")
	}

	// 4. 시맨틱 검색 테스트
	fmt.Println("\n[4/4] 시맨틱 검색 테스트...")
	qe := rag.NewQueryEngine(embedder, store)

	questions := []string{
		"스토리에서 가치 전환이란 무엇인가?",
		"파일 크기 제한 규칙은?",
		"연기 교육 프로그램 소개",
	}

	for _, q := range questions {
		formatted, err := qe.SearchAndFormat(ctx, q, 2)
		if err != nil {
			fmt.Printf("  ❌ 검색 실패: %v\n", err)
			continue
		}
		fmt.Printf("\n  🔍 질문: %s\n", q)
		fmt.Printf("  %s\n", formatted)
	}

	fmt.Println("\n🐙 지식창고 통합 테스트 완료!")
}
