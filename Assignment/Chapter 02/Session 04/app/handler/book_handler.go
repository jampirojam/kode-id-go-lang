package handler

import (
	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	GetAll(ctx *gin.Context)
	UpdateById(ctx *gin.Context)
	GetById(ctx *gin.Context)
	DeleteById(ctx *gin.Context)
	Create(ctx *gin.Context)
}