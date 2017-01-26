package models

import (
	"time"
)

//Log represents a worker log
type Log struct {
	WorkerID    string
	JobID       string
	Message     []byte
	CreatedTime time.Time
}
