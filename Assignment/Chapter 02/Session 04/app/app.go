package app

import (
	"ojam-test/c2/s4/app/config"
	"ojam-test/c2/s4/app/constant"
	"ojam-test/c2/s4/app/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Boot() {
	db := config.GetDatabaseConnection()
	router := gin.Default()
	Router(router, db)
	router.Run(constant.PORT_SERVER)
}

func Router(router *gin.Engine, db *gorm.DB) {
	controller.BookController(router, db)
}