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
	ID        string
	WorkerID  string
	StartTime time.Time
	EndTime   time.Time
	Output    []byte
}
