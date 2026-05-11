package e2e_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/insajin/autopus-adk/internal/api/handler/acting"
	"github.com/insajin/autopus-adk/pkg/acting/domain"
	"github.com/insajin/autopus-adk/pkg/acting/infrastructure/redis"
)

func TestWebSocketSync_E2E(t *testing.T) {
	// 1. Setup mock EventBroker and WebSocket Hub
	mockBroker := redis.NewMockEventBroker()
	hub := acting.NewWebSocketHub()
	
	// Start hub routing
	go hub.Run()
	
	// Connect Hub to Broker (subscribe to Redis channel in real env)
	go hub.ListenToBroker(mockBroker)

	// 2. Setup HTTP test server
	server := httptest.NewServer(http.HandlerFunc(hub.HandleWebSocket))
	defer server.Close()

	// Convert http:// to ws://
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// 3. Connect Client A
	connA, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	defer connA.Close()

	// 4. Connect Client B
	connB, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	defer connB.Close()

	// Allow time for clients to register
	time.Sleep(100 * time.Millisecond)

	// 5. Trigger an event (simulating someone completing a scene)
	testEvent := domain.ScriptProgressEvent{
		ScriptID:    "script-1",
		SceneID:     "scene-1",
		ProgressPct: 75,
		Timestamp:   time.Now().Unix(),
	}
	
	err = mockBroker.PublishScriptProgress(testEvent)
	assert.NoError(t, err)

	// 6. Verify Client A receives it
	connA.SetReadDeadline(time.Now().Add(2 * time.Second))
	_, msgA, err := connA.ReadMessage()
	assert.NoError(t, err)
	
	var receivedA domain.ScriptProgressEvent
	err = json.Unmarshal(msgA, &receivedA)
	assert.NoError(t, err)
	assert.Equal(t, 75, receivedA.ProgressPct)

	// 7. Verify Client B receives it simultaneously
	connB.SetReadDeadline(time.Now().Add(2 * time.Second))
	_, msgB, err := connB.ReadMessage()
	assert.NoError(t, err)
	
	var receivedB domain.ScriptProgressEvent
	err = json.Unmarshal(msgB, &receivedB)
	assert.NoError(t, err)
	assert.Equal(t, 75, receivedB.ProgressPct)
}
