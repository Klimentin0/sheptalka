package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Klimentin0/sheptalka/internal/auth"
	"github.com/Klimentin0/sheptalka/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"password"`
}

type userInputParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userInputParams := userInputParams{}
	err := decoder.Decode(&userInputParams)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	hashedUserPassword, err := auth.HashPassword(userInputParams.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sqlParams := database.CreateUserParams{
		HashedPassword: hashedUserPassword,
		Email:          userInputParams.Email,
	}

	dbUser, err := cfg.db.CreateUser(r.Context(), sqlParams)
	if err != nil {
		log.Printf("Error creating user %s", err)
	}
	jsonReturnUser := JsonReturn{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
	}
	user, err := json.Marshal(jsonReturnUser)
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
