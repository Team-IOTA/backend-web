package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func FindCustomNotes (objectid string) (*models.Model_CustomNotes, error) {
	var object models.Model_CustomNotes

	err := connections.DATABASE.Collection("CustomNotess").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, err
    }
	if object == (models.Model_CustomNotes{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
