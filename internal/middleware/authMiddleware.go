package middleware

import (
	"net/http"
	"strings"

	"github.com/anilrajput6441/mcp_project/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// OPTIONS requests are handled by CORS middleware, skip auth for them
        if c.Request.Method == "OPTIONS" {
            c.Next()
            return
        }

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Extract user info and attach to request context
		email, _ := claims["email"].(string)
		role, _ := claims["role"].(string)
		userID, _ := claims["_id"].(string)
		
		if email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token: missing email"})
			c.Abort()
			return
		}
		
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token: missing user ID"})
			c.Abort()
			return
		}

		c.Set("email", email)
		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}
