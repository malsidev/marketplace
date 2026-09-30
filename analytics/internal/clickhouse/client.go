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
		Addr: []string{"clickhouse:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
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
