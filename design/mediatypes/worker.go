package mediatypes

import (
	. "github.com/ferrariframework/ferrariserver/design/usertypes"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//Worker Defines a  worker media type
var Worker = MediaType("application/vnd.worker+json", func() {
	Description("A ferrari worker")

	Attributes(func() {
		Attribute("id", String, "ID of the worker", func() {
			Example("9m4e2mr0ui3e8a215n4g")
		})
		Attribute("name", String, "Name of the worker", func() {
			Example("My Awesome Worker")
		})

		Attribute("description", String, "Description of the worker", func() {
			Example("This worker process awesome things very easily")
		})

		Attribute("links", ResourceLinks, "Worker links")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("links")
	})
})

//Workers definition of workers media type
var Workers = MediaType("application/vnd.workers+json", func() {
	Description("A list of workers")
	Attributes(func() {
		Attribute("total_items", Integer, "Total number of items in the collection")
		Attribute("items", CollectionOf(Worker), "Response items")
		Attribute("total_pages", Integer, "Total number of pages in the collection")
		Attribute("links", ResourceLinks, "List of workers links")
	})

	View("default", func() {
		Attribute("total_items")
		Attribute("items")
		Attribute("total_pages")
		Attribute("links")
	})
})
