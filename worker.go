package main

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

// Create runs the create action.
func (c *WorkerController) Create(ctx *app.CreateWorkerContext) error {
	// WorkerController_Create: start_implement

	// Put your logic here

	// WorkerController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *WorkerController) Delete(ctx *app.DeleteWorkerContext) error {
	// WorkerController_Delete: start_implement

	// Put your logic here

	// WorkerController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *WorkerController) List(ctx *app.ListWorkerContext) error {
	// WorkerController_List: start_implement

	// Put your logic here

	// WorkerController_List: end_implement
	return nil
}

// Show runs the show action.
func (c *WorkerController) Show(ctx *app.ShowWorkerContext) error {
	// WorkerController_Show: start_implement

	// Put your logic here

	// WorkerController_Show: end_implement
	res := &app.Worker{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *WorkerController) Update(ctx *app.UpdateWorkerContext) error {
	// WorkerController_Update: start_implement

	// Put your logic here

	// WorkerController_Update: end_implement
	return nil
}
