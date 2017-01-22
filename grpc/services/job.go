package services

import (
	"golang.org/x/net/context"

	"github.com/ferrariframework/ferrariserver/grpc/gen"
	"github.com/ferrariframework/ferrariserver/services/job"
)

var _ gen.JobServiceServer = (*JobService)(nil)

//JobService implements a grpc JobService
type JobService struct {
	jobService job.Service
}

//NewJobService returns a new jobService
func NewJobService(jobService job.Service) *JobService {
	return &JobService{jobService: jobService}
}

//RegisterJob registers a job
func (j *JobService) RegisterJob(context.Context, *gen.Job) (*gen.Job, error) {
	return nil, nil
}

//RecordLog records a job log sent from a worker
func (j *JobService) RecordLog(gen.JobService_RecordLogServer) error {
	return nil
}

//RegisterJobResult registers the result of a Job
func (j *JobService) RegisterJobResult(context.Context, *gen.JobResult) (*gen.Job, error) {
	return nil, nil
}
