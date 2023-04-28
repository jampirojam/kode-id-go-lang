package controller

import (
	"ojam-test/c3/s4/app/config"
	handler "ojam-test/c3/s4/app/handler/impl"
	repository "ojam-test/c3/s4/app/repository/impl"
	service "ojam-test/c3/s4/app/service/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentController(router *gin.Engine, db *gorm.DB) {
	pr := repository.PhotoRepositoryImpl(db)
	cr := repository.CommentRepositoryImpl(db)
	s := service.CommentServiceImpl(pr, cr)
	h := handler.CommentHandlerImpl(s)

	group := router.Group("/v1/comment", config.Auth); {
		group.POST("/", h.Add)
		group.PUT("/:id", h.UpdateById)
		group.DELETE("/:id", h.DeleteById)
		group.GET("/:id", h.GetById)
		group.GET("/", h.GetAll)
	}
}