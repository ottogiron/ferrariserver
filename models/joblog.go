package models

import (
	"time"
)

//JobLog represents a worker log
type JobLog struct {
	ID          string    `json:"id"`
	JobID       string    `json:"job_id"`
	Message     string    `json:"message"`
	CreatedTime time.Time `json:"created_time"`
}
