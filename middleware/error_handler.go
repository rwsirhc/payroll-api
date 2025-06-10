package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler satisfies fiber.ErrorHandler (correct signature)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Default error code and message
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// If it's a Fiber error, use its code and message
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// Log the error (optional)
	log.Printf("[ERROR] %v", err)

	// Send JSON response
	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": message,
	})
}
