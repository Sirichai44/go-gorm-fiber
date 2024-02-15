package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string	`json:"author"`
	Description string	`json:"description"`
	Price       uint	`json:"price"`
}

func createBook(db *gorm.DB,book *Book)  error{
	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Book Created Successfully")
	return nil
}

func getBook(db *gorm.DB,id int) *Book {
	book := Book{}
	result := db.First(&book,id)

	if result.Error != nil {
		log.Fatalf("Error fetching book : %v",result.Error)
	}

	return &book
}

func updateBook(db *gorm.DB,book *Book)  error{
	result := db.Model(&book).Updates(book)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Book Updated Successfully")
	return nil
}


//soft delete
func deleteBook(db *gorm.DB,id int) error {
	result := db.Delete(&Book{},id)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Book Deleted Successfully")
	return nil
}

//hard delete
func deleteBookPermanently(db *gorm.DB,id int)  {
	result := db.Unscoped().Delete(&Book{},id)

	if result.Error != nil {
		log.Fatalf("Error deleting book : %v",result.Error)
	}

	fmt.Println("Book Deleted Successfully")
}

func getBooks(db *gorm.DB,) ([]Book,error) {
	books := []Book{}
	result := db.Find(&books)

	if result.Error != nil {
		log.Fatalf("Error fetching books : %v",result.Error)
	}
	//if no books found
	if len(books) == 0 {
		return nil,	fmt.Errorf("no books found")
	}

	return books,nil
}