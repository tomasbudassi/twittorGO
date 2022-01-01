package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tomasbudassi/twittorGO/models"
	"time"
)

func GenerarJWT(user models.Usuario) (string, error) {

	miClave := []byte("CursoGO_Twittor")
	payload := jwt.MapClaims{
		"email":            user.Email,
		"nombre":           user.Nombre,
		"apellidos":        user.Apellidos,
		"fecha_nacimiento": user.FechaNacimiento,
		"biografia":        user.Biografia,
		"ubicacion":        user.Ubicacion,
		"sitio_web":        user.SitioWeb,
		"_id":              user.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
