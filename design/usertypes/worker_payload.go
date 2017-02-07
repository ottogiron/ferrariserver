package usertypes

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//WorkerPayload represents a worker payload
var WorkerPayload = Type("WorkerPayload", func() {
	Attribute("name", String, "Worker name", func() {
		MinLength(2)
	})

	Attribute("description", String, "Worker description")
	Attribute("environment", HashOf(String, Any), "Worker environment variables")
})
