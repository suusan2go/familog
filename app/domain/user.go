package domain

import (
	"time"
)

// UserID value object
type UserID string

// User User using this service
type User struct {
	ID        UserID
	Name      string
	ImageUrl  string
	UpdatedAt time.Time
	CreatedAt time.Time
}

// NewUser create new User struct
func NewUser() User {
	return User{}
}
