package controller

import (
	"ojam-test/c3/s2/app/config"
	handler "ojam-test/c3/s2/app/handler/impl"
	repository "ojam-test/c3/s2/app/repository/impl"
	service "ojam-test/c3/s2/app/service/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductController(router *gin.Engine, db *gorm.DB) {
	r := repository.ProductRepositoryImpl(db)
	s := service.ProductServiceImpl(r)
	h := handler.ProductHandlerImpl(s)

	group := router.Group("/v1/product", config.Auth); {
		group.POST("/", h.Add)
		group.PUT("/:id", h.UpdateById)
		group.DELETE("/:id", h.DeleteById)
		group.GET("/:id", h.GetById)
		group.GET("/", h.GetAll)
	}
}