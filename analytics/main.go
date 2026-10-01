package main

import (
	"analytics/internal/clickhouse"
	"analytics/internal/kafka"
	"analytics/internal/models"
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	consumer := kafka.NewConsumer(
		[]string{"localhost:9092"},
		"log_users",
		"analytics",
	)
	defer consumer.Close()

	log.Println("kafka started")
	go consumer.Start(ctx)
	ch := clickhouse.New()

	err := ch.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to ClickHouse")

	request := models.Request{
		Timestamp:  time.Now(),
		UserID:     15,
		Method:     "GET",
		Path:       "/products",
		StatusCode: 200,
		DurationMs: 42.5,
	}

	err = ch.InsertRequest(ctx, request)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Request inserted")
}
