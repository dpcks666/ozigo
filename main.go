package main

import (
	"log"

	"ozigo/app"
	"ozigo/routes"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/opentracing/opentracing-go"
	fiberTracer "github.com/shareed2k/fiber_tracing"
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
		defer closer.Close()
	}

	// Register middlewares
	if a.Config.GetBool("APP_DEBUG") {
		// Debug utils - Pprof, Monitor
		debug := a.Server.Group("/debug")
		debug.Use("/pprof/*", pprof.New())
		debug.Use("/monitor", monitor.New())
	}
	// Recover
	a.Server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	// Logger
	a.Server.Use(logger.New(logger.Config{
		Next: routes.SkipperStatic,
	}))
	// Tracer
	a.Server.Use(fiberTracer.New(fiberTracer.Config{
		Tracer: tracer,
		Filter: routes.SkipperStatic,
	}))
	// Compress
	a.Server.Use(compress.New(compress.Config{
		Next: routes.SkipperStatic,
	}))
	// Etag
	a.Server.Use(etag.New(etag.Config{
		Next: routes.SkipperStatic,
	}))

	// Register routes
	routes.RegisterStatic(a.Server)
	routes.RegisterAPI(a.Server)
	routes.RegisterWeb(a.Server)

	// Start listening on the specified address
	log.Fatal(a.Server.Listen(":" + a.Config.GetString("APP_PORT")))
}
