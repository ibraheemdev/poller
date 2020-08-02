package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ibraheemdev/poller/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB :
var DB *mongo.Database

// Connect :
func Connect(config config.DatabaseConfig) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", config.Host, config.Port)))
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database(config.Name)
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
