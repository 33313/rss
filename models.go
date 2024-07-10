package main

import (
	"time"

	"github.com/myshkovsky/rss/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func modelUser(user database.User) User {
    return User{
    	ID:        user.ID,
    	CreatedAt: user.CreatedAt,
    	UpdatedAt: user.UpdatedAt,
    	Name:      user.Name,
    }
}
