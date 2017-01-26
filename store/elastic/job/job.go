package job

import (
	"context"

	"encoding/json"

	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/ferrariframework/ferrariserver/store/errortypes"
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

	j := &jobStore{
		refreshIndex: "true",
	}

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
		Refresh(j.refreshIndex).
		Id(id).
		BodyJson(job).
		Do(context.Background())

	if err != nil {
		return nil, errors.Wrapf(err, "elastic.Store Failed to index job %v", job)
	}

	job.ID = id
	return job, nil
}

func (j *jobStore) Get(id string) (*models.Job, error) {
	res, err := j.client.Get().
		Index(j.index).
		Type(j.docType).
		Refresh(j.refreshIndex).
		Id(id).
		Do(context.Background())

	if err != nil {
		if oelastic.IsNotFound(err) {
			return nil, errors.Wrapf(errortypes.NewNotFound(), "elastic.Store.Get(%id)", id)
		}
		return nil, errors.Wrapf(err, "elastic.Store Failed to get jobID=%s", id)
	}

	var job models.Job

	err = json.Unmarshal(*res.Source, &job)

	if err != nil {
		return nil, errors.Wrapf(err, "elastic.Store.Get(%id)", id)
	}
	job.ID = res.Id
	return &job, nil
}

func (j *jobStore) Update(id string, job *models.Job) error {

	_, err := j.client.Update().
		Index(j.index).
		Type(j.docType).
		Refresh(j.refreshIndex).
		Id(id).
		Doc(job).
		Do(context.Background())

	if err != nil {
		if oelastic.IsNotFound(err) {
			return errors.Wrapf(errortypes.NewNotFound(), "elastic.Store.Update(%id, %v)", job)
		}
		return errors.Wrapf(err, "elastic.Store.Update(%id, %v) ", id, job)
	}

	return nil
}
