package main

import (
	"analytics/internal/clickhouse"
	"analytics/internal/kafka"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	consumer := kafka.NewConsumer(
		[]string{"localhost:9092"},
		"log_users",
		"analytics",
	)
	clickhouse := clickhouse.New()

	for {

		batch, err := consumer.GetBatch(ctx, 10)
		if err != nil {
			log.Println("Kafka error:", err)
			return
		}

		for _, event := range batch {
			log.Printf("EVENT: %+v\n", event)
		}
		log.Println("Получили сообщений:", len(batch))

		err = clickhouse.InsertBatch(ctx, batch)
		if err != nil {
			log.Println("ClickHouse error:", err)
			return
		}

		log.Println("Batch inserted into ClickHouse")
	}
}
