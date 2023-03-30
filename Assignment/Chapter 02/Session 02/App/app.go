package app

import (
	"github.com/gin-gonic/gin"
	c "ojam-test/c2/s2/app/controller"
)

const LOCAL_HOST = "localhost:8188"

func Boot() {
	router := gin.Default()
	c.BookController(router)
	router.Run(LOCAL_HOST)
}