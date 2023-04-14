package graph

import (
	"context"

	"github.com/expitc/gqlgen-saga/graph/model"
)

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	user := &model.User{
		ID:   obj.UserID,
		Name: "User" + obj.UserID,
	}
	return user, nil
}
