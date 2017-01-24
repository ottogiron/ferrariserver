package elastic

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"

	"time"

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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := "test_index" + strconv.Itoa(r.Int())
	fmt.Println("Test index name:", index)
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

	j, clean := newTestStore(t)
	defer clean()

	testSaveJob := &models.Job{
		WorkerID:  "worker123",
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Output:    "Some happy output",
	}

	savedJob, err := j.Save(testSaveJob)

	if err != nil {
		t.Fatal("Failed to save test job for Get test", err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Job
		wantErr bool
	}{
		{"Get", args{savedJob.ID}, savedJob, false},
		{"Get Not Found", args{"1234564788478478744"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

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
	j, clean := newTestStore(t)
	defer clean()

	testSaveJob := &models.Job{
		WorkerID:  "worker123",
		StartTime: time.Now(),
		EndTime:   time.Now(),
		Output:    "Some happy output",
	}

	savedJob, err := j.Save(testSaveJob)

	if err != nil {
		t.Fatal("Failed to save test job for Get test", err)
	}

	savedJob.WorkerID = "worker321"
	type args struct {
		id  string
		job *models.Job
	}
	tests := []struct {
		name      string
		args      args
		wantSaved *models.Job
		wantErr   bool
	}{
		{"Update", args{savedJob.ID, savedJob}, savedJob, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := j.Update(tt.args.id, tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := j.Get(savedJob.ID)

			if err != nil {
				t.Errorf("jobStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.WorkerID != savedJob.WorkerID {
				t.Errorf("jobStore.Get() = %v want %v", got, tt.wantSaved)
			}

		})
	}
}
