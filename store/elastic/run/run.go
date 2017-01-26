package worker

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

//RunStore implementation of a run store
type runStore struct {
	client       *oelastic.Client
	index        string
	docType      string
	idGenerator  *snowflake.Snowflake
	refreshIndex string
}

//New returns a new instance of a run store
func New(options ...Option) store.Run {

	r := &runStore{
		refreshIndex: "true",
	}

	for _, option := range options {
		option(r)
	}

	return r
}

func (r *runStore) Save(run *models.Run) (*models.Run, error) {
	id, err := r.idGenerator.Mint()

	if err != nil {
		return nil, errors.Wrap(err, "elastic.Store Failed to generate id for job ")
	}

	_, err = r.client.Index().
		Index(r.index).
		Type(r.docType).
		Refresh(r.refreshIndex).
		Id(id).
		BodyJson(run).
		Do(context.Background())

	if err != nil {
		return nil, errors.Wrapf(err, "elastic.Store Failed to index job %v", run)
	}

	run.ID = id
	return run, nil
}

func (r *runStore) Get(id string) (*models.Run, error) {
	res, err := r.client.Get().
		Index(r.index).
		Type(r.docType).
		Refresh(r.refreshIndex).
		Id(id).
		Do(context.Background())

	if err != nil {
		if oelastic.IsNotFound(err) {
			return nil, errors.Wrapf(errortypes.NewNotFound(), "elastic.Store.Get(%id)", id)
		}
		return nil, errors.Wrapf(err, "elastic.Store Failed to get jobID=%s", id)
	}

	var job models.Run

	err = json.Unmarshal(*res.Source, &job)

	if err != nil {
		return nil, errors.Wrapf(err, "elastic.Store.Get(%id)", id)
	}
	job.ID = res.Id
	return &job, nil
}

func (r *runStore) Update(id string, worker *models.Run) error {

	_, err := r.client.Update().
		Index(r.index).
		Type(r.docType).
		Refresh(r.refreshIndex).
		Id(id).
		Doc(worker).
		Do(context.Background())

	if err != nil {
		if oelastic.IsNotFound(err) {
			return errors.Wrapf(errortypes.NewNotFound(), "elastic.Store.Update(%id, %v)", worker)
		}
		return errors.Wrapf(err, "elastic.Store.Update(%id, %v) ", id, worker)
	}

	return nil
}
