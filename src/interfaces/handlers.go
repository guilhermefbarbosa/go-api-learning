package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type HandlerContainer interface {
	GetHandlers() []Handler
}

type Handler interface {
	Routes(fiber.Router)
}
