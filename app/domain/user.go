package domain

// UserID value object
type UserID int64

// User User using this service
type User struct {
	ID UserID
}

// NewUser create new User Struct
func NewUser() User {
	return User{}
}
