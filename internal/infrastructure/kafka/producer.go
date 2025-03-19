package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) Publish(ctx context.Context, key string, value interface{}) error {
	valueBytes, err := json.Marshal(value)
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
