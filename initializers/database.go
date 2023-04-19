package initializers

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var sqlDB *sql.DB

func ConnectToDB() {
	var err error
	log.Print("Attempt to connect to DB")

	// get dns from env var
	dns := os.Getenv("DB_CONNECTION")

	if dns == "" {
		log.Fatal("environment DB_CONNECTION not specify or fail to get variable")
		return
	}

	// connect to db
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database")
		return
	}

	// set connection time with db from sql dialect
	sqlDB, err = DB.DB()

	if err != nil {
		log.Fatal("failed to set connection max time")
		return
	}

	defer sqlDB.SetConnMaxLifetime(time.Hour)

	// success log
	log.Print("DB Connected")
}
