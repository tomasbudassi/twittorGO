package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func BorroTweet(tweetID string, UserID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(tweetID)

	condicion := bson.M{
		"_id":    objID,
		"userId": UserID,
	}

	_, err := collection.DeleteOne(ctx, condicion)
	return err
}
