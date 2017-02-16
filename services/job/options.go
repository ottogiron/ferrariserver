package job

import (
	"context"

	"time"

	"github.com/ferrariframework/ferrariserver/store"
	"github.com/inconshreveable/log15"
)

//Option represents a JobService option
type Option func(jobService *job)

//SetContext sets the context of this  jobService
func SetContext(ctx context.Context) Option {
	return func(jobService *job) {
		jobService.ctx = ctx
	}
}

//SetLogger sets the logger of this jobService
func SetLogger(logger log15.Logger) Option {
	return func(jobService *job) {
		jobService.logger = logger
	}
}

//SetJobStore sets the jobStore of this service
func SetJobStore(jobStore store.Job) Option {
	return func(jobService *job) {
		jobService.jobStore = jobStore
	}
}

//SetJobLogStore sets the jobLogStore of this service
func SetJobLogStore(jobLogStore store.JobLog) Option {
	return func(jobService *job) {
		jobService.jobLogStore = jobLogStore
	}
}

//SetRecordLogs sets if the service should record logs
func SetRecordLogs(recordLogs bool) Option {
	return func(jobService *job) {
		jobService.recordLogs = recordLogs
	}
}

//SetRecordLogsInterval sets the interval to record logs
func SetRecordLogsInterval(recordLogsInterval time.Duration) Option {
	return func(jobService *job) {
		jobService.recordLogInterval = recordLogsInterval
	}
}
