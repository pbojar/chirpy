package main

import (
	"fmt"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, cfg.fileserverHits.Load())
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Here there be dragons!", nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	err := cfg.dbQueries.DeleteUsers(req.Context())
	if err != nil {
		log.Fatalf("Error deleting users: %v", err)
	}
	cfg.fileserverHits.Store(0)
	fmt.Fprintf(w, "Reset hits to %d", cfg.fileserverHits.Load())
}
