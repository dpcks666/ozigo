package main

import (
	"log"

	"ozigo/app"
	"ozigo/routes"

	"github.com/opentracing/opentracing-go"
)

func main() {
	// Init app
	a := app.Instance()

	// Auto-migrate database models
	if a.DB != nil {
		err := a.DB.MigrateModels()
		if err != nil {
			panic(err)
		}
	}

	// Register tracer
	tracer, closer, err := a.Config.GetTracerConfig().NewTracer()
	if err != nil {
		log.Println("failed to load tracer:", err.Error())
	} else {
		opentracing.SetGlobalTracer(tracer)
		a.Tracer = &tracer
		defer closer.Close()
	}

	// Register middlewares
	a.RegisterMiddlewares(routes.StaticSkipper)

	// Register routes
	routes.RegisterStatic(a.Echo)
	//routes.RegisterAPI(a.Echo)
	routes.RegisterWeb(a.Echo)

	// Start listening on the specified address
	a.Logger.Fatal(a.Start(":" + a.Config.GetString("APP_PORT")))
}
