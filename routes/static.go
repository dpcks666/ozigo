package routes

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func RegisterStatic(e *echo.Echo) {
	e.Static("/static", "./public")
	e.File("/robots.txt", "./public/robots.txt")
}

func StaticSkipper(c echo.Context) bool {
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
