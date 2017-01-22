package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("ferrariserver", func() {
	Title("Manages, runs, and monitors Ferrari Workers")
	Description("Provides an API to manage ferrari workers")
	Contact(func() {
		Name("Otto Giron")
		Email("ottog2486@gmail.com")
		URL("http://ottogiron.me")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/ferrariframework/ferrariserver/blob/master/LICENSE")
	})
	Docs(func() {
		Description("ferrari server guide")
		URL("http://github.com/ferrariframework/ferrariserver")
	})
	Host("localhost:8081")
	Scheme("http")
	BasePath("/v1")

})
