package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewTodosRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", getTodos)
	router.Delete("/{id}", deleteTodo)

	return router
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here are your todos"))
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Todo deleted"))
}
