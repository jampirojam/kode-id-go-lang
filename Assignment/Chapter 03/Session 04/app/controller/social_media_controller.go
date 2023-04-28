package controller

import (
	"ojam-test/c3/s4/app/config"
	handler "ojam-test/c3/s4/app/handler/impl"
	repository "ojam-test/c3/s4/app/repository/impl"
	service "ojam-test/c3/s4/app/service/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SocialMediaController(router *gin.Engine, db *gorm.DB) {
	r := repository.SocialMediaRepositoryImpl(db)
	s := service.SocialMediaServiceImpl(r)
	h := handler.SocialMediaHandlerImpl(s)

	group := router.Group("/v1/social-media", config.Auth); {
		group.POST("/", h.Add)
		group.PUT("/:id", h.UpdateById)
		group.DELETE("/:id", h.DeleteById)
		group.GET("/:id", h.GetById)
		group.GET("/", h.GetAll)
	}
}