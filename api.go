package main

import (
	"encoding/json"
	"log"
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
	type validateResponse struct {
		Error       string `json:"error"`
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(req.Body)
	chrp := chirp{}
	err := decoder.Decode(&chrp)
	if err != nil {
		log.Printf("Error decoding chirp: %s", err)
		w.WriteHeader(500)
		return
	}

	resp := validateResponse{}
	if len(chrp.Body) > 140 {
		w.WriteHeader(400)
		resp.Error = "Chirp is too long"
	} else {
		w.WriteHeader(200)
		resp.CleanedBody = utils.CleanChirp(chrp.Body)
	}

	dat, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dat)
}
