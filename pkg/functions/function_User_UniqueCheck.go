package functions

import(
    "classifie/pkg/connections"
	"classifie/pkg/models"
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"errors"
)

func UniqueCheck_User (object models.Model_User)error{
	model := models.Model_User {}

    err := connections.DATABASE.Collection("Users").FindOne(context.Background(), bson.M{"$or" : []interface{} { bson.M{"userid" : object.UserId},bson.M{"username" : object.UserName},bson.M{"email" : object.Email}, }}).Decode(&model)
	if err != nil {
		
		return nil
	}
	
	return errors.New(object.UserId + " " + object.UserName + " " + object.Email + " " + "are Unique or already exists")    
}