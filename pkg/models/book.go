package models

import (
	"gorm.io/gorm"
	"github.com/HaydnMeyburgh/booking-management-system/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Authro string `json:"authro"`
	Publication string `json:"publications"`
}

func init() {
	config.Connect()
	db.AutoMigrate(&Book{})
}