package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type EventType string

const (
	EventTypeCreateCompany EventType = "company.created"
	EventTypeUpdateCompany EventType = "company.updated"
)

type CompanyEvent struct {
	Type    EventType `json:"type"`
	Payload any       `json:"payload"`
}

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		log.Fatalf("Failed to connect to Kafka broker: %v", err)
	}
	defer conn.Close()

	createTopicIfNotExists(conn, topic)

	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func createTopicIfNotExists(conn *kafka.Conn, topic string) {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		log.Fatalf("Failed to read partitions: %v", err)
	}

	topicExists := false
	for _, p := range partitions {
		if p.Topic == topic {
			topicExists = true
			break
		}
	}

	if !topicExists {
		err = conn.CreateTopics(kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		})
		if err != nil {
			log.Fatalf("Failed to create topic %s: %v", topic, err)
		}
		log.Printf("Created topic %s", topic)
	} else {
		log.Printf("Topic %s already exists", topic)
	}
}

func (p *Producer) Publish(ctx context.Context, key string, eventType EventType, value interface{}) error {
	req := &CompanyEvent{
		Type:    eventType,
		Payload: value,
	}

	valueBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: valueBytes,
	}

	err = p.writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("Failed to publish message to Kafka: %v", err)
		return err
	}

	log.Printf("Published message to Kafka: %s", string(valueBytes))
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
