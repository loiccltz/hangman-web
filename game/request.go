package state

import(
	"net/http"
	"strings"
	"encoding/json"


)

type GuessRequest struct {
	Letter string `json:"letter"`
}


    // Préparation de la réponse à renvoyer au client
type GameStateResponse struct {
	Blanks string `json:"blanks"`
    Lives  int    `json:"lives"`
  }

  
  func HandleGuess(w http.ResponseWriter, r *http.Request) {
	  
	  if r.Method == http.MethodPost {
		  var guess GuessRequest
  
		  err := json.NewDecoder(r.Body).Decode(&guess)
		  if err != nil {
			  http.Error(w, "Erreur de décodage", http.StatusBadRequest)
			  return
		  }
  

		  if strings.Contains(state.Word, guess.Letter) {
			  UpdateBlanks(guess.Letter)
		  } else {

			  state.Lives--

		  }
  
		  // Préparer la réponse à envoyer au client
		  response := GameStateResponse{
			  Lives: state.Lives,
			  Blanks: GetBlanksDisplay(),
		  }
  
		  // Répondre avec les informations mises à jour
		  w.Header().Set("Content-Type", "application/json")
		  json.NewEncoder(w).Encode(response)
  
	  } else {
		  http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	  }
  }
  

