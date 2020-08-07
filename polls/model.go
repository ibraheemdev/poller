package polls

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ibraheemdev/poller/config/db"
	"github.com/ibraheemdev/poller/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Poll Document
type Poll struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title"`
	Password uuid.UUID          `json:"password" bson:"password"`
}

// PollParams : Valid poll params
type PollParams struct {
	Title string `json:"title" bson:"title"`
}

func validate(poll *PollParams) []error {
	v := &validator.Validator{}
	v.ValidatePresenceOf("Title", poll.Title)
	return v.Errors
}

func createPoll(poll *PollParams) (string, []error) {
	errs := validate(poll)
	if errs != nil {
		return "", errs
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pwd, err := uuid.NewRandom()
	if err != nil {
		log.Println(err.Error())
		return "", []error{err}
	}
	create := bson.M{"title": poll.Title, "password": pwd}
	res, err := db.DB.Collection("polls").InsertOne(ctx, create)
	if err != nil {
		log.Println(err.Error())
		return "", []error{err}
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func updatePoll(id string, poll *PollParams) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"title", poll.Title}}}}
	_, err := db.DB.Collection("polls").UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
