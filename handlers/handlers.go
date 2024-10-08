package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"rpsweb/logic"
	"strconv"
)

type Player struct{
	Name string
}

var player Player

const(
	templateDir = "templates/"
	templateBase = templateDir + "base.html"
)

// Funcion para enrutar
func Index(w http.ResponseWriter, r *http.Request) {
	restarValue()
	renderTemplate(w, "index.html", nil)
	
}

func NewGame(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "newGame.html", nil)

}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST"{
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		player.Name = r.Form.Get("name")
	}

	if player.Name == ""{
		http.Redirect(w,r, "/new", http.StatusFound)
	}

	fmt.Println(player.Name)
	renderTemplate(w,"game.html", player)

}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice,_ := strconv.Atoi( r.URL.Query().Get("c"))
	result := logic.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result,"", "  ")

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	fmt.Println(result)

	//fmt.Fprintln(w, "Jugar")

}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)

}


func renderTemplate(w http.ResponseWriter, page string, data any){

	//Must evitar errores de carga y si hay un error ejecuta un panic
	//Con templates se puede renderizar las plantillas
	tpl := template.Must(template.ParseFiles( templateBase, templateDir + page))
	
	err := tpl.ExecuteTemplate(w, "base", data)

	if err != nil {
		http.Error(w, "No se ha podido cargar", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	
}

//Reinicir valores
func restarValue(){
	player.Name = ""
	logic.ComputerScore = 0
	logic.PlayerScore = 0
}
