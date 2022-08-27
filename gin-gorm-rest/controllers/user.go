package controllers

import (
	"gin-gorm-rest/config"
	"gin-gorm-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome!",
	})
}

func GetAllUsers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Got all users",
		"data":    &users,
	})

}
