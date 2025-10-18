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
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatalf("error creating config: %v\n", err)
	}
	
	router := chi.NewRouter()
	authMiddleware := middlewares.NewAuthMiddleware(cfg.DB)

	router.Get("/ready", handlers.ReadinessCheck)
	router.Get("/err", handlers.ErrorReadinessCheck)

	// router.Get("/upgrade", handlers.HandlerUpgradeConnection)

	userHanlder := handlers.NewUserHandler(cfg.DB)
	router.Post("/users", userHanlder.CreateUser)
	router.Get("/login", userHanlder.Login)

	roomsHandler := handlers.NewRoomsHandler(cfg.DB, cfg.Hub)
	router.Post("/rooms", authMiddleware.EnsureAuth(roomsHandler.CreateRoom))
	router.Get("/rooms", authMiddleware.EnsureAuth(roomsHandler.GetAllRooms))

	socketHandler := handlers.NewSocketHandler(cfg.DB, &cfg.Hub.Upgrader, cfg.Hub)
	router.Get("/ws/{id}",authMiddleware.EnsureAuth(socketHandler.ConnectToRoom))
	
	server := http.Server{
		Addr:    ":" + cfg.Env.PORT,
		Handler: router,
	}

	fmt.Printf("Starting server on PORT: %v\n", cfg.Env.PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
