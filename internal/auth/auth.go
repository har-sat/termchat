package auth

import (
	"errors"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPwd), nil
}

func CompareWithHashedPassword(pwd, hashedPwd string) (error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}

func GetApiKey(r *http.Request) (string, error) {
	apiStr := r.Header.Get("Authorization")
	if apiStr == "" {
		return "", errors.New("No authorisation header found")
	}
	arr := strings.Split(apiStr, " ")
	if arr[0] != "Bearer" || len(arr) != 2 {
		return "", errors.New("invalid authorisation key")
	}
	return arr[1],nil
}