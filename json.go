package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(500)
		log.Println("error marshalling object: ", err)
	}

	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg error) {
	if code >= 500 {
		log.Printf("server error: %v\n", msg)
	}

	type errMsg struct {
		Error string	`json:"error"`
	}
	respondWithJSON(w, code, errMsg{Error: msg.Error()})
}