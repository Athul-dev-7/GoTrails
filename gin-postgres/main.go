package main

import (
	"gin-postgres/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
