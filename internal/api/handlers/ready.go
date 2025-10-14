package handlers

import (
	"net/http"

	"github.com/har-sat/termchat/utils"
)

func ReadinessCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, "let's get it")
}

func ErrorReadinessCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 500, "error check")
}