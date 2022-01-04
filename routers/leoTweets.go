package routers

import (
	"encoding/json"
	"github.com/tomasbudassi/twittorGO/bd"
	"net/http"
	"strconv"
)

func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	pagina := r.URL.Query().Get("pagina")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", 400)
		return
	}
	if len(pagina) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", 400)
		return
	}

	numPagina, err := strconv.Atoi(pagina)
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a 0", 400)
		return
	}

	pag := int64(numPagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)

	if !correcto {
		http.Error(w, "Error al leer los tweets", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
