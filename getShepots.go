package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) getShepotsHandler(w http.ResponseWriter, r *http.Request) {
	shepots, err := cfg.db.GetAllShepots(r.Context())
	if err != nil {
		log.Printf("Error creating user %s", err)
	}
	dat, err := json.Marshal(shepots)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
