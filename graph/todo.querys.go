package graph

import (
	"context"
	"fmt"

	"github.com/expitc/gqlgen-saga/graph/model"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, todoID uint) (*model.Todo, error) {
	var todo *model.Todo

	for _, t := range r.todos {
		if t.ID == todoID {
			todo = t
		}
	}

	if todo == nil {
		return nil, fmt.Errorf("todo not found")
	}

	return todo, nil
}

// UserTodos is the resolver for the userTodos field.
func (r *queryResolver) UserTodos(ctx context.Context, userID uint) ([]*model.Todo, error) {
	var todos []*model.Todo

	for _, todo := range r.todos {
		if todo.UserID == userID {
			todos = append(todos, todo)
		}
	}

	if todos == nil {
		return nil, fmt.Errorf("User have no todo")
	}

	return todos, nil
}
