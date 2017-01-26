package worker

import oelastic "gopkg.in/olivere/elastic.v3"
import "github.com/mattheath/kala/snowflake"

//Option defines an elastic job store option
type Option func(*runStore)

//SetClient sets an elastic client for this store
func SetClient(client *oelastic.Client) Option {
	return func(runStore *runStore) {
		runStore.client = client
	}
}

//SetIndex sets an elastic index for this store
func SetIndex(index string) Option {
	return func(runStore *runStore) {
		runStore.index = index
	}
}

//SetDocType sets an elastic docType for this store
func SetDocType(docType string) Option {
	return func(runStore *runStore) {
		runStore.docType = docType
	}
}

//SetIDGenerator sets an snoflake ID generator
func SetIDGenerator(idGenerator *snowflake.Snowflake) Option {
	return func(runStore *runStore) {
		runStore.idGenerator = idGenerator
	}
}

//SetRefreshIndex sets if the index should be immediatly refreshed
func SetRefreshIndex(refresh string) Option {
	return func(runStore *runStore) {
		runStore.refreshIndex = refresh
	}
}
