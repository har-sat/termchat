package main

import "net/http"

func readinessCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, "let's get it")
}

func errorReadinessCheck(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "oops!")
}