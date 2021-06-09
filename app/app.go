package app

import (
	"ozigo/config"
	"ozigo/database"

	"github.com/gorilla/sessions"
	tracer "github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
)

type App struct {
	*echo.Echo
	Config *config.Config
	DB     *database.Database
	Store  sessions.Store
	Tracer *opentracing.Tracer
}

var app *App

func init() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	// app struct init
	app = &App{
		Echo:   echo.New(),
		Config: config,
		DB:     database.New(config.GetDatabaseDialector()),
		Store:  sessions.NewCookieStore([]byte(config.GetString("SESSION_KEY"))),
	}

	// app settings
	app.Debug = config.GetBool("APP_DEBUG")
	app.HideBanner = true
}

func Instance() *App {
	return app
}

func (a *App) RegisterMiddlewares(skipper func(c echo.Context) bool) {
	// Debug utils - Pprof
	if a.Debug {
		pprof.Register(a.Echo, "/debug/pprof")
	}

	// Recover
	a.Use(middleware.Recover())

	// Logger
	a.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: skipper,
	}))

	// Tracer
	a.Use(tracer.TraceWithConfig(tracer.TraceConfig{
		Skipper: skipper,
		Tracer:  *a.Tracer,
	}))

	// Compress
	a.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: skipper,
	}))

	// Session
	a.Use(session.MiddlewareWithConfig(session.Config{
		Skipper: skipper,
		Store:   a.Store,
	}))

	// Request ID
	a.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Skipper: skipper,
	}))
}
