package bd

import (
	"context"
	"github.com/tomasbudassi/twittorGO/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertarRegistro(user models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancela el contexto de withTimeout

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	user.Password, _ = EncriptarPassword(user.Password)
	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
