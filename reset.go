package main

import (
	"net/http"
	"os"
)

func (cfg *apiConfig) resetDb(w http.ResponseWriter, r *http.Request) {
	dbPlatform := os.Getenv("PLATFORM")

	if dbPlatform != "dev" {
		http.Error(w, "403 Forbidden", http.StatusForbidden)
		return
	}

	if err := cfg.db.DeleteAllUsers(r.Context()); err != nil {
		http.Error(w, "Failed to reset database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
