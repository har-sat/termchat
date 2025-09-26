package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/har-sat/termchat/internal/server/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DB *database.Queries
	Upgrader websocket.Upgrader

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error loading env vars: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("port not mentioned in env")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalf("DB_URL not found")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v\n", err)
	}

	cfg := Config{
		DB: database.New(conn),
		Upgrader: websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
		},
	}
	router := chi.NewRouter()
	router.Get("/ready", cfg.readinessCheck)
	router.Get("/err", cfg.errorReadinessCheck)
	router.Get("/upgrade", cfg.HandlerUpgradeConnection)

	server := http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	fmt.Printf("Starting server on PORT: %v\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}