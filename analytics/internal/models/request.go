package models

import (
	"time"
)

type Request struct {
	Timestamp  time.Time
	UserID     int
	Method     string
	Path       string
	StatusCode int
	DurationMs float64
}
type LogEvent struct {
	Timestamp  time.Time
	UserID     string
	Method     string
	Path       string
	StatusCode int
	DurationMs float64
}
