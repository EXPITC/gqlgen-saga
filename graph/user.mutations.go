package graph

import (
	"context"
	"fmt"

	"github.com/expitc/gqlgen-saga/graph/model"
	"github.com/expitc/gqlgen-saga/initializers"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	user := &model.User{
		Name: input.Name,
	}

	// custom id if provide
	if input.ID != nil {
		// fist check if id available
		var foundUser *model.User
		getUser := initializers.DB.First(&foundUser, input.ID)

		if getUser.Error == nil {
			// return error if id already takes or we can automatic generated the next id
			// but it think is good for end user know that id has been taken.
			return nil, fmt.Errorf("ID has been taken.///record already exist.///%s", foundUser.Name)
		}

		// set the custom id, id available.
		user.ID = *input.ID
	}

	createUser := initializers.DB.Create(user)

	if createUser.Error != nil {
		return nil, fmt.Errorf("fail to create new user.///%s", createUser.Error)
	}

	return user, nil
}
