package main

import (
	"user-management/config"
	"user-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.UserRoutes(r)

	r.Run(":8080")
}
