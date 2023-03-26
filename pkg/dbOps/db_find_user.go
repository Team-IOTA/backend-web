package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func FindUser (username string) (*models.Model_User, error) {
	var object models.Model_User

	err := connections.DATABASE.Collection("Users").FindOne(context.Background(), bson.M{"username": username}).Decode(&object)
	if err != nil {
    	return nil, err
    }
	if object == (models.Model_User{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
