package services

import (
	"errors"
	"w2s-backend/database"
	"w2s-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserProfile(userId string) (*models.User, error) {
	res := database.FindOne("users", bson.M{"userId": userId}, nil)
	if err := res.Err(); err != nil {
		return nil, errors.New("user not found")
	}

	var user models.User
	if err := res.Decode(&user); err != nil {
		return nil, errors.New("failed to decode user")
	}

	user.Password = ""
	return &user, nil
}

func UpdateUserEmail(userId, newEmail string) error {

	if err := database.FindOne("users", bson.M{"email": newEmail}, nil).Err(); err == nil {
		return errors.New("email already in use")
	}

	update := bson.M{"$set": bson.M{"email": newEmail}}
	_, err := database.UpdateOne("users", bson.M{"userId": userId}, update)
	return err
}

func UpdateUserProfile(userId, name string) error {
	update := bson.M{"$set": bson.M{"name": name}}
	_, err := database.UpdateOne("users", bson.M{"userId": userId}, update)
	return err
}
