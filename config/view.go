package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/ace"
	"github.com/gofiber/template/amber"
	"github.com/gofiber/template/django"
	"github.com/gofiber/template/handlebars"
	"github.com/gofiber/template/html"
	"github.com/gofiber/template/jet"
	"github.com/gofiber/template/mustache"
	"github.com/gofiber/template/pug"
)

func (config *Config) getFiberViewsEngine() (viewsEngine fiber.Views) {
	config.SetDefault("FIBER_VIEWS_DIRECTORY", "resources/views")

	switch config.GetString("FIBER_VIEWS") {
	case "ace":
		engine := ace.New(config.GetString("FIBER_VIEWS_DIRECTORY"), ".ace")
		engine.Reload(config.GetBool("APP_DEBUG"))
		viewsEngine = engine
	case "amber":
		engine := amber.New(config.GetString("FIBER_VIEWS_DIRECTORY"), ".amber")
		engine.Reload(config.GetBool("APP_DEBUG"))
		viewsEngine = engine
	case "django":
		engine := django.New(config.GetString("FIBER_VIEWS_DIRECTORY"), ".django")
		engine.Reload(config.GetBool("APP_DEBUG"))
		viewsEngine = engine
	case "handlebars":
		engine := handlebars.New(config.GetString("FIBER_VIEWS_DIRECTORY"), ".hbs")
		engine.Reload(config.GetBool("APP_DEBUG"))
		viewsEngine = engine
	case "jet":
		engine := jet.New(config.GetString("FIBER_VIEWS_DIRECTORY"), ".jet")
		engine.Reload(config.GetBool("APP_DEBUG"))
		viewsEngine = engine
	case "mustache":
		engine := mustache.New(config.GetString("FIBER_VIEWS_DIRECTORY"), ".mustache")
		engine.Reload(config.GetBool("APP_DEBUG"))
		viewsEngine = engine
	case "pug":
		engine := pug.New(config.GetString("FIBER_VIEWS_DIRECTORY"), ".pug")
		engine.Reload(config.GetBool("APP_DEBUG"))
		viewsEngine = engine
	default:
		engine := html.New(config.GetString("FIBER_VIEWS_DIRECTORY"), ".html")
		engine.Reload(config.GetBool("APP_DEBUG"))
		viewsEngine = engine

	}
	return
}
