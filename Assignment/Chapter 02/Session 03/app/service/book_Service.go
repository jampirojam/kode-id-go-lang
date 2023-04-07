package service

import (
	"fmt"
	"log"
	"net/http"
	"ojam-test/c2/s3/app/config"
	"ojam-test/c2/s3/app/entity"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var db = config.GetDatabaseConnection()

func GetBook(ctx *gin.Context) {
	var books []entity.Book
	var book entity.Book
	sqlQuery := `SELECT * FROM book`

	result, err := db.Query(sqlQuery)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error" : err.Error()},
		)
		return
	}

	for result.Next() {
		err = result.Scan(&book.Id, &book.Title, &book.Author, &book.Publisher, &book.Year)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error" : err.Error()},
			)
		}

		books = append(books, book)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"content": books,
	})
}

func AddBook(ctx *gin.Context) {
	var book entity.Book
	var finalBook entity.Book
	sqlQuery := `
		INSERT INTO book(title, author, publisher, year) 
		VALUES($1, $2, $3, $4) 
		RETURNING id
	`

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	result := db.QueryRow(sqlQuery, book.Title, book.Author, book.Publisher, book.Year)
	err := result.Scan(&finalBook.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success added new book with Id: %d", finalBook.Id),
	})	
}

func UpdateBookById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	var book entity.Book
	sqlQuery := `
		UPDATE book
		SET title = $2, author = $3, publisher = $4, year = $5
		WHERE id = $1
		RETURNING id
	`

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.QueryRow(sqlQuery, id, book.Title, book.Author, book.Publisher, book.Year)

	err = result.Scan(&book.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated data with Id: %d", id),
		"data_update" : book,
	})	
}

func DeleteBookById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	var book entity.Book
	sqlQuery := `DELETE FROM book WHERE id = $1 RETURNING id`

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.QueryRow(sqlQuery, id)

	err = result.Scan(&book.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted data with Id: %d", id),
	})	
}

func GetBookById(ctx *gin.Context) {
	paramId := ctx.Param("id")
	var book entity.Book
	sqlQuery := `SELECT * FROM book WHERE id = $1`

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println(sqlQuery)
	result, err := db.Query(sqlQuery, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	for result.Next() {
		err = result.Scan(&book.Id, &book.Title, &book.Author, &book.Publisher, &book.Year)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error" : err.Error()},
			)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : book,
	})	
}