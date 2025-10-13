package middlewares

import (
	"fmt"
	"net/http"

	"github.com/har-sat/termchat/internal/auth"
	"github.com/har-sat/termchat/internal/config"
	"github.com/har-sat/termchat/internal/database"
	"github.com/har-sat/termchat/utils"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func EnsureAuth(next authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiStr, err := auth.GetApiKey(r)
		if err != nil {
			utils.RespondWithError(w, 401, fmt.Sprintf("Auth Error: %v", err))
			return
		}
		user, err := config.Cfg.DB.GetUserByApiKey(r.Context(), apiStr)
		if err != nil {
			utils.RespondWithError(w, 401, fmt.Sprintf("Invalid User: %v", err))
			return
		}
		next(w, r, user)
	}
}
