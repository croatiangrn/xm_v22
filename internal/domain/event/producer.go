package event

import "context"

// EventType represents the type of event being published (e.g., create, update, delete).
type EventType string

const (
	TypeCreateCompany EventType = "create_company"
	TypeUpdateCompany EventType = "update_company"
	TypeDeleteCompany EventType = "delete_company"
)

// ProducerInterface defines the interface for a Kafka producer.
type ProducerInterface interface {
	Publish(ctx context.Context, topic string, eventType EventType, payload interface{}) error
}
