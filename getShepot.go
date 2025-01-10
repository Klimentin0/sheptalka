package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Klimentin0/sheptalka/internal/database"
)

func (cfg *apiConfig) getShepotHandler(w http.ResponseWriter, r *http.Request) {

	shepotID := r.PathValue("shepotID")

	shepots, err := cfg.db.GetAllShepots(r.Context())
	if err != nil {
		log.Printf("Error creating shepots slice %s", err)
	}

	shepot := getShopot(shepots, shepotID)

	dat, err := json.Marshal(shepot)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}

func getShopot(shepots []database.Shepot, shepotID string) database.Shepot {
	var singleShepot database.Shepot
	for i, shepot := range shepots {
		mappedShepot := mapDbShepot(shepot)
		if mappedShepot.ID.String() == shepotID {
			singleShepot = shepots[i]
		}
	}
	return singleShepot
}
