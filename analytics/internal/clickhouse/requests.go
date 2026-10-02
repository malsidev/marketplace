package clickhouse

import (
	"analytics/internal/models"
	"context"
)

func (c *Client) InsertBatch(
	ctx context.Context,
	events []models.Request,
) error {
	batch, err := c.conn.PrepareBatch(
		ctx,
		"INSERT INTO requests",
	)
	if err != nil {
		return err
	}

	for _, event := range events {
		err := batch.Append(
			event.Timestamp,
			event.UserID,
			event.Method,
			event.Path,
			event.StatusCode,
			event.DurationMs,
		)

		if err != nil {
			return err
		}
	}

	return batch.Send()
}
