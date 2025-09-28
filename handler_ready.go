package main

import (
	"net/http"
)

func (cfg *Config) readinessCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "let's get it")
}

func (cfg *Config) errorReadinessCheck(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "error check")
}