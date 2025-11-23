package handlers

import (
	"net/http"

	"github.com/anilrajput6441/mcp_project/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)



func RegisterHandler(userCol *mongo.Collection) gin.HandlerFunc {
	return func (c *gin.Context){
		var body struct {
			Email string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
			FullName string `json:"full_name" binding:"required"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := services.RegisterUser(c, userCol, body.Email, body.Password, body.FullName)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "user registered"})
	}
}

func LoginHandler(userCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context){
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid data"})
			return
		}

		tokenRes, err := services.LoginUser(c,userCol,body.Email,body.Password)
		
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, tokenRes)
	}
}

func RefreshHandler(usersCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid data"})
			return
		}

		tokenRes, err := services.RefreshAccessToken(c, usersCol, body.RefreshToken)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, tokenRes)
	}
}
