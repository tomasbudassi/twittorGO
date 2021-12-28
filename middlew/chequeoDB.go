package middlew

import (
	"github.com/tomasbudassi/twittorGO/bd"
	"net/http"
)

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConexion() == 0 {
			http.Error(w, "Conexion perdida con la DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
