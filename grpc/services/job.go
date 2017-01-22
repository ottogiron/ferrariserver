package services

import (
	"golang.org/x/net/context"

	"io"

	"github.com/ferrariframework/ferrariserver/grpc/gen"
	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/services/job"
	"github.com/pkg/errors"
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
func (j *JobService) RecordLog(stream gen.JobService_RecordLogServer) error {
	for {
		jobLog, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(nil)
		}

		if err != nil {
			return errors.Wrap(err, "grpc.JobService Failed to record log")
		}
		return j.jobService.RecordLog(models.Log{
			WorkerID: jobLog.WorkerId,
			JobID:    jobLog.JobId,
			Message:  jobLog.Message,
		})
	}

}

//RegisterJobResult registers the result of a Job
func (j *JobService) RegisterJobResult(context.Context, *gen.JobResult) (*gen.Job, error) {
	return nil, nil
}
