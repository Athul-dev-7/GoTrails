package controllers

import (
	"fmt"
	"gin-gorm-rest/config"
	"gin-gorm-rest/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Got all users",
			"data":    &users,
		})
	}

}

func GetUser(ctx *gin.Context) {
	users := []models.User{}
	result := config.DB.Where("id = ?", ctx.Param("userId")).Take(&users)
	if result.Error != nil {
		log.Printf("Error while getting specified user, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Got a single user",
			"data":    &users,
		})
	}
}

func CreateUsers(ctx *gin.Context) {
	users := []models.User{}
	ctx.BindJSON(&users)
	result := config.DB.Create(&users)
	if result.Error != nil {
		log.Printf("Error while creating new user, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, gin.H{
			"status":       http.StatusCreated,
			"message":      "Created new user",
			"data":         &users,
			"rowsAffected": result.RowsAffected,
		})
	}

}

func UpdateUser(ctx *gin.Context) {
	users := []models.User{}
	result := config.DB.Where("id = ?", ctx.Param("userId")).First(&users)
	ctx.BindJSON(&users)
	result.Save(&users)
	if result.Error != nil {
		log.Printf("Error while updating user, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":       http.StatusOK,
			"message":      "Updated user",
			"data":         &users,
			"rowsAffected": result.RowsAffected,
		})
	}

}

func DeleteUser(ctx *gin.Context) {
	users := []models.User{}
	result := config.DB.Where("id = ?", ctx.Param("userId")).Delete(&users)
	fmt.Println(result)
	if result.Error != nil {
		log.Printf("Error while updating user, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":       http.StatusOK,
			"message":      "Deleted user",
			"rowsAffected": result.RowsAffected,
		})
	}

}

func DeleteAllUser(ctx *gin.Context) {
	users := []models.User{}
	result := config.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&users)
	if result.Error != nil {
		log.Printf("Error while deleting all user, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":       http.StatusOK,
			"message":      "Deleted all user",
			"rowsAffected": result.RowsAffected,
		})
	}

}
