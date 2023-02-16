package functions

import(
    "classifie/pkg/connections"
	"classifie/pkg/models"
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"errors"
)

func UniqueCheck_Summery (object models.Model_Summery)error{
	model := models.Model_Summery {}

    err := connections.DATABASE.Collection("Summerys").FindOne(context.Background(), bson.M{"$or" : []interface{} { bson.M{"summeryid" : object.SummeryId}, }}).Decode(&model)
	if err != nil {
		
		return nil
	}
	
	return errors.New(object.SummeryId + " " + "are Unique or already exists")    
}