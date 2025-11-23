package handlers

import (
	"github.com/anilrajput6441/mcp_project/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTask(taskCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context){
		var body struct{
			Title string `json:"title"`
			Description string `json:"description"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400,gin.H{"error":"invalid data"})
			return 
		}

		result, err := services.CreateTask(c, taskCol, body.Title, body.Description)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, result)

	}
}

func GetTasks(taskCol *mongo.Collection) gin.HandlerFunc {
	
	return func(c *gin.Context) {
		result, err := services.GetTasks(c, taskCol)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(200, result)
	}
}

func DeleteTask(taskCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := services.DeleteTask(c, taskCol, id)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "task deleted"})
	}
}

func UpdateTask(taskCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var body struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Status      string `json:"status"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid data"})
			return
		}

		result, err := services.UpdateTask(c, taskCol, id, body.Title, body.Description, body.Status)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, result)
	}
}