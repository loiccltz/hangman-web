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
	gs.InitGame()
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}

	t.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/play.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}

	// Vérifie la victoire avant d'exécuter la page
	if gs.GetBlanksDisplay() == "" { // Suppose que la fonction retourne une chaîne vide si aucun blanc n'existe
		http.Redirect(w, r, "/win", http.StatusSeeOther)
		return
	}

//	Blank := gs.GetBlanks()

	// on espace les underscores sinon tout est collé
//	BlankString := strings.Join(strings.Split(string(Blank), ""), " ")

	

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

func win(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/style.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}
	t.Execute(w, nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path {
	case "/":
		home(w,r)
	default:
		http.NotFound(w, r)
	}
}


func playHandler(w http.ResponseWriter, r *http.Request) {
     if r.Method == http.MethodGet {
         // rendu si method est get
         play(w,r)
     } else if r.Method == http.MethodPost{
		fmt.Println("Méthode bien passé")
         gs.HandleGuess(w,r) // quand le joueur rend un input
		 http.Redirect(w, r, "/play", http.StatusSeeOther)
    } else {
         http.Error(w, "Méthode non supportée", http.StatusMethodNotAllowed)
     }
 }

 func winHandler(w http.ResponseWriter, r *http.Request) {
	win(w, r)
}


func main() {
	fs := http.FileServer(http.Dir("../assets/"))
	
	http.HandleFunc("/", handler)
	http.HandleFunc("/play", playHandler) // création d'une autre fonction car on ne peut pas avoir handleguess et handler en meme temps
	http.HandleFunc("/win", winHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":8080", nil)
	
}