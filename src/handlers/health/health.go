package health

import (
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (this *HealthHandler) Routes(d fiber.Router) {
	Api := d.Group("/health")
	Api.Get("/", this.healthCheck)
}

func (this *HealthHandler) healthCheck(c *fiber.Ctx) error {
	return c.Status(200).SendString("I'm alive for sure.")
}
