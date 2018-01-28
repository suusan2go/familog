package domain

import (
	"time"
)

// UserID value object
type UserID int64

// User User using this service
type User struct {
	ID        UserID
	UpdatedAt time.Time
	CreatedAt time.Time
}

// NewUser create new User Struct
func NewUser() User {
	return User{}
}
