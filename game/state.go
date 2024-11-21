package state

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

type GameState struct {
	Lives int 
	Word  string
	Blanks []rune
	Blank string
}

var state GameState

// Initialisation du jeu
func InitGame() {
	state.Lives = 10
	state.Word = GetRandomWord()
	state.Blanks = make([]rune, len(state.Word))
	for i := range state.Blanks {
		state.Blanks[i] = '_'
	}
}

// Fonction pour obtenir un mot aléatoire
func GetRandomWord() string {
	file, err := os.Open("../hangman-classic/dictionnaries/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words[rand.Intn(len(words))]
}

// Getters 
func GetLives() int {
	return state.Lives
}

func GetWord() string {
	return state.Word
}

func GetBlanks() []rune {
	return state.Blanks
}


// Setters
func SetLives(lives int) {
	state.Lives = lives
}

func SetWord(word string) {
	state.Word = word
}

func SetBlanks(blanks []rune) {
	state.Blanks = blanks
}
