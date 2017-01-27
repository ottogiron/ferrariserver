package models

import (
	"time"
)

//Run defines a worker run
type Run struct {
	ID        string    `json:"id"`
	WorkerID  string    `json:"worker_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
