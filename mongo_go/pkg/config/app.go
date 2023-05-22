package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Replace the placeholder with your Atlas connection string
const uri = "mongodb+srv://Lokesh:lokesh@cluster0.bablpxy.mongodb.net/?retryWrites=true&w=majority"

// var db

func Connect() *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	usersCollection := client.Database("mongo").Collection("user")
	// user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// // insert the bson object using InsertOne()
	// result, err := usersCollection.InsertOne(context.TODO(), user)
	// fmt.Println(result)
	return usersCollection
}
