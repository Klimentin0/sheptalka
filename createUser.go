package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Klimentin0/sheptalka/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type userParams struct {
	Email string `json:"email"`
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userParams := userParams{}
	err := decoder.Decode(&userParams)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	dbUser, err := cfg.db.CreateUser(r.Context(), userParams.Email)
	if err != nil {
		log.Printf("Error creating user %s", err)
	}
	mappedUser := mapDbUser(dbUser)
	user, err := json.Marshal(mappedUser)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(user)
}

func mapDbUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
	}
}
