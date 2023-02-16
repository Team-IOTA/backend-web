package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
    "golang.org/x/net/context"
)

func FindallTimeStampData () (*[]models.Model_TimeStampData, error) {
	var objects []models.Model_TimeStampData
	results, err := connections.DATABASE.Collection("TimeStampDatas").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for results.Next(context.Background()) {
		var object models.Model_TimeStampData
		if err = results.Decode(&object); err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
