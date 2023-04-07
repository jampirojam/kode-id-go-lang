package app

import (
	ctrl "ojam-test/c2/s3/app/controller"
	ct "ojam-test/c2/s3/app/constant"
	"github.com/gin-gonic/gin"
)

func Boot() {
	router := gin.Default()
	ctrl.BookController(router)
	router.Run(ct.LOCAL_HOST)
}