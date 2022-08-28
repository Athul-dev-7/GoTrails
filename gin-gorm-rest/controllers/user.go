package controllers

import (
	"fmt"
	"log"
	"net/http"

	"gin-gorm-rest/config"
	"gin-gorm-rest/models"

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
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)
	// fmt.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
	if result.Error != nil {
		log.Printf("Error while getting all users, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	} else if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "0 users found",
			"data":    &users,
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
			"message": result.Error,
		})
	}
	if result.RowsAffected == 1 {
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
			"message": result.Error,
		})
	}
	if result.RowsAffected > 1 {
		ctx.IndentedJSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "Created new users",
			"data":    &users,
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
			"message": result.Error,
		})
	}
	if result.RowsAffected == 1 {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Updated user",
			"data":    &users,
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
			"message": result.Error,
		})
	}
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "No such user found",
		})
	}
	if result.RowsAffected == 1 {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User deleted",
		})
	}
}

func DeleteAllUser(ctx *gin.Context) {
	result := config.DB.Exec("DELETE FROM users")
	if result.Error != nil {
		log.Printf("Error while deleting all user, Reason : %v\n", result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": result.Error,
		})
	}
	if result.RowsAffected > 0 {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Deleted all user",
		})
	}
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "No records to delete",
		})
	}
}
