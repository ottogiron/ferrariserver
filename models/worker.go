package models

import (
	"time"
)

//Worker model for Worker
type Worker struct {
	ID          string            `json:"id"`
	Parallel    int               `json:"parallel"`
	Environment map[string]string `json:"environment"`
	CreatedTime time.Time         `json:"created_time"`
	UpdatedTime time.Time         `json:"updated_time"`
}
