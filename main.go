package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	cfg, err := CreateConfig()
	if err != nil {
		log.Fatalf("error creating config: %v\n", err)
	}

	router := chi.NewRouter()
	router.Get("/ready", cfg.readinessCheck)
	router.Get("/err", cfg.errorReadinessCheck)
	router.Get("/upgrade", cfg.HandlerUpgradeConnection)

	router.Post("/users", cfg.HandlerCreateUser)
	router.Post("/login", cfg.HandlerLogin)

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
