package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/har-sat/termchat/internal/api/ws"
	"github.com/har-sat/termchat/internal/database"
	"github.com/har-sat/termchat/utils"
)

type SocketHandler struct {
	db       *database.Queries
	Upgrader *websocket.Upgrader
	Hub      *ws.Hub
}

func NewSocketHandler(db *database.Queries, Upgrader *websocket.Upgrader, Hub *ws.Hub) *SocketHandler {
	return &SocketHandler{
		db:       db,
		Upgrader: Upgrader,
		Hub:      Hub,
	}
}

func (h *SocketHandler) ConnectToRoom(w http.ResponseWriter, r *http.Request, user database.User) {
	params := r.URL.Query()
	strID := params.Get("roomID")
	if strID == "" {
		utils.RespondWithError(w, 400, "invalid room ID (can't be \"\")")
		return
	}

	roomID, err := uuid.Parse(strID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error parsing roomID: %v", err))
		return
	}

	room, ok := h.Hub.Rooms[roomID]
	if !ok {
		utils.RespondWithError(w, 400, "Room is not active or doesn't exist")
		return
	}

	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error upgrading connection: %v\n", err))
		return
	}

	client := ws.Client{
		ID:      user.ID,
		Name:    user.Username,
		Conn:    conn,
		Receive: make(chan *ws.Message),
		Room:    room,
	}
	room.Join <- &client

	go client.Read()
	go client.Write()

	defer conn.Close()

}
