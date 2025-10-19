package ws

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID      uuid.UUID
	Name    string
	Conn    *websocket.Conn
	Receive chan *Message
	Room    *Room
}

type Message struct {
	Id         uuid.UUID `json:"id"`
	Content    string    `json:"content"`
	SenderId   uuid.UUID `json:"sender_id"`
	SenderName string    `json:"sender_name"`
	CreatedAt  time.Time `json:"created_at"`
	RoomId     uuid.UUID `json:"room_id"`
}

/*
Read reads from clients own socket and sends it to everyone else
While write sees what has been received in its socket and writes into the clients own socket

if there are two clients A, B and A sends a message to B
A types, sends it in his socket, the Read() reads from it, forwards it to the room.
the room forwards it to everyones receive channel
B gets it and writes to its socket
*/
func (c *Client) Read() {
	defer c.Conn.Close()
	for {
		var msg *Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading: %v", err)
			if c.Room != nil {
				c.Room.Leave <- c
			}
			return

		}
		c.Room.Forward <- msg
	}
}

func (c *Client) Write() {
	defer c.Conn.Close()
	for msg := range c.Receive {
		err := c.Conn.WriteJSON(&msg)
		if err != nil {
			if c.Room != nil {
				c.Room.Leave <- c
			}
			return
		}
	}
}
