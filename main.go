package main

import (
	"bookstore/controller"
	"bookstore/model"
	"fmt"
	"time"
)

func main() {
	waktu := time.Now()
	newBook := []model.BookStore{
		{
			ID:         1,
			Titles:     "Breaking The Habit Of Being Yourself",
			Authors:    "DR. Joe Dispenza",
			Quantities: 1,
			Available:  true,
			CreatedAt:  waktu.Local().String(),
			UpdatedAt:  waktu.Local().String(),
		},
		{
			ID:         2,
			Titles:     "The Attitude Book",
			Authors:    "Simon Tyler",
			Quantities: 1,
			Available:  true,
			CreatedAt:  waktu.Local().String(),
			UpdatedAt:  waktu.Local().String(),
		},
	}

	bookCtrl := controller.BookStoreDB{}

	insertBook := bookCtrl.Insert(&newBook, bookCtrl)
	if insertBook != nil {
		fmt.Println("SOMETHING WENT WRONG")
		return
	}

	getBookByID, err := bookCtrl.GetByID(1, bookCtrl)
	if err != nil {
		fmt.Println("Book ID Does Not Exist")
		return
	}
	fmt.Println("Book By ID ", getBookByID)

	deleteBook := bookCtrl.Delete(1, bookCtrl)
	if deleteBook != nil {
		fmt.Println("Unable To Delete This Book")
		return
	}

}
