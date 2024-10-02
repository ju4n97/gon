package v1

import (
	todosV1 "github.com/ju4n97/gon/apis/v1/todos"
	"github.com/labstack/echo/v4"
)

func NewGroup(g *echo.Group) {
	v1 := g.Group("/v1")

	todosV1.NewGroup(v1)
}
