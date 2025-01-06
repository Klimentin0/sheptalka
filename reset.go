package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) resetCountHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Number of hits %d", cfg.fileserverHits.Load())))
}
