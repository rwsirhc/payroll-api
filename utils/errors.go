// utils/errors.go
package utils

import "github.com/gofiber/fiber/v2"

func NewError(code int, message string) error {
	return fiber.NewError(code, message)
}
