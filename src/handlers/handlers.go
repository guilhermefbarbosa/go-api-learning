package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-api-learning/src/handlers/health"
	"go-api-learning/src/interfaces"
)

type HandlerContainer struct {
	Handlers []interfaces.Handler
	Server   *fiber.App
}

func NewHandlerContainer() interfaces.HandlerContainer {
	return &HandlerContainer{
		Handlers: []interfaces.Handler{
			health.NewHealthHandler(),
		},
	}
}

func (this *HandlerContainer) GetHandlers() []interfaces.Handler {
	return this.Handlers
}
