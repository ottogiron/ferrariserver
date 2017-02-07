package resources

import (
	. "github.com/ferrariframework/ferrariserver/design/mediatypes"
	. "github.com/ferrariframework/ferrariserver/design/usertypes"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var workerRegex = "WRK-[0-9a-v]{20}"

var _ = Resource("worker", func() {
	DefaultMedia(WorkerMedia)
	BasePath("/workers")
	Action("list", func() {
		Routing(
			GET(""),
		)

	})

	Action("show", func() {
		Routing(
			GET("/:workerID"),
		)

		Params(func() {
			Param("workerID", String, "Worker ID", func() {
				Pattern(workerRegex)
			})
		})

		Description("Retrieve a worker given an id")
		Response(OK)
		Response(NotFound)

	})

	Action("create", func() {
		Routing(
			POST(""),
		)

		Description("Create a new worker")
		Payload(WorkerPayload)

		Response(Created, WorkerMedia)

	})

	Action("update", func() {
		Routing(
			PUT("/:workerID"),
		)

		Params(func() {
			Param("workerID", String, "Worker ID", func() {
				Pattern(workerRegex)
			})
		})

		Payload(WorkerPayload)

		Response(NoContent)
		Response(NotFound)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:workerID"),
		)

		Params(func() {
			Param("workerID", String, "Worker ID", func() {
				Pattern(workerRegex)
			})
		})

		Response(NoContent)
		Response(NotFound)
	})

})
