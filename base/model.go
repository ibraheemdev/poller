package base

import (
	"context"
	"log"
	"time"

	"github.com/ibraheemdev/poller/config/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Find : finds the record with a particular ID
func Find(id string, model interface{}) error {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", ObjectID}}
	err = db.DB.Collection("polls").FindOne(ctx, filter).Decode(model)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
