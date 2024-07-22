package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)


const(
	templateDir = "templates/"
	templateBase = templateDir + "base.html"
)

// Funcion para enrutar
func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
	
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "newGame.html", nil)

}

func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w,"game.html", nil)

}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")

}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)

}


func renderTemplate(w http.ResponseWriter, page string, data any){

	//Must evitar errores de carga y si hay un error ejecuta un panic
	//Con templates se puede renderizar las plantillas
	tpl := template.Must(template.ParseFiles( templateBase, templateDir + page))
	
	err := tpl.ExecuteTemplate(w, "base", nil)

	if err != nil {
		http.Error(w, "No se ha podido cargar", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	
}
