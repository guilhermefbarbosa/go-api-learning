package interfaces

import "github.com/gofiber/fiber/v2"

type MiddlewareContainer interface {
	TokenAuth(c *fiber.Ctx) error
}
