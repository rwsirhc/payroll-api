package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rwsirhc/gofiber-api/handlers"
)

func RegisterUserRoutes(app *fiber.App) {
	user := app.Group("/users")
	user.Get("/", handlers.GetUsers)
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Put("/:id", handlers.UpdateUser)
	user.Delete("/:id", handlers.DeleteUser)
}
