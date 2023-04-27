package graph

import (
	"context"
	"fmt"

	"github.com/expitc/gqlgen-saga/graph/model"
	"github.com/expitc/gqlgen-saga/initializers"
)

// Todo is the resolver for the todo field in user .
func (r *userResolver) Todo(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	var user *model.User

	fetchTodo := initializers.DB.Preload("Todos").Find(&user, obj.ID)

	if fetchTodo.Error != nil {
		return nil, fmt.Errorf("fail to get get user todo.///%s", fetchTodo.Error)
	}

	return user.Todos, nil
}
