package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
    "golang.org/x/net/context"
)

func FindallUser () (*[]models.Model_User, error) {
	var objects []models.Model_User
	results, err := connections.DATABASE.Collection("Users").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for results.Next(context.Background()) {
		var object models.Model_User
		if err = results.Decode(&object); err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
