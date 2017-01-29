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

func newTestStore(t *testing.T) (store.Run, func()) {
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
	_, err = client.PutMapping().
		Index(index).
		Type(docType).
		BodyJson(map[string]interface{}{
			docType: map[string]interface{}{
				"properties": map[string]interface{}{
					"environment": map[string]interface{}{
						"type": "nested",
					},
				},
			},
		}).
		Do(context.Background())

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

func Test_runStore_Save(t *testing.T) {
	r, clean := newTestStore(t)
	defer clean()
	type args struct {
		run *models.Run
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Run
		wantErr bool
	}{
		{"Save", args{&models.Run{WorkerID: "worker123"}}, &models.Run{WorkerID: "worker123"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := r.Save(tt.args.run)
			if (err != nil) != tt.wantErr {
				t.Errorf("runStore.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID == "" {
				t.Errorf("runStore.Save() = got.ID = %v, want an id value", got.ID)
			}

			if !reflect.DeepEqual(got.WorkerID, tt.want.WorkerID) {
				t.Errorf("runStore.Save() = got.Output = %v, want %v", got.WorkerID, tt.want.WorkerID)
			}

		})
	}
}

func Test_runStore_Get(t *testing.T) {

	r, clean := newTestStore(t)
	defer clean()

	testSavedRun := &models.Run{
		WorkerID: "WorkerID123",
	}

	savedRun, err := r.Save(testSavedRun)

	if err != nil {
		t.Fatal("Failed to save test run for Get test", err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Run
		wantErr bool
	}{
		{"Get", args{savedRun.ID}, savedRun, false},
		{"Get Not Found", args{"1234564788478478744"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("runStore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got.ID != tt.want.ID {
				t.Errorf("runStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runStore_Update(t *testing.T) {
	r, clean := newTestStore(t)
	defer clean()

	testSavedRun := &models.Run{
		WorkerID: "worker123",
	}

	savedRun, err := r.Save(testSavedRun)

	if err != nil {
		t.Fatal("Failed to save test worker for Get test", err)
	}

	savedRun.WorkerID = "worker321"
	type args struct {
		id  string
		run *models.Run
	}
	tests := []struct {
		name      string
		args      args
		wantSaved *models.Run
		wantErr   bool
	}{
		{"Update", args{savedRun.ID, savedRun}, savedRun, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := r.Update(tt.args.id, tt.args.run)
			if (err != nil) != tt.wantErr {
				t.Errorf("workerStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := r.Get(savedRun.ID)

			if err != nil {
				t.Errorf("workerStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.WorkerID != savedRun.WorkerID {
				t.Errorf("workerStore.Get() = %v want %v", got, tt.wantSaved)
			}

		})
	}
}
