package services

import (
	"context"
	"errors"
	"time"

	"github.com/anilrajput6441/mcp_project/internal/models"
	"github.com/anilrajput6441/mcp_project/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


func RegisterUser(ctx context.Context, userCol *mongo.Collection, email,password,fullName string) error {
	//check if user already exists 
	count , _ := userCol.CountDocuments(ctx, bson.M{"email":email})

	if count > 0 {
		return errors.New("email already registered")
	}

	//hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(password),14)

	user := models.User{
		Email: email,
		Password: string(hash),
		FullName: fullName,
		Role: "user",
		CreatedAt: time.Now(),
	}

	_,err := userCol.InsertOne(ctx,user)
	return err

}

func LoginUser(ctx context.Context, userCol *mongo.Collection, email string, password string) (map[string]string, models.User, error){
	
	var user models.User
	err := userCol.FindOne(ctx,bson.M{"email":email}).Decode(&user)

	if err != nil {
		return nil, models.User{}, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, models.User{}, errors.New("invalid credentials")
	}

	// generate tokens
	access, _ := utils.GenerateAccessToken(user.Email, user.Role, user.ID)
	refresh, _ := utils.GenerateRefreshToken(user.Email)

	res := map[string]string{
		"access_token":  access,
		"refresh_token": refresh,
	}

	return res, user, nil
}


func RefreshAccessToken(ctx context.Context, usersCol *mongo.Collection, refreshToken string) (map[string]string, error) {
	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	email := claims["email"].(string)

	var user models.User
	err = usersCol.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// generate new access token
	access, _ := utils.GenerateAccessToken(user.Email, user.Role, user.ID)

	return map[string]string{
		"access_token": access,
	}, nil
}



func ChangePassword(ctx context.Context, userCol *mongo.Collection, email, oldPassword, newPassword string) error {
	

		if len(newPassword) < 6 {
			return errors.New("password too short")
		}
		var user models.User
		err := userCol.FindOne(ctx, bson.M{"email": email}).Decode(&user)
		if err != nil {
			return errors.New("user not found")
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
		if err != nil {
			return errors.New("invalid old password")
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 14)

		_, err = userCol.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"password": string(hash)}})
		if err != nil {
			return errors.New("failed to update password")
		}

		return nil
}
