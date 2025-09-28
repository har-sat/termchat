package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
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
