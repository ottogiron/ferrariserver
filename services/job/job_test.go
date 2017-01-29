package job

import (
	"context"
	"errors"
	"testing"

	"time"

	"reflect"

	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/inconshreveable/log15"
)

var testWorkerID = "worker123"
var testWorkerIDWithError = "worker123withError"

func newJobService(recordLogs bool, logsInterval time.Duration, t *testing.T) *Job {
	logger := log15.New()
	logger.SetHandler(log15.DiscardHandler())
	j := New(
		SetContext(context.Background()),
		SetLogger(logger),
		SetJobStore(&mockJobStore{}),
		SetJobLogStore(&mockJobLogStore{}),
		SetRecordLogs(recordLogs),
		SetRecordLogsInterval(logsInterval),
	)
	return j
}

func TestJob_Save(t *testing.T) {

	j := newJobService(false, 0, t)
	type args struct {
		job *models.Job
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Save job",
			args{
				&models.Job{
					WorkerID:  testWorkerID,
					RunID:     "run123",
					StartTime: time.Now(),
				},
			},
			testWorkerID,
			false,
		},
		{
			"Save job with error",
			args{
				&models.Job{
					WorkerID:  testWorkerIDWithError,
					RunID:     "run123",
					StartTime: time.Now(),
				},
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := j.Save(tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("Job.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Job.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJob_RecordLog(t *testing.T) {
	j := newJobService(false, 0, t)
	type args struct {
		log *models.JobLog
	}
	tests := []struct {
		name             string
		args             args
		wantRecordedLogs []*models.JobLog
		wantErr          bool
	}{
		{
			"Record Log",
			args{
				&models.JobLog{
					JobID:   "jobid123",
					Message: "Some message",
				},
			},
			[]*models.JobLog{
				&models.JobLog{
					JobID:   "jobid123",
					Message: "Some message",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := j.RecordLog(tt.args.log); (err != nil) != tt.wantErr {
				t.Errorf("Job.RecordLog() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(j.recordedLogs, tt.wantRecordedLogs) {
				t.Errorf("Job.RecordLog() recordedLogs = %v, want %v", j.recordedLogs, tt.wantRecordedLogs)
			}
		})
	}
}

func TestJob_startRecordingLogs(t *testing.T) {
	logInterval := time.Duration(500)
	setLogInterval := true
	j := newJobService(setLogInterval, logInterval, t)
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			j.startRecordingLogs()
		})
	}
}

type mockJobStore struct {
	store.Job
}

func (*mockJobStore) Save(model *models.Job) (*models.Job, error) {
	if model.WorkerID == testWorkerID {
		return &models.Job{ID: testWorkerID}, nil
	} else if model.WorkerID == testWorkerIDWithError {
		return nil, errors.New("Failed to save model")
	}
	return nil, nil
}

type mockJobLogStore struct {
	store.JobLog
}
