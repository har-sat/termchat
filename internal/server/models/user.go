package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/har-sat/termchat/internal/server/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string	`json:"username"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}


func DatabaseUserToUser(usr *database.User) User {
	return User{
		ID: usr.ID,
		Username: usr.Username,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
	}
}

