package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvs() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

}

var DB *gorm.DB

func ConnectDB() {

	var err error
	db_url := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect DB:", err)
	}

}
