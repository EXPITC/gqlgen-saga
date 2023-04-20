package graph

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/expitc/gqlgen-saga/graph/model"
	"github.com/expitc/gqlgen-saga/initializers"
)

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	var user *model.User

	// get user
	result := initializers.DB.First(&user, obj.UserID)

	// if user not found create user & update todos
	if result.Error != nil {
		// user not found create model user
		user = &model.User{
			Model: gorm.Model{
				ID: obj.UserID,
			},
			Name: fmt.Sprintf("User %d", obj.UserID),
		}

		// insert new model to db
		result := initializers.DB.Create(user)

		// error handel if fail to init new user or fail to insert new data to user
		if result.Error != nil {
			errMsg := errors.New("user not found & cannot create user")
			return nil, errMsg
		}

		updateUserTodo := initializers.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&obj).Updates(model.Todo{UserID: user.ID})

		if updateUserTodo.Error != nil {
			errMsg := errors.New("fail to create todo")
			initializers.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&user, &obj)
			return nil, errMsg
		}

		// success created user & update todo, return user
		return user, nil
	}

	// user found return user
	return user, nil
}
