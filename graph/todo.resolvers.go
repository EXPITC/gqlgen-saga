package graph

import (
	"context"
	"fmt"

	"github.com/expitc/gqlgen-saga/graph/model"
)

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	var user *model.User

	for _, u := range r.users {
		if u.ID == obj.UserID {
			user = u
		}
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}
