package common

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ApplicantCollection *mongo.Collection
	Ctx                 = context.TODO()
)

func InitMongoDB() {
	client := ConnectMongoDb()
	database := client.Database("job-app")
	ApplicantCollection = database.Collection("applicant")
}

func ConnectMongoDb() *mongo.Client {
	uri := "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged. mongodb")
	return client
}
