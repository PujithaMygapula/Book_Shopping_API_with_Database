package main

import (
	controller "BookShopDatabase/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getbooks", controller.GetBooksData)
	router.GET("/getbookbyname/:bookName", controller.GetBookByBookName)
	router.GET("/getbookbyid/:bookId", controller.GetBookByBookId)
	router.POST("/addbook", controller.AddBookData)
	router.DELETE("/deletebookbyid/:bookId", controller.DeleteBookByBookId)
	router.PATCH("/updatebooksoldcount/:bookId/:soldCount", controller.UpdateBookByBookId)
	router.Run("localhost:1211")
}
