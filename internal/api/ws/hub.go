package ws

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/har-sat/termchat/internal/database"
)

type Hub struct {
	Upgrader websocket.Upgrader
	Rooms map[uuid.UUID]*Room
}

func CreateHub() *Hub {
	return &Hub{
		Upgrader: websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
		},
		Rooms: make(map[uuid.UUID]*Room),
	}
}

func (hub *Hub) CreateRoom(r *database.Room) bool {
	_, ok := hub.Rooms[r.ID]
	if ok {
		return false
	}

	room := &Room{
		ID: r.ID,
		CreatedAt: r.CreatedAt,
		OwnerId: r.OwnerID,
		CreatorID: r.CreatorID,
		Clients: make(map[*Client]bool),
		Join: make(chan *Client),
		Leave: make(chan *Client),
		Forward: make(chan []byte),
	}
	hub.Rooms[r.ID] = room
	
	go room.RunRoom()
	return true
}