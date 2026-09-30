package models

import (
	"time"
)

type Request struct {
	Timestamp  time.Time
	UserID     int
	Method     string
	Path       string
	StatusCode uint16
	DurationMs float64
}
