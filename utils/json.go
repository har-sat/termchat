package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(500)
		log.Println("error marshalling object: ", err)
	}

	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Printf("server error: %v\n", msg)
	}

	type Err struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, Err{Error: msg})
}