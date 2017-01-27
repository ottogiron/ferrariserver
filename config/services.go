package config

import (
	"context"

	elastic "gopkg.in/olivere/elastic.v3"

	"time"

	jobservice "github.com/ferrariframework/ferrariserver/services/job"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
)

//ElasticClient returns a new instance of an elastic client
func ElasticClient(setSniff bool, urls ...string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(urls...),
		elastic.SetSniff(setSniff),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create elastic client urls=%v setSniff=%v", urls, setSniff)
	}
	return client, nil
}

//JobService Configures a new instance of a job service
func JobService(ctx context.Context, logger log15.Logger, jobStore store.Job, jobLogStore store.JobLog, recordLogs bool, recordLogsInterval time.Duration) jobservice.Service {
	clogger := logger.New("service", "job")

	return jobservice.New(
		jobservice.SetContext(ctx),
		jobservice.SetLogger(clogger),
		jobservice.SetJobStore(jobStore),
		jobservice.SetJobLogStore(jobLogStore),
		jobservice.SetRecordLogs(recordLogs),
		jobservice.SetRecordLogsInterval(recordLogsInterval),
	)
}
