package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-api-learning/internal/config"
	"go-api-learning/internal/config/listutils"
	"net/http"
)

func (mc *MiddlewareContainer) TokenAuth(c *fiber.Ctx) error {
	cfg := config.GetConfig()
	req := c.Request()
	authToken := string(req.Header.Peek("Authorization"))
	fmt.Println(listutils.Strcontains(cfg.Tokens, authToken))
	if !listutils.Strcontains(cfg.Tokens, authToken) {
		return c.Status(http.StatusUnauthorized).SendString("GO AWAY.")
	}
	return c.Next()
}
