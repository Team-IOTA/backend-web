package dbOps

import (
    "golang.org/x/net/context"
	"classifie/pkg/connections"
	"classifie/pkg/models"
)

func CreateSummery (object *models.Model_Summery) (*models.Model_Summery, error) {

	_, err := connections.DATABASE.Collection("Summerys").InsertOne(context.Background(), object)
	if err != nil {
		return nil, err
	}
	return object, nil
}