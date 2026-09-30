package clickhouse

import (
	"analytics/internal/models"
	"context"
)

func (c *Client) InsertRequest(
	ctx context.Context,
	request models.Request,
) error {
	return c.conn.Exec(
		ctx,
		`
		INSERT INTO requests
		(
			timestamp,
			user_id,
			method,
			path,
			status_code,
			duration_ms
		)
		VALUES (?, ?, ?, ?, ?, ?)
		`,
		request.Timestamp,
		request.UserID,
		request.Method,
		request.Path,
		request.StatusCode,
		request.DurationMs,
	)
}
