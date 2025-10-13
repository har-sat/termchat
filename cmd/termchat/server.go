package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/har-sat/termchat/internal/api/handlers"
	"github.com/har-sat/termchat/internal/api/middlewares"
	"github.com/har-sat/termchat/internal/config"
)

func main() {
	err := config.CreateConfig()
	if err != nil {
		log.Fatalf("error creating config: %v\n", err)
	}
	
	router := chi.NewRouter()
	router.Get("/ready", handlers.ReadinessCheck)
	router.Get("/err", handlers.ErrorReadinessCheck)
	router.Get("/upgrade", handlers.HandlerUpgradeConnection)

	router.Post("/users", handlers.HandlerCreateUser)
	router.Get("/login", handlers.HandlerLogin)

	router.Get("/rooms", middlewares.EnsureAuth(handlers.HanlderGetAllRooms))
	
	server := http.Server{
		Addr:    ":" + config.Cfg.Env.PORT,
		Handler: router,
	}

	fmt.Printf("Starting server on PORT: %v\n", config.Cfg.Env.PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
