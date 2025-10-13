package handlers

import (
	"net/http"

	"github.com/har-sat/termchat/internal/database"
	"github.com/har-sat/termchat/utils"
)

func HanlderGetAllRooms(w http.ResponseWriter, r *http.Request, usr database.User) {
	utils.RespondWithJSON(w, 200, "This response is authenticated")
}

