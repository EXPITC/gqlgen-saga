package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	log.Print("Attemp to connect to DB")

	// get dns from env var
	dns := os.Getenv("DB_CONNECTION")

	if dns == "" {
		log.Fatal("environment DB_CONNECTION not specify or fail to get variable")
		return
	}

	// connect to db
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
		return
	}

	// success log
	log.Print("DB Connected")
}
