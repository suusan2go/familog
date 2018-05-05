package domain

type UserRepository interface {
	Save(u *User) error
	FindByIds(ids []UserID) ([]User, error)
}
