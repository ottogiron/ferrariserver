package store

import "github.com/ferrariframework/ferrariserver/models"

//JobLog defines an interface for admnistering JobLog models
type JobLog interface {
	Save(logs []*models.JobLog) error
}
