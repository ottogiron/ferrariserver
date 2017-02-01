package usertypes

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//ResourceLink represents a resource link
var ResourceLink = Type("ResourceLink", func() {
	Attribute("href", String, func() {
		Description("Represents a link href")
		Example("https://api.ferrari.org/v1/workers/CARD-1SV265177X389440GKLJZIYY")
	})

	Attribute("rel", String, func() {
		Description("Represents a link rel")
		Example("self")
	})

	Attribute("method", String, func() {
		Description("Represents the link http Method")
		Example("GET")
	})
})

//ResourceLinks represents an array of resource link
var ResourceLinks = ArrayOf(ResourceLink)
