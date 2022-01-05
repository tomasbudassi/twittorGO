package routers

import (
	"github.com/tomasbudassi/twittorGO/bd"
	"github.com/tomasbudassi/twittorGO/models"
	"net/http"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var relacion models.Relacion
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(relacion)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar la relacion"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado borrar la relacion"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
