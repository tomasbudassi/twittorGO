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
	router.HandleFunc("/login", middlew.ChequeoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/ver_perfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificar_perfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leo_tweets", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminar_tweet", middlew.ChequeoDB(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) // permiso para que cualquiera pueda acceder (se puede limitar por ip, etc)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
