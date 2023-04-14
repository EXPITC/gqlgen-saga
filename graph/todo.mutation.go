package graph

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/expitc/gqlgen-saga/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Intn(100)),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "User" + input.UserID,
		},
		UserID: input.UserID,
	}

	r.todos = append(r.todos, todo)

	return todo, nil
}

// FinishTodo is the resolver for the finishTodo field.
func (r *mutationResolver) FinishTodo(ctx context.Context, input model.DoneTodo) (*model.Todo, error) {
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
