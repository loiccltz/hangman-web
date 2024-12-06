package state

import(
	"net/http"
	"strings"
	"encoding/json"
	"fmt"


)

type GameStateResponse struct {
	Blanks string `json:"blanks"`
    Lives  int    `json:"lives"`
	HangmanArt string `json:"hangmanArt"`
	GameOver    bool   `json:"gameOver"`
  }

  type GuessRequest struct {
	Letter string `json:"letter"`
}

  
  func HandleGuess(w http.ResponseWriter, r *http.Request) {

	  
	  if r.Method == http.MethodPost{
		  
		var guess GuessRequest

		err := json.NewDecoder(r.Body).Decode(&guess) // décode le json, c'est pour sa cela ne marchait paaaaas
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			fmt.Println("Erreur de décodage JSON :", err)
			return
		}

		  if strings.Contains(state.Word, guess.Letter) {
			  UpdateBlanks(guess.Letter)
		  } else {
			  state.Lives--
		  }

			maxLives := 10
			stage := maxLives - state.Lives
			hangmanArt := showHangman(stage)

  
		 
		  response := GameStateResponse{
			  Lives: state.Lives,
			  Blanks: GetBlanksDisplay(),
			  HangmanArt: hangmanArt,
			  GameOver: GetLives() <= 0,
		  }
  
		 
		  fmt.Println("Lettre devoné:", guess.Letter)
		  w.Header().Set("Content-Type", "application/json")
		  
		  json.NewEncoder(w).Encode(response)
  
	  } else {
		  http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	  }
  }
  

