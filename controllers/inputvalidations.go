package controllers

import (
	entity "BookShopDatabase/entities"
	"net/url"
)

func IsValid(book entity.Book) url.Values {
	errors := url.Values{}
	if book.BookName == "" {
		errors.Add("Book", "The Book Name must Required!!")
	}
	if book.Author == "" {
		errors.Add("Author", "The Author Name must Required!!")
	}
	if length := len(book.Author); !(length > 2 && length < 20) {
		errors.Add("Author", "The Author Name must be in between 2-20 Characters!!")
	}
	if length := len(book.BookName); !(length > 2 && length < 20) {
		errors.Add("BookName", "The BookName must be in between 2-20 Characters!!")
	}
	if book.BookID <= 0 {
		errors.Add("BookId", "The book Id cannot Negative")
	}
	if book.SoldCount < 0 {
		errors.Add("SoldCount", "The soldCount field cannot be Negative")
	}
	if book.ReviewScore == 0.0 {
		errors.Add("ReviewScore", "Provide a proper Review. The ReviewScore must Required!!")
	}
	return errors
}
