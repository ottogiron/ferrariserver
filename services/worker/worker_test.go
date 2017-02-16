package worker

import (
	"reflect"
	"testing"

	"github.com/ferrariframework/ferrariserver/models"
)

func TestWorker_Save(t *testing.T) {
	type args struct {
		worker models.Worker
	}
	tests := []struct {
		name    string
		w       *worker
		args    args
		want    *models.Worker
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &worker{}
			got, err := w.Save(tt.args.worker)
			if (err != nil) != tt.wantErr {
				t.Errorf("Worker.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Worker.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorker_Update(t *testing.T) {
	type args struct {
		worker models.Worker
	}
	tests := []struct {
		name    string
		w       *worker
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &worker{}
			if err := w.Update(tt.args.worker); (err != nil) != tt.wantErr {
				t.Errorf("Worker.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWorker_Delete(t *testing.T) {
	type args struct {
		worker models.Worker
	}
	tests := []struct {
		name    string
		w       *worker
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &worker{}
			if err := w.Delete(tt.args.worker); (err != nil) != tt.wantErr {
				t.Errorf("Worker.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
