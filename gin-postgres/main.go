package main

import (
	"gin-postgres/configs"
	"gin-postgres/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Connect()
	router := gin.Default()
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
