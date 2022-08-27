package main

import (
	"gin-gorm-rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Routes(router)

	router.Run()
}
