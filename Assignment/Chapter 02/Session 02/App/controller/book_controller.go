package controller

import (
	"github.com/gin-gonic/gin"
	service "ojam-test/second-session/service"
)

func BookController() {
	router := gin.Default()
	router.POST("/book/create", createBook)
}

func createBook(c *gin.Context) {
	service.CreateBookService
}