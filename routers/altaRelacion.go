package routers

import (
	"github.com/tomasbudassi/twittorGO/bd"
	"github.com/tomasbudassi/twittorGO/models"
	"net/http"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", 400)
		return
	}

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(relacion)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar la relacion"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la relacion"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
