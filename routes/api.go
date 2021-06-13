package routes

import (
	Controller "ozigo/app/controllers/api"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app *fiber.App) {
	api := app.Group("/api")
	registerRoles(api)
	//registerUsers()
}

func registerRoles(api fiber.Router) {
	roles := api.Group("/roles")
	roles.Get("/", Controller.GetAllRoles)
	roles.Get("/:id", Controller.GetRole)
	roles.Post("/", Controller.AddRole)
	roles.Put("/:id", Controller.EditRole)
	roles.Delete("/:id", Controller.DeleteRole)
}

//func registerUsers() {
//	users := app.Server.Group("/users")
//
//	users.Get("/", Controller.GetAllUsers())
//	users.Get("/:id", Controller.GetUser())
//	users.Post("/", Controller.AddUser())
//	users.Put("/:id", Controller.EditUser())
//	users.Delete("/:id", Controller.DeleteUser())
//}
