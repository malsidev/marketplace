package main

import (
	"analytics/internal/kafka"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	consumer := kafka.NewConsumer(
		[]string{"localhost:9092"},
		"actions",
		"analytics-servis",
	)
	defer consumer.Close()

	log.Println("kafka started")
	consumer.Start(ctx)
}
