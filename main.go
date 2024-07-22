package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {

	//Crear enrutador
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	PORT := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s", PORT)

	http.ListenAndServe(PORT, router)

}
