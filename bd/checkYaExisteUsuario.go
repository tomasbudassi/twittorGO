package bd

import (
	"context"
	"github.com/tomasbudassi/twittorGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func CheckYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	condition := bson.M{"email": email}
	var resultado models.Usuario

	err := collection.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
