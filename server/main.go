package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	// hc "hangmanweb/hangman-classic/functions"
	gs "hangmanweb/game"

)



func home(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/home.html"
	gs.InitGame()
	template.ParseFiles(fileName)
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}

	t.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/play.html"
	template.ParseFiles(fileName)
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}


	Blank := gs.GetBlanks()

	// on espace les underscores sinon tout est collé
	BlankString := strings.Join(strings.Split(string(Blank), ""), " ")

	

	data := gs.GameState {
		Lives : gs.GetLives(),
		Word : gs.GetWord(),
		// je fais passer en string sinon cela affiche les runes ( ASCII )
		BlanksDisplay : gs.GetBlanksDisplay(),

	}

	// debug
	println(data.Word)
	println(string(data.Blanks))

	t.Execute(w, data)
}

func handler(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
	case "/":
		home(w,r)
	case "/play":
		play(w,r)
	}
}

func main() {
	fs := http.FileServer(http.Dir("../assets/"))
	
	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe("", nil)
	
}