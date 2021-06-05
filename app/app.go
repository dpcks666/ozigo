package app

import (
	"ozigo/config"
	"ozigo/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/opentracing/opentracing-go"
	tracer "github.com/shareed2k/fiber_tracing"
)

type App struct {
	Server  *fiber.App
	Config  *config.Config
	DB      *database.Database
	Session *session.Store
	Tracer  *opentracing.Tracer
}

var app *App

func init() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	app = &App{
		Server:  fiber.New(config.GetFiberConfig()),
		Config:  config,
		DB:      database.New(config.GetDatabaseDialector()),
		Session: session.New(config.GetSessionConfig()),
	}
}

func Instance() *App {
	return app
}

func (a *App) RegisterMiddlewares(skipper func(c *fiber.Ctx) bool) {
	// Debug utils - Pprof, Monitor
	if a.Config.GetBool("APP_DEBUG") {
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
		Next: skipper,
	}))

	// Tracer
	a.Server.Use(tracer.New(tracer.Config{
		Tracer: *a.Tracer,
		Filter: skipper,
	}))

	// Compress
	a.Server.Use(compress.New(compress.Config{
		Next: skipper,
	}))

	// Etag
	a.Server.Use(etag.New(etag.Config{
		Next: skipper,
	}))
}
