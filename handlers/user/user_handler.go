package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/gofiber-api/models"
	"github.com/yourusername/gofiber-api/service"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := service.FetchAllUsers()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := service.FetchUserByID(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if user == nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := service.CreateUser(&u); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(u)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	success, err := service.ModifyUser(id, &u)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if !success {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(u)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	success, err := service.RemoveUser(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if !success {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.SendStatus(204)
}
