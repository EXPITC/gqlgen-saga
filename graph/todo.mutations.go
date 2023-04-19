package graph

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/expitc/gqlgen-saga/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {

	todo := &model.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}

	var user *model.User

	// create new user if not yet register
	for _, u := range r.users {
		if u.ID == input.UserID {
			user = nil
			break
		} else {
			user = &model.User{
				Model: gorm.Model{},
				Name:  fmt.Sprintf("User %d", input.UserID),
			}
		}
	}

	if len(r.users) == 0 {
		user = &model.User{
			Name: fmt.Sprintf("User %d", input.UserID),
		}
	}

	if user != nil {
		r.users = append(r.users, user)
	}

	// insert new todo and response
	r.todos = append(r.todos, todo)

	return todo, nil
}

// FinishTodo is the resolver for the finishTodo field.
func (r *mutationResolver) MarkTodo(ctx context.Context, input model.MarkTodo) (*model.Todo, error) {
	var finalResult *model.Todo

	for i, todo := range r.todos {
		if todo.ID == input.ID {
			r.todos[i].Done = input.Done
			finalResult = r.todos[i]
		}
	}

	if finalResult == nil {
		return nil, fmt.Errorf("todo not found")
	}

	return finalResult, nil
}
