package state

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
)

type GameState struct {
	Lives         int
	Word          string
	Blanks        []rune
	BlanksDisplay string
}

var state GameState

// Initialisation du jeu
func InitGame(difficulty string) {
	state.Lives = 10
	state.Word = removeAccents(GetRandomWord(difficulty))
	state.Blanks = make([]rune, len(state.Word))
	for i := range state.Blanks {
		state.Blanks[i] = '_'
	}
}	


func removeAccents(Word string) string {
	replacer := strings.NewReplacer(
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"à", "a", "â", "a", "ä", "a",
		"î", "i", "ï", "i",
		"ô", "o", "ö", "o",
		"ù", "u", "û", "u", "ü", "u",
		"ç", "c",
	)
	return replacer.Replace(Word)
}


// fonction utilisé pour le hangman de base
func showHangman(stage int) string {
	hangmanFile, err := os.Open("../hangman-classic/dictionnaries/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer hangmanFile.Close()

	scanner := bufio.NewScanner(hangmanFile)
	linesPerStage := 8

	startLine := stage * linesPerStage // la ligne de début pour cette étape
	endLine := startLine + linesPerStage

	currentLine := 0
	hangmanArt := ""

	for scanner.Scan() {

		if currentLine < startLine {
			currentLine++
			continue
		}

		if currentLine >= endLine {
			break
		}

		hangmanArt += scanner.Text() + "\n"
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return hangmanArt
}

func GetRandomWord(difficulty string) string { // fonction utiliser pour choisir un mot au hazard

	var filePath string

	switch difficulty { // permet de switch en fonction de la difficulté
	case "easy":
		filePath = "../hangman-classic/dictionnaries/words.txt" // easy
	case "medium":
		filePath = "../hangman-classic/dictionnaries/wordsMedium.txt" // medium
	case "hard":
		filePath = "../hangman-classic/dictionnaries/wordsHard.txt" // hard
	}

	file, err := os.Open(filePath)
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

// permet de convertir les runes en string pour pouvoir les afficher sur le front
func GetBlanksDisplay() string {
	var result []string
	for _, r := range state.Blanks {
		result = append(result, string(r))
	}
	return strings.Join(result, " ")
}

func UpdateBlanks(letter string) { // Remplace les underscores par la lettre, si trouvé
	for i, char := range state.Word {
		if string(char) == letter {
			state.Blanks[i] = rune(letter[0]) 
		}
	}
}

func ContainsBlanks(blanks []rune) bool {
	for _, char := range blanks {
		if char == '_' {
			return true
		}
	}
	return false
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
