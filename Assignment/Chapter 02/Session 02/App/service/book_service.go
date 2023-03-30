package service

import (
	"fmt"
	"net/http"
	"ojam-test/c2/s2/app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []model.Book{}

func CreateBook(ctx *gin.Context) {
	var book model.Book
	var id int

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if book.Id != 0 {
		for i := 0; i < len(books); i++ {
			if book.Id == books[i].Id {
				key, value := ErrorMessage(http.StatusConflict, fmt.Sprintf("Book with ID: %d found", id))
				ctx.IndentedJSON(http.StatusNotFound, gin.H{key: value})
				return
			}
		}
	}

	for i := 0; i < len(books); i++ {
		if i == (len(books) - 1) {
			id = books[i].Id + 1
		}
	}

	if id == 0 {
		id = 1
	}

	book.Id = id

	books = append(books, book)

	ctx.IndentedJSON(http.StatusCreated, book)
}

func GetBook(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
}

func GetBookById(ctx *gin.Context) {
	id := ctx.Param("id")

	book := findById(id)

	if book == nil {
		key, value := ErrorMessage(http.StatusNotFound, fmt.Sprintf("Book with ID: %s", id))
		ctx.IndentedJSON(http.StatusNotFound, gin.H{key: value})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}

func UpdateBookById(ctx *gin.Context) {
	id := ctx.Param("id")

	var update model.Book

	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := findById(id)

	if book == nil {
		key, value := ErrorMessage(http.StatusNotFound, fmt.Sprintf("Book with ID: %s", id))
		ctx.IndentedJSON(http.StatusNotFound, gin.H{key: value})
		return
	}

	*book = update
	book.Id, _ = strconv.Atoi(id)

	ctx.IndentedJSON(http.StatusOK, *book)
}

func DeleteBookById(ctx *gin.Context) {
	id := ctx.Param("id")

	book := findById(id)

	if book == nil {
		key, value := ErrorMessage(http.StatusNotFound, fmt.Sprintf("Book with ID: %s", id))
		ctx.IndentedJSON(http.StatusNotFound, gin.H{key: value})
		return
	}

	var newBook = []model.Book{}
	newId, _ := strconv.Atoi(id)
	var deleted bool

	for i := 0; i < len(books); i++ {
		if newId == books[i].Id {
			deleted = true
		} else {
			newBook = append(newBook, books[i])
		}
	}

	books = newBook

	if deleted {
		key, value := SuccessMessage(http.StatusAccepted, fmt.Sprintf("Book with ID: %s successfully deleted", id))
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			key:    value,
			"book": newBook,
		})
	}
}

func findById(id string) *model.Book {

	newId, _ := strconv.Atoi(id)

	for i, book := range books {
		if book.Id == newId {
			return &books[i]
		}
	}

	return nil
}
