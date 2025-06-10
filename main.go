package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rwsirhc/gofiber-api/config"
	"github.com/rwsirhc/gofiber-api/routes"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()
	routes.RegisterUserRoutes(app)

	app.Listen(":3000")
}
