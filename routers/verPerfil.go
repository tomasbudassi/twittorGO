package routers

import (
	"encoding/json"
	"github.com/tomasbudassi/twittorGO/bd"
	"net/http"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar busca el registro"+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)
}
