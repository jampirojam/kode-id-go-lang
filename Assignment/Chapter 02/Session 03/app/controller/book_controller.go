package controller

import (
	"ojam-test/c2/s3/app/service"
	
	"github.com/gin-gonic/gin"
)

func BookController(router *gin.Engine) {
	router.GET("/book", getBook)
	router.POST("/book", addBook)
	router.PUT("/book/:id", updateBookById)
	router.DELETE("/book/:id", deleteBookById)
	router.GET("/book/:id", getBookById)
}

func getBook(c *gin.Context) {
	service.GetBook(c)
}

func addBook(c *gin.Context) {
	service.AddBook(c)
}

func updateBookById(c *gin.Context) {
	service.UpdateBookById(c)
}

func deleteBookById(c *gin.Context) {
	service.DeleteBookById(c)
}

func getBookById(c *gin.Context) {
	service.GetBookById(c)
}