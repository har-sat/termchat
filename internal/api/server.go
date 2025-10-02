package api

import "github.com/har-sat/termchat/internal/config"

type Server struct {
	cfg config.Config
}

func NewServer(cfg config.Config) *Server {
	return &Server{cfg: cfg}
}
