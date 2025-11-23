package routes

import (
	"github.com/anilrajput6441/mcp_project/internal/handlers"
	"github.com/anilrajput6441/mcp_project/internal/middleware"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

func TaskRoutes(router *gin.Engine, db *mongo.Database) {
	taskCol := db.Collection("tasks")

	taskGroup := router.Group("/tasks")
	taskGroup.Use(middleware.AuthMiddleware()) // <--- PROTECTED
	{
		taskGroup.POST("/", handlers.CreateTask(taskCol))
		taskGroup.GET("/", handlers.GetTasks(taskCol))
		taskGroup.PUT("/:id", handlers.UpdateTask(taskCol))
		taskGroup.DELETE("/:id", handlers.DeleteTask(taskCol))
	}
}
