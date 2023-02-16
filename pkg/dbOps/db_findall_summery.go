package dbOps

import (
	"classifie/pkg/connections"
	"classifie/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
    "golang.org/x/net/context"
)

func FindallSummery () (*[]models.Model_Summery, error) {
	var objects []models.Model_Summery
	results, err := connections.DATABASE.Collection("Summerys").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for results.Next(context.Background()) {
		var object models.Model_Summery
		if err = results.Decode(&object); err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
