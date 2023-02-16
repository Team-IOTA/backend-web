package dbOps

import (
	"classifie/pkg/connections"
    "classifie/pkg/models"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

func UpdateTimeStampData (object *models.Model_TimeStampData) (*models.Model_TimeStampData, error) {

	result, err := connections.DATABASE.Collection("TimeStampDatas").UpdateOne(context.Background(), bson.M{"objectid": object.ObjectID}, bson.M{"$set": object})
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return nil, errors.New("Specified ID not found!")
	}

	return object, nil
}