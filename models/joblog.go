package models

import (
	"time"
)

//JobLog represents a worker log
type JobLog struct {
	WorkerID    string
	JobID       string
	Message     []byte
	CreatedTime time.Time
}
