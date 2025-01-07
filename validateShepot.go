package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Shepot struct {
	Body string `json:"body"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

type ValidResp struct {
	Valid bool `json:"valid"`
}

const maxShepotLgth = 140

func validateShepotHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	shepot := Shepot{}
	err := decoder.Decode(&shepot)
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

	if len(shepot.Body) > maxShepotLgth {
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

	respBody := ValidResp{
		Valid: true,
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
