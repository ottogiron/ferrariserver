package worker

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
	oelastic "gopkg.in/olivere/elastic.v3"
)

func newTestStore(t *testing.T) (store.Worker, func()) {
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

	_, err = client.CreateIndex(index).
		BodyJson(map[string]interface{}{
			"settings": map[string]interface{}{
				"number_of_shards": 1,
			},
		}).
		Do(context.Background())

	if err != nil {
		t.Fatal("Failed to create index ", err)
	}

	_, err = client.PutMapping().
		Index(index).
		Type(docType).
		BodyJson(map[string]interface{}{
			"properties": map[string]interface{}{
				"environment": map[string]interface{}{
					"type": "nested",
				},
			},
		}).
		Do(context.Background())

	if err != nil {
		t.Fatal("Failed to create document mappings ", err)
	}

	s := New(
		SetClient(client),
		SetIndex(index),
		SetDocType(docType),
		SetRefreshIndex("true"),
	)
	return s, func() {
		client.DeleteIndex(index).
			Do(context.Background())
	}
}

func Test_workerStore_Save(t *testing.T) {
	w, clean := newTestStore(t)
	defer clean()
	type args struct {
		worker *models.Worker
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Worker
		wantErr bool
	}{
		{"Save", args{&models.Worker{}}, &models.Worker{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := w.Save(tt.args.worker)
			if (err != nil) != tt.wantErr {
				t.Errorf("workerStore.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID == "" {
				t.Errorf("workerStore.Save() = got.ID = %v, want an id value", got.ID)
			}

			if !reflect.DeepEqual(got.Environment, tt.want.Environment) {
				t.Errorf("workerStore.Save() = got.Output = %v, want %v", got.Environment, tt.want.Environment)
			}

		})
	}
}

func Test_workerStore_Get(t *testing.T) {

	w, clean := newTestStore(t)
	defer clean()

	testSaveWorker := &models.Worker{
		Environment: map[string]string{"SOME_VAR": "some_value"},
	}

	savedWorker, err := w.Save(testSaveWorker)

	if err != nil {
		t.Fatal("Failed to save test worker for Get test", err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Worker
		wantErr bool
	}{
		{"Get", args{savedWorker.ID}, savedWorker, false},
		{"Get Not Found", args{"1234564788478478744"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := w.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("workerStore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got.ID != tt.want.ID {
				t.Errorf("workerStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_workerStore_Update(t *testing.T) {
	w, clean := newTestStore(t)
	defer clean()

	testSavedWorker := &models.Worker{
		Environment: map[string]string{"SOME_VAR": "some_value"},
	}

	savedWorker, err := w.Save(testSavedWorker)

	if err != nil {
		t.Fatal("Failed to save test worker for Get test", err)
	}

	savedWorker.Environment = map[string]string{"SOME_VAR": "some_new_value"}
	type args struct {
		id     string
		worker *models.Worker
	}
	tests := []struct {
		name      string
		args      args
		wantSaved *models.Worker
		wantErr   bool
	}{
		{"Update", args{savedWorker.ID, savedWorker}, savedWorker, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := w.Update(tt.args.id, tt.args.worker)
			if (err != nil) != tt.wantErr {
				t.Errorf("workerStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := w.Get(savedWorker.ID)

			if err != nil {
				t.Errorf("workerStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Environment["SOME_VAR"] != savedWorker.Environment["SOME_VAR"] {
				t.Errorf("workerStore.Get() = %v want %v", got, tt.wantSaved)
			}

		})
	}
}
