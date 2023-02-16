package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func FindTimeStampData (objectid string) (*models.Model_TimeStampData, error) {
	var object models.Model_TimeStampData

	err := connections.DATABASE.Collection("TimeStampDatas").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, err
    }
	if object == (models.Model_TimeStampData{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
