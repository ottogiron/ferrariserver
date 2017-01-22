package job

import "github.com/ferrariframework/ferrariserver/models"
import "fmt"

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
func (j *Job) Save(job models.Job) (string, error) {
	return "", nil
}

//RecordLog records a job log
func (j *Job) RecordLog(log models.Log) error {
	fmt.Println("workerID", log.WorkerID, "jobID", log.JobID)
	fmt.Println(string(log.Message))
	return nil
}
