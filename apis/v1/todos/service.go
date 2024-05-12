package todosV1

import (
	"context"

	dbsetup "github.com/jm2097/gon/internal/db"
	db "github.com/jm2097/gon/internal/db/codegen"
	"github.com/jm2097/gon/tools/pagination"
)

type TodosService interface {
	GetTodos() (*pagination.OffsetPagination[db.Todo], error)
	CreateTodo(todo *db.CreateTodoParams) (*db.Todo, error)
}

type todosServiceImpl struct {
}

func NewTodosService() TodosService {
	return &todosServiceImpl{}
}

func (s *todosServiceImpl) GetTodos() (*pagination.OffsetPagination[db.Todo], error) {
	var paginatedTodos *pagination.OffsetPagination[db.Todo]

	err := dbsetup.NewDatabaseConnection(func(q *db.Queries) error {
		todos, err := q.ListTodos(context.Background(), db.ListTodosParams{
			Limit:  10,
			Offset: 0,
		})
		if err != nil {
			return err
		}

		paginatedTodos = pagination.NewOffsetPagination(todos, len(todos), 1)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return paginatedTodos, nil
}

func (s *todosServiceImpl) CreateTodo(data *db.CreateTodoParams) (*db.Todo, error) {
	var insertedTodo *db.Todo

	err := dbsetup.NewDatabaseConnection(func(q *db.Queries) error {
		todo, err := q.CreateTodo(context.Background(), *data)
		if err != nil {
			return err
		}

		insertedTodo = &todo

		return nil
	})
	if err != nil {
		return nil, err
	}

	return insertedTodo, nil
}
