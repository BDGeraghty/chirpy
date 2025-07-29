package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

const badWordMask string = "****"

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"valid"`
	}
	

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	if params.Body == "" {
		respondWithError(w, http.StatusBadRequest, "Chirp is empty", nil)
		return
	}
	params.Body = containsBadWord(params.Body)
	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: params.Body,
	})
}


func containsBadWord(chirp string) string {
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	var cleanChirp strings.Builder
	badWordFound := false
	words := strings.Fields(strings.ToLower(chirp))
	for _, word := range words {
		for _, bad := range badWords {
			if word == bad {
				badWordFound = true
				break
			} 
		}
		if badWordFound {
			cleanChirp.WriteString(badWordMask + " ")
		} else {
			cleanChirp.WriteString(word + " ")
		}
	}
	return cleanChirp.String()
}
