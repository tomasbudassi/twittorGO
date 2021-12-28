package main

import (
	"github.com/tomasbudassi/twittorGO/bd"
	"github.com/tomasbudassi/twittorGO/handlers"
	"log"
)

func main() {
	if bd.CheckConexion() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Manejadores()
}
