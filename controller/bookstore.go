package controller

import (
	"bookstore/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type BookStoreCtrl interface {
	Insert(book model.BookStore) error
	GetByID(ID int, book model.BookStore) (bookByID model.BookStore, err error)
	Delete(ID int, book model.BookStore) (bookByID model.BookStore, err error)
}

type BookStoreDB struct{}

func (b *BookStoreDB) Insert(book *[]model.BookStore, bs BookStoreDB) error {
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return errors.New("cannot do marshalling")
	}

	insertBook := os.WriteFile("book.json", bookJSON, 0644)
	if insertBook != nil {
		return errors.New("cannot write into json")
	}
	return nil
}

func (b *BookStoreDB) GetByID(ID int, bs BookStoreDB) (getBookByID []model.BookStore, err error) {
	bookList, err := os.ReadFile("book.json")
	if err != nil {
		return []model.BookStore{}, errors.New("cannot read json")
	}

	book := []model.BookStore{}

	err = json.Unmarshal(bookList, &book)
	if err != nil {
		panic(err)
	}

	bookID := []model.BookStore{}
	for _, v := range book {
		if ID == 1 {
			bookID = append(bookID, v)
		}
	}
	return bookID, nil
}

func (b *BookStoreDB) Delete(ID int, bs BookStoreDB) error {
	bookList, err := os.ReadFile("book.json")
	if err != nil {
		return errors.New("cannot read json")
	}

	book := []model.BookStore{}

	err = json.Unmarshal(bookList, &book)
	if err != nil {
		panic(err)
	}

	for _, v := range book {
		if v.ID != ID {
			book = append(book, v)
		}
	}

	err = os.Remove("book.json")
	fmt.Println("FAILED TO REMOVE BOOK ? ", err)

	bookJSON, _ := json.Marshal(book)
	insertBook := os.WriteFile("book.json", bookJSON, 0644)
	if insertBook != nil {
		return errors.New("cannot write into json")
	}
	return nil
}
