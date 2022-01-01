package bd

import (
	"github.com/tomasbudassi/twittorGO/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {

	user, encontrado, _ := CheckYaExisteUsuario(email)
	if !encontrado {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
