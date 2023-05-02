package models

import (
	"gorm.io/gorm"
	"github.com/HaydnMeyburgh/booking-management-system/pkg/config"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"authro"`
	Publication string `json:"publications"`
}

func init() {
	config.DBConnection()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	DB.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	DB.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := DB.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	DB.Where("ID=?", ID).Delete(book)
	return book
}