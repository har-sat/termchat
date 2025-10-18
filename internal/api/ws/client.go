package ws

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID      uuid.UUID
	Name    string
	Conn    *websocket.Conn
	Receive chan []byte
	Room    *Room
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
		_, msg, err := c.Conn.ReadMessage()
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
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			if c.Room != nil {
				c.Room.Leave <- c
			}
			return
		}
	}
}
