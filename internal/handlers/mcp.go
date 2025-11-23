package handlers

import (
	"github.com/anilrajput6441/mcp_project/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MCPListTasks(taskCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		emailRaw, _ := c.Get("email")
		email := emailRaw.(string)

	

		tasks, err := services.GetTasksByEmail(c.Request.Context(), taskCol, email)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"tasks": tasks})
	}
}

func MCPCreateTask(taskCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

        // Get email from middleware
        emailRaw, _ := c.Get("email")
        email := emailRaw.(string)

		var body struct {
			Title       string `json:"title"`
			Description string `json:"description"`	
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		res, err := services.CreateTaskFromAI(c.Request.Context(), taskCol, email, body.Title, body.Description)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"id": res})
	}
}

func MCPUpdateTask(taskCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {


		emailRaw, _ := c.Get("email")
		email := emailRaw.(string)

		var body struct {
			Id          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Status      string `json:"status"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		res, err := services.UpdateTaskFromAI(c.Request.Context(), taskCol, email, body.Id, body.Title, body.Description, body.Status)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"result": res})
	}
}

func MCPDeleteTask(taskCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		emailRaw, _ := c.Get("email")
		email := emailRaw.(string)

		var body struct {
			Id    string `json:"id"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"})
			return
		}

		res, err := services.DeleteTaskFromAI(c.Request.Context(), taskCol, email, body.Id)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"result": res})
	}
}
