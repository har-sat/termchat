package main

import (
	"database/sql"
	"time"

	"github.com/gorilla/websocket"
	"github.com/har-sat/termchat/internal/server/database"
	_ "github.com/lib/pq"
)

type Config struct {
	DB       *database.Queries
	Upgrader websocket.Upgrader
	Clients  map[*websocket.Conn]bool
	Env      *Env
}

func CreateConfig() (*Config, error) {
	env, err := LoadEnv()
	if err != nil {
		return nil, err
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 10 * time.Second,
	}

	conn, err := sql.Open("postgres", env.DB_URL)
	if err != nil {
		return nil, err
	}

	db := database.New(conn)

	return &Config{
		DB:       db,
		Upgrader: upgrader,
		Clients:  make(map[*websocket.Conn]bool),
		Env:      env,
	}, nil
}
