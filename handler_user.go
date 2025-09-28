package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/termchat/internal/server/database"
	"github.com/har-sat/termchat/internal/server/models"
)

func (cfg *Config) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	parameters := params{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error decoding request params: %v\n", err))
		return
	}

	hashedPassword, err := HashPassword(parameters.Password)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error hashing password: %v\n", err))
		return
	}

	usr, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Username:  parameters.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	respondWithJSON(w, 201, models.DatabaseUserToUser(&usr))
}

func (cfg *Config) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	parameters := params{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error decoding request params: %v\n", err))
		return
	}

	user, err := cfg.DB.FindUserByUsername(r.Context(), parameters.Username)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("username %v doesn't exist\n", parameters.Username))
		return
	}

	if err := CompareWithHashedPassword(parameters.Password, user.Password); err != nil {
		respondWithError(w, 400, "Unauthorized - passwords don't match")
		return
	}

	respondWithJSON(w, 200, models.DatabaseUserToUser(&user))
}
