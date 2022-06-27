package main

import (
	"go-api-learning/src/handlers"
	"go-api-learning/src/http"
	"go-api-learning/src/middlewares"
)

func main() {
	middlewareContainer := middlewares.NewMiddlewareContainer()
	handlersWithServices := handlers.NewHandlerContainer()
	http.NewServer(handlersWithServices, middlewareContainer).Start()
}
