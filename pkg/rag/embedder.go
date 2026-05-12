// Package rag는 로컬 RAG(Retrieval-Augmented Generation) 파이프라인을 제공한다.
// 올라마 임베딩 + Pure Go 벡터 스토어로 외부 의존성 없이 동작한다.
package rag

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Embedder는 올라마 임베딩 API를 호출하여 텍스트를 벡터로 변환한다.
type Embedder struct {
	baseURL string
	model   string
	client  *http.Client
}

// NewEmbedder는 올라마 기반 임베딩 클라이언트를 생성한다.
// 기본 모델: nomic-embed-text (274MB, 768차원 벡터)
func NewEmbedder() *Embedder {
	return &Embedder{
		baseURL: "http://127.0.0.1:11434",
		model:   "nomic-embed-text",
		client:  &http.Client{Timeout: 30 * time.Second},
	}
}

// Embed는 단일 텍스트를 벡터(float64 슬라이스)로 변환한다.
func (e *Embedder) Embed(ctx context.Context, text string) ([]float64, error) {
	reqBody := embeddingRequest{
		Model:  e.model,
		Prompt: text,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("embed marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", e.baseURL+"/api/embeddings", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("embed request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("embed connect: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("embed read: %w", err)
	}

	var result embeddingResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("embed unmarshal: %w", err)
	}

	if len(result.Embedding) == 0 {
		return nil, fmt.Errorf("embed: empty vector returned")
	}

	return result.Embedding, nil
}

// EmbedBatch는 여러 텍스트를 순차적으로 벡터로 변환한다.
func (e *Embedder) EmbedBatch(ctx context.Context, texts []string) ([][]float64, error) {
	vectors := make([][]float64, 0, len(texts))
	for _, text := range texts {
		vec, err := e.Embed(ctx, text)
		if err != nil {
			return nil, fmt.Errorf("batch embed: %w", err)
		}
		vectors = append(vectors, vec)
	}
	return vectors, nil
}

// Ping은 올라마 임베딩 서버가 가동 중인지 확인한다.
func (e *Embedder) Ping(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", e.baseURL+"/api/tags", nil)
	if err != nil {
		return err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return fmt.Errorf("embedder not reachable: %w", err)
	}
	defer resp.Body.Close()
	return nil
}

// --- 올라마 임베딩 API 요청/응답 구조체 ---

type embeddingRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type embeddingResponse struct {
	Embedding []float64 `json:"embedding"`
}
