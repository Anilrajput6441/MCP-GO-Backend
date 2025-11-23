package main

import (
	"log"

	"github.com/anilrajput6441/mcp_project/internal/config"
	"github.com/anilrajput6441/mcp_project/internal/db"
	"github.com/anilrajput6441/mcp_project/internal/middleware"
	"github.com/anilrajput6441/mcp_project/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    // Load ENV variables
    config.LoadEnv()

    // Connect to database
    client, database := db.ConnectMongo()
    defer client.Disconnect(nil)

  


    // Initialize router
    router := gin.Default()

    router.Use(middleware.CORS())


	router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "health is good "})
    })

    // Register all routes
    routes.RegisterRoutes(router, database)

  
    // Start server
    log.Println("Server running on port 8080")
    router.Run(":8080")
}
