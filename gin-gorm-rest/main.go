package main

import (
	"gin-gorm-rest/config"
	"gin-gorm-rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	router := gin.Default()

	routes.Routes(router)

	router.Run()
}
