package app

import (
	"ozigo/config"
	"ozigo/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

type App struct {
	Server *fiber.App
	DB     *database.Database
	Store  *session.Store
	Logger *zap.Logger
	Bundle *i18n.Bundle
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
		Bundle: i18n.NewBundle(language.English), // default set
		Config: config,
	}
}

func Instance() *App {
	return app
}
