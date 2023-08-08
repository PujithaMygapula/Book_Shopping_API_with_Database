package controllers

import (
	database "BookShopDatabase/Database"
	entity "BookShopDatabase/entities"
	"strings"

	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooksData(c *gin.Context) {
	var books []entity.Book
	db := database.SetupDB()
	rows, err := db.Query("select * from books")

	if err != nil {
		return
	}
	for rows.Next() {
		var bookId int64
		var bookName string
		var author string
		var soldCount int64
		var reviewScore float64
		_ = rows.Scan(&bookId, &bookName, &author, &soldCount, &reviewScore)

		books = append(books, entity.Book{BookID: bookId, BookName: bookName, Author: author, SoldCount: soldCount, ReviewScore: reviewScore})
	}

	if len(books) == 0 {
		c.IndentedJSON(http.StatusOK, "No Book Record Found")
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func GetByBookName(book_name string) ([]entity.Book, error) {
	var books []entity.Book
	db := database.SetupDB()
	rows, err := db.Query("SELECT * from books WHERE bookname = $1", strings.ToLower(book_name))

	if err != nil {
		return nil, errors.New("no book data found with provided book name")
	}

	count := 0
	for rows.Next() {
		var bookId int64
		var bookName string
		var author string
		var soldCount int64
		var reviewScore float64
		count++

		_ = rows.Scan(&bookId, &bookName, &author, &soldCount, &reviewScore)
		books = append(books, entity.Book{BookID: bookId, BookName: bookName, Author: author, SoldCount: soldCount, ReviewScore: reviewScore})
	}

	if count == 0 {
		return nil, errors.New("no book data found with provided book name")
	}

	return books, nil
}

func GetBookByBookName(c *gin.Context) {
	book_name := c.Param("bookName")
	sameNameBooks, err := GetByBookName(book_name)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found with provided book name"})
		return
	}
	c.IndentedJSON(http.StatusOK, sameNameBooks)
}

func GetByBookId(book_id int) ([]entity.Book, error) {
	var book []entity.Book
	db := database.SetupDB()
	rows, err := db.Query("SELECT * from books WHERE bookid = $1", book_id)

	if err != nil {
		return nil, errors.New("no book data found with provided book id")
	}
	count := 0
	for rows.Next() {
		count++
		var bookId int64
		var bookName string
		var author string
		var soldCount int64
		var reviewScore float64

		_ = rows.Scan(&bookId, &bookName, &author, &soldCount, &reviewScore)
		book = append(book, entity.Book{BookID: bookId, BookName: bookName, Author: author, SoldCount: soldCount, ReviewScore: reviewScore})
	}

	if count == 0 {
		return nil, errors.New("no book data found with provided book id")
	}
	return book, nil
}

func GetBookByBookId(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.Param("bookId"))
	book, err := GetByBookId(bookId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Book data found with provided book id"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}

func AddBookData(c *gin.Context) {
	var newBook entity.Book
	db := database.SetupDB()

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong while Adding a Book"})
		return
	}
	checking := IsValid(newBook)
	if len(checking) != 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": checking})
		return
	}
	_, err := db.Exec("insert into books(bookid, bookname, author, soldcount, reviewscore) values($1,$2,$3,$4,$5)",
		newBook.BookID, strings.ToLower(newBook.BookName), newBook.Author, newBook.SoldCount, newBook.ReviewScore)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Some error occured while Adding a Book"})
		return
	}
	c.IndentedJSON(http.StatusOK, newBook)
}

func UpdateBookByBookId(c *gin.Context) {
	book_id, _ := strconv.Atoi(c.Param("bookId"))
	sold_count, _ := strconv.Atoi(c.Param("soldCount"))
	_, err := GetByBookId(book_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found"})
		return
	}

	db := database.SetupDB()
	_, err1 := db.Exec("UPDATE books set soldcount = $1 WHERE bookid = $2", sold_count, book_id)

	if err1 != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Some thing went wrong- while fetching data"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated Successfully !!"})
}

func DeleteByBookId(book_id int) (string, error) {
	db := database.SetupDB()
	rows, err := db.Query("SELECT * from books WHERE bookId = $1", book_id)

	if err != nil || !rows.Next() {
		return " ", errors.New("no book data found with provided book id")
	}

	_, err1 := db.Exec("delete from books where bookId = $1", book_id)
	if err1 != nil {
		return " ", err1
	}
	return "Book Data Deleted Successfully", nil
}

func DeleteBookByBookId(c *gin.Context) {
	book_id, _ := strconv.Atoi(c.Param("bookId"))
	msg, err := DeleteByBookId(book_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, msg)
}
