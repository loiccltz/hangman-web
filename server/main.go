package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

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

	//	Blank := gs.GetBlanks()

	// on espace les underscores sinon tout est collé
	//	BlankString := strings.Join(strings.Split(string(Blank), ""), " ")

	data := gs.GameState{
		Lives: gs.GetLives(),
		Word:  gs.GetWord(),
		// je fais passer en string sinon cela affiche les runes ( ASCII )
		BlanksDisplay: gs.GetBlanksDisplay(),
	}

	// debug
	println(data.Word)
	println(string(data.Blanks))

	t.Execute(w, data)
	return
}

func lose(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/lose.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}
	t.Execute(w, nil)
}

func win(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/win.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}
	t.Execute(w, nil)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// rendu si method est get
		play(w, r)

	} else if r.Method == http.MethodPost {
		fmt.Println("Méthode bien passé")
		gs.HandleGuess(w, r) // quand le joueur rend un input

		// Debug
		fmt.Println("Mot à trouver :", gs.GetWord())
		fmt.Println("Blanks affichés :", gs.GetBlanksDisplay())
		fmt.Println("Blanks nettoyés :", strings.ReplaceAll(gs.GetBlanksDisplay(), " ", ""))
		fmt.Println("Vies restantes :", gs.GetLives())

		if gs.GetLives() <= 0 {
			// Redirection vers la page de "lose"
			Red("/lose", w, r)
			return
		}

		if strings.ReplaceAll(gs.GetBlanksDisplay(), " ", "") == gs.GetWord() {
			fmt.Println("Mot trouvé avec succès !") // Debug
			fmt.Printf("Comparaison: '%s' == '%s'\n", strings.ReplaceAll(gs.GetBlanksDisplay(), " ", ""), gs.GetWord())

			// Redirection vers la page "win"
			Red("/win", w, r)
		}

	} else {
		http.Error(w, "Méthode non supportée", http.StatusMethodNotAllowed)
	}
}

func Red(s string, w http.ResponseWriter, r *http.Request) int {
	log.Println("Méthode:", r.Method, "URL:", r.URL.Path, "Statut:", http.StatusOK)
	http.Redirect(w, r, s, http.StatusSeeOther)
	return 0
}

func main() {
	fs := http.FileServer(http.Dir("../assets/"))

	http.HandleFunc("/", home)
	http.HandleFunc("/play", playHandler) // création d'une autre fonction car on ne peut pas avoir handleguess et handler en meme temps
	http.HandleFunc("/lose", lose)
	http.HandleFunc("/win", win)

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Serveur démarré sur le port 8080")
	http.ListenAndServe(":8080", nil)

}
