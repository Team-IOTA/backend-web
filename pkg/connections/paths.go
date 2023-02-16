package connections

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database


const DATABASE_URL = "mongodb+srv://RidmaTP:ridma9999@cluster0.eq3p8.mongodb.net/?retryWrites=true&w=majority"
const DATABASE_NAME = "classifie"

