package clickhouse

import (
	"context"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type Client struct {
	conn clickhouse.Conn
}

func New() *Client {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9002"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "default",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return &Client{conn: conn}
}

func (c *Client) Ping(ctx context.Context) error {
	return c.conn.Ping(ctx)
}
