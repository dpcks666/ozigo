package main

import (
	"ozigo/app"
	"ozigo/routes"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init app
	a := app.Instance()

	// Auto-migrate database models
	err := a.DB.MigrateModels()
	if err != nil {
		panic(err)
	}

	// Register tracer
	tracer, closer, err := a.Config.GetTracerConfig().NewTracer()
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	// Register middlewares
	// Recover
	a.Use(middleware.Recover())
	// Logger
	a.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: routes.StaticSkipper,
	}))
	// Tracer
	a.Use(jaegertracing.TraceWithConfig(jaegertracing.TraceConfig{
		Skipper: routes.StaticSkipper,
		Tracer:  tracer,
	}))
	// Compress
	a.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: routes.StaticSkipper,
	}))
	// Request ID
	a.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Skipper: routes.StaticSkipper,
	}))
	a.Use(session.Middleware(app.Instance().Store))

	// Register routes
	routes.RegisterStatic(a.Echo)
	//routes.RegisterAPI(a.Echo)
	routes.RegisterWeb(a.Echo)

	// Start listening on the specified address
	a.Logger.Fatal(a.Start(":" + a.Config.GetString("APP_PORT")))
}
