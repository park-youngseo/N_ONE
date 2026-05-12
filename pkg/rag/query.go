package rag

import (
	"context"
	"fmt"
	"strings"
)

// QueryEngine은 자연어 질문으로 지식창고를 검색하는 엔진이다.
type QueryEngine struct {
	embedder *Embedder
	store    *VectorStore
}

// NewQueryEngine은 시맨틱 검색 엔진을 생성한다.
func NewQueryEngine(embedder *Embedder, store *VectorStore) *QueryEngine {
	return &QueryEngine{embedder: embedder, store: store}
}

// Search는 자연어 질문을 벡터로 변환하여 유사한 문서를 검색한다.
// topK: 최대 결과 수 (기본 5), threshold: 최소 유사도 (기본 0.3)
func (qe *QueryEngine) Search(ctx context.Context, question string, topK int) ([]SearchResult, error) {
	if topK <= 0 {
		topK = 5
	}

	queryVec, err := qe.embedder.Embed(ctx, question)
	if err != nil {
		return nil, fmt.Errorf("query embed: %w", err)
	}

	results := qe.store.Search(queryVec, topK, 0.3)
	return results, nil
}

// SearchAndFormat은 검색 결과를 프롬프트에 삽입할 수 있는 텍스트로 포맷한다.
// 하네스의 판정관(GPT-5.4)에게 컨텍스트를 주입할 때 사용한다.
func (qe *QueryEngine) SearchAndFormat(ctx context.Context, question string, topK int) (string, error) {
	results, err := qe.Search(ctx, question, topK)
	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "[지식창고 검색 결과 없음]", nil
	}

	var sb strings.Builder
	sb.WriteString("## 지식창고 검색 결과\n\n")
	for i, r := range results {
		fmt.Fprintf(&sb, "### [%d] %s (유사도: %.2f%%)\n", i+1, r.Document.Source, r.Similarity*100)
		fmt.Fprintf(&sb, "%s\n\n", r.Document.Content)
	}
	return sb.String(), nil
}
