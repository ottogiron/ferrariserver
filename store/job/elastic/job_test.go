package elastic

import (
	"context"
	"reflect"
	"testing"

	"github.com/ferrariframework/ferrariserver/models"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/mattheath/kala/snowflake"
	oelastic "gopkg.in/olivere/elastic.v3"
)

func newTestStore(t *testing.T) (store.Job, func()) {
	client, err := oelastic.NewClient(
		oelastic.SetSniff(false),
	)

	if err != nil {
		t.Fatal("Failed to create new store elastic client ", err)
	}

	index := "test_index"
	docType := "test_doc_type"

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

func Test_jobStore_Save(t *testing.T) {
	j, clean := newTestStore(t)
	defer clean()
	type args struct {
		job *models.Job
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Job
		wantErr bool
	}{
		{"Save", args{&models.Job{WorkerID: "worker134"}}, &models.Job{WorkerID: "worker134"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := j.Save(tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobStore.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID == "" {
				t.Errorf("jobStore.Save() = got.ID = %v, want an id value", got.ID)
			}

			if !reflect.DeepEqual(got.Output, tt.want.Output) {
				t.Errorf("jobStore.Save() = got.Output = %v, want %v", got.Output, tt.want.Output)
			}

			if got.WorkerID != tt.want.WorkerID {
				t.Errorf("jobStore.Save() = got.WorkerID = %v, want %v", got.WorkerID, tt.want.WorkerID)
			}

			if got.StartTime != tt.want.StartTime {
				t.Errorf("jobStore.Save() = got.StartTime = %v, want %v", got.StartTime, tt.want.StartTime)
			}

			if got.EndTime != tt.want.EndTime {
				t.Errorf("jobStore.Save() = got.EndTime = %v, want %v", got.EndTime, tt.want.EndTime)
			}
		})
	}
}

func Test_jobStore_Get(t *testing.T) {
	type fields struct {
		client       *oelastic.Client
		index        string
		docType      string
		idGenerator  *snowflake.Snowflake
		refreshIndex string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Job
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobStore{
				client:       tt.fields.client,
				index:        tt.fields.index,
				docType:      tt.fields.docType,
				idGenerator:  tt.fields.idGenerator,
				refreshIndex: tt.fields.refreshIndex,
			}
			got, err := j.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobStore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jobStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jobStore_Update(t *testing.T) {
	type fields struct {
		client       *oelastic.Client
		index        string
		docType      string
		idGenerator  *snowflake.Snowflake
		refreshIndex string
	}
	type args struct {
		id  string
		job *models.Job
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Job
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobStore{
				client:       tt.fields.client,
				index:        tt.fields.index,
				docType:      tt.fields.docType,
				idGenerator:  tt.fields.idGenerator,
				refreshIndex: tt.fields.refreshIndex,
			}
			got, err := j.Update(tt.args.id, tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jobStore.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
