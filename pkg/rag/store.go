package rag

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"sync"
)

// Document는 벡터 스토어에 저장되는 단일 문서이다.
type Document struct {
	ID       string    `json:"id"`
	Content  string    `json:"content"`
	Source   string    `json:"source"`
	Vector   []float64 `json:"vector"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// SearchResult는 시맨틱 검색의 단일 결과이다.
type SearchResult struct {
	Document   Document
	Similarity float64
}

// VectorStore는 Pure Go 임베디드 벡터 데이터베이스이다.
// 외부 DB 서버 없이 로컬 JSON 파일로 영구 저장한다.
type VectorStore struct {
	mu       sync.RWMutex
	docs     []Document
	filePath string
}

// NewVectorStore는 디스크 기반 벡터 스토어를 생성한다.
// 기존 데이터가 있으면 자동으로 로드한다.
func NewVectorStore(filePath string) (*VectorStore, error) {
	store := &VectorStore{filePath: filePath}
	if err := store.load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("store load: %w", err)
	}
	return store, nil
}

// Add는 문서를 스토어에 추가한다.
func (s *VectorStore) Add(doc Document) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 동일 ID 문서가 존재하면 갱신
	for i, existing := range s.docs {
		if existing.ID == doc.ID {
			s.docs[i] = doc
			return
		}
	}
	s.docs = append(s.docs, doc)
}

// Search는 쿼리 벡터와 가장 유사한 문서를 반환한다.
// topK: 반환할 최대 결과 수, threshold: 최소 유사도 (0.0~1.0)
func (s *VectorStore) Search(queryVec []float64, topK int, threshold float64) []SearchResult {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []SearchResult
	for _, doc := range s.docs {
		sim := cosineSimilarity(queryVec, doc.Vector)
		if sim >= threshold {
			results = append(results, SearchResult{Document: doc, Similarity: sim})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Similarity > results[j].Similarity
	})

	if len(results) > topK {
		results = results[:topK]
	}
	return results
}

// Save는 전체 스토어를 디스크에 영구 저장한다.
func (s *VectorStore) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	data, err := json.MarshalIndent(s.docs, "", "  ")
	if err != nil {
		return fmt.Errorf("store marshal: %w", err)
	}
	return os.WriteFile(s.filePath, data, 0644)
}

// Count는 저장된 문서 수를 반환한다.
func (s *VectorStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.docs)
}

// load는 디스크에서 기존 데이터를 메모리로 로드한다.
func (s *VectorStore) load() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.docs)
}

// cosineSimilarity는 두 벡터 간의 코사인 유사도를 계산한다.
// 반환값: -1.0(정반대) ~ 1.0(완전 동일)
func cosineSimilarity(a, b []float64) float64 {
	if len(a) != len(b) || len(a) == 0 {
		return 0
	}
	var dot, normA, normB float64
	for i := range a {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	denom := math.Sqrt(normA) * math.Sqrt(normB)
	if denom == 0 {
		return 0
	}
	return dot / denom
}
