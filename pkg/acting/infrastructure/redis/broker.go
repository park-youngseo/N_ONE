package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/insajin/autopus-adk/pkg/acting/domain"
)

const ScriptProgressChannel = "acting:script:progress"

// RedisEventBroker implements domain.EventBroker using Redis Pub/Sub
type RedisEventBroker struct {
	client *redis.Client
}

// NewRedisEventBroker creates a new RedisEventBroker
func NewRedisEventBroker(client *redis.Client) *RedisEventBroker {
	return &RedisEventBroker{
		client: client,
	}
}

// PublishScriptProgress publishes a progress event to the Redis channel
func (b *RedisEventBroker) PublishScriptProgress(event domain.ScriptProgressEvent) error {
	ctx := context.Background() // For real app, pass proper context. Using Background for simplicity in this scaffold.
	
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return b.client.Publish(ctx, ScriptProgressChannel, payload).Err()
}

// SubscribeToScriptProgress subscribes to the Redis channel and returns a channel of events and a close function
func (b *RedisEventBroker) SubscribeToScriptProgress() (<-chan domain.ScriptProgressEvent, func(), error) {
	ctx := context.Background()
	pubsub := b.client.Subscribe(ctx, ScriptProgressChannel)

	// Verify connection
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return nil, nil, err
	}

	eventChan := make(chan domain.ScriptProgressEvent)
	
	// Start a goroutine to listen for messages
	go func() {
		ch := pubsub.Channel()
		for msg := range ch {
			var event domain.ScriptProgressEvent
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Printf("Failed to unmarshal script progress event: %v", err)
				continue
			}
			eventChan <- event
		}
	}()

	closeFunc := func() {
		pubsub.Close()
		close(eventChan)
	}

	return eventChan, closeFunc, nil
}

// MockEventBroker is an in-memory implementation for testing without Redis
type MockEventBroker struct {
	eventChan chan domain.ScriptProgressEvent
}

func NewMockEventBroker() *MockEventBroker {
	return &MockEventBroker{
		eventChan: make(chan domain.ScriptProgressEvent, 100),
	}
}

func (m *MockEventBroker) PublishScriptProgress(event domain.ScriptProgressEvent) error {
	select {
	case m.eventChan <- event:
	default:
		// Drop event if channel is full in mock
	}
	return nil
}

func (m *MockEventBroker) SubscribeToScriptProgress() (<-chan domain.ScriptProgressEvent, func(), error) {
	closeFunc := func() {
		// In a real mock with multiple subscribers, we'd need a fan-out mechanism.
		// Keeping it simple here.
	}
	return m.eventChan, closeFunc, nil
}
