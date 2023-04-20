package graph

import (
	"context"
	"fmt"

	"github.com/expitc/gqlgen-saga/graph/model"
	"github.com/expitc/gqlgen-saga/initializers"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {

	todo := &model.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	}

	// insert new todo and response
	create := initializers.DB.Create(todo)

	if create.Error != nil {
		// means no user associate create with omit and user resolver will handel the user
		create = initializers.DB.Omit("UserID").Create(todo)
		if create.Error != nil {
			return nil, fmt.Errorf("can't create todo.//%s", create.Error)
		}
	}

	return todo, nil
}

// FinishTodo is the resolver for the finishTodo field.
func (r *mutationResolver) MarkTodo(ctx context.Context, input model.MarkTodo) (*model.Todo, error) {
	var todo *model.Todo

	updateTodo := initializers.DB.Model(&todo).Where("id = ?", input.ID).Update("Done", input.Done)

	if updateTodo.Error != nil {
		return nil, fmt.Errorf("todo not found.//%s", updateTodo.Error)
	}

	getTodo := initializers.DB.First(&todo, input.ID)

	if getTodo.Error != nil {
		return nil, fmt.Errorf("cannot retrive todo.//%s", getTodo.Error)
	}

	return todo, nil
}
