package todosV1

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"
	db "github.com/jm2097/gon/internal/db/codegen"
	"github.com/jm2097/gon/tools/httperrors"
	"github.com/jm2097/gon/tools/validators"
)

type TodosHandler struct {
	todosService TodosService
}

func NewTodosHandler() *TodosHandler {
	return &TodosHandler{
		todosService: NewTodosService(),
	}
}

// TodoRequest defines the request payload for the Todo data model.
type todosRequestPayload struct {
	*db.CreateTodoParams
}

// TodoResponse defines the response payload for the Todo data model.
type todosResponsePayload struct {
	*db.Todo
}

func (p *todosRequestPayload) Bind(r *http.Request) error {
	if err := validators.ValidateModel(p.CreateTodoParams); err != nil {
		return err.Error
	}

	p.Title = strings.ToTitle(strings.TrimSpace(p.Title))

	return nil
}

func (p *todosResponsePayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newTodoPayloadResponse(todo *db.Todo) *todosResponsePayload {
	return &todosResponsePayload{todo}
}

func (s *TodosHandler) GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := s.todosService.GetTodos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, todos)
	}
}

func (s *TodosHandler) CreateTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &todosRequestPayload{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, httperrors.NewBadRequestError(err))
			return
		}

		newTodo, err := s.todosService.CreateTodo(data.CreateTodoParams)
		if err != nil {
			render.Render(w, r, httperrors.NewInternalServerError(err))
			return
		}

		render.Status(r, http.StatusCreated)

		if err := render.Render(w, r, newTodoPayloadResponse(newTodo)); err != nil {
			render.Render(w, r, httperrors.NewInternalServerError(err))
			return
		}
	}
}
