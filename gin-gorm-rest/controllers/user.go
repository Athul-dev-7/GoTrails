package controllers

import (
	"gin-gorm-rest/config"
	"gin-gorm-rest/models"
	"log"
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
	result := config.DB.Find(&users)
	if result.Error != nil {
		log.Printf("Error while getting all users, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Got all users",
		"data":    &users,
	})

}

func CreateUsers(ctx *gin.Context) {
	user := []models.User{}
	ctx.BindJSON(&user)
	result := config.DB.Create(&user)
	if result.Error != nil {
		log.Printf("Error while creating new user, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"status":       http.StatusCreated,
		"message":      "Created new user",
		"data":         &user,
		"rowsAffected": result.RowsAffected,
	})
}

func UpdateUser(ctx *gin.Context) {
	user := []models.User{}
	result := config.DB.Where("id = ?", ctx.Param("userId")).First(&user)
	ctx.BindJSON(&user)
	result.Save(&user)
	if result.Error != nil {
		log.Printf("Error while updating user, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Updated user",
		"data":         &user,
		"rowsAffected": result.RowsAffected,
	})

}
