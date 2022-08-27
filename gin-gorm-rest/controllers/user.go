package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome!",
	})
}
