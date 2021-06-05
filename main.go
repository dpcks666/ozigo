package main

import (
	"log"

	"ozigo/app"
	"ozigo/routes"
)

func main() {
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
		a.Tracer = &tracer
		defer closer.Close()
	}

	// Register middlewares
	a.RegisterMiddlewares(routes.SkipperStatic)

	// Register routes
	routes.RegisterStatic(a.Server)
	routes.RegisterAPI(a.Server)
	routes.RegisterWeb(a.Server)

	// Start listening on the specified address
	log.Fatal(a.Server.Listen(":" + a.Config.GetString("APP_PORT")))
}
