package main

import (
	"log"
	"net/http"

	"github.com/junolabsmobile/meaning-of-songs/internal/application"
	httpinfra "github.com/junolabsmobile/meaning-of-songs/internal/infrastructure/http"
	"github.com/junolabsmobile/meaning-of-songs/internal/infrastructure/repository/memory"
)

func main() {
	// Crear adaptador de salida (repositorio in-memory)
	songRepo := memory.NewSongRepository()

	// Crear servicio de aplicación inyectando el puerto
	songService := application.NewSongService(songRepo)

	// Crear router HTTP (adaptador de entrada)
	router := httpinfra.NewRouter(songService)

	// Arrancar servidor
	addr := ":8080"
	log.Printf("Servidor iniciado en http://localhost%s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
