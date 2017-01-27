package models

import (
	"time"
)

//JobStatus Represents a job final status
type JobStatus int

const (
	//JobStatusSucess represents a succeded job status
	JobStatusSucess JobStatus = 0
	//JobstatusFailed represents a failed job status
	JobstatusFailed JobStatus = 1
)

//Job represents a job model
type Job struct {
	ID        string    `json:"id"`
	WorkerID  string    `json:"worker_id"`
	RunID     string    `json:"run_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	JobStatus JobStatus `json:"job_status"`
	Output    string    `json:"output"`
}
