package dbOps

import (
    "golang.org/x/net/context"
	"classifie/pkg/connections"
	"classifie/pkg/models"
)

func CreateCustomNotes (object *models.Model_CustomNotes) (*models.Model_CustomNotes, error) {

	_, err := connections.DATABASE.Collection("CustomNotess").InsertOne(context.Background(), object)
	if err != nil {
		return nil, err
	}
	return object, nil
}