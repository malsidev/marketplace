package kafka

import (
	"context"
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


func (c *Consumer) Start(ctx context.Context) {
	for {
		message, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Kafka err", err)
			return
		}

		log.Println(
			"kafla message:",
			message.Topic,
			message.Partition,
			message.Offset,
			string(message.Value),
		)
	}
}
func (c *Consumer) Close() error {
	return c.reader.Close()
}
