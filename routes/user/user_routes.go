package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/rwsirhc/payroll-api/handlers/user"
)

func RegisterUserRoutes(app *fiber.App) {
	user := app.Group("/users")
	user.Get("/", handlers.GetUsers)
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Put("/:id", handlers.UpdateUser)
	user.Delete("/:id", handlers.DeleteUser)
}
