package main

import (
	"net/http"
	"encoding/json"

)

type Chirp struct {
	Body  string `json:"body"`
}

func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, r *http.Request) {


	var chirp Chirp
	if err := json.NewDecoder(r.Body).Decode(&chirp); err != nil {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid chirp format"}`))
		return
	}
	if len(chirp.Body) == 0 || len(chirp.Body) > 140 {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Chirp must be between 1 and 140 characters"}`))
		return
	}
	// Here you would add any additional validation logic for the chirp
	// For example, checking for prohibited words or patterns		
   w.Header().Set("Content-Type", "application/json; charset=utf-8")
   w.WriteHeader(http.StatusOK)
   json.NewEncoder(w).Encode(struct {
	   Valid bool `json:"valid"`
   }{
	   Valid: true,
   })

}
