package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
    "golang.org/x/net/context"
)

func FindallCustomNotes () (*[]models.Model_CustomNotes, error) {
	var objects []models.Model_CustomNotes
	results, err := connections.DATABASE.Collection("CustomNotess").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for results.Next(context.Background()) {
		var object models.Model_CustomNotes
		if err = results.Decode(&object); err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
