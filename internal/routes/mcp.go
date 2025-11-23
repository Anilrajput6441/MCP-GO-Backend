package routes

import (
	"github.com/anilrajput6441/mcp_project/internal/handlers"
	"github.com/anilrajput6441/mcp_project/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MCPRoutes(router *gin.Engine, db *mongo.Database) {
	taskCol := db.Collection("tasks")

	mcp := router.Group("/mcp/task")
	mcp.Use(middleware.AuthMiddleware())
	{
		mcp.POST("/list", handlers.MCPListTasks(taskCol))
		mcp.POST("/create", handlers.MCPCreateTask(taskCol))
		mcp.POST("/update/:id", handlers.MCPUpdateTask(taskCol))
		mcp.DELETE("/delete/:id", handlers.MCPDeleteTask(taskCol))
	}
}
