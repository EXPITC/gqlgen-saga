package graph

//go:generate go run github.com/99designs/gqlgen generate

import "github.com/expitc/gqlgen-saga/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// maybe if this in db we can create funct or var that will be called when we need todos tabel collections
	todos []*model.Todo
	todo  *model.Todo
	users []*model.User
	user  *model.User
}
