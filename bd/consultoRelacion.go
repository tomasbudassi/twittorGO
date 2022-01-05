package bd

import (
	"context"
	"fmt"
	"github.com/tomasbudassi/twittorGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func ConsultoRelacion(relacion models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("relacion")

	condicion := bson.M{
		"usuarioId":         relacion.UsuarioID,
		"usuarioRelacionID": relacion.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)

	err := collection.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
