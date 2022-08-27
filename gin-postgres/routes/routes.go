package routes

import (
	"gin-postgres/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/todos", controllers.GetAlltodos)

}

func welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "welcome to API",
	})
}
