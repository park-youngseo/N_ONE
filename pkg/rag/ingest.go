package rag

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Ingester는 문서를 벡터화하여 스토어에 투입하는 파이프라인이다.
type Ingester struct {
	embedder *Embedder
	store    *VectorStore
	chunkSize int // 청크당 최대 문자 수
	overlap   int // 청크 간 겹침 문자 수
}

// NewIngester는 문서 투입 파이프라인을 생성한다.
func NewIngester(embedder *Embedder, store *VectorStore) *Ingester {
	return &Ingester{
		embedder:  embedder,
		store:     store,
		chunkSize: 500,  // 500자 단위로 청킹
		overlap:   50,   // 50자 겹침 (문맥 유지)
	}
}

// IngestFile은 단일 파일을 읽어 벡터화하여 스토어에 저장한다.
func (ig *Ingester) IngestFile(ctx context.Context, filePath string) (int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("ingest read: %w", err)
	}

	content := string(data)
	chunks := ig.chunkText(content)
	count := 0

	for i, chunk := range chunks {
		if strings.TrimSpace(chunk) == "" {
			continue
		}
		vec, err := ig.embedder.Embed(ctx, chunk)
		if err != nil {
			return count, fmt.Errorf("ingest embed chunk %d: %w", i, err)
		}

		docID := fmt.Sprintf("%x-%d", sha256.Sum256([]byte(filePath)), i)
		doc := Document{
			ID:      docID,
			Content: chunk,
			Source:  filepath.Base(filePath),
			Vector:  vec,
			Metadata: map[string]string{
				"path":      filePath,
				"chunk_idx": fmt.Sprintf("%d", i),
			},
		}
		ig.store.Add(doc)
		count++
	}

	return count, nil
}

// IngestDir은 디렉토리 내 모든 텍스트 파일을 재귀적으로 투입한다.
func (ig *Ingester) IngestDir(ctx context.Context, dir string, exts []string) (int, error) {
	total := 0
	extSet := make(map[string]bool)
	for _, ext := range exts {
		extSet[ext] = true
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if !extSet[ext] {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		n, err := ig.IngestFile(ctx, path)
		if err != nil {
			return fmt.Errorf("ingest %s: %w", path, err)
		}
		total += n
		return nil
	})

	return total, err
}

// chunkText는 텍스트를 겹침 있는 청크로 분할한다.
func (ig *Ingester) chunkText(text string) []string {
	if len(text) <= ig.chunkSize {
		return []string{text}
	}

	var chunks []string
	for start := 0; start < len(text); {
		end := start + ig.chunkSize
		if end > len(text) {
			end = len(text)
		}
		chunks = append(chunks, text[start:end])
		start += ig.chunkSize - ig.overlap
	}
	return chunks
}
