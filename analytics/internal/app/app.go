package app

import (
	"context"
	"fmt"
	"log"

	"analytics/internal/clickhouse"
	"analytics/internal/kafka"
)

func Run(
	ctx context.Context,
	consumer *kafka.Consumer,
	clickhouse *clickhouse.Client,
) error {

	for {
		batch, err := consumer.GetBatch(ctx, 10)
		if err != nil {
			return fmt.Errorf("get kafka batch: %w", err)
		}

		for _, event := range batch {
			log.Printf("EVENT: %+v\n", event)
		}

		log.Println("Получили сообщений:", len(batch))

		if err := clickhouse.InsertBatch(ctx, batch); err != nil {
			return fmt.Errorf("insert clickhouse: %w", err)
		}

		log.Println("Batch inserted into ClickHouse")
	}
}
