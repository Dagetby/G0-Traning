package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var c *mongo.Client

type Hero struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Signed bool   `json:"signed"`
}

func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*Hero {
	var heroes []*Hero

	collections := client.Database("db_1").Collection("users")
	cur, err := collections.Find(context.TODO(), filter)
	check(err, "Error on Finding all the documents")

	for cur.Next(context.TODO()) {
		var hero Hero
		err = cur.Decode(&hero)
		check(err, "Error on decoding the document")
		heroes = append(heroes, &hero)
	}
	return heroes

}

func Connections() *mongo.Client {
	c = GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	check(err, "")
	return c
}
func GetClient() *mongo.Client {
	fmt.Println("Start the application...")

	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")

	client, err := mongo.NewClient(clientOptions)
	check(err, "")
	err = client.Connect(context.Background())
	check(err, "")

	fmt.Println("Connected to Mongodb")

	return client
}

func RemoveOneHero(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database("civilact").Collection("heroes")
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting one Hero", err)
	}
	return deleteResult.DeletedCount
}
