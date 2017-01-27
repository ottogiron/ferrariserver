package worker

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/mattheath/kala/snowflake"
	oelastic "gopkg.in/olivere/elastic.v3"
)

func newTestStore(t *testing.T) (store.JobLog, func()) {
	client, err := oelastic.NewClient(
		oelastic.SetSniff(false),
	)

	if err != nil {
		t.Fatal("Failed to create new store elastic client ", err)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := "test_index" + strconv.Itoa(r.Int())
	fmt.Println("Test index name:", index)
	docType := "test_doc_type"

	if err != nil {
		t.Fatal("Failed to put Environment mapping ", err)
	}

	idGenerator, err := snowflake.New(100)
	if err != nil {
		t.Fatal("Failed to create new store snowflake generator")
	}
	s := New(
		SetClient(client),
		SetIndex(index),
		SetDocType(docType),
		SetIDGenerator(idGenerator),
		SetRefreshIndex("true"),
	)
	return s, func() {
		client.DeleteIndex(index).
			Do(context.Background())
	}
}

func Test_jobLogStore_Save(t *testing.T) {
	j, clean := newTestStore(t)
	defer clean()

	type args struct {
		logs []*models.JobLog
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Bulk save", args{
			[]*models.JobLog{
				&models.JobLog{
					JobID:   "job123",
					Message: "Some message",
				},
				&models.JobLog{
					JobID:   "job345",
					Message: "Some message",
				},
				&models.JobLog{
					JobID:   "job3435",
					Message: "Some message",
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := j.Save(tt.args.logs); (err != nil) != tt.wantErr {
				t.Errorf("jobLogStore.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
