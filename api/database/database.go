package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDBURI() string {

	host := os.Getenv("OBJECT_DB_STORE_ADDR")
	uri := fmt.Sprintf("mongodb://%s", host)

	return uri

}

func ConnectToDB() *mongo.Client {

	uri := GetDBURI()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Println("db ping successful!")
	}

	return client

}

// this method inserts a single document in the database
func InsertDocument(document any) error {

	db := ConnectToDB()

	dbName := os.Getenv("OBJECT_DB_NAME")

	urlCollection := db.Database(dbName).Collection("urls")

	_, err := urlCollection.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}

	return nil

}

func FindDocument(filter bson.D, limit int) (*mongo.Cursor, error) {

	client := ConnectToDB()

	dbName := os.Getenv("OBJECT_DB_NAME")

	urlCollection := client.Database(dbName).Collection("urls")

	findOptions := options.FindOptions{}

	if limit != -1 {
		findOptions.SetLimit(int64(limit))
	}

	cursor, err := urlCollection.Find(context.TODO(), filter, &findOptions)
	if err != nil {
		log.Default().Println("(FindDocument) err in urlCollection.Find:", err)
		return nil, err
	}

	return cursor, nil

}
