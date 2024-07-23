package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {

	//Crear enrutador
	router := http.NewServeMux()

	// Manejador para archivos estáticos
	fs := http.FileServer(http.Dir("static"))

	//Ruta para acceder a los archivos estáticos
	router.Handle("/static/", http.StripPrefix("/static", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	PORT := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s", PORT)

	http.ListenAndServe(PORT, router)

}
