package acting

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/insajin/autopus-adk/pkg/acting/domain"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development
	},
	HandshakeTimeout: 5 * time.Second,
}

// WebSocketHub manages active WebSocket connections and broadcasts events
type WebSocketHub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan domain.ScriptProgressEvent
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.Mutex
}

func NewWebSocketHub() *WebSocketHub {
	return &WebSocketHub{
		broadcast:  make(chan domain.ScriptProgressEvent),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		clients:    make(map[*websocket.Conn]bool),
	}
}

// Run starts the hub loop to manage connections and broadcasting
func (h *WebSocketHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Printf("Client registered. Total clients: %d", len(h.clients))
			
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
			}
			h.mu.Unlock()
			log.Printf("Client unregistered. Total clients: %d", len(h.clients))
			
		case event := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				err := client.WriteJSON(event)
				if err != nil {
					log.Printf("WebSocket write error: %v", err)
					client.Close()
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}

// ListenToBroker connects the Hub to the EventBroker (e.g., Redis)
func (h *WebSocketHub) ListenToBroker(broker domain.EventBroker) {
	eventChan, closeFunc, err := broker.SubscribeToScriptProgress()
	if err != nil {
		log.Printf("Failed to subscribe to broker: %v", err)
		return
	}
	defer closeFunc()

	for event := range eventChan {
		// Forward events from Redis to the Hub's broadcast channel
		h.broadcast <- event
	}
}

// HandleWebSocket handles incoming HTTP requests and upgrades them to WebSocket
func (h *WebSocketHub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}

	h.register <- conn

	// Setup simple read loop to keep connection alive and detect disconnects
	// Also handles incoming Ping messages automatically by gorilla/websocket
	go func() {
		defer func() {
			h.unregister <- conn
		}()
		
		conn.SetReadLimit(512)
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		conn.SetPongHandler(func(string) error {
			conn.SetReadDeadline(time.Now().Add(60 * time.Second))
			return nil
		})

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("WebSocket error: %v", err)
				}
				break
			}
		}
	}()
}
