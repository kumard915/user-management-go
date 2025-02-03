package services

import (
	"context"
	"time"
	"user-management/config"
	"user-management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.DB.Collection("users").InsertOne(ctx, user)
	return err
}

func GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []models.User
	cursor, err := config.DB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &users)
	return users, err
}

func GetUserByID(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	objectID, _ := primitive.ObjectIDFromHex(id)
	err := config.DB.Collection("users").FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	return user, err
}

func UpdateUser(id string, user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := config.DB.Collection("users").UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": user})
	return err
}

func DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := config.DB.Collection("users").DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
