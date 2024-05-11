package todosV1

import (
	"github.com/go-chi/chi/v5"
)

func NewTodosRouter() chi.Router {
	router := chi.NewRouter()
	handler := NewTodosHandler()

	router.Get("/", handler.GetTodos())
	router.Post("/", handler.CreateTodo())

	return router
}
