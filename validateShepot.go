package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Shepot struct {
	Body string `json:"body"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

type ValidResp struct {
	Cleaned string `json:"cleaned_body"`
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
	cleanedBody := badWordCleaner(shepot.Body)
	respBody := ValidResp{
		Cleaned: cleanedBody,
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
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
