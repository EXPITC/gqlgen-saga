package graph

import (
	"context"
	"fmt"

	"github.com/expitc/gqlgen-saga/graph/model"
	"github.com/expitc/gqlgen-saga/initializers"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, input *model.PaginationRequest) ([]*model.User, error) {
	var users []*model.User
	batch := 0
	batchSize := 10

	if input != nil {
		batch = input.Batch
		batchSize = input.BatchSize
	}

	getUsers := initializers.DB.Scopes(paginate(batch, batchSize)).Find(&users)

	if getUsers.Error != nil {
		return nil, fmt.Errorf("cannot retrive users.//%s", getUsers.Error)
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID uint) (*model.User, error) {
	var user *model.User

	getUser := initializers.DB.Find(&user, userID)

	if getUser.Error != nil {
		return nil, fmt.Errorf("cannot retrive user.//%s", getUser.Error)
	}

	return user, nil
}
