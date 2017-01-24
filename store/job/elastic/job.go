package elastic

import (
	"context"

	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/mattheath/kala/snowflake"
	"github.com/pkg/errors"
	oelastic "gopkg.in/olivere/elastic.v3"
)

//JobStore implementation of a job store
type jobStore struct {
	client       *oelastic.Client
	index        string
	docType      string
	idGenerator  *snowflake.Snowflake
	refreshIndex string
}

//New returns a new instance of a job store
func New(options ...Option) store.Job {
	j := &jobStore{}
	j.refreshIndex = "true"
	for _, option := range options {
		option(j)
	}

	return j
}

func (j *jobStore) Save(job *models.Job) (*models.Job, error) {
	id, err := j.idGenerator.Mint()

	if err != nil {
		return nil, errors.Wrap(err, "elastic.Store Failed to generate id for job ")
	}

	_, err = j.client.Index().
		Index(j.index).
		Type(j.docType).
		Id(id).
		BodyJson(job).
		Refresh(j.refreshIndex).
		Do(context.Background())

	if err != nil {
		return nil, errors.Wrapf(err, "elastic.Store Failed to index job %v", job)
	}

	job.ID = id
	return job, nil
}
