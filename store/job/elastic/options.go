package elastic

import oelastic "gopkg.in/olivere/elastic.v3"
import "github.com/mattheath/kala/snowflake"

//Option defines an elastic job store option
type Option func(*jobStore)

//Client sets an elastic client for this store
func Client(client *oelastic.Client) Option {
	return func(jobStore *jobStore) {
		jobStore.client = client
	}
}

//Index sets an elastic index for this store
func Index(index string) Option {
	return func(jobStore *jobStore) {
		jobStore.index = index
	}
}

//DocType sets an elastic docType for this store
func DocType(docType string) Option {
	return func(jobStore *jobStore) {
		jobStore.docType = docType
	}
}

//IDGenerator sets an snoflake ID generator
func IDGenerator(idGenerator *snowflake.Snowflake) Option {
	return func(jobStore *jobStore) {
		jobStore.idGenerator = idGenerator
	}
}

//RefreshIndex sets if the index should be immediatly refreshed
func RefreshIndex(refresh string) Option {
	return func(jobStore *jobStore) {
		jobStore.refreshIndex = refresh
	}
}
