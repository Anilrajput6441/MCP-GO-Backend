package routes

import (
	"github.com/anilrajput6441/mcp_project/internal/handlers"
	"github.com/anilrajput6441/mcp_project/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(router *gin.Engine, db *mongo.Database) {
	userCol := db.Collection("users")
	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("", handlers.GetUsers(userCol))
		userGroup.GET("/", handlers.GetUsers(userCol))
		userGroup.PUT("/", handlers.UpdateUser(userCol))
		userGroup.DELETE("/", handlers.DeleteUser(userCol))
		userGroup.PUT("/change-password", handlers.ChangePassword(userCol))
	}
}