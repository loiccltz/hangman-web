package state

import(
	"net/http"
	"strings"
	"encoding/json"
	"fmt"


)

type GuessRequest struct {
	Letter string `json:"letter"`
}


   
type GameStateResponse struct {
	Blanks string `json:"blanks"`
    Lives  int    `json:"lives"`
  }

  
  func HandleGuess(w http.ResponseWriter, r *http.Request) {
	  
	  if r.Method == http.MethodPost {
		  var guess GuessRequest

		  
  
		  err := json.NewDecoder(r.Body).Decode(&guess)
		  if err != nil {
			  http.Error(w, "Invalid JSON", http.StatusBadRequest)
			  return
		  }

		  if strings.Contains(state.Word, guess.Letter) {
			  UpdateBlanks(guess.Letter)
		  } else {

			  state.Lives--

		  }
  
		 
		  response := GameStateResponse{
			  Lives: state.Lives,
			  Blanks: GetBlanksDisplay(),
		  }
  
		 
		  fmt.Println("Lettre devinée:", guess.Letter)
		  w.Header().Set("Content-Type", "application/json")
		  
		  json.NewEncoder(w).Encode(response)
  
	  } else {
		  http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	  }
  }
  

