package job

import (
	"context"
	"sync"

	"time"

	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
)

var _ Service = (*Job)(nil)

//Service defines a job service
type Service interface {
	Save(*models.Job) (string, error)
	RecordLog(*models.JobLog) error
}

//New retuns a new job service
func New(ctx context.Context, logger log15.Logger, jobStore store.Job, recordLogs bool, recordLogInterval time.Duration) *Job {
	service := &Job{
		jobStore:          jobStore,
		ctx:               ctx,
		recordLogInterval: recordLogInterval,
		logger:            logger,
	}
	if recordLogs {
		service.startRecordingLogs()
	}

	return service
}

//Job a job service implementation
type Job struct {
	mu                sync.Mutex
	jobStore          store.Job
	recordedLogs      []*models.JobLog
	ctx               context.Context
	recordLogInterval time.Duration
	logger            log15.Logger
}

//Save saves a new job
func (j *Job) Save(job *models.Job) (string, error) {
	r, err := j.jobStore.Save(job)
	if err != nil {
		err = errors.Wrap(err, "Failed to save job")
		return "", err
	}
	return r.ID, nil
}

//RecordLog records a job log
func (j *Job) RecordLog(log *models.JobLog) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.recordedLogs = append(j.recordedLogs, log)
	return nil
}

func (j *Job) startRecordingLogs() {
	go func() {
		c := time.Tick(j.recordLogInterval)
		for {
			select {
			case <-c:
			//Record logs in batch
			case <-j.ctx.Done():
				return
			}
		}

	}()
}
