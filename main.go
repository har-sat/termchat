package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

// type Config struct {
// 	DB *database.Queries
// }

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

	_, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v\n", err)
	}

	router := chi.NewRouter()
	router.Get("/ready", readinessCheck)
	router.Get("/err", errorReadinessCheck)

	server := http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	fmt.Println("Starting server on PORT: ", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}