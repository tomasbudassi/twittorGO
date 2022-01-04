package bd

import (
	"context"
	"github.com/tomasbudassi/twittorGO/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func ModificoRegistro(user models.Usuario, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(user.Nombre) > 0 {
		registro["nombre"] = user.Nombre
	}
	if len(user.Apellidos) > 0 {
		registro["apellidos"] = user.Apellidos
	}
	registro["fechaNacimiento"] = user.FechaNacimiento
	if len(user.Avatar) > 0 {
		registro["avatars"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		registro["banners"] = user.Banner
	}
	if len(user.Biografia) > 0 {
		registro["biografia"] = user.Biografia
	}
	if len(user.Ubicacion) > 0 {
		registro["ubicacion"] = user.Ubicacion
	}
	if len(user.SitioWeb) > 0 {
		registro["sitioWeb"] = user.SitioWeb
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := collection.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
