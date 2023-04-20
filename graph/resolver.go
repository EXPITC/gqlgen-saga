package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"gorm.io/gorm"

	"github.com/expitc/gqlgen-saga/graph/model"
)

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

func paginate(batch int, batchSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// default params
		// still put here if user provide params from input but still the value under 0
		if batch <= 0 {
			batch = 1
		}
		if batchSize <= 0 {
			batchSize = 10
		}

		// take example batch = 1 , batchSize = 10

		// offset will be 0 with limit 10 if we scale batch
		// to next which is 2 then,
		// offset wil be 2-1*10 = 10 with limit 10

		// batch 0/1 = offset 0
		// batch 2 offset 10
		// batch 3 offset 20 and so on with same limit 10 items/batch

		offset := (batch - 1) * batchSize

		// return offset with limit for scope
		return db.Offset(offset).Limit(batchSize)
	}
}
