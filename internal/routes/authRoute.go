package routes

import (
	"github.com/anilrajput6441/mcp_project/internal/handlers"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRoutes(router *gin.Engine, db *mongo.Database) {
	authGroup := router.Group("/auth")

	usersCol := db.Collection("users")

	authGroup.POST("/register", handlers.RegisterHandler(usersCol))
	authGroup.POST("/login", handlers.LoginHandler(usersCol))
	authGroup.POST("/refresh", handlers.RefreshHandler(usersCol))

}
