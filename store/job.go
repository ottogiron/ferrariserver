package store

import "github.com/ferrariframework/ferrariserver/models"

//Job a job store interface
type Job interface {
	Save(*models.Job) (*models.Job, error)
}
