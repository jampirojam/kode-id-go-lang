package controller

import (
	"ojam-test/c3/s2/app/config"
	handler "ojam-test/c3/s2/app/handler/impl"
	repository "ojam-test/c3/s2/app/repository/impl"
	service "ojam-test/c3/s2/app/service/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserController(router *gin.Engine, db *gorm.DB) {
	r := repository.UserRepositoryImpl(db)
	s := service.UserServiceImpl(r)
	h := handler.UserHandlerImpl(s)

	group := router.Group("/v1/user"); {
		group.POST("/", h.Register)
		group.POST("/login", h.Login)
		group.POST("/relog", config.RefreshAuth, h.Relog)
	}
}