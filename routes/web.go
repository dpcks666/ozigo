package routes

import (
	"net/http"
	"ozigo/app"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func RegisterWeb(e *echo.Echo) {
	// Homepage
	//web.Get("/", Controller.Index(a))

	// Panic test route, this brings up an error
	e.GET("/panic", func(c echo.Context) error {
		panic("Hi, I'm a panic error!")
	})

	// Test to load static, compiled assets
	e.GET("/test", func(c echo.Context) error {
		session, _ := app.Instance().Store.Get(c.Request(), "test")
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		app.Instance().Logger.Error(session.Values)
		session.Values["foo"] = "bar"
		session.Values["time"] = time.Now()
		app.Instance().Logger.Error(session.Values)
		session.Save(c.Request(), c.Response())
		return c.JSON(http.StatusOK, "test")
	})
}
