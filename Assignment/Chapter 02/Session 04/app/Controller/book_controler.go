package controller

import (
	handler "ojam-test/c2/s4/app/handler/impl"
	repository "ojam-test/c2/s4/app/repository/impl"
	service "ojam-test/c2/s4/app/service/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookController(router *gin.Engine, db *gorm.DB) {
	r := repository.NewBookRepository(db)
	s := service.NewBookService(r)
	h := handler.NewBookHandler(s)

	group := router.Group("/book"); {
		group.GET("/", h.GetAll)
		group.PUT("/:id", h.UpdateById)
		group.GET("/:id", h.GetById)
		group.DELETE("/:id", h.DeleteById)
		group.POST("/", h.Create)
	}
}