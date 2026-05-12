package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"
)

// OllamaAdapter implements ProviderAdapter for local Ollama API.
// 로컬 GPU에서 실행되는 올라마를 하네스의 전처리 전담 노동자로 연결한다.
type OllamaAdapter struct {
	baseURL string
	model   string
	client  *http.Client
}

// NewOllamaAdapter creates a new OllamaAdapter.
// 기본 엔드포인트: http://127.0.0.1:11434
// 기본 모델: qwen2.5:14b (코딩 특화, RTX 3060 최적)
func NewOllamaAdapter() *OllamaAdapter {
	return &OllamaAdapter{
		baseURL: "http://127.0.0.1:11434",
		model:   "qwen2.5:14b",
		client:  &http.Client{Timeout: 120 * time.Second},
	}
}

// Name returns "ollama".
func (a *OllamaAdapter) Name() string { return "ollama" }

// BuildCommand은 올라마 전용으로 사용하지 않음 (HTTP API 기반).
// 호환성을 위해 더미 커맨드를 반환한다.
func (a *OllamaAdapter) BuildCommand(ctx context.Context, task TaskConfig) *exec.Cmd {
	return exec.CommandContext(ctx, "echo", "ollama-uses-http-api")
}

// Generate sends a prompt to Ollama and returns the response.
// temperature=0.0 으로 환각을 최소화하며, 단순 전처리 작업에 최적화.
func (a *OllamaAdapter) Generate(ctx context.Context, prompt, systemPrompt string) (string, error) {
	model := a.model

	reqBody := ollamaRequest{
		Model:  model,
		Prompt: prompt,
		System: systemPrompt,
		Stream: false,
		Options: ollamaOptions{
			Temperature: 0.0, // 환각 방지: 창의성 0
			NumCtx:      4096,
		},
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("ollama marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/api/generate", bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("ollama request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama connect: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ollama read: %w", err)
	}

	var result ollamaResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("ollama unmarshal: %w", err)
	}

	return result.Response, nil
}

// ParseEvent parses Ollama output into a StreamEvent.
func (a *OllamaAdapter) ParseEvent(line []byte) (StreamEvent, error) {
	var raw struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(line, &raw); err != nil {
		return StreamEvent{}, fmt.Errorf("ollama parse: %w", err)
	}
	typ, subtype := splitEventType(raw.Type)
	return StreamEvent{
		Type:    typ,
		Subtype: subtype,
		Data:    json.RawMessage(append([]byte(nil), line...)),
	}, nil
}

// ExtractResult extracts the task result from an Ollama event.
func (a *OllamaAdapter) ExtractResult(event StreamEvent) TaskResult {
	return TaskResult{
		CostUSD:    0.0, // 로컬 처리: 비용 0원
		DurationMS: 0,
		Output:     string(event.Data),
	}
}

// Ping checks if Ollama server is running.
func (a *OllamaAdapter) Ping(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", a.baseURL+"/api/tags", nil)
	if err != nil {
		return err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("ollama not reachable: %w", err)
	}
	defer resp.Body.Close()
	return nil
}

// --- 올라마 API 요청/응답 구조체 ---

type ollamaRequest struct {
	Model   string        `json:"model"`
	Prompt  string        `json:"prompt"`
	System  string        `json:"system,omitempty"`
	Stream  bool          `json:"stream"`
	Options ollamaOptions `json:"options,omitempty"`
}

type ollamaOptions struct {
	Temperature float64 `json:"temperature"`
	NumCtx      int     `json:"num_ctx,omitempty"`
}

type ollamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}
