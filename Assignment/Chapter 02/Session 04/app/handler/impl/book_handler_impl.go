package impl

import (
	"net/http"
	"ojam-test/c2/s4/app/dto/request"
	"ojam-test/c2/s4/app/dto/response"
	"ojam-test/c2/s4/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Svc service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		Svc: bookService,
	}
}

func (handler *BookHandler) GetAll(ctx *gin.Context) {
	res, err := handler.Svc.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

func (handler *BookHandler) UpdateById(ctx *gin.Context) {
	var req request.BookRequest
	paramId := ctx.Param("id")

	id, _ := strconv.Atoi(paramId)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	res, err := handler.Svc.UpdateById(id, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

func (handler *BookHandler) GetById(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, _ := strconv.Atoi(paramId)

	res, err := handler.Svc.GetById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

func (handler *BookHandler) DeleteById(ctx *gin.Context) {
	paramId := ctx.Param("id")

	id, _ := strconv.Atoi(paramId)

	res, err := handler.Svc.DeleteById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}

func (handler *BookHandler) Create(ctx *gin.Context) {
	var req request.BookRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	res, err := handler.Svc.Create(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"context" : res,
	})
}