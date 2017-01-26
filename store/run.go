package store

import "github.com/ferrariframework/ferrariserver/models"

//Run worker run store interface
type Run interface {
	Save(*models.Run) (*models.Run, error)
	Get(id string) (*models.Run, error)
	Update(id string, model *models.Run) error
}
