package main

import (
	"analytics/internal/app"
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

	if err := app.Run(ctx, consumer, clickhouse); err != nil {
		log.Fatal(err)
	}
}
