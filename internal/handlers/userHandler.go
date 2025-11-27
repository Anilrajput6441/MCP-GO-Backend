package handlers

import (
	"net/http"

	"github.com/anilrajput6441/mcp_project/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(userCol *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		
		if userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user ID not found in token"})
			return
		}
		
		// Convert string ID to ObjectID
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID format"})
			return
		}
		
		user := models.User{}
		err = userCol.FindOne(c, bson.M{"_id": objectID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		// Remove password from response
		user.Password = ""
		c.JSON(http.StatusOK, user)
	}

}

func UpdateUser(userCol *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("userID")
		
		if userID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user ID not found in token"})
			return
		}
		
		// Convert string ID to ObjectID
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID format"})
			return
		}
		
		var updateData struct {
			FullName string `json:"full_name"`
			Email    string `json:"email"`
		}
		
		if err := ctx.ShouldBindJSON(&updateData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Build update document with only fields that are provided
		update := bson.M{}
		if updateData.FullName != "" {
			update["full_name"] = updateData.FullName
		}
		if updateData.Email != "" {
			update["email"] = updateData.Email
		}
		
		if len(update) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
			return
		}

		// Update the user
		_, err = userCol.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": update})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		// Fetch and return the updated user
		var updatedUser models.User
		err = userCol.FindOne(ctx, bson.M{"_id": objectID}).Decode(&updatedUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user updated but failed to fetch updated data"})
			return
		}
		
		// Remove password from response
		updatedUser.Password = ""
		ctx.JSON(http.StatusOK, updatedUser)
	}
}

func DeleteUser(userCol *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetString("userID")
		
		if userID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user ID not found in token"})
			return
		}
		
		// Convert string ID to ObjectID
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID format"})
			return
		}
		
		_, err = userCol.DeleteOne(ctx, bson.M{"_id": objectID})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
	}
}