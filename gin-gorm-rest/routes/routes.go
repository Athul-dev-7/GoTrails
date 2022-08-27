package routes

import (
	"gin-gorm-rest/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.Home)
}
