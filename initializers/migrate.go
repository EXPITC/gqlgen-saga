package initializers

import (
	"log"

	"github.com/expitc/gqlgen-saga/graph/model"
)

func SyncDatabase() {
	log.Print("Migrate models")

	if err := DB.AutoMigrate(&model.Todo{}, &model.User{}); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Model Migrated")
	}

	log.Print("Migrate Success")
}
