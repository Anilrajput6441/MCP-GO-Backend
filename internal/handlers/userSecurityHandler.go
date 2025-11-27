package handlers

import (
	"net/http"

	"github.com/anilrajput6441/mcp_project/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ChangePassword(userCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {

		email := c.GetString("email") // from JWT

		var body struct {
			OldPassword string `json:"oldPassword"`
			NewPassword string `json:"newPassword"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		if len(body.NewPassword) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password too short"})
			return
		}

		if err := services.ChangePassword(
			c.Request.Context(),
			userCol,
			email,
			body.OldPassword,
			body.NewPassword,
		); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "password updated"})
	}
}
