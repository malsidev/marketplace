package kafka

import (
	"analytics/internal/models"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic string, groupID string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})

	return &Consumer{
		reader: reader,
	}
}
func (c *Consumer) GetBatch(ctx context.Context, size int) ([]models.Request, error) {
	batch := make([]models.Request, 0, size)
	log.Println(batch)
	for len(batch) < size {
		message, err := c.reader.ReadMessage(ctx)
		log.Println("RAW KAFKA:", string(message.Value))
		if err != nil {
			return nil, err
		}

		var event models.Request

		err = json.Unmarshal(message.Value, &event)
		if err != nil {
			return nil, err
		}
		batch = append(batch, event)
	}
	return batch, nil
}
