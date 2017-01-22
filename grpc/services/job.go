package services

import (
	"golang.org/x/net/context"

	"io"

	"time"

	"github.com/ferrariframework/ferrariserver/grpc/gen"
	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/services/job"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
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
func (j *JobService) RegisterJob(ctx context.Context, job *gen.Job) (*gen.Job, error) {
	return &gen.Job{
		WorkerId:  job.WorkerId,
		Id:        uuid.NewV4().String(),
		StartTime: time.Now().Unix(),
	}, nil
}

//RecordLog records a job log sent from a worker
func (j *JobService) RecordLog(stream gen.JobService_RecordLogServer) error {
	for {
		jobLog, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&gen.Empty{})
		}

		if err != nil {
			return errors.Wrap(err, "grpc.JobService Failed to record log")
		}

		j.jobService.RecordLog(models.Log{
			WorkerID: jobLog.WorkerId,
			JobID:    jobLog.JobId,
			Message:  jobLog.Message,
		})
	}

}

//RegisterJobResult registers the result of a Job
func (j *JobService) RegisterJobResult(ctx context.Context, result *gen.JobResult) (*gen.Job, error) {
	return &gen.Job{
		WorkerId: result.WorkerId,
		Id:       result.JobId,
	}, nil
}
