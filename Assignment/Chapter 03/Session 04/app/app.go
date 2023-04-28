package app

import (
	"ojam-test/c3/s4/app/config"
	"ojam-test/c3/s4/app/constant"
	"ojam-test/c3/s4/app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "ojam-test/c3/s4/docs"
)

func Boot() {
	db := config.GetDatabaseConnection()
	router := gin.Default()
	Router(router, db)
	router.Run(constant.PORT_SERVER)
}

func Router(router *gin.Engine, db *gorm.DB) {
	controller.UserController(router, db)
	controller.PhotoController(router, db)
	controller.CommentController(router, db)
	controller.SocialMediaController(router, db)

	controller.SwaggerController(router)
}