package routes

import (
	"gin-gorm-rest/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.Home)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:userId", controllers.GetUser)
	router.POST("/users", controllers.CreateUsers)
	router.PUT("/users/:userId", controllers.UpdateUser)
	router.DELETE("/users/:userId", controllers.DeleteUser)
	router.DELETE("/users", controllers.DeleteAllUser)
}
