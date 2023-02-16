package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func FindSummery (objectid string) (*models.Model_Summery, error) {
	var object models.Model_Summery

	err := connections.DATABASE.Collection("Summerys").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, err
    }
	if object == (models.Model_Summery{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
