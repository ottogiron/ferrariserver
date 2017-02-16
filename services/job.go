package services

import "github.com/ferrariframework/ferrariserver/models"

//Job defines a job service
type Job interface {
	Save(*models.Job) (string, error)
	RecordLog(*models.JobLog) error
}
