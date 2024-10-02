package todosV1

import (
	"github.com/labstack/echo/v4"
)

func NewGroup(g *echo.Group) {
	handler := NewTodosHandler()

	group := g.Group("/todos")
	group.GET("", handler.GetTodos)
}
