package main

import (
	
	"ngodinginaja-be/config"
	"ngodinginaja-be/routes"
	// "ngodinginaja-be/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
	
}
