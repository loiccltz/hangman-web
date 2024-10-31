package main

import (
	"fmt"
	"net/http"
	"html/template"
	 hc "hangmanweb/hangman-classic/functions"
)



func home(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/home.html"
	template.ParseFiles(fileName)
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Erreur pendant le parsing", err)
		return
	}

	data := GameState {
		Lives : 10,
		Word : hc.GetRandomWord(),
	}

	t.Execute(w, data)
}

func play(w http.ResponseWriter, r *http.Request) {
	var fileName = "../templates/play.html"
	template.ParseFiles(fileName)
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