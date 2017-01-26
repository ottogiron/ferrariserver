package models

import (
	"time"
)

//Run defines a worker run
type Run struct {
	ID        string
	WorkerID  string
	StartTime time.Time
	EndTime   time.Time
}
