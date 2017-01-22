package job

import "github.com/ferrariframework/ferrariserver/models"

var _ Service = (*Job)(nil)

//Service defines a job service
type Service interface {
	Save(models.Job) (string, error)
	RecordLog(models.Log) error
}

//New retuns a new job service
func New() *Job {
	return &Job{}
}

//Job a job service implementation
type Job struct {
}

//Save saves a new job
func (j *Job) Save(models.Job) (string, error) {
	return "", nil
}

//RecordLog records a job log
func (j *Job) RecordLog(models.Log) error {
	return nil
}
