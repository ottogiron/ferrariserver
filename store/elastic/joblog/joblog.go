package worker

import (
	"context"

	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/mattheath/kala/snowflake"
	"github.com/pkg/errors"
	oelastic "gopkg.in/olivere/elastic.v3"
)

//jobLogStore implementation of a job store
type jobLogStore struct {
	client       *oelastic.Client
	index        string
	docType      string
	idGenerator  *snowflake.Snowflake
	refreshIndex string
	ctx          context.Context
}

//New returns a new instance of a worker store
func New(options ...Option) store.JobLog {

	j := &jobLogStore{
		refreshIndex: "true",
		ctx:          context.Background(),
	}

	for _, option := range options {
		option(j)
	}

	return j
}

func (j *jobLogStore) Save(logs []*models.JobLog) error {
	bulkService := j.client.Bulk().Index(j.index).Type(j.docType)

	for _, log := range logs {
		id, err := j.idGenerator.Mint()
		if err != nil {
			return errors.Wrap(err, "Failed to generate id when saving logs")
		}
		log.ID = id
		r := oelastic.NewBulkIndexRequest().Id(id).Doc(log)
		bulkService.Add(r)
	}

	r, err := bulkService.Do(j.ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to do bulk request for storing logs")
	}
	if r.Errors {
		return errors.New("Some of the stored logs failed to be saved")
	}
	return nil
}
