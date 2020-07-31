package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB :
var DB *mongo.Database

// Connect :
func Connect(dbConfig struct {
	Host string "yaml:\"host\""
	Port string "yaml:\"port\""
	Name string "yaml:\"name\""
}) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbConfig.Host+dbConfig.Port))
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database(dbConfig.Name)
	return client
}

// Disconnect :
func Disconnect(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
