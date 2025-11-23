package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(router *gin.Engine, db *mongo.Database) {
    // User Routes
    // Task Routes
	TaskRoutes(router, db)

  
    MCPRoutes(router, db)


	//auth routes
	AuthRoutes(router, db)

    // Example test endpoint
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
}
