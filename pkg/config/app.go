package config

import (
	"fmt"
	"log"

	"github.com/HaydnMeyburgh/booking-management-system/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	myDB_dsn := "host=localhost user=haydn password=tryguess dbname=bookstore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(myDB_dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database...", err)
	}

	fmt.Println("Database Connection successful!")
	DB = db
	DB.AutoMigrate(&models.Book{})
}
