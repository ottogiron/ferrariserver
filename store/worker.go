package store

import "github.com/ferrariframework/ferrariserver/models"

//Worker worker store interface
type Worker interface {
	Save(*models.Worker) (*models.Worker, error)
	Get(id string) (*models.Worker, error)
	Update(id string, model *models.Worker) error
}
