package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/BDGeraghty/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
		ExpiresInSeconds int64 `json:"expires_in_seconds"`
	}
	type response struct {
			ID        string    `json:"id"`
			Email     string    `json:"email"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Token     string    `json:"token"`
	}	

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	if params.ExpiresInSeconds <= 0 {
		params.ExpiresInSeconds = 3600 // Default to 1 hour if not specified
	}
	token, err := auth.MakeJWT(user.ID, cfg.secretKey, time.Duration(params.ExpiresInSeconds)*time.Second)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate JWT", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		ID:        user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	})
}
