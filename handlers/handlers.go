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

	router.HandleFunc("/subir_avatar", middlew.ChequeoDB(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtener_avatar", middlew.ChequeoDB(middlew.ValidoJWT(routers.ObtenerAvatar))).Methods("GET")
	router.HandleFunc("/subir_banner", middlew.ChequeoDB(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtener_banner", middlew.ChequeoDB(middlew.ValidoJWT(routers.ObtenerBanner))).Methods("GET")

	router.HandleFunc("/alta_relacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/baja_relacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consulta_relacion", middlew.ChequeoDB(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/lista_usuarios", middlew.ChequeoDB(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leo_tweets_seguidores", middlew.ChequeoDB(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) // permiso para que cualquiera pueda acceder (se puede limitar por ip, etc)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
