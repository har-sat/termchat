package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/har-sat/termchat/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, *database.User)

func GetApiKey(w http.ResponseWriter, usr database.User) (string, error) {
	apiStr := w.Header().Get("Authorization")
	if apiStr == "" {
		return "", errors.New("No authorisation header found")
	}
	arr := strings.Split(apiStr, " ")
	if arr[0] != "Bearer" || len(arr) != 2 {
		return "", errors.New("invalid authorisation key")
	}
	return arr[1],nil
}

// func EnsureAuth(next authHandler) http.HandlerFunc {
// }
