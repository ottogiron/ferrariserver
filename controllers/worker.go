package controllers

import (
	"github.com/ferrariframework/ferrariserver/app"
	"github.com/goadesign/goa"
)

// WorkerController implements the worker resource.
type WorkerController struct {
	*goa.Controller
}

// NewWorkerController creates a worker controller.
func NewWorkerController(service *goa.Service) *WorkerController {
	return &WorkerController{Controller: service.NewController("WorkerController")}
}

// List runs the list action.
func (c *WorkerController) List(ctx *app.ListWorkerContext) error {
	// WorkerController_List: start_implement

	// Put your logic here

	// WorkerController_List: end_implement
	return nil
}
