package app

import (
	"github.com/gin-gonic/gin"
	c "ojam-test/second-session/controller"
)

const LOCAL_HOST = "localhost:8188"

func Boot() {
	router := gin.Default()
	c.BookController
	router.Run(LOCAL_HOST)
}