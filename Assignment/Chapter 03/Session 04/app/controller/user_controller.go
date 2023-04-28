package controller

import (
	handler "ojam-test/c3/s4/app/handler/impl"
	repository "ojam-test/c3/s4/app/repository/impl"
	service "ojam-test/c3/s4/app/service/impl"


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
	}
}