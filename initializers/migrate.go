package initializers

import (
	"log"

	"github.com/expitc/gqlgen-saga/graph/model"
)

func SyncDatabase() {
	// defer sqlDB.Close()

	log.Print("Migrate models")

	if err := DB.AutoMigrate(&model.User{}, &model.Todo{}); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Model Migrated")
	}

	log.Print("Migrate Success")
}
