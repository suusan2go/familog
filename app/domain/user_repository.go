package domain

// UserRepository persist or find user from repository
type UserRepository interface {
	Save(*User) error
}
