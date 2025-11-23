package services

import (
	"context"
	"fmt"
	"time"

	"github.com/anilrajput6441/mcp_project/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTask(c *gin.Context, taskCol *mongo.Collection, title,description string) (interface{}, error) {
	userEmail := c.GetString("email")
	task := models.Task{
		Title: title,
		Description: description,
		Status: "todo",
		UserEmail: userEmail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	fmt.Println(task)

	res, err := taskCol.InsertOne(c,task)
	if err != nil {
		return nil, err
	}
	return gin.H{"id": res.InsertedID}, nil
}

func GetTasks(c *gin.Context, taskCol *mongo.Collection) ([]models.Task,error) {
	userEmail := c.GetString("email")
	

	cursor,err := taskCol.Find(c,bson.M{"user_email":userEmail})// why user
	
	if err != nil {
		return nil , err
	}
	

	var tasks []models.Task

	if err := cursor.All(c,&tasks); err != nil {
		return nil,err
	}
	return tasks,nil
}


func DeleteTask(c *gin.Context,taskCol *mongo.Collection,id string) error{
	userEmail := c.GetString("email")
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := taskCol.DeleteOne(c,bson.M{"_id":objId,"user_email":userEmail})
	return err
}

func UpdateTask(c *gin.Context, taskCol *mongo.Collection, id, title, description, status string) (string, error) {
	userEmail := c.GetString("email")
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objID, "user_email": userEmail}

	update := bson.M{
		"$set": bson.M{
			"title":       title,
			"description": description,
			"status":      status,
			"updated_at":  time.Now(),
		},
	}

	_, err := taskCol.UpdateByID(c, filter, update)
	if err != nil {
		return "", err
	}

	return "task updated", nil
}



////////////////////////////////////// MCP //////////////////////////////////////
//
//	MCP task services for AI
//
////////////////////////////////////// MCP //////////////////////////////////////




func GetTasksByEmail(ctx context.Context, taskCol *mongo.Collection, email string) ([]models.Task, error) {
	cursor, err := taskCol.Find(ctx, bson.M{"user_email": email})
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTaskFromAI(ctx context.Context, taskCol *mongo.Collection, email, title, description string) (interface{}, error) {
	task := models.Task{
		Title:       title,
		Description: description,
		Status:      "todo",
		UserEmail:   email,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := taskCol.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"id": res.InsertedID}, nil
}

func UpdateTaskFromAI(ctx context.Context, taskCol *mongo.Collection, email, id, title, description, status string) (interface{}, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objID, "user_email": email}

	update := bson.M{
		"$set": bson.M{
			"title":       title,
			"description": description,
			"status":      status,
			"updated_at":  time.Now(),
		},
	}

	_, err := taskCol.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return "task updated", nil
}

func DeleteTaskFromAI(ctx context.Context, taskCol *mongo.Collection, email, id string) (interface{},error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := taskCol.DeleteOne(ctx, bson.M{"_id": objID, "user_email": email})
	return "task deleted successfully", err
}




