package ws

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID
	Name      string
	Clients   map[*Client]bool
	CreatedAt time.Time
	OwnerId   uuid.UUID
	CreatorID uuid.UUID
	Join      chan *Client
	Leave     chan *Client
	Forward   chan *Message
}

func (r *Room) RunRoom() {
	for {
		select {
		case c := <-r.Join:
			r.Clients[c] = true
			log.Printf("Client %v joined room %v", c.Name, r.Name)
		case c := <-r.Leave:
			delete(r.Clients, c)
			close(c.Receive)
			err := c.Conn.Close()
			if err != nil {
				return
			}
		case m := <-r.Forward:
			for client := range r.Clients {
				client.Receive <- m
			}
		}
	}
}
