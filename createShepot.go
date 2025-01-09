package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Klimentin0/sheptalka/internal/database"
	"github.com/google/uuid"
)

type Shepot struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	User_id   uuid.UUID `json:"user_id"`
}

type shepotParams struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

type ValidResp struct {
	Cleaned string `json:"cleaned_body"`
}

func (cfg *apiConfig) shepotHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	shepotParams := shepotParams{}
	err := decoder.Decode(&shepotParams)
	if err != nil {
		respBody := ErrorResp{
			Error: "Error decoding json",
		}
		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(dat)
		return
	}

	if len(shepotParams.Body) > 140 {
		respBody := ErrorResp{
			Error: "Слишком длинное сообщение",
		}
		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dat)
		return
	}

	cleanedBody := badWordCleaner(shepotParams.Body)
	sqlParams := database.CreateShepotParams{
		Body:   cleanedBody,
		UserID: shepotParams.UserID,
	}

	dbShepot, err := cfg.db.CreateShepot(r.Context(), sqlParams)
	if err != nil {
		log.Printf("Error creating user %s", err)
	}

	shepot := mapDbShepot(dbShepot)

	dat, err := json.Marshal(shepot)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(dat)
}

func badWordCleaner(body string) string {
	split := strings.Split(body, " ")
	//Типо плоxие слова
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	var cleanedSplit []string
	for i := range split {
		for j := range badWords {
			if strings.ToLower(split[i]) == badWords[j] {
				split[i] = "****"
			}
		}
		cleanedSplit = append(cleanedSplit, split[i])
	}
	return strings.Join(cleanedSplit, " ")
}

func mapDbShepot(dbShepot database.Shepot) Shepot {
	return Shepot{
		ID:        dbShepot.ID,
		CreatedAt: dbShepot.CreatedAt,
		UpdatedAt: dbShepot.UpdatedAt,
		Body:      dbShepot.Body,
		User_id:   dbShepot.UserID,
	}
}
