package dbOps

import (
	"classifie/pkg/connections"
    "classifie/pkg/models"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

func UpdateUser (object *models.Model_User) (*models.Model_User, error) {

	result, err := connections.DATABASE.Collection("Users").UpdateOne(context.Background(), bson.M{"objectid": object.ObjectID}, bson.M{"$set": object})
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return nil, errors.New("Specified ID not found!")
	}

	return object, nil
}