package main

import (
	"encoding/json"
	"net/http"

	"github.com/pbojar/chirpy/internal/utils"
)

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
