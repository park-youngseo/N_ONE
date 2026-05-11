package domain

// ScriptProgressEvent represents a message published when a script's progress changes
type ScriptProgressEvent struct {
	ScriptID    string `json:"script_id"`
	SceneID     string `json:"scene_id"`
	ProgressPct int    `json:"progress_percent"`
	Timestamp   int64  `json:"timestamp"`
}

// EventBroker defines the interface for publishing and subscribing to events
type EventBroker interface {
	PublishScriptProgress(event ScriptProgressEvent) error
	SubscribeToScriptProgress() (<-chan ScriptProgressEvent, func(), error)
}
