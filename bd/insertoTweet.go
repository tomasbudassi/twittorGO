package bd

import (
	"context"
	"github.com/tomasbudassi/twittorGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertoTweet(tweet models.GraboTweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("tweet")

	registro := bson.M{
		"userId":  tweet.UserID,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}
	result, err := collection.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
