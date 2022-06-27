package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-api-learning/src/interfaces"
	"os"
	"os/signal"
	"syscall"
)

type HTTPServer struct {
	Server      *fiber.App
	Middlewares interfaces.MiddlewareContainer
}

func NewServer(handlers interfaces.HandlerContainer, mid interfaces.MiddlewareContainer) *HTTPServer {
	fiberServer := fiber.New()

	server := &HTTPServer{
		Server: fiberServer,
	}
	server.loadRoutes(handlers, mid)
	return server
}

func (hs *HTTPServer) loadRoutes(handlerContainer interfaces.HandlerContainer, middlewares interfaces.MiddlewareContainer) {
	
	Router := hs.Server.Use(middlewares.TokenAuth)
	for _, handler := range handlerContainer.GetHandlers() {
		handler.Routes(Router)
	}
}

func (hs *HTTPServer) Start() {
	server := hs.Server

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-quit
		fmt.Printf("Gracefully shutting down...")
		_ = server.Shutdown()
	}()

	if err := server.Listen(":8080"); err != nil {
		fmt.Printf("Error shutting down server: %s", err.Error())
	}

	fmt.Printf("Running cleanup tasks...")
}
