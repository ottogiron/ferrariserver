package config

import (
	"context"
	"math/rand"
	"time"

	"github.com/ferrariframework/ferrariserver/store"
	jobstore "github.com/ferrariframework/ferrariserver/store/elastic/job"
	joblog "github.com/ferrariframework/ferrariserver/store/elastic/joblog"
	"github.com/inconshreveable/log15"
	"github.com/mattheath/kala/snowflake"
	"github.com/pkg/errors"
	elastic "gopkg.in/olivere/elastic.v3"
)

//JobStore configures a new instance of a job store
func JobStore(ctx context.Context, index string, docType string, client *elastic.Client, idGenerator *snowflake.Snowflake) (store.Job, error) {

	err := createIndexIfDontExists(ctx, index, client)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create elastic index %s", index)
	}

	_, err = client.PutMapping().
		Index(index).
		Type(docType).
		BodyJson(map[string]interface{}{
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
				"worker_id": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
				"run_id": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
			},
		}).
		Do(ctx)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to put mappings for index %s and type %s", index, docType)
	}

	store := jobstore.New(
		jobstore.SetContext(ctx),
		jobstore.SetClient(client),
		jobstore.SetIndex(index),
		jobstore.SetDocType(docType),
		jobstore.SetIDGenerator(idGenerator),
		jobstore.SetRefreshIndex("true"),
	)
	return store, nil
}

//JobLogStore configures a new instance of a job store
func JobLogStore(ctx context.Context, index string, docType string, client *elastic.Client, idGenerator *snowflake.Snowflake) (store.JobLog, error) {

	err := createIndexIfDontExists(ctx, index, client)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create elastic index %s", index)
	}
	_, err = client.PutMapping().
		Index(index).
		Type(docType).
		BodyJson(map[string]interface{}{
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
				"job_id": map[string]interface{}{
					"type":  "string",
					"index": "not_analyzed",
				},
			},
		}).
		Do(ctx)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to put mappings for index %s and type %s", index, docType)
	}

	store := joblog.New(
		joblog.SetContext(ctx),
		joblog.SetClient(client),
		joblog.SetIndex(index),
		joblog.SetDocType(docType),
		joblog.SetIDGenerator(idGenerator),
		joblog.SetRefreshIndex("true"),
	)
	return store, nil
}

//SnowFlakeGenerator creates a new instance of an snowflake generator
func SnowFlakeGenerator() (*snowflake.Snowflake, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	workerID := r1.Uint32()
	generator, err := snowflake.New(workerID)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create new snowflake generator")
	}
	return generator, nil
}

func createIndexIfDontExists(ctx context.Context, name string, client *elastic.Client) error {

	exists, err := client.IndexExists(name).Do(ctx)

	if err != nil {
		return errors.Wrapf(err, "Failed to validate index existance %s", name)
	}
	if !exists {
		_, err = client.CreateIndex(name).
			BodyJson(map[string]interface{}{
				"settings": map[string]interface{}{
					"number_of_shards": 1,
				},
			}).
			Do(ctx)
		if err != nil {
			return errors.Wrapf(err, "Failed to create index with name %s", name)
		}
	} else {
		log15.Warn("Elastic Index already exists skipping creation", "name", name)
	}

	return nil
}
