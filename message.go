package pubsub

import "encoding/json"

// Message is an envelop for a JSON message with lazy unmarshalling for the
// payload.
type Message struct {
	Topic   string          `json:"topic"`
	Payload json.RawMessage `json:"payload"`
}
