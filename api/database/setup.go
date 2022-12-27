package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sherlock28/api-go-gim-gorm/api/models"
)

func ConnectDatabase() *gorm.DB {

	dbURL := "postgres://postgres:postgres@localhost:5432/test"

	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalln(err)
	}

	return database
}
