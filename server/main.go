package main

import (
	"fmt"
	"html/template"
	"net/http"

	// hc "hangmanweb/hangman-classic/functions"
	gs "hangmanweb/game"
)

func home(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/home.html"

	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}

	t.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/play.html"

	difficulty := r.URL.Query().Get("difficulty")

	gs.InitGame(difficulty) // charge la partie avec la difficulté récup grace a "r.URL.Query().Get"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}

	// initialisation des variables
	data := gs.GameState{
		Lives: gs.GetLives(),
		Word:  gs.GetWord(),
		// je fais passer en string sinon cela affiche les runes ( ASCII )
		BlanksDisplay: gs.GetBlanksDisplay(), // on espace les underscores sinon tout est collé
	}

	// debug
	println(data.Word)
	println(string(data.Blanks))

	t.Execute(w, data)
}

func lose(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/lose.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}

	data := gs.GameState{ // affiche le mot qui été a trouver
		Word: gs.GetWord(),
	}
	t.Execute(w, data)
}

func win(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/win.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}
	data := gs.GameState{ // affiche le mot qui été a trouver
		Word: gs.GetWord(),
	}
	t.Execute(w, data)
}

func handler(w http.ResponseWriter, r *http.Request) { // Redirect de fonctions
	switch r.URL.Path {
	case "/":
		home(w, r)
	case "/lose":
		lose(w, r)
	}

}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// rendu si method est get
		play(w, r)
	} else if r.Method == http.MethodPost {
		fmt.Println("Méthode bien passé")
		gs.HandleGuess(w, r) // quand le joueur rend un input
		if gs.GetLives() <= 0 {
			// Redirection vers la page de "lose"
			http.Redirect(w, r, "/lose", http.StatusSeeOther)
			return
		}
	} else {
		http.Error(w, "Méthode non supportée", http.StatusMethodNotAllowed)
	}
}

func main() {
	fs := http.FileServer(http.Dir("../assets/"))

	// Creation de nom pour appeler les fonctions
	http.HandleFunc("/", handler)
	http.HandleFunc("/play", playHandler) // création d'une autre fonction car on ne peut pas avoir handleguess et handler en meme temps
	http.HandleFunc("/lose", handler)
	http.HandleFunc("/win", win)

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":8080", nil)

}
