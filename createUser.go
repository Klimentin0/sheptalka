package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Klimentin0/sheptalka/internal/database"
)

type Params struct {
	Email string `json:"email"`
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	// ??
	var email sql.NullString
	if params.Email != "" {
		email = sql.NullString{String: params.Email, Valid: true}
	} else {
		email = sql.NullString{Valid: false}
	}
	user, err := cfg.db.CreateUser(r.Context(), email)
	if err != nil {
		log.Printf("Error creating user %s", err)
	}
	mappedUser, err := json.Marshal(mapDbUser(user))
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(mappedUser)
}

func mapDbUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
		Email:     dbUser.Email.String,
	}
}
