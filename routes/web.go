package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterWeb(app *fiber.App) {
	// Homepage
	//web.Get("/", Controller.Index(a))

	// Panic test route, this brings up an error
	app.Get("/panic", func(ctx *fiber.Ctx) error {
		panic("Hi, I'm a panic error!")
	})

	// Test to load static, compiled assets
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Render("test", fiber.Map{
			"Title": "ko",
			"Body":  "Body!!!",
		})
	})
}
