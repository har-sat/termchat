package handlers

import (
	"net/http"

	"github.com/har-sat/termchat/internal/api"
	"github.com/har-sat/termchat/utils"
)

func (cfg *api.Server) readinessCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, "let's get it")
}

func (cfg *Server) errorReadinessCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 500, "error check")
}