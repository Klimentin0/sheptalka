package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Klimentin0/sheptalka/internal/auth"
	"github.com/google/uuid"
)

type userLoginParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type JsonReturn struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userLoginParams := userLoginParams{}
	err := decoder.Decode(&userLoginParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	getUser, err := cfg.db.GetUser(r.Context(), userLoginParams.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	errorCheck := auth.CheckPasswordHash(userLoginParams.Password, getUser.HashedPassword)
	if errorCheck != nil {
		http.Error(w, errorCheck.Error(), http.StatusBadRequest)
		return
	}

	jsonReturnUser := JsonReturn{
		ID:        getUser.ID,
		CreatedAt: getUser.CreatedAt,
		UpdatedAt: getUser.UpdatedAt,
		Email:     getUser.Email,
	}
	dat, err := json.Marshal(jsonReturnUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
