package base

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Model : A record in a ModelCollection
type Model interface {
	Collection() ModelCollection
}

// ModelCollection : A collection of models in the database
type ModelCollection struct {
	*mongo.Collection
}

// Find : Finds the record with a particular ID
func (c ModelCollection) Find(id string, model Model) error {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", ObjectID}}
	err = c.FindOne(ctx, filter).Decode(model)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
