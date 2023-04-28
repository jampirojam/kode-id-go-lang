package handler

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	Add(ctx *gin.Context)
	UpdateById(ctx *gin.Context)
	DeleteById(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}