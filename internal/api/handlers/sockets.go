package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/har-sat/termchat/internal/database"
	"github.com/har-sat/termchat/utils"
)

type SocketHandler struct {
	db *database.Queries
	Upgrader websocket.Upgrader
}

func NewSocketHandler(db *database.Queries, Upgrader websocket.Upgrader) *SocketHandler {
	return &SocketHandler{
		db: db,
		Upgrader: Upgrader,
	}
}

func (h *SocketHandler) ConnectToRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := h.Upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error upgrading connection: %v\n", err))
		return
	}
	
	defer conn.Close()

}
