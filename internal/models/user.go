package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/termchat/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string	`json:"username"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	ApiKey	 string		`json:"api_key"`
}


func DatabaseUserToUser(usr *database.User) User {
	return User{
		ID: usr.ID,
		Username: usr.Username,
		ApiKey: usr.ApiKey,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}
}

