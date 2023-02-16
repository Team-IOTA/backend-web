package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

func DeleteTimeStampData (objectid string) (*models.Model_TimeStampData, error) {

    result, err := connections.DATABASE.Collection("TimeStampDatas").DeleteOne(context.Background(), bson.M{"objectid": objectid})
    if err != nil {
    	return nil, err
    }
    if result.DeletedCount < 1 {
    	return nil, errors.New("Specified ID not found!")
    }
    return nil, nil
}