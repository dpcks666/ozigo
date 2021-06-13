package app

import (
	"ozigo/config"
	"ozigo/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.uber.org/zap"
)

type App struct {
	Server *fiber.App
	DB     *database.Database
	Store  *session.Store
	Logger *zap.Logger
	Config *config.Config
}

var app *App

func init() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	app = &App{
		Server: fiber.New(config.GetFiberConfig()),
		DB:     database.New(config.GetDatabaseDialector()),
		Store:  session.New(config.GetSessionConfig()),
		Logger: zap.New(config.GetLoggerConfig()),
		Config: config,
	}
}

func Instance() *App {
	return app
}
