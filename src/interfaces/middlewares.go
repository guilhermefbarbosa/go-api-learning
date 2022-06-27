package middlewares

import "github.com/olaisaac/bigweld/src/interfaces"

type MiddlewareContainer struct{}

func NewMiddlewareContainer() interfaces.MiddlewareContainer {
	return &MiddlewareContainer{}
}
