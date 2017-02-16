package worker

import "github.com/ferrariframework/ferrariserver/models"

//Service defines a worker interface
type Service interface {
	Save(models.Worker) (*models.Worker, error)
	Update(models.Worker) error
	Delete(models.Worker) error
}

var _ Service = (*worker)(nil)

//worker implementation of a worker
type worker struct {
}

//Save saves a new worker
func (w *worker) Save(worker models.Worker) (*models.Worker, error) {
	panic("not implemented")
}

//Update updates a worker
func (w *worker) Update(worker models.Worker) error {
	panic("not implemented")
}

//Delete deletes a worker
func (w *worker) Delete(worker models.Worker) error {
	panic("not implemented")
}
