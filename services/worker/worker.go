package worker

import "github.com/ferrariframework/ferrariserver/models"

//Service defines a worker interface
type Service interface {
	Save(models.Worker) (*models.Worker, error)
	Update(models.Worker) error
	Delete(models.Worker) error
}

var _ Service = (*Worker)(nil)

//Worker implementation of a worker
type Worker struct {
}

//Save saves a new worker
func (w *Worker) Save(worker models.Worker) (*models.Worker, error) {
	panic("not implemented")
}

//Update updates a worker
func (w *Worker) Update(worker models.Worker) error {
	panic("not implemented")
}

//Delete deletes a worker
func (w *Worker) Delete(worker models.Worker) error {
	panic("not implemented")
}
