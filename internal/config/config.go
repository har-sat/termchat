package config

import (
	"database/sql"
	"time"

	"github.com/gorilla/websocket"
	"github.com/har-sat/termchat/internal/database"
	_ "github.com/lib/pq"
)

type Config struct {
	DB       *database.Queries
	Upgrader websocket.Upgrader
	Clients  map[*websocket.Conn]bool
	Env      *Env
}

var Cfg Config

// instead of this returning a config object, set the Cfg object that is global to the package,
// this way, all handlers don't have to be methods on the config struct
func CreateConfig() (error) {
	env, err := LoadEnv()
	if err != nil {
		return err
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 10 * time.Second,
	}

	conn, err := sql.Open("postgres", env.DB_URL)
	if err != nil {
		return err
	}

	db := database.New(conn)

	Cfg = Config{
		DB: db,
		Upgrader: upgrader,
		Clients: make(map[*websocket.Conn]bool),
		Env: env,
	}
	return  nil
}
