package models

import (
	"time"
)

//JobLog represents a worker log
type JobLog struct {
	ID          string
	JobID       string
	Message     string
	CreatedTime time.Time
}
