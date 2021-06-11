package app

import (
	"ozigo/config"
	"ozigo/database"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
)

type App struct {
	*echo.Echo
	DB     *database.Database
	Store  sessions.Store
	Config *config.Config
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
		DB:     database.New(config.GetDatabaseDialector()),
		Store:  sessions.NewCookieStore([]byte(config.GetString("APP_SECRET_KEY"))),
		Config: config,
	}

	// app debug settings
	app.Debug = config.GetBool("APP_DEBUG")
	app.HideBanner = !app.Debug
	if app.Debug {
		// middlewares - pprof
		pprof.Register(app.Echo, "/debug/pprof")
	}
}

func Instance() *App {
	return app
}
