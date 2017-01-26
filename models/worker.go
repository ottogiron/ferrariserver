package models

import (
	"time"
)

//Worker model for Worker
type Worker struct {
	ID          string
	Parallel    int
	Environment map[string]string
	CreatedTime time.Time
	UpdatedTime time.Time
}