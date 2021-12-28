package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tomasbudassi/twittorGO/middlew"
	"github.com/tomasbudassi/twittorGO/routers"
	"log"
	"net/http"
	"os"
)

func Manejadores() {
	router := mux.NewRouter()

	// Rutas
	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) // permiso para que cualquiera pueda acceder (se puede limitar por ip, etc)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
