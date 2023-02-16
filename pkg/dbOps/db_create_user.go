package dbOps

import (
    "golang.org/x/net/context"
	"classifie/pkg/connections"
	"classifie/pkg/models"
)

func CreateUser (object *models.Model_User) (*models.Model_User, error) {

	_, err := connections.DATABASE.Collection("Users").InsertOne(context.Background(), object)
	if err != nil {
		return nil, err
	}
	return object, nil
}