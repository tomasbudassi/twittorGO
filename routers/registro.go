package routers

import (
	"encoding/json"
	"github.com/tomasbudassi/twittorGO/bd"
	"github.com/tomasbudassi/twittorGO/models"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var user models.Usuario

	// El body unicamente puede ser leido una vez (es un objeto de tipo STREAM, se lee 1 vez y se destruye)
	err := json.NewDecoder(r.Body).Decode(&user) // almaceno el body en la variable user

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "La contraseña del usuario debe contener 6 caracteres como minimo", 400)
		return
	}

	_, userExiste, _ := bd.CheckYaExisteUsuario(user.Email)

	if userExiste {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertarRegistro(user)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el gistro de usuario"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
