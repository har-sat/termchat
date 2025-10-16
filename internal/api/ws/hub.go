package ws

import (
	"github.com/gorilla/websocket"
	"github.com/har-sat/termchat/internal/database"
)

type Hub struct {
	Upgrader websocket.Upgrader
	Rooms map[string]*Room
}

func CreateHub() *Hub {
	return &Hub{
		Upgrader: websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
		},
		Rooms: make(map[string]*Room),
	}
}

func (hub *Hub) CreateRoom(room *database.Room) bool {
	ID := room.ID.String()
	_, ok := hub.Rooms[ID]
	if ok {
		return false
	}

	hub.Rooms[ID] = &Room{
		ID: room.ID,
		CreatedAt: room.CreatedAt,
		OwnerId: room.OwnerID,
		CreatorID: room.CreatorID,
		Clients: make(map[*Client]bool),
	}
	// go room.RunRoom()
	return true
}