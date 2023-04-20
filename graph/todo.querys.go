package graph

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/expitc/gqlgen-saga/graph/model"
	"github.com/expitc/gqlgen-saga/initializers"
)

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, todoID uint) (*model.Todo, error) {
	var todo *model.Todo

	getTodo := initializers.DB.First(&todo, todoID)

	if getTodo.Error != nil {
		return nil, fmt.Errorf("cannot get todo.//%s", getTodo.Error)
	}

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, input *model.PaginationRequest) ([]*model.Todo, error) {
	var todos []*model.Todo

	// default params if user not provide
	batch := 0
	batchSize := 10

	// update params if user provide
	if input != nil {
		batch = input.Batch
		batchSize = input.BatchSize
	}

	getTodos := initializers.DB.Scopes(paginate(batch, batchSize)).Find(&todos)

	if getTodos.Error != nil {
		return nil, fmt.Errorf("cannot retrive todos.//%s", getTodos.Error)
	}

	return todos, nil
}

// UserTodos is the resolver for the userTodos field.
func (r *queryResolver) UserTodos(ctx context.Context, userID uint) ([]*model.Todo, error) {
	var todos []*model.Todo
	var user = &model.User{
		Model: gorm.Model{
			ID: userID,
		},
	}

	// use db realtion association base on user one to many relation with todo to get all todos that associate with
	association := initializers.DB.Model(&user).Association("Todo")

	// check if associate probably has been change
	if association.Error != nil {
		return nil, fmt.Errorf("association have been change or its not correct.//%s", association.Error)
	}

	// if its fine then find and store to tods
	association.Find(&todos)

	// return the result
	return todos, nil
}
