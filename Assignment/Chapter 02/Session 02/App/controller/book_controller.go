package controller

import (
	"github.com/gin-gonic/gin"
	"ojam-test/c2/s2/app/service"
)

func BookController(router *gin.Engine) {
	router.POST("/book/create", createBook)
	router.GET("/book", getBooks)
	router.GET("/book/:id", getBookById)
	router.PUT("/book/update/:id", updateBookById)
	router.DELETE("/book/delete/:id", deleteBookById)
}

func createBook(c *gin.Context) {
	service.CreateBook(c)
}

func getBooks(c *gin.Context) {
	service.GetBook(c)
}

func getBookById(c *gin.Context) {
	service.GetBookById(c)
}

func updateBookById(c *gin.Context) {
	service.UpdateBookById(c)
}

func deleteBookById(c *gin.Context) {
	service.DeleteBookById(c)
}