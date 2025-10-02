package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT   string
	DB_URL string
}

func LoadEnv() (*Env, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT not found in env")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		return nil, errors.New("DB_URL not found in env")
	}

	return &Env{
		PORT:   port,
		DB_URL: dbUrl,
	}, nil
}
