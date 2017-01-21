package main

import (
	_ "github.com/ferrariframework/ferrariserver/design"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	genapp "github.com/goadesign/goa/goagen/gen_app"
	genclient "github.com/goadesign/goa/goagen/gen_client"
	genjs "github.com/goadesign/goa/goagen/gen_js"
	genmain "github.com/goadesign/goa/goagen/gen_main"
	genschema "github.com/goadesign/goa/goagen/gen_schema"
	genswagger "github.com/goadesign/goa/goagen/gen_swagger"
)

func main() {
	codegen.ParseDSL()
	codegen.Run(
		genmain.NewGenerator(
			genmain.API(design.Design),
			genmain.Target("app"),
		),
		genswagger.NewGenerator(
			genswagger.API(design.Design),
		),
		genapp.NewGenerator(
			genapp.API(design.Design),
			genapp.OutDir("app"),
			genapp.Target("app"),
			genapp.NoTest(true),
		),
		genclient.NewGenerator(
			genclient.API(design.Design),
		),
		genschema.NewGenerator(
			genschema.API(design.Design),
		),
		genjs.NewGenerator(
			genjs.API(design.Design),
		),
	)
}
