package dbOps

import (
    "golang.org/x/net/context"
	"classifie/pkg/connections"
	"classifie/pkg/models"
)

func CreateTimeStampData (object *models.Model_TimeStampData) (*models.Model_TimeStampData, error) {

	_, err := connections.DATABASE.Collection("TimeStampDatas").InsertOne(context.Background(), object)
	if err != nil {
		return nil, err
	}
	return object, nil
}