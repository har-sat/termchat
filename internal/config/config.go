package config

import (
	"database/sql"

	"github.com/har-sat/termchat/internal/api/ws"
	"github.com/har-sat/termchat/internal/database"
	_ "github.com/lib/pq"
)

type Config struct {
	DB       *database.Queries
	Hub 	 *ws.Hub
	Env      *Env
}

func CreateConfig() (*Config, error) {
	env, err := LoadEnv()
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open("postgres", env.DB_URL)
	if err != nil {
		return nil, err
	}
	db := database.New(conn)

	hub := ws.CreateHub()
	cfg := Config{
		DB: db,
		Env: env,
		Hub: hub,
	}
	return  &cfg, nil
}
