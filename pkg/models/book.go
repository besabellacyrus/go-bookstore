package models

import (
	"cyrusbesabella/go-bookstore/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) (Book, error) {
	var book Book
	// First, find the book by ID
	if err := db.Where("ID = ?", Id).First(&book).Error; err != nil {
		// Return an error if the book is not found
		return Book{}, err
	}

	// Delete the book after it's found
	if err := db.Delete(&book).Error; err != nil {
		return Book{}, err
	}

	return book, nil
}