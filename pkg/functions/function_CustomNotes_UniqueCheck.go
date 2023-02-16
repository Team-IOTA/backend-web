package functions

import(
    "classifie/pkg/connections"
	"classifie/pkg/models"
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"errors"
)

func UniqueCheck_CustomNotes (object models.Model_CustomNotes)error{
	model := models.Model_CustomNotes {}

    err := connections.DATABASE.Collection("CustomNotess").FindOne(context.Background(), bson.M{"$or" : []interface{} { bson.M{"noteid" : object.NoteId}, }}).Decode(&model)
	if err != nil {
		
		return nil
	}
	
	return errors.New(object.NoteId + " " + "are Unique or already exists")    
}