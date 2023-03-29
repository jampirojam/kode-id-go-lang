package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
	m "ojam-test/second-session/model"
)

func CreateBookService(ctx *gin.Context) []m.Book {
	book := m.Book
	books := []m.Book{}

	if err := ctx.BindJSON(&book); err != nil {
		return
	}

	books = append(books, book)

	ctx.IntendedJSON(http.StatusCreated, book)

	return books
}