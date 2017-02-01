package resources

import (
	. "github.com/ferrariframework/ferrariserver/design/mediatypes"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("worker", func() {
	DefaultMedia(Worker)
	BasePath("/workers")
	Action("list", func() {
		Routing(
			GET(""),
		)

	})
})
