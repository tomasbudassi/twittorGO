package routers

import (
	"encoding/json"
	"github.com/tomasbudassi/twittorGO/bd"
	"github.com/tomasbudassi/twittorGO/jwt"
	"github.com/tomasbudassi/twittorGO/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	var user models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalida"+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
	}

	documento, existe := bd.IntentoLogin(user.Email, user.Password)
	if !existe {
		http.Error(w, "Usuario y/o contraseña invalida", 400)
		return
	}

	jwtKey, err := jwt.GenerarJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token de usuario"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Set cookie al user
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
