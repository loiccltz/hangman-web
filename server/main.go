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

func lose(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/lose.html"
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
	}

}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
        // rendu si méthode est GET
        play(w, r)
    } else if r.Method == http.MethodPost {
        fmt.Println("Méthode bien passée")
        gs.HandleGuess(w, r) // quand le joueur rend un input
        
        // Vérifier si le nombre de vies est 0 après avoir traité l'entrée
        if gs.GetLives() == 0 {
            // Rediriger vers la page de défaite
            http.Redirect(w, r, "/lose", http.StatusSeeOther)
            return
        }

        // Si le nombre de vies est toujours supérieur à 0, on continue normalement
        // (pas besoin de renvoyer d'en-têtes à ce stade)
    } else {
         http.Error(w, "Méthode non supportée", http.StatusMethodNotAllowed)
     }
 }


func main() {
	fs := http.FileServer(http.Dir("../assets/"))
	
	http.HandleFunc("/", handler)
	http.HandleFunc("/play", playHandler) // création d'une autre fonction car on ne peut pas avoir handleguess et handler en meme temps
	http.HandleFunc("/lose", lose)

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe("", nil)
	
}