package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/termchat/internal/api/ws"
	"github.com/har-sat/termchat/internal/database"
	"github.com/har-sat/termchat/utils"
)

type RoomsHandler struct {
	db *database.Queries
	hub *ws.Hub
}

func NewRoomsHandler(db *database.Queries, hub *ws.Hub) *RoomsHandler {
	return &RoomsHandler{db: db, hub: hub}
}

func (h *RoomsHandler) CreateRoom(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
	}
	var params parameters
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	room, err := h.db.CreateRoom(r.Context(), database.CreateRoomParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		CreatorID: user.ID,
		OwnerID:   user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("error creating room: %v", err))
		return
	}


	h.hub.CreateRoom(&room)
	utils.RespondWithJSON(w, 201, room)
}

func (h *RoomsHandler) GetAllRooms(w http.ResponseWriter, r *http.Request, _ database.User) {

	rooms, err := h.db.FindAllRooms(r.Context())
	if err != nil {
		utils.RespondWithJSON(w, 400, fmt.Sprintf("error getting rooms: %v", err))
		return
	}
	utils.RespondWithJSON(w, 200, rooms)
}
