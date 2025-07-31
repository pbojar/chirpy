package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pbojar/chirpy/internal/utils"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func handlerValidateChirp(w http.ResponseWriter, req *http.Request) {
	type chirp struct {
		Body string `json:"body"`
	}
	type validResponse struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(req.Body)
	chrp := chirp{}
	err := decoder.Decode(&chrp)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode chirp", err)
		return
	}

	const maxChirpLength = 140
	if len(chrp.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	profanity := []string{"kerfuffle", "sharbert", "fornax"}
	cleaned := utils.CleanChirp(chrp.Body, profanity)

	respondWithJSON(w, http.StatusOK, validResponse{
		CleanedBody: cleaned,
	})
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, req *http.Request) {
	type createUserReq struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	params := createUserReq{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode create user request", err)
		return
	}

	dbUser, err := cfg.dbQueries.CreateUser(req.Context(), params.Email)
	user := User(dbUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't create user", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}
