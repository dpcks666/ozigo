package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RegisterStatic(app *fiber.App) {
	config := fiber.Static{
		Compress:  true,
		ByteRange: true,
	}

	app.Static("/static", "./public", config)
	app.Static("/robots.txt", "./public/robots.txt", config)
}

func SkipperStatic(c *fiber.Ctx) bool {
	paths := []string{
		"/static",
		"/robots.txt",
	}
	for _, path := range paths {
		if strings.HasPrefix(c.Path(), path) {
			return true
		}
	}
	return false
}
