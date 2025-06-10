package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rwsirhc/payroll-api/config"
	middleware "github.com/rwsirhc/payroll-api/middleware"
	routes "github.com/rwsirhc/payroll-api/routes/user"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	config.ConnectDatabase()
	routes.RegisterUserRoutes(app)

	app.Listen(":3000")
}
