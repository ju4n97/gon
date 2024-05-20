package todosV1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodosHandler struct {
	todosService TodosService
}

func NewTodosHandler() *TodosHandler {
	return &TodosHandler{todosService: NewTodosService()}
}

func (s *TodosHandler) GetTodos(c echo.Context) error {
	todos, err := s.todosService.GetTodos()
	if err != nil {
		c.Error(err)
	}

	return c.JSON(http.StatusOK, todos)
}
