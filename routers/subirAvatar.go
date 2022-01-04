package routers

import (
	"github.com/tomasbudassi/twittorGO/bd"
	"github.com/tomasbudassi/twittorGO/models"
	"io"
	"net/http"
	"os"
	"strings"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	extension := strings.Split(handler.Filename, ".")[1]
	ruta := "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(ruta, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen"+err.Error(), 400)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen"+err.Error(), 400)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el avatar en la DB"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
