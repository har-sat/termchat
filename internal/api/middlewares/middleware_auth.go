package middlewares

import (
	"fmt"
	"net/http"

	"github.com/har-sat/termchat/internal/auth"
	"github.com/har-sat/termchat/internal/database"
	"github.com/har-sat/termchat/utils"
)

type AuthMiddleware struct {
	db *database.Queries
}

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func NewAuthMiddleware(db *database.Queries) *AuthMiddleware {
	return &AuthMiddleware{db: db}
}
func (m *AuthMiddleware) EnsureAuth(next authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiStr, err := auth.GetApiKey(r)
		if err != nil {
			utils.RespondWithError(w, 401, fmt.Sprintf("Auth Error: %v", err))
			return
		}
		user, err := m.db.GetUserByApiKey(r.Context(), apiStr)
		if err != nil {
			utils.RespondWithError(w, 401, fmt.Sprintf("Invalid User: %v", err))
			return
		}
		next(w, r, user)
	}
}
