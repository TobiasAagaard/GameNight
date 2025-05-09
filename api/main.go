package main

import (
	"game_night/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.SetupRoutes(server)
	server.Run("8080")
}
