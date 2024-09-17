package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectToDatastore() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: "",
		DB:       0,
	})

	status := client.Ping(context.TODO())
	if status.Val() != "PONG" {
		log.Default().Println("datastore is not running...")
	}

	return client

}

func SetKey(key, value string) {

	client := ConnectToDatastore()

	err := client.Set(context.TODO(), key, value, 0).Err()
	if err != nil {
		log.Default().Println("Failed to set value in cache", err.Error())
		return
	}

	log.Default().Println("Value written in cache")

}

func GetKey(key string) (string, bool) {

	client := ConnectToDatastore()

	value, err := client.Get(context.TODO(), key).Result()
	if err != nil {
		if err.Error() != redis.Nil.Error() {
			log.Default().Println("failed to get the key from cache:", err.Error())
		}
		return "", false
	}

	return value, true

}
