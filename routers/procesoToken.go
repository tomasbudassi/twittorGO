package routers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/tomasbudassi/twittorGO/bd"
	"github.com/tomasbudassi/twittorGO/models"
	"strings"
)

var Email string
var IDUsuario string

// Procesar token para extraer sus valores
func ProcesoToken(token string) (*models.Claim, bool, string, error) {

	miClave := []byte("CursoGO_Twittor")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	// Elimino el string "Bearer", queda el token limpio
	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.CheckYaExisteUsuario(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
